package com.pipixia.chat.ui.activity

import android.content.Intent
import android.content.pm.PackageManager
import android.os.Bundle
import android.text.Editable
import android.text.TextWatcher
import android.view.View
import android.widget.Button
import android.widget.EditText
import android.widget.TextView
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.ViewModelProvider
import com.pipixia.chat.Factory.LoginViewModelFactory
import com.pipixia.chat.R
import com.pipixia.chat.manager.NetworkPermission
import com.pipixia.chat.manager.NetworkPermission.PERMISSION_REQUEST_CODE
import com.pipixia.chat.view.LoginViewModel


class LoginActivity : AppCompatActivity() {
    private var editTextUserName: EditText? = null
    private var editTextPassword: EditText? = null
    private var buttonLogin: Button? = null
    private var buttonRegister:TextView?=null
    private var buttonForgot:TextView?=null
    private val REQUEST_CODE_REGISTER = 1
    private var viewModel: LoginViewModel? = null
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_login)
        viewModel = ViewModelProvider(this, LoginViewModelFactory(applicationContext))[LoginViewModel::class.java]

        editTextUserName = findViewById<EditText>(R.id.userName)
        editTextPassword = findViewById<EditText>(R.id.password)
        buttonLogin = findViewById<Button>(R.id.login)
        buttonRegister=findViewById<TextView>(R.id.newUser)
        buttonForgot=findViewById<TextView>(R.id.forgotPasswd)
        editTextUserName?.addTextChangedListener((object : TextWatcher {
            override fun beforeTextChanged(p0: CharSequence?, p1: Int, p2: Int, p3: Int) {

            }

            override fun onTextChanged(s: CharSequence, start: Int, before: Int, count: Int) {
                viewModel!!.setUserName(s.toString())
            }

            override fun afterTextChanged(p0: Editable?) {
            }
        }))

        editTextPassword?.addTextChangedListener(object : TextWatcher {
            override fun beforeTextChanged(p0: CharSequence?, p1: Int, p2: Int, p3: Int) {
            }

            override fun onTextChanged(s: CharSequence, start: Int, before: Int, count: Int) {
                viewModel!!.setPassword(s.toString())
            }

            override fun afterTextChanged(p0: Editable?) {
            }
        })
        buttonLogin?.setOnClickListener(View.OnClickListener { v: View? ->
            if (!NetworkPermission.hasNetworkPermission(context = this)) {
                NetworkPermission.requestNetworkPermission(this)
            }
            val userName = viewModel!!.getUserName()!!.value
            val password = viewModel!!.getPassword()!!.value
            Toast.makeText(this,userName+" "+password,Toast.LENGTH_SHORT).show()
            val (info,loginState)=viewModel!!.login()
            Toast.makeText(this,info,Toast.LENGTH_SHORT).show()
            if(loginState==true){
                val intent = Intent(this@LoginActivity, MainActivity::class.java)
                startActivityForResult(intent, REQUEST_CODE_REGISTER)
            }
        })
        buttonRegister?.setOnClickListener(View.OnClickListener { v: View? ->
            val intent = Intent(this@LoginActivity, RegisterActivity::class.java)
            startActivityForResult(intent, REQUEST_CODE_REGISTER)
        })
        buttonForgot?.setOnClickListener(View.OnClickListener { v: View? ->
            val intent = Intent(this@LoginActivity, ForgePasswordActivity::class.java)
            startActivityForResult(intent, REQUEST_CODE_REGISTER)
        })


    }
    override fun onRequestPermissionsResult(
        requestCode: Int,
        permissions: Array<String>,
        grantResults: IntArray
    ) {
        if (requestCode == PERMISSION_REQUEST_CODE) {
            if (grantResults.isNotEmpty() && grantResults[0] == PackageManager.PERMISSION_GRANTED) {
                // 权限已授权，可以发起网络请求
                // doSomething()
                Toast.makeText(
                    this,
                    R.string.success_work,
                    Toast.LENGTH_SHORT
                ).show()
            } else {
                // 权限被拒绝
                Toast.makeText(
                    this,
                    R.string.failed_work,
                    Toast.LENGTH_SHORT
                ).show()
                finish()
            }
        } else {
            super.onRequestPermissionsResult(requestCode, permissions, grantResults)
        }
    }
    override fun onActivityResult(requestCode: Int, resultCode: Int, data: Intent?) {
        super.onActivityResult(requestCode, resultCode, data)
        if (requestCode == REQUEST_CODE_REGISTER && resultCode == RESULT_OK) {
            // 注册成功，返回到LoginActivity
            Toast.makeText(this, applicationContext.getString(R.string.result_login), Toast.LENGTH_SHORT).show()
        }
    }


    companion object {
        // Used to load the 'chat' library on application startup.
        init {
            System.loadLibrary("chat")
        }
    }
}
