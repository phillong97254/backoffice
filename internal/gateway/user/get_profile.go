package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (i *impl) GetProfile(ctx context.Context, input GetProfileInput) (GetProfileOutput, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, i.baseURL+"/v1/users/"+url.PathEscape(input.ID), nil)
	if err != nil {
		return GetProfileOutput{}, err
	}
	resp, err := i.http.Do(req) // header Authorization do oauth2 transport tự gắn
	if err != nil {
		return GetProfileOutput{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return GetProfileOutput{}, fmt.Errorf("status %d", resp.StatusCode)
	}
	var out GetProfileOutput
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return GetProfileOutput{}, fmt.Errorf("decode: %w", err)
	}
	return out, nil
}
