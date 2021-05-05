# gomodbus-pm-311x
## Description
Use modbus-tcp to communicate PM-311x

## Build
```bash
# run on linux/arm7
GOOS=linux GOARCH=arm GOARM=7 go build -o read_modbus_tcp main.go
```

## Execute
![](docs/demo.png)