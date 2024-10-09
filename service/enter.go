/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 16:28
 */

package service

type (
	ServiceGroup struct {
		FilterateService
		FileService
	}
)

var (
	ServiceGroupApp = new(ServiceGroup)
)
