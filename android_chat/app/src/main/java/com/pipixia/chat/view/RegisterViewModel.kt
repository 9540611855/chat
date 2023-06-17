package com.pipixia.chat.view

import android.content.Context
import android.util.Log
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import com.pipixia.chat.R
import com.pipixia.chat.manager.TokenManager
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

class RegisterViewModel(private val context: Context): ViewModel()  {

    private val userName = MutableLiveData<String>()
    private val password = MutableLiveData<String>()
    private val confirmPassword = MutableLiveData<String>()
    private val realname = MutableLiveData<String>()
    private val mail = MutableLiveData<String>()

    fun verifyPassword():Boolean{
        return (password.value==confirmPassword.value)
    }

    fun verifyInput():Boolean{
        var userName = this.userName.value
        var password=this.password.value
        var confirmPassword = this.confirmPassword.value
        var realname=this.realname.value
        var mail = this.mail.value
        if (userName == null ||password==null||confirmPassword==null
            || realname==null||mail==null) {
            return false
        }
        return true
    }
    fun register(): Pair<String?, Boolean> {
        if(!verifyInput()){
          return  Pair(context.getString(R.string.no_input),false)
        }
        if(!verifyPassword()){
            return  Pair(context.getString(R.string.verify_password_fail),false)
        }
        val latch: CountDownLatch = CountDownLatch(1)

        val url=context.getString(R.string.url)+context.getString(R.string.register)

        val request = Request.Builder()
            .url(url)
            .post(
                RequestBody.create(
                    "application/json; charset=utf-8".toMediaTypeOrNull(),
                    JSONObject()
                        .put("username", userName.value.toString())
                        .put("password", password.value.toString())
                        .put("realname", realname.value.toString())
                        .put("mail", mail.value.toString())
                        .put("token", TokenManager.deviceToken.toString())
                        .toString()
                )
            ).build()
        var flag=true
        var rspInfo:String?=null
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
                        //TokenManager.setManagerJwtToken(jwtToken)
                       // rspInfo=context.getString(R.string.register_success)
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


    fun setMail(email: String) {
        this.mail.value = email
    }
    fun getMail(): MutableLiveData<String>? {
        return mail
    }
    fun setUserName(email: String) {
        this.userName.value = email
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


    fun setConfirmPassword(confirmPassword: String) {
        this.confirmPassword.value = confirmPassword
    }

    fun setRealname(realname: String) {
       this.realname.value = realname
    }

    fun getConfirmPassword(): MutableLiveData<String>? {
        return confirmPassword
    }

    fun getRealname(): MutableLiveData<String>? {
        return realname
    }
}