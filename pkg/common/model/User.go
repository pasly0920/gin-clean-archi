package model

type User struct {
	ID           uint64 `gorm:"primaryKey"`
	Username     string `gorm:"not null"`
	Name         string `gorm:"not null"`
	Nickname     string `gorm:"not null"`
	AvatarID     uint64
	PhoneNumber  string
	StudentYear  int
	StudentGroup string
	MajorCode    int
}

// users 테이블 이름을 명시해주는 코드도 추가합니다.
func (u *User) TableName() string {
	return "users"
}
