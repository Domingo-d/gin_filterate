/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 16:29
 */

package service

import (
	"filterate/model/request"
	"filterate/model/response"
)

type (
	FilterateService struct{}
)

func (receiver *FilterateService) Filter(info *request.FilterateReq) *response.FilterateRes {

	// filter(info.Str)

	res := &response.FilterateRes{}

	return res
}