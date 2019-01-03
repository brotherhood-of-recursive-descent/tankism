package internal_test

import (
	"os"
	"testing"

	"github.com/co0p/tankism/internal"
)

func Test_NewAssetManager_SHOULD_return_initialized_manager_WHEN_path_exists(t *testing.T) {

	path := "../test/assets"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatal("failed loading test data: ", err)
	}

	m, err := internal.NewAssetManager(path)

	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	if ok := m.HasAsset("tank_dark.png"); !ok {
		t.Errorf("expected assetManager to have loaded tank_dark.png")
	}
}

func Test_NewAssetManager_SHOULD_return_error_WHEN_path_does_not_exists(t *testing.T) {

	path := "../test/does/not/exists"
	_, err := internal.NewAssetManager(path)

	if err == nil {
		t.Errorf("expected err not to be nil, got nil")
	}

}
