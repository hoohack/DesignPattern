package com.java.hoohack.hystrix.before;

public class BService {

    public void getSingleData(int param) {
        System.out.println("getSingleData called, param: " + param);
    }

    public void getList() {
        System.out.println("getList called");
    }

}
