package report

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"HFish/utils/cache"
	"HFish/utils/log"
)

type CommonResponse struct {
	ResponseCode int             `json:"response_code"`
	VerboseMsg   string          `json:"verbose_msg"`
	Data         json.RawMessage `json:"data"`
}

func doHttpGet(reqUrl string) ([]byte, error) {
	resp, err := http.Get(reqUrl)
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "http get err", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "read body err", err)
		return nil, err
	}
	return body, nil
}

func doHttpPost(reqUrl, apikey, info string) ([]byte, error) {
	values := url.Values{}
	values.Add("apikey", apikey)
	values.Add("source", "hfish")
	values.Add("info", info)
	resp, err := http.PostForm(reqUrl, values)
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "http post err", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "read body err", err)
		return nil, err
	}
	return body, nil
}

func fetchIntelligenceData(ip string) (string, error) {
	if len(ip) == 0 {
		return "", fmt.Errorf("ERR_IP")
	}
	status, _ := cache.Get("ApikeyStatus")
	// 判断是否启用获取云端威胁情报
	if status == "0" {
		return "", fmt.Errorf("ERR_APIKEY")
	}
	apikey, _ := cache.Get("ApikeyInfo")
	url := fmt.Sprintf("https://api.threatbook.cn/v3/scene/ip_reputation?apikey=%s&resource=%s", apikey, ip)
	body, err := doHttpGet(url)
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "do http post err", err)
		return "", fmt.Errorf("ERR_HTTP")
	}

	var result CommonResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Pr("HFish", "127.0.0.1", "json unmarshal err", err)
		return "", fmt.Errorf("ERR_JSON")
	}
	if result.ResponseCode != 0 {
		return "", fmt.Errorf("ERR_RESP:%d:%s", result.ResponseCode, result.VerboseMsg)
	}

	intelligenceData, err := json.Marshal(result.Data)
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "json marshal err", err)
		return "", fmt.Errorf("ERR_JSON")
	}
	return string(intelligenceData), nil
}

func collectIntelligenceData(info string) error {
	status, _ := cache.Get("ApikeyStatus")
	if status == "0" {
		return fmt.Errorf("apikey disbale")
	}

	apikeyInfo, _ := cache.Get("ApikeyInfo")
	apikey, ok := apikeyInfo.(string)
	if !ok || apikey == "" {
		log.Pr("HFish", "127.0.0.1", "apikey illegal:", apikey)
		return fmt.Errorf("apikey illegal")
	}

	url := "https://x.threatbook.cn/co_intel_inf"
	body, err := doHttpPost(url, apikey, info)
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "do http post err", err)
		return err
	}

	var result CommonResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Pr("HFish", "127.0.0.1", "json unmarshal err", err)
		return err
	}
	if result.ResponseCode != 0 {
		log.Pr("HFish", "127.0.0.1", "collect intelligence fail:", result.VerboseMsg)
		return fmt.Errorf("err response code: %d", result.ResponseCode)
	}
	return nil
}
