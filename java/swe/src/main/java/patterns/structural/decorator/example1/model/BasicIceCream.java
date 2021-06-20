package patterns.structural.decorator.example1.model;


import patterns.structural.decorator.example1.interfaces.IceCream;

public class BasicIceCream implements IceCream {

    public BasicIceCream() {
        System.out.println("Creating a basic Ice-Cream!");
    }

    @Override
    public double cost() {
        return 0.50;
    }
}
