package middleware

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"deploy-test/domain"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery(c *gin.Context) {
	// catch panic error
	defer func() {
		if err := recover(); err != nil {
			// Check for a broken connection, as it is not really a
			// condition that warrants a panic stack trace.
			if ne, ok := err.(*net.OpError); ok {
				if se, ok := ne.Err.(*os.SyscallError); ok {
					if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
						c.Abort()
						return
					}
				}
			}

			zap.S().Error(fmt.Errorf("Recovery api req: %+v, err: %v", c.Request, err)) // 記錄下來

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    domain.ErrorServer.Code,
				"message": domain.ErrorServer.Message,
				"time":    fmt.Sprintf("%d", time.Now().Unix()),
			})
		}
	}()
	c.Next()
}
