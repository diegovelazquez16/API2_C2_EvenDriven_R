package launch

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	RegisterPagoModule(router)

}
