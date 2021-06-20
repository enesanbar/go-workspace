package patterns.behavioral.state.example1;

public class Main {

    public static void main(String[] args) {
        SodaVendingMachine sodaVendingMachine = new SodaVendingMachine(10);
        sodaVendingMachine.insertMoney();
        sodaVendingMachine.selectSoda();
    }
}
