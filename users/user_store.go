package users

type UserStore interface {
	CreateStore()
	UserCreator
	UserFinder
	UpdateResetTokenHash(user *User) error
	UpdatePasswordHash(user *User) error
}

type UserCreator interface {
	Create(user *User) error
}

type UserFinder interface {
	FindByEmail(email string) (User, error)
}

type UserPasswordResetter interface {
	UserFinder
	UpdateResetTokenHash(user *User) error
}

type UserPasswordUpdater interface {
	UserFinder
	UpdatePasswordHash(user *User) error
}
