package user

const (
	AppName = "user"
)

func NewUser(req *CreateUserRequest) *User {
	return &User{
		Data: req,
	}
}

func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{
		Page:
	}
}