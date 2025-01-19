// models/global_notification_model.go
package models

// GlobalNotificationModel 全局通知表
type GlobalNotificationModel struct {
  Model
  Title   string `gorm:"size:32" json:"title"`
  Icon    string `gorm:"size:256" json:"icon"`  //小图标
  Content string `gorm:"size:64" json:"content"`
  Href    string `gorm:"size:256" json:"href"` // 用户点击消息，然后去进行一个跳转
}
