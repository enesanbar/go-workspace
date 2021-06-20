package patterns.structural.adapter.example2;

public class SocketObjectAdapterImpl implements SocketAdapter {

    private Socket socket = new Socket();

    @Override
    public Volt get120Volt() {
        return socket.getVolt();
    }

    @Override
    public Volt get12Volt() {
        return convertVolt(socket.getVolt(), 10);
    }

    @Override
    public Volt get3Volt() {
        return convertVolt(socket.getVolt(), 40);
    }

    @Override
    public Volt get1Volt() {
        return convertVolt(socket.getVolt(), 120);
    }

    private Volt convertVolt(Volt volt, int i) {
        return new Volt(volt.getVolt() / i);
    }
}
