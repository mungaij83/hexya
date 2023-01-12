package models

import (
	"github.com/hexya-erp/hexya/src/models/loader"
	"github.com/hexya-erp/hexya/src/models/security"
	"time"
)

const (
	freeTransientPeriod = 1 * time.Minute
)

func BootStrap() {
	log.Info("Bootstrapping models")
	if Registry.bootstrapped == true {
		log.Info("Trying to bootstrap models twice ! Skipping..")
		return
	}
	loadManualSequencesFromDB()
	Registry.Lock()
	defer Registry.Unlock()
	log.Info("Migrate models")
	Registry.migrate() // Migrate all models
	createModelLinks()
	log.Info("Bootstrapping models security...")
	setupSecurity()
	log.Info("Register transient worker...")
	RegisterWorker(NewWorkerFunction(FreeTransientModels, freeTransientPeriod))
	Registry.bootstrapped = true
}

// loadManualSequencesFromDB fetches manual sequences from DB and updates registry
func loadManualSequencesFromDB() {
	sequences := loader.GetAdapter().Sequences("%_manseq")
	for _, dbSeq := range sequences {
		seq := &Sequence{
			JSON:      dbSeq.Name,
			Start:     dbSeq.StartValue,
			Increment: dbSeq.Increment,
		}
		Registry.addSequence(seq)
	}
}

//

// setupSecurity adds execution permission to:
// - the admin group for all methods
// - to CRUD methods to call "Load"
// - to "Create" method to call "Write"
// - to execute CRUD on context models
func setupSecurity() {
	for _, repository := range Registry.registryByTableName {
		model, ok := repository.GetModel()
		if !ok {
			log.Warn("Model is not initialized for", "table_name", repository.TableName())
			continue
		}
		loadMeth, loadExists := model.Methods().Get("Load")
		fetchMeth, fetchExists := model.Methods().Get("Fetch")
		writeMeth, writeExists := model.Methods().Get("Write")
		for _, meth := range model.Methods().Registry() {
			meth.AllowGroup(security.GroupAdmin)
			if loadExists && unauthorizedMethods[meth.Name()] {
				loadMeth.AllowGroup(security.GroupEveryone, meth)
			}
			if writeExists && meth.Name() == "Create" {
				writeMeth.AllowGroup(security.GroupEveryone, meth)
			}
		}
		if fetchExists {
			loadMeth.AllowGroup(security.GroupEveryone, fetchMeth)
		}
	}
	updateContextModelsSecurity()
}

// updateContextModelsSecurity synchronizes the methods permissions of context models with their base model.
func updateContextModelsSecurity() {
	for _, repo := range Registry.registryByTableName {
		model, ok := repo.GetModel()
		if !ok {
			continue
		}
		if !repo.isContext() {
			continue
		}
		baseModel := model.Fields().MustGet("Record").RelatedModel
		for _, methName := range []string{"Create", "Load", "Write", "Unlink"} {
			method := model.Methods().MustGet(methName)
			method.AllowGroup(security.GroupEveryone, baseModel.Methods().MustGet(methName))
			for grp := range baseModel.Methods().MustGet(methName).Groups {
				method.AllowGroup(grp)
			}
			for cGroup := range baseModel.Methods().MustGet(methName).GroupsCallers {
				method.AllowGroup(cGroup.Group, cGroup.Caller)
			}
		}
		model.Methods().MustGet("Load").AllowGroup(security.GroupEveryone, baseModel.Methods().MustGet("Create"))
	}
}

func createModelLinks() {
	for _, mi := range Registry.registryByTableName {
		mdl, ok := mi.GetModel()
		if !ok {
			log.Debug("Model not initialized", "table_name", mi.TableName())
			continue
		}
		for _, fi := range mdl.Fields().NameRegistry() {
			var (
				relatedMI *loader.Model
				ok        bool
			)
			if !fi.FieldType.IsRelationType() {
				continue
			}
			relatedMI, ok = Registry.Get(fi.RelatedModelName)
			if !ok {
				log.Panic("Unknown related model in field declaration", "model", mdl.Name(), "field", fi.Name(), "relatedName", fi.RelatedModelName)
			}
			if fi.FieldType.IsReverseRelationType() {
				fi.JsonReverseFK = relatedMI.Fields().MustGet(fi.ReverseFK).JSON()
			}
			fi.RelatedModel = relatedMI
		}
		mdl.Bootstrapped(true)
	}
}
