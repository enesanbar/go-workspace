package patterns.behavioral.mediator.example1;

public interface ATCMediator {

    void sendMessage(String message, AirCraft airCraft);

    void addAirCrafts(AirCraft ...airCraft);
}
