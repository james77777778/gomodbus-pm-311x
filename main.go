package main

import (
	"fmt"
	"encoding/binary"

	"github.com/goburrow/modbus"
)

func main(){
	client := modbus.TCPClient("192.168.255.1:502")

	// model info
	byteData, _ := client.ReadInputRegisters(0x0202, 1)
	valueData := binary.BigEndian.Uint16(byteData)
	fmt.Println("Model Name:", valueData)

	byteData, _ = client.ReadInputRegisters(0x0203, 1)
	valueData = binary.BigEndian.Uint16(byteData)
	fmt.Println("Model Type:", valueData, "(1: 50Hz, 2: 60Hz)")

	// 1
	byteData, _ = client.ReadInputRegisters(0x1201, 1)
	valueData = binary.BigEndian.Uint16(byteData)
	fmt.Println("V_1:", float32(valueData) * 0.1)
	byteData, _ = client.ReadInputRegisters(0x1203, 1)
	valueData = binary.BigEndian.Uint16(byteData)
	fmt.Println("I_1:", float32(valueData) * 0.1)

	// 2
	byteData, _ = client.ReadInputRegisters(0x1212, 1)
	valueData = binary.BigEndian.Uint16(byteData)
	fmt.Println("V_2:", float32(valueData) * 0.1)
	byteData, _ = client.ReadInputRegisters(0x1214, 1)
	valueData = binary.BigEndian.Uint16(byteData)
	fmt.Println("I_2:", float32(valueData) * 0.1)

	// 3
	byteData, _ = client.ReadInputRegisters(0x1223, 1)
	valueData = binary.BigEndian.Uint16(byteData)
	fmt.Println("V_3:", float32(valueData) * 0.1)
	byteData, _ = client.ReadInputRegisters(0x1225, 1)
	valueData = binary.BigEndian.Uint16(byteData)
	fmt.Println("I_3:", float32(valueData) * 0.1)

	// 2
	byteData, _ = client.ReadInputRegisters(0x1234, 1)
	valueData = binary.BigEndian.Uint16(byteData)
	fmt.Println("V_4:", float32(valueData) * 0.1)
	byteData, _ = client.ReadInputRegisters(0x1236, 1)
	valueData = binary.BigEndian.Uint16(byteData)
	fmt.Println("I_4:", float32(valueData) * 0.1)
}