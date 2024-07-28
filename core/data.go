package core

import "sync"

var core_data map[string]string = make(map[string]string)
var mu sync.Mutex
