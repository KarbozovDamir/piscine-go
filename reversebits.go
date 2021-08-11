package piscine

func ReverseBits(oct byte) byte {
	return oct/16 + (oct%16)*16
}
