package structure
=================

Golang does not support import cycles. In order to have a clear understanding what goes where
this project follows the following alphabetical naming convention:
    
top level directories:
 * `app` - contains the individual apps, depends on game and lib
 * `game` - contains all game specific constructs, depends on lib and asset
 * `lib` - provides generic constructs used by gama
 * `media` - contains assets
 
__dependencies point downwards__

this could look like:
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


or see ![packagestructure](https://raw.githubusercontent.com/co0p/tankism/master/docs/packagestructure.png) 
