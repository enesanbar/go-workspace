package patterns.structural.bridge.example1;

public abstract class Vehicle {

    protected Workshop workshop;
    protected Workshop workshop2;

    public Vehicle(Workshop workshop, Workshop workshop2) {
        this.workshop = workshop;
        this.workshop2 = workshop2;
    }

    abstract public void manufacture();

}
