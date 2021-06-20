package patterns.structural.bridge.example1;

public class Car extends Vehicle {

    public Car(Workshop workshop, Workshop workshop2) {
        super(workshop, workshop2);
    }

    @Override
    public void manufacture() {
        System.out.println("Manufacturing a car...");
        workshop.make();
        workshop2.make();
    }

}
