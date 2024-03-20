package generator

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"sync"
	"time"
)

type ProcessorInterface interface {
	Process(payload string) ([]byte, error)
}

type HTMLProcessor struct {
	chromiumURL string
}

func NewHTMLProcessor(url string) ProcessorInterface {
	return HTMLProcessor{url}
}

func (hp HTMLProcessor) Process(payload string) ([]byte, error) {
	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), hp.chromiumURL)
	defer cancel()

	// log the CDP messages so that you can find the one to use.
	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	start := time.Now()
	var buf []byte

	err := chromedp.Run(ctx,
		// the navigation will trigger the "page.EventLoadEventFired" event too,
		// so we should add the listener after the navigation.
		chromedp.Navigate("about:blank"),
		// set the page content and wait until the page is loaded (including its resources).
		chromedp.ActionFunc(func(ctx context.Context) error {
			lctx, cancel := context.WithCancel(ctx)
			chromedp.ListenTarget(lctx, func(e interface{}) {
				switch ev := e.(type) {
				case *network.EventResponseReceived:
					if !strings.HasPrefix(ev.Response.URL, "data:") {
						log.Printf("Cache %s: %v", ev.Response.URL, ev.Response.FromDiskCache)
					}
				case *page.EventLoadEventFired:
					// It's a good habit to remove the event listener if we don't need it anymore.
					cancel()
					wg.Done()
				}
			})
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			return page.SetDocumentContent(frameTree.Frame.ID, payload).Do(ctx)
		}),
		// wait for the page.EventLoadEventFired
		chromedp.ActionFunc(func(ctx context.Context) error {
			wg.Wait()
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
