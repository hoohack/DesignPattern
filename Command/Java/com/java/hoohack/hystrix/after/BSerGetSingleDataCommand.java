package com.java.hoohack.hystrix.after;

public class BSerGetSingleDataCommand implements SimpleHystrixCommand {
    private BService bService;

    private int param;

    public BSerGetSingleDataCommand(BService bService, int param) {
        this.bService = bService;
        this.param = param;
    }

    @Override
    public void execute() {
        try {
            bService.getSingleData(param);
        } catch (Exception e) {
            fallback();
        }
    }

    @Override
    public void fallback() {
        System.out.println("BSerGetSingleDataCommand fallback called");
    }
}
