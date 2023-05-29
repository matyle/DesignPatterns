package ch8

import (
	"log"
	"testing"
)

func TestFactory(t *testing.T) {
	log.Println("FactoryRun[float64](1, 23) = ", FactoryRun[float64](1, 23))
}
