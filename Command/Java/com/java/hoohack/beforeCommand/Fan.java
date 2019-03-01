package com.java.hoohack.beforeCommand;

public class Fan {

    public String getButtonTurnOnName() {
        return "fanTurnOn";
    }

    public String getButtonTurnOffName() {
        return "fanTurnOff";
    }

    public String getButtonSetSpeedUpName() {
        return "fanSpeedUp";
    }

    public String getButtonSetSpeedDownName() {
        return "fanSpeedDown";
    }

    public void on() {
        System.out.println("turn on fan");
    }

    public void off() {
        System.out.println("turn off fan");
    }

    public void speedUp() {
        System.out.println("speed up fan");
    }

    public void speedDown() {
        System.out.println("speed down fan");
    }

}
