package mocks

// ApplicationStub é um stub simples da interface Application
// Útil para testes onde não precisamos verificar chamadas específicas
type ApplicationStub struct {
	SetUpError    error
	RunError      error
	TearDownError error
}

// NewApplicationStub cria uma nova instância do stub
func NewApplicationStub() *ApplicationStub {
	return &ApplicationStub{}
}

// SetUp implementa Application.SetUp
func (s *ApplicationStub) SetUp() error {
	return s.SetUpError
}

// Run implementa Application.Run
func (s *ApplicationStub) Run() error {
	return s.RunError
}

// TearDown implementa Application.TearDown
func (s *ApplicationStub) TearDown() error {
	return s.TearDownError
}
