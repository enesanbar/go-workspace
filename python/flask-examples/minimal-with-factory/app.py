from flask import Flask


def create_app() -> Flask:
    app = Flask(__name__)

    @app.get("/")
    def hello_world():
        return 'catalog world'

    return app
