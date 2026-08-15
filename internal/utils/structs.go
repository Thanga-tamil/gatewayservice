package utils

import "time"

type Registeruser struct {
	RegisterData string
}

type Data struct {
	UserId 	           string
	NickName           string `json:"name"`
	Email              string
	MobileNumber       string
	ContactPermission  bool
	UserType           string
}

type Users struct {
	UserId     string   `gorm:"uniqueIndex;not null" json:"user_id,omitempty"`
	Nickname   string   `gorm:"not null" json:"nickname,omitempty"`
	IsDeleted   bool    `gorm:"not null" json:"is_deleted,omitempty"`
	ContactPermission   bool    `gorm:"not null" json:"contact_permission,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	EmailId  string   `gorm:"not null" json:"email_id,omitempty"`
	UserType  string   `gorm:"not null" json:"user_type,omitempty"`
	MobileNumber   string   `gorm:"not null" json:"mobile_number,omitempty"`
}
