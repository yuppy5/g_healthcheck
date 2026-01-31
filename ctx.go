package smartctx

import (
	"sync"

	"github.com/yuppy5/g_healthcheck/healthcheck"

	"github.com/gin-gonic/gin"
)

var onceDo sync.Once

func AddHealthcheckHandler(r *gin.Engine) {
	onceDo.Do(func() {
		// healthcheck handler
		r.GET("/is_running", healthcheck.Process)
		r.HEAD("/is_running", healthcheck.Process)
		r.GET("/healthcheck.html", healthcheck.HealthCheck)
		r.HEAD("/healthcheck.html", healthcheck.HealthCheck)
	})
}
