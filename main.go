package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/rayfiyo/autoALCv3/cmd"
	"github.com/rayfiyo/autoALCv3/cmd/check"
	"github.com/rayfiyo/autoALCv3/cmd/tasks"
)

func main() {
	fmt.Print("現在，付与されるスキルポイントが正しくない場合が報告されてます．詳しくは README を参照してください．\n\n")

	// Chrome のインスタンス作成（画面サイズをPCとして行う）
	// /* Release:
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("enable-automation", false),
		chromedp.WindowSize(1024, 576),
	)
	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf)) // */

	/* Debug: Log only
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf)) // */

	/* Debug: No Headless
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("enable-automation", false),
		chromedp.WindowSize(1024, 576),
	)
	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf)) // */

	defer cancel()

	// アクセス
	log.Println("Start     access to top page")
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://nanext.alcnanext.jp/anetn/Student/stlogin/index/nit-ariake/")); err != nil {
		log.Panic("Failed access to top page\n", err)
	}
	log.Println("Finish of access to top page")

	// ログイン
	if err := cmd.Login(ctx); err != nil {
		log.Panic(err)
	}
	if err := check.URL(ctx, "https://nanext.alcnanext.jp/anetn/Student/StTop"); err != nil {
		log.Panic("ID or Password is likely wrong.\n", err)
	}

	// コースの選択
	course := 0
	subcourse := -1
	fmt.Printf("コースの選択\nPowerWords Hybridコース -------------> 1\nTOEIC(R) L&R テスト 500点突破コース -> 2\n")
	fmt.Scanln(&course)
Loop:
	for i := 0; ; i++ {
		switch course {
		case 1:
			// 1~6の範囲じゃないなら繰り返す
			for subcourse < 1 || subcourse > 6 {
				fmt.Print("サブコースの選択（1~6）: ")
				fmt.Scanln(&subcourse)
			}
			break Loop
		case 2:
			// 1~5の範囲じゃないなら繰り返す
			for subcourse < 1 || subcourse > 5 {
				fmt.Print("サブコースの選択（1~5）: ")
				fmt.Scanln(&subcourse)
			}
			break Loop
		default:
			fmt.Print("1 or 2: ")
			fmt.Scanln(&course)
		}
	}

	// コース・サブコースの遷移
	if err := cmd.Navigate(ctx, course, subcourse); err != nil {
		log.Panic(err)
	}

	// 遷移の確認
	if err := check.URL(ctx, "https://nanext.alcnanext.jp/anetn/Student/StUnitList"); err != nil {
		log.Panic("Failure to select a course or sub-course\n", err)
	}

	// リンクがある行数（ノード数）を取得
	nodeNum, err := cmd.NodeCount(ctx, `//*[@id="nan-contents"]/div[7]/div/table/tbody/tr`, chromedp.AtLeast(1))
	if err != nil {
		log.Panic(err)
	}

	// ユニットの選択と処理
	for i := 1; i < nodeNum+1; i++ {
		log.Printf("- * - * - * - * -\n")
		log.Printf("%d/%d 開始\n", i, nodeNum)

		id, stcnt, err := tasks.GetInfo(ctx, i)
		if err != nil {
			log.Panic(err)
		}

		if id.CId != "ex.TC1" {
			if err := tasks.Submit(id, stcnt); err != nil {
				log.Panic(err)
			}
		}

		log.Printf("%d/%d 完了\n", i, nodeNum)
		log.Printf("- * - * - * - * -\n\n")
	}

	fmt.Println("すべての処理が完了しました．")
}
