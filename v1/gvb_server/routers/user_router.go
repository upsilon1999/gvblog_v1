package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

//获取siteInfo配置信息
func (router RouterGroup) UserRouter() {
	UserApi:= api.ApiGroupApp.UserApi
	user := router.Group("user")
	{
	   //邮箱或用户名登录
	   user.POST("emailLogin", UserApi.EmailLoginView)
	   //获取用户列表
	   user.GET("list",middleware.JwtAuth(),UserApi.UserListView)
	   //管理员修改用户权限
	   user.POST("updateRole",middleware.JwtAdmin(),UserApi.UserUpdateRoleView)
	   //用户修改密码
	   user.POST("updatePwd",middleware.JwtAuth(),UserApi.UserUpdatePassword)
	   //用户注销
	   user.POST("logout",middleware.JwtAuth(),UserApi.LogoutView)
	   //创建用户
	   user.POST("create",middleware.JwtAdmin(),UserApi.UserCreateView)
	   //删除用户
	   user.DELETE("remove",middleware.JwtAdmin(),UserApi.UserRemoveView)


	}
   
  }