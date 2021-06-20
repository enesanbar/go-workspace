package patterns.structural.bridge.example1;

public class Main {

    public static void main(String[] args) {
        Vehicle bmw = new Car(new Make(), new Assemble());
        bmw.manufacture();

        Vehicle bmx = new Bike(new Make(), new Assemble());
        bmx.manufacture();
    }

}
