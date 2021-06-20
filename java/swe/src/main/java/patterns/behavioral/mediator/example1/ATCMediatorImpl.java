package patterns.behavioral.mediator.example1;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class ATCMediatorImpl implements ATCMediator {

    private List<AirCraft> airCrafts = new ArrayList<>();

    public ATCMediatorImpl() {

    }

    @Override
    public void sendMessage(String message, AirCraft airCraft) {
        airCrafts
                .stream()
                .filter(a -> a != airCraft)
                .forEach(a -> a.receive(message));
//        for (AirCraft a : airCrafts) {
//            // message should not be received by the sending aircraft
//            if (a != airCraft) {
//                a.receive(message);
//            }
//        }
    }

    @Override
    public void addAirCrafts(AirCraft ...airCrafts) {
        this.airCrafts.addAll(Arrays.asList(airCrafts));
    }
}
