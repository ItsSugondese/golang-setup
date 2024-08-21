package global_gin_context

import (
	"github.com/gin-gonic/gin"
)

type globalGinContext struct {
	Context *gin.Context
}

var GlobalGinContext *globalGinContext

func NewGlobalGinContext(context *gin.Context) {
	GlobalGinContext = &globalGinContext{
		Context: context,
	}
}
