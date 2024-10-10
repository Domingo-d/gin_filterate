/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/30 11:54
 */

package core

import (
	"filterate/global"
	"filterate/initialize"
	"log"
	"net/http"
)

func RunServer() {
	global.AhoCorasick = initialize.NewAhoCorasick()
	if nil == global.AhoCorasick.ReadPattern(global.FilterateName) {
		global.AhoCorasick.Build()
	}

	global.Router = initialize.Routers()

	s := &http.Server{Addr: "0.0.0.0:8080", Handler: global.Router}

	log.Fatalf("----------- %v ------------", s.ListenAndServe())

	log.Println("服务退出")
}
