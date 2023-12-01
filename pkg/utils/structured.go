/*
Copyright 2023 The KubeAGI Authors.

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

package utils

import (
	"encoding/json"
	"fmt"
	"reflect"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// UnstructuredToStructured convert unstructed object to a structured targe(must be a pointer)
func UnstructuredToStructured(unstructuredObj *unstructured.Unstructured, target any) error {
	// Check if the target is a pointer
	if reflect.TypeOf(target).Kind() != reflect.Ptr {
		return fmt.Errorf("target must be a pointer")
	}

	// Convert the unstructured object to JSON byte array
	jsonBytes, err := unstructuredObj.MarshalJSON()
	if err != nil {
		return err
	}

	// Unmarshal the JSON byte array into the structured custom resource type
	err = json.Unmarshal(jsonBytes, target)
	if err != nil {
		return err
	}

	return nil
}
