package main

import (
	"ReadWriteExell/fun_ReadWrite"
	"fmt"
)

func main() {

	fileName := "E:\\MyGo\\ThucHanh\\readAndwrite\\ReadExell\\Sofe_QuyetToan.xlsm"
	shName := "Sheet2"
	mang := fun_ReadWrite.ReadExell(fileName, shName)
	fmt.Println(mang)
}
