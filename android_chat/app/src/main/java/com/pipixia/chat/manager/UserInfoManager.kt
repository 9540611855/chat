package com.pipixia.chat.manager

data class UserInfoManager(val name: String, val id: Int, val email: String,val jwt:String)

object UserInfoManagerSingleton {
   public var userInfoManager:UserInfoManager?=null
    fun SetUserInfoManager(info: UserInfoManager) {
        userInfoManager = info
    }

}