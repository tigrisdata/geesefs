// Copyright 2021 Yandex LLC
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

package cfg

import (
	"os"
	"syscall"
)

func redirectStdout(target *os.File) error {
	return syscall.Dup3(int(target.Fd()), int(os.Stdout.Fd()), 0)
}

func redirectStderr(target *os.File) error {
	return syscall.Dup3(int(target.Fd()), int(os.Stderr.Fd()), 0)
}
