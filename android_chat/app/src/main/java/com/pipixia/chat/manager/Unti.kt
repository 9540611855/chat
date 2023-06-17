package com.pipixia.chat.manager

import android.app.Activity
import android.content.Context
import android.view.inputmethod.InputMethodManager


object Unti {
    fun hideKeyboard(activity: Activity) {
        val imm = activity.getSystemService(Context.INPUT_METHOD_SERVICE) as InputMethodManager
        val view = activity.currentFocus
        if (view != null) {
            imm.hideSoftInputFromWindow(view.windowToken, 0)
        }
    }
}