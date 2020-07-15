# Schematic
Schematic is a library for Dragonfly implementing schematics support. The Schematic type implements
Dragonfly's `world.Structure` interface, so that schematics may be pasted with a high throughput.

## Installation
Schematic requires at least Go 1.13. The library may be installed using:
```
go get github.com/df-mc/schematic
```

## Usage
The basic method to get and use a `Schematic` struct is using the `FromReader` function:
```go
package main

import (
	"github.com/df-mc/dragonfly/dragonfly/world"
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

## Module documentation
Schematic has a very simple interface. The package documentation can be found 
[here](https://pkg.go.dev/github.com/df-mc/schematic?tab=overview).

## Contributing
We use JetBrains Space to manage our issues, pull requests and code reviews, but we welcome contributions
through GitHub issues and pull requests.

## Contact
[![Chat on Discord](https://img.shields.io/badge/Chat-On%20Discord-738BD7.svg?style=for-the-badge)](https://discord.gg/evzQR4R)
