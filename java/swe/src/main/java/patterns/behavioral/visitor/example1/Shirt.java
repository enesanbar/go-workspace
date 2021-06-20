package patterns.behavioral.visitor.example1;

public class Shirt implements Visitable {

    private double price;

    public Shirt(double price) {
        this.price = price;
    }

    public double getPrice() {
        return price;
    }

    public void setPrice(double price) {
        this.price = price;
    }

    @Override
    public double accept(Visitor visitor) {
        return visitor.visitor(this);
    }

}
