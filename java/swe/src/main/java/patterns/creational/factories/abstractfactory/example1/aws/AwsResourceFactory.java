package patterns.creational.factories.abstractfactory.example1.aws;

import patterns.creational.factories.abstractfactory.example1.Instance;
import patterns.creational.factories.abstractfactory.example1.ResourceFactory;
import patterns.creational.factories.abstractfactory.example1.Storage;

//Factory implementation for Google cloud platform resources
public class AwsResourceFactory implements ResourceFactory {

	@Override
	public Instance createInstance(Instance.Capacity capacity) {
		return new Ec2Instance(capacity);
	}

	@Override
	public Storage createStorage(int capMib) {
		return new S3Storage(capMib);
	}


}
