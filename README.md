# go-types -- assorted useful go types

## Bytes

Simple counter-like type with fancy `.String()` formatter.
One may find useful `.Rate(secs)` or `.BitRate(secs)` formatters too.

The code like
```go
…
import (
	"time"
	"fmt"
	. "xxx/types"
)
…
	var traffic Bytes
	var t0, t1 time.Time
…
	for t0 = time.Now() {
		data := endpoint.Read()
		traffic.Add(uint64(len(data)))
		…
	}
	t1 = time.Now()

	fmt.Printf("Got %s in %s at %s.\n",
		traffic, t2.Sub(t1), traffic.BitRate(t2.Sub(t1).Seconds()))
…
```
will show up something like
```
Got 123.456MiB in 12.123456s at 81.465Mi bps.
```

100% test covered

# EOF #
