package patterns.creational.singleton;

/**
 * Singleton pattern using lazy initialization holder class. This ensures that, we have a lazy initialization
 * without worrying about synchronization.
 */
public class LazyRegistryHolder {

	private LazyRegistryHolder() {
		System.out.println("In LazyRegistryIODH singleton");
	}
	
	private static class RegistryHolder {
		static LazyRegistryHolder INSTANACE = new LazyRegistryHolder();
	}
	
	public static LazyRegistryHolder getInstance() {
		return RegistryHolder.INSTANACE;
	}
}
