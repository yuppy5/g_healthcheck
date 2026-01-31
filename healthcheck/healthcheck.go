package healthcheck

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var (
	HealthCheckFile string
)

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	HealthCheckFile = dir + "/" + "healthcheck.html"
}

func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat 获取文件信息
	if err != nil {
		return false
	}
	return true
}

func HealthCheck(c *gin.Context) {
	if exists(HealthCheckFile) {
		c.String(http.StatusOK, "%s", "健康检查文件存在, 检查通过")
	} else {
		c.String(http.StatusServiceUnavailable, "%s", "健康检查文件不存在, 检查不通过")
	}
}
