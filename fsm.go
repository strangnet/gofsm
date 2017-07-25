//    Copyright 2017 Patrick Strang <strangnet@gmail.com>
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Package gofsm implements a finite state machine
//
// It is based on the work done by Jake Gordon in JSM:
// https://github.com/jakesgordon/javascript-state-machine
//
// Also influenced by a Python implementation of JSM:
// https://github.com/mriehl/fysom
//
// and another Go library that used an earlier version of JSM as its base
// https://github.com/looplab/fsm
//
package gofsm

import (
	"sync"
)

// FSM is the state machine.
type FSM struct {
	state   string
	pending bool

	transitions map[string]Transition
	methods     map[string]Method

	stateMu sync.RWMutex
	transMu sync.Mutex
}

// Transition defines a command/event (Name) that holds which valid states
// that it can affect. A Transition can define multiple sources (From) but
// only a single destination (To).
type Transition struct {
	Name string
	From []string
	To   string
}

// Transitions is a helper type that holds a slice of Transitions.
type Transitions []Transition

// Method is a function type for the lifecycle events to call.
type Method func(*Transition)

// Methods is helper type that maps a string to a Method.
type Methods map[string]Method

// NewFSM initializes a new state machine.
func NewFSM(init string, transitions []Transition, methods Methods) *FSM {
	fsm := &FSM{
		state:       init,
		transitions: make(map[string]Transition),
		methods:     make(map[string]Method),
	}

	return fsm
}

// State returns the current state of an fsm.
func (f *FSM) State() string {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	return f.state
}

// SetState explicitly sets a new current state regardless.
// of current state
func (f *FSM) SetState(state string) {
	f.stateMu.Lock()
	defer f.stateMu.Unlock()
	f.state = state
	return
}

// Is returns whether the current state is the supplied state.
func (f *FSM) Is(state string) bool {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	return state == f.state
}

func (f *FSM) isPending() bool {
	return f.pending
}

// Can checks if a transition is valid from the current state.
func (f *FSM) Can(transition string) bool {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()

	return !f.isPending()
}

// Cannot checks if a transition is not valid from the current state.
func (f *FSM) Cannot(transition string) bool {
	return !f.Can(transition)
}
