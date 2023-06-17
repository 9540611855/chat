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

class ForgePasswordModel  (private val context: Context): ViewModel(){

    private val password = MutableLiveData<String>()
    private val email = MutableLiveData<String>()
    private val code = MutableLiveData<String>()


    fun setCode(code: String){
        this.code.value=code
    }
    fun setPassword(password: String) {
        this.password.value = password
    }

    fun setEmail(email: String) {
        this.email.value = email
    }
    //发送验证码逻辑
    fun sendCode():Pair<String?, Boolean> {
        var flag=false
        var rspInfo: String? =context.getString(R.string.server_error)

        do{
            //验证输入
            if(this.email.value==null){
                rspInfo=context.getString(R.string.no_input)
                break
            }
            //todo 验证邮箱格式


            //发送数据包
            val url= context.getString(R.string.url)+context.getString(R.string.send_mail)

            val request = Request.Builder()
                .url(url)
                .post(
                    RequestBody.create(
                        "application/json; charset=utf-8".toMediaTypeOrNull(),
                        JSONObject()
                            .put("email", email.value.toString())
                            .put("verification_code", "")
                            .toString()
                    )
                ).build()

            val latch: CountDownLatch = CountDownLatch(1)
            val client = OkHttpClient()
            client.newCall(request).enqueue(object : Callback {
                override fun onFailure(call: Call, e: IOException) {
                    // 处理请求失败的情况
                    flag=false
                    rspInfo= context.getString(R.string.no_work)
                    latch.countDown()
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

        return  Pair(rspInfo,flag)
    }


    fun editPassword():Pair<String?, Boolean> {
        var flag=false
        var rspInfo: String? =context.getString(R.string.server_error)

        do{
            //验证输入
            if(this.email.value==null||this.password.value==null||this.code.value==null){
                rspInfo=context.getString(R.string.no_input)
                break
            }
            //todo 验证邮箱格式


            //发送数据包
            val url= context.getString(R.string.url)+context.getString(R.string.forge_password)

            val request = Request.Builder()
                .url(url)
                .post(
                    RequestBody.create(
                        "application/json; charset=utf-8".toMediaTypeOrNull(),
                        JSONObject()
                            .put("email", email.value.toString())
                            .put("verification_code",code.value.toString() )
                            .put("password",password.value.toString() )
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

        return  Pair(rspInfo,flag)
    }

}