package flag

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/plugins/log_stash"
)

func Makemigrations() {
  var err error
  //使用gorm提供的方法进行表迁移，即根据model生成数据库表
  global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
  global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
  // 生成四张表的表结构
  err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
    AutoMigrate(
      &models.BannerModel{},
      &models.TagModel{},
      &models.MessageModel{},
      &models.AdvertModel{},
      &models.UserModel{},
      &models.CommentModel{},
      &models.ArticleModel{},
      &models.MenuModel{},
      &models.UserCollectModel{},
      &models.MenuBannerModel{},
      &models.FadeBackModel{},
      &models.LoginDataModel{},
      &models.ChatModel{},
      &log_stash.LogStashModel{},
    )
  if err != nil {
    global.Log.Error("[ error ] 生成数据库表结构失败")
    return
  }
  global.Log.Info("[ success ] 生成数据库表结构成功！")
}
