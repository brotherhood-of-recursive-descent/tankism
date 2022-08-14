ECS entities, components and systems
====================================

Entities are simple IDs that are used to group components together. Components contain 
the data upon which systems execute behaviour. 

When modelling objects in tankism we ask ourselves: "what beahviour does the object contain"

So for example a barrel contains of
- a position in the world 
- a visual representation
- a colliding body
- health level (full or not)

And now systems pick up on those and 
- render the visual representation at the given position
- transfer the health to the other colliding body 

In order to get an overview on identify what things actually consist of the following random list of thoughts might help

Objects in the game are
- Pickups (barrel)
- Tanks (human controllable, AI controllable)
- Ground (grass, sand, ...)
- Bullets
- Explosions
- Lights

Bahaviours are:
- movement
- Targeting AI
- PathFinding AI
- Input Controlling


