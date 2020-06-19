package alert

import (
	"HFish/utils/try"
	"strings"
	"HFish/utils/send"
	"bytes"
	"net/http"
	"HFish/utils/log"
	"encoding/json"
	"HFish/view/data"
	"github.com/gin-gonic/gin"
	"HFish/error"
	"HFish/utils/cache"
	"HFish/utils/passwd"
)

func AlertMail(model string, typex string, agent string, ipx string, country string, region string, city string, infox string) {
	// 判断邮件通知
	try.Try(func() {
		// 只有新加入才会发送邮件通知
		if (model == "new") {
			status, _ := cache.Get("MailConfigStatus")

			// 判断是否启用通知
			if status == "1" {
				info, _ := cache.Get("MailConfigInfo")
				config := strings.Split(info.(string), "&&")

				if (country == "本地地址") {
					region = ""
					city = ""
				} else if (country == "局域网") {
					region = ""
					city = ""
				}

				// 判断是否开启脱敏
				passwdConfigStatus, _ := cache.Get("PasswdConfigStatus")

				if (passwdConfigStatus == "1") {
					if (typex == "FTP" || typex == "SSH") {
						// 获取脱敏加密字符
						passwdConfigInfo, _ := cache.Get("PasswdConfigInfo")

						arr := strings.Split(infox, "&&")

						infox = arr[0] + "&&" + passwd.Desensitization(arr[1], passwdConfigInfo.(string))
					}
				}

				text := `
				<div><b>Hi，上钩了！</b></div>
				<div><b><br /></b></div>
				<div><b>集群名称：</b>` + agent + `</div>
				<div><b>攻击IP：</b>` + ipx + `</div>
				<div><b>地理信息：</b>` + country + ` ` + region + ` ` + city + `</div>
				<div><b>上钩内容：</b>` + infox + `</div>
				<div><br /></div>
				<div><span style="color: rgb(128, 128, 128); font-size: 10px;">(HFish 自动发送)</span></div>
				`

				send.SendMail(config[4:], "[HFish]提醒你，"+typex+"有鱼上钩!", text, config)
			}
		}
	}).Catch(func() {
	})
}

func AlertWebHook(id string, model string, typex string, projectName string, agent string, ipx string, country string, region string, city string, infox string, time string) {
	// 判断 WebHook 通知
	try.Try(func() {
		status, _ := cache.Get("HookConfigStatus")

		// 判断是否启用通知
		if status == "1" {
			info, _ := cache.Get("HookConfigInfo")

			song := make(map[string]interface{})
			song["id"] = id
			song["model"] = model
			song["project"] = projectName
			song["type"] = typex
			song["agent"] = agent
			song["ip"] = ipx
			song["country"] = country
			song["region"] = region
			song["city"] = city
			song["info"] = infox
			song["time"] = time

			bytesData, _ := json.Marshal(song)

			reader := bytes.NewReader(bytesData)

			request, _ := http.NewRequest("POST", info.(string), reader)
			request.Header.Set("Content-Type", "application/json;charset=UTF-8")

			client := http.Client{}
			resp, err := client.Do(request)

			if err != nil {
				log.Pr("HFish", "127.0.0.1", "WebHook 调用失败", err)
			} else {
				log.Pr("HFish", "127.0.0.1", "WebHook 调用成功")
			}

			defer resp.Body.Close()
			//defer request.Body.Close()
		}
	}).Catch(func() {
	})
}

// 大数据展示
func AlertDataWs(model string, typex string, projectName string, agent string, ipx string, country string, region string, city string, time string) {
	if (model == "new") {
		// 拼接字典
		d := data.MakeDataJson("center_data", map[string]interface{}{
			"type":        typex,
			"projectName": projectName,
			"agent":       agent,
			"ipx":         ipx,
			"country":     country,
			"region":      region,
			"city":        city,
			"time":        time,
		})

		// 发送到客户端
		data.Send(gin.H{
			"code": error.ErrSuccessCode,
			"msg":  error.ErrSuccessMsg,
			"data": d,
		})
	}
}
