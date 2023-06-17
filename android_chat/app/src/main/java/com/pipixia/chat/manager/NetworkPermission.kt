package com.pipixia.chat.manager
import android.Manifest
import android.content.Context
import android.content.pm.PackageManager
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.ActivityCompat

object  NetworkPermission {
    const val PERMISSION_REQUEST_CODE = 123 // 可以自定义一个请求码

    fun hasNetworkPermission(context: Context): Boolean {
        return ActivityCompat.checkSelfPermission(
            context,
            Manifest.permission.INTERNET
        ) == PackageManager.PERMISSION_GRANTED
    }

    fun requestNetworkPermission(activity: AppCompatActivity) {
        ActivityCompat.requestPermissions(
            activity,
            arrayOf(Manifest.permission.INTERNET),
            PERMISSION_REQUEST_CODE
        )
    }
}