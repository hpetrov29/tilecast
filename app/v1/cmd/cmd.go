package cmd

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func ValidateCoordinates(coord string) (float64, float64, error) {
	coord = strings.ReplaceAll(coord, " ", "")
	parts := strings.Split(coord, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid coordinate format, expected 'latitude,longitude'")
	}

	// Parse latitude and longitude as float64
	lat, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid latitude: %v", err)
	}

	lng, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid longitude: %v", err)
	}

	if lat < -90 || lat > 90 {
		return 0, 0, fmt.Errorf("latitude out of range: must be between -90 and 90")
	}
	
	if lng < -180 || lng > 180 {
		return 0, 0, fmt.Errorf("longitude out of range: must be between -180 and 180")
	}

	return lat, lng, nil
}

func Main() {
	topLeft := flag.String("top-left", "", "Top-left coordinates (latitude,longitude)")
	bottomRight := flag.String("bottom-right", "", "Bottom-right coordinates (latitude,longitude)")

	flag.Parse()

	if *topLeft == "" || *bottomRight == "" {
		log.Fatal("Both top-left and bottom-right coordinates must be provided.")
	}

	topLat, topLng, err := ValidateCoordinates(*topLeft)
	if err != nil {
		log.Fatalf("Invalid top-left coordinates: %v\n", err)
	}

	bottomLat, bottomLng, err := ValidateCoordinates(*bottomRight)
	if err != nil {
		log.Fatalf("Invalid bottom-right coordinates: %v\n", err)
	}

	fmt.Printf("Fetching mbtiles for the region:\n")
	fmt.Printf("Top-left: Latitude: %f, Longitude: %f\n", topLat, topLng)
	fmt.Printf("Bottom-right: Latitude: %f, Longitude: %f\n", bottomLat, bottomLng)

	// Here, you would call other functions to fetch the tiles and serve them

}