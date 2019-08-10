package is

import (
	"HFish/utils/conf"
)

func Rpc() bool {
	rpcStatus := conf.Get("rpc", "status")

	if rpcStatus == "2" {
		return true
	} else {
		return false
	}

}
