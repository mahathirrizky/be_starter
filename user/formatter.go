package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

func FormatUser(user UserTable, token string) UserFormatter {
	return UserFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	}
}
