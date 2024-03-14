package components

import "laatoo.io/sdk/server/core"

type ExpressionsManager interface {
	RegisterExpression(ctx core.ServerContext, expression core.Expression) error
	GetExpressionValue(ctx core.RequestContext, expression core.Expression, vars core.StringMap) (interface{}, error)
}
