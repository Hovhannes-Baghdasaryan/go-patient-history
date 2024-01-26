package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Patient holds the schema definition for the Patient entity.
type PatientEntity struct {
	ent.Schema
}

// Fields of the Patient.
func (PatientEntity) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
		field.String("surname"),
		field.String("patronymic").Optional(),
		field.Int("age").Positive().Max(100),
		field.Enum("gender").Values("male", "female"),
		field.String("nationality"),
	}
}
