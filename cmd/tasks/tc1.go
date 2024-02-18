// TC1 (TOEIC(R) L&R テスト 500点突破コース) の処理

package tasks

import (
	"context"
	// "fmt"
	"log"
	// "strings"
	"time"

	"github.com/chromedp/chromedp"
	"golang.org/x/xerrors"
)

func TC1(ctx context.Context, unitNum, nodeNum int) error {
	log.Printf("Start tasks\n")

	if err := chromedp.Run(ctx,
		// ここから書く２
		// ユニットの選択と処理
		chromedp.Click(``, chromedp.NodeVisible),
	); err != nil {
		return xerrors.Errorf("Failed to select units: %w", err)
	}

	time.Sleep(10 * time.Second)
	log.Printf("Finish tasks\n\n")
	return nil
}
