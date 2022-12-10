package domain

type UserCore struct {
	ID       uint
	Fullname string
	Email    string
	Password string
	Role     uint
	Point    int
	Token    string
}

type Repository interface {
	AddUser(newUser UserCore) (UserCore, error)
	GetUser(existUser UserCore) (UserCore, error)
	GetMyUser(userID uint) (UserCore, error)
	GetAllUser(email string) ([]UserCore, error)
}

type Service interface {
	Register(newUser UserCore) (UserCore, error) //user
	Login(existUser UserCore) (UserCore, error)  //admin & user
	MyProfile(userID uint) (UserCore, error)     //user
	AllUser(email string) ([]UserCore, error)    //admin
}
