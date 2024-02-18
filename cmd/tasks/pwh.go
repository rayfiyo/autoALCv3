// PWH (PowerWords Hybridコース) コースの処理

package tasks

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"golang.org/x/xerrors"
)

func killUnit(ctx context.Context) error {
	time.Sleep(6 * time.Second) // 読み込み待ち

	var targets []*target.Info
	var err error
	if targets, err = chromedp.Targets(ctx); err != nil {
		return xerrors.Errorf("Failed to make a new target: %w", err)
	}

	// [ToDO]forの読み込み改善
	var homeID, unitID target.ID = "", ""
	for _, t := range targets {
		if t.Type == "page" && t.URL != "about:blank" { // 拡張機能の background_page や，about:blank を除外
			if homeID == "" { // 元のページのIDを保存
				homeID = t.TargetID
			}
			if t.URL != "https://nanext.alcnanext.jp/anetn/Student/StUnitList#" { // ユニットページのIDを選択
				unitID = t.TargetID
				break
			}
		}
	}
	if homeID == "" || unitID == "" {
		return xerrors.New("Page fails to load or is too slow")
	}

	ctx, _ = chromedp.NewContext(ctx, chromedp.WithTargetID(unitID))

	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="nan-contents-cover-buttons"]/div/div[1]/button`, chromedp.NodeVisible),
	); err != nil {
		return xerrors.Errorf("Failed to click on start button: %w", err)
	}

	// [ToDo]改善
	// 本当はhomeIDに戻るべきだけど，お行儀悪く新規で開く
	// ctx, _ = chromedp.NewContext(ctx, chromedp.WithTargetID(homeID))
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://nanext.alcnanext.jp/anetn/Student/StUnitList"),
	); err != nil {
		log.Fatal("err@debugURL: Failed to location url")
	}
	return nil
}

func PWH(ctx context.Context, unitNum, nodeNum int) error {
	log.Printf("Start tasks\n")
	time.Sleep(1 * time.Second) // 読み込み待ち

	// ユニットの選択と処理
	for i := 1; i < nodeNum; i++ {

		// ユニットの仕分けの為にリンクテキストを取得
		linkText := "txt"
		if err := chromedp.Run(ctx,
			chromedp.TextContent(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(i)+`]/td[3]`, &linkText),
		); err != nil {
			return xerrors.Errorf("Failed to filter input: %w", err)
		}
		linkText = strings.TrimSpace(linkText)

		// ユニット毎の選択と処理
		if linkText == "" { // 実力テスト・確認テスト
			// Unit の選択
			if err := chromedp.Run(ctx,
				chromedp.Click(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(i)+`]/td[2]/span[2]/a`,
					chromedp.NodeVisible),
			); err != nil {
				return xerrors.Errorf("Failed to click on test unit: %w", err)
			}

			// Unit の処理
			if err := killUnit(ctx); err != nil {
				return err
			}
		} else if linkText == "-" { // ドリル
			// Unit の選択
			if err := chromedp.Run(ctx,
				chromedp.Click(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(i)+`]/td[1]/a`,
					chromedp.NodeVisible),
			); err != nil {
				return xerrors.Errorf("Failed to click on drill unit: %w", err)
			}

			// Unit の処理
			if err := killUnit(ctx); err != nil {
				return err
			}
		} else if linkText == "インプット / Input" {
			// インプット は 何もしない
		} else {
			return xerrors.Errorf("Unexpected Units: %s", linkText)
		}
	}

	time.Sleep(12 * time.Minute)
	log.Printf("Finish tasks\n\n")
	return nil
}
