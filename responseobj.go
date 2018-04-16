package RestfulRouter

import "github.com/gin-gonic/gin"

type ResponseObj struct {
	Code ResponseCode `json:"code"`
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data"`
}

type ResponseCode int8

const (
	_ ResponseCode = iota
	RSUCCESS
	RERROR
	ROTHER
)

func GinResponseObj(o *ResponseObj) gin.H {
	return gin.H{
		"code": o.Code,
		"msg":  o.Msg,
		"data": o.Data,
	}
}
