package gapi

import (
	"context"
	"time"

	db "github.com/TunKeyy/BankingGO/db/sqlc"
	"github.com/TunKeyy/BankingGO/pb"
	"github.com/TunKeyy/BankingGO/util"
	"github.com/TunKeyy/BankingGO/val"
	"github.com/TunKeyy/BankingGO/worker"
	"github.com/hibiken/asynq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}
	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	// TODO: use db transaction
	taskPayload := &worker.PayloadSendVerifyEmail{
		Username: user.Username,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}
	err = server.taskDistributor.DistributeTaskVerifyEmail(ctx, taskPayload, opts...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to distribute task to send verify email: %s", err)
	}

	rsp := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return rsp, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.Username); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	if err := val.ValidatePassword(req.Password); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}
	if err := val.ValidateEmail(req.Email); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	if err := val.ValidateFullName(req.FullName); err != nil {
		violations = append(violations, fieldViolation("full_name", err))
	}
	return violations
}
