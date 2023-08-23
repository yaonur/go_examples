package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"os"
	"postgoflt/models"
	"postgoflt/storage"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	bookModel := &models.Books{}
	err := context.BodyParser(&bookModel)
	fmt.Println("air is working")
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}
	err = r.DB.Create(&bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "request failed"})
		return err
	}
	context.Status(http.StatusCreated).JSON(&fiber.Map{"message": "request success"})
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}
	err := r.DB.Find(&bookModels).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "request failed"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "bookModels fetched successfully", "data": bookModels})
	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := &models.Books{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id cannot be empty"})
		return nil
	}
	err := r.DB.Delete(&bookModel, id).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not delete book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "book delete succesfully"})
	return nil
}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {
	id := context.Params("id")
	fmt.Println("the Id is: ", id)
	bookModel := &models.Books{}
	err := r.DB.First(&bookModel, id).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "request failed"})
		return err
	}
	context.Status(http.StatusOK).JSON(&bookModel)
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Post("/create_book", r.CreateBook) // Use r.CreateBook directly without wrapping it.
	api.Get("/get_book/:id", r.GetBookByID)
	api.Get("/get_books", r.GetBooks)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	err = models.MigrateBoolModels(db)
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":3000")
}
