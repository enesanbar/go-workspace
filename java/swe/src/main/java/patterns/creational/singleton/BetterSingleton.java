package patterns.creational.singleton;


/**
 * Singleton is thread safe with "synchronized" keyword
 *
 * We have declared the obj volatile which ensures that
 * multiple threads offer the obj variable correctly
 * when it is being initialized to Singleton instance.
 * This method drastically reduces the overhead of calling the synchronized method every time.
 */
public class BetterSingleton {

    private static volatile BetterSingleton obj;

    private BetterSingleton() {}

    public static BetterSingleton getInstance() {
        if (obj == null) {
            // To make thread safe
            synchronized (BetterSingleton.class) {
                // check again as multiple threads can reach above step
                if (obj==null)
                    obj = new BetterSingleton();
            }
        }
        return obj;
    }

}
