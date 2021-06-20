package patterns.behavioral.visitor.example1;

public class Mai {

    public static void main(String[] args) {
        TaxVisitor taxVisitor = new TaxVisitor();

        Jacket jacket = new Jacket(45.99);
        TShirt tShirt = new TShirt(12.49);
        Shirt shirt = new Shirt(22.45);

        System.out.println("Jacket price: " + jacket.accept(taxVisitor));
        System.out.println("T-Shirt price: " + tShirt.accept(taxVisitor));
        System.out.println("Shirt price: " + shirt.accept(taxVisitor));
    }
}
