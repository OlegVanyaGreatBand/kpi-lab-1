package gomodule

import (
	"fmt"
	"github.com/google/blueprint"
	"github.com/roman-mazur/bood"
	"path"
)
//cd $workDir &&
var (
	goDoc = pctx.StaticRule("doc", blueprint.RuleParams{
		Command:          "echo $pkg > $htmlFile && cd $workDir && go doc $pkg > $htmlFile",
		Description:      "godoc $pkg",
	},  "pkg", "htmlFile", "workDir")
)

type goDocModuleType struct {
	blueprint.SimpleName

	properties struct {
		Pkg         string
		Srcs        []string
		// Exclude patterns.
		TestSrcs    []string
		VendorFirst bool
		Deps        []string
	}
}

func (gd *goDocModuleType) DynamicDependencies(blueprint.DynamicDependerModuleContext) []string {
	return gd.properties.Deps
}

func (gd *goDocModuleType) GenerateBuildActions(ctx blueprint.ModuleContext) {
	name := ctx.ModuleName()
	config := bood.ExtractConfig(ctx)
	config.Info.Printf("Adding build & test actions for go binary module '%s'", name)
	outputPath := path.Join(config.BaseOutputDir, "docs", "my-docs.txt")
	var buildInputs []string
	inputErrors := false

	for _, src := range gd.properties.Srcs {
		if matches, err := ctx.GlobWithDeps(src, gd.properties.TestSrcs); err == nil {
			buildInputs = append(buildInputs, matches...)
		} else {
			ctx.PropertyErrorf("godocSrcs", "Cannot resolve files that match pattern %s", src)
			inputErrors = true
		}
	}

	if inputErrors {
		return
	}
	if gd.properties.VendorFirst {
		vendorDirPath := path.Join(ctx.ModuleDir(), "vendor")
		ctx.Build(pctx, blueprint.BuildParams{
			Description: fmt.Sprintf("Vendor dependencies of %s", name),
			Rule:        goVendor,
			Outputs:     []string{vendorDirPath},
			Implicits:   []string{path.Join(ctx.ModuleDir(), "../go.mod")},
			Optional:    true,
			Args: map[string]string{
				"workDir": ctx.ModuleDir(),
				"name":    name,
			},
		})
		buildInputs = append(buildInputs, vendorDirPath)
	}

	ctx.Build(pctx, blueprint.BuildParams{
		Description: fmt.Sprintf("Build %s as Go binary", name),
		Rule:        goDoc,
		Outputs:     []string{outputPath},
		Implicits:   buildInputs,
		Args: map[string]string{
			"workDir":    ctx.ModuleDir(),
			"pkg":        gd.properties.Pkg,
			"htmlFile":   outputPath,
		},
	})

}

func SimpleDocFactory() (blueprint.Module, []interface{}) {
	DocType := &goDocModuleType{}
	return DocType, []interface{}{&DocType.SimpleName.Properties, &DocType.properties}
}
