package com.pipixia.chat.ui.activity

import android.content.Intent
import android.os.Bundle
import android.text.Editable
import android.text.TextWatcher
import android.view.View
import android.widget.Button
import android.widget.EditText
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.ViewModelProvider
import com.pipixia.chat.Factory.RegisterViewModelFactory
import com.pipixia.chat.R
import com.pipixia.chat.view.RegisterViewModel


class RegisterActivity : AppCompatActivity() {
    private var editTextUserName: EditText? = null
    private var editTextPassword: EditText? = null
    private var editTextMail: EditText? = null
    private var editTextConfirmPassword: EditText? = null
    private var editTextRealName: EditText? = null
    private var buttonRegister: Button? = null
    private var viewModel: RegisterViewModel? = null

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_register)
        viewModel = ViewModelProvider(
            this,
            RegisterViewModelFactory(applicationContext)
        )[RegisterViewModel::class.java]

        editTextUserName = findViewById<EditText>(R.id.userName)
        editTextPassword = findViewById<EditText>(R.id.password)
        editTextMail = findViewById<EditText>(R.id.mail)
        editTextConfirmPassword=findViewById<EditText>(R.id.confirmPassword)
        editTextRealName=findViewById<EditText>(R.id.realname)
        buttonRegister=findViewById<Button>(R.id.register)

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

        editTextMail?.addTextChangedListener((object : TextWatcher {
            override fun beforeTextChanged(p0: CharSequence?, p1: Int, p2: Int, p3: Int) {

            }

            override fun onTextChanged(s: CharSequence, start: Int, before: Int, count: Int) {
                viewModel!!.setMail(s.toString())
            }

            override fun afterTextChanged(p0: Editable?) {
            }
        }))

        editTextConfirmPassword?.addTextChangedListener(object : TextWatcher {
            override fun beforeTextChanged(p0: CharSequence?, p1: Int, p2: Int, p3: Int) {
            }

            override fun onTextChanged(s: CharSequence, start: Int, before: Int, count: Int) {
                viewModel!!.setConfirmPassword(s.toString())
            }

            override fun afterTextChanged(p0: Editable?) {
            }
        })

        editTextRealName?.addTextChangedListener(object : TextWatcher {
            override fun beforeTextChanged(p0: CharSequence?, p1: Int, p2: Int, p3: Int) {
            }

            override fun onTextChanged(s: CharSequence, start: Int, before: Int, count: Int) {
                viewModel!!.setRealname(s.toString())
            }

            override fun afterTextChanged(p0: Editable?) {
            }
        })

        buttonRegister?.setOnClickListener(View.OnClickListener { v: View? ->

            val (info,registerState)=viewModel!!.register()
            Toast.makeText(this,info, Toast.LENGTH_SHORT).show()
            if(registerState==true){
                val resultIntent = Intent()
                setResult(RESULT_OK, resultIntent)
                finish()
            }


        })


    }





}