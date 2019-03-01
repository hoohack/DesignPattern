package com.java.hoohack.hystrix.after;

public class ASerGetSingleDataCommand implements SimpleHystrixCommand {
    private AService aService;

    private int param;

    public ASerGetSingleDataCommand(AService aService, int param) {
        this.aService = aService;
        this.param = param;
    }

    @Override
    public void execute() {
        try {
            aService.getSingleData(param);
        } catch (Exception e) {
            fallback();
        }
    }

    @Override
    public void fallback() {
        System.out.println("ASerGetSingleDataCommand fallback called");
    }
}
