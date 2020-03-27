// Package schematic implements the parsing of Schematic files. Due to the format of schematics, only blocks
// with an ID smaller than 256 are supported. This means that new blocks are not supported.
// Due to new blocks not being supported, this package does not provide methods to save blocks to a schematic.
package schematic
