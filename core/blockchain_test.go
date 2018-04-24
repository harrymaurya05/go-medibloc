// Copyright 2018 The go-medibloc Authors
// This file is part of the go-medibloc library.
//
// The go-medibloc library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-medibloc library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-medibloc library. If not, see <http://www.gnu.org/licenses/>.

package core_test

import (
	"testing"

	"github.com/medibloc/go-medibloc/core"
	testUtil "github.com/medibloc/go-medibloc/util/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getBlockChain(t *testing.T) *core.BlockChain {
	m := NewMockMedlet(t)

	bc, err := core.NewBlockChain(m.Config())
	require.NoError(t, err)
	err = bc.Setup(m.Genesis(), m.Storage())
	require.NoError(t, err)
	return bc
}

func getBlockSlice(blockMap map[testUtil.BlockID]*core.Block, id testUtil.BlockID) []*core.Block {
	return []*core.Block{blockMap[id]}
}

func TestBlockChain_OnePath(t *testing.T) {
	bc := getBlockChain(t)

	idxToParent := []testUtil.BlockID{testUtil.GenesisID, 0, 1, 2, 3, 4, 5}
	blockMap := testUtil.NewBlockTestSet(t, idxToParent)

	err := bc.PutVerifiedNewBlocks(blockMap[0], getBlockSlice(blockMap, 0), getBlockSlice(blockMap, 0))
	assert.NotNil(t, err)

	err = bc.PutVerifiedNewBlocks(testUtil.GenesisBlock, getBlockSlice(blockMap, 0), getBlockSlice(blockMap, 0))
	assert.Nil(t, err)

	for _, idx := range idxToParent[2:] {
		blocks := getBlockSlice(blockMap, idx)
		err = bc.PutVerifiedNewBlocks(blockMap[testUtil.BlockID(idx-1)], blocks, blocks)
		assert.Nil(t, err)
		err = bc.SetTailBlock(blocks[0])
		assert.Nil(t, err)
		assert.Equal(t, blocks[0].Hash(), bc.MainTailBlock().Hash())
	}
}

func TestBlockChain_Tree(t *testing.T) {
	tests := []struct {
		tree []testUtil.BlockID
	}{
		{[]testUtil.BlockID{testUtil.GenesisID, 0, 0, 1, 1, 1, 1}},
		{[]testUtil.BlockID{testUtil.GenesisID, 0, 0, 0, 0, 1, 2, 2, 3}},
		{[]testUtil.BlockID{testUtil.GenesisID, 0, 1, 1, 2, 2, 3, 3}},
	}
	// Put one-by-one
	for _, test := range tests {
		bc := getBlockChain(t)
		blockMap := testUtil.NewBlockTestSet(t, test.tree)
		for idx, parentID := range test.tree {
			blocks := getBlockSlice(blockMap, testUtil.BlockID(idx))
			err := bc.PutVerifiedNewBlocks(blockMap[parentID], blocks, blocks)
			assert.Nil(t, err)
		}
	}
	// Put all
	for _, test := range tests {
		bc := getBlockChain(t)
		blockMap := testUtil.NewBlockTestSet(t, test.tree)
		notTail := make(map[testUtil.BlockID]bool)
		for _, idx := range test.tree {
			notTail[idx] = true
		}
		allBlocks := make([]*core.Block, 0)
		tailBlocks := make([]*core.Block, 0)
		for id, block := range blockMap {
			allBlocks = append(allBlocks, block)
			if !notTail[id] {
				tailBlocks = append(tailBlocks, block)
			}
		}
		/* TODO handle when tailBlocks are wrong
		err := bc.PutVerifiedNewBlocks(genesisBlock, allBlocks, allBlocks)
		assert.NotNil(t, err)
		*/
		err := bc.PutVerifiedNewBlocks(testUtil.GenesisBlock, allBlocks, tailBlocks)
		assert.Nil(t, err)
	}
}
