package v1

import (
	"context"
	"github.com/danyatalent/misis-client-server-2024/backend/question-service/internal/models"
	pb "github.com/danyatalent/protos/gen/go/question-service"
	"google.golang.org/grpc"
)

type QuestionService interface {
	GetRandom(ctx context.Context) (*models.Question, error)
}

type QuestionServiceServer struct {
	pb.UnimplementedQuestionsServer
	u QuestionService
}

func Register(server *grpc.Server, q QuestionService) {
	pb.RegisterQuestionsServer(server, &QuestionServiceServer{
		u: q,
	})
}

//func StartGRPCServer(addr string, uc QuestionService, l *slog.Logger) error {
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", addr))
//	if err != nil {
//		return err
//	}
//	s := grpc.NewServer()
//	pb.RegisterQuestionsServer(s, &QuestionServiceServer{
//		u: uc,
//		l: l,
//	})
//	l.Info("grpc server listening on " + addr)
//	if err := s.Serve(lis); err != nil {
//		return err
//	}
//	return nil
//}
