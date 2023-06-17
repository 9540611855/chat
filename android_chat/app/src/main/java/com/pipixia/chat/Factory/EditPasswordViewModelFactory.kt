package com.pipixia.chat.Factory

import android.content.Context
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import com.pipixia.chat.view.EditPasswordModel

class EditPasswordViewModelFactory (private val context: Context) : ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(EditPasswordModel::class.java)) {
            return EditPasswordModel(context) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}