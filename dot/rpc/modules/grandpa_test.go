// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package modules

import (
	"errors"
	"net/http"
	"testing"

	"github.com/ChainSafe/gossamer/dot/rpc/modules/mocks"
	"github.com/ChainSafe/gossamer/dot/types"
	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/ChainSafe/gossamer/lib/crypto/ed25519"
	"github.com/ChainSafe/gossamer/lib/grandpa"
	"github.com/ChainSafe/gossamer/lib/keystore"

	"github.com/stretchr/testify/assert"
)

func TestGrandpaModule_ProveFinality(t *testing.T) {
	testHash := common.NewHash([]byte{0x01, 0x02})
	testHashSlice := []common.Hash{testHash, testHash, testHash}

	mockBlockFinalityAPI := mocks.NewBlockFinalityAPI(t)
	mockBlockAPI := mocks.NewBlockAPI(t)
	mockBlockAPI.On("SubChain", testHash, testHash).Return(testHashSlice, nil)
	mockBlockAPI.On("HasJustification", testHash).Return(true, nil)
	mockBlockAPI.On("GetJustification", testHash).Return([]byte("test"), nil)

	mockBlockAPIHasJustErr := mocks.NewBlockAPI(t)
	mockBlockAPIHasJustErr.On("SubChain", testHash, testHash).Return(testHashSlice, nil)
	mockBlockAPIHasJustErr.On("HasJustification", testHash).Return(false, nil)

	mockBlockAPIGetJustErr := mocks.NewBlockAPI(t)
	mockBlockAPIGetJustErr.On("SubChain", testHash, testHash).Return(testHashSlice, nil)
	mockBlockAPIGetJustErr.On("HasJustification", testHash).Return(true, nil)
	mockBlockAPIGetJustErr.On("GetJustification", testHash).Return(nil, errors.New("GetJustification error"))

	mockBlockAPISubChainErr := mocks.NewBlockAPI(t)
	mockBlockAPISubChainErr.On("SubChain", testHash, testHash).Return(nil, errors.New("SubChain error"))

	grandpaModule := NewGrandpaModule(mockBlockAPISubChainErr, mockBlockFinalityAPI)
	type fields struct {
		blockAPI         BlockAPI
		blockFinalityAPI BlockFinalityAPI
	}
	type args struct {
		r   *http.Request
		req *ProveFinalityRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expErr error
		exp    ProveFinalityResponse
	}{
		{
			name: "SubChain Err",
			fields: fields{
				grandpaModule.blockAPI,
				grandpaModule.blockFinalityAPI,
			},
			args: args{
				req: &ProveFinalityRequest{
					blockHashStart: testHash,
					blockHashEnd:   testHash,
					authorityID:    uint64(21),
				},
			},
			expErr: errors.New("SubChain error"),
		},
		{
			name: "OK Case",
			fields: fields{
				mockBlockAPI,
				mockBlockFinalityAPI,
			},
			args: args{
				req: &ProveFinalityRequest{
					blockHashStart: testHash,
					blockHashEnd:   testHash,
					authorityID:    uint64(21),
				},
			},
			exp: ProveFinalityResponse{
				[]uint8{0x74, 0x65, 0x73, 0x74},
				[]uint8{0x74, 0x65, 0x73, 0x74},
				[]uint8{0x74, 0x65, 0x73, 0x74}},
		},
		{
			name: "HasJustification Error",
			fields: fields{
				mockBlockAPIHasJustErr,
				mockBlockFinalityAPI,
			},
			args: args{
				req: &ProveFinalityRequest{
					blockHashStart: testHash,
					blockHashEnd:   testHash,
					authorityID:    uint64(21),
				},
			},
			exp: ProveFinalityResponse(nil),
		},
		{
			name: "GetJustification Error",
			fields: fields{
				mockBlockAPIGetJustErr,
				mockBlockFinalityAPI,
			},
			args: args{
				req: &ProveFinalityRequest{
					blockHashStart: testHash,
					blockHashEnd:   testHash,
					authorityID:    uint64(21),
				},
			},
			exp: ProveFinalityResponse(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm := &GrandpaModule{
				blockAPI:         tt.fields.blockAPI,
				blockFinalityAPI: tt.fields.blockFinalityAPI,
			}
			res := ProveFinalityResponse(nil)
			err := gm.ProveFinality(tt.args.r, tt.args.req, &res)
			if tt.expErr != nil {
				assert.EqualError(t, err, tt.expErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.exp, res)
		})
	}
}

func TestGrandpaModule_RoundState(t *testing.T) {
	var kr, _ = keystore.NewEd25519Keyring()
	var voters grandpa.Voters

	for _, k := range kr.Keys {
		voters = append(voters, types.GrandpaVoter{
			Key: *k.Public().(*ed25519.PublicKey),
			ID:  1,
		})
	}

	mockBlockAPI := mocks.NewBlockAPI(t)
	mockBlockFinalityAPI := mocks.NewBlockFinalityAPI(t)
	mockBlockFinalityAPI.On("GetVoters").Return(voters)
	mockBlockFinalityAPI.On("GetSetID").Return(uint64(0))
	mockBlockFinalityAPI.On("GetRound").Return(uint64(2))
	mockBlockFinalityAPI.On("PreVotes").Return([]ed25519.PublicKeyBytes{
		kr.Alice().Public().(*ed25519.PublicKey).AsBytes(),
		kr.Bob().Public().(*ed25519.PublicKey).AsBytes(),
		kr.Charlie().Public().(*ed25519.PublicKey).AsBytes(),
		kr.Dave().Public().(*ed25519.PublicKey).AsBytes(),
	})
	mockBlockFinalityAPI.On("PreCommits").Return([]ed25519.PublicKeyBytes{
		kr.Alice().Public().(*ed25519.PublicKey).AsBytes(),
		kr.Bob().Public().(*ed25519.PublicKey).AsBytes(),
	})

	type fields struct {
		blockAPI         BlockAPI
		blockFinalityAPI BlockFinalityAPI
	}
	type args struct {
		r   *http.Request
		req *EmptyRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expErr error
		exp    RoundStateResponse
	}{
		{
			name: "GetJustification Error",
			fields: fields{
				mockBlockAPI,
				mockBlockFinalityAPI,
			},
			args: args{
				req: &EmptyRequest{},
			},
			exp: RoundStateResponse{
				SetID: 0x0,
				Best: RoundState{
					Round:           0x2,
					TotalWeight:     0x9,
					ThresholdWeight: 0x6,
					Prevotes: Votes{
						CurrentWeight: 0x4,
						Missing: []string{
							"5Ck2miBfCe1JQ4cY3NDsXyBaD6EcsgiVmEFTWwqNSs25XDEq",
							"5E2BmpVFzYGd386XRCZ76cDePMB3sfbZp5ZKGUsrG1m6gomN",
							"5CGR8FbjxeV31JKaUUuVUgasW79k8xFGdoh8WG5MokEc78qj",
							"5E9ZP1w5qat63KrWEJLkh7aDr2fPTbu3UhetAjxeyBojKHYH",
							"5Cjb197EXcHehjxuyKUCF3wJm86owKiuKCzF18DcMhbgMhPX",
						},
					},
					Precommits: Votes{
						CurrentWeight: 0x2,
						Missing: []string{
							"5DbKjhNLpqX3zqZdNBc9BGb4fHU1cRBaDhJUskrvkwfraDi6",
							"5ECTwv6cZ5nJQPk6tWfaTrEk8YH2L7X1VT4EL5Tx2ikfFwb7",
							"5Ck2miBfCe1JQ4cY3NDsXyBaD6EcsgiVmEFTWwqNSs25XDEq",
							"5E2BmpVFzYGd386XRCZ76cDePMB3sfbZp5ZKGUsrG1m6gomN",
							"5CGR8FbjxeV31JKaUUuVUgasW79k8xFGdoh8WG5MokEc78qj",
							"5E9ZP1w5qat63KrWEJLkh7aDr2fPTbu3UhetAjxeyBojKHYH",
							"5Cjb197EXcHehjxuyKUCF3wJm86owKiuKCzF18DcMhbgMhPX",
						},
					},
				},
				Background: []RoundState{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm := &GrandpaModule{
				blockAPI:         tt.fields.blockAPI,
				blockFinalityAPI: tt.fields.blockFinalityAPI,
			}
			res := RoundStateResponse{}
			err := gm.RoundState(tt.args.r, tt.args.req, &res)
			if tt.expErr != nil {
				assert.EqualError(t, err, tt.expErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.exp, res)
		})
	}
}
