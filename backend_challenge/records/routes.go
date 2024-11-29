package records

import (
	"challengeFinal/shared"

	"github.com/gin-gonic/gin"
)

func AddRecordsRoutes(r *gin.Engine){
	r.GET("/records", shared.AuthMiddleware(), GetRecordsHandler)
	r.POST("/save-config", shared.AuthMiddleware(), SaveUserConfig)
}