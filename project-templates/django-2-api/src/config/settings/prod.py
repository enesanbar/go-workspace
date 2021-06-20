from .base import *

DEBUG = False

# keep the secret key used in production secret!
SECRET_KEY = '4k2fejzgm&4gbtjv9=u(#%whf%c=no8d*cxd54=4g@dd%lf%#n'

ADMINS = (
    ('Enes Anbar', 'enesanbar@gmail.com'),
)

# A list of strings representing the host/domain names that this Django site can serve.
ALLOWED_HOSTS = ['mysite.com', '.mysite.com']

DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.postgresql_psycopg2',
        'HOST': 'postgres',
        'PORT': 5432,
        'NAME': 'mysite',
        'USER': 'mysite',
        'PASSWORD': 'password',
    }
}

# redirect HTTP request to HTTPs ones.
SECURE_SSL_REDIRECT = True


CSRF_COOKIE_SECURE = True
