package users

type UserStore interface {
	CreateStore()
	UserCreator
	UserPasswordResetter
	UserFinder
}

type UserCreator interface {
	Create(user *User) error
}

type UserFinder interface {
	FindByEmail(email string) (User, error)
}

type UserPasswordResetter interface {
	UpdateResetTokenHashByEmail(email string, token string) error
}
