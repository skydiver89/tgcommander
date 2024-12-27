#!/bin/bash
systemctl stop tgcommander.service
systemctl disable tgcommander.service
systemctl daemon-reload
exit 0
