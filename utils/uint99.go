package utils

import "fmt"

type UInt99 struct {
	Value    uint8
	Overflow int
}

func normalize(current uint8, value int) (uint8, int) {
	total := int(current) + value
	val := ((total % 100) + 100) % 100
	overflow := 0

	if value > 0 {
		overflow = total / 100
	} else {
		if int(current) == 0 {
			overflow = abs(value) / 100
		} else if abs(value) >= int(current) {
			overflow = 1 + (abs(value)-int(current))/100
		}
	}

	return uint8(val), abs(overflow)
}

func (u *UInt99) Add(n int) {
	val, ov := normalize(u.Value, n)
	u.Value = val
	u.Overflow += ov
}

func (u *UInt99) Sub(n int) {
	val, ov := normalize(u.Value, -n)
	u.Value = val
	u.Overflow += ov
}

func (u UInt99) String() string {
	return fmt.Sprintf("%d", u)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
