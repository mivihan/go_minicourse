package main

import (
	"context"
	"fmt"
	"net"

	"go_minicourse/HW4/proto"
	"go_minicourse/HW4/server/elastic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type accountServer struct {
	proto.UnimplementedAccountServiceServer
}

var storage elastic.AccountStorage

func (s accountServer) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.Account, error) {
	exists, err := storage.IsExistsAccount(req.Name, ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка проверки существования аккаунта: %w", err)
	}
	if !exists {
		return nil, status.Errorf(codes.NotFound, "аккаунт %s не найден", req.Name)
	}
	account, err := storage.GetAccount(req.Name, ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения аккаунта: %w", err)
	}
	return account, nil
}

func (s accountServer) CreateAccount(ctx context.Context, req *proto.CreateAccountRequest) (*proto.Account, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "имя не может быть пустым")
	}
	exists, err := storage.IsExistsAccount(req.Name, ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка проверки существования аккаунта: %w", err)
	}
	if exists {
		return nil, status.Errorf(codes.PermissionDenied, "аккаунт уже существует")
	}
	newAccount := &proto.Account{
		Name:   req.Name,
		Amount: req.Amount,
	}
	err = storage.PatchAccount(newAccount, ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка обновления аккаунта: %w", err)
	}
	return newAccount, nil
}

func (s accountServer) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest) (*emptypb.Empty, error) {
	exists, err := storage.IsExistsAccount(req.Name, ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка проверки существования аккаунта: %w", err)
	}
	if !exists {
		return nil, status.Errorf(codes.NotFound, "аккаунт %s не найден", req.Name)
	}
	err = storage.DeleteAccount(req.Name, ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка удаления аккаунта: %w", err)
	}
	return &emptypb.Empty{}, nil
}

func (s accountServer) PatchAccount(ctx context.Context, req *proto.PatchAccountRequest) (*proto.Account, error) {
	exists, err := storage.IsExistsAccount(req.Name, ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка проверки существования аккаунта: %w", err)
	}
	if !exists {
		return nil, status.Errorf(codes.NotFound, "аккаунт %s не найден", req.Name)
	}
	patchedAccount := &proto.Account{
		Name:   req.Name,
		Amount: req.Amount,
	}
	err = storage.PatchAccount(patchedAccount, ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка обновления аккаунта: %w", err)
	}
	return patchedAccount, nil
}

func (s accountServer) RenameAccount(ctx context.Context, req *proto.RenameAccountRequest) (*proto.Account, error) {
	if req.NewName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "новое имя не может быть пустым")
	}
	exists, err := storage.IsExistsAccount(req.OldName, ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка проверки существования аккаунта: %w", err)
	}
	if !exists {
		return nil, status.Errorf(codes.NotFound, "аккаунт %s не найден", req.OldName)
	}
	exists, err = storage.IsExistsAccount(req.NewName, ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка проверки существования аккаунта: %w", err)
	}
	if exists {
		return nil, status.Errorf(codes.PermissionDenied, "аккаунт %s уже существует", req.NewName)
	}
	oldAccount, err := storage.GetAccount(req.OldName, ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения аккаунта: %w", err)
	}
	newAccount := &proto.Account{
		Name:   req.NewName,
		Amount: oldAccount.Amount,
	}
	err = storage.DeleteAccount(req.OldName, ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка удаления аккаунта: %w", err)
	}
	err = storage.PatchAccount(newAccount, ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка обновления аккаунта: %w", err)
	}
	return newAccount, nil
}

func main() {
	if err := storage.Init(); err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", "0.0.0.0:5445")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterAccountServiceServer(grpcServer, &accountServer{})

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
