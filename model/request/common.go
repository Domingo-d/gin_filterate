/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 15:48
 */

package request

type (
	FilterateReq struct {
		Str string `json:"str"`
	}

	SignInReq struct {
		UserName string `json:"username"`
		PassWord string `json:"password"`
	}

	SignUpReq struct {
		UserName string `json:"username"`
		PassWord string `json:"password"`
	}

	SignOutReq struct {
		UserName string `json:"username"`
	}
)
