package clients

import (
	"context"
	"fmt"
)

// LoadInitialClients returns an initial Client Repository state populated with clients from the internal storage.
func LoadInitialClients(ctx context.Context, p ClientStore) ([]*Client, error) {
	all, err := p.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get clients: %v", err)
	}

	// mark previously connected clients as disconnected with current time
	now := now()
	for _, cur := range all {
		if cur.DisconnectedAt == nil {
			cur.SetDisconnected(&now)
			err := p.Save(ctx, cur)
			if err != nil {
				return nil, fmt.Errorf("failed to save client: %v", err)
			}
		}
	}

	return all, nil
}
