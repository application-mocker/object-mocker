package tree

import (
	"object-mocker/pkg/model"
	"object-mocker/utils"
	"regexp"
	"strings"
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

	Data []*model.Data `json:"data"`
}

func newNode(parent *Node, scope string) *Node {
	return &Node{
		Parent:   parent,
		Scope:    scope,
		Children: map[string]*Node{},
		Data:     []*model.Data{},
	}
}

// NewRoot return a root node of tree
func NewRoot() *Node {
	return newNode(nil, "")
}

// Node return special Node from node. And scope muse follow in ScopeRegex.
func (n *Node) Node(scope string) *Node {
	if !scopeRegex.Match([]byte(scope)) {
		return nil
	}
	if n.Children[scope] == nil {
		n.Children[scope] = newNode(n.Parent, scope)
	}
	return n.Children[scope]
}

// NodeWithScopes return special Node by scopes in format: "scope1.scope2.scope3".
// All the scopes will be split by ScopeSeparator.
func (n *Node) NodeWithScopes(scopes string) *Node {
	if strings.Index(scopes, ScopeSeparator) == -1 {
		return n.Node(scopes)
	}
	firstScopeEndIndex := strings.Index(scopes, ScopeSeparator)
	currentCope := scopes[0:firstScopeEndIndex]
	if ns := n.Node(currentCope); ns != nil {
		return ns.NodeWithScopes(scopes[firstScopeEndIndex+1:])
	}
	return nil
}

// RemoveNode will remove node from parent. If scope not exits, return nil.
func (n *Node) RemoveNode(scope string) *Node {
	if child, ok := n.Children[scope]; ok {
		delete(n.Children, scope)
		return child
	}

	return nil
}

// Prune will delete all empty node in the tree.
func (n *Node) Prune() {
	oldMap := n.Children
	for scope, node := range oldMap {
		node.Prune()

		if len(node.Data) == 0 && len(node.Children) == 0 {
			n.RemoveNode(scope)
		}
	}
}

func (n *Node) SetChildrenParent() {
	for _, child := range n.Children {
		child.Parent = n
		child.SetChildrenParent()
	}
}

func (n *Node) ToJson() (string, error) {
	return utils.ToJson(n)
}

func (n *Node) String() string {
	if r, err := n.ToJson(); err == nil {
		return r
	}
	return ""
}
