package elastic

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	"go_minicourse/HW4/proto"
	"github.com/elastic/go-elasticsearch/v8"
)

type AccountStorage struct {
	client *elasticsearch.TypedClient
}

var ErrConnection = errors.New("ошибка соединения с Elasticsearch")

func (a *AccountStorage) Init() error {
	cert, err := os.ReadFile("HW4/server/elastic/http_ca.crt")
	if err != nil {
		return fmt.Errorf("ошибка чтения сертификата: %w", err)
	}

	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{"https://10.244.1.125:9200/"},
		APIKey:    "UzJ2bl81QUJDdTJ2aWhRam0tMXk6al9nZlZnck5SMkN0eUM1S2J0RW9TZw==",
		CACert:    cert,
	})
	if err != nil {
		return ErrConnection
	}

	exists, err := client.Indices.Exists("accounts").Do(context.Background())
	if err != nil {
		log.Fatalf("Ошибка проверки индекса: %s", err)
		return ErrConnection
	}

	if !exists {
		_, err = client.Indices.Create("accounts").Do(context.Background())
		if err != nil {
			return ErrConnection
		}
	}

	a.client = client
	return nil
}

func (a *AccountStorage) IsExistsAccount(name string, ctx context.Context) (bool, error) {
	resp, err := a.client.Exists("accounts", name).Do(ctx)
	if err != nil {
		return false, err
	}
	return resp, nil
}

func (a *AccountStorage) GetAccount(name string, ctx context.Context) (*proto.Account, error) {
	resp, err := a.client.Get("accounts", name).Do(ctx)
	if err != nil {
		return nil, err
	}

	var account proto.Account
	if err := json.Unmarshal(resp.Source_, &account); err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *AccountStorage) PatchAccount(account *proto.Account, ctx context.Context) error {
	_, err := a.client.Index("accounts").Id(account.Name).Request(account).Do(ctx)
	return err
}

func (a *AccountStorage) DeleteAccount(name string, ctx context.Context) error {
	_, err := a.client.Delete("accounts", name).Do(ctx)
	return err
}
