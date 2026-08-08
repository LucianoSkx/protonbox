package main

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed assets/icon.png
var iconData []byte

//go:embed assets/logo.png
var logoData []byte

var resourceIcon = fyne.NewStaticResource("icon.png", iconData)
var resourceLogo = fyne.NewStaticResource("logo.png", logoData)
