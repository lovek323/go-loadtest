package loadtest

func LoadTest(function func() error) error {
	return function()
}
