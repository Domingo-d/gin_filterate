package service

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/argon2"
	"net/http"
	"server/global"
	"server/global/common"
	"server/model"
	"server/model/request"
	"server/utils"
)

type (
	UserService struct{}
)

func (userService *UserService) SignUp(info *request.SignUpReq) (int, error) {
	ret, err := global.Redis.HExists(context.TODO(), common.REDIS_KEY_USER_INFO, info.UserName).Result()
	if nil != err {
		global.Logger.Error("玩家注册 redis 查询账号返回错误")
	}

	if ret {
		return http.StatusBadRequest, errors.New("账户已存在")
	}

	salt := make([]byte, 16)
	rand.Read(salt)

	hash := argon2.IDKey([]byte(info.PassWord), salt, 1, 64*1024, 4, 32)
	userInfo := &model.UserInfo{UserName: info.UserName, PasswdHash: hash, Salt: salt}

	err = global.HSet(common.REDIS_KEY_USER_INFO, info.UserName, userInfo)
	if nil != err {
		return http.StatusInternalServerError, errors.New("redis 写入失败")
	}

	return http.StatusOK, err
}

func (userService *UserService) SignIn(info *request.SignInReq) (int, string, error) {
	userInfo, err := global.HGet[model.UserInfo](common.REDIS_KEY_USER_INFO, info.UserName)
	if nil != err || nil == userInfo {
		global.Logger.Error("玩家登录 redis 获取账号信息返回错误", zap.Error(err), zap.String("username", info.UserName))
		return http.StatusBadRequest, "", errors.New("账号存在")
	}

	hash := argon2.IDKey([]byte(info.PassWord), userInfo.Salt, 1, 64*1024, 4, 32)
	if !bytes.Equal(userInfo.PasswdHash, hash) {
		return http.StatusBadRequest, "", errors.New("密码错误")
	}

	token := utils.GenerateToken(info.UserName)
	userInfo.Token = token
	global.HSet(common.REDIS_KEY_USER_INFO, info.UserName, userInfo)

	return http.StatusOK, token, nil
}

func (userService *UserService) SignOut(info *request.SignOutReq) (int, error) {
	userInfo, err := global.HGet[model.UserInfo](common.REDIS_KEY_USER_INFO, info.UserName)
	if nil != err || nil == userInfo {
		global.Logger.Error("玩家登录 redis 获取账号信息返回错误", zap.Error(err), zap.String("username", info.UserName))
		return http.StatusBadRequest, errors.New("账号错误")
	}

	userInfo.Token = ""
	_ = global.HSet(common.REDIS_KEY_USER_INFO, info.UserName, userInfo)

	return http.StatusOK, nil
}
