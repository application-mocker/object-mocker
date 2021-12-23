package handle

import "github.com/gin-gonic/gin"

func ErrorCtx(ctx *gin.Context, code int, err error) {
	errObj := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	ctx.JSON(code, errObj)

}

