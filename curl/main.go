package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func start(sID string, subCourse, unit int) {
	// クライアント作成・リクエストヘッダの作成
	client := &http.Client{}
	data := strings.NewReader(
		`{"Qtype":"","VId":"ALC","CId":"TC1","SId":"TC1_S` + fmt.Sprint(subCourse) + `","UId":"TC1_S` + fmt.Sprint(subCourse) + `_U00` + fmt.Sprint(unit) + `-1","SessionId":"` + sID + `"}`)
	// log.Println(data)
	req, err := http.NewRequest("POST", "https://nanext.alcnanext.jp/anetn/api/HistoryApi/registStartHistory", data)
	if err != nil {
		log.Fatal(err)
	}

	// Cookieなどの追加
	req.Header.Set("Content-Type", "application/JSON")
	req.Header.Set("Cookie", "ASP.NET_SessionId="+sID)

	// 送信
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 受信
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s\n", bodyText)
	// {"Result":"0","Estep":"","SDate":"20240226231235558"}
}

func end(sID string, subCourse, unit int) {
	sdate := "2024............."
	fmt.Print("SDateを入力: ")
	fmt.Scanln(&sdate)

	// クライアント作成・リクエストヘッダの作成
	client := &http.Client{}
	data := strings.NewReader(
		`{"FId":"02","LCD":"1","LInfo":{"FID02":{"StepSection02":[{"SOrder":"1","SFlag":"1","Voca":"1"},{"SOrder":"2","SFlag":"1","Voca":"1"},{"SOrder":"3","SFlag":"1","Voca":"1"}]}},"SDate":"` + sdate + `","Skill":"30,0,0,0,0,0","VId":"ALC","CId":"TC1","SId":"TC1_S` + fmt.Sprint(subCourse) + `","UId":"TC1_S` + fmt.Sprint(subCourse) + `_U00` + fmt.Sprint(unit) + `-1","SessionId":"` + sID + `"}`)
	// log.Println(data)
	req, err := http.NewRequest("POST", "https://nanext.alcnanext.jp/anetn/api/HistoryApi/registLearnHistory", data)
	if err != nil {
		log.Fatal(err)
	}

	// Cookieなどの追加
	req.Header.Set("Content-Type", "application/JSON")
	req.Header.Set("Cookie", "ASP.NET_SessionId="+sID)

	// 送信
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 受信
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s\n", bodyText)
	// {"Result":"0","EDate":"20240226231503751","TTime":"148"}
}

func main() {
	fmt.Println("TC1 をやります")

	// Cookieの入力
	fmt.Print("SessionIDの入力: ")
	sID := "hoge"
	fmt.Scanln(&sID)

	// サブコースの入力
	fmt.Print("サブコースの入力: ")
	subCourse := 1
	fmt.Scanln(&subCourse)

	// ユニット数の入力
	fmt.Print("ユニット数の入力: ")
	unitCount := 28
	fmt.Scanln(&unitCount)

	// デバック用
	// sID = ""
	subCourse = 1
	unitCount = 28

	for i := 8; i < unitCount+1; i++ {
		for {
			start(sID, subCourse, i)

			ok := "no"
			fmt.Print("正常ならokを入力@start:")
			fmt.Scanln(&ok)
			if ok == "ok" {
				for {
					time.Sleep(6 * time.Second)
					end(sID, subCourse, i)

					ok := "no"
					fmt.Print("正常ならokを入力@end:")
					fmt.Scanln(&ok)
					if ok == "ok" {
						break
					}
				}
				break
			}

			time.Sleep(3 * time.Second)
		}

		time.Sleep(6 * time.Second)
		log.Printf("%dユニットが完了しました", i)
	}
}
