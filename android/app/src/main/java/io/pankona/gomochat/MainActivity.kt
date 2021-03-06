package io.pankona.gomochat

import android.support.v7.app.AppCompatActivity
import android.os.Bundle
import android.text.method.ScrollingMovementMethod
import android.view.View
import android.widget.Button
import android.widget.EditText
import android.widget.TextView

import gomochat.*

class MainActivity : AppCompatActivity(), ReceiveMessageListener, View.OnClickListener {

    lateinit var log: TextView
    lateinit var name: EditText
    lateinit var message: EditText
    lateinit var send: Button
    lateinit var client: Client

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        log = this.findViewById(R.id.log) as TextView
        log.movementMethod = ScrollingMovementMethod()
        name = this.findViewById(R.id.name) as EditText
        message = this.findViewById(R.id.message) as EditText
        send = this.findViewById(R.id.send) as Button
        send.setOnClickListener(this)

        client = Gomochat.newClient()
        client.addReceiveMessageListener(this)
    }

    override fun onResume() {
        super.onResume()

        try {
            client.connect("ws://192.168.0.106:8080/ws")
        } catch (e: Exception) {
            e.printStackTrace()
        }
    }

    override fun onReceiveMessage(msg: String) {
        log.append("$msg\n")
    }

    override fun onClick(v: View) {
        val n = name.text
        val m = message.text

        val msg = "$n: $m"
        if (msg.isEmpty()) {
            return
        }
        try {
            client.sendMessage(msg)
            message.setText("")
        } catch (e: Exception) {
            log.append("send error: "+e.message+"\n")
        }
    }
}
