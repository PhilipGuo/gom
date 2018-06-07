package server

import (
	"sync"
)

//
//
type SessionBucket struct {
	smap map[string]*Session

	mu *sync.RWMutex
}

// NewSessionBucket
//
func NewSessionBucket() *SessionBucket {
	return &SessionBucket{
		smap: make(map[string]*Session),
		mu:   new(sync.RWMutex),
	}
}

// Get session by addr
//
func (sbk *SessionBucket) Get(remoteaddr string) (bool, *Session) {
	if nil == sbk || nil == sbk.smap {
		return false, nil
	}
	s, ok := sbk.smap[remoteaddr]
	return ok, s
}

// Put session into map by addr
//
func (sbk *SessionBucket) Put(remoteaddr string, s *Session) bool {
	if nil == sbk || nil == sbk.smap {
		return false
	}

	sbk.smap[remoteaddr] = s
	return true
}

// Delete session by addr
//
func (sbk *SessionBucket) Delete(remoteaddr string) bool {
	if nil == sbk || nil == sbk.smap {
		return false
	}
	delete(sbk.smap, remoteaddr)
	return true
}
