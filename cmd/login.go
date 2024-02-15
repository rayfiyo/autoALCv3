// ログイン処理
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/chromedp"
	"github.com/manifoldco/promptui"
	//"github.com/rayfiyo/autoALCv3/debug"
)

func Login(ctx context.Context) {
	log.Printf("Start Login\n")

	// var checkText string
	// クレデンシャルの入力要求
	id := "ID"
	fmt.Printf("ID:")
	fmt.Scan(&id)
	passwd := passwdInputer("Password: ")

	// アカウントとパスワードの入力
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="AccountId"]`, chromedp.NodeVisible),
		input.InsertText(id),
		chromedp.Click(`//*[@id="Password"]`, chromedp.NodeVisible),
		input.InsertText(passwd),
		chromedp.Click(`//*[@id="BtnLogin"]`, chromedp.NodeVisible),
	); err != nil {
		log.Panic("Error: In the login process")
	}
	time.Sleep(1 * time.Second)

	log.Printf("Finish Login\n\n")
}

func passwdInputer(labelMessage string) string {
	validate := func(input string) error {
		return nil
	}

	prompt := promptui.Prompt{
		Label:    labelMessage,
		Validate: validate,
		Mask:     '*',
	}

	passwd, err := prompt.Run()
	if err != nil {
		log.Panic("Error: Failed to run prompt")
	}

	return passwd
}
