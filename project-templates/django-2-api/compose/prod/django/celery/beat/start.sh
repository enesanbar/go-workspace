#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

rm -f './celerybeat.pid'
celery beat -A config --loglevel=INFO