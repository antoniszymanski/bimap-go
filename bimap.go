/*
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package bimap

import (
	"iter"
	"maps"
)

type BiMap[K, V comparable] struct {
	forward map[K]V
	inverse map[V]K
}

func New[K, V comparable](size int) BiMap[K, V] {
	return BiMap[K, V]{
		forward: make(map[K]V, max(0, size)),
		inverse: make(map[V]K, max(0, size)),
	}
}

func From[K, V comparable](items map[K]V) BiMap[K, V] {
	bm := New[K, V](len(items))
	for key, val := range items {
		bm.Set(key, val)
	}
	return bm
}

func Collect[K, V comparable](seq iter.Seq2[K, V]) BiMap[K, V] {
	bm := New[K, V](0)
	bm.Insert(seq)
	return bm
}

func (bm BiMap[K, V]) Get(key K) V {
	return bm.forward[key]
}

func (bm BiMap[K, V]) Lookup(key K) (V, bool) {
	v, ok := bm.forward[key]
	return v, ok
}

func (bm BiMap[K, V]) Has(key K) bool {
	_, ok := bm.forward[key]
	return ok
}

func (bm BiMap[K, V]) GetInverse(val V) K {
	return bm.inverse[val]
}

func (bm BiMap[K, V]) LookupInverse(val V) (K, bool) {
	v, ok := bm.inverse[val]
	return v, ok
}

func (bm BiMap[K, V]) HasInverse(val V) bool {
	_, ok := bm.inverse[val]
	return ok
}

func (bm BiMap[K, V]) Set(key K, val V) {
	bm.forward[key] = val
	bm.inverse[val] = key
}

func (bm BiMap[K, V]) Delete(key K) {
	delete(bm.inverse, bm.forward[key])
	delete(bm.forward, key)
}

func (bm BiMap[K, V]) DeleteInverse(val V) {
	delete(bm.forward, bm.inverse[val])
	delete(bm.inverse, val)
}

func (bm BiMap[K, V]) DeleteFunc(del func(K, V) bool) {
	for k, v := range bm.forward {
		if del(k, v) {
			bm.Delete(k)
		}
	}
}

func (bm BiMap[K, V]) Clear() {
	clear(bm.forward)
	clear(bm.inverse)
}

func (bm BiMap[K, V]) Equal(other BiMap[K, V]) bool {
	return maps.Equal(bm.forward, other.forward)
}

func (bm BiMap[K, V]) EqualFunc(other BiMap[K, V], eq func(V, V) bool) bool {
	return maps.EqualFunc(bm.forward, other.forward, eq)
}

func (bm BiMap[K, V]) Clone() BiMap[K, V] {
	return BiMap[K, V]{
		forward: maps.Clone(bm.forward),
		inverse: maps.Clone(bm.inverse),
	}
}

func (bm BiMap[K, V]) Reverse() BiMap[V, K] {
	return BiMap[V, K]{
		forward: maps.Clone(bm.inverse),
		inverse: maps.Clone(bm.forward),
	}
}

func (bm BiMap[K, V]) Copy(dst BiMap[K, V]) {
	for key, val := range bm.forward {
		dst.Set(key, val)
	}
}

func (bm BiMap[K, V]) Size() int {
	return len(bm.forward)
}

func (bm BiMap[K, V]) All() iter.Seq2[K, V] {
	return maps.All(bm.forward)
}

func (bm BiMap[K, V]) Keys() iter.Seq[K] {
	return maps.Keys(bm.forward)
}

func (bm BiMap[K, V]) Values() iter.Seq[V] {
	return maps.Values(bm.forward)
}

func (bm BiMap[K, V]) Insert(seq iter.Seq2[K, V]) {
	for key, val := range seq {
		bm.Set(key, val)
	}
}
