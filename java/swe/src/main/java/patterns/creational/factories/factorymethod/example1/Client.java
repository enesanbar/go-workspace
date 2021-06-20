package patterns.creational.factories.factorymethod.example1;

import patterns.creational.factories.factorymethod.example1.message.Message;

import java.util.AbstractCollection;

public class Client {

	public static void main(String[] args) {
		printMessage(new JSONMessageCreator());
		printMessage(new TextMessageCreator());
	}
	
	public static void printMessage(MessageCreator creator) {
		Message msg = creator.getMessage();
		System.out.println(msg);
	}
}
