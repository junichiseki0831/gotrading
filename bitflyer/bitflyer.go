package bitflyer

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseURL = "https://api.bitflyer.com/v1/"

type APIClient struct {
	key        string
	secret     string
	httpClient *http.Client
}

func New(key, secret string) *APIClient {
	APIClient := &APIClient{key, secret, &http.Client{}}
	return APIClient
}

// リクエストヘッダー作成と認証作成
func (api APIClient) header(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)   //タイムスタンプ取得
	message := timestamp + method + endpoint + string(body) //bitflyerの使用項目通りに成型

	// 鍵付きハッシュメッセージ認証コード実装
	mac := hmac.New(sha256.New, []byte(api.secret))
	mac.Write([]byte(message))
	sign := hex.EncodeToString(mac.Sum(nil))
	return map[string]string{
		"ACCESS-KEY":       api.key,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}
}

//httpリクエスト
func (api *APIClient) doRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
	baseURL, err := url.Parse(baseURL) //urlが正しいか確認
	if err != nil {
		return
	}
	apiURL, err := url.Parse(urlPath) //引数のurlパスを確認
	if err != nil {
		return
	}
	endpoint := baseURL.ResolveReference(apiURL).String() //エンドポイント作成
	log.Printf("action=doRequest endpoint=%s", endpoint)  //ログに残す

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data)) //リクエストの作成(POST)
	if err != nil {
		return
	}

	//query追加
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode() //エンコード忘れず

	//ヘッダー情報追加
	for key, value := range api.header(method, req.URL.RequestURI(), data) {
		req.Header.Add(key, value)
	}

	resp, err := api.httpClient.Do(req) //リクエストを渡す
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() //エラーがなければ終了

	body, err = ioutil.ReadAll(resp.Body) //レスポンスを返す
	if err != nil {
		return nil, err
	}
	return body, nil

}

// レスポンスデータの定義
type Balance struct {
	CurrentCode string  `json:"currency_code"`
	Amount      float64 `json:"amount"`
	Available   float64 `json:"available"`
}

// apiへのアクセス
func (api *APIClient) GetBalance() ([]Balance, error) {
	url := "me/getbalance"
	resp, err := api.doRequest("GET", url, map[string]string{}, nil)
	log.Printf("url=%s resp=%s", url, string(resp))
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	var balance []Balance
	err = json.Unmarshal(resp, &balance)
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	return balance, nil
}
