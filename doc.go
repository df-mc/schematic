// Package schematic implements the parsing of Schematic files for Dragonfly. Due to the format of schematics,
// only blocks with an ID smaller than 256 are supported. This means that new blocks are not supported.
// Due to new blocks not being supported, this package does not provide methods to save blocks to a schematic.
//
// The Schematic struct provided by the schematic package implements Dragonfly's world.Structure interface,
// allowing for a high throughput for schematic pasting and easy usage. See World.BuildStructure for pasting a
// schematic.
package schematic
