package internal

import (
	"errors"
	"time"
)

type Semaphore interface {
	Acquire() error
	Release() error
	Len() int
}

type semaphore struct {
	chann chan struct{}
}

func NewSemaphore(tickets int) Semaphore {
	return &semaphore{chann: make(chan struct{}, tickets)}
}

func (s *semaphore) Acquire() error {
	select {
	case s.chann <- struct{}{}:
		return nil
	case <-time.After(time.Millisecond * 3000):
		return errors.New("fail acquire")
	}
}

func (s *semaphore) Release() error {
	select {
	case <-s.chann:
		return nil
	case <-time.After(time.Millisecond * 3000):
		return errors.New("fail release")
	}
}

func (s *semaphore) Len() int {
	return len(s.chann)
}
