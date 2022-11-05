[![tankism](https://github.com/co0p/tankism/actions/workflows/ci.yml/badge.svg)](https://github.com/co0p/tankism/actions/workflows/ci.yml)

# tankism

A top down panzer game written in Go.


## devlog


### 10/2022 first visual effect - particles
> see a5bc3050ed29510dca9ceca1f3b589e18385e414

Yeah, we got some particles on the screen!

To create an emitter entity attach the emitter components

```go
// see app/particles/main.go
redEmitter := p.EntityManager.NewEntity()
redEmitter.AddComponent(&components.ParticleEmitter{
	Color:        lib.ColorRed,

	Lifetime_min: 10,
	Lifetime_max: 100,

	Velocity_min:  1,
	Velocity_max:  1,

	Direction_min: 0,
	Direction_max: 90,
})
redEmitter.AddComponent(&components.Transform{X: ... })

```

![particles!](https://raw.githubusercontent.com/co0p/tankism/master/docs/particles.gif) 



### 10/2022 Relative Positioning
> see 37193f42f3c4e66f7140bc9ea7614b6e01a35808

We implemented relative positioning of entities this evening. It is basically building a tree of transform components. With each entity only having max one transform and one transform only having one parent you get a scene graph.

To demonstrate relative positioning we decided to build a planet demo. As a bonus each planet as an emitting light attached to it.

```go
// see app/position/main.go
//...
sun := s.EntityManager.NewEntity()
sunTransform := components.Transform{}
sun.AddComponents(&sunTransform, ... )

earth := s.EntityManager.NewEntity()
earthTransform := components.Transform{OffsetX: 300, OffsetY: 0}
earthTransform.AddParent(&sunTransform)

moon := s.EntityManager.NewEntity()
moonTransform := components.Transform{OffsetX: 150, OffsetY: 150}
moonTransform.AddParent(&earthTransform)

```

![planets with lights, relative positioned](https://raw.githubusercontent.com/co0p/tankism/master/docs/relativePositioning.png) 



### 08/2022 SpriteAnimations making boom
> see e1a38829ba41fd49d4c6c1dfdd1aefb79291aea4

The first action has landed. Loading sprites and rendering a list of images.

![cycling through sprites](https://raw.githubusercontent.com/co0p/tankism/master/docs/tank_boom.gif) 


### 07/2022 Ambient Lighting + LightingTexture

> see 5836834a79dc1b6507f941bedf472cae370f60b2

We finaly have some color dynamics in the game. Basically we draw every lighting source into a texture
and then use a shader to merge everything. 

```go
package main

var AmbientColor vec4

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	// ( color of lightingmap + ambient color ) * destination color
	return (imageSrc0UnsafeAt(texCoord) + AmbientColor) * (color)
}
```

In the there are 3 lights in the scene. Pointlights on top of the barrel and the crate. In the right bottom corner
there is a circle light which is bigger and more diffuse. To demonstrate the different effects you can achieve 
with different ambient light colors, we cycle through some color variations ...  

![cycling through ambient lights](https://raw.githubusercontent.com/co0p/tankism/master/docs/tank_lighting.gif) 


### 05/2022 Adding first AI behaviour, the "observer"

> see de790ae4dd2528046a20e21b7da82968f84523cd

This month we managed to draw the map, introduced ZIndex to the SpriteRenderSystem and added
the first AI behaviour. A Tank rotates in the direction of the player.

![enemy follows](https://raw.githubusercontent.com/co0p/tankism/master/docs/tank_follow.gif) 


### 04/2022 Making the tank move Asteroids style

> see 4ef22abd3572bcd0eacb827181e0e58b160c6699

Finally some action on the screen. The tank is now keyboard controllable and moves across the screen.

![tank moves](https://raw.githubusercontent.com/co0p/tankism/master/docs/tank_move.gif) 


### 04/2022 Drawing a tank using Entity Component System(s) (ECS)

> see 08892283bc67aabc710d610a02a6155b8704f25a

Today we implemented a basic ECS system placing a tank entity on the world, drawing a sprite and making the tank shake. 

The generic ECS part can be found at ```/lib/ecs/*``` and the actual systems and components of tankism at ```/game/ecs/*```.

Instead of having each game object being responsible for drawing itself and updating the game logic, there are systems that are responsible for the behavior. The data is stored in components. And the entities are a means of grouping components. This separation of concerns allows to decouple systems from other systems and entities from other entities. This is also an example of data-oriented design. 

Below is the simple sprite render system:
```go
type SpriteRenderer struct {
	EntityManager ecs.EntityManager
}

func (s *SpriteRenderer) Update() error { return nil }
func (s *SpriteRenderer) Draw(screen *ebiten.Image) {

	entities := s.EntityManager.FindByComponents(components.SpriteType, components.TranslateType)
	for _, e := range entities {

		op := &ebiten.DrawImageOptions{}

		translate := e.GetComponent(components.TranslateType).(*components.Translate)
		x := translate.X
		y := translate.Y
		scale := translate.Scale

		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(x, y)

		sprite := e.GetComponent(components.SpriteType).(*components.Sprite)
		img := sprite.Image
		screen.DrawImage(img, op)
	}
}
```

![empty window](https://raw.githubusercontent.com/co0p/tankism/master/docs/ecs.gif) 


### 03/2022 Scene managment

> see 360d4bac97ead6067e44a4f016ca182cae33db35

Tankism is a big for-loop. In order to put some structure to the code, we decided to have individual scenes managing 
certain aspects of the computer game. The **Start-Scene** is responsible for loading assets and the **Game-Scene** is responsible
for managing the actual game.  

Our initial thought is something along the lines:

![scenes](https://user-images.githubusercontent.com/3327413/156920274-1e68ee0d-8453-43d9-86fe-5ba9b362e932.png) 


### 03/2022 Folder structure

> see 36813f5b94e36e06781b702731a05ff17800f7d1 

Structuring go apps is a bit different than say Java packages. We tried to follow Bil Kennedy's advice of 
package-oriented design and came up with the following folder structure.

The main goal was to make the dependencies visible in the code utilizing the folder structure. Dependencies point downwards.

```
app/
├─ client/
│  ├─ gamescene/...
│  ├─ startscene/...
│  ├─ main.go
├─ server/
│  ├─ main.go
game/
├─ objects/
│  ├─ ...
│  ├─ tank.go
├─ ui/
│  ├─ ...
│  ├─ exitbutton.go
├─ colors.go
lib/
├─ ecs/...
├─ ui/...
├─ gameobject.go
media/
├─ images/...
├─ sounds/...
```

More about it at [packagestructure](https://raw.githubusercontent.com/co0p/tankism/master/docs/packagestructure.md) 


### 03/2022 First image drawn

> see 43bf36088d9817475df42d7d8e318ac6970776e8

After more than 3 years of being dormant, we revived the idea of writing a top-down tank shooter. This time
using the ebiten library. In our first remote session we managed to create an empty window :-) 

![empty window](https://raw.githubusercontent.com/co0p/tankism/master/docs/emptywindow.png) 

## development

 * install dependencies ```go install```
 * run tests ```go test ./...```
 * run game ```go run cmd/tankism/main.go```
 
## game design

 * screens: [docs/screens.md](docs/screens.md)
 * classes: [docs/classes.md](docs/classes.md)
 * pickups: [docs/pickups.md](docs/pickups.md)


## Modes
 
 * Duel vs AI (one enemy at a time)
 * Duel vs one human over network
 * Deathmatch vs humans over network up to 16 players


## Other Features

* fog of war
* random level generation

  
## Nice to have
 
 * limited field of view
 * destructible environment 
 * anti cheat system
 * particle effects
 * shop / economy... buy skins but no features
 * call in artillery strike


# Credits

Thank you https://kenney.nl for the amazing sprites and tilesheets: https://kenney.nl/assets/topdown-tanks-redux


# Resources

 * http://gameprogrammingpatterns.com/ a nicely writting book explaining programming patterns and how to utilize them in game development
 * https://www.youtube.com/channel/UCvcCbgIFwmjYaK3Zb_4Cgpw a youtube video series about game development in golang
 * http://lazyfoo.net/tutorials/SDL/ sdl2 tutorials
 * https://www.richardlord.net/blog/ecs/what-is-an-entity-framework.html from oop entity framework to component architecture
