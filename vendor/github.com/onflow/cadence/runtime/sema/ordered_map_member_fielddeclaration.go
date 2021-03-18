// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Based on https://github.com/wk8/go-ordered-map, Copyright Jean Rougé
 *
 */

package sema

import (
	"container/list"

	"github.com/onflow/cadence/runtime/ast"
)

// MemberAstFieldDeclarationOrderedMap
//
type MemberAstFieldDeclarationOrderedMap struct {
	pairs map[*Member]*MemberAstFieldDeclarationPair
	list  *list.List
}

// NewMemberAstFieldDeclarationOrderedMap creates a new MemberAstFieldDeclarationOrderedMap.
func NewMemberAstFieldDeclarationOrderedMap() *MemberAstFieldDeclarationOrderedMap {
	return &MemberAstFieldDeclarationOrderedMap{
		pairs: make(map[*Member]*MemberAstFieldDeclarationPair),
		list:  list.New(),
	}
}

// Get returns the value associated with the given key.
// Returns nil if not found.
// The second return value indicates if the key is present in the map.
func (om *MemberAstFieldDeclarationOrderedMap) Get(key *Member) (result *ast.FieldDeclaration, present bool) {
	var pair *MemberAstFieldDeclarationPair
	if pair, present = om.pairs[key]; present {
		return pair.Value, present
	}
	return
}

// GetPair returns the key-value pair associated with the given key.
// Returns nil if not found.
func (om *MemberAstFieldDeclarationOrderedMap) GetPair(key *Member) *MemberAstFieldDeclarationPair {
	return om.pairs[key]
}

// Set sets the key-value pair, and returns what `Get` would have returned
// on that key prior to the call to `Set`.
func (om *MemberAstFieldDeclarationOrderedMap) Set(key *Member, value *ast.FieldDeclaration) (oldValue *ast.FieldDeclaration, present bool) {
	var pair *MemberAstFieldDeclarationPair
	if pair, present = om.pairs[key]; present {
		oldValue = pair.Value
		pair.Value = value
		return
	}

	pair = &MemberAstFieldDeclarationPair{
		Key:   key,
		Value: value,
	}
	pair.element = om.list.PushBack(pair)
	om.pairs[key] = pair

	return
}

// Delete removes the key-value pair, and returns what `Get` would have returned
// on that key prior to the call to `Delete`.
func (om *MemberAstFieldDeclarationOrderedMap) Delete(key *Member) (oldValue *ast.FieldDeclaration, present bool) {
	var pair *MemberAstFieldDeclarationPair
	pair, present = om.pairs[key]
	if !present {
		return
	}

	om.list.Remove(pair.element)
	delete(om.pairs, key)
	oldValue = pair.Value

	return
}

// Len returns the length of the ordered map.
func (om *MemberAstFieldDeclarationOrderedMap) Len() int {
	return len(om.pairs)
}

// Oldest returns a pointer to the oldest pair.
func (om *MemberAstFieldDeclarationOrderedMap) Oldest() *MemberAstFieldDeclarationPair {
	return listElementToMemberAstFieldDeclarationPair(om.list.Front())
}

// Newest returns a pointer to the newest pair.
func (om *MemberAstFieldDeclarationOrderedMap) Newest() *MemberAstFieldDeclarationPair {
	return listElementToMemberAstFieldDeclarationPair(om.list.Back())
}

// Foreach iterates over the entries of the map in the insertion order, and invokes
// the provided function for each key-value pair.
func (om *MemberAstFieldDeclarationOrderedMap) Foreach(f func(key *Member, value *ast.FieldDeclaration)) {
	for pair := om.Oldest(); pair != nil; pair = pair.Next() {
		f(pair.Key, pair.Value)
	}
}

// MemberAstFieldDeclarationPair
//
type MemberAstFieldDeclarationPair struct {
	Key   *Member
	Value *ast.FieldDeclaration

	element *list.Element
}

// Next returns a pointer to the next pair.
func (p *MemberAstFieldDeclarationPair) Next() *MemberAstFieldDeclarationPair {
	return listElementToMemberAstFieldDeclarationPair(p.element.Next())
}

// Prev returns a pointer to the previous pair.
func (p *MemberAstFieldDeclarationPair) Prev() *MemberAstFieldDeclarationPair {
	return listElementToMemberAstFieldDeclarationPair(p.element.Prev())
}

func listElementToMemberAstFieldDeclarationPair(element *list.Element) *MemberAstFieldDeclarationPair {
	if element == nil {
		return nil
	}
	return element.Value.(*MemberAstFieldDeclarationPair)
}
