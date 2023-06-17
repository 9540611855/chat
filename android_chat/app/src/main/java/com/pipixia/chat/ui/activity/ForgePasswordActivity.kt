package com.pipixia.chat.ui.activity

import android.content.Intent
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.text.Editable
import android.text.TextWatcher
import android.view.View
import android.widget.Button
import android.widget.EditText
import android.widget.Toast
import androidx.lifecycle.ViewModelProvider
import com.pipixia.chat.Factory.ForgePasswordViewModelFactory
import com.pipixia.chat.R
import com.pipixia.chat.view.ForgePasswordModel

class ForgePasswordActivity : AppCompatActivity() {

    private var editCode: EditText? = null
    private var editEmail: EditText? = null
    private var editPassword: EditText? = null
    private var buttonSendCode: Button? = null
    private var buttonSubmit: Button? = null

    private var viewModel: ForgePasswordModel? = null
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_forge_password)
        viewModel = ViewModelProvider(
            this,
            ForgePasswordViewModelFactory(applicationContext)
        )[ForgePasswordModel::class.java]
        editPassword=findViewById<EditText>(R.id.newPassword)
        editCode = findViewById<EditText>(R.id.code)
        editEmail = findViewById<EditText>(R.id.email)
        buttonSendCode=findViewById<Button>(R.id.sendCode)
        buttonSubmit=findViewById<Button>(R.id.submit)

        editEmail?.addTextChangedListener(object : TextWatcher {
            override fun beforeTextChanged(p0: CharSequence?, p1: Int, p2: Int, p3: Int) {
            }

            override fun onTextChanged(s: CharSequence, start: Int, before: Int, count: Int) {
                viewModel!!.setEmail(s.toString())
            }

            override fun afterTextChanged(p0: Editable?) {
            }
        })

        editPassword?.addTextChangedListener(object : TextWatcher {
            override fun beforeTextChanged(p0: CharSequence?, p1: Int, p2: Int, p3: Int) {
            }

            override fun onTextChanged(s: CharSequence, start: Int, before: Int, count: Int) {
                viewModel!!.setPassword(s.toString())
            }

            override fun afterTextChanged(p0: Editable?) {
            }
        })
        editCode?.addTextChangedListener(object : TextWatcher {
            override fun beforeTextChanged(p0: CharSequence?, p1: Int, p2: Int, p3: Int) {
            }

            override fun onTextChanged(s: CharSequence, start: Int, before: Int, count: Int) {
                viewModel!!.setCode(s.toString())
            }

            override fun afterTextChanged(p0: Editable?) {
            }
        })



        buttonSendCode?.setOnClickListener(View.OnClickListener { v: View? ->
            val (info,State)=viewModel!!.sendCode()
            Toast.makeText(this,info, Toast.LENGTH_SHORT).show()

        })

        buttonSubmit?.setOnClickListener(View.OnClickListener { v: View? ->
            val (info,State)=viewModel!!.editPassword()
            Toast.makeText(this,info, Toast.LENGTH_SHORT).show()
            if(State){
                val resultIntent = Intent()
                setResult(RESULT_OK, resultIntent)
                finish()
            }
        })



    }





}