# Raft Basics

## What is Raft?

> Raft is a consensus algorithm used in distributed systems to ensure that
> multiple nodes agree on a single source of truth, even in the presense of
> failures

Essentially, it makes sure that there are no conflicts between nodes, and we
decide on one leader to make decisions. It's almost like selecting a president
to make decisions on behalf of the people.

### Who uses it?

Kubernetes ! One of the biggest companies in the world.

## The state machine

We have 3 states that the nodes can be in

| Role      | Responsibilities                                                            |
| --------- | --------------------------------------------------------------------------- |
| Follower  | Passive, waits for messages, and times out if theres no leader contact      |
| Candidate | Starts an election, and asks peers to votes                                 |
| Leader    | Handles clent requests, replicates logs, and send heartbeats to other nodes |

## How does leader elections work?

Followers start an election timer. If it expires, the node becomes a candidate
and begins requesting votes. If a majority quorum is granted, meaning that more
than 50% of the nodes agree on this node, it becomes the Leader. The leader
begins sending heartbeats (AppendEntries RPCs), which helps maintain authority.

```
            +-------------------+
            |                   |
            |   (heartbeat)     |
            v                   |
      +-----------+   timeout   |
      | Follower  | -----------+
      +-----------+             |
            |                   |
            | election timeout  |
            v                   |
      +------------+            |
      | Candidate  |            |
      +------------+            |
        |   ^     |             |
        |   |     |             |
 votes  |   | no  | yes (majority)
denied  |   |     v
        |   | +-----------+
        |   +-|  Leader   |
        |     +-----------+
        |           |
        |<-- appendEntries
        v           |
   +-----------+ <--+
   | Follower  |
   +-----------+
```

## How am I implementing it?

Im putting it into several files.

`types.go`, this is what defines roles, terms, log entries, and RPC messages.

`nodes.go`, This is the core struct, and helps with role switching and such. It
helps makes a good scaffold

`election.go`, this is the general election logic. Candidates request votes,
collects majority votes, and things of that nature

`heartbeat.go`, this is the pulses the leader sends out. It resets the follower
timers, to make sure that they dont start an election

`rpc.go`, abstracts the transportation, works in memory and over HTTP.

## How do we deal with failure?

If the leader fails, the heartbeat will not be transmitted anymore. Therefore, a
candidate election will sbe started, and a new leader will therefore be
selected.

## What after?

Log replication, Applying comitted entries to the KV store, and a client-facing
api!
