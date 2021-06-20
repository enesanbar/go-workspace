from locust import User, task, between, HttpUser, constant
import logging
logging.info("this log message will go wherever the other locust log messages go")


class MyUser(HttpUser):
    @task
    def my_task(self):
        logging.debug("Executing my_task")
        self.client.get("/")
        print("executing my_task")

    wait_time = constant(0)
