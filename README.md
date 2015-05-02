# shellsocket
Simple Go TCP Server to bind bash to a socket

### Starting
```
Allans-MacBook-Pro$ ./shellsocket -p=2222
```

### Usage:
```
Allans-MacBook-Pro$ telnet localhost 2222
Trying ::1...
Connected to localhost.
Escape character is '^]'.
# ls /tmp
README.md
shellsocket
shellsocket.go
# date
Sat May  2 14:05:23 EDT 2015
# exit
```
