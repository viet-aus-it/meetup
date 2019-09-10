package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func (s *server) digNewHole() string {
	s.metrics.holes_available.Add(1)
	logrus.Println("digged a new hole")
	return fmt.Sprintf("a hole has been digged")
}

func (s *server) closeHole() string {
	s.metrics.holes_available.Dec()
	logrus.Println("closed a hole")
	return fmt.Sprintf("a hole has been closed")
}

func (s *server) setHole(holes int) string {
	s.metrics.holes_available.Set(float64(holes))
	logrus.Println("discovered holes")
	return fmt.Sprintf("set to %s holes", holes)
}