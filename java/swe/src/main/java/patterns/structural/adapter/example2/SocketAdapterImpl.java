package patterns.structural.adapter.example2;

public class SocketAdapterImpl extends Socket implements SocketAdapter {

    @Override
    public Volt get120Volt() {
        return getVolt();
    }

    @Override
    public Volt get12Volt() {
        return converVolt(getVolt(), 10);
    }

    @Override
    public Volt get3Volt() {
        return converVolt(getVolt(), 40);
    }

    @Override
    public Volt get1Volt() {
        return converVolt(getVolt(), 120);
    }

    private Volt converVolt(Volt volt, int i) {
        return new Volt(volt.getVolt() / i);
    }
}
