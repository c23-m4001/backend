package data_type

type EnumValidator interface {
	IsValid() bool
	GetValidValuesString() string
}
