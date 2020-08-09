package repo

import (
	"database/sql"
	"sort"
	"time"

	"github.com/ansel1/merry"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/mreider/koto/backend/common"
)

type Node struct {
	ID            string       `db:"id"`
	Address       string       `db:"address"`
	AdminID       string       `db:"admin_id"`
	AdminName     string       `db:"admin_name"`
	AdminAvatarID string       `db:"admin_avatar_id"`
	CreatedAt     time.Time    `db:"created_at"`
	ApprovedAt    sql.NullTime `db:"approved_at"`
	DisabledAt    sql.NullTime `db:"disabled_at"`
	Details       string       `db:"details"`
	PostLimit     int          `db:"post_limit"`
}

type ConnectedNode struct {
	Node        Node
	MinDistance int
	Count       int
}

type NodeRepo interface {
	NodeExists(address string) (bool, error)
	AddNode(address, details string, nodeAdmin User, postLimit int) (string, error)
	AllNodes() ([]Node, error)
	Nodes(user User) ([]Node, error)
	Node(nodeID string) (*Node, error)
	ApproveNode(nodeID string) error
	RemoveNode(nodeID string) error
	ConnectedNodes(user User) ([]ConnectedNode, error)
	SetNodePostLimit(nodeAdminID, nodeID string, postLimit int) error
	AssignUserToNode(userID, nodeID string) error
	UserNodes(userIDs []string) (map[string][]string, error)
}

type nodeRepo struct {
	db *sqlx.DB
}

var (
	ErrNodeNotFound = common.ErrNotFound.WithMessage("node not found")
)

func NewNodes(db *sqlx.DB) NodeRepo {
	return &nodeRepo{
		db: db,
	}
}

func (r *nodeRepo) NodeExists(address string) (bool, error) {
	var nodeID string
	err := r.db.Get(&nodeID, `select id from nodes where address = $1`,
		address)
	if err != nil {
		if merry.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, merry.Wrap(err)
	}
	return true, nil
}

func (r *nodeRepo) AddNode(address, details string, nodeAdmin User, postLimit int) (string, error) {
	nodeID, err := uuid.NewV4()
	if err != nil {
		return "", merry.Wrap(err)
	}

	if postLimit < 0 {
		postLimit = 0
	}

	_, err = r.db.Exec(`
		insert into nodes(id, address, admin_id, created_at, details, post_limit) 
		VALUES ($1, $2, $3, $4, $5, $6)`,
		nodeID.String(), address, nodeAdmin.ID, common.CurrentTimestamp(), details, postLimit)
	if err != nil {
		return "", merry.Wrap(err)
	}
	return nodeID.String(), nil
}

func (r *nodeRepo) AllNodes() ([]Node, error) {
	var nodes []Node
	err := r.db.Select(&nodes, `
			select n.id, n.address, n.admin_id, n.created_at, n.approved_at, n.disabled_at, n.details,
				   u.name admin_name, u.avatar_thumbnail_id admin_avatar_id, post_limit
			from nodes n
				inner join users u on u.id = n.admin_id`)
	for i := range nodes {
		nodes[i].Address = common.CleanPublicURL(nodes[i].Address)
	}
	return nodes, merry.Wrap(err)
}

func (r *nodeRepo) Nodes(user User) ([]Node, error) {
	var nodes []Node
	err := r.db.Select(&nodes, `
		select n.id, n.address, n.admin_id, n.created_at, n.approved_at, n.disabled_at, n.details,
				   u.name admin_name, u.avatar_thumbnail_id admin_avatar_id, post_limit
		from nodes n
			inner join users u on u.id = n.admin_id
		where n.admin_id = $1`, user.ID)
	for i := range nodes {
		nodes[i].Address = common.CleanPublicURL(nodes[i].Address)
	}
	return nodes, merry.Wrap(err)
}

func (r *nodeRepo) Node(nodeID string) (*Node, error) {
	var node Node
	err := r.db.Get(&node, `
		select id, address, admin_id, created_at, approved_at, disabled_at, details, post_limit
		from nodes
		where id = $1`, nodeID)
	if err != nil {
		if merry.Is(err, sql.ErrNoRows) {
			return nil, ErrNodeNotFound.Here()
		}
		return nil, merry.Wrap(err)
	}

	node.Address = common.CleanPublicURL(node.Address)
	return &node, nil
}

func (r *nodeRepo) ApproveNode(nodeID string) error {
	_, err := r.db.Exec(`
		update nodes
		set approved_at = $1
		where id = $2`,
		common.CurrentTimestamp(), nodeID)
	return merry.Wrap(err)
}

func (r *nodeRepo) RemoveNode(nodeID string) error {
	return common.RunInTransaction(r.db, func(tx *sqlx.Tx) error {
		_, err := tx.Exec(`
			delete from user_nodes
			where node_id = $1`,
			nodeID)
		if err != nil {
			return merry.Wrap(err)
		}

		_, err = tx.Exec(`
			delete from nodes
			where id = $1`,
			nodeID)
		return merry.Wrap(err)
	})
}

func (r *nodeRepo) ConnectedNodes(user User) (connectedNodes []ConnectedNode, err error) {
	type friend struct {
		MinDistance int
		Count       int
	}
	type friendPair struct {
		UserID   string `db:"user_id"`
		FriendID string `db:"friend_id"`
	}
	currentLevel := []string{user.ID}
	nextPairs := make([]friendPair, 0)
	friends := map[string]friend{user.ID: {MinDistance: 0, Count: 1}}
	friendPairs := map[friendPair]struct{}{}
	distance := 0
	for len(currentLevel) > 0 {
		nextPairs = nextPairs[:0]
		distance++
		query, args, err := sqlx.In(`
			select user_id, friend_id
			from friends
			where user_id in (?)`,
			currentLevel)
		if err != nil {
			return nil, merry.Wrap(err)
		}
		query = r.db.Rebind(query)
		err = r.db.Select(&nextPairs, query, args...)
		if err != nil {
			return nil, merry.Wrap(err)
		}

		currentLevel = currentLevel[:0]
		for _, pair := range nextPairs {
			if _, ok := friendPairs[pair]; ok {
				continue
			}
			if _, ok := friendPairs[friendPair{UserID: pair.FriendID, FriendID: pair.UserID}]; ok {
				continue
			}
			friendPairs[pair] = struct{}{}
			if f, ok := friends[pair.FriendID]; !ok {
				friends[pair.FriendID] = friend{MinDistance: distance, Count: 1}
				currentLevel = append(currentLevel, pair.FriendID)
			} else {
				friends[pair.FriendID] = friend{MinDistance: f.MinDistance, Count: f.Count + 1}
			}
		}
	}

	var nodes []Node
	err = r.db.Select(&nodes, `
		select id, address, admin_id, created_at, approved_at, disabled_at, details, post_limit
		from nodes
		where approved_at is not null and disabled_at is null`)
	if err != nil {
		return nil, merry.Wrap(err)
	}

	connectedNodes = make([]ConnectedNode, 0, 10)
	for _, node := range nodes {
		if friend, ok := friends[node.AdminID]; ok && (node.PostLimit <= 0 || friend.MinDistance < node.PostLimit) {
			node.Address = common.CleanPublicURL(node.Address)
			connectedNodes = append(connectedNodes, ConnectedNode{
				Node:        node,
				MinDistance: friend.MinDistance,
				Count:       friend.Count,
			})
		}
	}

	return connectedNodes, nil
}

func (r *nodeRepo) SetNodePostLimit(nodeAdminID, nodeID string, postLimit int) error {
	if postLimit < 0 {
		postLimit = 1
	}

	_, err := r.db.Exec(`
		update nodes
		set post_limit = $1
		where id = $2 and admin_id = $3`,
		postLimit, nodeID, nodeAdminID)
	return merry.Wrap(err)
}

func (r *nodeRepo) AssignUserToNode(userID, nodeID string) error {
	now := common.CurrentTimestamp()
	_, err := r.db.Exec(`
			insert into user_nodes(user_id, node_id, created_at, updated_at)
			values($1, $2, $3, $4)
			on conflict (user_id, node_id) do update set updated_at = $4;`,
		userID, nodeID, now, now)
	if err != nil {
		return merry.Wrap(err)
	}
	return nil
}

func (r *nodeRepo) UserNodes(userIDs []string) (map[string][]string, error) {
	query, args, err := sqlx.In(`
		select un.user_id, n.address node_address
		from user_nodes un
			inner join nodes n on n.id = un.node_id
		where un.user_id in (?)`, userIDs)
	if err != nil {
		return nil, merry.Wrap(err)
	}
	query = r.db.Rebind(query)
	var nodes []struct {
		UserID      string `db:"user_id"`
		NodeAddress string `db:"node_address"`
	}
	err = r.db.Select(&nodes, query, args...)
	if err != nil {
		return nil, merry.Wrap(err)
	}
	result := make(map[string][]string)
	for _, node := range nodes {
		result[node.NodeAddress] = append(result[node.NodeAddress], node.UserID)
	}
	for _, nodeUserIDs := range result {
		sort.Slice(nodeUserIDs, func(i, j int) bool {
			return nodeUserIDs[i] < nodeUserIDs[j]
		})
	}
	return result, nil
}
