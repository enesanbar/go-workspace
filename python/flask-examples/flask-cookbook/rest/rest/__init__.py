import ccy as ccy
from flask import Flask, request, render_template
from flask_injector import FlaskInjector
from flask_sqlalchemy import SQLAlchemy

from main import bootstrap_app
from rest.catalog.views import catalog

db = SQLAlchemy()


def create_app() -> Flask:
    app = Flask(__name__)
    app.secret_key = 'some_random_key'

    app.register_blueprint(catalog)

    # setup dependency injection
    app_context = bootstrap_app()
    FlaskInjector(app, modules=[], injector=app_context.injector)
    app.injector = app_context.injector

    # application level filter
    @app.template_filter('format_currency')
    def format_currency_filter(amount):
        currency_code = ccy.countryccy(request.accept_languages.best[-2:]) or 'USD'
        return '{0} {1}'.format(currency_code, amount)

    @app.errorhandler(404)
    def page_not_found(e):
        return render_template('404.html'), 404

    return app



