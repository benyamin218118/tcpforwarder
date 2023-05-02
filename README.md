# tcp forwarder

its a simple tcp forwarding tool to forward incoming tcp connections to a remote service at a remote host

## Download and usage :
```
wget https://github.com/benyamin218118/tcpforwarder/raw/main/tcpforwarder
chmod +x ./tcpforwarder

./tcpforwarder -lPort LISTEN_PORT -rHost REMOTE_SERVER_HOST -rPort REMOTE_SERVER_PORT
```

## Simple Usage :
```
forward incoming tcp connection on port 4444 to 1.2.3.4 port 80 :
$ ./tcpforwarder -lPort 4444 -rHost 1.2.3.4 -rPort 80
```

## Forward a Port Range :
```
forward incoming tcp connection on port range 4444-5555 to example.com 4444-5555 :
$ ./tcpforwarder -lPort 4444-5555 -rHost example.com
```
## Detailed Example :

**we want to forward incoming tcp connections from our ir vps to a service running on port 9090 on a usa vps**
- first we need to choose a listen port for the ir vps like 8080 ( we will accept the incoming tcp connections on it )
- our usa vps ip address is 44.55.66.77 and that service port is 9090 as mentioned before ( `you can use domain address instead of the ip too` )
- so we need to replace the variables in `./tcpforwarder -lPort LISTEN_PORT -rHost REMOTE_SERVER_HOST -rPort REMOTE_SERVER_PORT` and run it on the ir vps:

> ./tcpforwarder -lPort 8080 -rHost 44.55.66.77 -rPort 9090

now we can use the ir vps ip address and 8080 port instead of 44.55.66.77 and 9090 for connecting to that service running on the usa vps. 

` client > x.x.x.x:8080 -> 44.55.66.77:9090`

you can use `screen` for keeping the process alive or write a `systemd unit file`
( you can install `screen` tool using `apt install screen -y` on ubuntu)

# FAQ
- can we use domain address instead of ip address?
```
Yes, you just need to use domain address in the rHost param value.
in case of the example above, if the domain address is sub.domain.com then it will become this:
> ./tcpforwarder -lPort 8080 -rHost sub.domain.com -rPort 9090
```

- why does it log too many open files sometimes? how to fix it?
```
The "Too Many Open Files" error indicates that this process has reached its max open socket limit.
you can check the current open file limit (open socket limit in this case) using  `ulimit -a | grep open`

to fix this issue you need to change this limit to a higher number before running the tcpforwarder
for example to set the limit to 10240 :
ulimit -n 10240
```

- what is this `can't assign requested address` error about? how to solve it?
```
it may happen when you are forwarding a wide range of ports, like 250-65534 which is almost all available ports! just leave 10000 ports away and it will work alright.
```
