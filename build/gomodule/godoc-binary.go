package gomodule

import (
	"fmt"
	"github.com/google/blueprint"
	"github.com/roman-mazur/bood"
	"path"
)

var (
	goDoc = pctx.StaticRule("doc", blueprint.RuleParams{
		Command:          "cd $workDir && go doc $pkg > $htmlFile",
		Description:      "godoc $pkg",
	}, "workDir", "pkg", "htmlFile")
)

type goDocModuleType struct {
	blueprint.SimpleName

	properties struct {
		Pkg         string
		Srcs        []string
		Deps        []string
	}
}

func (gt *goDocModuleType) DynamicDependencies(blueprint.DynamicDependerModuleContext) []string {
	return gt.properties.Deps
}

func (gt *goDocModuleType) GenerateBuildActions(ctx blueprint.ModuleContext) {
	name := ctx.ModuleName()
	config := bood.ExtractConfig(ctx)
	config.Info.Printf("Adding build & test actions for go binary module '%s'", name)

	outputPath := path.Join(config.BaseOutputDir, "docs", "my-docs.html")

	var buildInputs []string
	inputErrors := false

	for _, src := range gt.properties.Srcs {
		if matches, err := ctx.GlobWithDeps(src, nil); err == nil {
			buildInputs = append(buildInputs, matches...)
		} else {
			ctx.PropertyErrorf("godocSrcs", "Cannot resolve files that match pattern %s", src)
			inputErrors = true
		}
	}

	if inputErrors {
		return
	}

	//if gt.properties.VendorFirst {
	//	vendorDirPath := path.Join(ctx.ModuleDir(), "vendor")
	//	ctx.Build(pctx, blueprint.BuildParams{
	//		Description: fmt.Sprintf("Vendor dependencies of %s", name),
	//		Rule:        goVendor,
	//		Outputs:     []string{vendorDirPath},
	//		Implicits:   []string{path.Join(ctx.ModuleDir(), "../go.mod")},
	//		Optional:    true,
	//		Args: map[string]string{
	//			"workDir": ctx.ModuleDir(),
	//			"name":    name,
	//		},
	//	})
	//	buildInputs = append(buildInputs, vendorDirPath)
	//	testInputs = append(testInputs, vendorDirPath)
	//}

	ctx.Build(pctx, blueprint.BuildParams{
		Description: fmt.Sprintf("Build %s as Go binary", name),
		Rule:        goDoc,
		Outputs:     []string{outputPath},
		Implicits:   buildInputs,
		Args: map[string]string{
			"htmlFile":   outputPath,
			"workDir":    ctx.ModuleDir(),
			"pkg":        gt.properties.Pkg,
		},
	})

}

func SimpleDocFactory() (blueprint.Module, []interface{}) {
	DocType := &goDocModuleType{}
	return DocType, []interface{}{&DocType.SimpleName.Properties, &DocType.properties}
}
