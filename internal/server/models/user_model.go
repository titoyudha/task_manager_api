package models

type User struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Avatar    string    `json:"avatar"`
	Tasks     []Task    `gorm:"foreignKey:AssignedTo" json:"tasks"`
	Comments  []Comment `gorm:"foreignKey:UserID" json:"comments"`
}
