package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"object-mocker/pkg/tree"
	"object-mocker/utils"
	"strings"
)

type Handler struct {
	root *tree.Node
}

func NewHandler(root *tree.Node) *Handler {
	utils.Logger.Infof("Inti a handler with root-node: {%s}", root)
	return &Handler{
		root: root,
	}
}

func (h *Handler) GetData(ctx *gin.Context) {
	path := pathValue(ctx)
	id := ctx.Query("id")
	var data interface{}
	if len(id) != 0 {

		specialData := h.root.DataWithScopes(path, id)
		if specialData.Id != id {
			// not found
			ErrorCtx(ctx, http.StatusNotFound, fmt.Errorf("not found data with id: %s in scope: %s", id, path))
			return
		}
		data = specialData
	} else {
		// list
		data = h.root.DatasWithScopes(path)
	}

	ctx.JSON(http.StatusOK, data)
}

func (h *Handler) CreateData(ctx *gin.Context) {
	var valueData map[string]interface{}

	path := pathValue(ctx)

	err := ctx.BindJSON(&valueData)
	if err != nil {
		ErrorCtx(ctx, http.StatusBadRequest, err)
		return
	}

	data := h.root.NewData(path, valueData)
	ctx.JSON(http.StatusOK, data)
}

func (h *Handler) DeleteData(ctx *gin.Context) {

	id := ctx.Query("id")
	path := pathValue(ctx)
	data := h.root.DeleteData(path, id)
	if data.Id != id {
		ErrorCtx(ctx, http.StatusNotFound, fmt.Errorf("not found data with id: %s in scope: %s", id, path))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (h *Handler) ListAllNode(ctx *gin.Context) {
	node := h.root.DeepClone()

	ctx.JSON(http.StatusOK, node)
}

func (h *Handler) GetNode(ctx *gin.Context) {
	path := pathValue(ctx)

	node := h.root.NodeWithScopes(path)

	ctx.JSON(http.StatusOK, node)
}

func pathValue(ctx *gin.Context) string {
	path := ctx.Param("path")
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}
	return path
}

func ErrorCtx(ctx *gin.Context, code int, err error) {
	errObj := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}

	ctx.JSON(code, errObj)

}
