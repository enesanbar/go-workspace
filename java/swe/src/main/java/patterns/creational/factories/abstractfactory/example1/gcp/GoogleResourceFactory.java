package patterns.creational.factories.abstractfactory.example1.gcp;

import patterns.creational.factories.abstractfactory.example1.Instance;
import patterns.creational.factories.abstractfactory.example1.ResourceFactory;
import patterns.creational.factories.abstractfactory.example1.Storage;

//Factory implementation for Google cloud platform resources
public class GoogleResourceFactory implements ResourceFactory {

	@Override
	public Instance createInstance(Instance.Capacity capacity) {
		return new GoogleComputeEngineInstance(capacity);
	}

	@Override
	public Storage createStorage(int capMib) {
		return new GoogleCloudStorage(capMib);
	}
	

}
