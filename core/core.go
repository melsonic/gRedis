package core

import "sync"

var (
	core_data map[string]any = make(map[string]any)
	mu        sync.Mutex
)
