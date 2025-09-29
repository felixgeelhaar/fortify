//go:build example
// +build example

	tm := timeout.New[*http.Response](timeout.Config{
		DefaultTimeout: time.Second * 30,
		OnTimeout: func() {
			log.Println("Request timed out")
		},
	})
