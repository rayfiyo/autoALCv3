// 遷移チェック
package check

import (
	"context"

	"github.com/chromedp/chromedp"
	"golang.org/x/xerrors"
)

func URL(ctx context.Context, expectedURL string) error {
	receivedURL := "http://dummy"
	if err := chromedp.Run(ctx, chromedp.Location(&receivedURL)); err != nil {
		return xerrors.Errorf("Transition check fails on URL: %w", err)
	}

	if expectedURL == receivedURL {
		return nil
	} else {
		return xerrors.Errorf("Transition check fails: Expected URL is %s, but the one received was %s", expectedURL, receivedURL)
	}
}
