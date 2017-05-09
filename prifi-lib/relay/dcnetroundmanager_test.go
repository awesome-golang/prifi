package relay

import (
	"testing"
)

func TestDCNetRound(test *testing.T) {

	data := make([]byte, 101)
	window := 10
	dcmr := NewDCNetRoundManager(window)

	if dcmr.CurrentRound() != 1 {
		test.Error("Should be in round 1")
	}
	if !dcmr.CurrentRoundIsStill(1) {
		test.Error("Should still be in round 1")
	}

	//requesting the next downstream round to send should not return an open round
	if dcmr.NextDownStreamRoundToSent() != 1 {
		test.Error("NextDownStreamRoundToSent should be equal to 1", dcmr.NextDownStreamRoundToSent())
	}
	//but should still return the same number
	if dcmr.NextDownStreamRoundToSent() != 1 {
		test.Error("NextDownStreamRoundToSent should still be equal to 1", dcmr.NextDownStreamRoundToSent())
	}

	//opening another round should not change current round
	dcmr.OpenRound(1)
	if dcmr.CurrentRound() != 1 {
		test.Error("Should be in round 1")
	}

	//requesting the next downstream round to send should not return an open round
	if dcmr.NextDownStreamRoundToSent() != 2 {
		test.Error("NextDownStreamRoundToSent should be equal to 2", dcmr.NextDownStreamRoundToSent())
	}

	//setting a round to closed should skip it
	s := make(map[int32]bool, 2)
	s[2] = false
	s[4] = false
	dcmr.SetStoredRoundSchedule(s)
	if dcmr.storedRoundsSchedule == nil || len(dcmr.storedRoundsSchedule) != len(s) || dcmr.storedRoundsSchedule[0] != s[0] {
		test.Error("dcmr.storedRoundsSchedule should be s")
	}
	if dcmr.NextDownStreamRoundToSent() != 3 {
		test.Error("NextDownStreamRoundToSent should be equal to 3", dcmr.NextDownStreamRoundToSent())
	}

	//should be able to open a round while skipping another round
	dcmr.OpenRound(3)
	dcmr.OpenRound(5)
	if dcmr.CurrentRound() != 1 {
		test.Error("Should be in round 1")
	}
	dcmr.CloseRound(1)
	if dcmr.CurrentRound() != 3 {
		test.Error("Should be in round 3")
	}
	dcmr.CloseRound(3)
	if dcmr.CurrentRound() != 5 {
		test.Error("Should be in round 5", dcmr.CurrentRound())
	}

	_ = data

}
