## Description  
This telegram bot can be used to launch commands or scripts and see their output  
## Build  
Simply run:  
```
$ make
```
Bin file will appear in ./build directory  
To build deb-package run  
```
$ make deb
```
## Usage  
```
$ tgcommander -h
Usage: tgcommander [-c configFile] [-h] [-v]
  -c string
        Path to config file (default "config.yaml")
  -h    Show this help
  -v    Show version information
```
## Configuration file  
Configuration file is YAML-file  
```
telegram:
  token: YOUR_TG_BOT_TOKEN
  #list of user IDs who is allowed to use bot
  #can be found in @userinfobot
  users: []
  #message to send to non-allowed users
  declineMessage: "go away"
#list of buttons
buttons:
    #name of button
  - name: name1
    #row number on keyboard
    row: 0
    #command
    command: "ls"
    #arguments
    arguments: ["-l", "-a"]
    #send output of command
    output: true
  - name: name2
    row: 1
    command: "./test.sh"
  - name: name3
    row: 1
    command: "echo"
    arguments: ["Hello world"]
    output: true
```
## Install on debian-based systems with systemd  
Download or build deb-package. Then:  
```
$ sudo apt install deb-package
```
Then you should edit config /etc/tgcommander/config.yaml  
Start bot:
```
$ sudo systemctl start tgcommander.service
```
Stop bot:
```
$ sudo systemctl stop tgcommander.service  
```
Enable autostart:  
```
$ sudo systemctl enable tgcommander.service
```
