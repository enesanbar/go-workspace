#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset
set -o xtrace

celery worker -A config --loglevel=DEBUG --concurrency=16 -Ofair
