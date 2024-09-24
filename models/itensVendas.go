package models

type ItensVendaBase struct {
	ProdutoID  int     `json:"produto_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	Quantidade int     `json:"quantidade" validate:"nonnil" gorm:"not null"`
	Preco      float64 `json:"preco" validate:"nonnull" gorm:"not null"`
}

type ItensVenda struct {
	ItemID  uint `gorm:primaryKey`
	VendaID int  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	ItensVendaBase
	Produto Produto `gorm:"not null"`
}
