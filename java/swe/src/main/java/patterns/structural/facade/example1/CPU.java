package patterns.structural.facade.example1;

public class CPU {
    public void freeze() {
        System.out.println("Computer freezing...");
    }
    public void jump(long position) {
        System.out.println("Jumping to..." + position);
    }
    public void execute() {
        System.out.println("Computer executing commands...");
    }

}
