# Schematic
Schematic is a library for Dragonfly implementing schematics support. The Schematic type implements
Dragonfly's `world.Structure` interface, so that schematics may be pasted with a high throughput.

## Installation
Schematic requires at least Go 1.18. The library may be installed using:
```
go get github.com/df-mc/schematic
```

## Usage
The basic method to get and use a `Schematic` struct is using the `FromReader` function:
```go
package main

import (
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/schematic"
	"os"
)

func main() {
	file, _ := os.Open("file.schematic")
	s, _ := schematic.FromReader(file)
	
	var w *world.World
	w.BuildStructure(world.BlockPos{}, s)
}
```

## Documentation
[![Go Reference](https://pkg.go.dev/badge/github.com/df-mc/schematic.svg)](https://pkg.go.dev/github.com/df-mc/schematic)

## Contact
[![Discord Banner 2](https://discordapp.com/api/guilds/623638955262345216/widget.png?style=banner2)](https://discord.gg/U4kFWHhTNR)