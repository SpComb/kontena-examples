#!/bin/sh

set -ue

# TODO: handle restarts
/usr/sbin/nginx # daemonizes
/bin/confd -watch
