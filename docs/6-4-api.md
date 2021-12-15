![image-20210730173724481](http://img.threatbook.cn/hfish/20210730173725.png)

> API配置页面是管理HFish api_key和查阅API样例的页面

通过HFish的API接口，能够实现把蜜罐数据同步给其它应用或设备，从而实现更自主丰富的数据展示和安全设备联动。

！> 相关SDK我们希望能够得到社区用户的支持，大家能够把自己对蜜罐数据的使用方案贡献出来，帮助更多人的使用。



### 已经支持的三种API接口调用示例



- 获取攻击来源

<!-- tabs:start -->

#### **cURL**

```curl
curl --location --request POST 'https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY' \
--header 'Content-Type: application/json' \
--data '{
  "start_time": 0,
  "end_time": 0,
  "intranet": "0",
  "threat_label": [
    "Scanner"
  ]
}'
```

#### **Python**

```python
import requests
import json

url = "https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY"

payload = json.dumps({
  "start_time": 0,
  "end_time": 0,
  "intranet": "0",
  "threat_label": [
    "Scanner"
  ]
})
headers = {
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)

```

#### **GO**

```go
package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY"
  method := "POST"

  payload := strings.NewReader(`{
  "start_time": 0,
  "end_time": 0,
  "intranet": "0",
  "threat_label": [
    "Scanner"
  ]
}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}

```

#### **JAVA**

```java
OkHttpClient client = new OkHttpClient().newBuilder()
  .build();
MediaType mediaType = MediaType.parse("application/json");
RequestBody body = RequestBody.create(mediaType, "{\n  \"start_time\": 0,\n  \"end_time\": 0,\n  \"intranet\": \"0\",\n  \"threat_label\": [\n    \"Scanner\"\n  ]\n}");
Request request = new Request.Builder()
  .url("https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY")
  .method("POST", body)
  .addHeader("Content-Type", "application/json")
  .build();
Response response = client.newCall(request).execute();
```

#### **JavaScript**

```javascript
var settings = {
  "url": "https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY",
  "method": "POST",
  "timeout": 0,
  "headers": {
    "Content-Type": "application/json"
  },
  "data": JSON.stringify({
    "start_time": 0,
    "end_time": 0,
    "intranet": "0",
    "threat_label": [
      "Scanner"
    ]
  }),
};

$.ajax(settings).done(function (response) {
  console.log(response);
});
```

#### **PHP**

```php
<?php

$curl = curl_init();

curl_setopt_array($curl, array(
  CURLOPT_URL => 'https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY',
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_ENCODING => '',
  CURLOPT_MAXREDIRS => 10,
  CURLOPT_TIMEOUT => 0,
  CURLOPT_FOLLOWLOCATION => true,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => 'POST',
  CURLOPT_POSTFIELDS =>'{
  "start_time": 0,
  "end_time": 0,
  "intranet": "0",
  "threat_label": [
    "Scanner"
  ]
}',
  CURLOPT_HTTPHEADER => array(
    'Content-Type: application/json'
  ),
));

$response = curl_exec($curl);

curl_close($curl);
echo $response;

```

#### **Shell**

```shell
wget --no-check-certificate --quiet \
  --method POST \
  --timeout=0 \
  --header 'Content-Type: application/json' \
  --body-data '{
  "start_time": 0,
  "end_time": 0,
  "intranet": "0",
  "threat_label": [
    "Scanner"
  ]
}' \
   'https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY'
```

<!-- tabs:end -->


- 获取攻击详情


<!-- tabs:start -->

#### **cURL**

```curl
curl --location --request POST 'https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY' \
--header 'Content-Type: application/json' \
--data '{
  "start_time": 0,
  "end_time": 0,
  "page_no": 1,
  "page_size": 100,
  "threat_label": ["Scanner"],
  "client_id": [],
  "service_name": [],
  "info_confirm": "1"
}'
```

#### **Python**

```python
import requests
import json

url = "https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY"

payload = json.dumps({
  "start_time": 0,
  "end_time": 0,
  "page_no": 1,
  "page_size": 100,
  "threat_label": [
    "Scanner"
  ],
  "client_id": [],
  "service_name": [],
  "info_confirm": "1"
})
headers = {
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)

```

#### **GO**

```go
package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY"
  method := "POST"

  payload := strings.NewReader(`{
  "start_time": 0,
  "end_time": 0,
  "page_no": 1,
  "page_size": 100,
  "threat_label": ["Scanner"],
  "client_id": [],
  "service_name": [],
  "info_confirm": "1"
}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
```

#### **JAVA**

```java
OkHttpClient client = new OkHttpClient().newBuilder()
  .build();
MediaType mediaType = MediaType.parse("application/json");
RequestBody body = RequestBody.create(mediaType, "{\n  \"start_time\": 0,\n  \"end_time\": 0,\n  \"page_no\": 1,\n  \"page_size\": 100,\n  \"threat_label\": [\"Scanner\"],\n  \"client_id\": [],\n  \"service_name\": [],\n  \"info_confirm\": \"1\"\n}");
Request request = new Request.Builder()
  .url("https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY")
  .method("POST", body)
  .addHeader("Content-Type", "application/json")
  .build();
Response response = client.newCall(request).execute();
```

#### **JavaScript**

```javascript
var settings = {
  "url": "https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY",
  "method": "POST",
  "timeout": 0,
  "headers": {
    "Content-Type": "application/json"
  },
  "data": JSON.stringify({
    "start_time": 0,
    "end_time": 0,
    "page_no": 1,
    "page_size": 100,
    "threat_label": [
      "Scanner"
    ],
    "client_id": [],
    "service_name": [],
    "info_confirm": "1"
  }),
};

$.ajax(settings).done(function (response) {
  console.log(response);
});
```

#### **PHP**

```php
<?php

$curl = curl_init();

curl_setopt_array($curl, array(
  CURLOPT_URL => 'https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY',
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_ENCODING => '',
  CURLOPT_MAXREDIRS => 10,
  CURLOPT_TIMEOUT => 0,
  CURLOPT_FOLLOWLOCATION => true,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => 'POST',
  CURLOPT_POSTFIELDS =>'{
  "start_time": 0,
  "end_time": 0,
  "page_no": 1,
  "page_size": 100,
  "threat_label": ["Scanner"],
  "client_id": [],
  "service_name": [],
  "info_confirm": "1"
}',
  CURLOPT_HTTPHEADER => array(
    'Content-Type: application/json'
  ),
));

$response = curl_exec($curl);

curl_close($curl);
echo $response;

```

#### **Shell**

```shell
wget --no-check-certificate --quiet \
  --method POST \
  --timeout=0 \
  --header 'Content-Type: application/json' \
  --body-data '{
  "start_time": 0,
  "end_time": 0,
  "page_no": 1,
  "page_size": 100,
  "threat_label": ["Scanner"],
  "client_id": [],
  "service_name": [],
  "info_confirm": "1"
}' \
   'https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY'
```

<!-- tabs:end -->



- 获取该IP攻击使用的账号信息


<!-- tabs:start -->

#### **cURL**

```curl
curl --location --request POST 'https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY' \
--header 'Content-Type: application/json' \
--data '{
  "start_time": 0,
  "end_time": 0,
  "intranet": "0",
  "threat_label": [
    "Scanner"
  ]
}'
```

#### **Python**

```python
import requests
import json

url = "https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY"

payload = json.dumps({
  "start_time": 0,
  "end_time": 0,
  "attack_ip": []
})
headers = {
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)

```

#### **GO**

```go
package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY"
  method := "POST"

  payload := strings.NewReader(`{
  "start_time": 0,
  "end_time": 0,
  "attack_ip": []
}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
```

#### **JAVA**

```java
OkHttpClient client = new OkHttpClient().newBuilder()
  .build();
MediaType mediaType = MediaType.parse("application/json");
RequestBody body = RequestBody.create(mediaType, "{\n  \"start_time\": 0,\n  \"end_time\": 0,\n  \"attack_ip\": []\n}");
Request request = new Request.Builder()
  .url("https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY")
  .method("POST", body)
  .addHeader("Content-Type", "application/json")
  .build();
Response response = client.newCall(request).execute();
```

#### **JavaScript**

```javascript
var settings = {
  "url": "https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY",
  "method": "POST",
  "timeout": 0,
  "headers": {
    "Content-Type": "application/json"
  },
  "data": JSON.stringify({
    "start_time": 0,
    "end_time": 0,
    "attack_ip": []
  }),
};

$.ajax(settings).done(function (response) {
  console.log(response);
});
```

#### **PHP**

```php
<?php

$curl = curl_init();

curl_setopt_array($curl, array(
  CURLOPT_URL => 'https://Server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY',
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_ENCODING => '',
  CURLOPT_MAXREDIRS => 10,
  CURLOPT_TIMEOUT => 0,
  CURLOPT_FOLLOWLOCATION => true,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => 'POST',
  CURLOPT_POSTFIELDS =>'{
  "start_time": 0,
  "end_time": 0,
  "attack_ip": []
}',
  CURLOPT_HTTPHEADER => array(
    'Content-Type: application/json'
  ),
));

$response = curl_exec($curl);

curl_close($curl);
echo $response;

```

#### **Shell**

```shell
wget --no-check-certificate --quiet \
  --method POST \
  --timeout=0 \
  --header 'Content-Type: application/json' \
  --body-data '{
  "start_time": 0,
  "end_time": 0,
  "attack_ip": []
}' \
   'https://server_IP/api/v1/attack/ip?api_key=YOUR_API_KEY'
```

<!-- tabs:end -->

