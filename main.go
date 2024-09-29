/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 15:26
 */

package main

import (
	"filterate/initialize"
	"net/http"
)

func main() {
	Router := initialize.Routers()
	s := &http.Server{Addr: "0:0:0:0:9999", Handler: Router}
	s.ListenAndServe()
}