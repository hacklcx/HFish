package report

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"HFish/core/dbUtil"
	"HFish/utils/cache"
	"HFish/utils/log"
)

type CommonResponse struct {
	ResponseCode int             `json:"response_code"`
	VerboseMsg   string          `json:"verbose_msg"`
	Content      json.RawMessage `json:"content"`
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

func readIntelligence(ip string) (string, error) {
	result, err := dbUtil.DB().Table("hfish_intelligence").
		Fields("source", "detail").
		Where("ip", ip).
		First()
	if err != nil{
		log.Pr("HFish", "127.0.0.1", "从数据库获取威胁情报失败", err)
		return "", err
	}
	if len(result) == 0 {
		log.Pr("HFish", "127.0.0.1", "数据库没有该IP对应的威胁情报")
		return "", fmt.Errorf("no this ip intelligence")
	}
	return result["source"].(string) + "&&" + result["detail"].(string), nil
}

func insertIntelligence(source, ip, detail string) error {
	_, err := dbUtil.DB().Table("hfish_intelligence").Data(map[string]interface{}{
		"ip":     ip,
		"source": source,
		"detail": detail,
		"update_time": time.Now().Format("2006-01-02 15:04:05"),
	}).InsertGetId()
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "insert intelligence err", err)
	}
	return err
}

func fetchIntelligenceData(ip string) (string, error) {
	if len(ip) == 0 {
		return "", fmt.Errorf("ERR_IP")
	}
	data, err := readIntelligence(ip)
	if err == nil {
		return data, nil
	}
	status, _ := cache.Get("ApikeyStatus")
	// 判断是否启用获取云端威胁情报
	if status == "0" {
		return "", fmt.Errorf("ERR_APIKEY")
	}
	apikeyInfo, _ := cache.Get("ApikeyInfo")
	apikey, _ := apikeyInfo.(string)

	// apikeyinfo信息是由来源+地址+key组成的, &&分隔
	apiArr := strings.Split(apikey, "&&")
	if len(apiArr) != 3 {
		return "", fmt.Errorf("ERR_APIKEY")
	}
	if apiArr[0] == "xplt" {
		data, err := fetchIntelligenceDataFromXplt(apiArr[1], apiArr[2], ip)
		if err == nil {
			insertIntelligence("xplt", ip, data)
			data = "xplt&&" + data
		}
		return data, err
	} else if apiArr[0] == "tip" {
		data, err := fetchIntelligenceDataFromTip(apiArr[1], apiArr[2], ip)
		if err == nil {
			insertIntelligence("tip", ip, data)
			data = "tip&&" + data
		}
		return data, err
	}
	return "", fmt.Errorf("ERR_APIKEY")
}

func fetchIntelligenceDataFromXplt(server, apikey, ip string) (string, error) {
	url := fmt.Sprintf("%s/v3/scene/ip_reputation?apikey=%s&resource=%s", server, apikey, ip)
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

func fetchIntelligenceDataFromTip(server, apikey, ip string) (string, error) {
	url := fmt.Sprintf("%s/api/v3/intelligence_search?apiKey=%s&data=%s", server, apikey, ip)
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

	intelligenceData, err := json.Marshal(result.Content)
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
	apikeys := strings.Split(apikey, "&&")
	if len(apikeys) != 3 {
		log.Pr("HFish", "127.0.0.1", "apikey illegal:", apikey)
		return fmt.Errorf("format apikey error")
	}
	var key string
	url := "https://x.threatbook.cn/co_intel_inf"
	if apikeys[0] == "xplt" {
		key = apikeys[2]
	}
	body, err := doHttpPost(url, key, info)
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

func FetchIntelligenceData(source, server, apikey, ip string) (string, error) {
	if source == "xplt" {
		return fetchIntelligenceDataFromXplt(server, apikey, ip)
	} else if source == "tip" {
		return fetchIntelligenceDataFromTip(server, apikey, ip)
	}
	return "", fmt.Errorf("illegal source")
}
