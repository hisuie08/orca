package fakeexecutor

import (
	"orca/internal/executor"
)

var _ executor.FileSystem = (*FakeFilesystem)(nil)

const (
	opWrite  = "WriteFile"
	opMkdir  = "CreateDir"
	opRmFile = "RemoveFile"
	opRmDir  = "RemoveDir"
)

type FakeFilesystem struct {
	Issued          []string
	Done            []string
	AllowSideEffect bool
}

func (f *FakeFilesystem) WriteFile(path string, data []byte) error {
	op := opWrite + ":" + path
	return f.recordOp(op)
}

func (f *FakeFilesystem) CreateDir(path string) error {
	op := opMkdir + ":" + path
	return f.recordOp(op)
}

func (f *FakeFilesystem) RemoveFile(path string) error {
	op := opRmFile + ":" + path
	return f.recordOp(op)
}

func (f *FakeFilesystem) RemoveDir(path string) error {
	op := opRmDir + ":" + path
	return f.recordOp(op)
}

func (f *FakeFilesystem) recordOp(op string) error {
	f.Issued = append(f.Issued, op)
	if f.AllowSideEffect {
		f.Done = append(f.Done, op)
	}
	return nil
}
