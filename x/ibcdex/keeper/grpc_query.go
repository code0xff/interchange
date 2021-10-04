package keeper

import (
	"github.com/code0xff/interchange/x/ibcdex/types"
)

var _ types.QueryServer = Keeper{}
