package patterns.behavioral.command.example1;

public class KirbyCharacterReceiver {

    private String name;

    public KirbyCharacterReceiver(String name) {
        this.name = name;
    }

    public void jumpUp() {
        System.out.println(name + " jumping up");
    }

    public void moveDown() {
        System.out.println(name + " moving down");
    }

    public void moveLeft() {
        System.out.println(name + " moving left");
    }

    public void moveRight() {
        System.out.println(name + " moving right");
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
