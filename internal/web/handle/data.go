package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"object-mocker/pkg/tree"
	"object-mocker/utils"
	"strings"
)

type DataHandler struct {
	root *tree.Node
}

func NewHandler(root *tree.Node) *DataHandler {
	utils.Logger.Infof("Inti a handler with root-node: {%s}", root)
	return &DataHandler{
		root: root,
	}
}

func (h *DataHandler) GetData(ctx *gin.Context) {
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

func (h *DataHandler) CreateData(ctx *gin.Context) {
	path := pathValue(ctx)

	var valueData map[string]interface{}
	err := ctx.BindJSON(&valueData)
	if err != nil {
		ErrorCtx(ctx, http.StatusBadRequest, err)
		return
	}

	data := h.root.NewData(path, valueData)
	ctx.JSON(http.StatusOK, data)
}

func (h *DataHandler) DeleteData(ctx *gin.Context) {

	id := ctx.Query("id")
	path := pathValue(ctx)
	data := h.root.DeleteData(path, id)
	if data.Id != id {
		ErrorCtx(ctx, http.StatusNotFound, fmt.Errorf("not found data with id: %s in scope: %s", id, path))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (h *DataHandler) UpdateData(ctx *gin.Context) {
	id := ctx.Query("id")
	path := pathValue(ctx)

	var valueData map[string]interface{}
	err := ctx.BindJSON(&valueData)
	if err != nil {
		ErrorCtx(ctx, http.StatusBadRequest, err)
		return
	}

	data := h.root.UpdateData(path, id, valueData)
	if data.Id != id {
		ErrorCtx(ctx, http.StatusNotFound, fmt.Errorf("not found data with id: %s in scope: %s", id, path))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (h *DataHandler) ListAllNode(ctx *gin.Context) {
	node := h.root.DeepClone()

	ctx.JSON(http.StatusOK, node)
}

func (h *DataHandler) GetNode(ctx *gin.Context) {
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

