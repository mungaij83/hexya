package models

import (
	"fmt"
	"github.com/hexya-erp/hexya/src/models/loader"
	"github.com/hexya-erp/hexya/src/tools/strutils"
)

// A Sequence holds the metadata of a DB sequence
//
// There are two types of sequences: those created before bootstrap
// and those created after. The former will be created and updated at
// bootstrap and cannot be modified afterwards. The latter will be
// created, updated or dropped immediately.
type Sequence struct {
	JSON      string
	Increment int64
	Start     int64
	boot      bool
}

// CreateSequence creates a new Sequence in the database and returns a pointer to it
func CreateSequence(name string, increment, start int64) *Sequence {
	var boot bool
	suffix := "manseq"
	if !Registry.bootstrapped {
		boot = true
		suffix = "bootseq"
	}
	json := fmt.Sprintf("%s_%s", strutils.SnakeCase(name), suffix)
	seq := &Sequence{
		JSON:      json,
		Increment: increment,
		Start:     start,
		boot:      boot,
	}
	if !boot {
		// Create the sequence on the fly if we already bootstrapped.
		// Otherwise, this will be done in Bootstrap
		loader.GetAdapter().CreateSequence(seq.JSON, seq.Increment, seq.Start)
	}
	Registry.addSequence(seq)
	return seq
}

// Drop this sequence and removes it from the database
func (s *Sequence) Drop() {
	Registry.Lock()
	defer Registry.Unlock()
	delete(Registry.sequences, s.JSON)
	if Registry.bootstrapped {
		// Drop the sequence on the fly if we already bootstrapped.
		// Otherwise, this will be done in Bootstrap
		if s.boot {
			log.Panic("Boot Sequences cannot be dropped after bootstrap")
		}
		loader.GetAdapter().DropSequence(s.JSON)
	}
}

// Alter alters this sequence by changing next number and/or increment.
// Set a parameter to 0 to leave it unchanged.
func (s *Sequence) Alter(increment, restart int64) {
	var boot bool
	if !Registry.bootstrapped {
		boot = true
	}
	if s.boot && !boot {
		log.Panic("Boot Sequences cannot be modified after bootstrap")
	}
	if restart > 0 {
		s.Start = restart
	}
	if increment > 0 {
		s.Increment = increment
	}
	if !boot {
		loader.GetAdapter().AlterSequence(s.JSON, increment, restart)
	}
}

// NextValue returns the next value of this Sequence
func (s *Sequence) NextValue() int64 {
	return loader.GetAdapter().NextSequenceValue(s.JSON)
}
