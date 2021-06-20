package patterns.structural.bridge.example1;

public class Make implements Workshop {

    @Override
    public void make() {
        System.out.println("Making...");
    }

}
