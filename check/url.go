// 遷移チェック
package check

import (
	"context"
	"errors"

	"github.com/chromedp/chromedp"
	//"github.com/rayfiyo/autoALCv3/debug"
)

func URL(ctx context.Context, expectedURL string) error {
	// getURL := "https://dummy"
	var getURL string
	if err := chromedp.Run(ctx, chromedp.Location(&getURL)); err != nil {
		return errors.New("Transition check fails on URL\n")
	}

	if expectedURL == getURL {
		return nil
	} else {
		return errors.New("Transition check fails\n" + "get text: " + getURL + "\n")
	}
}
