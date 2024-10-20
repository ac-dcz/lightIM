package cache

import "sync"

type EdgeInfo struct {
	rwE      sync.RWMutex
	etcdKeys map[int64]string

	rwU        sync.RWMutex
	edgeOnline map[int64]map[int64]struct{}
}

func NewEdgeInfo() *EdgeInfo {
	return &EdgeInfo{
		etcdKeys:   make(map[int64]string),
		edgeOnline: make(map[int64]map[int64]struct{}),
	}
}

func (e *EdgeInfo) UpdateEtcdKey(key int64, value string) {
	e.rwE.Lock()
	defer e.rwE.Unlock()
	e.etcdKeys[key] = value
}

func (e *EdgeInfo) GetEtcdKey(edgeId int64) (string, bool) {
	e.rwE.RLock()
	defer e.rwE.RUnlock()
	key, ok := e.etcdKeys[edgeId]
	return key, ok
}

func (e *EdgeInfo) AddOnline(edgeId, uId int64) {
	e.rwU.Lock()
	defer e.rwU.Unlock()
	onlines, ok := e.edgeOnline[edgeId]
	if !ok {
		onlines = make(map[int64]struct{})
		e.edgeOnline[edgeId] = onlines
	}
	onlines[uId] = struct{}{}
}

func (e *EdgeInfo) RemOnline(edgeId, uId int64) {
	e.rwU.Lock()
	defer e.rwU.Unlock()
	onlines, ok := e.edgeOnline[edgeId]
	if !ok {
		onlines = make(map[int64]struct{})
		e.edgeOnline[edgeId] = onlines
	}
	delete(onlines, uId)
}

func (e *EdgeInfo) IsOnline(edgeId, uId int64) bool {
	e.rwU.RLock()
	defer e.rwU.RUnlock()
	if onlines, ok := e.edgeOnline[edgeId]; !ok {
		return false
	} else if _, ok := onlines[uId]; !ok {
		return false
	}
	return true
}

func (e *EdgeInfo) EdgeRouteByUID(uId int64) (int64, string, bool) {
	e.rwU.RLock()
	defer e.rwU.RUnlock()
	e.rwE.RLock()
	defer e.rwE.RUnlock()
	for edgeId, onlines := range e.edgeOnline {
		if _, ok := onlines[uId]; ok {
			return edgeId, e.etcdKeys[edgeId], true
		}
	}
	return 0, "", false
}
