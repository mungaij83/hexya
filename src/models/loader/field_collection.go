package loader

import "github.com/hexya-erp/hexya/src/models"

type FieldsCollection struct {
	*models.FieldsCollection
}

// CreateDate returns a pointer to the CreateDate Field.
func (c FieldsCollection) CreateDate() *models.Field {
	return c.MustGet("CreateDate")
}

// CreateUID returns a pointer to the CreateUID Field.
func (c FieldsCollection) CreateUID() *models.Field {
	return c.MustGet("CreateUID")
}

// DisplayName returns a pointer to the DisplayName Field.
func (c FieldsCollection) DisplayName() *models.Field {
	return c.MustGet("DisplayName")
}

// ID returns a pointer to the ID Field.
func (c FieldsCollection) ID() *models.Field {
	return c.MustGet("ID")
}

// LastUpdate returns a pointer to the LastUpdate Field.
func (c FieldsCollection) LastUpdate() *models.Field {
	return c.MustGet("LastUpdate")
}

// WriteDate returns a pointer to the WriteDate Field.
func (c FieldsCollection) WriteDate() *models.Field {
	return c.MustGet("WriteDate")
}

// WriteUID returns a pointer to the WriteUID Field.
func (c FieldsCollection) WriteUID() *models.Field {
	return c.MustGet("WriteUID")
}

func (c FieldsCollection) GetField(name string) *models.Field {
	return c.MustGet(name)
}
