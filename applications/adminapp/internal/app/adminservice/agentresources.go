package adminservice

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"log"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	adminmodel "theterriblechild/CloudApp/applications/adminapp/model"
	adminvalidator "theterriblechild/CloudApp/applications/adminapp/validator"
	commontype "theterriblechild/CloudApp/common"
	"theterriblechild/CloudApp/common/model"
	"theterriblechild/CloudApp/tools/authentication/cloudappprincipal"
	"theterriblechild/CloudApp/tools/utils/validator"
)

type AgentResource struct {
	agentDal         dal.IAgentDal
	principalManager *cloudappprincipal.PrincipalManager
}

func (instance *AgentResource) RegisterAgent(ctx context.Context, request *adminmodel.CreateAgentRequest) (r *adminmodel.CreateAgentResponse, err error) {
	log.Println(request)
	principal, _ := instance.principalManager.GetPrincipal(ctx)
	err = validator.Validate(
		ctx,
		&cloudappprincipal.PrincipalValidator{
			Principal: principal,
		},
		&adminvalidator.UserAccessValidtor{
			Principal: principal,
			AccountId: request.AccountId,
		},
	)
	if err != nil {
		return
	}
	agent := dal.Agent{ID: uuid.New().String(), AccountID: request.AccountId, Name: request.AgentName}
	instance.agentDal.CreateAgent(&agent)
	r = &adminmodel.CreateAgentResponse{Agent: &model.Agent{Id: agent.ID, AccountId: agent.ID, Name: agent.Name}}
	return r, err
}

func (instance *AgentResource) ListAgents(ctx context.Context, request *adminmodel.ListAgentsRequest) (r *adminmodel.ListAgentsResponse, err error) {
	log.Println(request)
	principal, _ := instance.principalManager.GetPrincipal(ctx)
	err = validator.Validate(
		ctx,
		&cloudappprincipal.PrincipalValidator{
			Principal: principal,
		},
		&adminvalidator.UserAccessValidtor{
			Principal: principal,
			AccountId: request.AccountId,
		},
	)
	if err != nil {
		return
	}
	agents, err := instance.agentDal.ListAgents(request.AccountId)
	log.Println(err)
	r = &adminmodel.ListAgentsResponse{Agents: make([]*model.Agent, 0)}
	for _, agent := range agents {
		r.Agents = append(r.Agents, &model.Agent{Id: agent.ID, Name: agent.Name})
	}
	return r, err
}

func (instance *AgentResource) UpdateAgent(ctx context.Context, request *model.Agent) (r *model.Agent, err error) {
	log.Println(request)
	principal, _ := instance.principalManager.GetPrincipal(ctx)
	err = validator.Validate(
		ctx,
		&cloudappprincipal.PrincipalValidator{
			Principal: principal,
		},
		&adminvalidator.UserAccessValidtor{
			Principal: principal,
			AccountId: request.AccountId,
		},
	)
	if err != nil {
		return
	}
	r = &model.Agent{}
	return r, err
}

func (instance *AgentResource) DeleteAgent(ctx context.Context, request *model.Agent) (r *commontype.Empty, err error) {
	log.Println(request)
	principal, _ := instance.principalManager.GetPrincipal(ctx)
	err = validator.Validate(
		ctx,
		&cloudappprincipal.PrincipalValidator{
			Principal: principal,
		},
		&adminvalidator.UserAccessValidtor{
			Principal: principal,
			AccountId: request.AccountId,
		},
	)
	if err != nil {
		return
	}
	instance.agentDal.DeleteAgent(request.Id)
	r = &commontype.Empty{}
	return r, err
}
