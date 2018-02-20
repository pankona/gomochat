package io.pankona.gomochat;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;

import gomochat.*;

public class MainActivity extends AppCompatActivity implements ReceiveMessageListener {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        GomoChatClient client = Gomochat.newGomoChatClient();
        client.addReceiveMessageListener(this);

        try {
            client.connect("127.0.0.1", 8080);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    @Override
    public void onReceiveMessage(String msg) {
        // TODO: implement

    }
}
