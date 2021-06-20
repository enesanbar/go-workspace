package patterns.behavioral.observer.example1;

public interface Subject {

    void register(Observer observer);
    void unregister(Observer observer);
    void notifyObservers();
    Object getUpdate(Observer observer);
}
