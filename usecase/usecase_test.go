package usecase

import(
	"testing"
	"github.com/stretchr/testify/suite"
	"assessment-tennis-player/http/response"
)

type TestSuite struct {
	suite.Suite
	usecase InterfaceUseCase
	data []*response.DataContainer
	flag int
}

func (s *TestSuite) SetupSuite() {
	s.usecase = NewUseCase(nil, s.data, s.flag)
}

func TestInit (t *testing.T)  {
	BallReady = 1
	SetContainer = 1
	suite.Run(t, new( TestSuite ))
}

func (s *TestSuite) TestGetContainerSucces() {
	//Test Container according as the set variable setContainer
	data := s.usecase.PutBallToContainerUseCase()
	s.Suite.NotEmpty(data)
	s.Suite.Equal( len(data), SetContainer )
}

func (s *TestSuite) TestGetContainerFailed() {
	//Test Container according as the set variable setContainer
	data := s.usecase.PutBallToContainerUseCase()
	s.Suite.NotEmpty(data)
	s.Suite.NotEqual( len(data), 5 )
}

func (s *TestSuite) TestBallOnContainerSuccess() {
	expectedBall := 1
	//create container
	_ = s.usecase.PutBallToContainerUseCase()
	
	//Test Ball insert
	dataBall := s.usecase.PutBallToContainerUseCase()
	s.Suite.Equal( dataBall[0].Ball, expectedBall )
}

func (s *TestSuite) TestBallOnContainerFailed() {
	expectedBall := 5

	//Test Container according as the set variable
	_ = s.usecase.PutBallToContainerUseCase()
	
	//Test Ball insert
	dataBall := s.usecase.PutBallToContainerUseCase()
	s.Suite.NotEqual( dataBall[0].Ball, expectedBall )
}