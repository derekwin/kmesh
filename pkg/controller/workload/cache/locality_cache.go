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

package cache

import (
	"sync"
)

type LocalityCache interface {
	GetOrCreateRegion(region string) uint32
	GetOrCreateZone(zone string) uint32
	GetOrCreateSubzone(subzone string) uint32
}

type localityCache struct {
	mutex sync.RWMutex
	// keyed by "locality"->region_id
	Region  map[string]uint32
	Zone    map[string]uint32
	Subzone map[string]uint32
}

func NewLocalityCache() *localityCache {
	return &localityCache{
		Region:  make(map[string]uint32),
		Zone:    make(map[string]uint32),
		Subzone: make(map[string]uint32),
	}
}

func getOrNewId(m map[string]uint32, newLocality string) uint32 {
	if id, ok := m[newLocality]; ok {
		return id
	}
	id := uint32(len(m))
	m[newLocality] = id
	return id
}

func (lc *localityCache) GetOrCreateRegion(region string) uint32 {
	lc.mutex.RLock()
	defer lc.mutex.RUnlock()

	return getOrNewId(lc.Region, region)

}

func (lc *localityCache) GetOrCreateZone(zone string) uint32 {
	lc.mutex.RLock()
	defer lc.mutex.RUnlock()

	return getOrNewId(lc.Zone, zone)
}

func (lc *localityCache) GetOrCreateSubzone(subzone string) uint32 {
	lc.mutex.RLock()
	defer lc.mutex.RUnlock()

	return getOrNewId(lc.Subzone, subzone)
}
