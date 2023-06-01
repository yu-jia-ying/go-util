package syncs

type CallBack func(interface{}) (interface{}, error)

func Future(param interface{}, f CallBack) func() (interface{}, error) {
	var result interface{}
	var err error

	c := make(chan struct{})

	go func(chan struct{}) {
		defer close(c)
		result, err = f(param)
	}(c)

	return func() (interface{}, error) {
		<-c
		return result, err
	}
}
