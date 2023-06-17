package com.pipixia.chat.ui.activity

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.widget.EditText
import android.widget.ImageView
import android.widget.TextView
import androidx.lifecycle.ViewModelProvider
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.pipixia.chat.Factory.AddFriendViewModelFactory
import com.pipixia.chat.R
import com.pipixia.chat.manager.Unti
import com.pipixia.chat.ui.adapter.AddFriendListAdapter
import com.pipixia.chat.view.AddFriendModel

class AddFriendActivity : AppCompatActivity() {
    private var editTextSearchText: EditText? = null
    private var editImageViewSearch: ImageView? = null
    private var viewModel: AddFriendModel? = null
    var headerTitle: TextView? =null
    var recyclerView: RecyclerView? =null
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_add_friend)
        init()

    }

    fun init(){
        initView()
        initText()
        initRecyClerView()
    }
    fun initRecyClerView(){
        recyclerView=findViewById(R.id.recyclerView)
        recyclerView?.apply {
            setHasFixedSize(true)
            layoutManager = LinearLayoutManager(context)
            adapter = viewModel?.let { AddFriendListAdapter(viewModel = viewModel,it.addFriendItemList) }
        }
    }
    fun initView(){
        viewModel = ViewModelProvider(
            this,
            AddFriendViewModelFactory(applicationContext)
        )[AddFriendModel::class.java]

    }


    fun initText(){

        editTextSearchText=findViewById<EditText>(R.id.search_text)
        editImageViewSearch=findViewById<ImageView>(R.id.search)


        editImageViewSearch?.setOnClickListener{search()}
    }
    fun search(){
        Unti.hideKeyboard(this)
        val key = editTextSearchText?.text?.trim().toString()
        viewModel?.setAddText(key)
        viewModel?.search()
    }

}