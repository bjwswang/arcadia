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

package main

import (
	"k8s.io/klog/v2"

	"github.com/kubeagi/arcadia/apiserver/config"
	"github.com/kubeagi/arcadia/apiserver/service"

	_ "github.com/kubeagi/arcadia/apiserver/pkg/oidc"
)

func main() {
	conf := config.NewServerFlags()

	klog.Infof("listening server on port: %d", conf.Port)
	service.NewServerAndRun(conf)
}
