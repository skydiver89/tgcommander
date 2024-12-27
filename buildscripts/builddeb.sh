#!/bin/bash
dt=`git show --no-patch --format=%ci`
dt=${dt//-/.}
dt=${dt//:/.}
dt=${dt// /-}
dt=`echo $dt | cut -c 1-19`
fpm \
    -s dir -t deb \
    -p ./build/tgcommander-$dt-x86_64.deb \
    -f \
    -n tgcommander \
    --license MIT \
    -v $dt \
    -a x86_64 \
    -m "skydiver89 <maslov140@gmail.com>" \
    --description "Telegram bot command executor" \
    --url "https://github.com/skydiver89/tgcommander" \
    --after-install ./buildscripts/postinst.sh \
    --before-remove ./buildscripts/preremove.sh \
    ./build/tgcommander=/usr/local/bin/tgcommander \
    config.yaml=/etc/tgcommander/config.yaml \
    ./buildscripts/tgcommander.service=/etc/systemd/system/tgcommander.service
exit 0
