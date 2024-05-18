package user

type CreateUserRq struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type UserResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}
