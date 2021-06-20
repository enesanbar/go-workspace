package patterns.behavioral.template.example1;

public class Main {

    public static void main(String[] args) {

        Game game = new EndlessRunnerGame();
        game.play();

        Game normalGame = new FootballGame();
        normalGame.play();

    }
}
