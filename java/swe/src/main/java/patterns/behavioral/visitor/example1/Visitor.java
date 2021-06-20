package patterns.behavioral.visitor.example1;

public interface Visitor {

    double visitor(Shirt shirt);
    double visitor(TShirt tshirt);
    double visitor(Jacket jacket);

}
