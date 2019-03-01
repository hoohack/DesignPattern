package com.java.hoohack.hystrix.before;

public class SimpleHystrix {

    private AService aService;

    private BService bService;

    public SimpleHystrix() {
        this.aService = new AService();
        this.bService = new BService();
    }

    public void call(String callName, int param) {
        if (callName.equals("aServiceGetSingleData")) {
            try {
                aService.getSingleData(param);
            } catch (Exception e) {
                System.out.println("aService getSingleData exception");
            }
        } else if (callName.equals("aServiceGetList")) {
            try {
                aService.getList();
            } catch (Exception e) {
                System.out.println("aService getList exception");
            }
        } else if (callName.equals("bServiceGetSingleData")) {
            try {
                bService.getSingleData(param);
            } catch (Exception e) {
                System.out.println("bService getSingleData exception");
            }
        } else if (callName.equals("bServiceGetList")) {
            try {
                bService.getList();
            } catch (Exception e) {
                System.out.println("bService getList exception");
            }
        }
    }

}
