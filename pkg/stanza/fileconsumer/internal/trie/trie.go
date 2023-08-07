// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// TRIE data structure inspired by https://github.com/dghubble/trie
/*
	This differs from the original trie.
   	This has been modified to detect partial matches as well.
   	For eg.
		If we add "ABCD" to this trie, and try to check if "ABCDEF" is in the trie,
		it will return true because that's how fingerprint matching works in current implementation.
*/

package trie // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/trie"

type Trie[T comparable] struct {
	children map[byte]*Trie[T]
	value    T
	parent   *Trie[T]
}

// NewTrie allocates and returns a new *Trie.
func NewTrie[T comparable]() *Trie[T] {
	return &Trie[T]{}
}

func (trie *Trie[T]) HasKey(key []byte) (bool, T) {
	node := trie
	isPresent := false
	var nullValue T
	matched := 0
	for _, r := range key {
		node = node.children[r]
		if node == nil {
			return isPresent, nullValue
		}
		matched += 1
		// check for any ending node in our current path
		isPresent = isPresent || !node.isNull()
		// We have reached end of the current path and all the previous characters have matched
		// break from here
		if node.isLeaf() && node != trie {
			break
		}
	}
	if matched < len(key) {
		// not a match for value as the key length exceeds the trie depth
		return isPresent, nullValue
	}
	return isPresent, node.value
}

// Put inserts the key into the trie
func (trie *Trie[T]) Put(key []byte, value T) {
	var nullValue T // This is used to assign empty value to nodes
	var last *Trie[T]
	node := trie
	shouldPush := false
	for _, r := range key {
		child, ok := node.children[r]
		if !ok {
			if node.children == nil {
				node.children = map[byte]*Trie[T]{}
			}
			child = NewTrie[T]()
			child.parent = node
			node.children[r] = child
		}
		node = child
		if !node.isNull() {
			last = node
			shouldPush = true
		}
	}
	node.value = value
	if shouldPush {
		node.value = last.value
		last.value = nullValue
	}
}

// Delete removes keys from the Trie. Returns true if node was found for the given key.
// If the node or any of its ancestors
// becomes childless as a result, it is removed from the trie.
func (trie *Trie[T]) Delete(key []byte) bool {
	var path []*Trie[T] // record ancestors to check later
	node := trie
	for _, b := range key {
		path = append(path, node)
		node = node.children[b]
		if node == nil {
			// node does not exist
			return false
		}
	}
	return trie.DeleteNode(node)
}

func (trie *Trie[T]) DeleteNode(node *Trie[T]) bool {
	var nullValue T // This is used to assign empty value to nodes
	if node.isNull() {
		// someonce called Delete() on the node which doesn't have any value
		// exit straight away
		return false
	}
	node.value = nullValue
	// if leaf, remove it from its parent's children map. Repeat for ancestor path.
	if node.isLeaf() {
		// iterate backwards over path
		for node != trie {
			parent := node.parent

			// Find the node in parent's children with matching key
			for key, child := range parent.children {
				if node == child {
					// delete the key from parent's map
					delete(parent.children, key)
					break
				}
			}
			if !parent.isLeaf() {
				// parent has other children, stop
				break
			}
			parent.children = nil
			if !parent.isNull() {
				// Parent has a value, stop
				break
			}
			node = node.parent
		}
	}
	return true // node (internal or not) existed and its value was nil'd
}

func (trie *Trie[T]) isNull() bool {
	var nullValue T
	return trie.value == nullValue
}

func (trie *Trie[T]) isLeaf() bool {
	return len(trie.children) == 0
}

type Value interface{}
