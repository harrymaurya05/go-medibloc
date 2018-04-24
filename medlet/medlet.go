package medlet

import (
	"github.com/medibloc/go-medibloc/core"
	"github.com/medibloc/go-medibloc/core/pb"
	"github.com/medibloc/go-medibloc/medlet/pb"
	mednet "github.com/medibloc/go-medibloc/net"
	"github.com/medibloc/go-medibloc/rpc"
	"github.com/medibloc/go-medibloc/storage"
	"github.com/medibloc/go-medibloc/util/logging"
	m "github.com/rcrowley/go-metrics"
	"github.com/sirupsen/logrus"
)

var (
	metricsMedstartGauge   = m.GetOrRegisterGauge("med.start", nil)
	transactionManagerSize = 1280
)

// Medlet manages blockchain services.
type Medlet struct {
	bs         *core.BlockSubscriber
	config     *medletpb.Config
	miner      *core.Miner
	netService mednet.Service
	rpc        rpc.GRPCServer
	txMgr      *core.TransactionManager
	storage    storage.Storage
	genesis    *corepb.Genesis
}

type rpcBridge struct {
	bm    *core.BlockManager
	txMgr *core.TransactionManager
}

// BlockManager return core.BlockManager
func (rb *rpcBridge) BlockManager() *core.BlockManager {
	return rb.bm
}

func (rb *rpcBridge) TransactionManager() *core.TransactionManager {
	return rb.txMgr
}

// New returns a new medlet.
func New(config *medletpb.Config) (*Medlet, error) {
	return &Medlet{
		config: config,
	}, nil
}

// Setup sets up medlet.
func (m *Medlet) Setup() {
	var err error
	logging.Console().Info("Setting up Medlet...")

	m.netService, err = mednet.NewMedService(m)
	if err != nil {
		logging.Console().WithFields(logrus.Fields{
			"err": err,
		}).Fatal("Failed to setup net service.")
	}

	m.storage, err = storage.NewLeveldbStorage(m.Config().Chain.Datadir)
	if err != nil {
		logging.Console().WithFields(logrus.Fields{
			"err": err,
		}).Fatal("Failed to create leveldb storage.")
	}

	logging.Console().Info("Set up Medlet.")
}

// Start starts the services of the medlet.
func (m *Medlet) Start() {
	if err := m.netService.Start(); err != nil {
		logging.Console().WithFields(logrus.Fields{
			"err": err,
		}).Fatal("Failed to start net service.")
		return
	}

	metricsMedstartGauge.Update(1)

	m.txMgr = core.NewTransactionManager(m, transactionManagerSize)
	m.txMgr.RegisterInNetwork(m.netService)
	m.txMgr.Start()

	bp, bc, err := core.GetBlockPoolBlockChain(m.storage)
	if err != nil {
		logging.Console().WithFields(logrus.Fields{
			"err": err,
		}).Fatal("Failed to create block pool or block chain")
		return
	}
	m.bs = core.StartBlockSubscriber(m.netService, bp, bc)

	if m.Config().Chain.StartMine {
		m.miner = core.StartMiner(m.netService, bc, m.txMgr)
	}

	m.rpc = rpc.NewServer(&rpcBridge{bm: m.bs.BlockManager(), txMgr: m.txMgr}, m.config.Rpc.RpcListen[0]) // TODO choose index
	m.rpc.Start()
	m.rpc.RunGateway(m.config.Rpc.HttpListen[0]) // TODO choose index

	logging.Console().Info("Started Medlet.")
}

// Stop stops the services of the medlet.
func (m *Medlet) Stop() {
	if m.netService != nil {
		m.netService.Stop()
		m.netService = nil
	}

	m.txMgr.Stop()

	m.bs.StopBlockSubscriber()
	if m.miner != nil {
		m.miner.StopMiner()
	}

	m.rpc.Stop()

	logging.Console().Info("Stopped Medlet.")
}

// Config returns medlet configuration.
func (m *Medlet) Config() *medletpb.Config {
	return m.config
}

// Storage returns storage.
func (m *Medlet) Storage() storage.Storage {
	return m.storage
}

// Genesis returns genesis config.
func (m *Medlet) Genesis() *corepb.Genesis {
	return m.genesis
}
