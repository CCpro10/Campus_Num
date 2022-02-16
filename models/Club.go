package models

import (
	"log"
	"main/conf"
	"time"
)

//社团信息
type ClubInfo struct {
	ID           uint  `gorm:"primarykey"`
	ClubId       int64 //社团注册时的账号
	CreatedAt    time.Time
	AvatarAddr   string `json:"avatar_addr"`                     //社团头像url地址
	Introduction string `form:"introduction"json:"introduction"` //社团简介
	Email        string `form:"email"json:"email"`               //电子邮箱
	ClubName     string `form:"club_name"json:"club_name"`       //社团名称
	Password     string `form:"password"json:"password"`         //密码

}

//存在则返回true
func ExistClub(field string, value string) bool {
	var c ClubInfo
	e := DB.Where(field+" = ?", value).Find(&c).Error
	if e != nil {
		log.Println(e)
	}
	log.Println(c)
	if c.ID != 0 {
		return true
	}
	return false
}

//创建社团信息和默认头像
func CreateClubInfo(clubInfo *ClubInfo) bool {
	//创建时的具体类型一定要明确
	clubInfo.AvatarAddr = conf.Config.Oss.DefaultAvatarUrl
	e := DB.Create(clubInfo).Error
	if e != nil {
		return false
	}
	//创建默认头像
	var avatar Avatar
	avatar.ClubId = clubInfo.ID
	avatar.PictureName = "avatar/default.jfif"
	avatar.PictureAddr = conf.Config.Oss.DefaultAvatarUrl
	e = DB.Create(&avatar).Error
	if e != nil {
		return false
	}
	return true
}