import injector
from flask_injector import RequestScope
from sqlalchemy.engine import Connection, Engine
from sqlalchemy.orm import Session


request = injector.ScopeDecorator(RequestScope)


class Db(injector.Module):
    def __init__(self, engine: Engine) -> None:
        self._engine = engine

    @request
    @injector.provider
    def connection(self) -> Connection:
        return self._engine.connect()

    @request
    @injector.provider
    def session(self, connection: Connection) -> Session:
        return Session(bind=connection)
