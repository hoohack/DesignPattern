package com.java.hoohack.hystrix.before;

public class SimpleHystrixRun {

    public static void main(String[] args) {
        SimpleHystrix simpleHystrix = new SimpleHystrix();
        simpleHystrix.call("aServiceGetSingleData", 1);
        simpleHystrix.call("aServiceGetList", 0);

        simpleHystrix.call("bServiceGetSingleData", 1);
        simpleHystrix.call("bServiceGetList", 0);
    }

}
