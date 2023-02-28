# tcp forwarder

its a simple tcp forwarding tool

## Download and usage :
```
wget https://github.com/benyamin218118/tcpforwarder/raw/main/tcpforwarder
chmod +x ./tcpforwarder

./tcpforwarder -lPort LISTEN_PORT -rHost REMOTE_SERVER_HOST -rPort REMOTE_SERVER_PORT
```

## examples :

> **we want to forward tcp connection from our ir server to a service running on port 9090 on a usa vps**
- **we choose a listen port for the ir vps like 8080**
- **our usa vps ip address is 44.55.66.77 and that service port is 9090 as mentioned before**
- **so we need to replace the variables in `./tcpforwarder -lPort LISTEN_PORT -rHost REMOTE_SERVER_HOST -rPort REMOTE_SERVER_PORT` and it becomes** :

(don't forget, we need to run this on our ir vps)
> ./tcpforwarder -lPort 8080 -rHost 44.55.66.77 -rPort 9090

no we can use the ir server ip address and 8080 port instead of 44.55.66.77 and 9090 for connecting to that service running on the usa vps. 

` client > x.x.x.x:8080 -> 44.55.66.77:9090`


you can use `screen` for keeping the process alive ( you can install it with `apt install screen -y` on ubuntu)