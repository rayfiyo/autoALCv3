package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/rayfiyo/autoALCv3/check"
	"github.com/rayfiyo/autoALCv3/cmd"
)

func main() {
	// Chrome のインスタンス作成
	/* Release:
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("enable-automation", false),
		chromedp.WindowSize(1024, 576),
	)
	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf)) // */

	/* Debug: Log only
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf)) // */

	// /* Debug: No Headless
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("enable-automation", false),
		chromedp.WindowSize(1024, 576),
	)
	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf)) // */

	defer cancel()

	// access
	log.Println("Start access to top page")
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://nanext.alcnanext.jp/anetn/Student/stlogin/index/nit-ariake/")); err != nil {
		log.Panic("Error: Failed access to top page\n" + fmt.Sprintln(err))
	}
	fmt.Println("End of access to top page")

	// login
	if err := cmd.Login(ctx); err != nil {
		log.Panic(err)
	}
	if err := check.URL(ctx, "https://nanext.alcnanext.jp/anetn/Student/StTop"); err != nil {
		log.Panic(fmt.Sprintln(err) + "ID か パスワードが間違っている可能性があります")
	}

	// select question
	course := 0
	subcourse := 0
	fmt.Printf("コースの選択\nPowerWords Hybridコース -------------> 1\nTOEIC(R) L&R テスト 500点突破コース -> 2\n")
	fmt.Scanln(&course)
Loop:
	for i := 0; ; i++ {
		switch course {
		case 1:
			subcourse = 6
			break Loop
		case 2:
			subcourse = 5
			break Loop
		default:
			fmt.Printf("1 or 2: ")
			fmt.Scanln(&course)
		}
	}

	if err := cmd.Navigate(ctx, course, subcourse); err != nil {
		log.Panic(err)
	}

	log.Println(`cat qnaSets/?.csv > "qnaSets/逆allSecQ.csv"`)
	log.Println("All done.")
}
