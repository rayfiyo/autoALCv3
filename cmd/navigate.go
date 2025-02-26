// コースまでの遷移 (navigate)

package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"golang.org/x/xerrors"
)

func pwh(ctx context.Context, subcID int) error {
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="DivAllSubCourseTable"]/table/tbody/tr/td[2]/table/tbody/tr[`+fmt.Sprint(subcID)+`]/td[1]/a`,
			chromedp.NodeVisible),
	); err != nil {
		return xerrors.Errorf("Failed to select subcourse pwh: %w", err)
	}
	return nil
}

func tc1(ctx context.Context, subcID int) error {
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="DivAllSubCourseTable"]/table/tbody/tr/td[2]/table/tbody/tr[`+fmt.Sprint(subcID)+`]/td[1]/a`,
			chromedp.NodeVisible),
	); err != nil {
		return xerrors.Errorf("Failed to select subcourse tc1: %w", err)
	}
	return nil
}

func Navigate(ctx context.Context, crsID, subcID int) error {
	log.Println("Start  navigation")

	// コースの選択
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="LbtSubCourseLink_`+fmt.Sprint(crsID)+`"]`, chromedp.NodeVisible),
	); err != nil {
		return xerrors.Errorf("Failed to select course: %w", err)
	}
	time.Sleep(600 * time.Millisecond)

	// サブコースの選択
	switch crsID {
	case 1:
		// PowerWords Hybridコース
		if err := pwh(ctx, subcID); err != nil {
			return err
		}
	case 2:
		// TOEIC(R) L&R テスト 500点突破コース
		if err := tc1(ctx, subcID); err != nil {
			return err
		}
	}

	time.Sleep(2 * time.Second) // 遷移待ち

	log.Println("Finish navigation")
	return nil
}
