#!/usr/bin/env bash
bash /etc/scripts/wait-for-it.sh go-app-db:3306
/usr/bin/supervisord -c /etc/supervisor.conf