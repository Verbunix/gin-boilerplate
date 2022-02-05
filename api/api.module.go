package api

import (
	"docsign/api/test"
	"docsign/utils/di"
)

func Module() di.Module {
	return di.Module{
		Imports: di.Modules{
			test.Module(),
		},
	}
}
