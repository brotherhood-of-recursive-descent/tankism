package resources

import _ "embed"

//go:embed start_background.jpg
var BackgroundImage []byte

//go:embed tanks/tank_green.png
var TankImage []byte

//go:embed tanks/tank_bigRed.png
var BigTankImage []byte

//go:embed objects/treeGreen_large.png
var BigTreeImage []byte

//go:embed objects/bulletSand3_outline.png
var BulletSandOutline []byte

//go:embed objects/oilSpill_large.png
var OilSpillLarge []byte

//go:embed objects/barricadeMetal.png
var Barricade []byte
