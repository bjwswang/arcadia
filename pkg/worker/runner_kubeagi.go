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

package worker

import (
	"context"
	"errors"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	arcadiav1alpha1 "github.com/kubeagi/arcadia/api/base/v1alpha1"
)

const (
	defaultCoreLibraryCLIImage = "kubeagi/core-library-cli:v0.0.1"
)

var _ ModelRunner = (*CoreLibraryCLIRunner)(nil)

// CoreLibraryCLIRunner utilizes  core-library-cli(https://github.com/kubeagi/core-library/tree/main/libs/cli) to run model services
// Mainly for reranking,whisper,etc..
type CoreLibraryCLIRunner struct {
	c client.Client
	w *arcadiav1alpha1.Worker
}

func NewCoreLibraryCLIRunner(c client.Client, w *arcadiav1alpha1.Worker) (ModelRunner, error) {
	return &CoreLibraryCLIRunner{
		c: c,
		w: w,
	}, nil
}

// Device used when running model
func (runner *CoreLibraryCLIRunner) Device() Device {
	return DeviceBasedOnResource(runner.w.Spec.Resources.Limits)
}

// NumberOfGPUs utilized by this runner
func (runner *CoreLibraryCLIRunner) NumberOfGPUs() string {
	return NumberOfGPUs(runner.w.Spec.Resources.Limits)
}

// Build a model runner instance
func (runner *CoreLibraryCLIRunner) Build(ctx context.Context, model *arcadiav1alpha1.TypedObjectReference) (any, error) {
	if model == nil {
		return nil, errors.New("nil model")
	}

	img := defaultCoreLibraryCLIImage
	if runner.w.Spec.Runner.Image != "" {
		img = runner.w.Spec.Runner.Image
	}

	// read worker address
	mountPath := "/data/models"
	container := &corev1.Container{
		Name:            "runner",
		Image:           img,
		ImagePullPolicy: runner.w.Spec.Runner.ImagePullPolicy,
		Command: []string{
			"python", "kubeagi_cli/cli.py", "serve", "--host", "0.0.0.0", "--port", "21002",
		},
		Env: []corev1.EnvVar{
			// Only reranking supported for now
			{Name: "RERANKING_MODEL_PATH", Value: fmt.Sprintf("%s/%s", mountPath, model.Name)},
		},
		Ports: []corev1.ContainerPort{
			{Name: "http", ContainerPort: 21002},
		},
		VolumeMounts: []corev1.VolumeMount{
			{Name: "models", MountPath: mountPath},
		},
		Resources: runner.w.Spec.Resources,
	}

	return container, nil
}
