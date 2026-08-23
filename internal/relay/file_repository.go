package observatory

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

type FileRepository struct {
	path       string
	mu         sync.Mutex
	beforeSave func()
	beforeLock func()
}

func OpenFileRepository(path string) (*FileRepository, error) {
	if path == "" {
		return nil, errors.New("path required")
	}
	return &FileRepository{path: path}, nil
}
func (r *FileRepository) Load(ctx context.Context, id string) (Workspace, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// BUG_BASE: cancellation is not rechecked after a save waits for the lock.
	data, err := os.ReadFile(r.path)
	if errors.Is(err, os.ErrNotExist) {
		return Workspace{}, ErrWorkspaceNotFound
	}
	if err != nil {
		return Workspace{}, err
	}
	w, err := DecodeWorkspace(data)
	if err != nil {
		return Workspace{}, err
	}
	if w.ID != id {
		return Workspace{}, ErrWorkspaceNotFound
	}
	return CloneWorkspace(w), nil
}
func (r *FileRepository) Save(ctx context.Context, w Workspace) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if r.beforeLock != nil {
		r.beforeLock()
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.beforeSave != nil {
		r.beforeSave()
	}
	data, err := EncodeWorkspace(w)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(r.path), 0755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(r.path), ".relay-")
	if err != nil {
		return err
	}
	name := tmp.Name()
	defer os.Remove(name)
	if _, err = tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err = tmp.Close(); err != nil {
		return err
	}
	return os.Rename(name, r.path)
}
func (r *FileRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	return os.Remove(r.path)
}
func (r *FileRepository) SetBeforeSave(fn func()) { r.beforeSave = fn }
func (r *FileRepository) SetBeforeLock(fn func()) { r.beforeLock = fn }
