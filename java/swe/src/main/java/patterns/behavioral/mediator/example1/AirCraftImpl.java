package patterns.behavioral.mediator.example1;

public class AirCraftImpl extends AirCraft {

    public AirCraftImpl(ATCMediator mediator, String name) {
        super(mediator, name);
    }

    @Override
    public void send(String message) {
        System.out.println(this.name + " is sending a message: " + message);
        mediator.sendMessage(message, this);
    }

    @Override
    public void receive(String message) {
        System.out.println(this.name + " received a message: " + message);
    }
}
