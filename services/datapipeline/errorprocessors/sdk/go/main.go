package errorprocessors

import (
	"laatoo/sdk/modules/datapipeline"
	"laatoo/sdk/server/core"
)

type ErrorRecordsStore interface {
	GetErrorRecords(ctx core.RequestContext) []*datapipeline.PipelineRecord
}
