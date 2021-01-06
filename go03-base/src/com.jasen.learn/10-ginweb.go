package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, JasenWeb")
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/user/list",getUser)
		v1.GET("/user/info",getUser)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}

func getUser(c *gin.Context)  {

}