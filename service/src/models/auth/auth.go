package auth

import (
	authGoCore "auth-go-core"
	"auth-go-core/strategies/do"
	"auth-go-core/strategies/github"
	"auth-go-core/strategies/google"
)

func SelectStrategy(name string, data *authGoCore.StrategyData) *authGoCore.AuthGoCore {
	switch name {
	case "do":
		return &authGoCore.AuthGoCore{
			Data:     *data,
			Strategy: &do.Strategy {},
		}
	case "google":
		return &authGoCore.AuthGoCore{
			Data: *data,
			Strategy: &google.Strategy{},
		}
	case "github": {
		return &authGoCore.AuthGoCore{
			Data: *data,
			Strategy: &github.Strategy{},
		}
	}
	default:
		return nil
	}
}