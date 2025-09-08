package raft

import "time"

type Role int

const (
	Follower Role = iota
	Candidate
	Leader
)

type Term uint64
type Index uint64

type LogEntry struct {
	Term    Term
	Index   Index
	Command []byte
}

// RPC req/res types

type RequestVoteReq struct {
	Term         Term
	CandidateID  string
	LastLogIndex Index
	LastLogTerm  Term
}

type RequestVoteResp struct {
	Term        Term
	VoteGranted bool
}

type AppendEntriesReq struct {
	Term         Term
	LeaderID     string
	PrevLogIndex Index
	PrevLogTerm  Term
	Entries      []LogEntry
	LeaderCommit Index
}

type AppendEntriesResp struct {
	Term    Term
	Success bool
}

type Config struct {
	ElectionTimeoutMin time.Duration
	ElectionTimeoutMax time.Duration
	HeartbeatInterval  time.Duration
}
