package patterns.structural.proxy.example1;

public class Main {

    public static void main(String[] args) {
        Bank bank = new ProxyBank();
        try {
            bank.withdraw("enes");
            bank.withdraw("James");
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
