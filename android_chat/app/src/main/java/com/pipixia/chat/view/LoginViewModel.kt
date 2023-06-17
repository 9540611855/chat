package com.pipixia.chat.view

import android.content.Context
import android.util.Log
import android.widget.Toast
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import com.google.gson.Gson
import com.pipixia.chat.R
import com.pipixia.chat.manager.TokenManager
import com.pipixia.chat.manager.UserInfoManager
import com.pipixia.chat.manager.UserInfoManagerSingleton
import okhttp3.Call
import okhttp3.Callback
import okhttp3.FormBody
import okhttp3.MediaType
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


class LoginViewModel(private val context: Context):ViewModel() {
    private val userName = MutableLiveData<String>()
    private val password = MutableLiveData<String>()

    fun setUserName(userName: String) {
        this.userName.value = userName
    }

    fun setPassword(password: String) {
        this.password.value = password
    }

    fun getUserName(): MutableLiveData<String>? {
        return userName
    }

    fun getPassword(): MutableLiveData<String>? {
        return password
    }
    fun  isValid(): Boolean {

        var userName = this.userName.value
        var password=this.password.value
        if (userName != null &&password!=null) {
                return true

        }
        return false

    }
    fun login(): Pair<String?, Boolean> {
        var flag=false
        var rspInfo: String? =context.getString(R.string.server_error)
        // 发送HTTP请求到后端并验证用户
        var userName = this.userName.value
        var password=this.password.value
        if(!isValid()){
            return Pair(context.getString(R.string.no_input),false)
        }
        if(TokenManager.deviceToken==null)
        {
            TokenManager.setManagerDeviceToken()
        }
        val latch: CountDownLatch = CountDownLatch(1)

        val url= context.getString(R.string.url)+context.getString(R.string.login)
        val request = Request.Builder()
            .url(url)
            .post(
                RequestBody.create(
                    "application/json; charset=utf-8".toMediaTypeOrNull(),
                    JSONObject()
                        .put("username", userName.toString())
                        .put("password", password.toString())
                        .put("token", TokenManager.deviceToken.toString())
                        .toString()
                )
            ).build()

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
                           //val jsonObject = JSONObject(response.body!!.string())
                            //val jwtToken = jsonObject.getString("success")
                            //TokenManager.setManagerJwtToken(jwtToken)
                            val gson = Gson()
                            var user=gson.fromJson(response.body!!.string(), UserInfoManager::class.java)
                            UserInfoManagerSingleton.SetUserInfoManager(user)
                            rspInfo=context.getString(R.string.login_success)
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
    return  Pair(rspInfo,flag)
    }

}