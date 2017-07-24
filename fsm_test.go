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
	"reflect"
	"testing"
)

func TestNewFSM(t *testing.T) {
	type args struct {
		init        string
		transitions []Transition
		methods     Methods
	}
	tests := []struct {
		name string
		args args
		want *FSM
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFSM(tt.args.init, tt.args.transitions, tt.args.methods); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFSM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSM_State(t *testing.T) {
	tests := []struct {
		name string
		f    *FSM
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.State(); got != tt.want {
				t.Errorf("FSM.State() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSM_SetState(t *testing.T) {
	type args struct {
		state string
	}
	tests := []struct {
		name string
		f    *FSM
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.SetState(tt.args.state)
		})
	}
}

func TestFSM_Is(t *testing.T) {
	type args struct {
		state string
	}
	tests := []struct {
		name string
		f    *FSM
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Is(tt.args.state); got != tt.want {
				t.Errorf("FSM.Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSM_IsPending(t *testing.T) {
	tests := []struct {
		name string
		f    *FSM
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.IsPending(); got != tt.want {
				t.Errorf("FSM.IsPending() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSM_Can(t *testing.T) {
	type args struct {
		transition string
		args       []interface{}
	}
	tests := []struct {
		name    string
		f       *FSM
		args    args
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.Can(tt.args.transition, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FSM.Can() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FSM.Can() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSM_Cannot(t *testing.T) {
	type args struct {
		transition string
	}
	tests := []struct {
		name    string
		f       *FSM
		args    args
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.Cannot(tt.args.transition)
			if (err != nil) != tt.wantErr {
				t.Errorf("FSM.Cannot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FSM.Cannot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSM_Transition(t *testing.T) {
	tests := []struct {
		name    string
		f       *FSM
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.Transition(); (err != nil) != tt.wantErr {
				t.Errorf("FSM.Transition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFSM_doTransition(t *testing.T) {
	tests := []struct {
		name    string
		f       *FSM
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.doTransition(); (err != nil) != tt.wantErr {
				t.Errorf("FSM.doTransition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_transitionerStruct_transition(t *testing.T) {
	type args struct {
		f *FSM
	}
	tests := []struct {
		name    string
		t       transitionerStruct
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.transition(tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("transitionerStruct.transition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFSM_beforeEventCallback(t *testing.T) {
	type args struct {
		e *Event
	}
	tests := []struct {
		name    string
		f       *FSM
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.beforeEventCallback(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("FSM.beforeEventCallback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFSM_leaveStateCallback(t *testing.T) {
	type args struct {
		e *Event
	}
	tests := []struct {
		name    string
		f       *FSM
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.leaveStateCallback(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("FSM.leaveStateCallback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFSM_enterStateCallback(t *testing.T) {
	type args struct {
		e *Event
	}
	tests := []struct {
		name    string
		f       *FSM
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.enterStateCallback(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("FSM.enterStateCallback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFSM_afterEventCallback(t *testing.T) {
	type args struct {
		e *Event
	}
	tests := []struct {
		name    string
		f       *FSM
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.afterEventCallback(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("FSM.afterEventCallback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEvent_Abort(t *testing.T) {
	type args struct {
		err []error
	}
	tests := []struct {
		name string
		e    *Event
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.Abort(tt.args.err...)
		})
	}
}

func TestEvent_Async(t *testing.T) {
	tests := []struct {
		name string
		e    *Event
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.Async()
		})
	}
}
