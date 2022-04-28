# tankism

A top down panzer game written in Go.


## devlog

### 04/2022 Making the tank move Asteroids style

> see 4ef22abd3572bcd0eacb827181e0e58b160c6699

Finally some action on the screen. The tank is now keyboard controllable and moves across the screen.

![empty window](https://raw.githubusercontent.com/co0p/tankism/master/docs/tank_move.gif) 


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
