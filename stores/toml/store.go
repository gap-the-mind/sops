package toml

import (
	"fmt"

	"github.com/pelletier/go-toml"
	"go.mozilla.org/sops/v3"
	"go.mozilla.org/sops/v3/stores"
)

// Store handles storage of TOML data
type Store struct {
}

// LoadEncryptedFile loads an encrypted secrets file onto a sops.Tree object
func (store *Store) LoadEncryptedFile(in []byte) (sops.Tree, error) {
	return sops.Tree{}, nil
}

func tomlTreeValueToTreeValue(tree interface{}) interface{} {

	switch node := tree.(type) {
	case []*toml.Tree:
		return []interface{}{}
	case *toml.Tree:
		return tomlTreeToTreeBranch(node)

	default:
		return node
	}
}

func tomlTreeToTreeBranch(tree *toml.Tree) sops.TreeBranch {
	var root sops.TreeBranch

	fmt.Printf("%#v", tree)

	for _, key := range tree.Keys() {
		item := sops.TreeItem{
			Key:   key,
			Value: tomlTreeValueToTreeValue(tree.Get(key)),
		}

		root = append(root, item)
	}
	return root
}

// LoadPlainFile loads plaintext toml file bytes onto a sops.TreeBranches object
func (store *Store) LoadPlainFile(in []byte) (sops.TreeBranches, error) {
	tree, err := toml.LoadBytes(in)

	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling input TOML: %s", err)
	}

	branches := sops.TreeBranches{tomlTreeToTreeBranch(tree)}

	fmt.Printf("%#v\n", branches)

	return branches, nil
}

// EmitEncryptedFile returns the encrypted bytes of the toml file corresponding to a
// sops.Tree runtime object
func (store *Store) EmitEncryptedFile(in sops.Tree) ([]byte, error) {
	return nil, nil
}

// EmitPlainFile returns the plaintext bytes of the toml file corresponding to a
// sops.TreeBranches runtime object
func (store *Store) EmitPlainFile(in sops.TreeBranches) ([]byte, error) {
	return nil, nil
}

// EmitValue returns bytes corresponding to a single encoded value
// in a generic interface{} object
func (store *Store) EmitValue(v interface{}) ([]byte, error) {
	return nil, nil
}

// EmitExample returns the bytes corresponding to an example complex tree
func (store *Store) EmitExample() []byte {
	bytes, err := store.EmitPlainFile(stores.ExampleComplexTree.Branches)
	if err != nil {
		panic(err)
	}
	return bytes
}
