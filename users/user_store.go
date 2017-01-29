package users

type UserStore interface {
	CreateStore()
	UserCreator
  UserUpdater
	UserFinder
}

type UserCreator interface {
	Create(user *User) error
}

type UserUpdater interface {
  UpdateResetTokenHashByEmail(email string, token string) error
}

type UserFinder interface {
	FindByEmail(email string) (User, error)
}
