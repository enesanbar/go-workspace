#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

celery worker -A config --loglevel=INFO --concurrency=16 -Ofair