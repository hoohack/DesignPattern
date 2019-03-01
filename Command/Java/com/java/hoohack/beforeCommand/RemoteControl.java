package com.java.hoohack.beforeCommand;

public class RemoteControl {

    private Light light;

    private Fan fan;

    public RemoteControl() {
        light = new Light();
        fan = new Fan();
    }

    public void buttonPressed(String buttonName) {
        if (buttonName.equals(light.getButtonTurnOnName())) {
            light.on();
        } else if (buttonName.equals(light.getButtonTurnOffName())) {
            light.off();
        } else if (buttonName.equals(fan.getButtonTurnOnName())) {
            fan.on();
        } else if (buttonName.equals(fan.getButtonTurnOffName())) {
            fan.off();
        } else if (buttonName.equals(fan.getButtonSetSpeedUpName())) {
            fan.speedUp();
        } else if (buttonName.equals(fan.getButtonSetSpeedDownName())) {
            fan.speedDown();
        }
    }
}
