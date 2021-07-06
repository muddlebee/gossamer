package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/dgraph-io/badger/v2"

	"github.com/stretchr/testify/require"
)

func iterateDB(db *badger.DB, cb func(*badger.Item)) {
	txn := db.NewTransaction(false)
	itr := txn.NewIterator(badger.DefaultIteratorOptions)

	for itr.Rewind(); itr.Valid(); itr.Next() {
		cb(itr.Item())
	}
}
func runPruneCmd(t *testing.T, configFile, prunedDBPath string) {
	ctx, err := newTestContext(
		"Test state trie offline pruning  --prune-state",
		[]string{"config", "pruned-db-path", "bloom-size", "retain-blocks"},
		[]interface{}{configFile, prunedDBPath, "256", int64(5)},
	)
	if err != nil {
		t.Fatal(err)
	}

	command := pruningCommand
	err = command.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPruneState(t *testing.T) {
	var (
		inputDBPath   = "../../tests/data/db"
		configFile    = "../../tests/data/db/config.toml"
		prunedDBPath  = fmt.Sprintf("%s/%s", t.TempDir(), "pruned")
		storagePrefix = "storage"
	)

	inputDB, err := badger.Open(badger.DefaultOptions(inputDBPath).WithReadOnly(true))
	require.NoError(t, err)

	nonStorageKeys := make(map[string]interface{})
	var numStorageKeys int

	getKeysInputDB := func(item *badger.Item) {
		key := string(item.Key())
		if strings.HasPrefix(key, storagePrefix) {
			numStorageKeys++
			return
		}
		nonStorageKeys[key] = nil
	}
	iterateDB(inputDB, getKeysInputDB)

	err = inputDB.Close()
	require.NoError(t, err)

	t.Log("Total keys in input DB", numStorageKeys+len(nonStorageKeys), "storage keys", numStorageKeys)

	t.Log("pruned DB path", prunedDBPath)

	runPruneCmd(t, configFile, prunedDBPath)

	prunedDB, err := badger.Open(badger.DefaultOptions(prunedDBPath))
	require.NoError(t, err)

	nonStorageKeysPruned := make(map[string]interface{})
	var numStorageKeysPruned int

	getKeysPrunedDB := func(item *badger.Item) {
		key := string(item.Key())
		if strings.HasPrefix(key, storagePrefix) {
			numStorageKeysPruned++
			return
		}
		nonStorageKeysPruned[key] = nil
	}
	iterateDB(prunedDB, getKeysPrunedDB)

	t.Log("Total keys in pruned DB", len(nonStorageKeysPruned)+numStorageKeysPruned, "storage keys", numStorageKeysPruned)
	require.Equal(t, len(nonStorageKeysPruned), len(nonStorageKeys))

	// Check all non storage keys are present.
	for k := range nonStorageKeys {
		_, ok := nonStorageKeysPruned[k]
		require.True(t, ok)
	}

	err = prunedDB.Close()
	require.NoError(t, err)
}