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
// It is heavily based on the work done by Jake Gordon in:
// https://github.com/jakesgordon/javascript-state-machine
//
// Also influenced by a Python implementation:
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

type transitioner interface {
	transition(*FSM) error
}

// FSM is the state machine
type FSM struct {
	state string

	transitions map[tKey]string
	methods     map[mKey]Method

	transition      func()
	transitionerObj transitioner

	stateMu sync.RWMutex
	transMu sync.Mutex
}

type Transition struct {
	Name string
	From []string
	To   string
}

type Transitions []Transition

type Method func(*Transition)

type Methods map[string]Method

// NewFSM initializes a new state machine
func NewFSM(init string, transitions []Transition, methods Methods) *FSM {
	fsm := &FSM{
		state:           init,
		transitions:     make(map[tKey]string),
		methods:         make(map[mKey]Method),
		transitionerObj: &transitionerStruct{},
	}

	allStates := make(map[string]bool)
	allTransitions := make(map[string]bool)
	for _, t := range transitions {
		for _, from := range t.From {
			fsm.transitions[tKey{t.Name, from}] = t.To
			allStates[from] = true
			allStates[t.To] = true
		}
		allTransitions[t.Name] = true
	}

	for name, fn := range methods {
		var target string
		var cbType int

		switch {
		default:
			target = name
			if _, ok := allStates[target]; ok {
				cbType = callbackEnterState
			} else if _, ok := allTransitions[target]; ok {
				cbType = callbackAfterTransition
			}
		}

		if cbType != callbackNone {
			fsm.methods[mKey{target, cbType}] = fn
		}
	}

	return fsm
}

func (f *FSM) State() string {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	return f.state
}

func (f *FSM) SetState(state string) {
	f.stateMu.Lock()
	defer f.stateMu.Unlock()
	f.state = state
	return
}

func (f *FSM) Is(state string) bool {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	return state == f.state
}

func (f *FSM) IsPending() bool {
	return false
}

func (f *FSM) Can(transition string, args ...interface{}) (bool, error) {
	f.stateMu.RLock()
	defer f.stateMu.RUnlock()
	to, ok := f.transitions[tKey{transition, f.state}]
	if !ok || (f.transition != nil) {
		return false, nil
	}

	e := &Event{f, transition, f.state, to, nil, args, false, false}
	err := f.beforeEventCallback(e)

	return err == nil, err
}

func (f *FSM) Cannot(transition string) (bool, error) {
	ok, err := f.Can(transition)
	return !ok, err
}

func (f *FSM) Transition() error {
	f.transMu.Lock()
	defer f.transMu.Unlock()
	return f.doTransition()
}

func (f *FSM) doTransition() error {
	return f.transitionerObj.transition(f)
}

type transitionerStruct struct{}

func (t transitionerStruct) transition(f *FSM) error {
	if f.transition == nil {
		return errors.New("Error")
	}
	f.transition()
	f.transition = nil
	return nil
}

func (f *FSM) beforeEventCallback(e *Event) error {

	return nil
}

func (f *FSM) leaveStateCallback(e *Event) error {

	return nil
}

func (f *FSM) enterStateCallback(e *Event) error {

	return nil
}

func (f *FSM) afterEventCallback(e *Event) error {

	return nil
}

const (
	callbackNone int = iota
	callbackBeforeTransition
	callbackLeaveState
	callbackEnterState
	callbackAfterTransition
)

type mKey struct {
	target string
	cbType int
}

type tKey struct {
	transition string

	src string
}

type Event struct {
	FSM   *FSM
	Event string
	From  string
	To    string
	Err   error
	Args  []interface{}

	aborted bool
	async   bool
}

func (e *Event) Abort(err ...error) {
	e.aborted = true
	if len(err) > 0 {
		e.Err = err[0]
	}
}

func (e *Event) Async() {
	e.async = true
}
