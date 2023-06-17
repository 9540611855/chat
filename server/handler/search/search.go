package search

import (
	"net/http"
	"pipiChat/configs"
	"pipiChat/handler/auth"
	model "pipiChat/models"

	"github.com/gin-gonic/gin"
)

type Search struct {
	Token    string `form:"token" json:"token" xml:"token" binding:"required"`
	UserId   int64  `form:"user_id" json:"user_id" xml:"user_id" binding:"required"`
	FindText string `form:"find_id" json:"find_id" xml:"find_id" binding:"required"`
}

func (s Search) SearchUser() []string {
	var selectUsers []model.User
	result := configs.DB.Find(&selectUsers, "user_name = ?", s.FindText)

	if result.RowsAffected == 0 || result.Error != nil {
		return nil
	}
	usernames := make([]string, len(selectUsers))
	for i, user := range selectUsers {
		usernames[i] = user.UserName
	}

	return usernames
}

func (s Search) SearchGroup() []string {
	var selectGroup []model.GroupInfo
	result := configs.DB.Find(&selectGroup, "group_name = ?", s.FindText)

	if result.RowsAffected == 0 || result.Error != nil {
		return nil
	}
	groupnames := make([]string, len(selectGroup))
	for i, user := range selectGroup {
		groupnames[i] = user.GroupName
	}

	return groupnames
}

func (s Search) SearchHandler(c *gin.Context) {
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//1.验证token
	if !auth.CheckJwtTokenTime(s.Token, s.UserId) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "jwt token 错误"})
		return
	}
	//2.搜索用户
	userList := s.SearchUser()
	//3.搜索群组
	groupList := s.SearchGroup()
	c.JSON(http.StatusOK, gin.H{"user": userList, "group": groupList})
	//c.JSON(http.StatusOK, gin.H{"success": jwtToken})
}
