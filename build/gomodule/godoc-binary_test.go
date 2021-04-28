package gomodule

import (
	"bytes"
	"github.com/google/blueprint"
	"github.com/roman-mazur/bood"
	"strings"
	"testing"
)

func TestSimpleDocFactory(t *testing.T) {
	ctx := blueprint.NewContext()

	ctx.MockFileSystem(map[string][]byte{
		"Blueprints": []byte(`
			go_doc {
 			name: "test-out",
			pkg: ".",
			srcs: ["test-src.go"],
			}
		`),
		"test-src.go": nil,
	})

	ctx.RegisterModuleType("go_doc", SimpleDocFactory)

	cfg := bood.NewConfig()

	_, errs := ctx.ParseBlueprintsFiles(".", cfg)
	if len(errs) != 0 {
		t.Fatalf("Syntax errors in the test blueprint file: %s", errs)
	}

	_, errs = ctx.PrepareBuildActions(cfg)
	if len(errs) != 0 {
		t.Errorf("Unexpected errors while preparing doc actions: %s", errs)
	}
	buffer := new(bytes.Buffer)
	if err := ctx.WriteBuildFile(buffer); err != nil {
		t.Errorf("Error writing ninja file: %s", err)
	} else {
		text := buffer.String()
		t.Logf("Gennerated ninja doc file:\n%s", text)
		if !strings.Contains(text, "out/docs/my-docs.txt: ") {
			t.Errorf("Generated ninja file does not have build of the doc module")
		}
		if !strings.Contains(text, " test-src.go") {
			t.Errorf("Generated ninja file does not have source dependency")
		}
		if !strings.Contains(text, "build out/docs/my-docs.txt: g.gomodule.doc") {
			t.Errorf("Generated ninja file does not have test doc rule")
		}
	}
}

func TestSimpleDocFactoryOptional(t *testing.T) {
	ctx := blueprint.NewContext()

	ctx.MockFileSystem(map[string][]byte{
		"Blueprints": []byte(`
			go_doc {
 			name: "test-out",
			pkg: ".",
			srcs: ["test-src.go"],
			optional: true
			}
		`),
		"test-src.go": nil,
	})

	ctx.RegisterModuleType("go_doc", SimpleDocFactory)

	cfg := bood.NewConfig()

	_, errs := ctx.ParseBlueprintsFiles(".", cfg)
	if len(errs) != 0 {
		t.Fatalf("Syntax errors in the test blueprint file: %s", errs)
	}

	_, errs = ctx.PrepareBuildActions(cfg)
	if len(errs) != 0 {
		t.Errorf("Unexpected errors while preparing doc actions: %s", errs)
	}
	buffer := new(bytes.Buffer)
	if err := ctx.WriteBuildFile(buffer); err != nil {
		t.Errorf("Error writing ninja file: %s", err)
	} else {
		text := buffer.String()
		t.Logf("Gennerated ninja doc file:\n%s", text)
		if strings.Contains(text, "default out/docs/my-docs.txt") {
			t.Errorf("Generated ninja file with default statement")
		}
	}
}
