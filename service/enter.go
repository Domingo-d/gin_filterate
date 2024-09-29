/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 16:28
 */

package service

type (
	ServiceGroup struct {
		FilterateService
	}
)

var (
	ServiceGroupApp = new(ServiceGroup)
)