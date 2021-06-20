package patterns.structural.adapter.example2;

public class Socket {

    public Volt getVolt() {
        return new Volt(120);
    }

}
