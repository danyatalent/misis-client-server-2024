package v1

import (
	"context"
	pb "github.com/danyatalent/protos/gen/go/question-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *QuestionServiceServer) GetQuestion(ctx context.Context, req *pb.Empty) (*pb.QuestionResponse, error) {
	q, err := s.u.GetRandom(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "error getting random question")
	}
	resp := &pb.QuestionResponse{
		Text:    q.Text,
		Answer:  q.Answer,
		Author:  q.Author,
		Comment: q.Comment,
	}
	return resp, nil
}
