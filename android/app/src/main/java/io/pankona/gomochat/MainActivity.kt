package io.pankona.gomochat

import android.support.v7.app.AppCompatActivity
import android.os.Bundle

import gomochat.*

class MainActivity : AppCompatActivity(), ReceiveMessageListener {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        val client = Gomochat.newClient()
        client.addReceiveMessageListener(this)

        try {
            client.connect("127.0.0.1", 8080)
        } catch (e: Exception) {
            e.printStackTrace()
        }
    }

    override fun onReceiveMessage(msg: String) {
        // TODO: implement
    }
}
