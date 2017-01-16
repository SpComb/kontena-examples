#!/bin/sh

set -ue

/usr/sbin/nginx # daemonizes
/bin/confd -watch
