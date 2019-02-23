package dal

import (
	"theterriblechild/CloudApp/applications/storageapp/internal/model"
)

type AgentDal interface {
	CreateAgent(*model.Agent) error
	GetAgent(string) (*model.Agent, error)
	GetAgentByOwner(string) ([]*model.Agent, error)
	UpdateAgent(string, *model.Agent) error
	DeleteAgent(string) error
}