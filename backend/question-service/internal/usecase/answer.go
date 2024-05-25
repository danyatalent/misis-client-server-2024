package usecase

import (
	"context"
	pb "github.com/danyatalent/protos/gen/go/answer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

type AnswerUseCase struct {
	l *slog.Logger
	pb.AnswerClient
}

func NewAnswer(l *slog.Logger, addr string) *AnswerUseCase {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.Error("Failed to create client", "error", err)
		return nil
	}
	return &AnswerUseCase{
		l:            l,
		AnswerClient: pb.NewAnswerClient(conn),
	}
}

func (uc *AnswerUseCase) CheckAnswer(ctx context.Context, in *pb.AnswerRequest) (*pb.AnswerResponse, error) {
	const op = "AnswerUseCase.CheckAnswer"
	uc.l = uc.l.With(op)

	out, err := uc.AnswerClient.CheckAnswer(ctx, in)
	if err != nil {
		uc.l.Info(err.Error())
		return &pb.AnswerResponse{}, err
	}
	return out, nil
}
