package patterns.creational.factories.eaxmple;

import patterns.creational.factories.eaxmple.interfaces.HamburgerStore;
import patterns.creational.factories.eaxmple.model.CheeseBurger;
import patterns.creational.factories.eaxmple.model.Hamburger;
import patterns.creational.factories.eaxmple.model.JamHamburgerstore;
import patterns.creational.factories.eaxmple.model.MozHamburgerStore;

public class Main {

    public static void main(String[] args) {

        CheeseBurger cheeseBurger = new CheeseBurger();

        HamburgerStore mozambicanBurgerStore = new MozHamburgerStore();
        HamburgerStore jamaicanBurgerStore = new JamHamburgerstore();

        Hamburger hamburger = mozambicanBurgerStore.orderHamburger("cheese");
        System.out.println("Paulo ordered " + hamburger.getName() + "\n" );


        hamburger = jamaicanBurgerStore.orderHamburger("veggie");
        System.out.println("James Bond ordered: " + hamburger.getName() + "\n");




    }
}
