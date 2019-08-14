package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seo-content/database/mongo_connect"
)

func IndexApi(context *gin.Context) {
	context.String(http.StatusOK, "It works in index")
}

func GetCatagory(context *gin.Context) {
	username := mongo_connect.Find()
	context.String(http.StatusOK, "It works in GetCatagory %v", username)
}
