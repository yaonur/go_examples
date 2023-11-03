package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Load local HTML file
	htmlContent, err := ioutil.ReadFile("test.html")
	if err != nil {
		fmt.Println("Error reading HTML file:", err)
		return
	}

	var pdfBuf []byte

	err = chromedp.Run(ctx, convertHTMLToPDF(string(htmlContent), &pdfBuf))
	if err != nil {
		fmt.Println("Error converting HTML to PDF:", err)
		return
	}

	err = ioutil.WriteFile("output.pdf", pdfBuf, 0644)
	if err != nil {
		fmt.Println("Error writing PDF file:", err)
		return
	}

	fmt.Println("Converted HTML to PDF and saved as output.pdf")
}

func convertHTMLToPDF(html string, pdfBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate("data:text/html," + html),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().Do(ctx)
			if err != nil {
				return err
			}
			*pdfBuf = buf
			return nil
		}),
	}
}
