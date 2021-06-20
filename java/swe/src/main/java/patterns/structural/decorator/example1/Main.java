package patterns.structural.decorator.example1;

import patterns.structural.decorator.example1.interfaces.IceCream;
import patterns.structural.decorator.example1.model.BasicIceCream;
import patterns.structural.decorator.example1.model.MintIceCream;
import patterns.structural.decorator.example1.model.VanillaIceCream;

public class Main {

    public static void main(String[] args) {

        IceCream basicIceCream = new BasicIceCream();
        System.out.println("Basic Ice-cream cost $" + basicIceCream.cost());

        //Add Vanilla to the order

        IceCream vanilla = new VanillaIceCream(basicIceCream); // wrapping vanilla
        System.out.println("Adding Vanilla - cost is: $" + vanilla.cost());


        //Add Mint to the order
        IceCream mint = new MintIceCream(vanilla);
        System.out.println("Adding Mint - cost is: $" + mint.cost());


    }
}
