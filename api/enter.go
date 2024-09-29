/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 15:37
 */

package api

import "filterate/service"

type ApiGroup struct {
	FilterateApi
}

var (
	ApiGroupApp = new(ApiGroup)

	filterateService = service.ServiceGroupApp.FilterateService
)