package loader

type FieldsCollections struct {
	*FieldsCollection
}

// CreateDate returns a pointer to the CreateDate Field.
func (c FieldsCollections) CreateDate() *Field {
	return c.MustGet("CreateDate")
}

// CreateUID returns a pointer to the CreateUID Field.
func (c FieldsCollections) CreateUID() *Field {
	return c.MustGet("CreateUID")
}

// DisplayName returns a pointer to the DisplayName Field.
func (c FieldsCollections) DisplayName() *Field {
	return c.MustGet("DisplayName")
}

// ID returns a pointer to the ID Field.
func (c FieldsCollections) ID() *Field {
	return c.MustGet("ID")
}

// LastUpdate returns a pointer to the LastUpdate Field.
func (c FieldsCollections) LastUpdate() *Field {
	return c.MustGet("LastUpdate")
}

// WriteDate returns a pointer to the WriteDate Field.
func (c FieldsCollections) WriteDate() *Field {
	return c.MustGet("WriteDate")
}

// WriteUID returns a pointer to the WriteUID Field.
func (c FieldsCollections) WriteUID() *Field {
	return c.MustGet("WriteUID")
}

func (c FieldsCollections) GetField(name string) *Field {
	return c.MustGet(name)
}
