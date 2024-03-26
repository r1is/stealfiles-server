package controller

import (
	"fmt"
	"net/http"

	"stealfiles-server/common"
	"stealfiles-server/utils"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

type Req struct {
	Code string `json:"code"`
}

type Msgdata struct{}

func (m Msgdata) GetMsg(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(http.StatusOK, Resp{Code: -1, Msg: fmt.Sprintf("系统500了...,%v", err)})
		}
	}()

	// {"code":"密文"}
	var req Req
	resp := Resp{}
	// 解析请求中的json数据
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.Code = -1
		resp.Msg = "json error ..."
		ctx.JSON(http.StatusOK, resp)
		return
	}

	// 解密
	timeCode, err := utils.Sm4_d(common.ContectKey, req.Code)
	if err != nil {
		resp.Code = -1
		resp.Msg = fmt.Sprintf("decrypt error: %v", err)
		ctx.JSON(http.StatusOK, resp)
		return
	}

	fmt.Println("timeCode: ", timeCode)

	if !utils.Validate(timeCode) {
		resp.Code = -1
		resp.Msg = "Invalid key"
		ctx.JSON(http.StatusOK, resp)
		return
	}

	sm4Key := common.ContectKey + timeCode
	text, err := utils.Sts()
	if err != nil {
		resp.Code = -1
		resp.Msg = "get sts error ..."
		ctx.JSON(http.StatusOK, resp)
		return
	}

	fmt.Println("text: ", text)

	//加密下发OSS的
	resp.Msg, err = utils.Sm4_e(sm4Key, text)
	if err != nil {
		resp.Code = -1
		resp.Msg = "encrypt error ..."
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp.Code = 0

	ctx.JSON(http.StatusOK, resp)

}
