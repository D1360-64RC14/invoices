package interfaces

import "github.com/gin-gonic/gin"

type RouterController interface {
	AttachTo(group *gin.RouterGroup)
}
