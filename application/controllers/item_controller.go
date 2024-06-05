package controllers



type IItemController interface {
	FindAll(ctx *gin.Context)
}

type IItemController struct {
	service services.IItemService
}

func NewItemController(service services.IItemService) IItemController {
	return &IItemController{service: service}
}

func (c *NewItemController) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": Unexpected error})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}
