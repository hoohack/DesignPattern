package com.java.hoohack.hystrix.after;

public interface SimpleHystrixCommand {
    void execute();
    void fallback();
}
