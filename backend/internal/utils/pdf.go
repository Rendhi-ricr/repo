package utils

import (
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

// ValidatePDF memeriksa apakah file PDF valid dan tidak rusak
func ValidatePDF(filePath string) error {
	conf := model.NewDefaultConfiguration()
	return api.ValidateFile(filePath, conf)
}

// GetPDFPageCount mengembalikan jumlah halaman dalam file PDF
func GetPDFPageCount(filePath string) (int, error) {
	return api.PageCountFile(filePath)
}

// SplitPDF memecah file PDF menjadi satu file per halaman
// Hasilnya disimpan di outputDir
func SplitPDF(filePath string, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}
	// Split file menjadi single pages
	return api.SplitFile(filePath, outputDir, 1, nil)
}
