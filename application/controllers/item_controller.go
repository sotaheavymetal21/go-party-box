package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sotaheavymetal21/go-party-box/services"
)

// IItemControllerは、アイテムコントローラーのインターフェースを定義します。
type IItemController interface {
	// FindAllは、全てのアイテムを取得し、HTTPレスポンスとして返すメソッドです。
	FindAll(ctx *gin.Context)
}

// ItemControllerは、IItemControllerインターフェースを実装する具体的なコントローラーです。
type ItemController struct {
	service services.IItemService // アイテムサービスへの依存
}

// NewItemControllerは、新しいItemControllerを初期化して返します。
func NewItemController(service services.IItemService) IItemController {
	// serviceを初期値として持つItemControllerを返します。
	return &ItemController{service: service}
}

// FindAllは、全てのアイテムを取得し、HTTPレスポンスとして返すメソッドです。
func (c *ItemController) FindAll(ctx *gin.Context) {
	// サービス層のFindAllメソッドを呼び出して、全てのアイテムを取得します。
	items, err := c.service.FindAll()
	if err != nil {
		// エラーが発生した場合、HTTPステータス500とエラーメッセージを返します。
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	// 成功した場合、HTTPステータス200とアイテムデータを返します。
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}
