package captain

import (
	"context"
)

type Account struct {
	UUID         *string `json:"uuid"`
	FriendlyName *string `json:"friendly_name"`
}

func (c *Client) GetAccounts(ctx context.Context) ([]*Account, error) {
	req, err := c.NewRequest("GET", "/v1/accounts", nil)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	err = c.Do(ctx, req, &accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}