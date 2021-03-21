package gomodule

import (
	"fmt"
	"github.com/google/blueprint"
	"github.com/roman-mazur/bood"
	"path"
)

var (
	goDoc = pctx.StaticRule("doc", blueprint.RuleParams{
		Command:          "cd $workDir && go doc -all -u $pkg > $docFile",
		Description:      "godoc $pkg",
	},   "workDir","pkg", "docFile")
)

type goDocModuleType struct {
	blueprint.SimpleName

	properties struct {
		Pkg         string
		Srcs        []string
		// Exclude patterns.
		SrcsExclude []string
		Deps        []string
	}
}

func (gd *goDocModuleType) DynamicDependencies(blueprint.DynamicDependerModuleContext) []string {
	return gd.properties.Deps
}

func (gd *goDocModuleType) GenerateBuildActions(ctx blueprint.ModuleContext) {
	name := ctx.ModuleName()
	config := bood.ExtractConfig(ctx)
	config.Info.Printf("Adding doc actions for go module '%s'", name)
	outputPath := path.Join(config.BaseOutputDir, "docs", "my-docs.txt")
	var buildInputs []string
	inputErrors := false

	for _, src := range gd.properties.Srcs {
		if matches, err := ctx.GlobWithDeps(src, gd.properties.SrcsExclude); err == nil {
			buildInputs = append(buildInputs, matches...)
		} else {
			ctx.PropertyErrorf("srcs", "Cannot resolve files that match pattern %s", src)
			inputErrors = true
		}
	}

	if inputErrors {
		return
	}

	ctx.Build(pctx, blueprint.BuildParams{
		Description: fmt.Sprintf("Generating docs for %s", name),
		Rule:        goDoc,
		Outputs:     []string{outputPath},
		Implicits:   buildInputs,
		Args: map[string]string{
			"workDir":    ctx.ModuleDir(),
			"pkg":        gd.properties.Pkg,
			"docFile":   outputPath,
		},
	})
}

func SimpleDocFactory() (blueprint.Module, []interface{}) {
	DocType := &goDocModuleType{}
	return DocType, []interface{}{&DocType.SimpleName.Properties, &DocType.properties}
}
