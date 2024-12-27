#!/bin/bash
user=root
while true; do
    read -p "Do you wish to run bot with root privilegies?(yes/no) " yn
    case $yn in
        [Yy]* ) break;;
        [Nn]* ) user=$SUDO_USER ; break;;
        * ) echo "Please answer yes or no.";;
    esac
done
sed -i "s/User=root/User=$user/g" /etc/systemd/system/tgcommander.service
systemctl daemon-reload
exit 0
