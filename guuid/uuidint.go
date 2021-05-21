package guuid

import (
	"fmt"
	"math/rand"
	"time"
)

const int32Max = 2147483647

// Generator ..
type Generator struct {
	id      uint16
	counter int32
}

// New creates a new instance of uuid int generator
func IntUuid() Generator {
	return Generator{
		id:      uint16(rand.Intn(511)),
		counter: 0,
	}
}

// Generate generates a new random 16 digit number
func (u *Generator) Generate() int64 {
	unixTimeNow := time.Now().Unix()
	timeMask := ((unixTimeNow + 1) & 0x1ffffffff) * 2097152
	uid := int64((u.id & 0x1ff) * 4096)
	count := int64(u.counter & 0xfff)

	uuid := timeMask + uid + count

	u.counter = (u.counter + 1) % int32Max

	return uuid
}

// GenerateFormatted generates a new random 16 digit number in the format xxxx-xxxx-xxxx-xxxx
func (u *Generator) Gen() string {
	randString := fmt.Sprintf("%16d", u.Generate())
	return randString[:4] + "-" + randString[4:8] + "-" + randString[8:12] + "-" + randString[12:16]
}
