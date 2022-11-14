# config/database.py
import logging

from masoniteorm.connections import ConnectionResolver

DATABASES = {
  "default": "mysql",
  "mysql": {
    "host": "127.0.0.1",
    "driver": "mysql",
    "database": "masonite",
    "user": "root",
    "password": "root",
    "port": 3306,
    "log_queries": False,
    "options": {
      #
    }
  },
  "postgres": {
    "host": "127.0.0.1",
    "driver": "postgres",
    "database": "masonite",
    "user": "root",
    "password": "",
    "port": 5432,
    "log_queries": False,
    "options": {
      #
    }
  },
  "sqlite": {
    "driver": "sqlite",
    "database": "masonite.sqlite3",
  }
}

DB = ConnectionResolver().set_connection_details(DATABASES)
logger = logging.getLogger('masoniteorm.connection.queries')
logger.setLevel(logging.DEBUG)

handler = logging.StreamHandler()
file_handler = logging.FileHandler('queries.log')

logger.addHandler(handler)
logger.addHandler(file_handler)
