#!/bin/bash
set -e
wget https://raw.githubusercontent.com/freedomeofdevelopers/fod/master/domains
sed -i '1s/^/+/' domains
sed -e ':a' -e 'N' -e '$!ba' -e 's/\n./\n+./g' domains > domains-new
rm domains
mv  /etc/privoxy/domains  /etc/privoxy/domains-old
mv domains-new /etc/privoxy/domains
systemctl restart privoxy.service
