package patterns.behavioral.template.example1;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class EndlessRunnerGame extends Game {
    @Override
    void initialize() {
        System.out.println("Endless Runner initializing...");
    }

    @Override
    void startPlay() {
        System.out.println("Endless Runner starting...");
        playerWantsNewCharacter();
    }

    @Override
    void endPlay() {
        System.out.println("Endless Runner Ending...");
    }

    @Override
    protected void addNewCharacterToTheGame() {
        System.out.println("Adding new Character to the game");
    }

    public boolean playerWantsNewCharacter() {
        String answer = getUserInput();
        return answer.toLowerCase().startsWith("y");
    }

    private String getUserInput() {
        String answer = null;

        System.out.println("Would you like to add a new character to the game? (y/n)? ");

        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        try {
            answer = in.readLine();

        }catch (IOException ioe) {
            System.out.println("IO Error");
        }
        if (answer == null) {
            return "no";
        }

        return answer;
    }


    //Add more methods...
}
