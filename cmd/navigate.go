// 問題選択までの遷移（ナビゲーション）
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"golang.org/x/xerrors"
	//"github.com/rayfiyo/autoALCv3/debug"
)

func Navigate(ctx context.Context, course int, subcourse int) error {
	log.Printf("Start Navigate\n")

	// コースの選択
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="LbtSubCourseLink_`+fmt.Sprint(course)+`"]`, chromedp.NodeVisible),
	); err != nil {
		return xerrors.Errorf("Failed to select course: %w", err)
	}
	time.Sleep(100 * time.Millisecond)

	for i := 1; i < subcourse+1; i++ {
		nodes := []*cdp.Node{}

		// サブコースの選択
		if err := chromedp.Run(ctx,
			chromedp.Click(`//*[@onclick="javascript: GoToStUnitList_Click('PWH_L0`+fmt.Sprint(i)+`')"]`, chromedp.NodeVisible),
		); err != nil {
			return xerrors.Errorf("Failed to select subcourse: %w", err)
		}
		time.Sleep(1 * time.Second)

		// ユニット数の回収
		if err := chromedp.Run(ctx,
			chromedp.Nodes(`//*[@class="label label-success"]`, &nodes, chromedp.AtLeast(1)),
		); err != nil {
			return xerrors.Errorf("Failed to get units count: %w", err)
		}
		units := len(nodes)
		fmt.Println(units)

		// ユニットの選択
		if err := chromedp.Run(ctx, chromedp.Nodes("a", &nodes)); err != nil {
			return xerrors.Errorf("Failed to select units: %w", err)
		}
		for i := range nodes {
			log.Println(nodes[len(nodes)-i].AttributeValue("href"))
		}

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
	}

	time.Sleep(10 * time.Second)
	log.Printf("Finish Navigate\n\n")
	return nil
}
