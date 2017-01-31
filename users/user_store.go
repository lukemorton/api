package users

type UserStore interface {
	CreateStore()
	UserCreator
	UserFinder
	UpdateResetTokenHash(user *User)
	UpdatePasswordHash(user *User)
}

type UserCreator interface {
	Create(user *User) error
}

type UserFinder interface {
	FindByEmail(email string) (User, error)
}

type UserPasswordResetter interface {
	UserFinder
	UpdateResetTokenHash(user *User)
}

type UserPasswordUpdater interface {
	UserFinder
	UpdatePasswordHash(user *User)
}
