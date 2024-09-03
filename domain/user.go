package domain

type User struct {
	ID       int    `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"-" gorm:"column:password"`
	Photo    string `json:"photo" gorm:"column:photo"`
}

type Login struct {
	Email    string `json:"email" gorm:"column:email"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}
