package main

import (
	"log"

	"github.com/zHenriqueGN/PDFSanitizer/pkg/controller"
	"github.com/zHenriqueGN/PDFSanitizer/pkg/models"
)

func main() {
	folders, err := controller.GetRootFolders()
	if err != nil {
		log.Fatal(err)
	}

	PDFsMap := make(map[string]models.PDFFile)

	for _, folder := range folders {
		err = controller.MapPDFFiles(folder, &PDFsMap)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = controller.CreateSanitizedPDFsFolder(PDFsMap)
	if err != nil {
		log.Fatal(err)
	}
}
