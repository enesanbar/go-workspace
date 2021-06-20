package solid.ocp.example.specifications;

import solid.ocp.example.Color;
import solid.ocp.example.Product;
import solid.ocp.example.Specification;

public class ColorSpecification implements Specification<Product> {

    private Color color;

    public ColorSpecification(Color color) {
        this.color = color;
    }

    @Override
    public boolean isSatisfied(Product p) {
        return p.color == color;
    }
}