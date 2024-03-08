package tasks

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/rayfiyo/autoALCv3/model"
)

func start(sID string, id model.Id) string {
	// クライアント作成・リクエストヘッダの作成
	client := &http.Client{}
	data := strings.NewReader(
		`{"Qtype":"","VId":"ALC","CId":"` + id.CId +
			`","SId":"` + id.SId +
			`","UId":"` + id.UId +
			`","SessionId":"` + sID + `"}`,
	)

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

	return fmt.Sprintf("%s\n", bodyText) // ex.{"Result":"0","Estep":"","SDate":"20240226231235558"}
}

func end(sID string, id model.Id, sdate string) string {
	// クライアント作成・リクエストヘッダの作成
	client := &http.Client{}
	data := strings.NewReader(
		`{"FId":"02","LCD":"1","LInfo":{"FID02":{"StepSection02":[{"SOrder":"1","SFlag":"1","Voca":"1"},{"SOrder":"2","SFlag":"1","Voca":"1"},{"SOrder":"3","SFlag":"1","Voca":"1"},{"SOrder":"4","SFlag":"1","Voca":"1"},{"SOrder":"5","SFlag":"1","Voca":"1"}]}},"SDate":"` + sdate + `","Skill":"50,0,0,0,0,0","VId":"ALC","CId":"` + id.CId +
			`","SId":"` + id.SId +
			`","UId":"` + id.UId +
			`","SessionId":"` + sID + `"}`,
	)

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

	return fmt.Sprintf("%s\n", bodyText) // ex.{"Result":"0","EDate":"20240226231503751","TTime":"148"}
}

func Submit(sID string, id model.Id) {
	log.Printf("Start of unit submission\n")
	sdate := "ex.20240226231235558"

	for {
		res := start(sID, id) // ex.{"Result":"0","Estep":"","SDate":"20240226231235558"}
		if string([]rune(res)[1:13]) == `"Result":"0"` && !strings.Contains(res, "null") {
			for i := 0; i+1 < len(res); i++ {
				if string([]rune(res)[i:i+1]) == "S" {
					sdate = string([]rune(res)[i+8 : len(res)-3])
				}
			}
			break
		}
		time.Sleep(600 * time.Millisecond)

		log.Println(res)
	}
	time.Sleep(1 * time.Second)

	for {
		res := end(sID, id, sdate)
		if string([]rune(res)[1:13]) == `"Result":"0"` && !strings.Contains(res, "null") {
			break
		}
		time.Sleep(1 * time.Second)

		log.Println(res)
	}
	time.Sleep(1 * time.Second)

	log.Printf("End of unit submission\n\n")
}
