// 問題選択のうち，サブコースまでの遷移（ナビゲーション）
package navigate

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"golang.org/x/xerrors"
)

func SubCourse(ctx context.Context, subCourse int) error {
	log.Printf("Start navigating subcourse\n")

	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@onclick="javascript: GoToStUnitList_Click('PWH_L0`+fmt.Sprint(subCourse)+`')"]`,
			chromedp.NodeVisible),
	); err != nil {
		return xerrors.Errorf("Failed to select subcourse: %w", err)
	}
	time.Sleep(1 * time.Second) // もっと待っても良いかも

	log.Printf("Finish navigating subcourse\n\n")
	return nil
}
