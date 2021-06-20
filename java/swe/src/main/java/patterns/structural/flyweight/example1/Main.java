package patterns.structural.flyweight.example1;

import java.util.Random;

public class Main {

    private static final String[] colors = {"Red", "Blue", "Purple", "Brown"};

    public static void main(String[] args) {
        for (int i = 0; i < 20; i++) {
            Circle circle = (Circle) ShapeFactory.getCircle(getRandomColor());
            circle.setX(getRandomX());
            circle.setY(getRandomY());
            circle.draw();
        }
    }

    private static String getRandomColor() {
        return colors[(int) (Math.random() * colors.length)];
    }

    private static int getRandomX() {
        return new Random().nextInt();
    }

    private static int getRandomY() {
        return new Random().nextInt();
    }

}
