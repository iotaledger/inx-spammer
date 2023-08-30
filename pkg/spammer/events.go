package spammer

import (
	"github.com/iotaledger/hive.go/runtime/event"
)

// SpamStats are stats for a single spam block.
type SpamStats struct {
	Tipselection float32 `json:"tipselect"`
	ProofOfWork  float32 `json:"pow"`
}

// AvgSpamMetrics are average metrics of the created spam.
type AvgSpamMetrics struct {
	NewBlocks              uint32  `json:"newBlocks"`
	AverageBlocksPerSecond float32 `json:"avgBlocks"`
}

// Events are the events issued by the spammer.
type Events struct {
	// Fired when a single spam block is issued.
	SpamPerformed *event.Event1[*SpamStats]
	// Fired when average spam metrics were updated by the worker.
	AvgSpamMetricsUpdated *event.Event1[*AvgSpamMetrics]
}
