package patterns.behavioral.mediator.example1;

public class Main {

    public static void main(String[] args) {
        ATCMediator mediator = new ATCMediatorImpl();
        AirCraft boing1 = new AirCraftImpl(mediator, "Boing 1");
        AirCraft helicopter = new AirCraftImpl(mediator, "Helicopter");
        AirCraft boing707 = new AirCraftImpl(mediator, "Boing 707");

        mediator.addAirCrafts(boing1, helicopter, boing707);
        boing1.send("Hello from " + boing1.getName());
    }
}
