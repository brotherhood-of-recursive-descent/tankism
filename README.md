# tankism

top down panzer game written in go

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
