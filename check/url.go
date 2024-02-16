// 遷移チェック
package check

import (
	"context"

	"github.com/chromedp/chromedp"
	"golang.org/x/xerrors"
	//"github.com/rayfiyo/autoALCv3/debug"
)

func URL(ctx context.Context, expectedURL string) error {
	// getURL := "https://dummy"
	var getURL string
	if err := chromedp.Run(ctx, chromedp.Location(&getURL)); err != nil {
		return xerrors.Errorf("Transition check fails on URL: %w", err)
	}

	if expectedURL == getURL {
		return nil
	} else {
		return xerrors.Errorf("Transition check fails: Getting text is %s", getURL)
	}
}
