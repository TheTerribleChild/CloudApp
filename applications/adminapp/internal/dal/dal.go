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

type IAgentDal interface {
	CreateAgent(agent *Agent) error
	GetAgentByName(accountId string, agentName string) (agent Agent, err error)
	ListAgents(accountId string) (agents []Agent, err error)
	UpdateAgent(agent *Agent) error
	DeleteAgent(agentId string) error
}
