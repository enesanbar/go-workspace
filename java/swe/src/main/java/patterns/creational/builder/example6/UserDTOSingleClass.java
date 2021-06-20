package patterns.creational.builder.example6;

import java.time.LocalDate;
import java.time.Period;

//Product class
public class UserDTOSingleClass {

    private String name;

    private String address;

    private String age;

    public String getName() {
        return name;
    }

    public String getAddress() {
        return address;
    }

    public String getAge() {
        return age;
    }

    private void setName(String name) {
        this.name = name;
    }

    private void setAddress(String address) {
        this.address = address;
    }

    private void setAge(String age) {
        this.age = age;
    }

    @Override
    public String toString() {
        return "name=" + name + "\nage=" + age + "\naddress=" + address ;
    }
    //Get builder instance
    public static UserDTOBuilder getBuilder() {
        return new UserDTOBuilder();
    }
    //Builder
    public static class UserDTOBuilder {

        private String firstName;

        private String lastName;

        private String age;

        private String address;

        private UserDTOSingleClass userDTOSingleClass;

        public UserDTOBuilder withFirstName(String fname) {
            this.firstName = fname;
            return this;
        }

        public UserDTOBuilder withLastName(String lname) {
            this.lastName = lname;
            return this;
        }

        public UserDTOBuilder withBirthday(LocalDate date) {
            age = Integer.toString(Period.between(date, LocalDate.now()).getYears());
            return this;
        }

        public UserDTOBuilder withAddress(Address address) {
            this.address = address.getHouseNumber() + " " +address.getStreet()
                    + "\n"+address.getCity()+", "+address.getState()+" "+address.getZipcode();

            return this;
        }

        public UserDTOSingleClass build() {
            this.userDTOSingleClass = new UserDTOSingleClass();
            userDTOSingleClass.setName(firstName+" " + lastName);
            userDTOSingleClass.setAddress(address);
            userDTOSingleClass.setAge(age);
            return this.userDTOSingleClass;
        }

        public UserDTOSingleClass getUserDTO() {
            return this.userDTOSingleClass;
        }
    }
}
