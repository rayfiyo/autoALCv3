// 遷移チェック
package check

import (
	"context"
	"errors"

	"github.com/chromedp/chromedp"
	//"github.com/rayfiyo/autoALCv3/debug"
)

func Content(ctx context.Context, xpath string, expectedText string) error {
	getText := "dummy"
	if err := chromedp.Run(ctx, chromedp.Text(xpath, &getText)); err == nil {
		return errors.New("Transition check fails on Content\n")
	}

	if expectedText == getText {
		return nil
	} else {
		return errors.New("Transition check fails\n" + "get text: " + getText + "\n")
	}
}
