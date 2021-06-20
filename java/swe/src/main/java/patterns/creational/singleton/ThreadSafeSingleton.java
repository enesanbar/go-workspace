package patterns.creational.singleton;


/**
 * Singleton is thread safe with "synchronized" keyword
 *
 * The main disadvantage of this is method is that
 * using synchronized every time while creating the singleton object is expensive and
 * may decrease the performance of your program.
 * However if performance of getInstance() is not critical for your application,
 * this method provides a clean and simple solution.
 */
public class ThreadSafeSingleton {

    private static ThreadSafeSingleton obj;

    private ThreadSafeSingleton() {
    }

    // Only one thread can execute this at a time
    public static synchronized ThreadSafeSingleton getInstance() {
        if (obj == null)
            obj = new ThreadSafeSingleton();
        return obj;
    }

}
