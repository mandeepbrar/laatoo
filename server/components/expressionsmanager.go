package components

import "laatoo.io/sdk/server/core"

type ExpressionsManager interface {
	RegisterExpression(ctx core.ServerContext, expression *core.Expression) error
	GetExpressionValue(ctx core.ServerContext, expression *core.Expression) (interface{}, error)
}
