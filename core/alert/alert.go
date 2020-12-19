package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"HFish/error"
	"HFish/utils/cache"
	"HFish/utils/geo"
	"HFish/utils/log"
	"HFish/utils/passwd"
	"HFish/utils/send"
	"HFish/utils/try"
	"HFish/view/data"
)

func AlertSyslog(model string, projectName string, typex string, agent string, ipx string, country string, region string, city string, infox string, time string) {
	// Judge syslog notification
	try.Try(func() {
		// Only new members will send syslog notifications
		if (model == "new") {
			status, _ := cache.Get("SyslogConfigStatus")

			// Determine whether to enable notification
			if status == "1" {
				info, _ := cache.Get("SyslogConfigInfo")
				configs := strings.Split(info.(string), "&&")

				if (country == "Local address") {
					region = ""
					city = ""
				} else if (country == "local area network") {
					region = ""
					city = ""
				}

				// Determine whether to turn on desensitization
				passwdConfigStatus, _ := cache.Get("PasswdConfigStatus")

				if (passwdConfigStatus == "1") {
					if (typex == "FTP" || typex == "SSH") {
						// Get masked encrypted characters
						passwdConfigInfo, _ := cache.Get("PasswdConfigInfo")

						arr := strings.Split(infox, "&&")

						infox = arr[0] + "&&" + passwd.Desensitization(arr[1], passwdConfigInfo.(string))
					}
				}

				text := fmt.Sprintf("project: %s, type: %s, agent: %s, ip: %s, geo: %s, info: %s, time: %s",
					projectName, typex, agent, ipx, geo.Format(country, region, city, "-"), infox, time)

				log.Pr("HFish", "127.0.0.1", "alert syslog:", text)
				for _, v := range configs {
					config := strings.Split(v, ":")
					send.SendSyslog(config[0], config[1], config[2], text)
				}
			}
		}
	}).Catch(func() {
	})
}

func AlertMail(model string, typex string, agent string, ipx string, country string, region string, city string, infox string) {
	// Judgment email notification
	try.Try(func() {
		// Only new members will send email notifications
		if (model == "new") {
			status, _ := cache.Get("MailConfigStatus")

			// Determine whether to enable notification
			if status == "1" {
				info, _ := cache.Get("MailConfigInfo")
				config := strings.Split(info.(string), "&&")

				if (country == "Local address") {
					region = ""
					city = ""
				} else if (country == "local area network") {
					region = ""
					city = ""
				}

				// Determine whether to turn on desensitization
				passwdConfigStatus, _ := cache.Get("PasswdConfigStatus")

				if (passwdConfigStatus == "1") {
					if (typex == "FTP" || typex == "SSH") {
						// 获取脱敏加密字符
						passwdConfigInfo, _ := cache.Get("PasswdConfigInfo")

						arr := strings.Split(infox, "&&")

						infox = arr[0] + "&&" + passwd.Desensitization(arr[1], passwdConfigInfo.(string))
					}
				}

				geoInfo := geo.Format(country, region, city, " ")
				text := `
				<div><b>Hi, you got the bait! </b></div>
				<div><b><br /></b></div>
				<div><b>Cluster name:</b>` + agent + `</div>
				<div><b>Attack IP:</b>` + ipx + `</div>
				<div><b>Geographic information:</b>` + geoInfo + `</div>
				<div><b>Hook content:</b>` + infox + `</div>
				<div><br /></div>
				<div><span style="color: rgb(128, 128, 128); font-size: 10px;">(HFish auto delivery)</span></div>
				`

				send.SendMail(config[5:], "[HFish] Remind you, "+typex+" A fish is on the bait!", text, config)
			}
		}
	}).Catch(func() {
	})
}

func AlertWebHook(id string, model string, typex string, projectName string, agent string, ipx string, country string, region string, city string, infox string, time string) {
	// Judge WebHook notification
	try.Try(func() {
		status, _ := cache.Get("HookConfigStatus")

		// Determine whether to enable notification
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
				log.Pr("HFish", "127.0.0.1", "WebHook Call failed", err)
			} else {
				log.Pr("HFish", "127.0.0.1", "WebHook Successful call")
			}

			defer resp.Body.Close()
			//defer request.Body.Close()
		}
	}).Catch(func() {
	})
}

// Big data display
func AlertDataWs(model string, typex string, projectName string, agent string, ipx string, country string, region string, city string, time string) {
	if (model == "new") {
		// Splicing dictionary
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

		// Send to client
		data.Send(error.ErrSuccessWithData(d))
	}
}
