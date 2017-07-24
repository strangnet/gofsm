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
	tests := []struct {
		name string
		want *FSM
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFSM(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFSM() = %v, want %v", got, tt.want)
			}
		})
	}
}