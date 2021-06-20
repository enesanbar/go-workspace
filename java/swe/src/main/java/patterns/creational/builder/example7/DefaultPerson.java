package patterns.creational.builder.example7;

public class DefaultPerson implements Person {

    // required
    private final String firstName;
    private final String lastName;

    // optional
    protected String phoneNumber;
    protected String address;
    protected int age;

    public DefaultPerson(String firstName, String lastName) {
        this.firstName = firstName;
        this.lastName = lastName;
    }

    public DefaultPerson(Builder builder) {
        this.firstName = builder.getFirstName();
        this.lastName = builder.getLastName();
    }

    @Override
    public String getFirstName() {
        return firstName;
    }

    @Override
    public String getLastName() {
        return lastName;
    }

    @Override
    public String getPhoneNumber() {
        return null;
    }

    @Override
    public int getAge() {
        return 0;
    }

    @Override
    public String getAddress() {
        return null;
    }

    class Builder extends DefaultPerson {

        public Builder(String firstName, String lastName) {
            super(firstName, lastName);
        }

        public Builder phone(String phone) {
            this.phoneNumber = phone;
            return this;
        }

        public Builder address(String address) {
            this.address = address;
            return this;
        }

        public Builder age(int age) {
            this.age = age;
            return this;
        }

        public Person build() {
            return new DefaultPerson(this);
        }
    }
}
