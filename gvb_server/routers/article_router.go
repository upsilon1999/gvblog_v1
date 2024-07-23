package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

//获取siteInfo配置信息
func (router RouterGroup) ArticleRouter() {
	articleApi:= api.ApiGroupApp.ArticleApi
	article := router.Group("article")
	{
	   //添加文章
	   article.POST("create",middleware.JwtAuth(), articleApi.ArticleCreateView)
	   //获取文章列表
	   article.GET("list",middleware.JwtAuth(), articleApi.ArticleListView)
	   //通过id获取文章详情
	   article.GET("detail/:id",middleware.JwtAuth(), articleApi.ArticleDetailView)
	   //通过title获取文章详情
	   article.GET("detail",middleware.JwtAuth(), articleApi.ArticleDetailByTitleView)
	   //文章日历
	   article.GET("calendar",middleware.JwtAuth(), articleApi.ArticleCalendarView)
	}
   
  }