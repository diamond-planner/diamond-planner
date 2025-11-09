package dp

import (
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/diamond-planner/diamond-planner/bsm"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

// MockApp implements GameImportServiceApp for tests.
type MockApp struct {
	findFirstRecordByDataFunc    func(any, string, any) (*core.Record, error)
	FindRecordsByFilterFunc      func(collectionModelOrIdentifier any, filter string, sort string, limit int, offset int, params ...dbx.Params) ([]*core.Record, error)
	findCollectionByNameOrIdFunc func(string) (*core.Collection, error)
	saveFunc                     func(core.Model) error
}

func (m *MockApp) FindRecordsByFilter(collectionModelOrIdentifier any, filter string, sort string, limit int, offset int, params ...dbx.Params) ([]*core.Record, error) {
	return m.FindRecordsByFilterFunc(collectionModelOrIdentifier, filter, sort, limit, offset, params...)
}

func (m *MockApp) Logger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}

func (m *MockApp) ExpandRecord(record *core.Record, expands []string, optFetchFunc core.ExpandFetchFunc) map[string]error {
	// TODO: implement
	return nil
}

func (m *MockApp) FindFirstRecordByData(collectionModelOrIdentifier any, field string, value any) (*core.Record, error) {
	return m.findFirstRecordByDataFunc(collectionModelOrIdentifier, field, value)
}

func (m *MockApp) FindCollectionByNameOrId(nameOrId string) (*core.Collection, error) {
	return m.findCollectionByNameOrIdFunc(nameOrId)
}

func (m *MockApp) Save(record core.Model) error {
	return m.saveFunc(record)
}

func TestCreateOrUpdateField_NewLocation(t *testing.T) {
	mockApp := &MockApp{
		findFirstRecordByDataFunc: func(any, string, any) (*core.Record, error) {
			return nil, errors.New("not found")
		},
		findCollectionByNameOrIdFunc: func(_ string) (*core.Collection, error) {
			collection := core.NewCollection("möp", LocationCollection)
			return collection, nil
		},
		saveFunc: func(_ core.Model) error {
			return nil
		},
	}

	testTeamCollection := core.NewCollection("möp", "teams")
	service := GameImportService{App: mockApp}
	testField := bsm.Field{BSMID: 123, Name: "Test Field"}
	testTeam := core.NewRecord(testTeamCollection)

	location, err := service.createOrUpdateField(testTeam, testField)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if location == nil {
		t.Fatal("Expected location to be not nil")
	}
	if location.Name() != "Test Field" {
		t.Errorf("Expected name to be 'Test Field', got: %v", location.Name())
	}
	if location.BSMID() != 123 {
		t.Errorf("Expected BSMID '123', got: %v", location.BSMID())
	}
}

func TestCreateOrUpdateField_UpdateExistingLocation(t *testing.T) {
	testLocationCollection := core.NewCollection("möp", "locations")

	existingRecord := core.NewRecord(testLocationCollection)
	existingRecord.Set("name", "PreviousLocationName")
	existingRecord.Set("bsm_id", 123)

	mockApp := &MockApp{
		findFirstRecordByDataFunc: func(any, string, any) (*core.Record, error) {
			return existingRecord, nil
		},
		saveFunc: func(_ core.Model) error {
			return nil
		},
	}

	testTeamCollection := core.NewCollection("möp", "teams")
	service := GameImportService{App: mockApp}
	testField := bsm.Field{BSMID: 123, Name: "UpdatedLocationName"}
	testTeam := core.NewRecord(testTeamCollection)

	location, err := service.createOrUpdateField(testTeam, testField)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if location == nil {
		t.Fatal("Expected location to be not nil")
	}
	if location.BSMID() != 123 {
		t.Errorf("Expected BSMID '123', got: %v", location.BSMID())
	}
	if location.Name() != "UpdatedLocationName" {
		t.Errorf("Expected Name 'UpdatedLocationName', got: %v", location.Name())
	}
}

func TestCreateOrUpdateField_SaveError(t *testing.T) {
	mockApp := &MockApp{
		findFirstRecordByDataFunc: func(any, string, any) (*core.Record, error) {
			return nil, errors.New("not found")
		},
		findCollectionByNameOrIdFunc: func(_ string) (*core.Collection, error) {
			return &core.Collection{}, nil
		},
		saveFunc: func(_ core.Model) error {
			return errors.New("save failed")
		},
	}

	testTeamCollection := core.NewCollection("möp", "teams")
	service := GameImportService{App: mockApp}
	testField := bsm.Field{BSMID: 123, Name: "Failing Field"}
	testTeam := core.NewRecord(testTeamCollection)

	location, err := service.createOrUpdateField(testTeam, testField)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if location != nil {
		t.Errorf("Expected nil location, got: %+v", location)
	}
}
