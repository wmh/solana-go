// Copyright 2021 github.com/gagliardetto
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package computebudget

import (
	"testing"

	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"
)

func TestSetLoadedAccountsDataSizeLimit(t *testing.T) {
	accounts := []*ag_solanago.AccountMeta{}
	instruction := NewSetLoadedAccountsDataSizeLimitInstruction(1024 * 1024)

	inst, err := instruction.ValidateAndBuild()
	require.NoError(t, err)

	data, err := inst.Data()
	require.NoError(t, err)

	decoded, err := DecodeInstruction(accounts, data)
	require.NoError(t, err)

	require.Equal(t, Instruction_SetLoadedAccountsDataSizeLimit, decoded.TypeID.Uint8())
	require.Equal(t, uint32(1024*1024), decoded.Impl.(*SetLoadedAccountsDataSizeLimit).AccountDataSizeLimit)
}

func TestSetLoadedAccountsDataSizeLimitValidation(t *testing.T) {
	// Test with zero limit (should fail)
	instruction := NewSetLoadedAccountsDataSizeLimitInstruction(0)
	_, err := instruction.ValidateAndBuild()
	require.Error(t, err)
	require.Contains(t, err.Error(), "AccountDataSizeLimit parameter is not set")

	// Test with valid limit
	instruction = NewSetLoadedAccountsDataSizeLimitInstruction(1024)
	_, err = instruction.ValidateAndBuild()
	require.NoError(t, err)
}

func TestSetLoadedAccountsDataSizeLimitSerialization(t *testing.T) {
	original := SetLoadedAccountsDataSizeLimit{
		AccountDataSizeLimit: 2048,
	}

	data, err := ag_binary.MarshalBin(original)
	require.NoError(t, err)

	var decoded SetLoadedAccountsDataSizeLimit
	err = ag_binary.UnmarshalBin(data, &decoded)
	require.NoError(t, err)

	require.Equal(t, original.AccountDataSizeLimit, decoded.AccountDataSizeLimit)
}