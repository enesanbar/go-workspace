package patterns.creational.singleton;


/**
 * Here we have created instance of singleton in static initializer.
 * JVM executes static initializer when the class is loaded and
 * hence this is guaranteed to be thread safe.
 *
 * Use this method only when your singleton class is light and
 * is used throughout the execution of your program.
 */
public class ClassicEagerSingleton {

    private static ClassicEagerSingleton obj = new ClassicEagerSingleton();

    private ClassicEagerSingleton() {
    }

    public static ClassicEagerSingleton getInstance() {
        return obj;
    }

}
