#!/bin/bash

set -e

/entrypoint.sh echo done >/dev/null

php occ maintenance:install --admin-pass=admin >/dev/null

php -r 'include "config/config.php"; echo json_encode($CONFIG);'