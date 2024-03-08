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

func start(sID string, id model.Id) (string, error) {
	// クライアント新規作成
	client := &http.Client{}
	data := strings.NewReader(
		`{"Qtype":"","VId":"ALC","CId":"` + id.CId +
			`","SId":"` + id.SId +
			`","UId":"` + id.UId +
			`","SessionId":"` + sID + `"}`,
	)

	// リクエスト新規作成
	req, err := http.NewRequest("POST", "https://nanext.alcnanext.jp/anetn/api/HistoryApi/registStartHistory", data)
	if err != nil {
		return "err", err
	}

	// リクエストヘッダの追加
	req.Header.Set("Content-Type", "application/JSON")
	req.Header.Set("Cookie", "ASP.NET_SessionId="+sID)

	// 送信
	resp, err := client.Do(req)
	if err != nil {
		return "err", err
	}
	defer resp.Body.Close()

	// 受信
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "err", err
	}

	return fmt.Sprintf("%s\n", bodyText), nil // ex.{"Result":"0","Estep":"","SDate":"20240226231235558"}
}

func end(sID string, id model.Id, sdate string) (string, error) {
	// クライアント新規作成
	client := &http.Client{}
	data := strings.NewReader(
		`{"FId":"02","LCD":"1","LInfo":{"FID02":{"StepSection02":[{"SOrder":"1","SFlag":"1","Voca":"1"},{"SOrder":"2","SFlag":"1","Voca":"1"},{"SOrder":"3","SFlag":"1","Voca":"1"},{"SOrder":"4","SFlag":"1","Voca":"1"},{"SOrder":"5","SFlag":"1","Voca":"1"}]}},"SDate":"` + sdate + `","Skill":"50,0,0,0,0,0","VId":"ALC","CId":"` + id.CId +
			`","SId":"` + id.SId +
			`","UId":"` + id.UId +
			`","SessionId":"` + sID + `"}`,
	)

	// リクエスト新規作成
	req, err := http.NewRequest("POST", "https://nanext.alcnanext.jp/anetn/api/HistoryApi/registLearnHistory", data)
	if err != nil {
		return "err", err
	}

	// リクエストヘッダの追加
	req.Header.Set("Content-Type", "application/JSON")
	req.Header.Set("Cookie", "ASP.NET_SessionId="+sID)

	// 送信
	resp, err := client.Do(req)
	if err != nil {
		return "err", err
	}
	defer resp.Body.Close()

	// 受信
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "err", err
	}

	return fmt.Sprintf("%s\n", bodyText), nil // ex.{"Result":"0","EDate":"20240226231503751","TTime":"148"}
}

func Submit(sID string, id model.Id) error {
	log.Printf("Start of unit submission\n")
	sdate := "ex.20240226231235558"

	for i := 0; ; i++ {
		res, err := start(sID, id) // ex.{"Result":"0","Estep":"","SDate":"20240226231235558"}
		if err != nil {
			return xerrors.Errorf("Error on request to start unit: %w", err)
		}

		// 期待するレスポンスであれば，SDateの値のみを抽出
		if string([]rune(res)[1:13]) == `"Result":"0"` && !strings.Contains(res, "null") {
			for i := 0; i+1 < len(res); i++ {
				if string([]rune(res)[i:i+1]) == "S" {
					sdate = string([]rune(res)[i+8 : len(res)-3])
				}
			}
			break
		}

		// 待ち時間
		time.Sleep(600 * time.Millisecond)

		// ３回以上失敗するとレスポンスを標準出力
		if i > 3 {
			log.Println(res)
		}

		// 失敗しすぎるとエラーとして処理
		if i > 8 || string([]rune(res)[2:8]) == `Result` {
			return xerrors.Errorf("Error or unexpected value returned from server")
		}
	}
	time.Sleep(1 * time.Second)

	for i := 0; ; i++ {
		res, err := end(sID, id, sdate)
		if err != nil {
			return xerrors.Errorf("Error on request to finish unit: %w", err)
		}

		// 期待するレスポンスなら終了
		if string([]rune(res)[1:13]) == `"Result":"0"` && !strings.Contains(res, "null") {
			break
		}

		// 待ち時間
		time.Sleep(1 * time.Second)

		// ３回以上失敗するとレスポンスを標準出力
		if i > 3 {
			log.Println(res)
		}

		// 失敗しすぎるとエラーとして処理
		if i > 8 || string([]rune(res)[2:8]) == `Result` {
			return xerrors.Errorf("Error or unexpected value returned from server")
		}
	}
	time.Sleep(1 * time.Second)

	log.Printf("End of unit submission\n\n")
	return nil
}
