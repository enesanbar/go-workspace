package patterns.structural.adapter.example2;

public interface SocketAdapter {
    Volt get120Volt();
    Volt get12Volt();
    Volt get3Volt();
    Volt get1Volt();
}
