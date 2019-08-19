package routers

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//加载日志和中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}
