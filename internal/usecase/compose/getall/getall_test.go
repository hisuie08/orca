package getall

import (
	"errors"
	"orca/errs"
	"orca/internal/context"
	"orca/model/compose"
	"testing"
)

type fakeDockerInspector struct {
	ComposeMap map[string]*compose.ComposeSpec
	ErrMap     map[string]error
}

func (f *fakeDockerInspector) Compose(dir string) (*compose.ComposeSpec, error) {
	if err, ok := f.ErrMap[dir]; ok {
		return nil, err
	}
	if c, ok := f.ComposeMap[dir]; ok {
		return c, nil
	}
	return nil, errs.ErrComposeNotFound
}

type fakeFSInspector struct {
	DirsMap map[string][]string
	Err     error
}

func (f *fakeFSInspector) Dirs(path string) ([]string, error) {
	if f.Err != nil {
		return nil, f.Err
	}
	return f.DirsMap[path], nil
}
func TestGetAllCompose_AllFound(t *testing.T) {
	ctx := context.New().WithRoot("/root")

	fs := &fakeFSInspector{
		DirsMap: map[string][]string{
			"/root": {"/root/a", "/root/b"},
		},
	}

	docker := &fakeDockerInspector{
		ComposeMap: map[string]*compose.ComposeSpec{
			"/root/a": {},
			"/root/b": {},
		},
	}

	result, err := getAllCompose(&ctx, docker, fs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(*result) != 2 {
		t.Fatalf("expected 2 compose, got %d", len(*result))
	}

	if _, ok := (*result)["a"]; !ok {
		t.Errorf("compose a not found")
	}
	if _, ok := (*result)["b"]; !ok {
		t.Errorf("compose b not found")
	}
}

func TestGetAllCompose_SkipNotFound(t *testing.T) {
	ctx := context.New().WithRoot("/root")

	fs := &fakeFSInspector{
		DirsMap: map[string][]string{
			"/root": {"/root/a", "/root/b"},
		},
	}

	docker := &fakeDockerInspector{
		ComposeMap: map[string]*compose.ComposeSpec{
			"/root/a": {},
		},
	}

	result, err := getAllCompose(&ctx, docker, fs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(*result) != 1 {
		t.Fatalf("expected 1 compose, got %d", len(*result))
	}

	if _, ok := (*result)["a"]; !ok {
		t.Errorf("compose a should exist")
	}
}

func TestGetAllCompose_FilesystemError(t *testing.T) {
	ctx := context.New().WithRoot("/root")

	fs := &fakeFSInspector{
		Err: errors.New("fs error"),
	}

	docker := &fakeDockerInspector{}

	_, err := getAllCompose(&ctx, docker, fs)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}


func TestGetAllCompose_DockerError(t *testing.T) {
    ctx := context.New().WithRoot("/root")

    fs := &fakeFSInspector{
        DirsMap: map[string][]string{
            "/root": {"/root/a"},
        },
    }

    docker := &fakeDockerInspector{
        ErrMap: map[string]error{
            "/root/a": errors.New("docker down"),
        },
    }

    _, err := getAllCompose(&ctx, docker, fs)
    if err == nil {
        t.Fatal("expected error, got nil")
    }
}