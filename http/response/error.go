package response

import (
	"github.com/gin-gonic/gin"
)

// Error エラーが発生したときのレスポンス
func Error(context *gin.Context, statusCode int, errorCode int, err error) {
	context.JSON(statusCode, gin.H{
		"code":  errorCode,
		"error": err.Error(),
	})
}
