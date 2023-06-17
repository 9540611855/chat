package com.pipixia.chat.ui.widget

import android.content.Context
import android.util.AttributeSet
import android.view.View
import android.widget.Button
import android.widget.EditText
import android.widget.ImageView
import android.widget.RelativeLayout
import android.widget.TextView
import com.pipixia.chat.R
import com.pipixia.chat.data.AddFriendItem

class AddFriendListItemView(context: Context?, attrs: AttributeSet? = null) : RelativeLayout(context, attrs) {
    var userName: TextView?=null
    var type:TextView?=null
    init {
        View.inflate(context, R.layout.view_add_friend_item, this)
        type=findViewById<TextView>(R.id.type)
        userName=findViewById<TextView>(R.id.userName)
    }

    fun bindView(addFriendItem: AddFriendItem) {
        if (addFriendItem.type) {
            //好友
            type?.text=context.getString(R.string.friend)


        } else {
            //群组
            type?.text=context.getString(R.string.group)
        }
        userName?.text = addFriendItem.userName
    }

    private fun addFriend(userName: String) {

    }
}