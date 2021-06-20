package patterns.behavioral.interpreter.example1;

public class Main {

    public InterpreterContext interpreterContext;

    public Main(InterpreterContext interpreterContext) {
        this.interpreterContext = interpreterContext;
    }

    public String interpret(String str) {
        Expression expression = null;

        if (str.toLowerCase().contains("hex")) {
            expression = new IntToHexExpression(Integer.parseInt(str.substring(0, str.indexOf(" "))));
        } else if (str.toLowerCase().contains("bin")) {
            expression = new IntToBinaryExpression(Integer.parseInt(str.substring(0, str.indexOf(" "))));
        } else {
            return str;
        }

        return expression.interpreter(interpreterContext);
    }

    public static void main(String[] args) {
        String first = "13 in Binary";
        String second = "28 in Hex";

        Main interpreter = new Main(new InterpreterContext());
        System.out.println("First: " + interpreter.interpret(first));
        System.out.println("Second: " + interpreter.interpret(second));

    }
}
