package txpool

import (
	"time"

	"github.com/hashicorp/go-hclog"
	"konsta.live/types"
)

const (
	reApplyCoolDown = 5000 * time.Millisecond
)

type ReApplyConfig struct {
	ReApplyCountDown uint64
}

// TxReApplyPool is a module that handles pending transactions.
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
type TxReApplyPool struct {
	logger           hclog.Logger // all the primaries sorted by max gas price
	executables      *pricedQueue
	txpool           *TxPool
	txreaaplystorage *TxReApplyStoragePool
	reApplyCountDown uint64
}

// NewTxReApplyPool returns a new pool for processing incoming transactions.
func NewTxReApplyPool(
	logger hclog.Logger,
	txpool *TxPool,
	txreaaplystorage *TxReApplyStoragePool,
	config *ReApplyConfig,
) (*TxReApplyPool, error) {
	reApplyPool := &TxReApplyPool{
		logger:           logger.Named("tx-reapply-pool"),
		executables:      newPricedQueue(),
		txpool:           txpool,
		txreaaplystorage: txreaaplystorage,
		reApplyCountDown: config.ReApplyCountDown,
	}

	// thread run there
	// Attach the event manager
	//pool.eventManager = newEventManager(pool.logger)

	return reApplyPool, nil
}

func reApplyFunction(logger hclog.Logger, txpool *TxPool, txstorage *TxReApplyStoragePool) {
	for i := 0; i < 100000; i++ {
		time.Sleep(2 * time.Second)
		logger.Info("reapply tx`s", txstorage.Length())
		if txstorage.executables.length() != 0 {
			var txs = txstorage.executables.pop()
			err := txpool.AddTx(txs)
			if err != nil {
				logger.Error("reapply tx apply error")
			}
			txstorage.executables.clear()
			txstorage.Push(txs)
		}
	}
}

// Prepare generates all the transactions
// ready for execution. (primaries)
func (p *TxReApplyPool) Prepare() {
	// clear from previous round
	if p.executables.length() != 0 {
		p.executables.clear()
	}

	go reApplyFunction(p.logger, p.txpool, p.txreaaplystorage)
}

// Push removes the given transaction from the
// associated promoted queue (account).
// Will update executables with the next primary
// from that account (if any).
func (p *TxReApplyPool) Push(tx *types.Transaction) {

	if tx != nil {
		p.executables.push(tx)
	}
}

// Length returns the total number of all promoted transactions.
func (p *TxReApplyPool) Length() uint64 {
	return p.executables.length()
}
