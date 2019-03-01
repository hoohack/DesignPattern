package com.java.hoohack.useCommand;

public class RemoteControl {

    private Command command;

    public void command(Command command) {
        this.command = command;
    }

    public void buttonPressed() {
        this.command.execute();
    }

}
