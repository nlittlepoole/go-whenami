[![forthebadge](https://forthebadge.com/images/badges/contains-technical-debt.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/certified-snoop-lion.svg)](https://forthebadge.com)
# go-whenami
[![Go Report Card](https://goreportcard.com/badge/github.com/nlittlepoole/go-whenami)](https://goreportcard.com/report/github.com/nlittlepoole/go-whenami)
`go-whenami` is a package that provides functionality to lookup the timezone of a given coordinate.

```golang
import (
    "log"
    "github.com/nlittlepoole/go-whenami"
)

func main(){
    lat := -77.0283 // float64
    long := 38.389 // float64
    timezone, err = whenami.WhenAmI(lat, long)
    if err != nil{
      log.Warn(err)
    }
}
```
![](https://media.giphy.com/media/hYJymOkDJYYBa/giphy.gif)

## Overview
This is a port of the Python project [tzgeo](https://pypi.python.org/pypi/tzgeo/0.0.4). You can use this library to quickly associate a 
timezone with an (Latitude, Longitude) coordinate. In particular, this is helpful when parsing server or CDN logs that don't have timezone information 
but have coordinates provided via GeoIP.

## What's Included
### Server
I've included a small REST microservice built with the Gin framework. It exposes two endpoints. The first is a health check endpoint
for scenarios where youy want to host this in a production system. The second is a `/whenami/` GET endpoint that takes latitude and longitude 
as query parameters. It returns the timezone and any encountered errors.

### Dockerfile
I've also created a Dockerfile that can be used to build spatialite  and golang on top of Alpine. The container I've written also runs the
server but feel free to rip any of that ode for another purpose. 

## Todo

- [ ] Self host the tzgeo.sqlite file
- [ ] Add tests
- [ ] Create a base container for Spatialite + Golang on Alpine
