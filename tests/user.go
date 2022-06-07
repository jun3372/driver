package tests

type User struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username" gorm:"column:username;type:varchar;size:255;index"`
	Password  string `json:"password" gorm:"column:password;comment:'密码'"`
	Sex       string `json:"sex" gorm:"column:sex"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}
