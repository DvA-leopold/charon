package main

import (
	"charon/core"
	"charon/methods"
	"net/http"
)

func init() {
	core.ReadConfig("its_stub.txt")
	if address, exist := core.Config["HTTPListen"]; exist {
		core.InitHTTPServer(address,
			map[string]func(w http.ResponseWriter, r *http.Request) {
				"/auth/": methods.Auth,
				"/updInfo/": methods.UpdateInfo,
				"/addPod/": methods.AddPod,
				"/attachPod/": methods.AttachPod,
				"/webdav/": methods.RedirectWebdav,
			})
	}
	//core.InitTCPServer(core.Config["TCPListen"])
}

func main() {
	core.WaitGroup.Wait()
}
