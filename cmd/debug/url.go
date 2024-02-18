package debug

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
	"golang.org/x/xerrors"
)

func URL(ctx context.Context) (string, error) {
	var url string
	if err := chromedp.Run(ctx,
		chromedp.Location(&url),
	); err != nil {
		return url, xerrors.Errorf("Failed to location URL: %w", err)
	}

	return fmt.Sprintf("DebugURL: %s\n", url), nil
}
