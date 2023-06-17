package init

import (
	"fmt"
	"pipiChat/configs"
	model "pipiChat/models"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (err error) {
	preString := strings.Join([]string{configs.DBUser, configs.DBPassword}, ":")

	dsn := preString + "@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	configs.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		err = fmt.Errorf("init database err %v", err)
		return err
	}
	configs.DB.AutoMigrate(&model.User{})
	configs.DB.AutoMigrate(&model.UserState{})
	configs.DB.AutoMigrate(&model.GroupInfo{})
	configs.DB.AutoMigrate(&model.ChatInformation{})
	configs.DB.AutoMigrate(&model.Friend{})
	configs.DB.AutoMigrate(&model.GroupUser{})
	configs.DB.AutoMigrate(&model.RegMail{})
	configs.DB.AutoMigrate(&model.FriendGroup{})
	configs.DB.AutoMigrate(&model.ChatInformationFriend{})
	return nil
}
