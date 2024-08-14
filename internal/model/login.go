/**
 * @Auth: Kerbos - 0xkerbos@gmail.com
 * @Repo: https://github.com/imkerbos
 * @Date: 10/11/2023 - 10:20 pm
 * @File: internal/model/login.go
 * @Desc:
 */

package model

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Captcha  string `form:"captcha" binding:"required"`
}
