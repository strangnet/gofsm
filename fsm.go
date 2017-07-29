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
	"errors"
	"sync"
)

// FSM is the state machine.
type FSM struct {
	state   string
	pending bool

	transitions map[transitionKey]Transition
	methods     map[string]Method

	states map[string]bool

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
		transitions: make(map[transitionKey]Transition),
		methods:     make(map[string]Method),
		states:      make(map[string]bool),
	}

	fsm.addState(init)

	for _, t := range transitions {
		for _, from := range t.From {
			fsm.transitions[transitionKey{t.Name, from}] = t
			fsm.addState(from)
		}
		fsm.addState(t.To)
	}

	fsm.methods = methods

	return fsm
}

func (f *FSM) addState(state string) {
	if !f.states[state] {
		f.states[state] = true
	}
}

// State returns the current state of an fsm.
func (f *FSM) State() string {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	return f.state
}

// Is returns whether the current state is the supplied state.
func (f *FSM) Is(state string) bool {
	return state == f.State()
}

func (f *FSM) isPending() bool {
	return f.pending
}

// Can checks if a transition is valid from the current state.
// TODO: Proper implementation
func (f *FSM) Can(transition string) bool {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()

	return !f.isPending()
}

// Cannot checks if a transition is not valid from the current state.
func (f *FSM) Cannot(transition string) bool {
	return !f.Can(transition)
}

// AllStates returns all the available states
func (f *FSM) AllStates() []string {
	var states []string
	for s := range f.states {
		states = append(states, s)
	}

	return states
}

// AllTransitions returns all the transitions
func (f *FSM) AllTransitions() []string {
	f.transMu.Lock()
	defer f.transMu.Unlock()

	var t []string
	for _, v := range f.transitions {
		t = append(t, v.Name)
	}

	return t
}

// Transitions returns all available transition from the current state
func (f *FSM) Transitions() []string {
	f.transMu.Lock()
	defer f.transMu.Unlock()

	current := f.State()

	t := make([]string, len(f.transitions))
	for _, v := range f.transitions {
		for _, from := range v.From {
			if current == from {
				t = append(t, v.Name)
				break
			}
		}
	}

	return t
}

// Transit executes a transition
func (f *FSM) Transit(transition string) error {
	f.transMu.Lock()
	defer f.transMu.Unlock()

	f.stateMu.RLock()
	defer f.stateMu.RUnlock()

	t, ok := f.transitions[transitionKey{transition, f.state}]
	if !ok {
		return errors.New("Transition is not valid")
	}

	// TODO: clean up this placeholder routine just to get the test passing
	for _, method := range f.methods {
		method(&t)
	}

	f.stateMu.RUnlock()

	f.stateMu.Lock()
	f.state = t.To
	f.stateMu.Unlock()

	f.stateMu.RLock()

	return nil
}

func (f *FSM) beginTransit() {
	f.pending = true
}

func (f *FSM) endTransit(result interface{}) interface{} {
	f.pending = false
	return result
}

func (f *FSM) failTransit(err error) error {
	f.pending = false
	return err
}

// transitionKey is the key used for mapping the transition in the fsm
type transitionKey struct {
	Name string
	From string
}
