package com.pipixia.chat.ui.fragment

import android.content.Intent
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ImageView
import android.widget.TextView
import androidx.swiperefreshlayout.widget.SwipeRefreshLayout
import com.pipixia.chat.R
import com.pipixia.chat.ui.activity.AddFriendActivity
import com.pipixia.chat.ui.activity.RegisterActivity

class ContactFragment : Fragment() {

    var headerTitle: TextView? =null
    var add: ImageView?=null
    var swipeRefreshLayout: SwipeRefreshLayout?=null
    var views:View?=null
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

    }

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        // Inflate the layout for this fragment

        views =inflater.inflate(R.layout.fragment_contact, container, false)
        init()
        return views
    }

    fun init(){
        initHeader()
        initSwipeRefreshLayout()
    }


    fun initHeader(){
        headerTitle= views?.findViewById<TextView>(R.id.headerTitle)
        headerTitle?.text=getString(R.string.contact)


        add= views?.findViewById<ImageView>(R.id.add)
        add?.visibility = View.VISIBLE

        /*
        * 需要去做 添加朋友以及群组==|
        * */
        val intent = Intent(context, AddFriendActivity::class.java)
        add?.setOnClickListener { context?.startActivity(intent) }
    }

    fun initSwipeRefreshLayout() {
        swipeRefreshLayout=view?.findViewById<SwipeRefreshLayout>(R.id.swipeRefreshLayout)
        swipeRefreshLayout?.apply {
            setColorSchemeResources(R.color.qq_blue)
            isRefreshing = true

            /*
            * 需要去做
            * */
            //setOnRefreshListener { presenter.loadContacts() }
        }

    }

}