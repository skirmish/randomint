package randomint

import (
	"testing"
	"math/rand"
)

func TestRandom(t *testing.T) {
	r := NewRandomInt(1)
	for i := 0; i < 100; i++ {
		r.Intn(5000)
	}
	r.Intn(0)
}

func TestRandom1K(t *testing.T) {
	r := NewRandomInt(1000)
	for i := 0; i < 100; i++ {
		r.Intn(5000)
	}
}

func TestRandom3(t *testing.T) {
	r := NewRandomInt(1)
	for i := 0; i < 100; i++ {
		r.Intn(3)
	}
}

func TestRandomInt32(t *testing.T) {
	r := NewRandomInt(1)
	for i := 0; i < 100; i++ {
		r.Int32()
	}
}

func BenchmarkRandom1(b *testing.B) {
	b.ReportAllocs()
	r := NewRandomInt(1)
	for n := b.N; n > 0; n-- {
		r.Intn(5000)
	}
	count:=int64(b.N)*4
	b.SetBytes(count)
}

func BenchmarkRandom256(b *testing.B) {
	//b.SkipNow()
	b.ReportAllocs()
	r := NewRandomInt(256)
	for n := b.N; n > 0; n-- {
		r.Intn(5000)
	}
	count:=int64(b.N)*4
	b.SetBytes(count)
}

func BenchmarkRandom8K(b *testing.B) {
	b.ReportAllocs()
	r := NewRandomInt(8192)
	for n := b.N; n > 0; n-- {
		r.Intn(5000)
	}
	count:=int64(b.N)*4
	b.SetBytes(count)
}

func BenchmarkRandom64K(b *testing.B) {
	b.ReportAllocs()
	r := NewRandomInt(256 * 256)
	for n := b.N; n > 0; n-- {
		r.Intn(5000)
	}
	count:=int64(b.N)*4
	b.SetBytes(count)
}

func BenchmarkRandom16M(b *testing.B) {
	b.ReportAllocs()
	r := NewRandomInt(256 * 256 * 256)
	for n := b.N; n > 0; n-- {
		r.Intn(5000)
	}
	count:=int64(b.N)*4
	b.SetBytes(count)
}

func BenchmarkStdLibRand(b *testing.B) {
	b.ReportAllocs()
	r := rand.New(rand.NewSource(1))
	for n := b.N; n > 0; n-- {
		r.Intn(5000)
	}
	count:=int64(b.N)*4
	b.SetBytes(count)
}
