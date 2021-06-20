import socket

from .base import *

# SECURITY WARNING: keep the secret key used in production secret!
SECRET_KEY = '4k2fejzgm&4gbtjv9=u(#%whf%c=no8d*cxd54=4g@dd%lf%#n'

# SECURITY WARNING: don't run with debug turned on in production!
DEBUG = True

# Database
# https://docs.djangoproject.com/en/1.11/ref/settings/#databases

DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.sqlite3',
        'NAME': os.path.join(BASE_DIR, 'db.sqlite3'),
    }
}

# enable the following apps in the local environment
INSTALLED_APPS += (
    'django_extensions',
    'debug_toolbar',
)

# required by debug_toolbar
MIDDLEWARE += ('debug_toolbar.middleware.DebugToolbarMiddleware',)
INTERNAL_IPS = ['127.0.0.1', '172.20.0.1', '172.18.0.1']
ip = socket.gethostbyname(socket.gethostname())
INTERNAL_IPS += [ip[:-1] + '1']

# Output emails to the console.
EMAIL_BACKEND = 'django.core.mail.backends.console.EmailBackend'

GRAPH_MODELS = {
  'all_applications': True,
  'group_models': True,
}

SHELL_PLUS_PRINT_SQL = True
