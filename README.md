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
	var t0 time.Time
	var dt time.Duration
…
	for t0 = time.Now();; {
		data := endpoint.Read()
		traffic.Add(uint64(len(data)))
		…
	}
	dt = time.Now().Sub(t0)

	fmt.Printf("Got %s in %s at %s.\n",
		traffic, dt, traffic.BitRate(dt.Seconds()))
…
```
will show up something like
```
Got 123.456MiB in 12.123456s at 81.465Mi bps.
```

100% test covered

# EOF #
