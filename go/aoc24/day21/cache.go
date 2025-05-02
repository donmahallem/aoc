package day21

type Cache = map[uint16]uint

func HashId2(start *Point, end *Point, depth uint8) uint16 {
	var value uint16 = uint16(depth)
	value = (value * 4) + uint16(start.X)
	value = (value * 4) + uint16(start.Y)
	value = (value * 4) + uint16(end.X)
	value = (value * 4) + uint16(end.Y)
	return value
}

// takes current start and end points and the remaining depth
// and creates a hashKey
func HashId(start *Point, end *Point, remainingDepth uint8) uint16 {
	tmpValue := (start.X * 4) + start.Y
	tmpValue = (tmpValue * 4) + end.X
	tmpValue = (tmpValue * 4) + end.Y

	// <<8 as 4*4*4*4 = 256
	return uint16(remainingDepth)<<8 + uint16(tmpValue)
}
