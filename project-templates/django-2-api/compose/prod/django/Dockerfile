FROM python:3.6.2

ENV PYTHONUNBUFFERED 1

# required for the 'python manage.py compilemessage' to work
RUN apt-get update && apt-get install -y gettext libgettextpo-dev

# Requirements have to be pulled and installed here, otherwise caching won't work
COPY ./requirements /requirements
RUN pip install -r /requirements/prod.txt

COPY ./compose/prod/django/uwsgi.ini /

COPY ./compose/prod/django/start.sh /start.sh
RUN sed -i 's/\r//' /start.sh
RUN chmod +x /start.sh

COPY ./compose/prod/django/celery/worker/start.sh /start-celeryworker.sh
RUN sed -i 's/\r//' /start-celeryworker.sh
RUN chmod +x /start-celeryworker.sh

COPY ./compose/prod/django/celery/beat/start.sh /start-celerybeat.sh
RUN sed -i 's/\r//' /start-celerybeat.sh
RUN chmod +x /start-celerybeat.sh

COPY ./compose/prod/django/celery/flower/start.sh /start-celeryflower.sh
RUN sed -i 's/\r//' /start-celeryflower.sh
RUN chmod +x /start-celeryflower.sh

COPY ./src /app
WORKDIR /app
