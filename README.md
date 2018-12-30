# tankism
top down panzer game written in sdl2

## Tech Stack
Target is linux pc platform leveraging sdl2. Written in XXX. Get the lastest nightly build from here: XXXX . 

## tank classes
 
You start as a recruit. After a few kills you become a major. If you are able to survive and kill even more enemies you are promoted to general. 

__recruit__
 ![recruit](https://raw.githubusercontent.com/co0p/tankism/master/docs/recruit.png)

 * one cannon
 * 100hp 

__major__
 ![major](https://raw.githubusercontent.com/co0p/tankism/master/docs/major.png)

 * min. 5 kills
 * two cannons
 * 130hp

__general__
 ![general](https://raw.githubusercontent.com/co0p/tankism/master/docs/general.png)
 
 * min: 15 kills
 * three cannons !!!
 * 200hp 

## pickups

 * ![ammo](https://raw.githubusercontent.com/co0p/tankism/master/docs/pickup_ammo.png) More damage per shot
 * ![armor](https://raw.githubusercontent.com/co0p/tankism/master/docs/pickup_armor.png) Better armor protection
 * ![repair](https://raw.githubusercontent.com/co0p/tankism/master/docs/pickup_repair.png) Get your hp back

## Modes
 
 * Duel vs AI (one enemy at a time)
 * Duel vs one human over network
 * Deathmatch vs humans over network up to 16 players


## Other Features

* fog of war
* random level generation

  
## Nice to have
 
 * limmited field of view
 * destructable environment 
 * anti cheat system
 * particle effects
 * shop / economy... buy skins but no features


# Credits

Thank you https://kenney.nl for the amazing sprites and tilesheets: https://kenney.nl/assets/topdown-tanks-redux


# Resources

 * http://gameprogrammingpatterns.com/ a nicely writting book explaining programming patterns and how to utilize them in game development
 * https://www.youtube.com/channel/UCvcCbgIFwmjYaK3Zb_4Cgpw a youtube video series about game development in golang
 * http://lazyfoo.net/tutorials/SDL/ sdl2 tutorials
 * https://www.richardlord.net/blog/ecs/what-is-an-entity-framework.html from oop entity framework to component architecture
