package com.example.mcclient

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.widget.Button
import android.widget.EditText
import android.widget.TextView
import mcclient.Core

class MainActivity : AppCompatActivity() {
    private lateinit var hostInput: EditText
    private lateinit var pathInput: EditText
    private lateinit var connectButton: Button
    private lateinit var rttLog: TextView
    private val client = Core.NewClient()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        hostInput = findViewById(R.id.hostInput)
        pathInput = findViewById(R.id.pathInput)
        connectButton = findViewById(R.id.connectButton)
        rttLog = findViewById(R.id.rttLog)

        connectButton.setOnClickListener {
            val host = hostInput.text.toString()
            val path = pathInput.text.toString()
            Thread {
                client.Start(host, path)
            }.start()
        }
    }
}
