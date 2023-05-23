package model

import "capstone/data_type"

const TransactionTableName = "transactions"

var transaction BaseModel = &Transaction{}

type Transaction struct {
	Id         string         `db:"id"`
	CategoryId string         `db:"category_id"`
	WalletId   string         `db:"wallet_id"`
	UserId     string         `db:"user_id"`
	Name       string         `db:"name"`
	Amount     float64        `db:"amount"`
	Date       data_type.Date `db:"date"`
	Timestamp

	Category *Category `db:"-"`
}

func (m Transaction) TableName() string {
	return TransactionTableName
}

func (m Transaction) TableIds() []string {
	return []string{"id"}
}

func (m Transaction) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          m.Id,
		"category_id": m.CategoryId,
		"wallet_id":   m.WalletId,
		"user_id":     m.UserId,
		"name":        m.Name,
		"amount":      m.Amount,
		"date":        m.Date,
		"created_at":  m.CreatedAt,
		"updated_at":  m.UpdatedAt,
	}
}

type TransactionQueryOption struct {
	QueryOption

	CategoryId *string
	UserId     *string
	WalletId   *string

	Phrase *string
}

func (o *TransactionQueryOption) SetDefault() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"*"}
	}

	o.QueryOption.SetDefault()

	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "name", Direction: "asc"},
		}
	}

	o.translateSorts(transaction, o.translateSort)
}
