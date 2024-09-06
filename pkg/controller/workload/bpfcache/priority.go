/*
 * Copyright The Kmesh Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package bpfcache

import (
	"github.com/cilium/ebpf"
)

const (
	MaxPrio       = 7
	MaxSizeOfPrio = 1000
)

type PrioKey struct {
	ServiceId uint32
	Rank      uint32
}

type PrioValue struct {
	Count   uint32                // count of current prio
	UidList [MaxSizeOfPrio]uint32 // workload_uid to backend
}

func (c *Cache) PrioUpdate(key *PrioKey, value *PrioValue) error {
	log.Debugf("PrioUpdate [%#v], [%#v]", *key, *value)
	return c.bpfMap.KmeshPrio.Update(key, value, ebpf.UpdateAny)
}

func (c *Cache) PrioDelete(key *PrioKey) error {
	log.Debugf("PrioDelete [%#v]", *key)
	return c.bpfMap.KmeshPrio.Delete(key)
}

func (c *Cache) PrioLookup(key *PrioKey, value *PrioValue) error {
	log.Debugf("PrioLookup [%#v]", *key)
	return c.bpfMap.KmeshPrio.Lookup(key, value)
}
