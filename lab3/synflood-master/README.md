watch -n 0.01 "netstat -tn src :3333 | grep 127.0.0.1:3333"
sudo ./synflood-master -h 127.0.0.1 -p 3333
go run ../kerberos.go
go run ../client.go