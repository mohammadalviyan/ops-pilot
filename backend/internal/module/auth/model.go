package auth

type User struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
	Status       string
	RoleCode     string
}

type AuthUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (u User) ToAuthUser() AuthUser {
	return AuthUser{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Role:  u.RoleCode,
	}
}
