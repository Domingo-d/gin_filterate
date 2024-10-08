/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 16:29
 */

package service

import (
	"filterate/global"
	"filterate/model/request"
	"filterate/model/response"
)

type (
	FilterateService struct{}
)

func (receiver *FilterateService) Filter(info *request.FilterateReq) *response.FilterateRes {
	res := &response.FilterateRes{Str: global.AhoCorasick.SearchAndReplace(info.Str)}

	return res
}
