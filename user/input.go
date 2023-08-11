package user

type RegisterUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type CheckUserInput struct {
	Email string `json:"email" binding:"required,email"`
}

type CheckUsernameInput struct {
	Username string `json:"username" binding:"required,username"`
}

type FormCreateUserInput struct {
	Name       string `form:"name" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Occupation string `form:"occupation" binding:"required"`
	Password   string `form:"password" binding:"required"`
	Error      error
}

type FormUpdateUserInput struct {
	ID         int
	Name       string `form:"name" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Occupation string `form:"occupation" binding:"required"`
	Error      error
}
