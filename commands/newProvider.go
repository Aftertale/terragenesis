package commands

import (
	"github.com/spf13/cobra"
)

var _ cmder = (*genProviderCmd)(nil)

type genProviderCmd stuct {
	sourceProviderDefinition string

	definitionType string

	*baseCmd
}

func newGenerateProviderCmd(){
	cc := &genProviderCmd{}

	return cc
}
