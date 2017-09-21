// Custom random package because we just need random numbers and not all the baggage that hangs around with the
// standard random package.It obviously doesn't do all that the standard package does either.
// This has been mildly optimized for speed of getting integer random numbers.

package randomint

import (
	"crypto/rand"
)

type RandomInt struct {
	bytes  []byte
	offset int
	len    int
}

func NewRandomInt(bufferSize int) *RandomInt {
	r := &RandomInt{}
	r.bytes = make([]byte, bufferSize*4)
	r.offset = 0
	r.len = bufferSize * 4
	r.read()
	return r
}

// Intn returns an int32 random number.
func (r *RandomInt) Int32() int32 {
	var value int32
	// End of buffer?
	if r.offset+4 >= r.len {
		r.read()
	}
	value = int32(r.bytes[r.offset])
	value += int32(r.bytes[r.offset+1]) << 8
	value += int32(r.bytes[r.offset+2]) << 16
	value += int32(r.bytes[r.offset+3]) << 24
	r.offset += 4
	return value
}

// Intn returns an int random number in [0,max).
func (r *RandomInt) Intn(max int) int {
	if max == 0 {
		return 0
	}
	// End of buffer?
	if r.offset+4 >= r.len {
		r.read()
	}
	var value int
	if max < 256 {
		value = int(r.bytes[r.offset])
		r.offset += 1
	} else {
		value = int(r.bytes[r.offset])
		value += int(r.bytes[r.offset+1]) << 8
		value += int(r.bytes[r.offset+2]) << 16
		value += int(r.bytes[r.offset+3]) << 24
		r.offset += 4
	}

	if value < 0 {
		value = -value
	}
	if value >= max {
		value %= max
	}
	return value
}

func (r *RandomInt) read() (count int, err error) {
	count, err = rand.Read(r.bytes)
	r.offset = 0
	return
}
