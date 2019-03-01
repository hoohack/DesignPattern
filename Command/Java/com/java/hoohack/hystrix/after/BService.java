package com.java.hoohack.hystrix.after;

public class BService {

    public void getSingleData(int param) {
        System.out.println("BService getSingleData called, param: " + param);
    }

    public void getList() {
        System.out.println("BService getList called");
    }

}
