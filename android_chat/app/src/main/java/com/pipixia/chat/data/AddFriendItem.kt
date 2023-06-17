package com.pipixia.chat.data

data class AddFriendItem(val userName: String, val type: Boolean = false)
/*
* type:false 的话是群组
* type:true  的话是好友
* */
