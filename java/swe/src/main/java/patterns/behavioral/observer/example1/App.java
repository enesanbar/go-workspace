package patterns.behavioral.observer.example1;

public class App {

    public static void main(String[] args) {
        EmailTopic topic = new EmailTopic();

        Observer firstObserver = new EmailTopicSubscriber("First Onserber");
        Observer secondObserver = new EmailTopicSubscriber("Second Onserber");
        Observer thirdObserver = new EmailTopicSubscriber("Third Onserber");

        topic.register(firstObserver);
        topic.register(secondObserver);
        topic.register(thirdObserver);

        firstObserver.setSubject(topic);
        secondObserver.setSubject(topic);
        thirdObserver.setSubject(topic);

        firstObserver.update();

        topic.postMessage("Hello Observers");
        topic.unregister(thirdObserver);
        topic.postMessage("Hello Observers");

    }

}
