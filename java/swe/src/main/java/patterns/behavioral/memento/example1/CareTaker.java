package patterns.behavioral.memento.example1;

import java.util.ArrayList;
import java.util.List;

public class CareTaker {

    private List<Memento> states = new ArrayList<>();

    public void addMemento(Memento memento) {
        states.add(memento);
    }

    public Memento getMemento(int index) {
        return states.get(index);
    }

}
