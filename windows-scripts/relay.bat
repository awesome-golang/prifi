@echo off
cd ..\\dissent\\
go run main.go config.go relay.go trusteeServer.go -relay -nclients=2 -ntrustees=1 -t1host=localhost:9000 -t2host=localhost:9000 -reportlimit=10