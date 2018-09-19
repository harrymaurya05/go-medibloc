// Copyright (C) 2018  MediBloc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package trie_test

import (
	"encoding/hex"
	"testing"

	"github.com/medibloc/go-medibloc/storage"
	"github.com/stretchr/testify/require"

	"github.com/medibloc/go-medibloc/common/trie"
	"github.com/medibloc/go-medibloc/core"
	"github.com/medibloc/go-medibloc/util/testutil"
	"github.com/stretchr/testify/assert"
)

func testBatch(t *testing.T, batch *trie.Batch) {
	err := batch.BeginBatch()
	assert.Nil(t, err)

	err = batch.BeginBatch()
	assert.Equal(t, core.ErrBeginAgainInBatch, err)

	err = batch.RollBack()
	assert.Nil(t, err)

	err = batch.RollBack()
	assert.Equal(t, core.ErrNotBatching, err)
}

func TestTrieBatch(t *testing.T) {
	s := testutil.GetStorage(t)

	trieBatch, err := trie.NewBatch(nil, s)
	assert.Nil(t, err)

	testBatch(t, trieBatch)

	err = trieBatch.BeginBatch()
	assert.Nil(t, err)

	err = trieBatch.BeginBatch()
	assert.Equal(t, core.ErrBeginAgainInBatch, err)

	key1, _ := hex.DecodeString("2ed1b10c")
	val1 := []byte("leafmedi1")
	val2 := []byte("leafmedi2")

	testGetValue := func(key []byte, val []byte) {
		getVal, err := trieBatch.Get(key1)
		if val == nil {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, val, getVal)
		}
	}

	err = trieBatch.Put(key1, val1)
	assert.Nil(t, err)
	testGetValue(key1, val1)

	err = trieBatch.Commit()
	assert.Nil(t, err)

	testGetValue(key1, val1)

	err = trieBatch.BeginBatch()
	assert.Nil(t, err)

	err = trieBatch.Put(key1, val2)
	assert.Nil(t, err)
	testGetValue(key1, val2)
	trieBatch.RollBack()
	testGetValue(key1, val1)

	err = trieBatch.BeginBatch()
	assert.Nil(t, err)

	err = trieBatch.Delete(key1)
	assert.Nil(t, err)
	testGetValue(key1, nil)
	trieBatch.RollBack()
	testGetValue(key1, val1)
}

func TestTrieBatch_Clone(t *testing.T) {
	s := testutil.GetStorage(t)

	trBatch1, err := trie.NewBatch(nil, s)
	assert.Nil(t, err)

	trBatch2, err := trBatch1.Clone()

	err = trBatch1.BeginBatch()
	assert.Nil(t, err)

	_, err = trBatch1.Clone()
	assert.Equal(t, core.ErrCannotCloneOnBatching, err)

	key1, _ := hex.DecodeString("2ed1b10c")
	val1 := []byte("leafmedi1")

	err = trBatch1.Put(key1, val1)
	assert.Nil(t, err)

	err = trBatch1.Commit()
	assert.Nil(t, err)

	_, err = trBatch1.Get(key1)
	assert.Nil(t, err)
	_, err = trBatch2.Get(key1)
	assert.Equal(t, core.ErrNotFound, err)

	err = trBatch2.BeginBatch()
	assert.Nil(t, err)

	err = trBatch2.Put(key1, val1)
	assert.Nil(t, err)

	err = trBatch2.Commit()
	assert.Nil(t, err)
	getVal1, err := trBatch1.Get(key1)
	assert.Nil(t, err)
	getVal2, err := trBatch2.Get(key1)
	assert.Nil(t, err)
	assert.Equal(t, getVal1, getVal2)
}

func TestTrieBatch_DiskWriteCount(t *testing.T) {
	testPairs := []*testPair{
		{[]byte{0, 0, 0, 1}, "node0"},
		{[]byte{0, 0, 0, 1, 1, 2}, "node1"},
		{[]byte{0, 0, 0, 1, 1, 3}, "node2"},
		{[]byte{0, 0, 0, 1, 1, 5, 5, 5, 1, 0}, "node3"},
		{[]byte{0, 0, 0, 1, 1, 5, 5, 5, 2, 0}, "node4"},
	}

	storTr, err := storage.NewMemoryStorage()
	require.NoError(t, err)
	tr, err := trie.NewTrie(nil, storTr)
	require.NoError(t, err)

	for _, pair := range testPairs {
		putPairToTrie(t, tr, pair)
	}

	for _, pair := range testPairs {
		getAndEqual(t, tr, pair)
	}

	cnt := 0
	storTr.Data().Range(func(key, value interface{}) bool {
		cnt++
		return true
	})
	t.Log("Number of nodes in trie storage:", cnt)

	storTb, err := storage.NewMemoryStorage()
	require.NoError(t, err)
	tb, err := trie.NewBatch(nil, storTb)
	require.NoError(t, err)

	require.NoError(t, tb.BeginBatch())
	for _, pair := range testPairs {
		putPairToTrieBatch(t, tb, pair)
	}
	require.NoError(t, tb.Commit())

	for _, pair := range testPairs {
		getAndEqualBatch(t, tb, pair)
	}

	cnt = 0
	storTb.Data().Range(func(key, value interface{}) bool {
		cnt++
		return true
	})
	t.Log("Number of nodes in trie batch storage:", cnt)

}

func TestTrieBatch_DiskWriteCount2(t *testing.T) {
	testPairs := []*testPair{
		{[]byte{0, 0, 0, 1}, "node0"},
		{[]byte{0, 0, 0, 1, 1, 2}, "node1"},
		{[]byte{0, 0, 0, 1, 1, 3}, "node2"},
		{[]byte{0, 0, 0, 1, 1, 5, 5, 5, 1, 0}, "node3"},
		{[]byte{0, 0, 0, 1, 1, 5, 5, 5, 2, 0}, "node4"},
	}

	storTr, err := storage.NewMemoryStorage()
	require.NoError(t, err)
	tr, err := trie.NewTrie(nil, storTr)
	require.NoError(t, err)

	for _, pair := range testPairs {
		putPairToTrie(t, tr, pair)
	}
	require.NoError(t, tr.Delete(trie.RouteToKey(testPairs[4].route)))

	for _, pair := range testPairs[:4] {
		getAndEqual(t, tr, pair)
	}

	cnt := 0
	storTr.Data().Range(func(key, value interface{}) bool {
		cnt++
		return true
	})
	t.Log("Number of nodes in trie storage:", cnt)

	storTb, err := storage.NewMemoryStorage()
	require.NoError(t, err)
	tb, err := trie.NewBatch(nil, storTb)
	require.NoError(t, err)

	require.NoError(t, tb.BeginBatch())
	for _, pair := range testPairs {
		putPairToTrieBatch(t, tb, pair)
	}
	require.NoError(t, tb.Delete(trie.RouteToKey(testPairs[4].route)))
	require.NoError(t, tb.Commit())

	for _, pair := range testPairs[:4] {
		getAndEqualBatch(t, tb, pair)
	}

	cnt = 0
	storTb.Data().Range(func(key, value interface{}) bool {
		cnt++
		return true
	})
	t.Log("Number of nodes in trie batch storage:", cnt)

}
