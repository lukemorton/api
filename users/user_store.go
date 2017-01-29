package users

type UserStore interface {
	CreateStore()
	UserCreator
	UserFinder
}

type UserCreator interface {
	Create(user *User) error
}

type UserFinder interface {
	FindByEmail(email string) (User, error)
}
