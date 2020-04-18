package dal

type IRegistrationDal interface {
	RegisterAccountAndUser(account *Account, user *User) error
}

type IUserDal interface {
	CreateUser(user *User) error 
	UpdateUser(user *User) error
	GetUserByEmail(email string) (user User, err error)
	GetUserByID(userID string) (user User, err error)
}

type IAccountDal interface {
	CreateAccount(account *Account) error 
}