package cachestore

import (
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
	"theterriblechild/CloudApp/common/model"
	cacheutil "theterriblechild/CloudApp/tools/utils/cache"
	hashutil "theterriblechild/CloudApp/tools/utils/hash"
)

type CachedAdminDB struct {
	dal.AdminDB
	CacheClient cacheutil.ICacheClient
	TTL         int
}

const (
	userByIDPreFix = "UserByID"
)

func (instance *CachedAdminDB) Close() error {
	return instance.AdminDB.Close()
}

func (instance *CachedAdminDB) GetUserByID(userID string) (user model.User, err error) {
	cachedUser := model.User{}
	cacheKey := getKey(userByIDPreFix, userID)
	err = instance.CacheClient.GetJsonDecompress(cacheKey, &cachedUser)
	if cachedUser.Id == userID {
		user = cachedUser
		return
	}
	user, err = instance.AdminDB.GetUserByID(userID)
	if err == nil {
		instance.CacheClient.SetJsonCompress(cacheKey, user, instance.TTL)
	}
	return
}

func getKey(prefix string, key string) string {
	return "AdminDB_" + hashutil.GetHashString(prefix+"_"+key)
}
