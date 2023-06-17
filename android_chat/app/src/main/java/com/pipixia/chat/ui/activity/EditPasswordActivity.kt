package com.pipixia.chat.ui.activity

import android.os.Bundle
import android.text.Editable
import android.text.TextWatcher
import android.view.View
import android.widget.Button
import android.widget.EditText
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.ViewModelProvider
import com.pipixia.chat.Factory.EditPasswordViewModelFactory
import com.pipixia.chat.R
import com.pipixia.chat.view.EditPasswordModel

class EditPasswordActivity : AppCompatActivity() {

    private var editTextPassword: EditText? = null
    private var editTextConfirmPassword: EditText? = null
    private var editPassword: Button? = null
    private var viewModel: EditPasswordModel? = null
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_edit_password)
        viewModel = ViewModelProvider(
            this,
            EditPasswordViewModelFactory(applicationContext)
        )[EditPasswordModel::class.java]
        editTextPassword = findViewById<EditText>(R.id.password)
        editTextConfirmPassword=findViewById<EditText>(R.id.confirmPassword)
        editPassword=findViewById<Button>(R.id.editPassword)
        editTextPassword?.addTextChangedListener(object : TextWatcher {
            override fun beforeTextChanged(p0: CharSequence?, p1: Int, p2: Int, p3: Int) {
            }

            override fun onTextChanged(s: CharSequence, start: Int, before: Int, count: Int) {
                viewModel!!.setPassword(s.toString())
            }

            override fun afterTextChanged(p0: Editable?) {
            }
        })
        editTextConfirmPassword?.addTextChangedListener(object : TextWatcher {
            override fun beforeTextChanged(p0: CharSequence?, p1: Int, p2: Int, p3: Int) {
            }

            override fun onTextChanged(s: CharSequence, start: Int, before: Int, count: Int) {
                viewModel!!.setConfirmPassword(s.toString())
            }

            override fun afterTextChanged(p0: Editable?) {
            }
        })

        editPassword?.setOnClickListener(View.OnClickListener { v: View? ->

            val State=viewModel!!.verifyPassword()
            if(State==true){
                val (info,editPasswordState)=viewModel!!.editPassword()
                Toast.makeText(this,info, Toast.LENGTH_SHORT).show()
                if (editPasswordState){
                    //修改成功 返回相关逻辑
                }
            }
        })

    }





}