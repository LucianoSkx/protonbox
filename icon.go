package main

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed assets/icon.png
var iconData []byte

var resourceIcon = fyne.NewStaticResource("icon.png", iconData)
