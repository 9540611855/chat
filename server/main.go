package main

import (
	"fmt"
	"pipiChat/handler/auth"
	"pipiChat/handler/friend"
	"pipiChat/handler/group"
	"pipiChat/handler/search"
	"pipiChat/handler/user"
	init_ "pipiChat/init"

	"github.com/gin-gonic/gin"
)

func main() {
	//init DB
	err := init_.InitDB()
	if err != nil {
		fmt.Println("[*]open mysql fali")
		return
	}
	//register gin func
	r := gin.Default()
	//user handle
	r.POST("/user/edit_password", user.EditPasswordHandle)
	r.POST("/user/forge_password", user.ForgePasswordHandle)
	r.POST("/user/login", user.UserLoginHandle)
	r.POST("/user/register", user.UserRegisterHandle)
	r.POST("/user/replace_token", user.RepalceTokenHandle)
	//group handle
	r.POST("/group/approve_join_request", group.ApproveJoinRequest{}.ApproveJoinRequestHandler)
	r.POST("/group/chat", group.SaveChatRequest{}.SaveChatInformationHandler)
	r.POST("/group/create_group", group.CreateGroupHandler)
	r.POST("/group/join_group", group.JoinGroup{}.JoinGroupHandler)
	r.POST("/group/kickout_user", group.KickoutUserRequest{}.KickoutUserHandler)
	r.POST("/group/leave-group", group.LeaveGroup{}.LeaveGroupHandler)
	//friend handle
	r.POST("/friend/accept_friend", friend.AcceptFriend{}.AcceptFriendHandle)
	r.POST("/friend/add_friend", friend.AddFriend{}.AddFriendHandle)
	r.POST("/friend/block_friend", friend.BlockFriend{}.BlockFriendHandle)
	r.POST("/friend/flash_message", friend.FlashMessage{}.FlashFriendHandle)
	r.POST("/friend/send_message", friend.FriendMessage{}.SendMessageHandle)
	//auth
	r.POST("/auth/send_mail", auth.SendMailHandle)
	r.POST("/auth/heartbeat", auth.HeartbeatHandler)
	r.POST("/auth/captcha_verify", auth.CaptchaVerifyHandle)
	r.POST("/auth/captcha_generate", auth.GenerateCaptchaHandler)
	r.POST("/search/search", search.Search{}.SearchHandler)
	r.Run()
}
