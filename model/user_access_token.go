package model

import "capstone/data_type"

const UserAccessTokenTableName = "user_access_tokens"

var _ BaseModel = &UserAccessToken{}

type UserAccessToken struct {
	Id           string             `db:"id"`
	UserId       string             `db:"user_id"`
	Revoked      bool               `db:"revoked"`
	ExpiredAt    data_type.DateTime `db:"expired_at"`
	IpAddress    string             `db:"ip_address"`
	Longitude    float64            `db:"longitude"`
	Latitude     float64            `db:"latitude"`
	LocationName string             `db:"location_name"`
	Timestamp
}

func (m UserAccessToken) TableName() string {
	return UserAccessTokenTableName
}

func (m UserAccessToken) TableIds() []string {
	return []string{"id"}
}

func (m UserAccessToken) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            m.Id,
		"user_id":       m.UserId,
		"revoked":       m.Revoked,
		"expired_at":    m.ExpiredAt,
		"ip_address":    m.IpAddress,
		"longitude":     m.Longitude,
		"latitude":      m.Latitude,
		"location_name": m.LocationName,
	}
}
