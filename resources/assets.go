package resources

import _ "embed"

//go:embed start_background.jpg
var BackgroundImage []byte

//go:embed tanks/tank_green.png
var TankImage []byte

//go:embed tanks/tank_bigRed.png
var BigTankImage []byte
