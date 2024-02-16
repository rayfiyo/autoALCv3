// 問題選択のうち，コースまでの遷移（ナビゲーション）
package navigate

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"golang.org/x/xerrors"
)

func MainCourse(ctx context.Context, course int) error {
	log.Printf("Start navigating course\n")

	// コースの選択
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="LbtSubCourseLink_`+fmt.Sprint(course)+`"]`, chromedp.NodeVisible),
	); err != nil {
		return xerrors.Errorf("Failed to select course: %w", err)
	}
	time.Sleep(100 * time.Millisecond)

	log.Printf("Finish navigating course\n\n")
	return nil
}
