package dto

type CreateProductInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type UpdateProductInput struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateSessionInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateSessionOutput struct {
	AccessToken string `json:"acess_token"`
}
