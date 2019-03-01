package com.java.hoohack.hystrix.after;

public class ASerGetListCommand implements SimpleHystrixCommand {
    private AService aService;

    public ASerGetListCommand(AService aService) {
        this.aService = aService;
    }

    @Override
    public void execute() {
        try {
            aService.getList();
        } catch (Exception e) {
            this.fallback();
        }
    }

    @Override
    public void fallback() {
        System.out.println("ASerGetListCommand called fallback");
    }
}
