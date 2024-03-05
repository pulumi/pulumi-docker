// Copyright 2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import "github.com/pulumi/pulumi-go-provider/infer"

// These enum values are derived from
// https://github.com/docker/cli/blob/d1b88930/cli/command/manifest/util.go#L22-L51

var _ = (infer.Enum[Platform])((*Platform)(nil))

type Platform string

func (Platform) Values() []infer.EnumValue[Platform] {
	return []infer.EnumValue[Platform]{
		{Value: "darwin/386"},
		{Value: "darwin/amd64"},
		{Value: "darwin/arm"},
		{Value: "darwin/arm64"},
		{Value: "dragonfly/amd64"},
		{Value: "freebsd/386"},
		{Value: "freebsd/amd64"},
		{Value: "freebsd/arm"},
		{Value: "linux/386"},
		{Value: "linux/amd64"},
		{Value: "linux/arm"},
		{Value: "linux/arm64"},
		{Value: "linux/mips64"},
		{Value: "linux/mips64le"},
		{Value: "linux/ppc64le"},
		{Value: "linux/riscv64"},
		{Value: "linux/s390x"},
		{Value: "netbsd/386"},
		{Value: "netbsd/amd64"},
		{Value: "netbsd/arm"},
		{Value: "openbsd/386"},
		{Value: "openbsd/amd64"},
		{Value: "openbsd/arm"},
		{Value: "plan9/386"},
		{Value: "plan9/amd64"},
		{Value: "solaris/amd64"},
		{Value: "windows/386"},
		{Value: "windows/amd64"},
	}
}

func (p Platform) String() string {
	return string(p)
}
