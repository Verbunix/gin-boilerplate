package test

func NewTestService() *TestService {
	return &TestService{}
}

type TestService struct{}

func (t *TestService) ping() string {
	return "pong"
}
