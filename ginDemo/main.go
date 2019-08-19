package ginDemo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Student struct {
	Name string
	Age int
}
//动态及正则路由
func main() {
	router := gin.Default()
	router.GET("/user/:name",func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK,"Hello %s",name)
	})
	router.GET("/user/:name/*action", func(d *gin.Context) {
		name := d.Param("name")
		action := d.Param("action")
		message := name + " is " + action
		d.String(http.StatusOK,message)
	})
	router.Run(":8080")
}


//路由组
/*func main()  {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/login",hello)
		v1.GET("/submit",hello)
		v1.GET("/read",hello)
	}
	router.Run(":8080")
}

func hello(content *gin.Context) {
	content.String(http.StatusOK,"ok")
}*/

//单个路由中间件
/*func main() {
	router := gin.Default()
	router.GET("/someGet",middleware1,middleware2,handler)
	router.Run(":8080")
}

func handler(c *gin.Context) {
	log.Println("exec handler")
}

func middleware1(c *gin.Context) {
	log.Println("exec middleware1")
	c.Next()
}
func middleware2(c *gin.Context) {
	log.Println("arrive at middleware2")
	// 执行该中间件之前，先跳到流程的下一个方法
	c.Next()
	// 流程中的其他逻辑已经执行完了
	log.Println("exec middleware2")

	//你可以写一些逻辑代码
}*/


//url查询参数,使用Query和有默认参数的DefaultQuery
/*func main() {
	router := gin.Default()
	router.GET("/welcome",func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname","Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK,"Hello %s %s",firstname,lastname)
	})
	router.Run(":8080")
}*/

//表单和Body参数
/*func main() {
	router := gin.Default()
	router.POST("/form_post", func(c *gin.Context) {
		// 获取post过来的message内容
		// 获取的所有参数内容的类型都是 string
		message := c.PostForm("message")
		// 如果不存在，使用第二个当做默认内容
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200,gin.H{
			"status": "posted",
			"message":message,
			"nick":nick,
		})
	})
	router.Run(":8080")
}*/

//上传文件
/*func main() {
	router := gin.Default()
	// 设置文件上传大小 router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// 处理单一的文件上传
	router.POST("/upload", func(c *gin.Context) {
		file,_:=c.FormFile("file")
		log.Println(file.Filename)
		c.String(http.StatusOK,fmt.Sprintf("'%s' uploaded!",file.Filename))
	})
	// 处理多个文件的上传
	router.POST("/uploads", func(b *gin.Context) {
		form, _ := b.MultipartForm()
		// 拿到集合
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
		}
		b.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8080")
}*/
//其他格式的数据
/*一些复杂的场景下，如用户直接 POST一段 json字符串到应用中
，我们需要获取原始数据，这时需要用 c.GetRawData来获取原始字节
 */

/* func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		// 获取原始字节
		d, err := c.GetRawData()
		if err!=nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
		c.String(200, "ok")
	})
	router.Run(":8080")
} */
