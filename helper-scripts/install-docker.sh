#!/bin/sh

add_line_to_file_if_not_exist(){
    if test -e "$2"
    then 
        grep -qxF "$1" "$2" || echo "$1" | sudo tee -a "$2"
    else
        echo "$1" | sudo tee -a "$2"
    fi
}

add_line_to_file_if_not_exist 'Acquire::http::Proxy::download.docker.com "http://fodev.org:8118/";' /etc/apt/apt.conf.d/proxy.conf

# Create the directory for docker installation process
DOCKER_INSTALLER_LOCATION="$HOME/docker-install-helper/"
mkdir -p "$DOCKER_INSTALLER_LOCATION"
cd "$DOCKER_INSTALLER_LOCATION" || exit

# get docker installation script
curl -x http://fodev.org:8118 https://get.docker.com -o get-docker.sh
# Create a curlrc file for proxying and install docker
add_line_to_file_if_not_exist 'proxy="http://fodev.org:8118/"' .curlrc
sudo CURL_HOME="$DOCKER_INSTALLER_LOCATION" sh get-docker.sh
# Setup proxy for docker image pulling
sudo mkdir -p /etc/systemd/system/docker.service.d

HTTP_PROXY_CONTENT="[Service]
Environment='HTTPS_PROXY=http://fodev.org:8118'"
add_line_to_file_if_not_exist "$HTTP_PROXY_CONTENT" /etc/systemd/system/docker.service.d/http-proxy.conf

sudo systemctl daemon-reload
sudo systemctl restart docker

# install the residue installer helper content
rm -r "$DOCKER_INSTALLER_LOCATION"
