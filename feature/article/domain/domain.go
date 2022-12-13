package domain

type ContentCore struct {
	ID        uint
	Article   string
	Point_Art int
	User_ID   uint
	Fullname  string
}

type User struct {
	ID       uint
	Fullname string
	Point    int
}

type Repository interface {
	Add(newItem ContentCore) (ContentCore, error)
	GetAll() ([]ContentCore, error)
	GetMy(userID uint, key uint) ([]ContentCore, error)
	GetMyAll(userID uint) ([]ContentCore, error)
	Edit(point ContentCore, contentID uint, user uint) (ContentCore, User, error)
}

type Service interface {
	Post(newItem ContentCore) (ContentCore, error)
	GetAllContent() ([]ContentCore, error)
	GetMyContent(userID uint, key uint) ([]ContentCore, error)
	GetMyAllContent(userID uint) ([]ContentCore, error)
	EditPoint(point ContentCore, contentID uint, user uint) (ContentCore, User, error)
}
