package user_ser

import (
	"errors"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils"
)

const Avatar = "/uploads/avatar/default.png"

func (UserService) CreateUser(userName, nickName, password string, role ctype.Role, email string, ip string) error {
	// 判断用户名是否存在
	var userModel models.UserModel
	count := global.DB.Take(&userModel, "user_name = ?", userName).RowsAffected
	if count!=0 {
		return errors.New("用户名已存在")
	}
	// 对密码进行hash
	hashPwd := utils.HashPwd(password)

	// 头像问题
	// 1. 默认头像
	// 2. 随机选择头像


	//根据ip获取注册用户的地址信息
	addr := utils.GetAddr(ip)
	// 入库
	err := global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     Avatar,
		IP:         ip,
		Addr:       addr,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		return err
	}
	return nil
}