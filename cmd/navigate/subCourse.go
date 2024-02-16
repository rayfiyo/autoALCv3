// 問題選択のうち，サブコースまでの遷移（ナビゲーション）
package navigate

import (
	"context"
	"fmt"
	"log"
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
	units := 0
	if v, err := nodeCount(ctx,
		`//*[@class="label label-success"]`, chromedp.AtLeast(7),
	); err != nil {
		return err
	} else {
		units = v
	}
	fmt.Printf("units: %d\n", units)
	/*
		if err := chromedp.Run(ctx,
			chromedp.Nodes(`//*[@class="label label-success"]`, &nodes, chromedp.AtLeast(1)),
		); err != nil {
			return xerrors.Errorf("Failed to get units count: %w", err)
		}
		units := len(nodes)
	*/

	// リンクがある行数（ノード数）を取得
	nodes := 0
	// for i := 1; i < 100; i++ {
	if v, err := nodeCount(ctx,
		`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr`, chromedp.AtLeast(7),
		// `//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(i)+`]/td`, chromedp.AtLeast(7),
	); err != nil {
		return err
	} else {
		nodes = v
	}
	// }
	fmt.Printf("nodes: %d\n", nodes)

	// ユニットの選択
	/*
		for i := 1; i < units+1; i++ {
			if err := chromedp.Run(ctx,
				chromedp.Nodes(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[`+fmt.Sprint(i)+`]/td`, &nodes),
			); err != nil {
				return xerrors.Errorf("Failed to select units: %w", err)
			}
			for _, n := range nodes {
				log.Println(n.Attributes)
			}
			fmt.Println("ok")
		}
	*/

	/*
		unitID := "JT01"
		for i := 0; i < units; i++ {
			if err := chromedp.Run(ctx,
	*/
	//chromedp.Text(`//*/td[@rowspan="2"][1]`, &unitID),
	/*
			); err != nil {
				return xerrors.Errorf("Failed to select subcourse: %w", err)
			}
		}
	*/

	/*
		// a11y.QueryAXTree().Do(ctx)
		var nodeID cdp.NodeID
		fmt.Println(a11y.QueryAXTree().WithRole("link").WithNodeID(nodeID))
		fmt.Println(nodeID)


			chromedp.Evaluate(`document.querySelector('`+selector+`').getAttribute('role')`, &role),
			// var a11y *accessibility.GetPartialAXTreeParams
			a11y := accessibility.GetPartialAXTree(accessibility.GetPartialAXTree().
				WithNodeID(accessibility.GetAXNodeAndAncestors().NodeID))
			if course == 1 && unitID[:1] == "U" {
			}
			_ = a11y
			// func (p QueryAXTreeParams) WithRole(role string) *QueryAXTreeParams {
	*/

	fmt.Println("end")
	time.Sleep(10 * time.Second)
	log.Printf("Finish navigating subcourse\n\n")
	return nil
}

func nodeCount(ctx context.Context, sel interface{}, opts ...chromedp.QueryOption) (int, error) {
	nodes := []*cdp.Node{}
	if err := chromedp.Run(ctx,
		chromedp.Nodes(sel, &nodes, opts...),
	); err != nil {
		return -1, xerrors.Errorf("Failed to get nodes count: %w", err)
	}
	log.Print(nodes)
	return len(nodes), nil
}
