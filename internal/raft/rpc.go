package raft

import "context"

// abstraction of communication between raft nodes
type Transport interface {

	//requestvote RPC to a peer
	RequestVote(ctx context.Context, peerID string, req *RequestVoteReq) (*RequestVoteResp, error)

	//AppendEntries RPC to a peer
	AppendEntries(ctx context.Context, peerID string, req *AppendEntriesReq) (*AppendEntriesResp, error)
}

