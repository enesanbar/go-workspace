run-all-tests:
	./manage.py test -- --integration

run-unit-tests:
	pytest -svv --cov=rentomatic --cov-report=term-missing


run-local:
	FLASK_CONFIG="development" flask run

run-app:
	./manage.py compose up -d --build
	./manage.py init-postgres
