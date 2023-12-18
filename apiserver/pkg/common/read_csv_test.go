/*
Copyright 2023 KubeAGI.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package common

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

const csvData = `a,a,a
b,b,b
c,c,c
d,d,d
e,e,e
f,f,f
g,g,g
i,i,i
1,1,1
2,2,2
3,3,3
4,4,4
5,5,5
6,6,6`

func TestReadCSV(t *testing.T) {
	type input struct {
		startLine, size int64
		exp             [][]string
		lines           int64
		expErr          error
	}

	reader := bytes.NewReader([]byte(csvData))
	for _, tc := range []input{
		{
			1, 1, [][]string{{"a", "a", "a"}}, 14, io.EOF,
		},
		{
			1, 2, [][]string{{"a", "a", "a"}, {"b", "b", "b"}}, 14, io.EOF,
		},
		{
			2, 1, [][]string{{"b", "b", "b"}}, 14, io.EOF,
		},
		{
			2, 2, [][]string{{"b", "b", "b"}, {"c", "c", "c"}}, 14, io.EOF,
		},
		{
			9, 10, [][]string{{"1", "1", "1"}, {"2", "2", "2"}, {"3", "3", "3"}, {"4", "4", "4"}, {"5", "5", "5"}, {"6", "6", "6"}}, 14, io.EOF,
		},
		{
			14, 1, [][]string{{"6", "6", "6"}}, 14, io.EOF,
		},
		{
			14, 2, [][]string{{"6", "6", "6"}}, 14, io.EOF,
		},
		{
			8, 3, [][]string{{"i", "i", "i"}, {"1", "1", "1"}, {"2", "2", "2"}}, 14, io.EOF,
		},
		{
			1, 15, [][]string{
				{"a", "a", "a"},
				{"b", "b", "b"},
				{"c", "c", "c"},
				{"d", "d", "d"},
				{"e", "e", "e"},
				{"f", "f", "f"},
				{"g", "g", "g"},
				{"i", "i", "i"},
				{"1", "1", "1"},
				{"2", "2", "2"},
				{"3", "3", "3"},
				{"4", "4", "4"},
				{"5", "5", "5"},
				{"6", "6", "6"},
			}, 14, io.EOF,
		},
		{
			1, 14, [][]string{
				{"a", "a", "a"},
				{"b", "b", "b"},
				{"c", "c", "c"},
				{"d", "d", "d"},
				{"e", "e", "e"},
				{"f", "f", "f"},
				{"g", "g", "g"},
				{"i", "i", "i"},
				{"1", "1", "1"},
				{"2", "2", "2"},
				{"3", "3", "3"},
				{"4", "4", "4"},
				{"5", "5", "5"},
				{"6", "6", "6"},
			}, 14, io.EOF,
		},
		{
			15, 2, nil, 14, io.EOF,
		},
		{
			// page=1, size=3
			1, 3, [][]string{
				{"a", "a", "a"},
				{"b", "b", "b"},
				{"c", "c", "c"},
			}, 14, io.EOF,
		},
		{
			// page=2,size=3
			4, 3, [][]string{
				{"d", "d", "d"},
				{"e", "e", "e"},
				{"f", "f", "f"},
			}, 14, io.EOF,
		},
		{
			// page=3,size=3
			7, 3, [][]string{
				{"g", "g", "g"},
				{"i", "i", "i"},
				{"1", "1", "1"},
			}, 14, io.EOF,
		},
		{
			// page=4,size=3
			10, 3, [][]string{
				{"2", "2", "2"},
				{"3", "3", "3"},
				{"4", "4", "4"},
			}, 14, io.EOF,
		},
		{
			// page=5,size=3
			13, 3, [][]string{
				{"5", "5", "5"},
				{"6", "6", "6"},
			}, 14, io.EOF,
		},
	} {
		r, totalLines, err := ReadCSV(reader, tc.startLine, tc.size)
		if err != tc.expErr || !reflect.DeepEqual(tc.exp, r) || totalLines != tc.lines {
			t.Fatalf("input: %+v expect %v get %v, expect error %v get %v, expect lines %d got %d", tc, tc.exp, r, tc.expErr, err, tc.lines, totalLines)
		}
		_, _ = reader.Seek(0, io.SeekStart)
	}
}
