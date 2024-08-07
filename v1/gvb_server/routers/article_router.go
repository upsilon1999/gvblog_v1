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
	   //获取文章标签
	   article.GET("tags",middleware.JwtAuth(),articleApi.ArticleTagListView)
	   //文章更新
	   article.PUT("update",middleware.JwtAuth(), articleApi.ArticleUpdateView)
	   //文章批量删除
	   article.DELETE("remove",middleware.JwtAuth(),articleApi.ArticleRemoveView)
	   //文章列表高亮搜索
	   article.GET("highlist",middleware.JwtAuth(), articleApi.ArticleHighListView)
	   //支持标题、内容简介、内容搜索，标签搜索、排序搜索、分页搜索，但是只高亮标题
	   article.GET("hiagTitle",middleware.JwtAuth(),articleApi.ArticleHighTitleView)
	   //收藏文章,或者取消收藏
	   article.POST("collect",middleware.JwtAuth(),articleApi.ArticleCollCreateView)
	   //获取用户的收藏列表
	   article.GET("collectList",middleware.JwtAuth(),articleApi.ArticleCollListView)
	   //通过id列表删除收藏数据
	   article.DELETE("collect",middleware.JwtAuth(),articleApi.ArticleCollBatchRemoveView)
	   //全文搜索
	   article.GET("fullText",middleware.JwtAuth(),articleApi.FullTextSearchView)

	}
   
  }