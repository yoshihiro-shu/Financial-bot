package testutils_test

import (
	"strings"
	"testing"

	"github.com/yoshihiro-shu/financial-bot/internal/testutils"
)

func TestGetProjectRoot(t *testing.T) {
	root, err := testutils.GetProjectRoot()
	if err != nil {
		t.Errorf("error is %s", err)
	}

	if !strings.HasSuffix(root, "financial-bot") {
		t.Errorf("root is %s", root)
	}
}
