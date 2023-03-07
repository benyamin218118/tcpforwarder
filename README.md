# tcp forwarder

its a simple tcp forwarding tool to forward incoming tcp connections to a remote service at a remote host

`you can use a domain as the remote host too.`

## Download and usage :
```
wget https://github.com/benyamin218118/tcpforwarder/raw/main/tcpforwarder
chmod +x ./tcpforwarder

./tcpforwarder -lPort LISTEN_PORT -rHost REMOTE_SERVER_HOST -rPort REMOTE_SERVER_PORT
```

## examples :

**we want to forward incoming tcp connections from our ir vps to a service running on port 9090 on a usa vps**
- first we need to choose a listen port for the ir vps like 8080 ( we will accept the incoming tcp connections on it )
- our usa vps ip address is 44.55.66.77 and that service port is 9090 as mentioned before
- so we need to replace the variables in `./tcpforwarder -lPort LISTEN_PORT -rHost REMOTE_SERVER_HOST -rPort REMOTE_SERVER_PORT` and run it on the ir vps:

> ./tcpforwarder -lPort 8080 -rHost 44.55.66.77 -rPort 9090

now we can use the ir vps ip address and 8080 port instead of 44.55.66.77 and 9090 for connecting to that service running on the usa vps. 

` client > x.x.x.x:8080 -> 44.55.66.77:9090`


you can use `screen` for keeping the process alive or write a `systemd unit file`
( you can install `screen` tool using `apt install screen -y` on ubuntu)
