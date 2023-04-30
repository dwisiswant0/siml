package siml

import (
	"sync"

	"net/http"
)

var (
	wg     sync.WaitGroup
	client *http.Client
)
