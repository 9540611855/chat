package com.pipixia.chat.view

import android.content.Context
import android.util.Log
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import com.pipixia.chat.R
import com.pipixia.chat.data.AddFriendItem
import com.pipixia.chat.manager.TokenManager
import com.pipixia.chat.manager.UserInfoManager
import com.pipixia.chat.manager.UserInfoManagerSingleton
import okhttp3.Call
import okhttp3.Callback
import okhttp3.MediaType.Companion.toMediaTypeOrNull
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.RequestBody
import okhttp3.Response
import org.json.JSONException
import org.json.JSONObject
import java.io.IOException
import java.util.concurrent.CountDownLatch
import java.util.concurrent.TimeUnit

class AddFriendModel (private val context: Context): ViewModel() {

    private val addText = MutableLiveData<String>()
    private val  userList=mutableListOf<String>()
    private  val groupList=mutableListOf<String>()
    val addFriendItemList = mutableListOf<AddFriendItem>()
    fun setAddText(text: String) {
        this.addText.value = text
    }
    fun toAddFriendItemList(){
        addFriendItemList.addAll(userList.flatMap { listOf(AddFriendItem(it, type = true)) })
        addFriendItemList.addAll(groupList.flatMap { listOf(AddFriendItem(it, type = false)) })
    }
    fun search(): Pair<String?, Boolean> {
        val jwtToken = UserInfoManagerSingleton.userInfoManager?.jwt
        val userId = UserInfoManagerSingleton.userInfoManager?.id
        var flag = false

        var rspInfo: String? = context.getString(R.string.server_error)
        val url = context.getString(R.string.url) + context.getString(R.string.search)
        do {
            if(jwtToken==null){
                rspInfo=context.getString(R.string.jwt_token_error)
                flag=false
                break
            }
            val request = Request.Builder()
                .url(url)
                .post(
                    RequestBody.create(
                        "application/json; charset=utf-8".toMediaTypeOrNull(),
                        JSONObject()
                            .put("find_id", addText.value.toString())
                            .put("token", jwtToken)
                            .put("user_id",userId)
                            .toString()
                    )
                ).build()
            val latch: CountDownLatch = CountDownLatch(1)
            val client = OkHttpClient()
            client.newCall(request).enqueue(object : Callback {
                override fun onFailure(call: Call, e: IOException) {
                    // 处理请求失败的情况
                    flag=false
                    latch.countDown()
                    rspInfo= context.getString(R.string.no_work)
                }

                override fun onResponse(call: Call, response: Response) {
                    // 保存用户会话信息
                    if(response.code!=200){
                        flag=false
                        try {
                            val jsonObject = JSONObject(response.body!!.string())
                            val error = jsonObject.getString("error")
                            Log.e("LoginViewModel",error)
                            rspInfo=error
                            latch.countDown()
                            // TODO 处理错误信息
                        } catch (e: JSONException) {
                            latch.countDown()
                            e.printStackTrace()
                        } catch (e: IOException) {
                            latch.countDown()
                            e.printStackTrace()
                        }
                    }else{
                        flag=true
                        try {
                            val jsonObject = JSONObject(response.body!!.string())
                            val user = jsonObject.getJSONArray("user")
                            for (i in 0 until user.length()) {
                                val name =  user.getString(i)
                                userList.add(name)
                            }
                            val group = jsonObject.getJSONArray("group")
                            for (i in 0 until group.length()) {
                                val name =  group.getString(i)
                                groupList.add(name)
                            }
                            rspInfo=context.getString(R.string.success_search)
                            latch.countDown()

                            // TODO 处理错误信息
                        } catch (e: JSONException) {
                            latch.countDown()

                            e.printStackTrace()

                        } catch (e: IOException) {
                            latch.countDown()

                            e.printStackTrace()
                        }
                    }

                }
            })


            (latch.await(10, TimeUnit.SECONDS))

        }while (false)
        toAddFriendItemList()
        return Pair(rspInfo,flag)
    }
    fun add(friendName:String):Pair<String?, Boolean>{
        val jwtToken = TokenManager.getManagerJwtToken()
        val userId = UserInfoManagerSingleton.userInfoManager?.id
        var flag = false
        var rspInfo: String? = context.getString(R.string.server_error)
        val url = context.getString(R.string.url) + context.getString(R.string.add_friend)
        do {
            if(jwtToken==null){
                rspInfo=context.getString(R.string.jwt_token_error)
                flag=false
                break
            }
            val request = Request.Builder()
                .url(url)
                .post(
                    RequestBody.create(
                        "application/json; charset=utf-8".toMediaTypeOrNull(),
                        JSONObject()
                            .put("friend_name", friendName)
                            .put("jwt_token", jwtToken)
                            .put("user_id",userId)
                            .put("friend_id",0)
                            .put("add_type",1)
                            .toString()
                    )
                ).build()


        val latch: CountDownLatch = CountDownLatch(1)
        val client = OkHttpClient()
        client.newCall(request).enqueue(object : Callback {
            override fun onFailure(call: Call, e: IOException) {
                // 处理请求失败的情况
                flag=false
                latch.countDown()
                rspInfo= context.getString(R.string.no_work)
            }

            override fun onResponse(call: Call, response: Response) {
                // 保存用户会话信息
                if(response.code!=200){
                    flag=false
                    try {
                        val jsonObject = JSONObject(response.body!!.string())
                        val error = jsonObject.getString("error")
                        Log.e("LoginViewModel",error)
                        rspInfo=error
                        latch.countDown()
                        // TODO 处理错误信息
                    } catch (e: JSONException) {
                        latch.countDown()
                        e.printStackTrace()
                    } catch (e: IOException) {
                        latch.countDown()
                        e.printStackTrace()
                    }
                }else{
                    flag=true
                    try {
                        val jsonObject = JSONObject(response.body!!.string())
                        rspInfo = jsonObject.getString("success")
                        latch.countDown()

                        // TODO 处理错误信息
                    } catch (e: JSONException) {
                        latch.countDown()

                        e.printStackTrace()

                    } catch (e: IOException) {
                        latch.countDown()

                        e.printStackTrace()
                    }
                }

            }
        })


        (latch.await(10, TimeUnit.SECONDS))

    }while (false)
        return Pair(rspInfo,flag)
    }
}