package generator

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

type URLProcessor struct {
	chromiumURL string
}

func NewURLProcessor(url string) ProcessorInterface {
	return URLProcessor{url}
}

func (up URLProcessor) Process(payload string) ([]byte, error) {
	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), up.chromiumURL)
	defer cancel()

	// log the CDP messages so that you can find the one to use.
	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancel()

	start := time.Now()
	var buf []byte

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(payload),
		// set the page content and wait until the page is loaded (including its resources).
		chromedp.ActionFunc(func(ctx context.Context) error {
			lctx, _ := context.WithCancel(ctx)
			chromedp.ListenTarget(lctx, func(e interface{}) {
				switch ev := e.(type) {
				case *network.EventResponseReceived:
					if !strings.HasPrefix(ev.Response.URL, "data:") {
						log.Printf("Cache %s: %v", ev.Response.URL, ev.Response.FromDiskCache)
					}
				}
			})
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().
				WithPaperWidth(8.27).
				WithPaperHeight(11.69).
				WithPrintBackground(true).
				Do(ctx)
			if err != nil {
				return err
			}
			finish := time.Now()
			log.Printf("Generation time: %s", finish.Sub(start))

			return nil
		}),
	)

	return buf, err
}
