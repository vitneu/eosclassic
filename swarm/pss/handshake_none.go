// Copyright 2018 The eosclassic & go-ethereum Authors
// This file is part of the eosclassic library.
//
// The eosclassic library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The eosclassic library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the eosclassic library. If not, see <http://www.gnu.org/licenses/>.

// +build nopsshandshake

package pss

const (
	IsActiveHandshake = false
)

func NewHandshakeParams() interface{} {
	return nil
}
