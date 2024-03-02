// ユニットを解く（タスクする）

package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"golang.org/x/xerrors"
)

// UId: TC1_S1_U004-1
func GetUId(ctx context.Context, unitNum, nodeNum int) (string, error) {
	log.Printf("Start tasks\n")
	time.Sleep(1 * time.Second) // 読み込み待ち

	// ユニットの選択と処理
	for i := 1; i < nodeNum; i++ {

		// ユニットの仕分けの為に，３列目のリンクテキストを取得
		linkText := "txt"
		if err := chromedp.Run(ctx,
			chromedp.TextContent(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(i)+`]/td[3]`, &linkText),
		); err != nil {
			return "", xerrors.Errorf("Failed to filter input: %w", err)
		}
		linkText = strings.TrimSpace(linkText)

		// 修了済みか確認する為に．ステータスの文字列を取得
		status := "txt"
		if err := chromedp.Run(ctx,
			chromedp.TextContent(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(i)+`]/td[2]`, &status),
		); err != nil {
			return "", xerrors.Errorf("Failed to filter input: %w", err)
		}
		status = strings.TrimSpace(status)

		// ユニット毎の選択と処理
		if linkText == "" { // 実力テスト・確認テスト
			// やらなくて良さげなので，何もしない
			// chromedp.Click(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(i)+`]/td[2]/span[2]/a`,
		} else if linkText == "-" { // ドリル
			if status != "修了 / Completed" { // 修了済みではないなら解く
				// ShowLearnPage() の引数より UId を取得
				var nodes []*cdp.Node
				if err := chromedp.Run(ctx,
					chromedp.Nodes(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(i)+`]/td[1]/a`, &nodes),
				); err != nil {
					return "", xerrors.Errorf("Failed to click on drill unit: %w", err)
				}

				for _, n := range nodes {
					log.Println(string([]rune(n.AttributeValue("onclick"))[15:29]))
				}
			}
		} else if linkText == "インプット / Input" {
			// インプット はしなくて良い
		} else {
			return "", xerrors.Errorf("Unexpected Units: %s", linkText)
		}
	}

	time.Sleep(12 * time.Minute)
	log.Printf("Finish tasks\n\n")
	return "UId: hogehoge", nil
}
