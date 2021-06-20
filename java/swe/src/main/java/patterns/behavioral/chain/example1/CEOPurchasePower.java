package patterns.behavioral.chain.example1;

public class CEOPurchasePower extends PurchasePower {

    @Override
    protected double getAllowable() {
        return BASE * 100;
    }

    @Override
    protected String getRole() {
        return "CEO";
    }

}
