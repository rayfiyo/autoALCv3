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
	"golang.org/x/xerrors"
)

func passwdInputer(labelMessage string) (string, error) {
	// validate := func(input string) error {
	// return nil
	// }

	prompt := promptui.Prompt{
		Label: labelMessage,
		// Validate: validate,
		Mask: '*',
	}

	passwd, err := prompt.Run()
	if err != nil {
		return passwd, xerrors.Errorf("Failed to run prompt: %w", err)
	}

	return passwd, nil
}

func Login(ctx context.Context) error {
	log.Printf("Start Login\n")

	// クレデンシャルの入力要求
	id := "ID"
	fmt.Print("ID: ")
	fmt.Scan(&id)

	if passwd, err := passwdInputer("Password: "); err != nil {
		return err
	} else {
		// アカウントとパスワードを入力
		if err := chromedp.Run(ctx,
			chromedp.Click(`//*[@id="AccountId"]`, chromedp.NodeVisible),
			input.InsertText(id),
			chromedp.Click(`//*[@id="Password"]`, chromedp.NodeVisible),
			input.InsertText(passwd),
			chromedp.Click(`//*[@id="BtnLogin"]`, chromedp.NodeVisible),
		); err != nil {
			return xerrors.Errorf("Login process failed: %w\n", err)
		}
		time.Sleep(1 * time.Second)
	}

	log.Printf("Finish Login\n\n")
	return nil
}
