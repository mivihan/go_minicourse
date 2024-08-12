package main

import (
	"go_minicourse/HW2/dto/"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ServerConfig struct {
	Host string
	Port int
}

type Command struct {
	Name    string
	Execute func(ServerConfig) error
}

var commands = []Command{
	{
		Name: "Создать новый аккаунт",
		Execute: func(config ServerConfig) error {
			var req dto.CreateAccountRequest
			fmt.Print("Введите имя нового аккаунта: ")
			_, _ = fmt.Scan(&req.Name)
			fmt.Print("Введите баланс: ")
			_, _ = fmt.Scan(&req.Amount)

			data, err := json.Marshal(req)
			if err != nil {
				return fmt.Errorf("json marshal failed: %w", err)
			}

			_, err = sendRequest(config, "api/account/create", "POST", data)
			if err != nil {
				return err
			}

			fmt.Println("Аккаунт создан")
			return nil
		},
	},
	{
		Name: "Получить информацию об аккаунте",
		Execute: func(config ServerConfig) error {
			var name string
			fmt.Print("Введите имя аккаунта: ")
			_, _ = fmt.Scan(&name)

			resp, err := sendRequest(config, fmt.Sprintf("api/account?name=%s", name), "GET", nil)
			if err != nil {
				return err
			}

			var response dto.GetAccountResponse
			if err := json.Unmarshal(resp, &response); err != nil {
				return fmt.Errorf("json unmarshal failed: %w", err)
			}

			fmt.Printf("Имя: %s, Баланс: %d\n", response.Name, response.Amount)
			return nil
		},
	},
	{
		Name: "Изменить баланс",
		Execute: func(config ServerConfig) error {
			var req dto.PatchAccountRequest
			fmt.Print("Введите имя аккаунта: ")
			_, _ = fmt.Scan(&req.Name)
			fmt.Print("Введите новую сумму: ")
			_, _ = fmt.Scan(&req.Amount)

			data, err := json.Marshal(req)
			if err != nil {
				return fmt.Errorf("json marshal failed: %w", err)
			}

			_, err = sendRequest(config, "api/account", "PATCH", data)
			if err != nil {
				return err
			}

			fmt.Println("Баланс изменен")
			return nil
		},
	},
	{
		Name: "Удалить аккаунт",
		Execute: func(config ServerConfig) error {
			var req dto.DeleteAccountRequest
			fmt.Print("Введите имя аккаунта: ")
			_, _ = fmt.Scan(&req.Name)

			data, err := json.Marshal(req)
			if err != nil {
				return fmt.Errorf("json marshal failed: %w", err)
			}

			_, err = sendRequest(config, "api/account", "DELETE", data)
			if err != nil {
				return err
			}

			fmt.Println("Аккаунт удален")
			return nil
		},
	},
	{
		Name: "Переименовать аккаунт",
		Execute: func(config ServerConfig) error {
			var req dto.ChangeAccountRequest
			fmt.Print("Введите имя аккаунта: ")
			_, _ = fmt.Scan(&req.Name)
			fmt.Print("Введите новое имя: ")
			_, _ = fmt.Scan(&req.NewName)

			data, err := json.Marshal(req)
			if err != nil {
				return fmt.Errorf("json marshal failed: %w", err)
			}

			_, err = sendRequest(config, "api/account/rename", "POST", data)
			if err != nil {
				return err
			}

			fmt.Println("Аккаунт переименован")
			return nil
		},
	},
}

func sendRequest(config ServerConfig, path, method string, data []byte) ([]byte, error) {
	url := fmt.Sprintf("http://%s:%d/%s", config.Host, config.Port, path)
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http request create failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("read body failed: %w", err)
		}
		return nil, fmt.Errorf("response error: %s", string(body))
	}

	return io.ReadAll(resp.Body)
}

func main() {
	host := flag.String("host", "http://localhost", "server host")
	port := flag.Int("port", 8080, "server port")
	secretKey := flag.String("secret-key", "", "admin secret key")
	flag.Parse()

	config := ServerConfig{
		Host: *host,
		Port: *port,
	}

	if *secretKey != "" {
		fmt.Println("Попытка получить список всех аккаунтов...")
		resp, err := sendRequest(config, fmt.Sprintf("api/accounts?secret-key=%s", *secretKey), "GET", nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var accounts []dto.GetAccountResponse
		if err := json.Unmarshal(resp, &accounts); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Всего аккаунтов: %d\n", len(accounts))
		for _, acc := range accounts {
			fmt.Printf("Имя: %s, Баланс: %d\n", acc.Name, acc.Amount)
		}
		os.Exit(0)
	}

	for {
		showMenu(config)
	}
}

func showMenu(config ServerConfig) {
	fmt.Println()
	for i, cmd := range commands {
		fmt.Printf("%d > %s\n", i+1, cmd.Name)
	}
	fmt.Println("0 > Выход")
	fmt.Print("Выберите команду: ")

	var choice int
	_, _ = fmt.Scan(&choice)
	if choice < 0 || choice > len(commands) {
		fmt.Println("Неверный выбор")
		return
	}

	if choice == 0 {
		os.Exit(0)
	}

	if err := commands[choice-1].Execute(config); err != nil {
		fmt.Println(err)
	}
}
