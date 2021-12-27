package handle

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func MockHttpCode(ctx *gin.Context) {
	code := ctx.Param("code")
	codeInt, err := strconv.ParseInt(code, 10, 0)
	if err != nil {
		ErrorCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(int(codeInt), map[string]interface{}{
		"code": codeInt,
	})
}

func MockHttpBodyByteSize(ctx *gin.Context) {
	size := ctx.Param("size")
	sizeInt, err := strconv.ParseInt(size, 10, 0)
	if err != nil {
		ErrorCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	bf := bytes.Buffer{}
	for i := 0; i < int(sizeInt); i++ {
		bf.WriteByte('0')
	}

	if _, err := ctx.Writer.Write(bf.Bytes()); err != nil {
		ErrorCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
}
