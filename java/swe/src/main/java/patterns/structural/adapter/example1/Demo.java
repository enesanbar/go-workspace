package patterns.structural.adapter.example1;

public class Demo {

    public void test()
    {
        Square sq = new Square(11);
        Rectangle adapter = new SquareToRectangleAdapter(sq);
    }

}
