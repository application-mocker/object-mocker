package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func MockHttpCode(ctx *gin.Context) {
	code := ctx.Param("code")
	codeInt, err := strconv.ParseInt(code, 10, 0)
	if err != nil {
		ErrorCtx(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(int(codeInt), map[string]interface{}{
		"code": codeInt,
	})

}
