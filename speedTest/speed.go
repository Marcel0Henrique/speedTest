package speedtest

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

// Run the service
func Run() {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://fast.com/pt/`),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`body > footer`),
		// find and click "Example" link
		//chromedp.Click(`.result-label`, chromedp.NodeVisible),
		// retrieve the text of the textarea
		chromedp.Value(`#your-speed-message`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}
