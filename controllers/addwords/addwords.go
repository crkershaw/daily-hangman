package addwords

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {

	addwords := route.Group("/add")
	addwords.GET("/", addwords_page)
}

func addwords_page(c *gin.Context) {
	c.HTML(http.StatusOK, "addwords.html", []string{})
}
