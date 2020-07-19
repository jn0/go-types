package types

import "strconv"

// Simple counter-like type with fancy .String() formatter.
// One may find useful .Rate(secs) or .BitRate(secs) formatters too.
type Bytes uint64

const (
	BITS_PER_BYTE = 8
	JUSTBYTES     = "B"
	KILOBYTE      = Bytes(1) << 10
	KILOBYTES     = "KiB"
	MEGABYTE      = Bytes(1) << 20
	MEGABYTES     = "MiB"
	GIGABYTE      = Bytes(1) << 30
	GIGABYTES     = "GiB"
	TERABYTE      = Bytes(1) << 40
	TERABYTES     = "TiB"
	PETABYTE      = Bytes(1) << 50
	PETABYTES     = "PiB"
	EXABYTE       = Bytes(1) << 60
	EXABYTES      = "EiB"
	BYTES_MAX     = Bytes(0xffffffffffffffff)
)

func (v *Bytes) Add(n uint64) *Bytes {
	*v += Bytes(n)
	return v
}

func (v *Bytes) Set(n uint64) *Bytes {
	*v = Bytes(n)
	return v
}

func (v Bytes) Uint() uint64                { return uint64(v) }
func (v Bytes) Float() float64              { return float64(v) }
func (v *Bytes) IncrementBy(n Bytes) *Bytes { return v.Add(n.Uint()) }
func (v *Bytes) Assign(n Bytes) *Bytes      { return v.Set(n.Uint()) }

func (v Bytes) format(scale Bytes) string {
	return strconv.FormatFloat(float64(v)/float64(scale), 'f', 3, 64)
}
func (v Bytes) String() string {
	if v >= EXABYTE {
		return v.format(EXABYTE) + EXABYTES
	}
	if v >= PETABYTE {
		return v.format(PETABYTE) + PETABYTES
	}
	if v >= TERABYTE {
		return v.format(TERABYTE) + TERABYTES
	}
	if v >= GIGABYTE {
		return v.format(GIGABYTE) + GIGABYTES
	}
	if v >= MEGABYTE {
		return v.format(MEGABYTE) + MEGABYTES
	}
	if v >= KILOBYTE {
		return v.format(KILOBYTE) + KILOBYTES
	}
	return strconv.FormatUint(uint64(v), 10) + JUSTBYTES
}

func (v Bytes) Rate(seconds float64) string {
	return Bytes(float64(v)/seconds).String() + "/s"
}

// calculate bits-per-second value
func (v Bytes) Bps(seconds float64) (r float64) {
	r = v.Float()
	if v > PETABYTE {
		r /= seconds
		r *= float64(BITS_PER_BYTE)
	} else {
		r *= float64(BITS_PER_BYTE)
		r /= seconds
	}
	return
}
func (v Bytes) BitRate(seconds float64) string {
	r := Bytes(v.Bps(seconds)).String()
	return r[:len(r)-1] + " bps"
}

// vim:set ai:EOF //
