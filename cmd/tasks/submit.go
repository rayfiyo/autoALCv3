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
	return string(bodyText), fmt.Sprintln(data), nil
}

func end(id model.Id, stCnt int, sdate string) (string, string, error) {
	// ステップの数に応じた数の偽装データを作成
	solvedStepData := ""
	for i := 0; i < stCnt; i++ {
		if i > 0 {
			solvedStepData += `,`
		}
		// SOrder は自然数で管理しているのでインクリメント
		solvedStepData += `{"SOrder":"` + fmt.Sprint(i+1) + `","SFlag":"1","Voca":"1"}`
	}

	// スキルポイントの調整 // [NOTE] スキルポイントは要改善
	skillPointData := "   L,S,R,W,G,V"
	if strings.Contains(id.UId, "PWH") {
		skillPointData = "0,0,0,0,0,10"
	} else if strings.Contains(id.UId, "TC") {
		skillPointData = "10,0,0,0,0,0"
	} else if strings.Contains(id.UId, "JT") {
		skillPointData = "0,0,0,0,0,0"
	} else if strings.Contains(id.UId, "KT") {
		skillPointData = "0,0,0,0,0,0" // 点数なし
	} else {
		skillPointData = "0,0,0,0,0,0"
		log.Println("未登録の種類です，スキルポイントを０で設定します．")
	}

	// クライアント新規作成
	client := &http.Client{}
	data := strings.NewReader(
		`{"FId":"02","LCD":"1","LInfo":{"FID02":{"StepSection02":[` + solvedStepData +
			`]}},"SDate":"` + sdate +
			`","Skill":"` + skillPointData + `","VId":"ALC","CId":"` + id.CId +
			`","SId":"` + id.SId +
			`","UId":"` + id.UId +
			`","SessionId":"` + id.SessId + `"}`,
	)

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
	return string(bodyText), fmt.Sprintln(data), nil
}

func Submit(id model.Id, stCnt int) error {
	log.Println("Start  of unit submission")
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
					sdate = string(res[j+8 : len(res)-2])
				}
			}

			// sdate が null（不正な値）ではないか調べる
			if sdate == "null" {
				return xerrors.Errorf("SDate is null @start: %w\n Response: %s\nSend data: %s\n", err, res, data)
			}

			// sdate が 17文字ではない文字数（不正な値）ではないか調べる．Count の "" は Unicode コードポイントの数 + 1 を返す
			if strings.Count(sdate, "")-1 != 17 {
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

	log.Println("Finish of unit submission")
	return nil
}
