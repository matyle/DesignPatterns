package ch14

import "testing"

func TestOriObserverRun(t *testing.T) {
	ObserverV1()
	ObserverV2()
	ObserverV3()
	DelegateMain()
}
