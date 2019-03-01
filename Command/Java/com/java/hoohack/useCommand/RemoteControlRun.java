package com.java.hoohack.useCommand;

public class RemoteControlRun {

    public static void main(String[] args) {
        Light light = new Light();


        RemoteControl remoteControl = new RemoteControl();
        remoteControl.command(new LightOnCommand(light));
        remoteControl.buttonPressed();

        remoteControl.command(new LightOffCommand(light));
        remoteControl.buttonPressed();
    }

}
