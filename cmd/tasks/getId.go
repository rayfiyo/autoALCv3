// ユニットを解くために必要なIdを取得

package tasks

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

type Id struct {
	CId string // ex. TC1
	SId string // ex. TC1_S1
	UId string // ex. TC1_S1_U003-1
}

func GetId(ctx context.Context, selNode int) (string, error) {
	log.Printf("Start tasks\n")
	time.Sleep(1 * time.Second) // 読み込み待ち

	// ユニットの仕分けの為に，３列目のリンクテキストを取得
	// linkText := "txt"
	// if err := chromedp.Run(ctx,
	// chromedp.TextContent(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(selNode)+`]/td[3]`, &linkText),
	// ); err != nil {
	// return "", xerrors.Errorf("Failed to filter input: %w", err)
	// }
	// linkText = strings.TrimSpace(linkText)

	// 修了済みか確認する為に．ステータスの文字列を取得
	status2 := "修了（ステータスが２列目にある場合）"
	status4 := "-（ステータスが４列目にある場合）"
	if err := chromedp.Run(ctx,
		chromedp.TextContent(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(selNode)+`]/td[2]`, &status2),
		chromedp.TextContent(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(selNode)+`]/td[4]`, &status4),
	); err != nil {
		return "", xerrors.Errorf("Failed to get status: %w", err)
	}
	status2 = strings.TrimSpace(status2)
	status4 = strings.TrimSpace(status4)

	// 修了済みではないなら，ShowLearnPage() の引数より Id を取得
	id := Id{CId: "ex.TC1", SId: "ex.TC1_S1", UId: "ex.TC1_S1_U003-1"}
	if status2 != "修了 / Completed" && status4 != "修了 / Completed" {
		var nodes []*cdp.Node
		if err := chromedp.Run(ctx,
			chromedp.Nodes(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(selNode)+`]//a`, &nodes),
		); err != nil {
			return "", xerrors.Errorf("Failed to click on drill unit: %w", err)
		}

		rawValue := "ex.PWH_L03_JT01-1','JT01"
		for _, n := range nodes {
			rawValue = string([]rune(n.AttributeValue("onclick"))[15:36]) // SIdが３桁の場合を考え少し長めに切る
		}

		for i := 0; i+1 < len(rawValue); i++ {
			if string([]rune(rawValue)[i:i+1]) == "_" {
				if id.CId == "ex.TC1" { // id.CIdが初期値なら最初の"_"
					id.CId = string([]rune(rawValue)[:i])
				} else {
					id.SId = string([]rune(rawValue)[:i])
				}
			} else if string([]rune(rawValue)[i:i+1]) == "'" {
				id.UId = string([]rune(rawValue)[:i])
			}
		}

		// log.Print("2: ", status2, ", 4: ", status4, "\n")
		// log.Printf("CId: %s, SId: %s, UId %s\n", id.CId, id.SId, id.UId)
	} else {
		log.Println("修了済み")
	}

	log.Printf("ユニット%dのタスク完了", selNode)
	// ユニット毎の選択と処理
	// if linkText == "" { // 実力テスト・確認テスト
	// log.Println("実力テスト・確認テストはスキップします")
	// やらなくて良さげなので，何もしない
	// chromedp.Click(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(i)+`]/td[2]/span[2]/a`,
	// } else if linkText == "-" { // ドリル
	// if status != "修了 / Completed" { // 修了済みではないなら解く
	// ShowLearnPage() の引数より UId を取得
	// var nodes []*cdp.Node
	// if err := chromedp.Run(ctx,
	// chromedp.Nodes(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(i)+`]/td[1]/a`, &nodes),
	// ); err != nil {
	// return "", xerrors.Errorf("Failed to click on drill unit: %w", err)
	// }

	// for _, n := range nodes {
	// log.Println(string([]rune(n.AttributeValue("onclick"))[15:29]))
	// }
	// }
	// log.Printf("ユニット%d が終了", i)
	// } else if linkText == "インプット / Input" {
	// log.Println("インプットはスキップします")
	// } else {
	// return "", xerrors.Errorf("Unexpected Units: %s", linkText)
	// }
	// }

	log.Printf("Finish tasks\n\n")
	time.Sleep(1 * time.Minute)
	return "id.UId", nil
}
