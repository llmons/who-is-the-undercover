package repo

import (
	"github.com/google/wire"
	"github.com/llmons/who-is-the-undercover/internal/base/data"
)

var ProviderSetRepo = wire.NewSet(
	data.NewData,
	data.NewDB,
	NewUndercoverRepo,
)
