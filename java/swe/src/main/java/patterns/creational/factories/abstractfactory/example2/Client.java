package patterns.creational.factories.abstractfactory.example2;

import org.w3c.dom.Document;

import javax.xml.parsers.DocumentBuilder;
import javax.xml.parsers.DocumentBuilderFactory;
import javax.xml.parsers.ParserConfigurationException;

public class Client {

    public static void main(String[] args) throws ParserConfigurationException {
        DocumentBuilderFactory factory = DocumentBuilderFactory.newInstance();
        System.out.println("Using factory class " + factory.getClass());

        DocumentBuilder builder = factory.newDocumentBuilder();
        System.out.println("Got builder class " + builder.getClass());

        Document document = builder.newDocument();
        System.out.println("Got document class " + document.getClass());

    }
}
