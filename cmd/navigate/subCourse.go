// 問題選択のうち，サブコースまでの遷移（ナビゲーション）
package navigate

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"golang.org/x/xerrors"
	//"github.com/rayfiyo/autoALCv3/cmd/debug"
)

func SubCourse(ctx context.Context, subCourse int) error {
	log.Printf("Start navigating subcourse\n")

	// サブコースの選択
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@onclick="javascript: GoToStUnitList_Click('PWH_L0`+fmt.Sprint(subCourse)+`')"]`, chromedp.NodeVisible),
	); err != nil {
		return xerrors.Errorf("Failed to select subcourse: %w", err)
	}
	time.Sleep(1 * time.Second)

	// ユニット数を取得
	unitNum, err := nodeCount(ctx,
		`//*[@class="label label-success"]`, chromedp.AtLeast(7),
	)
	if err != nil {
		return err
	}

	// リンクがある行数（ノード数）を取得
	nodeNum, err := nodeCount(ctx,
		`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr`, chromedp.AtLeast(7),
	)
	if err != nil {
		return err
	}

	// ユニットの選択と処理
	if unitNum != nodeNum { // PowerWords Hybridコース 用
		for i := 1; i < nodeNum; i++ {
			linkText := "txt"
			if err := chromedp.Run(ctx,
				// chromedp.Text(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(i)+`]/td[3]`, &linkText),
				chromedp.TextContent(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(i)+`]/td[3]`, &linkText),
			); err != nil {
				return xerrors.Errorf("Failed to filter input: %w", err)
			}
			if !strings.Contains(linkText, "Input") {
				// ここから書く１
				// Input以外のユニットの選択と処理
				log.Print(i, linkText, "\n")
			}
		}
	} else { // TOEIC(R) L&R テスト 500点突破コース 用
		if err := chromedp.Run(ctx,
			// ここから書く２
			// ユニットの選択と処理
			chromedp.Click(``, chromedp.NodeVisible),
		); err != nil {
			return xerrors.Errorf("Failed to select units: %w", err)
		}
	}

	fmt.Printf("unitNum: %d\n", unitNum)
	fmt.Printf("nodeNum: %d\n", nodeNum)

	time.Sleep(10 * time.Second)
	log.Printf("Finish navigating subcourse\n\n")
	return nil
}

func nodeCount(ctx context.Context, sel interface{}, opts ...chromedp.QueryOption) (int, error) {
	nodes := []*cdp.Node{}
	err := chromedp.Run(ctx, chromedp.Nodes(sel, &nodes, opts...))
	if err != nil {
		return -1, xerrors.Errorf("Failed to get nodes count: %w", err)
	}
	return len(nodes), nil
}
