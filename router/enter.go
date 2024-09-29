/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 15:55
 */

package router

type (
	RouterGroup struct {
		FilterateApi
	}
)

var (
	RouterGroupApp = new(RouterGroup)
)