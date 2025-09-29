//go:build example
// +build example

	tm := timeout.New[string](timeout.Config{
		DefaultTimeout: time.Second * 5,
		OnTimeout: func() {
			fmt.Println("Operation timed out")
		},
	})
