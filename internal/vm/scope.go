package vm

import (
	"arena"
	"fmt"

	obj "git.sr.ht/~madcapjake/rhi/internal/object"
)

const INIT_BUCKET_SIZE int = 16

const LOAD_EXPAND_THRESHOLD float32 = 0.5
const LOAD_SHRINK_THRESHOLD float32 = 0.35
const FNV_OFFSET_BASIS uint64 = 0xCBF29CE484222325
const FNV_PRIME uint64 = 0x00000100000001B3

func hashValue(value obj.Hashable, limit int) int {
	hash := FNV_OFFSET_BASIS
	for _, b := range value.HashBytes() {
		hash = hash ^ uint64(b)
		hash = hash * FNV_PRIME
	}
	return int(hash % uint64(limit))
}

type Scope struct {
	length  int
	buckets [][]ScopeEntry
}

type ScopeEntry struct {
	key   obj.Hashable
	value obj.Any
}

func NewScopeSized(a *arena.Arena, i int) Scope {
	return Scope{buckets: arena.MakeSlice[[]ScopeEntry](a, 20, 20)}
}

func NewScope(a *arena.Arena) Scope {
	return NewScopeSized(a, INIT_BUCKET_SIZE)
}

func (sc Scope) Get(key obj.Hashable) (obj.Any, bool) {
	hash := hashValue(key, len(sc.buckets))
	for _, v := range sc.buckets[hash] {
		if v.key == key {
			return v.value, true
		}
	}
	return nil, false
}

func (sc Scope) loadFactor() float32 {
	return float32(sc.length) / float32(len(sc.buckets))
}

func (sc *Scope) Set(key obj.Hashable, value obj.Any) {
	hash := hashValue(key, len(sc.buckets))
	// check if key is already added, if yes, overwrite it
	for i, lbl := range sc.buckets[hash] {
		if lbl.key == key {
			sc.buckets[hash][i].value = value
			return
		}
	}
	sc.buckets[hash] = append(sc.buckets[hash], ScopeEntry{key, value})
	sc.length += 1
	if sc.loadFactor() > LOAD_EXPAND_THRESHOLD {
		sc.expand()
	}
}

func (sc *Scope) adjustBuckets(amount int) [][]ScopeEntry {
	buckets := make([][]ScopeEntry, amount)
	for _, bucket := range sc.buckets {
		for _, lbl := range bucket {
			newHash := hashValue(lbl.key, len(buckets))
			buckets[newHash] = append(buckets[newHash],
				ScopeEntry{lbl.key, lbl.value})
		}
	}
	return buckets
}

func (sc *Scope) expand() error {
	sc.buckets = sc.adjustBuckets(len(sc.buckets) * 2)
	return nil
}

func (sc *Scope) shrink() error {
	sc.buckets = sc.adjustBuckets(len(sc.buckets) / 2)
	return nil
}

func (sc *Scope) Delete(key obj.Hashable) error {
	hash := hashValue(key, len(sc.buckets))
	for i, lbl := range sc.buckets[hash] {
		if lbl.key == key {
			current := sc.buckets[hash]
			current[i] = current[len(current)-1]
			current = current[:len(current)-1]
			sc.length -= 1
			sc.buckets[hash] = current
			if sc.loadFactor() < LOAD_SHRINK_THRESHOLD {
				sc.shrink()
			}
			return nil
		}
	}
	return fmt.Errorf("key error")
}
