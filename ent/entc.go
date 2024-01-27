//go:build ignore

package main

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"fmt"
	constant "go-patient-history/libs/common/constant/logger"
	logger "go-patient-history/libs/common/logger/main"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{gen.FeatureVersionedMigration},
	})
	if err != nil {
		logger.LogError(logger.LoggerPayload{FuncName: constant.GenerateSchema, Message: fmt.Sprintf("Running ent codegen: %v", err.Error())})
	}
}
