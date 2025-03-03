package tests

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/hpetrov29/tilecast/app/v1/cmd"
)

// TestFlagParsing tests the flag parsing and validation for coordinates
func TestFlagParsing(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		expectedTopLat  float64
		expectedTopLng  float64
		expectedBotLat  float64
		expectedBotLng  float64
		expectErrorLat     bool
		expectErrorLng     bool
	}{
		{
			name:           "Valid coordinates",
			args:           []string{"cmd/tilecast/main.go", "--top-left", "37.7749,-122.4194", "--bottom-right", "37.8044,-122.2711"},
			expectedTopLat: 37.7749,
			expectedTopLng: -122.4194,
			expectedBotLat: 37.8044,
			expectedBotLng: -122.2711,
			expectErrorLat:    false,
			expectErrorLng:    false,
		},
		{
			name:           "Missing bottom-right coordinate",
			args:           []string{"cmd/tilecast/main.go", "--top-left", "37.7749,-122.4194"},
			expectedTopLat: 37.7749,
			expectedTopLng: -122.4194,
			expectedBotLat: 0,
			expectedBotLng: 0,
			expectErrorLat:    false,
			expectErrorLng:    true,
		},
		{
			name:           "Invalid coordinates",
			args:           []string{"cmd/tilecast/main.go", "--top-left", "100.0,-122.4194", "--bottom-right", "37.8044,-190.0"},
			expectedTopLat: 100.0,
			expectedTopLng: -122.4194,
			expectedBotLat: 37.8044,
			expectedBotLng: -190.0,
			expectErrorLat:    true,
			expectErrorLng:    true,
		},
	}

	// Iterate through test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset the flag set for each test to avoid "flag redefined" error
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

			// Define the flags within the test case
			topLeft := flag.String("top-left", "", "Top-left coordinates (latitude,longitude)")
			bottomRight := flag.String("bottom-right", "", "Bottom-right coordinates (latitude,longitude)")
			// Simulate setting the command-line arguments
			os.Args = tt.args

			// Call flag.Parse to parse the arguments
			flag.Parse()
			fmt.Println(*bottomRight)

			// Validate top-left coordinates
			topLat, topLng, err := cmd.ValidateCoordinates(*topLeft)
			if (err != nil) != tt.expectErrorLat {
				t.Errorf("Expected error %v, got %v", tt.expectErrorLat, err)
			}
			if !tt.expectErrorLat && (topLat != tt.expectedTopLat || topLng != tt.expectedTopLng) {
				t.Errorf("Expected top-left coordinates (%f, %f), got (%f, %f)", tt.expectedTopLat, tt.expectedTopLng, topLat, topLng)
			}

			// Validate bottom-right coordinates
			botLat, botLng, err := cmd.ValidateCoordinates(*bottomRight)
			if (err != nil) != tt.expectErrorLng {
				t.Errorf("Expected error %v, got %v", tt.expectErrorLng, "no error")
			}
			if !tt.expectErrorLng && (botLat != tt.expectedBotLat || botLng != tt.expectedBotLng) {
				t.Errorf("Expected bottom-right coordinates (%f, %f), got (%f, %f)", tt.expectedBotLat, tt.expectedBotLng, botLat, botLng)
			}
		})
	}
}
