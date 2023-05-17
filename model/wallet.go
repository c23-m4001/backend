package model

import "capstone/data_type"

const WalletTableName = "wallets"

var wallet BaseModel = &Wallet{}

type Wallet struct {
	Id          string                   `db:"id"`
	UserId      string                   `db:"user_id"`
	Name        string                   `db:"name"`
	TotalAmount float64                  `db:"total_amount"`
	LogoType    data_type.WalletLogoType `db:"logo_type"`
	Timestamp
}

func (m Wallet) TableName() string {
	return WalletTableName
}

func (m Wallet) TableIds() []string {
	return []string{"id"}
}

func (m Wallet) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":           m.Id,
		"user_id":      m.UserId,
		"name":         m.Name,
		"total_amount": m.TotalAmount,
		"logo_type":    m.LogoType,
		"created_at":   m.CreatedAt,
		"updated_at":   m.UpdatedAt,
	}
}

type WalletQueryOption struct {
	QueryOption
	Phrase *string
	UserId *string
}

func (o *WalletQueryOption) SetDefault() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"*"}
	}

	o.QueryOption.SetDefault()

	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "name", Direction: "asc"},
		}
	}

	o.translateSorts(wallet, o.translateSort)
}
