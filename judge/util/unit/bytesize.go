package unit

type Byte int64

const (
	// Decimal
	KB Byte = 1000
	MB      = 1000 * KB
	GB      = 1000 * MB
	TB      = 1000 * GB
	PB      = 1000 * TB

	// Binary
	KiB Byte = 1024
	MiB      = 1024 * KiB
	GiB      = 1024 * MiB
	TiB      = 1024 * GiB
	PiB      = 1024 * TiB
)
