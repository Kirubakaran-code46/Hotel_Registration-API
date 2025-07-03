package main

import (
	"HOTEL-REGISTRY_API/common"
	"fmt"
)

func main() {

	lpath, lErr := common.GetFileBase64("DOC20250701105303_6632.png")

	if lErr != nil {
		fmt.Println(lErr)
	}
	fmt.Println("lpath", lpath)
}
