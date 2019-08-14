package router

import (
	"github.com/gin-gonic/gin"
	"seo-content/api"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.GET("/", api.IndexApi)
	router.GET("/content-management", api.GetCatagory)
	return router
}

//router := gin.Default()
//router.GET("/someGet", getting)
//router.POST("/somePost", posting)
//router.PUT("/somePut", putting)
//router.DELETE("/someDelete", deleting)
//router.PATCH("/somePatch", patching)
//router.HEAD("/someHead", head)
//router.OPTIONS("/someOptions", options)

// 一 、 获取路径中的参数
//此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
//router.GET("/user/:name", func(c *gin.Context) {
//	name := c.Param("name")
//	c.String(http.StatusOK, "Hello %s", name)
//})

// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
//router.GET("/user/:name/*action", func(c *gin.Context) {
//	name := c.Param("name")
//	action := c.Param("action")
//	message := name + " is " + action
//	c.String(http.StatusOK, message)
//})

//二 、 获取Get参数
//router.GET("/welcome", func(c *gin.Context) {
//	// 此方法可以为默认值
//	firstname := c.DefaultQuery("first","Guest")
//	lastname :=  c.Request.URL.Query().Get("lastname") // 可以简写 c.Query
//	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
//})

// 三、 获取Post参数
//router.POST("/form_post", func(c *gin.Context) {
//	message := c.PostForm("message")
//	nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值
//	fmt.Println("1",message)
//	fmt.Println("2",nick)
//	c.JSON(200, gin.H{
//		"status":  "posted",
//		"message": message,
//		"nick":    nick,
//	})
//})

// 四 / get + post
//router.POST("/post", func(c *gin.Context) {
//	id := c.Query("id")
//	page := c.DefaultQuery("page", "0")
//	name := c.PostForm("name")
//	message := c.PostForm("message")   //postform 需要理解
//
//	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
//})

// 五、上传文件
// 给表单限制上传大小 (默认 32 MiB)
// router.MaxMultipartMemory = 8 << 20  // 8 MiB
//router.POST("/upload", func(c *gin.Context) {
//	// 单文件
//	file, _ := c.FormFile("file")
//	log.Println(file.Filename)

// 上传文件到指定的路径
// c.SaveUploadedFile(file, dst)

//	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
//})

// 六 路由分组
//v1 := router.Group("/v1")
//{
//	v1.POST("/login", loginEndpoint)
//	v1.POST("/submit", submitEndpoint)
//	v1.POST("/read", readEndpoint)
//}
//
//v2 := router.Group("/v2")
//{
//	v2.POST("/login", loginEndpoint)
//	v2.POST("/submit", submitEndpoint)
//	v2.POST("/read", readEndpoint)
//}

// 七 使用中间件
// 创建一个不包含中间件的路由器
//r := gin.New()
//
//// 全局中间件
//// 使用 Logger 中间件
//r.Use(gin.Logger())
//
//// 使用 Recovery 中间件
//r.Use(gin.Recovery())
//
//// 路由添加中间件，可以添加任意多个
//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
//
//// 路由组中添加中间件
//// authorized := r.Group("/", AuthRequired())
//// exactly the same as:
//authorized := r.Group("/")
//// per group middleware! in this case we use the custom created
//// AuthRequired() middleware just in the "authorized" group.
//authorized.Use(AuthRequired())
//{
//	authorized.POST("/login", loginEndpoint)
//	authorized.POST("/submit", submitEndpoint)
//	authorized.POST("/read", readEndpoint)
//
//	// nested group
//	testing := authorized.Group("testing")
//	testing.GET("/analytics", analyticsEndpoint)
//}

//八  创建 日志文件
// 禁用控制台颜色
//gin.DisableConsoleColor()

// 创建记录日志的文件
//f, _ := os.Create("gin.log")
//gin.DefaultWriter = io.MultiWriter(f)

// 如果需要将日志同时写入文件和控制台，请使用以下代码
// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

//router.GET("/ping", func(c *gin.Context) {
//	c.String(200, "pong")
//})

//router.Run(":8080")

//router.GET("/ping", func(c *gin.Context) {
//	c.JSON(200, gin.H{
//		"message": "pong",
//	})
//})
