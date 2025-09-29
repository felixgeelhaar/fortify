//go:build example
// +build example

	tm := timeout.New[*APIResponse](timeout.Config{
		DefaultTimeout: time.Second * 5,
		OnTimeout: func() {
			fmt.Println("   ⏱ Timeout")
		},
	})
