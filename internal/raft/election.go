package raft

import (
	"context"
	"sync"
	"time"
)

func (n *Node) runCandidate() {
	n.mu.Lock()
	n.term++
	term := n.term
	n.votedFor = n.id
	peers := append([]string(nil), n.peers...)
	n.mu.Unlock()

	votes := 1
	majority := (len(peers)+1)/2 + 1

	var wg sync.WaitGroup
	voteCh := make(chan bool, len(peers))

	for _, peer := range peers {
		wg.Add(1)
		go func(peerID string) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
			defer cancel()

			req := &RequestVoteReq{
				Term:         term,
				CandidateID:  n.id,
				LastLogIndex: 0, //TODO: hook into real log
				LastLogTerm:  0,
			}

			resp, err := n.transport.RequestVote(ctx, peerID, req)
			if err != nil || resp == nil {
				voteCh <- false
				return
			}

			n.mu.Lock()
			if resp.Term > n.term {
				n.term = resp.Term
				n.role = Follower
				n.votedFor = ""
			}

			n.mu.Unlock()
			voteCh <- resp.VoteGranted
		}(peer)
	}

	go func() {
		wg.Wait()
		close(voteCh)

	}()

	for granted := range voteCh {
		if granted {
			votes++
		}
		if votes >= majority {
			n.mu.Lock()
			n.role = Leader
			n.mu.Unlock()
			return
		}
	}

	n.mu.Lock()
	n.role = Follower
	n.mu.Unlock()
}

