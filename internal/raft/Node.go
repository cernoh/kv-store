package raft

import (
	"math/rand"
	"sync"
	"time"
)

type Node struct {
	mu sync.Mutex

	id       string
	peers    []string
	role     Role
	term     Term
	votedFor string

	log []LogEntry

	commitIndex Index
	lastApplied Index

	//timing
	electionTimeout   time.Duration
	heartbeatInterval time.Duration
	resetElectionCh   chan struct{}

	stopCh chan struct{}
}

func NewNode(id string, peers []string, cfg Config) *Node {

	return &Node{
		id:                id,
		peers:             peers,
		role:              Follower,
		term:              0,
		votedFor:          "",
		log:               make([]LogEntry, 0),
		electionTimeout:   randomElectionTimeout(cfg.ElectionTimeoutMin, cfg.ElectionTimeoutMax),
		heartbeatInterval: cfg.HeartbeatInterval,
		resetElectionCh:   make(chan struct{}),
		stopCh:            make(chan struct{}),
	}
}

func randomElectionTimeout(min, max time.Duration) time.Duration {
	delta := int(max.Milliseconds() - min.Milliseconds())
	return min + time.Duration(rand.Intn(delta))*time.Millisecond
}

func (n *Node) Run() {
	for {
		n.mu.Lock()
		role := n.role
		n.mu.Unlock()

		switch role {
		case Follower:
			n.runFollower()
		case Candidate:
			n.runCandidate()
		case Leader:
			n.runLeader()
		}

		select {
		case <-n.stopCh:
			return
		default:
		}
	}
}

func (n *Node) runFollower() {
	timeout := time.NewTimer(n.electionTimeout)
	defer timeout.Stop()

	select {
	case <-n.resetElectionCh: // heartbeat received
		return
	case <-timeout.C: // election timeout
		n.mu.Lock()
		n.role = Candidate
		n.mu.Unlock()
	case <-n.stopCh:
		return
	}
}

// candidate starts election (stubbed for now)
func (n *Node) runCandidate() {
	n.mu.Lock()
	n.term++
	n.votedFor = n.id
	n.mu.Unlock()

	// TODO: send RequestVote RPCs
	time.Sleep(300 * time.Millisecond)

	// For now: just promote self to leader if no competition
	n.mu.Lock()
	n.role = Leader
	n.mu.Unlock()
}

// leader sends heartbeats
func (n *Node) runLeader() {
	ticker := time.NewTicker(n.heartbeatInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// TODO: send AppendEntries RPCs to peers
			// For now, just reset self
		case <-n.stopCh:
			return
		}
	}
}

func (n *Node) Stop() {
	close(n.stopCh)
}
