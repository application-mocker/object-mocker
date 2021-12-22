package tree

import (
	"object-mocker/pkg/model"
	"object-mocker/utils"
	"regexp"
	"strings"
	"sync"
)

const (
	ScopeSeparator = "/"
	ScopeRegex     = "^[a-zA-Z0-9\\_\\<\\>]*$"
)

var scopeRegex *regexp.Regexp

func init() {
	scopeRegex, _ = regexp.Compile(ScopeRegex)
}

// Node is one of tree node. Special, root node haven't parent-node.
// For safe to json, parent will ignore when to json. All the operator of node is not thread-safe.
type Node struct {
	Parent   *Node            `json:"-"`
	Scope    string           `json:"scope"`
	Children map[string]*Node `json:"children"`

	Data map[string]*model.Data `json:"data"`

	freeze *sync.RWMutex
}

// ToJsonWithPrune will freeze node and prune empty node.
func ToJsonWithPrune(n *Node) string {
	if n == nil {
		return ""
	}

	n.FreezeNode()
	defer n.UnFreeze()
	str, err := n.ToJson()
	if err != nil {
		return ""
	}

	return str
}

func (n *Node) FreezeNode() {
	n.freeze.Lock()
	for _, nodes := range n.Children {
		nodes.FreezeNode()
	}
}

func (n *Node) UnFreeze() {
	n.freeze.Unlock()
	for _, nodes := range n.Children {
		nodes.UnFreeze()
	}
}

func newNode(parent *Node, scope string) *Node {
	return &Node{
		Parent:   parent,
		Scope:    scope,
		Children: map[string]*Node{},
		Data:     map[string]*model.Data{},
		freeze:   &sync.RWMutex{},
	}
}

// NewRoot return a root node of tree
func NewRoot() *Node {
	return newNode(nil, "")
}

// Node return special Node from node. And scope muse follow in ScopeRegex.
// the node should a
func (n *Node) node(scope string) *Node {
	n.freeze.RLock()
	defer n.freeze.RUnlock()
	if !scopeRegex.Match([]byte(scope)) {
		return nil
	}
	if n.Children[scope] == nil {
		n.Children[scope] = newNode(n.Parent, scope)
	}
	return n.Children[scope]
}

// nodeWithScopes return special Node by scopes in format: "scope1.scope2.scope3".
// All the scopes will be split by ScopeSeparator.
func (n *Node) nodeWithScopes(scopes string) *Node {
	n.freeze.RLock()
	defer n.freeze.RUnlock()
	if strings.Index(scopes, ScopeSeparator) == -1 {
		return n.node(scopes)
	}
	firstScopeEndIndex := strings.Index(scopes, ScopeSeparator)
	currentCope := scopes[0:firstScopeEndIndex]
	if ns := n.node(currentCope); ns != nil {
		return ns.nodeWithScopes(scopes[firstScopeEndIndex+1:])
	}
	return nil
}

// removeNode will remove node from parent. If scope not exits, return nil.
func (n *Node) removeNode(scope string) *Node {
	if child, ok := n.Children[scope]; ok {
		delete(n.Children, scope)
		return child
	}

	return nil
}

// prune will delete all empty node in the tree.
func (n *Node) prune() {
	oldMap := n.Children
	for scope, node := range oldMap {
		node.prune()

		if len(node.Data) == 0 && len(node.Children) == 0 {
			n.removeNode(scope)
		}
	}
}

// SetChildrenParent set children-parent of n to n
func (n *Node) SetChildrenParent() {
	for _, child := range n.Children {
		child.Parent = n
		child.SetChildrenParent()
	}
}

// ToJson string
func (n *Node) ToJson() (string, error) {
	return utils.ToJson(n)
}

//String return json string
func (n *Node) String() string {
	if r, err := n.ToJson(); err == nil {
		return r
	}
	return ""
}

// AppendData to append a new data
func (n *Node) AppendData(scopes string, data *model.Data) *model.Data {
	n.freeze.RLock()
	defer n.freeze.RUnlock()
	node := n.nodeWithScopes(scopes)
	if node == nil {
		utils.Logger.Error("[AppendData]: not found node: {%s}", scopes)
		return nil
	}

	node.Data[data.Id] = data
	return data
}

// NewData package the value to new data and append it.
func (n *Node) NewData(scopes string, value map[string]interface{}) *model.Data {
	n.freeze.RLock()
	defer n.freeze.RUnlock()
	data, err := model.NewDataWithDataValue(value)
	if err != nil {
		utils.Logger.Error(err)
		return nil
	}

	return n.AppendData(scopes, data)
}

// DeleteData delete data from n
func (n *Node) DeleteData(scopes, id string) *model.Data {
	n.freeze.RLock()
	defer n.freeze.RUnlock()
	node := n.nodeWithScopes(scopes)
	if node == nil {
		utils.Logger.Error("[DeleteData]: Not found node: {%s}", scopes)
		return nil
	}
	if _, ok := node.Data[id]; !ok {
		utils.Logger.Error("[DeleteData]: node found data: {%s} in node :{%s}", id, scopes)
		return nil
	}

	node.Data[id].Delete()
	data := node.Data[id]
	delete(node.Data, id)

	return data
}

// UpdateData update a data with scopes and id
func (n *Node) UpdateData(scopes, id string, value map[string]interface{}) *model.Data {
	n.freeze.RLock()
	defer n.freeze.RUnlock()
	node := n.nodeWithScopes(scopes)
	if node == nil {
		utils.Logger.Error("[UpdateData]: Not found node: {%s}", scopes)
		return nil
	}
	if _, ok := node.Data[id]; !ok {
		utils.Logger.Error("[UpdateData]: node found data: {%s} in node :{%s}", id, scopes)
		return nil
	}
	node.Data[id].DataValue = value
	node.Data[id].UpdateAt = utils.NowTime()

	return node.Data[id]
}
