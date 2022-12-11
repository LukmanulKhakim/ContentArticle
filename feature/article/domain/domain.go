package domain

type ContentCore struct {
	ID        uint
	Article   string
	Point_Art int
	User_ID   uint
	Fullname  string
}

type Repository interface {
	Add(newItem ContentCore) (ContentCore, error)
	GetAll() ([]ContentCore, error)
	GetMy(userID uint, contentID uint) ([]ContentCore, error)
	Edit(point ContentCore, contentID uint) (ContentCore, error)
}

type Service interface {
	Post(newItem ContentCore) (ContentCore, error)
	GetAllContent() ([]ContentCore, error)
	GetMyContent(userID uint, contentID uint) ([]ContentCore, error)
	EditPoint(point ContentCore, contentID uint) (ContentCore, error)
}
