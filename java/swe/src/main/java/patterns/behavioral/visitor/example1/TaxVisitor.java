package patterns.behavioral.visitor.example1;

import java.text.DecimalFormat;

public class TaxVisitor implements Visitor {

    DecimalFormat decimalFormat = new DecimalFormat("#,##");

    public TaxVisitor() {
    }

    @Override
    public double visitor(Shirt shirt) {
        System.out.println("Shirt final price with tax:");
        return Double.parseDouble(decimalFormat.format(shirt.getPrice() * .10 + shirt.getPrice()));
    }

    @Override
    public double visitor(TShirt tshirt) {
        System.out.println("TShirt final price with tax:");
        return Double.parseDouble(decimalFormat.format(tshirt.getPrice() * .18 + tshirt.getPrice()));
    }

    @Override
    public double visitor(Jacket jacket) {
        System.out.println("Jacket final price with tax:");
        return Double.parseDouble(decimalFormat.format(jacket.getPrice() * .14 + jacket.getPrice()));
    }

}
