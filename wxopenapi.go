package wxopenapi

func CGIBin(opts ...option) *cgiBin {
	self := &cgiBin{}
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func Connect(opts ...option) *connect {
	self := &connect{}
	for _, opt := range opts {
		opt(self)
	}
	return self
}
