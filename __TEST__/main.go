package main

import (
	"pump/src/help"
	"pump/src/install"
	"pump/src/initmod"
	"pump/src/mod"
)

func main() {
	help.Handle()       // örnek fonksiyon çağrısı
	install.Handle()
	initmod.Handle()
	mod.Handle()
}
