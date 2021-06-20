package patterns.behavioral.state.example1;

public class SoldOutState implements State {

    private SodaVendingMachine sodaVendingMachine;

    public SoldOutState(SodaVendingMachine sodaVendingMachine) {
        this.sodaVendingMachine = sodaVendingMachine;
    }

    @Override
    public void insertMoney() {
        System.out.println("Machine sold out");
    }

    @Override
    public void ejectMoney() {
        System.out.println("Insert money first before ejecting");
    }

    @Override
    public void select() {
        System.out.println("No soda available");
    }

    @Override
    public void dispense() {
        System.out.println("Sold out");
    }

    @Override
    public String toString() {
        return "Sold out";
    }
}
