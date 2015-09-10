package faker

import "testing"

func TestCodeIsbn10(t *testing.T) {
	testMatchRx(t, Code().Isbn10, `^\d{9}-[\d|X]$`)
}

func TestCodeIsbn13(t *testing.T) {
	testMatchRx(t, Code().Isbn13, `^\d{12}-\d$`)
}

func TestCodeEan13(t *testing.T) {
	testMatchRx(t, Code().Ean13, `\d{13}$`)
}

func TestCodeEan8(t *testing.T) {
	testMatchRx(t, Code().Ean8, `\d{8}$`)
}

func TestCodeRut(t *testing.T) {
	testMatchRx(t, Code().Rut, `^\d{8}-(\d|k|K)$`)
}

func TestCodeAbn(t *testing.T) {
	testMatchRx(t, Code().Abn, `^\d{11}$`)
}
