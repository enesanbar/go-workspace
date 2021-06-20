package patterns.structural.bridge.example1;

public class Bike extends Vehicle {

    public Bike(Workshop workshop, Workshop workshop2) {
        super(workshop, workshop2);
    }

    @Override
    public void manufacture() {
        System.out.println("Manufacturing a bike...");
        workshop.make();
        workshop2.make();
    }

}
