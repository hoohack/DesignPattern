package com.java.hoohack.beforeCommand;

public class Client {

    public static void main(String[] args) {
        RemoteControl remoteControl = new RemoteControl();
        remoteControl.buttonPressed("lightTurnOn");
        remoteControl.buttonPressed("lightTurnOff");
        remoteControl.buttonPressed("fanTurnOn");
        remoteControl.buttonPressed("fanTurnOff");
        remoteControl.buttonPressed("fanSpeedUp");
        remoteControl.buttonPressed("fanSpeedDown");
    }

}
