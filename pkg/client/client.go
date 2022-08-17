package client

// User model
type User struct {
	ID   string
	Name string
}

type UserService interface {
	GetUser(id string) (User, error)
}

type userService struct {
	// some important fields here
}

func Init() UserService {
	return &userService{}
}

func (u *userService) GetUser(id string) (User, error) {
	// fetches user from remote service by ID
	return User{
		ID:   id,
		Name: "Name " + id,
	}, nil
}
