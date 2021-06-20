package patterns.creational.builder.example4;

public class CodeBuilder {

    private Class theClass = new Class();

    public CodeBuilder(String rootName)
    {
        theClass.name = rootName;
    }

    public CodeBuilder addField(String name, String type)
    {
        theClass.fields.add(new Field(name, type));
        return this;
    }

    @Override
    public String toString() {
        return theClass.toString();
    }

    public static void main(String[] args) {
        CodeBuilder cb = new CodeBuilder("Person")
                .addField("name", "String")
                .addField("age", "int");
        System.out.println(cb.toString());
    }
}
