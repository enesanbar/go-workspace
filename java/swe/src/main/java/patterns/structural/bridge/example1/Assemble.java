package patterns.structural.bridge.example1;

public class Assemble implements Workshop {

    @Override
    public void make() {
        System.out.print("..also ");
        System.out.println("assembled");
    }

}
