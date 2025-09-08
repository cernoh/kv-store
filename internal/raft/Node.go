type Role int

const (
	Follower Role = iota
	Candidate
	Leader
)

type Term uint64
type Index uint64

type Node struct {
	mu sync.Mutex

	id       string
	peers    []string
	role     Role
	term     Term
	votedFor string

	electionTimeout   time.Duration
	heartbeatInterval time.Duration
	resetElectionCh   chan struct{}

	stopCh chan struct{}
}
