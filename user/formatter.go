package user

type UserFormatter struct {
	ID         int			 `json:"id"`       
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	Password   string        `json:"password"`
	Token	   string 		 `json:"token"`
}

func FormatterUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
		Token: token,
	}

	return formatter
}