package api

import "math"

// EffectiveOffset translated negative offsets into positive ones.
func EffectiveOffset(reqOffset int, total int) (offset int) {
	switch {
	case reqOffset < -total:
		return 0
	case reqOffset < 0:
		return int(math.Max(float64(total+reqOffset), 0))
	default:
		return reqOffset
	}
}

// EffectiveLimit computes a hard limit based on the offset and total available items if there is a
// limit set.
// Input: non-negative offset
func EffectiveLimit(limit int, offset int, total int) int {
	if offset < 0 {
		panic("input offset has to be non-negative")
	}
	switch {
	case limit <= 0:
		return -1
	case limit > total-offset:
		return total - offset
	default:
		return limit
	}
}

// EffectiveOffsetNLimit chains EffectiveOffset and EffectiveLimit together.
func EffectiveOffsetNLimit(reqOffset int, reqLimit int, totalItems int) (offset int, limit int) {
	offset = EffectiveOffset(reqOffset, totalItems)
	limit = EffectiveLimit(reqLimit, offset, totalItems)
	return offset, limit
}
