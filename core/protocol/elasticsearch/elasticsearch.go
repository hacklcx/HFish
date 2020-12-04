package elasticsearch

import (
    "fmt"
    "net/http"
    "time"
    "strings"
    "HFish/core/report"
    "HFish/core/rpc/client"
    "HFish/utils/is"
)

// Config represents the configuration information.
type Config struct {
	LogFile        string  `json:"logfile"`
	UseRemote      bool    `json:"use_remote"`
	Remote         Remote  `json:"remote"`
	HpFeeds        HpFeeds `json:"hpfeeds"`
	InstanceName   string  `json:"instance_name"`
	Anonymous      bool    `json:"anonymous"`
	SensorIP       string  `json:"honeypot_ip"`
	SpoofedVersion string  `json:"spoofed_version"`
	PublicIpUrl    string  `json:"public_ip_url"`
}

// Remote is a struct used to contain the details for a remote server connection
type Remote struct {
	URL     string `json:"url"`
	UseAuth bool   `json:"use_auth"`
	Auth    Auth   `json:"auth"`
}

type HpFeeds struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Channel string `json:"channel"`
	Ident   string `json:"ident"`
	Secret  string `json:"secret"`
	Enabled bool   `json:"enabled"`
}

// Auth contains the details in case basic auth is to be used when connecting
// to the remote server
type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Conf holds the global config
var Conf Config

// Attack is a struct that contains the details of an attack entry
type Attack struct {
	SourceIP  string    `json:"source"`
	Timestamp time.Time `json:"@timestamp"`
	URL       string    `json:"url"`
	Method    string    `json:"method"`
	Form      string    `json:"form"`
	Payload   string    `json:"payload"`
	Headers   Headers   `json:"headers"`
	Type      string    `json:"type"`
	SensorIP  string    `json:"honeypot"`
}

// Headers contains the filtered headers of the HTTP request
type Headers struct {
	UserAgent      string `json:"user_agent"`
	Host           string `json:"host"`
	ContentType    string `json:"content_type"`
	AcceptLanguage string `json:"accept_language"`
}

func FakeBanner(w http.ResponseWriter, r *http.Request) {
	printInfo(r)
	response := fmt.Sprintf(`{
        "status" : 200,
        "name" : "%s",
        "cluster_name" : "elasticsearch",
        "version" : {
            "number" : "%s",
            "build_hash" : "89d3241d670db65f994242c8e838b169779e2d4",
            "build_snapshot" : false,
            "lucene_version" : "4.10.2"
        },
        "tagline" : "You Know, for Search"
    }`, Conf.InstanceName, Conf.SpoofedVersion)
	WriteResponse(w, response)
	return
}

// FakeNodes presents a fake /_nodes result
// TODO: Change IP Address with actual server IP address
func FakeNodes(w http.ResponseWriter, r *http.Request) {
	printInfo(r)
	response := fmt.Sprintf(`
	{
        "cluster_name" : "elasticsearch",
        "nodes" : {
            "x1JG6g9PRHy6ClCOO2-C4g" : {
              "name" : "%s",
              "transport_address" : "inet[/
			%s:9300]",
              "host" : "elk",
              "ip" : "127.0.1.1",
              "version" : "%s",
              "build" : "89d3241",
              "http_address" : "inet[/%s:9200]",
              "os" : {
                "refresh_interval_in_millis" : 1000,
                "available_processors" : 12,
                "cpu" : {
                  "total_cores" : 24,
                  "total_sockets" : 48,
                  "cores_per_socket" : 2
                }
              },
              "process" : {
                "refresh_interval_in_millis" : 1000,
                "id" : 2039,
                "max_file_descriptors" : 65535,
                "mlockall" : false
              },
              "jvm" : {
                "version" : "1.7.0_65"
              },
              "network" : {
                "refresh_interval_in_millis" : 5000,
                "primary_interface" : {
                  "address" : "%s",
                  "name" : "eth0",
                  "mac_address" : "08:01:c7:3F:15:DD"
                }
              },
              "transport" : {
                "bound_address" : "inet[/0:0:0:0:0:0:0:0:9300]",
                "publish_address" : "inet[/%s:9300]"
              },
              "http" : {
                "bound_address" : "inet[/0:0:0:0:0:0:0:0:9200]",
                "publish_address" : "inet[/%s:9200]",
                "max_content_length_in_bytes" : 104857600
              }}
            }
        }`, Conf.InstanceName, Conf.SensorIP, Conf.SpoofedVersion, Conf.SensorIP, Conf.SensorIP, Conf.SensorIP, Conf.SensorIP)
	WriteResponse(w, response)
	return
}

// FakeSearch returns fake search results
func FakeSearch(w http.ResponseWriter, r *http.Request) {
	printInfo(r)
	response := fmt.Sprintf(`
	{
        "took" : 6,
        "timed_out" : false,
        "_shards" : {
            "total" : 6,
            "successful" : 6,
            "failed" : 0
        },
        "hits" : {
            "total" : 1,
            "max_score" : 1.0,
            "hits" : [ {
                "_index" : ".kibana",
                "_type" : "index-pattern",
                "_id" : "logstash-*",
                "_score" : 1.0,
                "_source":{"title":"logstash-*","timeFieldName":"@timestamp","customFormats":"{}","fields":"[{\"type\":\"string\",\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"name\":\"host\",\"count\":0},{\"type\":\"string\",\"indexed\":false,\"analyzed\":false,\"name\":\"_source\",\"count\":0},{\"type\":\"string\",\"indexed\":true,\"analyzed\":false,\"doc_values\":false,\"name\":\"message.raw\",\"count\":0},{\"type\":\"string\",\"indexed\":false,\"analyzed\":false,\"name\":\"_index\",\"count\":0},{\"type\":\"string\",\"indexed\":true,\"analyzed\":false,\"doc_values\":false,\"name\":\"@version\",\"count\":0},{\"type\":\"string\",\"indexed\":true,\"analyzed\":true,\"doc_values\":false,\"name\":\"message\",\"count\":0},{\"type\":\"date\",\"indexed\":true,\"analyzed\":false,\"doc_values\":false,\"name\":\"@timestamp\",\"count\":0},{\"type\":\"string\",\"indexed\":true,\"analyzed\":false,\"name\":\"_type\",\"count\":0},{\"type\":\"string\",\"indexed\":true,\"analyzed\":false,\"name\":\"_id\",\"count\":0},{\"type\":\"string\",\"indexed\":true,\"analyzed\":false,\"doc_values\":false,\"name\":\"host.raw\",\"count\":0},{\"type\":\"geo_point\",\"indexed\":true,\"analyzed\":false,\"doc_values\":false,\"name\":\"geoip.location\",\"count\":0}]"}
            }]
        }
    }`)
	WriteResponse(w, response)
	return
}

func WriteResponse(w http.ResponseWriter, d string) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(d))
	return
}

func printInfo(r *http.Request) {
	info := "URL:" + r.URL.String() + "&&Method:" + r.Method + "&&RemoteAddr:" + r.RemoteAddr

	arr := strings.Split(r.RemoteAddr, ":")

	// 判断是否为 RPC 客户端
	if is.Rpc() {
		go client.ReportResult("ES", "ES蜜罐", arr[0], info, "0")
	} else {
		go report.ReportEs("ES蜜罐", "本机", arr[0], info)
	}
}

func Start(address string) {
	// Create the handlers
	http.HandleFunc("/", FakeBanner)
	http.HandleFunc("/_nodes", FakeNodes)
	http.HandleFunc("/_search", FakeSearch)

	// Start the server
	http.ListenAndServe(address, nil)
}
