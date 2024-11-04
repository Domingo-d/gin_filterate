/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 16:29
 */

package service

import (
	"server/global"
	"server/model/request"
	"server/model/response"
)

type (
	FilterateService struct{}
)

func (receiver *FilterateService) Filter(info *request.FilterateReq) *response.FilterateRes {
	// *解引用可以获取原对象,创建一个有原对象的临时新引用, 此时原指针可以重新指向新对象
	tmpAho := *global.AhoCorasick
	res := &response.FilterateRes{Str: tmpAho.SearchAndReplace(info.Str)}

	return res
}
