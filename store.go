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

import "github.com/google/uuid"

// Store is the interface for wallet stores.  It is used to store and access data provided by wallets, both wallets themselves
// as well as keys inside the wallets.
type Store interface {
	// Name provides the name of the store
	Name() string

	// StoreWallet stores wallet data.  It will fail if it cannot store the data.
	StoreWallet(walletID uuid.UUID, walletName string, data []byte) error

	// RetrieveWallet retrieves wallet data for all wallets.
	RetrieveWallets() <-chan []byte

	// RetrieveWallet retrieves wallet data for a wallet with a given name.
	// It will fail if it cannot retrieve the data.
	RetrieveWallet(walletName string) ([]byte, error)

	// RetrieveWalletByID retrieves wallet data for a wallet with a given ID.
	// It will fail if it cannot retrieve the data.
	RetrieveWalletByID(walletID uuid.UUID) ([]byte, error)

	// StoreAccount stores account data.  It will fail if it cannot store the data.
	StoreAccount(walletID uuid.UUID, accountID uuid.UUID, accountName string, data []byte) error

	// RetrieveAccounts retrieves account information for all accounts.
	RetrieveAccounts(walletID uuid.UUID) <-chan []byte

	// RetrieveAccount retrieves account data for a wallet with a given name.
	// It will fail if it cannot retrieve the data.
	RetrieveAccount(walletID uuid.UUID, accountName string) ([]byte, error)

	// RetrieveAccountByID retrieves account data for a wallet with a given ID.
	// It will fail if it cannot retrieve the data.
	RetrieveAccountByID(walletID uuid.UUID, accountID uuid.UUID) ([]byte, error)
}
