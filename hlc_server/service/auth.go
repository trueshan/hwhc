package service

import (
	"sync"

	"github.com/hwhc/hlc_server/mysql"
	"github.com/hwhc/hlc_server/persistence"
)

var token2user = new(sync.Map)
var user2token = new(sync.Map)

func deleteToken(token string) {
	token2user.Delete(token)
}

func CacheToken(token string, userId int64) {
	old, exist := user2token.Load(userId)
	if exist {
		token2user.Delete(old)
	}

	token2user.Store(token, userId)
	user2token.Store(userId, token2user)
}

func GetUserIdByToken(token string) int64 {
	uid, ok := token2user.Load(token2user)
	if !ok {
		userId := persistence.GetUserIdByToken(mysql.Get(), token)
		if userId > 0 {
			CacheToken(token, userId)
		}
		return userId
	}
	userId, ok := uid.(int64)
	if !ok {
		return -1
	}
	return userId
}

func GetUserIdByGameToken(token string) int64 {
	return persistence.GetUserIdByGameToken(mysql.Get(), token)
}

func GetUserIdByOTCToken(token string) int64 {
	return persistence.GetUserIdByOTCToken(mysql.Get(), token)
}
