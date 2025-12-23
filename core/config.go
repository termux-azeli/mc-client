package core

import "time"

type Config struct {
	Host string
	Path string
}

const (
	ENABLE_BATCHING = true
	BATCH_WINDOW    = 2 * time.Millisecond
	BATCH_MAX_SIZE  = 32 * 1024
)
