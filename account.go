// Copyright © 2019 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"github.com/google/uuid"
	types "github.com/wealdtech/go-eth2-types"
)

// Account is the interface for all Ethereum 2 accounts.
type Account interface {
	// ID provides the ID for the account.
	ID() uuid.UUID

	// Name provides the name for the account.
	Name() string

	// PublicKey provides the public key for the account.
	PublicKey() types.PublicKey

	// Path provides the path for the account.
	// Can be empty if the account is not derived from a path.
	Path() string

	// Lock locks the account.  A locked account cannot sign.
	Lock()

	// Unlock unlocks the account.  An unlocked account can sign.
	Unlock([]byte) error

	// Sign signs data with the account.
	Sign(data []byte, domain uint64) (types.Signature, error)
}

// AccountPrivateKeyProvider is the interface for accounts that can provide a private key.
type AccountPrivateKeyProvider interface {
	// PrivateKey provides the private key for the account.
	PrivateKey() types.PrivateKey
}
