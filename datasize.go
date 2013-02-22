package main

import (
	"fmt"
	"strconv"
)

type Datasize int64

const (
	Byte = Datasize(1 << (10 * iota))
	KiloByte
	MegaByte
	GigaByte
	TeraByte
)

func (f Datasize) String() string {
	value := float64(0)
	suffix := "B"
	switch {
	case f >= TeraByte:
		value = float64(f) / float64(TeraByte)
		suffix = "T"
	case f >= GigaByte:
		value = float64(f) / float64(GigaByte)
		suffix = "G"
	case f >= MegaByte:
		value = float64(f) / float64(MegaByte)
		suffix = "M"
	case f >= KiloByte:
		value = float64(f) / float64(KiloByte)
		suffix = "K"
	}
	return fmt.Sprintf("%01.3f%s", value, suffix)
}

func (f *Datasize) MarshalGoption(s string) error {
	prefix, suffix := s[0:len(s)-1], s[len(s)-1:]
	num, err := strconv.ParseFloat(prefix, 64)
	if err != nil {
		return err
	}
	multiplier := float64(1)
	switch suffix {
	case "B":
		multiplier = float64(Byte)
	case "K":
		multiplier = float64(KiloByte)
	case "M":
		multiplier = float64(MegaByte)
	case "G":
		multiplier = float64(GigaByte)
	case "T":
		multiplier = float64(TeraByte)
	default:
		return fmt.Errorf("Unknown unit %s", suffix)
	}
	*f = Datasize(int64(multiplier * num))
	return nil
}
