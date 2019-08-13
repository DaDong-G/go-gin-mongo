package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexApi(context *gin.Context) {
	context.String(http.StatusOK, "It works in index")
}