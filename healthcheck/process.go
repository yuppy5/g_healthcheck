package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Process(c *gin.Context) {
	c.String(http.StatusOK, "%s", "程序正常运行中")
}
