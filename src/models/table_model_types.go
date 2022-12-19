package models

import "github.com/hexya-erp/hexya/src/models/types"

// Different types of models using go interface

type TableModel interface {
	TableName() string
}

type OrderedTableModel interface {
	OrderFields() []string
}

type MixinModelType interface {
	ParentTableName() string
}

type ManualModelType interface {
	ExistingTableName() string
}

type TransientTableModel interface {
	BaseMixin() string
	TimeoutMinutes() uint
}

type EnumWrapper interface {
	Values() types.Selection
}
