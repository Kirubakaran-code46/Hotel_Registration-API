package main

import (
	s3filehandler "HOTEL-REGISTRY_API/common/S3FileHandler"
	"HOTEL-REGISTRY_API/helpers"
	"fmt"
)

func main() {

	lDebug := new(helpers.HelperStruct)

	file, err := s3filehandler.S3FileDownload(lDebug, "DOC20250704160627_3543.PNG")
	if err != nil {
		fmt.Println("###", err)
	} else {
		fmt.Println("file->", file)
	}
}
