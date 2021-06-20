package patterns.behavioral.observer.example1;

import java.util.List;

public class EmailTopicSubscriber implements Observer{

    private String name;
    private Subject topic;

    public EmailTopicSubscriber(String name) {
        this.name = name;
    }

    @Override
    public void update() {
        String message = (String) topic.getUpdate(this);
        if (message == null) {
            System.out.println(name + " no new message on this topic");
        } else {
            System.out.println(name + "Retrieving message: " + message);
        }
    }

    @Override
    public void setSubject(Subject subject) {
        this.topic = subject;
    }
}
