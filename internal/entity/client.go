package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNameIsRequired         = errors.New("Name is required")
	ErrEmailIsRequired        = errors.New("Email is required")
	ErrAccountNotBelongClient = errors.New("account does not belong to client")
)

type Client struct {
	ID        string
	Name      string
	Email     string
	Accounts  []*Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(name, email string) (*Client, error) {
	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := client.Validate()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return ErrNameIsRequired
	}

	if c.Email == "" {
		return ErrEmailIsRequired
	}
	return nil
}

func (c *Client) Update(name, email string) error {
	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()

	err := c.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AddAccount(account *Account) error {
	if account.Client.ID != c.ID {
		return ErrAccountNotBelongClient
	}

	c.Accounts = append(c.Accounts, account)
	return nil
}
