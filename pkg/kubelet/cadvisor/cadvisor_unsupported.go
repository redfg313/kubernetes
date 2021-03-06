// +build !cgo !linux

/*
Copyright 2015 Google Inc. All rights reserved.

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

package cadvisor

import (
	"errors"

	cadvisorApi "github.com/google/cadvisor/info/v1"
	cadvisorApiV2 "github.com/google/cadvisor/info/v2"
)

type cadvisorUnsupported struct {
}

var _ Interface = new(cadvisorUnsupported)

func New(port uint) (Interface, error) {
	return &cadvisorUnsupported{}, nil
}

var unsupportedErr = errors.New("cAdvisor is unsupported in this build")

func (self *cadvisorUnsupported) DockerContainer(name string, req *cadvisorApi.ContainerInfoRequest) (cadvisorApi.ContainerInfo, error) {
	return cadvisorApi.ContainerInfo{}, unsupportedErr
}

func (self *cadvisorUnsupported) ContainerInfo(name string, req *cadvisorApi.ContainerInfoRequest) (*cadvisorApi.ContainerInfo, error) {
	return nil, unsupportedErr
}

func (self *cadvisorUnsupported) MachineInfo() (*cadvisorApi.MachineInfo, error) {
	return nil, unsupportedErr
}

func (self *cadvisorUnsupported) VersionInfo() (*cadvisorApi.VersionInfo, error) {
	return nil, unsupportedErr
}

func (self *cadvisorUnsupported) DockerImagesFsInfo() (cadvisorApiV2.FsInfo, error) {
	return cadvisorApiV2.FsInfo{}, unsupportedErr
}
