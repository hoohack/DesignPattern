package com.java.hoohack.hystrix.after;

public class SimpleHystrixRun {

    public static void main(String[] args) {
        AService aService = new AService();

        SimpleHystrix simpleHystrix = new SimpleHystrix();
        simpleHystrix.setSimpleHystrixCommand(new ASerGetSingleDataCommand(aService, 1));
        simpleHystrix.call();

        simpleHystrix.setSimpleHystrixCommand(new ASerGetListCommand(aService));
        simpleHystrix.call();

        BService bService = new BService();
        simpleHystrix.setSimpleHystrixCommand(new BSerGetSingleDataCommand(bService, 2));
        simpleHystrix.call();

        simpleHystrix.setSimpleHystrixCommand(new BSerGetListCommand(bService));
        simpleHystrix.call();
    }

}
