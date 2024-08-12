package main

import (
	"go_minicourse/HW3/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

type ConnInfo struct {
	ctx    context.Context
	client proto.AccountServiceClient
}

type CmdInfo struct {
	Cmd     string
	Execute func(ConnInfo) error
}

var cmdList = []CmdInfo{
	{
		Cmd: "Создать новый аккаунт",
		Execute: func(info ConnInfo) error {
			var req proto.CreateAccountRequest

			fmt.Print("Введите имя нового аккаунта: ")
			_, _ = fmt.Scan(&req.Name)

			fmt.Print("Введите баланс: ")
			_, _ = fmt.Scan(&req.Amount)

			resp, err := info.client.CreateAccount(info.ctx, &req)
			if err != nil {
				return err
			}

			fmt.Printf("Аккаунт %s с балансом %v создан\n", resp.Name, resp.Amount)
			return nil
		},
	},
	{
		Cmd: "Получить информацию об аккаунте",
		Execute: func(info ConnInfo) error {
			var req proto.GetAccountRequest

			fmt.Print("Введите имя аккаунта: ")
			_, _ = fmt.Scan(&req.Name)

			resp, err := info.client.GetAccount(info.ctx, &req)
			if err != nil {
				return err
			}

			fmt.Printf("Имя: %s, Баланс: %d\n", resp.Name, resp.Amount)
			return nil
		},
	},
	{
		Cmd: "Изменить баланс",
		Execute: func(info ConnInfo) error {
			var req proto.PatchAccountRequest

			fmt.Print("Введите имя аккаунта: ")
			_, _ = fmt.Scan(&req.Name)

			fmt.Print("Введите новую сумму: ")
			_, _ = fmt.Scan(&req.Amount)

			resp, err := info.client.PatchAccount(info.ctx, &req)
			if err != nil {
				return err
			}

			fmt.Printf("Баланс аккаунта %s теперь %v\n", resp.Name, resp.Amount)
			return nil
		},
	},
	{
		Cmd: "Удалить аккаунт",
		Execute: func(info ConnInfo) error {
			var req proto.DeleteAccountRequest

			fmt.Print("Введите имя аккаунта: ")
			_, _ = fmt.Scan(&req.Name)

			_, err := info.client.DeleteAccount(info.ctx, &req)
			if err != nil {
				return err
			}

			fmt.Printf("Аккаунт удален\n")
			return nil
		},
	},
	{
		Cmd: "Переименовать аккаунт",
		Execute: func(info ConnInfo) error {
			var req proto.RenameAccountRequest

			fmt.Print("Введите имя аккаунта: ")
			_, _ = fmt.Scan(&req.OldName)

			fmt.Print("Введите новое имя: ")
			_, _ = fmt.Scan(&req.NewName)

			resp, err := info.client.RenameAccount(info.ctx, &req)
			if err != nil {
				return err
			}

			fmt.Printf("Аккаунт %s переименован в %s\n", req.OldName, resp.Name)
			return nil
		},
	},
}

func main() {
	host := flag.String("host", "go.grpc.umu-art.ru", "server api host")
	port := flag.Int("port", 5445, "server port")
	secret := flag.String("secret-key", "", "Ключ админа")

	flag.Parse()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%v", *host, *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := proto.NewAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	connInfo := ConnInfo{
		ctx:    ctx,
		client: client,
	}

	if len(*secret) > 0 {
		fmt.Println("Попытка получить список всех аккаунтов...")
		resp, err := client.GetAllAccounts(ctx, &proto.GetAllAccountsRequest{SecretKey: *secret})

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Всего аккаунтов: %d\n", len(resp.Accounts))
			for _, account := range resp.Accounts {
				fmt.Printf("Имя: %s, Баланс: %d\n", account.Name, account.Amount)
			}
		}
		os.Exit(0)
	}

	for {
		executeCommand(connInfo)
	}
}

func executeCommand(connInfo ConnInfo) {
	fmt.Println()

	for i, cmd := range cmdList {
		fmt.Printf("%d > %s\n", i+1, cmd.Cmd)
	}
	fmt.Println("0 > Выход")
	fmt.Print("Выберите команду: ")

	var cmdIndex int
	_, _ = fmt.Scan(&cmdIndex)
	if cmdIndex < 0 || cmdIndex > len(cmdList) {
		fmt.Println("Неверный индекс")
		return
	}

	if cmdIndex == 0 {
		os.Exit(0)
	}

	err := cmdList[cmdIndex-1].Execute(connInfo)
	if err != nil {
		fmt.Println(err.Error())
	}
}
