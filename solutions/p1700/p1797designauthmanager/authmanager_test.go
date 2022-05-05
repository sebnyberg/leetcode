package p1797designauthmanager

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TestAction func(t *testing.T, m *AuthenticationManager)

func Renew(tokenID string, currentTime int) TestAction {
	return func(t *testing.T, m *AuthenticationManager) {
		m.Renew(tokenID, currentTime)
	}
}

func CountUnexpiredTokens(currentTime int, want int) TestAction {
	return func(t *testing.T, m *AuthenticationManager) {
		res := m.CountUnexpiredTokens(currentTime)
		require.Equal(t, want, res)
	}
}

func Generate(tokenID string, currentTime int) TestAction {
	return func(t *testing.T, m *AuthenticationManager) {
		m.Generate(tokenID, currentTime)
	}
}

func TestAuthManager(t *testing.T) {
	for _, tc := range []struct {
		name    string
		ttl     int
		actions []TestAction
	}{
		{
			"testcase 60", 13,
			[]TestAction{
				Renew("ajvy", 1),
				CountUnexpiredTokens(3, 0),
				CountUnexpiredTokens(4, 0),
				Generate("fuzxq", 5),
				Generate("izmry", 7),
				Renew("puv", 12),
				Generate("ybiqb", 13),
				Generate("gm", 14),
				CountUnexpiredTokens(15, 4),
				CountUnexpiredTokens(18, 3),
				CountUnexpiredTokens(19, 3),
				Renew("ybiqb", 21),
				CountUnexpiredTokens(23, 2),
				CountUnexpiredTokens(25, 2),
				CountUnexpiredTokens(26, 2),
				Generate("aqdm", 28),
				CountUnexpiredTokens(29, 2),
				Renew("puv", 30),
			},
		},
		{
			"testcase 1", 6,
			[]TestAction{
				CountUnexpiredTokens(6, 0),
				CountUnexpiredTokens(5, 0),
				CountUnexpiredTokens(4, 0),
				Generate("kxlq", 9),
				Renew("avem", 10),
				CountUnexpiredTokens(15, 0),
			},
		},
		{
			"example 1", 5,
			[]TestAction{
				Renew("aaa", 1),
				Generate("aaa", 2),
				CountUnexpiredTokens(6, 1),
				Generate("bbb", 7),
				Renew("aaa", 8),
				Renew("bbb", 10),
				CountUnexpiredTokens(15, 0),
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := Constructor(tc.ttl)
			for _, action := range tc.actions {
				action(t, &m)
			}
		})
	}
}
