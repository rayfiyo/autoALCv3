// ユニットを処理する POST を投稿 (submit)

package tasks

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/rayfiyo/autoALCv3/model"
	"golang.org/x/xerrors"
)

func start(id model.Id) (string, string, error) {
	// クライアント新規作成
	client := &http.Client{}
	data := strings.NewReader(
		`{"Qtype":"","VId":"ALC","CId":"` + id.CId +
			`","SId":"` + id.SId +
			`","UId":"` + id.UId +
			`","SessionId":"` + id.SessId + `"}`,
	)

	// リクエスト新規作成
	req, err := http.NewRequest("POST", "https://nanext.alcnanext.jp/anetn/api/HistoryApi/registStartHistory", data)
	if err != nil {
		return "err", fmt.Sprintln(data), err
	}

	// リクエストヘッダの追加
	req.Header.Set("Content-Type", "application/JSON")
	req.Header.Set("Cookie", "ASP.NET_SessionId="+id.SessId)

	// 送信
	resp, err := client.Do(req)
	if err != nil {
		return "err", fmt.Sprintln(data), err
	}
	defer resp.Body.Close()

	// 受信
	bodyText, err := io.ReadAll(resp.Body) // 第一返り値は []byte
	if err != nil {
		return fmt.Sprintf("%s\n", bodyText), fmt.Sprintln(data), err
	}

	// ex.{"Result":"0","Estep":null,"SDate":"20240310233816521"}
	return fmt.Sprintf("%s\n", bodyText), fmt.Sprintln(data), nil
}

func end(id model.Id, stCnt int, sdate string) (string, string, error) {
	stepSection := ""
	stCnt = 5
	for i := 0; i < stCnt; i++ {
		if i > 0 {
			stepSection += `,`
		}
		// SOrder は自然数で管理しているのでインクリメント
		stepSection += `{"SOrder":"` + fmt.Sprint(i+1) + `","SFlag":"1","Voca":"1"}`
	}
	// {"SOrder":"1","SFlag":"1","Voca":"1"},
	// {"SOrder":"2","SFlag":"1","Voca":"1"},
	// {"SOrder":"3","SFlag":"1","Voca":"1"},
	// {"SOrder":"4","SFlag":"1","Voca":"1"},
	// {"SOrder":"5","SFlag":"1","Voca":"1"}

	// クライアント新規作成
	client := &http.Client{}
	data := strings.NewReader(
		`{"FId":"02","LCD":"1","LInfo":{"FID02":{"StepSection02":[` + stepSection +
			`]}},"SDate":"` + sdate +
			`","Skill":"` + fmt.Sprint(10*stCnt) + `,0,0,0,0,0","VId":"ALC","CId":"` + id.CId +
			`","SId":"` + id.SId +
			`","UId":"` + id.UId +
			`","SessionId":"` + id.SessId + `"}`,
	)

	// log.Print("stCnd: ", stCnt, "\n")
	// log.Print("send: ", data, "\n")
	// log.Print("want: ", `__{"FId":"02","LCD":"1","LInfo":{"FID02":{"StepSection02":[{"SOrder":"1","SFlag":"1","Voca":"1"},{"SOrder":"2","SFlag":"1","Voca":"1"},{"SOrder":"3","SFlag":"1","Voca":"1"},{"SOrder":"4","SFlag":"1","Voca":"1"},{"SOrder":"5","SFlag":"1","Voca":"1"}]}},"SDate":"`+sdate+`","Skill":"50,0,0,0,0,0","VId":"ALC","CId":"`+id.CId+`","SId":"`+id.SId+`","UId":"`+id.UId+`","SessionId":"`+id.SessId+`"}`, "\n")

	// time.Sleep(12 * time.Minute)

	// リクエスト新規作成
	req, err := http.NewRequest("POST", "https://nanext.alcnanext.jp/anetn/api/HistoryApi/registLearnHistory", data)
	if err != nil {
		return "err", fmt.Sprintln(data), err
	}

	// リクエストヘッダの追加
	req.Header.Set("Content-Type", "application/JSON")
	req.Header.Set("Cookie", "ASP.NET_SessionId="+id.SessId)

	// 送信
	resp, err := client.Do(req)
	if err != nil {
		return "err", fmt.Sprintln(data), err
	}
	defer resp.Body.Close()

	// 受信
	bodyText, err := io.ReadAll(resp.Body) // 第一返り値は []byte
	if err != nil {
		return fmt.Sprintf("%s\n", bodyText), fmt.Sprintln(data), err
	}

	// ex.{"Result":"0","EDate":"20240226231503751","TTime":"148"}
	return fmt.Sprintf("%s\n", bodyText), fmt.Sprintln(data), nil
}

func Submit(id model.Id, stCnt int) error {
	log.Printf("Start of unit submission\n")
	sdate := "ex.20240226231235558"

	for i := 0; ; i++ {
		res, data, err := start(id) // ex.{"Result":"0","Estep":"","SDate":"20240226231235558"}
		if err != nil {
			return xerrors.Errorf("Error on request to start unit: %w\n Response: %s\nSend data: %s\n", err, res, data)
		}

		// 期待するレスポンスなら，SDateの値のみを抽出
		if string(res[1:13]) == `"Result":"0"` {
			for j := 0; j+1 < len(res); j++ {
				if string(res[j:j+1]) == "S" {
					sdate = string(res[j+8 : len(res)-3])
				}
			}
			if sdate == "null" { // sdate が不正な値（厳密にはnullのみ）ではないか調べる
				return xerrors.Errorf("SDate is null @start: %w\n Response: %s\nSend data: %s\n", err, res, data)
			}
			break
		}

		// 待ち時間
		time.Sleep(600 * time.Millisecond)

		// ４～６回目はレスポンスを標準出力
		if i > 3 && i < 7 {
			log.Println(res)
		}

		// ７回目以降は失敗（エラー）として処理
		if i > 6 {
			return xerrors.Errorf("Number of attempts exceeded. The format of the data sent is correct, but the value may be incorrect. @start\n Response: %s\nSend data: %s\n", res, data)
		}

		// Result から始まらないレスポンスだとエラーとして処理
		if string(res[2:8]) != `Result` {
			return xerrors.Errorf("Unexpected response. It is possible that there is an error on the server side or that the data sent is in an invalid format. @start\n Response: %s\nSend data: %s\n", res, data)
		}
	}
	time.Sleep(1 * time.Second)

	for i := 0; ; i++ {
		res, data, err := end(id, stCnt, sdate)
		if err != nil {
			return xerrors.Errorf("Error on request to finish unit: %w\n Response: %s\nSend data: %s\n", err, res, data)
		}

		// 期待するレスポンスなら終了
		if string(res[1:13]) == `"Result":"0"` && !strings.Contains(res, "null") {
			break
		}

		// 待ち時間

		// 待ち時間
		time.Sleep(1 * time.Second)

		// ４～６回目はレスポンスを標準出力
		if i > 3 && i < 7 {
			log.Println(res)
		}

		// ７回目以降は失敗（エラー）として処理
		if i > 6 {
			return xerrors.Errorf("Number of attempts exceeded. The format of the data sent is correct, but the value may be incorrect. @finish\n Response: %s\nSend data: %s\n", res, data)
		}

		// Result から始まらないレスポンスだとエラーとして処理
		if string(res[2:8]) != `Result` {
			return xerrors.Errorf("Unexpected response. It is possible that there is an error on the server side or that the data sent is in an invalid format. @finish\n Response: %s\nSend data: %s\n", res, data)
		}
	}
	time.Sleep(1 * time.Second)

	log.Printf("Finish of unit submission\n\n")
	return nil
}
