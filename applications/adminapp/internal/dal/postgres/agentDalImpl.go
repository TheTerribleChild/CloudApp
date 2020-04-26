package postgres

import (
	"github.com/jmoiron/sqlx"
	"log"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	reflectutil "theterriblechild/CloudApp/tools/utils/reflect"
)

type AgentDalImpl struct {
	DB *sqlx.DB
}

func (instance *AgentDalImpl) CreateAgent(agent *dal.Agent) error {
	fieldValMap := reflectutil.GetTagValueAndFieldValueByTagName(agent, dal.DBTag)
	log.Println(fieldValMap)
	query, args, err := psql.Insert(dal.AgentTable).SetMap(fieldValMap).ToSql()
	result, err := instance.DB.Exec(query, args...)
	log.Println(result, err)
	return err
}

func (instance *AgentDalImpl) GetAgentByName(accountId string, agentName string) (agent dal.Agent, err error) {
	sql, args, err := psql.Select("*").From(dal.AgentTable).Where("account_id = ? and name = ?", accountId, agentName).ToSql()
	agent = dal.Agent{}
	if err = instance.DB.Unsafe().Get(&agent, sql, args...); err != nil {
		log.Println(err)
	}
	return
}

func (instance *AgentDalImpl) ListAgents(accountId string) (agents []dal.Agent, err error) {
	sql, args, err := psql.Select("*").From(dal.AgentTable).Where("account_id = ?", accountId).ToSql()
	if err != nil {
		return
	}
	if rows, err := instance.DB.Queryx(sql, args...); err == nil {
		for rows.Next() {
			var agent dal.Agent
			err = rows.StructScan(&agent)
			agents = append(agents, agent)
		}
	}
	return
}

func (instance *AgentDalImpl) UpdateAgent(agent *dal.Agent) error {
	fieldValMap := reflectutil.GetTagValueAndFieldValueByTagName(agent, dal.DBTag)
	log.Println(fieldValMap)
	sql, args, err := psql.Update(dal.AgentTable).SetMap(fieldValMap).Where("id = ?", agent.ID).ToSql()
	result, err := instance.DB.Exec(sql, args...)
	log.Println(result, err)
	return err
}

func (instance *AgentDalImpl) DeleteAgent(agentId string) error {
	sql, args, err := psql.Delete(dal.AgentTable).Where("id = ?", agentId).ToSql()
	result, err := instance.DB.Exec(sql, args...)
	log.Println(result, err)
	return err
}
