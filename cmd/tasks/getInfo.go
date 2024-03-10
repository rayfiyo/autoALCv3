// ShowLearnPageの引数から，CId，SId，UId，stcnt（ユニットのステップ数）の情報を取得

package tasks

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/rayfiyo/autoALCv3/model"
	"golang.org/x/xerrors"
)

func GetInfo(ctx context.Context, selNode int) (model.Id, int, error) {
	log.Printf("Start of getting Id\n")

	// 返り値の初期化
	id := model.Id{UId: "ex.TC1_S1_U003-1", SId: "ex.TC1_S1", CId: "ex.TC1", SessId: "barbar"}
	stCnt := -1

	// 読み込み待ち
	time.Sleep(1 * time.Second)

	// Cookie(ASP.NET_SessionId) の取得のために，それが記述してあるHTMLタグのnodeを取得
	var nodes []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes(`//*[@id="HidSessionId"]`, &nodes)); err != nil {
		log.Panic("Failed to get nodes of HidSessionId: %w", err)
	}

	// Cookie(ASP.NET_SessionId) の取得
	for _, n := range nodes {
		id.SessId = n.AttributeValue("value")
	}
	if id.SessId == "" {
		log.Panic("Failed to get session ID.")
	}

	// 修了済みか確認する為に．ステータスの文字列を取得
	status2 := "例: 修了（ステータスが２列目にある場合用）"
	status4 := "例: -（ステータスが４列目にある場合用）"
	if err := chromedp.Run(ctx,
		chromedp.TextContent(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(selNode)+`]/td[2]`, &status2),
		chromedp.TextContent(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(selNode)+`]/td[4]`, &status4),
	); err != nil {
		return id, stCnt, xerrors.Errorf("Failed to get status: %w", err)
	}
	status2 = strings.TrimSpace(status2)
	status4 = strings.TrimSpace(status4)

	// 修了済みではないなら，ShowLearnPage() の引数より Id を取得
	if status2 != "修了 / Completed" && status4 != "修了 / Completed" {
		var nodes []*cdp.Node
		if err := chromedp.Run(ctx,
			chromedp.Nodes(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(selNode)+`]//a`, &nodes),
		); err != nil {
			return id, stCnt, xerrors.Errorf("Failed to click on drill unit: %w", err)
		}

		// [note] slpArgって，マップで持った方がいいよね？

		// ex.     "PWH_L03_U024-2", "U024",   "10",           "",         "UNIT024",   "2", "&STCnt1=&STCnt2=&STCnt3=&STCnt4="}
		slpArg := []string{"unitId", "unitNo", "unitDivision", "unitType", "unitTitle", "deviceType", "unitTrainingCount"}
		for _, n := range nodes {
			// ex. "ShowLearnPage('PWH_L03_U024-2','U024','10', '','UNIT024', '2', '&STCnt1=&STCnt2=&STCnt3=&STCnt4=')"
			onclickValue := n.AttributeValue("onclick")
			if onclickValue == "" {
				return id, stCnt, xerrors.Errorf("onclickValue is empty")
			}

			// "(" の位置
			head := strings.Index(onclickValue, "(")
			if head == -1 {
				return id, stCnt, xerrors.Errorf(`There is no "(".`)
			}

			// ")" の位置
			tail := strings.LastIndex(onclickValue, ")")
			if tail == -1 {
				return id, stCnt, xerrors.Errorf(`There is no ")".`)
			}

			// ( と ) の中の文字列（つまり，ShowLearnPage の引数）だけの文字列（配列ではない）にする
			monoSlpArg := onclickValue[head+1 : tail]

			// "'" を削除
			monoSlpArg = strings.ReplaceAll(monoSlpArg, "'", "")

			// 引数に応じた配列にする（","で区切る）
			slpArg = strings.Split(monoSlpArg, ",")
		}

		// 適切に切り，空白を削除して代入
		id.UId = strings.TrimSpace(slpArg[0])
		id.SId = strings.TrimSpace(id.UId[:strings.LastIndex(id.UId, "_")])
		id.CId = strings.TrimSpace(id.SId[:strings.LastIndex(id.SId, "_")])

		// slpArg の末尾の配列
		if strings.TrimSpace(slpArg[len(slpArg)-1]) == "" {
			return id, stCnt, xerrors.Errorf("ShowLearnPage[%d] is empty", len(slpArg)-1)
		}

		// &STCnt.= の数を数える
		stCnt = strings.Count(slpArg[len(slpArg)-1], "&STCnt")

		// stCnt が１未満だったら（実力テストなど），１として処理
		if stCnt < 1 {
			stCnt = 1
		}
	} else {
		log.Println("修了済み") // id は初期値のまま
	}

	log.Printf("End of getting ID\n\n")
	return id, stCnt, nil
}
