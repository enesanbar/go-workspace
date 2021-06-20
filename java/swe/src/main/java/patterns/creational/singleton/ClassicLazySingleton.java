package patterns.creational.singleton;


/**
 * The main problem with this singleton is that it is not thread safe
 *
 * If more than one thread tries to call the singleton at the same time,
 * more than one copies will be instantiated
 */
public class ClassicLazySingleton {

    private static ClassicLazySingleton obj;

    // private constructor to force use of
    // getInstance() to create Singleton object
    private ClassicLazySingleton() {
    }

    public static ClassicLazySingleton getInstance() {
        if (obj == null)
            obj = new ClassicLazySingleton();
        return obj;
    }

}
