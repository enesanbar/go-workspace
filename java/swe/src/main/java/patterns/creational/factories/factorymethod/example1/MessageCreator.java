package patterns.creational.factories.factorymethod.example1;

import patterns.creational.factories.factorymethod.example1.message.Message;

/**
 * This is our abstract "creator". 
 * The abstract method createMessage() has to be implemented by
 * its subclasses.
 */
public abstract class MessageCreator {

	/**
	 * This is called by clients
	 * @return A {@link Message}
	 */
	public Message getMessage() {
		Message msg = createMessage();
		
		msg.addDefaultHeaders();
		msg.encrypt();
		
		return msg;
	}

	/**
	 * Subclasses must provide implementation for this and
	 * return a specific Message subclass
	 * @return A concrete {@link Message}
	 */
	protected abstract Message createMessage();
}
