package com.pipixia.chat.manager

import java.util.UUID

object TokenManager  {
    var jwtToken: String? = null
    var deviceToken:String?=null
    fun setManagerJwtToken(token: String?) {
        jwtToken = token
    }

    fun getManagerJwtToken(): String? {
        return jwtToken
    }
    fun setManagerDeviceToken(){
        deviceToken  = "your_value4"
    }
    fun getManagerDeviceToken(): String? {
        return deviceToken
    }
}