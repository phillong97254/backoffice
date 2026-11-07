package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (i *impl) PasswordToken(ctx context.Context, input PasswordTokenInput) (PasswordTokenOutput, error) {
	form := url.Values{"grant_type": {"password"}, "username": {input.Username}, "password": {input.Password}}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, i.tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return PasswordTokenOutput{}, fmt.Errorf("password token: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := i.http.Do(req)
	if err != nil {
		return PasswordTokenOutput{}, fmt.Errorf("password token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PasswordTokenOutput{}, ErrInvalidCredentials
	}
	var out PasswordTokenOutput
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return PasswordTokenOutput{}, fmt.Errorf("password token: decode: %w", err)
	}
	return out, nil
}
