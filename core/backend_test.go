// Copyright 2020 Ka-Hing Cheung
// Copyright 2020 Databricks
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

package core

import "sync"

type TestBackend struct {
	StorageBackend
	ListBlobsFunc           func(param *ListBlobsInput) (*ListBlobsOutput, error)
	HeadBlobFunc            func(param *HeadBlobInput) (*HeadBlobOutput, error)
	MultipartBlobAddFunc    func(param *MultipartBlobAddInput) (*MultipartBlobAddOutput, error)
	MultipartBlobCopyFunc   func(param *MultipartBlobCopyInput) (*MultipartBlobCopyOutput, error)
	MultipartBlobCommitFunc func(param *MultipartBlobCommitInput) (*MultipartBlobCommitOutput, error)
	capabilities            *Capabilities
	err                     error
	lock                    sync.RWMutex
}

func NewTestBackend(init *TestBackend) *TestBackend {
	t := &TestBackend{}
	t.lock.Lock()
	defer t.lock.Unlock()
	t.StorageBackend = init.StorageBackend
	t.ListBlobsFunc = init.ListBlobsFunc
	t.HeadBlobFunc = init.HeadBlobFunc
	t.MultipartBlobAddFunc = init.MultipartBlobAddFunc
	t.MultipartBlobCopyFunc = init.MultipartBlobCopyFunc
	t.MultipartBlobCommitFunc = init.MultipartBlobCommitFunc
	t.capabilities = init.capabilities
	t.err = init.err
	return t
}

func (s *TestBackend) Init(bucket string) error {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.StorageBackend == nil {
		return nil
	}
	return s.StorageBackend.Init(bucket)
}

func (s *TestBackend) Capabilities() *Capabilities {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.StorageBackend == nil {
		if s.capabilities == nil {
			s.capabilities = &Capabilities{
				Name:             "s3",
				MaxMultipartSize: 5 * 1024 * 1024 * 1024,
			}
		}
		return s.capabilities
	}
	return s.StorageBackend.Capabilities()
}

func (s *TestBackend) SetErr(err error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.err = err
}

func (s *TestBackend) Delegate() interface{} {
	return s
}

func (s *TestBackend) HeadBlob(param *HeadBlobInput) (*HeadBlobOutput, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.HeadBlobFunc != nil {
		return s.HeadBlobFunc(param)
	}
	if s.err != nil {
		return nil, s.err
	}
	return s.StorageBackend.HeadBlob(param)
}

func (s *TestBackend) ListBlobs(param *ListBlobsInput) (*ListBlobsOutput, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.ListBlobsFunc != nil {
		return s.ListBlobsFunc(param)
	}
	if s.err != nil {
		return nil, s.err
	}
	return s.StorageBackend.ListBlobs(param)
}

func (s *TestBackend) DeleteBlob(param *DeleteBlobInput) (*DeleteBlobOutput, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.err != nil {
		return nil, s.err
	}
	return s.StorageBackend.DeleteBlob(param)
}

func (s *TestBackend) DeleteBlobs(param *DeleteBlobsInput) (*DeleteBlobsOutput, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.err != nil {
		return nil, s.err
	}
	return s.StorageBackend.DeleteBlobs(param)
}

func (s *TestBackend) RenameBlob(param *RenameBlobInput) (*RenameBlobOutput, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.err != nil {
		return nil, s.err
	}
	return s.StorageBackend.RenameBlob(param)
}

func (s *TestBackend) CopyBlob(param *CopyBlobInput) (*CopyBlobOutput, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.err != nil {
		return nil, s.err
	}
	return s.StorageBackend.CopyBlob(param)
}

func (s *TestBackend) GetBlob(param *GetBlobInput) (*GetBlobOutput, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.err != nil {
		return nil, s.err
	}
	return s.StorageBackend.GetBlob(param)
}

func (s *TestBackend) PutBlob(param *PutBlobInput) (*PutBlobOutput, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.err != nil {
		return nil, s.err
	}
	return s.StorageBackend.PutBlob(param)
}

func (s *TestBackend) MultipartBlobBegin(param *MultipartBlobBeginInput) (*MultipartBlobCommitInput, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.err != nil {
		return nil, s.err
	}
	return s.StorageBackend.MultipartBlobBegin(param)
}

func (s *TestBackend) MultipartBlobAdd(param *MultipartBlobAddInput) (*MultipartBlobAddOutput, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.MultipartBlobAddFunc != nil {
		return s.MultipartBlobAddFunc(param)
	}
	if s.err != nil {
		return nil, s.err
	}
	return s.StorageBackend.MultipartBlobAdd(param)
}

func (s *TestBackend) MultipartBlobCopy(param *MultipartBlobCopyInput) (*MultipartBlobCopyOutput, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.MultipartBlobCopyFunc != nil {
		return s.MultipartBlobCopyFunc(param)
	}
	if s.err != nil {
		return nil, s.err
	}
	return s.StorageBackend.MultipartBlobCopy(param)
}

func (s *TestBackend) MultipartBlobAbort(param *MultipartBlobCommitInput) (*MultipartBlobAbortOutput, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.err != nil {
		return nil, s.err
	}
	return s.StorageBackend.MultipartBlobAbort(param)
}

func (s *TestBackend) MultipartBlobCommit(param *MultipartBlobCommitInput) (*MultipartBlobCommitOutput, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.MultipartBlobCommitFunc != nil {
		return s.MultipartBlobCommitFunc(param)
	}
	if s.err != nil {
		return nil, s.err
	}
	return s.StorageBackend.MultipartBlobCommit(param)
}

func (s *TestBackend) MultipartExpire(param *MultipartExpireInput) (*MultipartExpireOutput, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.err != nil {
		return nil, s.err
	}
	return s.StorageBackend.MultipartExpire(param)
}
