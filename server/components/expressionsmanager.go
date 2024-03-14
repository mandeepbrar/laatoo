package components

import (
	"laatoo.io/sdk/ctx"
	"laatoo.io/sdk/datatypes"
	"laatoo.io/sdk/server/core"
)

type ExpressionsManager interface {
	RegisterExpression(ctx core.ServerContext, expression core.Expression, dtype datatypes.DataType) error
	GetExpressionValue(ctx ctx.Context, expression core.Expression, vars core.StringMap) (interface{}, error)
}
