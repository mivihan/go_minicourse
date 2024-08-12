package main

import (
	"go_minicourse/HW3/proto"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"sync"
)

type accountServer struct {
	proto.UnimplementedAccountServiceServer
}

var (
	secretKey = uuid.New().String()
	accountDB = make(map[string]*proto.Account)
	mutex     = &sync.RWMutex{}
)

func (accountServer) GetAccount(ctx context.Context, req *proto.AccountRequest) (*proto.AccountResponse, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	account, exists := accountDB[req.Name]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "account %s not found", req.Name)
	}

	return &proto.AccountResponse{Name: account.Name, Balance: account.Balance}, nil
}

func (accountServer) CreateAccount(ctx context.Context, req *proto.NewAccountRequest) (*proto.AccountResponse, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}

	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := accountDB[req.Name]; exists {
		return nil, status.Errorf(codes.PermissionDenied, "account already exists")
	}

	newAccount := &proto.Account{
		Name:    req.Name,
		Balance: req.Balance,
	}

	accountDB[req.Name] = newAccount

	return &proto.AccountResponse{Name: newAccount.Name, Balance: newAccount.Balance}, nil
}

func (accountServer) DeleteAccount(ctx context.Context, req *proto.RemoveAccountRequest) (*emptypb.Empty, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := accountDB[req.Name]; !exists {
		return nil, status.Errorf(codes.NotFound, "account %s not found", req.Name)
	}

	delete(accountDB, req.Name)

	return &emptypb.Empty{}, nil
}

func (accountServer) PatchAccount(ctx context.Context, req *proto.UpdateAccountRequest) (*proto.AccountResponse, error) {
	mutex.Lock()
	defer mutex.Unlock()

	account, exists := accountDB[req.Name]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "account %s not found", req.Name)
	}

	account.Balance = req.Balance

	return &proto.AccountResponse{Name: account.Name, Balance: account.Balance}, nil
}

func (accountServer) RenameAccount(ctx context.Context, req *proto.ChangeAccountNameRequest) (*proto.AccountResponse, error) {
	if req.NewName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}

	mutex.Lock()
	defer mutex.Unlock()

	account, exists := accountDB[req.OldName]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "account %s not found", req.OldName)
	}

	if _, exists := accountDB[req.NewName]; exists {
		return nil, status.Errorf(codes.PermissionDenied, "account %s already exists", req.NewName)
	}

	account.Name = req.NewName
	accountDB[req.NewName] = account
	delete(accountDB, req.OldName)

	return &proto.AccountResponse{Name: account.Name, Balance: account.Balance}, nil
}

func (accountServer) GetAllAccounts(ctx context.Context, req *proto.AllAccountsRequest) (*proto.AllAccountsResponse, error) {
	if req.SecretKey != secretKey {
		return nil, status.Errorf(codes.PermissionDenied, "invalid secret key")
	}

	mutex.RLock()
	defer mutex.RUnlock()

	var accounts []*proto.Account
	for _, account := range accountDB {
		accounts = append(accounts, &proto.Account{
			Name:    account.Name,
			Balance: account.Balance,
		})
	}

	return &proto.AllAccountsResponse{Accounts: accounts}, nil
}

func main() {
	println("Секретный ключ:", secretKey)

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
