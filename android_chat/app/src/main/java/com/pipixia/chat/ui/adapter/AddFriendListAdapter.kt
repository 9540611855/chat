package com.pipixia.chat.ui.adapter

import android.content.Context
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import android.widget.Toast
import androidx.recyclerview.widget.RecyclerView
import com.pipixia.chat.R
import com.pipixia.chat.data.AddFriendItem
import com.pipixia.chat.view.AddFriendModel

class AddFriendListAdapter(val viewModel: AddFriendModel?, val addFrienditems: MutableList<AddFriendItem>) : RecyclerView.Adapter<AddFriendListAdapter.ViewHolder>() {
    inner class ViewHolder (view: View):RecyclerView.ViewHolder(view){
        val type:TextView=view.findViewById(R.id.type)
        val userName:TextView=view.findViewById(R.id.userName)
    }

    lateinit var context: Context

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ViewHolder {
        context = parent.context
        val inflater = LayoutInflater.from(context)
        val binding = inflater.inflate(R.layout.view_add_friend_item, parent, false)
        val viewHolder=ViewHolder(binding)
        viewHolder.itemView.setOnClickListener{
            val position=viewHolder.adapterPosition
            val friendItem=addFrienditems[position]
            val (info,_)= viewModel!!.add(friendItem.userName)
            Toast.makeText(context,info, Toast.LENGTH_SHORT).show()
        }
        return ViewHolder(binding)
    }

    override fun onBindViewHolder(holder: ViewHolder, position: Int) {
            val friendItem=addFrienditems[position]
            holder.userName.text=friendItem.userName
            holder.userName.isEnabled=false
            if (friendItem.type){
                holder.type.text=context.getString(R.string.friend)
            }else{
                holder.type.text=context.getString(R.string.group)
            }
        holder.type.isEnabled=false
    }

    override fun getItemCount(): Int {
        return addFrienditems.size
    }

}