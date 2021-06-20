#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

celery -A config flower --port=5555