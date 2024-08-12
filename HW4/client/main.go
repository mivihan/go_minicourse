package main

import (
	"go_minicourse/HW4/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

type ConnectionInfo struct {
	ctx    context.Context
	client proto.AccountServiceClient
}

type CommandInfo struct {
	Command string
	Execute func(ConnectionInfo) error
}

var commands = []CommandInfo{
	{
		Command: "Создать новый аккаунт",
		Execute: func(info ConnectionInfo) error {
			var request proto.CreateAccountRequest

			fmt.Print("Введите имя нового аккаунта: ")
			_, _ = fmt.Scan(&request.Name)

			fmt.Print("Введите баланс: ")
			_, _ = fmt.Scan(&request.Amount)

			resp, err := info.client.CreateAccount(info.ctx, &request)
			if err != nil {
				return err
			}

			fmt.Printf("Аккаунт %s с балансом %v успешно создан.\n", resp.Name, resp.Amount)
			return nil
		},
	},
	{
		Command: "Получить информацию об аккаунте",
		Execute: func(info ConnectionInfo) error {
			var request proto.GetAccountRequest

			fmt.Print("Введите имя аккаунта: ")
			_, _ = fmt.Scan(&request.Name)

			resp, err := info.client.GetAccount(info.ctx, &request)
			if err != nil {
				return err
			}

			fmt.Printf("Детали аккаунта - Имя: %s, Баланс: %d\n", resp.Name, resp.Amount)
			return nil
		},
	},
	{
		Command: "Изменить баланс",
		Execute: func(info ConnectionInfo) error {
			var request proto.PatchAccountRequest

			fmt.Print("Введите имя аккаунта: ")
			_, _ = fmt.Scan(&request.Name)

			fmt.Print("Введите новую сумму: ")
			_, _ = fmt.Scan(&request.Amount)

			resp, err := info.client.PatchAccount(info.ctx, &request)
			if err != nil {
				return err
			}

			fmt.Printf("Баланс аккаунта %s обновлен до %v\n", resp.Name, resp.Amount)
			return nil
		},
	},
	{
		Command: "Удалить аккаунт",
		Execute: func(info ConnectionInfo) error {
			var request proto.DeleteAccountRequest

			fmt.Print("Введите имя аккаунта: ")
			_, _ = fmt.Scan(&request.Name)

			_, err := info.client.DeleteAccount(info.ctx, &request)
			if err != nil {
				return err
			}

			fmt.Println("Аккаунт успешно удален.")
			return nil
		},
	},
	{
		Command: "Переименовать аккаунт",
		Execute: func(info ConnectionInfo) error {
			var request proto.RenameAccountRequest

			fmt.Print("Введите текущее имя аккаунта: ")
			_, _ = fmt.Scan(&request.OldName)

			fmt.Print("Введите новое имя: ")
			_, _ = fmt.Scan(&request.NewName)

			resp, err := info.client.RenameAccount(info.ctx, &request)
			if err != nil {
				return err
			}

			fmt.Printf("Аккаунт %s переименован в %s\n", request.OldName, resp.Name)
			return nil
		},
	},
}

func main() {
	hostVal := flag.String("host", "go.grpc-es.umu-art.ru", "адрес сервера")
	portVal := flag.Int("port", 5445, "порт сервера")

	flag.Parse()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%v", *hostVal, *portVal), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	c := proto.NewAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	connectionInfo := ConnectionInfo{
		ctx,
		c,
	}

	for {
		displayMenu(connectionInfo)
	}
}

func displayMenu(connectionInfo ConnectionInfo) {
	fmt.Println("\n--- Меню ---")
	for i, command := range commands {
		fmt.Printf("%d > %s\n", i+1, command.Command)
	}
	fmt.Println("0 > Выход")
	fmt.Print("Выберите команду: ")

	var commandIndex int
	_, _ = fmt.Scan(&commandIndex)
	if commandIndex < 0 || commandIndex > len(commands) {
		fmt.Println("Неверный индекс, попробуйте снова.")
		return
	}

	if commandIndex == 0 {
		fmt.Println("Выход из программы.")
		os.Exit(0)
	}

	err := commands[commandIndex-1].Execute(connectionInfo)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err.Error())
	}
}
