@echo off

start "service1" go run main.go --server_address :8001 &
start "service2" go run main.go --server_address :8002 &
start "service3" go run main.go --server_address :8003
pause