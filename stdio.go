package main

import (
	"errors"
	"os"
)

type stdio struct{}

func (s *stdio) Read(p []byte) (n int, err error) {
	return os.Stdin.Read(p)
}

func (s *stdio) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func (s *stdio) Close() error {
	return errors.Join(os.Stdin.Close(), os.Stdout.Close())
}
