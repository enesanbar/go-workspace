package patterns.structural.proxy.example1;

public class RealBank implements Bank {
    @Override
    public void withdraw(String clientName) throws Exception {
        System.out.println(clientName + " withdrawing from the ATM...");
    }
}
