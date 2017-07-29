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
package gofsm

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestNewFSM(t *testing.T) {
	fsm := NewFSM(
		"closed",
		Transitions{
			{},
		},
		Methods{},
	)

	if fsm.State() != "closed" {
		t.Error("Expected state was 'closed'")
	}
}

func TestNewFSM_AllStates(t *testing.T) {
	fsm := NewFSM(
		"closed",
		Transitions{
			{Name: "open", From: []string{"closed"}, To: "open"},
			{Name: "close", From: []string{"open"}, To: "closed"},
		},
		Methods{},
	)

	found := fsm.AllStates()
	sort.StringSlice(found).Sort()
	expected := []string{"open", "closed"}
	sort.StringSlice(expected).Sort()
	if !reflect.DeepEqual(found, expected) {
		t.Errorf("Expected AllStates() to return the defined states. Got = %v, want = %v", found, expected)
	}
}

func TestNewFSM_AllTransitions(t *testing.T) {
	fsm := NewFSM(
		"closed",
		Transitions{
			{Name: "open", From: []string{"closed"}, To: "open"},
			{Name: "close", From: []string{"open"}, To: "closed"},
		},
		Methods{},
	)

	found := fsm.AllTransitions()
	sort.StringSlice(found).Sort()
	expected := []string{"open", "close"}
	sort.StringSlice(expected).Sort()
	if !reflect.DeepEqual(found, expected) {
		t.Errorf("Expected AllTransitions() to return the defined transitions. Got = %v, want = %v", found, expected)
	}
}

func TestIs_ReturnsTrueForCurrentState(t *testing.T) {
	fsm := NewFSM(
		"closed",
		Transitions{
			{},
		},
		Methods{},
	)

	if !fsm.Is(fsm.State()) {
		t.Error("Current state mismatch")
	}
}

func TestTransition_WithValidTransition(t *testing.T) {
	fsm := NewFSM(
		"closed",
		Transitions{
			{Name: "open", From: []string{"closed"}, To: "open"},
			{Name: "close", From: []string{"open"}, To: "closed"},
		},
		Methods{},
	)

	err := fsm.Transit("open")
	if err != nil {
		t.Error("Transitioning should have worked")
	}
	if fsm.State() != "open" {
		t.Error("State after the transition was expected to be 'open'")
	}
}

func TestTransition_WithInvalidTransition(t *testing.T) {
	fsm := NewFSM(
		"closed",
		Transitions{
			{Name: "open", From: []string{"closed"}, To: "open"},
			{Name: "close", From: []string{"open"}, To: "closed"},
		},
		Methods{},
	)

	err := fsm.Transit("push")
	if err == nil {
		t.Error("Transitioning should not have worked")
	}
	if fsm.State() != "closed" {
		t.Errorf("State should still be 'closed', was %v", fsm.State())
	}
}

func TestFSM_GenericMethods(t *testing.T) {
	beforeTransition := false
	leaveState := false
	enterState := false
	afterTransition := false

	fsm := NewFSM(
		"closed",
		Transitions{
			{Name: "open", From: []string{"closed"}, To: "open"},
		},
		Methods{
			"before_transition": func(t *Transition) {
				beforeTransition = true
			},
			"leave_state": func(t *Transition) {
				leaveState = true
			},
			"enter_state": func(t *Transition) {
				enterState = true
			},
			"after_transition": func(t *Transition) {
				afterTransition = true
			},
		},
	)

	fsm.Transit("open")
	if !(beforeTransition && leaveState && enterState && afterTransition) {
		t.Error("Expected all methods to be called")
	}

}

func ExampleFSM_State() {
	fsm := NewFSM(
		"closed",
		Transitions{
			{Name: "open", From: []string{"closed"}, To: "open"},
			{Name: "close", From: []string{"open"}, To: "closed"},
		},
		Methods{},
	)
	fmt.Println(fsm.State())
	// Output: closed
}
