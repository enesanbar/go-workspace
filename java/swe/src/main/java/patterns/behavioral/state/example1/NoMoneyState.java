package patterns.behavioral.state.example1;

public class NoMoneyState implements State {

    private SodaVendingMachine sodaVendingMachine;

    public NoMoneyState(SodaVendingMachine sodaVendingMachine) {
        this.sodaVendingMachine = sodaVendingMachine;
    }

    @Override
    public void insertMoney() {
        System.out.println("You inserted money");
        sodaVendingMachine.setState(sodaVendingMachine.getHasMoneyState());
    }

    @Override
    public void ejectMoney() {
        System.out.println("You haven't inserted any money yet.");
    }

    @Override
    public void select() {
        System.out.println("You selected, but there's no money");
    }

    @Override
    public void dispense() {
        System.out.println("Pay me first");
    }

    @Override
    public String toString() {
        return "Waiting for a dollar...";
    }
}
