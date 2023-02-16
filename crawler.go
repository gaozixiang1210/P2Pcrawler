package crawler

import (
	"context"
	"database/sql"
	"github.com/ethereum/go-ethereum/p2p/enode"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"sync"
	"time"
)

const (
	Interval          = 6 * time.Second  //crawl interval for each node
	RoundInterval     = 10 * time.Second //crawl interval for each round
	DefaultTimeout    = 1 * time.Hour    //check interval for all nodes
	DefaultChanelSize = 512
	seedCount         = 30
	seedMaxAge        = 5 * 24 * time.Hour
)

type Crawler struct {
	BootNodes    []*enode.Node // BootNodes is the set of nodes that the crawler will start from.
	CurrentNodes nodeSet       // CurrentNodes is the set of nodes that the crawler is currently crawling.

	ReqCh    chan *enode.Node // ReqCh is the channel that the crawler uses to send requests to the workers.
	OutputCh chan *Node       // OutputCh is the channel that the crawler uses to send requests to the filter.

	leveldb   *enode.DB          // leveldb is the database that the crawler uses to store the nodes.
	db        *sql.DB            // db is the database that the crawler uses to store the nodes.
	tableName string             // tableName is the name of the table that the crawler will use to store the nodes.
	mu        sync.Mutex         // mu is the mutex that protects the crawler.
	ctx       context.Context    // ctx is the context that the crawler uses to cancel all crawl.
	cancel    context.CancelFunc // cancel is the function that the crawler uses to cancel all crawl.

	logger *zap.Logger // logger is the logger that the crawler uses to log the information.
}
