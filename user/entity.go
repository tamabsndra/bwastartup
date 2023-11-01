package user

import "time"

type User struct {
	ID         int			         
	Name       string         
	Email      string         
	Password   string         
	Role       string         
	Occupation string         
	Avatar     string         
	CreatedAt  time.Time         
	UpdatedAt  time.Time         
}
