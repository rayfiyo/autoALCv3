// 問題選択までの遷移（ナビゲーション）
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	//"github.com/rayfiyo/autoALCv3/debug"
)

func Navigate(ctx context.Context, course int, subcourse int) {
	log.Printf("Start Navigate\n")

	// コースの選択
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="LbtSubCourseLink_`+fmt.Sprint(course)+`"]`, chromedp.NodeVisible),
	); err != nil {
		log.Fatal("Failed to select course")
	}
	time.Sleep(100 * time.Millisecond)

	// onclick="ShowLearnPage('PWH_L03_U001-2','U001','10', '','UNIT001', '2', '&STCnt1=&STCnt2=&STCnt3=&STCnt4=')"
	//*[@id="nan-contents"]/div[7]/div/table/tbody/tr[214]/td[1]/a
	for i := 1; i < subcourse+1; i++ {

		// サブコースの選択
		if err := chromedp.Run(ctx,
			chromedp.Click(`//*[@onclick="javascript: GoToStUnitList_Click('PWH_L0`+fmt.Sprint(i)+`')"]`, chromedp.NodeVisible),
		); err != nil {
			log.Fatal("Failed to select subcourse")
		}
		time.Sleep(1 * time.Second)

		// ユニット数の回収
		var nodes []*cdp.Node
		if err := chromedp.Run(ctx,
			chromedp.Nodes(`//*[@id="nan-contents"]/div[7]/div/table/tbody/tr`, &nodes, chromedp.AtLeast(1)),
		); err != nil {
			log.Fatal("Failed to select subcourse")
		}
		units := 0
		if course == 1 {
			units = len(nodes) - 100
		} else {
			units = len(nodes)
		}
		fmt.Println(units)

		// ユニットの選択
	}

	time.Sleep(10 * time.Second)
	log.Printf("Finish Navigate\n\n")
}
