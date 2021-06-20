package solid.ocp.example.specifications;

import solid.ocp.example.Product;
import solid.ocp.example.Size;
import solid.ocp.example.Specification;

public class SizeSpecification implements Specification<Product>
{
    private Size size;

    public SizeSpecification(Size size) {
        this.size = size;
    }

    @Override
    public boolean isSatisfied(Product p) {
        return p.size == size;
    }

}