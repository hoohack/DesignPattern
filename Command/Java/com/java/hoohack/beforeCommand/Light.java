package com.java.hoohack.beforeCommand;

public class Light {

    public String getButtonTurnOnName() {
        return "lightTurnOn";
    }

    public String getButtonTurnOffName() {
        return "lightTurnOff";
    }

    public void on() {
        System.out.println("turn on light");
    }

    public void off() {
        System.out.println("turn off light");
    }

}
