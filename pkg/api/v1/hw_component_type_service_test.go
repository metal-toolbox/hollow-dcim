package hollow_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	hollow "go.metalkube.net/hollow/pkg/api/v1"
)

func TestHWComponentTypeServiceCreate(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		hct := hollow.HardwareComponentType{Name: "unit-test"}
		jsonResponse := json.RawMessage([]byte(`{"message": "resource created", "uuid":"00000000-0000-0000-0000-000000001234"}`))

		c := mockClient(string(jsonResponse), respCode)
		res, err := c.HardwareComponentType.Create(ctx, hct)
		if !expectError {
			assert.Equal(t, "00000000-0000-0000-0000-000000001234", res.String())
		}

		return err
	})
}

func TestHWComponentTypeServiceList(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		hct := []hollow.HardwareComponentType{{UUID: uuid.New(), Name: "unit-test-1"}, {UUID: uuid.New(), Name: "unit-test-2"}}
		jsonResponse, err := json.Marshal(hct)
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, err := c.HardwareComponentType.List(ctx, nil)
		if !expectError {
			assert.ElementsMatch(t, hct, res)
		}

		return err
	})
}
