package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Result{
		Code:  200,
		Error: "",
		Data:  data,
	})
}
