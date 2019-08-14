package main

import (
 	"seo-content/router"
)

func main() {
	// 初始化引擎
	//engine := gin.Default()
	router := router.InitRouters()
	router.Run(":8888")
}
