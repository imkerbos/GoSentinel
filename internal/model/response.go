/**
 * @Auth: Kerbos - 0xkerbos@gmail.com
 * @Repo: https://github.com/imkerbos
 * @Date: 10/11/2023 - 10:20 pm
 * @File: internal/model/response.go
 * @Desc:
 */

package model

// JsonResponse 用于返回通用的JSON响应
type JsonResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
