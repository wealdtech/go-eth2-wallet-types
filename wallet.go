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
)

// Wallet is the interface for wallets.
// A wallet contains one or more accounts.  Each account has its own security mechanism.
type Wallet interface {
	// ID provides the ID for the wallet.
	ID() uuid.UUID

	// Name provides the name for the wallet.
	Name() string

	// Type provides the type of the wallet.
	Type() string

	// Version provides the version of the wallet.
	Version() uint

	// Lock locks the wallet.  A locked account cannot create new accounts.
	Lock()

	// Unlock unlocks the wallet.  An unlocked account can create new accounts.
	Unlock([]byte) error

	// IsUnlocked returns true if the wallet is unlocked.
	IsUnlocked() bool

	// CreateAccount creates a new account in the wallet.
	// The only rule for names is that they cannot start with an underscore (_) character.
	// This will error if an account with the name already exists.
	CreateAccount(name string, passphrase []byte) (Account, error)

	// ImportAccount creates a new account in the wallet from an existing private key.
	// The only rule for names is that they cannot start with an underscore (_) character.
	// This will error if an account with the name already exists.
	ImportAccount(name string, key []byte, passphrase []byte) (Account, error)

	// Accounts provides all accounts in the wallet.
	Accounts() <-chan Account

	// AccountByName provides a single account from the wallet given its name.
	// This will error if the account is not found.
	AccountByName(name string) (Account, error)

	// Key returns the wallet's key.
	// This may or may not be present, depending on the wallet type.
	Key() ([]byte, error)

	// Export exports the entire wallet, protected by an additional passphrase.
	Export(passphrase []byte) ([]byte, error)

	// Import exports the entire wallet, protected by an additional passphrase.
	// Import(passphrase []byte) ([]byte, error)
}
