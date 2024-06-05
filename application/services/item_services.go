package services

import (
	"github.com/sotaheavymetal21/go-party-box/models"
	"github.com/sotaheavymetal21/go-party-box/repositories"
)

// IItemServiceは、アイテムサービスのインターフェースを定義します。
type IItemService interface {
	// FindAllは、全てのアイテムを取得するメソッドです。
	FindAll() (*[]models.Item, error)
}

// ItemServiceは、IItemServiceインターフェースを実装する具体的なサービスです。
type ItemService struct {
	repository repositories.IItemRepository // アイテムリポジトリへの依存
}

// NewItemServiceは、新しいItemServiceを初期化して返します。
func NewItemService(repository repositories.IItemRepository) IItemService {
	// repositoryを初期値として持つItemServiceを返します。
	return &ItemService{repository: repository}
}

// FindAllは、全てのアイテムを取得するメソッドです。
func (s *ItemService) FindAll() (*[]models.Item, error) {
	// repositoryのFindAllメソッドを呼び出し、結果を返します。
	return s.repository.FindAll()
}
