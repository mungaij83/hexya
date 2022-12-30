package loader

import (
	"github.com/hexya-erp/hexya/src/models/security"
	"sync"
)

// A RecordRule allow to grant a Group some permissions
// on a selection of records.
// - If Global is true, then the RecordRule applies to all Groups
// - Condition is the filter to apply on the model to retrieve
// the records on which to allow the Perms permission.
type RecordRule struct {
	Name      string
	Global    bool
	Group     *security.Group
	Condition *Condition
	Perms     security.Permission
}

// A RecordRuleRegistry keeps a list of RecordRule. It is meant
// to be attached to a model.
type recordRuleRegistry struct {
	sync.RWMutex
	rulesByName  map[string]*RecordRule
	rulesByGroup map[string][]*RecordRule
	globalRules  map[string]*RecordRule
}

// AddRule registers the given RecordRule to the registry with the given name.
func (rrr *recordRuleRegistry) addRule(rule *RecordRule) {
	rrr.Lock()
	defer rrr.Unlock()
	rrr.rulesByName[rule.Name] = rule
	if rule.Global {
		rrr.globalRules[rule.Name] = rule
	} else {
		rrr.rulesByGroup[rule.Group.ID()] = append(rrr.rulesByGroup[rule.Group.ID()], rule)
	}
}

// RemoveRule removes the RecordRule with the given name
// from the rule registry.
func (rrr *recordRuleRegistry) removeRule(name string) {
	rrr.Lock()
	defer rrr.Unlock()
	rule, exists := rrr.rulesByName[name]
	if !exists {
		log.Warn("Trying to remove non-existent record rule", "name", name)
		return
	}
	delete(rrr.rulesByName, name)
	if rule.Global {
		delete(rrr.globalRules, name)
	} else {
		newRuleSlice := make([]*RecordRule, len(rrr.rulesByGroup[rule.Group.ID()])-1)
		i := 0
		for _, r := range rrr.rulesByGroup[rule.Group.ID()] {
			if r.Name == rule.Name {
				continue
			}
			newRuleSlice[i] = r
			i++
		}
		rrr.rulesByGroup[rule.Group.ID()] = newRuleSlice
	}
}

// newRecordRuleRegistry returns a pointer to a new RecordRuleRegistry instance
func newRecordRuleRegistry() *recordRuleRegistry {
	return &recordRuleRegistry{
		rulesByName:  make(map[string]*RecordRule),
		rulesByGroup: make(map[string][]*RecordRule),
		globalRules:  make(map[string]*RecordRule),
	}
}
