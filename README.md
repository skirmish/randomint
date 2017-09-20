# randomint
Random Integers library, fairly random and fairly quick.

Wraps around "crypto/rand" as the source for the random source.

I needed something that would give me random integers, lots of them very quickly.
This solved the issue and it's pretty fast, it is about a 10% improvement in Go 1.9 over
the "math/rand" package and doesn't suffer from the lack of randomness I was feeling that
package gave me.

Overall it probably wasn't worth the effort to build, but hey it exists and it works.

```
r := NewRandomInt(256)  // specify a decent buffer size (1 is very slow, so maybe 256 entries at least).
val := r.Intn(123456)   // get a random value.
```