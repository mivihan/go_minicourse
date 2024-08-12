package account

import (
	"go_minicourse/HW2/dto/"
	"go_minicourse/HW2/server/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

// Handler структура для обработки запросов
type Handler struct {
	accounts  map[string]*model.Account
	mutex     *sync.RWMutex
	secretKey string
}

// New создает новый экземпляр Handler
func New(secretKey string) *Handler {
	return &Handler{
		accounts:  make(map[string]*model.Account),
		mutex:     &sync.RWMutex{},
		secretKey: secretKey,
	}
}

// Actuator проверка состояния сервера
func (h *Handler) Actuator(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// GetAll возвращает список всех аккаунтов
func (h *Handler) GetAll(c echo.Context) error {
	if c.QueryParam("secret-key") != h.secretKey {
		return c.NoContent(http.StatusForbidden)
	}

	h.mutex.RLock()
	defer h.mutex.RUnlock()

	var accountList []dto.GetAccountResponse
	for _, acc := range h.accounts {
		accountList = append(accountList, dto.GetAccountResponse{
			Name:   acc.Name,
			Amount: acc.Amount,
		})
	}

	return c.JSON(http.StatusOK, accountList)
}

// CreateAccount создает новый аккаунт
func (h *Handler) CreateAccount(c echo.Context) error {
	var req dto.CreateAccountRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	if req.Name == "" {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.mutex.Lock()
	defer h.mutex.Unlock()

	if _, exists := h.accounts[req.Name]; exists {
		return c.String(http.StatusForbidden, "account already exists")
	}

	h.accounts[req.Name] = &model.Account{
		Name:   req.Name,
		Amount: req.Amount,
	}

	return c.NoContent(http.StatusCreated)
}

// GetAccount возвращает информацию об аккаунте
func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParam("name")

	h.mutex.RLock()
	defer h.mutex.RUnlock()

	acc, exists := h.accounts[name]
	if !exists {
		return c.String(http.StatusNotFound, "account "+name+" not found")
	}

	return c.JSON(http.StatusOK, dto.GetAccountResponse{
		Name:   acc.Name,
		Amount: acc.Amount,
	})
}

// DeleteAccount удаляет аккаунт
func (h *Handler) DeleteAccount(c echo.Context) error {
	var req dto.DeleteAccountRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	h.mutex.Lock()
	defer h.mutex.Unlock()

	if _, exists := h.accounts[req.Name]; !exists {
		return c.String(http.StatusNotFound, "account "+req.Name+" not found")
	}

	delete(h.accounts, req.Name)
	return c.NoContent(http.StatusOK)
}

// PatchAccount обновляет баланс аккаунта
func (h *Handler) PatchAccount(c echo.Context) error {
	var req dto.PatchAccountRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	h.mutex.Lock()
	defer h.mutex.Unlock()

	acc, exists := h.accounts[req.Name]
	if !exists {
		return c.String(http.StatusNotFound, "account "+req.Name+" not found")
	}

	acc.Amount = req.Amount
	return c.NoContent(http.StatusOK)
}

// ChangeAccount изменяет имя аккаунта
func (h *Handler) ChangeAccount(c echo.Context) error {
	var req dto.ChangeAccountRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	if req.NewName == "" {
		return c.String(http.StatusBadRequest, "empty new name")
	}

	h.mutex.Lock()
	defer h.mutex.Unlock()

	acc, exists := h.accounts[req.Name]
	if !exists {
		return c.String(http.StatusNotFound, "account "+req.Name+" not found")
	}

	if _, newExists := h.accounts[req.NewName]; newExists {
		return c.String(http.StatusForbidden, "account "+req.NewName+" already exists")
	}

	acc.Name = req.NewName
	h.accounts[req.NewName] = acc
	delete(h.accounts, req.Name)

	return c.NoContent(http.StatusOK)
}
