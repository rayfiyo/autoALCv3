package debug

import (
	"context"
	"os"

	"github.com/chromedp/chromedp"
	"golang.org/x/xerrors"
)

func Pic(ctx context.Context) {
	var buf []byte

	// Get Screen
	if err := chromedp.Run(ctx,
		chromedp.FullScreenshot(&buf, 90),
	); err != nil {
		xerrors.Errorf("Failed to capture a screenshot: %w", err)
	}

	// Save Screen
	if err := os.WriteFile(
		"fullScreenshot.png", buf, 0o644,
	); err != nil {
		xerrors.Errorf("Failed to output the screenshot: %w", err)
	}
}
