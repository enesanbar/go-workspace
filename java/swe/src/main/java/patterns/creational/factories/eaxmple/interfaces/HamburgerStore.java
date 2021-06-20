package patterns.creational.factories.eaxmple.interfaces;

import patterns.creational.factories.eaxmple.model.Hamburger;

public abstract class HamburgerStore {


    public Hamburger orderHamburger(String type){
        Hamburger burger;

        //We now user our factory! Not the if statements.
        burger =  createHamburger(type);  //factory.createHamburger(type);

        burger.prepare();
        burger.cook();
        burger.box();

        return burger;

    }

    abstract public Hamburger createHamburger(String type);
}
