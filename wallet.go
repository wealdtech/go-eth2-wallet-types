// Copyright 2019, 2020 Weald Technology Trading
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

	// Accounts provides all accounts in the wallet.
	Accounts() <-chan Account

	// AccountByID provides a single account from the wallet given its ID.
	// This will error if the account is not found.
	AccountByID(id uuid.UUID) (Account, error)

	// AccountByName provides a single account from the wallet given its name.
	// This will error if the account is not found.
	AccountByName(name string) (Account, error)
}

// WalletAccountCreator is the interface for wallets that can create accounts.
type WalletAccountCreator interface {
	// CreateAccount creates a new account in the wallet.
	// The only rule for names is that they cannot start with an underscore (_) character.
	// This will error if an account with the name already exists.
	CreateAccount(name string, passphrase []byte) (Account, error)
}

// WalletDistributedAccountCreator is the interface for wallets that can create distributed accounts.
type WalletDistributedAccountCreator interface {
	// CreateDistributedAccount creates a new distributed account in the wallet.
	// The only rule for names is that they cannot start with an underscore (_) character.
	// This will error if an account with the name already exists.
	CreateDistributedAccount(name string, particpants uint32, signingThreshold uint32, passphrase []byte) (Account, error)
}

// WalletKeyProvider is the interface for wallets that can provide a key.
type WalletKeyProvider interface {
	// Key returns the wallet's key.
	Key() ([]byte, error)
}

// WalletExporter is the interface for wallets that can export themselves.
type WalletExporter interface {
	// Export exports the entire wallet, protected by an additional passphrase.
	Export(passphrase []byte) ([]byte, error)
}

// WalletAccountImporter is the interface for wallets that can import accounts.
type WalletAccountImporter interface {
	// ImportAccount creates a new account in the wallet from an existing private key.
	// The only rule for names is that they cannot start with an underscore (_) character.
	// This will error if an account with the name already exists.
	ImportAccount(name string, key []byte, passphrase []byte) (Account, error)
}

// WalletDistributedAccountImporter is the interface for wallets that can import distributed accounts.
type WalletDistributedAccountImporter interface {
	// ImportDistributedAccount creates a new distributed account in the wallet from provided data.
	// The only rule for names is that they cannot start with an underscore (_) character.
	// This will error if an account with the name already exists.
	ImportDistributedAccount(name string, privatekey []byte, threshold uint32, verificationVector [][]byte, participants map[uint64]string, passphrase []byte) (Account, error)
}
