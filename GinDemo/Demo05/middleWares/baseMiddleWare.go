package middleWares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleWares(c *gin.Context) {
	fmt.Println(" Before ")
	fmt.Println(time.Now())
	fmt.Println("Request url: ", c.Request.URL)
	c.Next()
	fmt.Println("After ")

}

func TestMiddleWares(c *gin.Context) {
	c.Set("url", c.Request.URL)
	fmt.Printf("%T", c.Request.URL)
}

func GoroutineTestMiddleWares(c *gin.Context) {
	cpp := c.Copy()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Done ! in  path" + cpp.Request.URL.Path)
	}()
}
