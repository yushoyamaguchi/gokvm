package eventfd

type EventNotifier struct {
	rfd int
	wfd int
}

func (e EventNotifier) EventNotifierInit() error {
	eventfd, err := Create()
	if err != nil {
		return err
	}
	e.rfd = eventfd.fd
	e.wfd = eventfd.fd
	return nil
}

func (e EventNotifier) GetWfd() int {
	return e.wfd
}
