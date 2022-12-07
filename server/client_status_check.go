package chserver

import (
	"context"
	"time"

	"github.com/cloudradar-monitoring/rport/server/clients"
	"github.com/cloudradar-monitoring/rport/share/comm"
	"github.com/cloudradar-monitoring/rport/share/logger"
)

type ClientsStatusCheckTask struct {
	log         *logger.Logger
	cr          *clients.ClientRepository
	th          time.Duration // Threshold after which a client to server ping is considered outdated.
	pingTimeout time.Duration // Don't wait longer than pingTimeout for a response
}

// NewClientsStatusCheckTask pings all active clients and marks them disconnected on ping failure
func NewClientsStatusCheckTask(log *logger.Logger, cr *clients.ClientRepository, th time.Duration, pingTimeout time.Duration) *ClientsStatusCheckTask {
	return &ClientsStatusCheckTask{
		log:         log.Fork("clients-status-check"),
		cr:          cr,
		th:          th,
		pingTimeout: pingTimeout,
	}
}

func (t *ClientsStatusCheckTask) Run(ctx context.Context) error {
	timerStart := time.Now()
	var dueClients []*clients.Client
	var confirmedClients = 0
	var now = time.Now()
	for _, c := range t.cr.GetAllActive() {
		// Shorten the threshold aka make heartbeat older than it is because the ping response is stored after this check.
		// Clients would get checked only every second time otherwise.
		if c.LastHeartbeatAt != nil && now.Sub(*c.LastHeartbeatAt) < t.th-10*time.Second {
			// Skip all clients having sent a heartbeat from client to server recently
			confirmedClients++
			continue
		}
		dueClients = append(dueClients, c)
	}
	if len(dueClients) == 0 {
		// Nothing to do
		t.log.Debugf("ended after %s, no clients to ping", time.Since(timerStart))
		return nil
	}
	maxWorkers := 100
	if maxWorkers > len(dueClients) {
		maxWorkers = len(dueClients)
	}
	clientsToPing := make(chan *clients.Client, len(dueClients))
	results := make(chan bool, len(dueClients))
	for w := 1; w <= maxWorkers; w++ {
		go t.PingClients(clientsToPing, results)
	}
	for _, dueClient := range dueClients {
		clientsToPing <- dueClient
	}
	close(clientsToPing)
	var dead = 0
	var alive = 0
	for a := 0; a < len(dueClients); a++ {
		if <-results {
			alive++
		} else {
			dead++
		}
	}
	t.log.Debugf("ended after %s, skipped: %d, pinged: %d, alive: %d, dead: %d", time.Since(timerStart), confirmedClients, len(dueClients), alive, dead)
	return nil
}

func (t *ClientsStatusCheckTask) PingClients(clientsToPing <-chan *clients.Client, results chan<- bool) {
	for cl := range clientsToPing {
		ok, response, rtt, err := comm.PingConnectionWithTimeout(cl.Connection, t.pingTimeout)
		//t.log.Debugf("ok=%s, error=%s, response=%s", ok, err, response)
		var now = time.Now()
		//Old clients cannot respond properly to a ping request yet
		if !ok && err == nil && string(response) == "unknown request" {
			t.log.Debugf("ping to %s [%s] succeeded in %s. client < 0.8.2", cl.Name, cl.ID, rtt)
			cl.LastHeartbeatAt = &now
			results <- true
			continue
		}
		// Only an empty response confirms the ping
		if ok && err == nil && len(response) == 0 {
			t.log.Debugf("ping to %s [%s] succeeded in %s. client >= 0.8.2", cl.Name, cl.ID, rtt)
			cl.LastHeartbeatAt = &now
			results <- true
			continue
		}
		// None of the above. Ping must have failed or timed out.
		t.log.Infof("ping to %s [%s] failed: %s", cl.Name, cl.ID, err)

		cl.SetDisconnected(&now)

		cl.Close()
		results <- false
	}
}
