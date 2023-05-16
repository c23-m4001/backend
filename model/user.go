package model

const UserTableName = "users"

var _ BaseModel = &User{}

type User struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Timestamp
}

func (m User) TableName() string {
	return UserTableName
}

func (m User) TableIds() []string {
	return []string{"id"}
}

func (m User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         m.Id,
		"name":       m.Name,
		"email":      m.Email,
		"password":   m.Password,
		"created_at": m.CreatedAt,
		"updated_at": m.UpdatedAt,
	}
}
