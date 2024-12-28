#!/bin/bash
ver=`git describe --tags --abbrev=0`
fpm \
    -s dir -t deb \
    -p ./build/tgcommander-$ver-x86_64.deb \
    -f \
    -n tgcommander \
    --license MIT \
    -v $ver \
    -a x86_64 \
    -m "skydiver89 <maslov140@gmail.com>" \
    --description "Telegram bot command executor" \
    --url "https://github.com/skydiver89/tgcommander" \
    --after-install ./buildscripts/postinst.sh \
    --before-remove ./buildscripts/preremove.sh \
    ./build/tgcommander_linux_x86_64=/usr/local/bin/tgcommander \
    config.yaml=/etc/tgcommander/config.yaml \
    ./buildscripts/tgcommander.service=/etc/systemd/system/tgcommander.service
exit 0
