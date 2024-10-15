package backend_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/docker/compose/v2/pkg/api"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

type BackendTestSuite struct {
	suite.Suite
	context              context.Context
	Success              bool
	SuccessMap           map[string]bool
	CumulativeSuccessMap map[string]bool
}

func (suite *BackendTestSuite) SetupSuite() {
	suite.context = context.Background()
	suite.Success = true
	suite.SuccessMap = make(map[string]bool)
	suite.CumulativeSuccessMap = make(map[string]bool)

	identifier := tc.StackIdentifier("integration_test")
	compose, err := tc.NewDockerComposeWith(tc.WithStackFiles("../deployment/docker-compose/docker-compose.yml"), identifier)
	require.NoError(suite.T(), err, "Error creating docker compose")

	suite.T().Cleanup(func() {
		require.NoError(suite.T(), compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	ctx, cancel := context.WithCancel(context.Background())
	suite.T().Cleanup(cancel)

	require.NoError(suite.T(), compose.Up(ctx, tc.WithRecreate(api.RecreateNever), tc.Wait(true)), "compose.Up()")
}

func (suite *BackendTestSuite) TearDownSuite() {
	// need to write the test results to a file
}

func (suite *BackendTestSuite) TestSignUp() {
	signUpRequest := map[string]string{
		"username": "testuser",
		"password": "testpassword",
		"role":     "user",
	}
	requestBody, err := json.Marshal(signUpRequest)
	require.NoError(suite.T(), err, "Error marshaling signup request")

	resp, err := http.Post("http://localhost:3000/api/v1/signup", "application/json", bytes.NewBuffer(requestBody))
	require.NoError(suite.T(), err, "Error making signup request")
	defer resp.Body.Close()

	require.Equal(suite.T(), http.StatusCreated, resp.StatusCode, "Expected status code 201 Created")

	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	require.NoError(suite.T(), err, "Error decoding signup response")

	require.Equal(suite.T(), "User created successfully", responseBody["message"], "Expected success message")
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(BackendTestSuite))
}
