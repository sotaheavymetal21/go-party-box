// パッケージ名の定義
package repositories

// インターフェース IItemRepository を定義
// FindAll メソッドを持ち、*[]models.Item 型と error 型を返す
type IItemRepository interface {
	FindAll() (*[]models.Item, error)
}

// ItemMemoryRepository 構造体を定義
// items フィールドは models.Item 型のスライス
type ItemMemoryRepository struct {
	items []models.Item
}

// NewItemMemoryRepository 関数を定義
// 引数として models.Item 型のスライスを受け取り、IItemRepository 型を返す
// ItemMemoryRepository 構造体を初期化し、そのポインタを返す
func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

// FindAll メソッドを定義
// ItemMemoryRepository 型のレシーバ (r) を持つ
// *[]models.Item 型と error 型を返す
// レシーバの items フィールドのポインタと nil を返す
func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}
