package txpool

import (
	"github.com/hashicorp/go-hclog"
	"konsta.live/types"
)

type ReApplyStorageConfig struct {
}

// TxReApplyStoragePool is a module that handles pending transactions.
// All transactions are handled within their respective accounts.
// An account contains 2 queues a transaction needs to go through:
// - 1. Enqueued (entry point)
// - 2. Promoted (exit point)
// (both queues are min nonce ordered)
//
// When consensus needs to process promoted transactions,
// the pool generates a queue of "executable" transactions. These
// transactions are the first-in-line of some promoted queue,
// ready to be written to the state (primaries).
type TxReApplyStoragePool struct {
	logger      hclog.Logger // all the primaries sorted by max gas price
	executables *pricedQueue
}

// NewTxReApplyStoragePool returns a new pool for processing incoming transactions.
func NewTxReApplyStoragePool(
	logger hclog.Logger,
	config *ReApplyStorageConfig,
) (*TxReApplyStoragePool, error) {
	reApplyPool := &TxReApplyStoragePool{
		logger:      logger.Named("tx-reapply-storage"),
		executables: newPricedQueue(),
	}

	// thread run there
	// Attach the event manager
	//pool.eventManager = newEventManager(pool.logger)

	return reApplyPool, nil
}

// Prepare generates all the transactions
// ready for execution. (primaries)
func (p *TxReApplyStoragePool) Prepare() {
	// clear from previous round
	if p.executables.length() != 0 {
		p.executables.clear()
	}
}

// Push removes the given transaction from the
// associated promoted queue (account).
// Will update executables with the next primary
// from that account (if any).
func (p *TxReApplyStoragePool) Push(tx *types.Transaction) {

	if tx != nil {
		p.executables.push(tx)
	}
}

// Length returns the total number of all promoted transactions.
func (p *TxReApplyStoragePool) Length() uint64 {
	return p.executables.length()
}
