package utypes

import "strconv"

// Simple counter-like type with fancy .String() formatter.
// The value becomes formatted as for "%7.3f" with binary SI suffix followed by `B`.
// Suffixes are from `Ki` (kilo-, 2^10) to `Ei` (exa-, 2^60).
//
// It is supposed to declare a
// 	var traffic Bytes
// and use `traffic.Add(size)` to accumulate the value, then
// use (implicitely) `traffic.String()` to make pretty print.
//
// One may find useful .Rate(secs) or .BitRate(secs) formatters too.
//
// The formatted Bytes value takes up to 10 character positions.
// The formatted Rate() value takes up to 12 c.p.
// The formatted BitRate() value takes up to 13 c.p.
//
// NOTE: They are "traffic bytes"! The "disk bytes" are usually decimal, not binary.
type Bytes uint64

const (
	BITS_PER_BYTE        = 8
	JUSTBYTES            = "B"
	KILOBYTE             = Bytes(1) << 10
	KILOBYTES            = "KiB"
	MEGABYTE             = Bytes(1) << 20
	MEGABYTES            = "MiB"
	GIGABYTE             = Bytes(1) << 30
	GIGABYTES            = "GiB"
	TERABYTE             = Bytes(1) << 40
	TERABYTES            = "TiB"
	PETABYTE             = Bytes(1) << 50
	PETABYTES            = "PiB"
	EXABYTE              = Bytes(1) << 60
	EXABYTES             = "EiB"
	BYTES_MAX            = Bytes(0xffffffffffffffff)
	BYTES_VALUE_BITSIZE  = 64
	BYTES_VALUE_FORMAT   = 'f'
	BYTES_VALUE_DECIMALS = 3
	BYTES_PER_SECOND     = "/s"
	BITS_PER_SECONDS     = " bps"
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
	return strconv.FormatFloat(
		float64(v)/float64(scale),
		BYTES_VALUE_FORMAT,
		BYTES_VALUE_DECIMALS,
		BYTES_VALUE_BITSIZE,
	)
}

// format byte count value with binary SI suffixes
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

// format bytes-per-second value with binary SI suffixes (KiB to EiB)
func (v Bytes) Rate(seconds float64) string {
	return Bytes(float64(v)/seconds).String() + BYTES_PER_SECOND
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

// format bits-per-second value with binary SI suffixes
func (v Bytes) BitRate(seconds float64) string {
	r := Bytes(v.Bps(seconds)).String()
	return r[:len(r)-1] + BITS_PER_SECONDS
}

// vim:set ai:EOF //
