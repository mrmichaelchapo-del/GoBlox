// server/setup.go
// GoBlox 🐿️ — Server Setup
// First Go file! Reads the app manifest and boots the server.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// AppNeeds mirrors the "appNeeds" section of package.json5
type AppNeeds struct {
	Schemas string `json:"schemas"`
	API     string `json:"api"`
	Blocks  string `json:"blocks"`
	Requires struct {
		Go    string `json:"go"`
		JSON5 bool   `json:"json5"`
	} `json:"requires"`
}

// Config mirrors the parts of the manifest we care about at startup
type Config struct {
	Name     string   `json:"name"`
	Version  string   `json:"version"`
	AppNeeds AppNeeds `json:"appNeeds"`
}

func main() {
	fmt.Println("🐿️  GoBlox server starting...")

	// For now, just prove the server boots.
	// Next step: read package.json5 (JSON5 — needs a parser!)
	cfg := Config{
		Name:    "GoBlox",
		Version: "0.1.0",
	}
	cfg.AppNeeds.Schemas = "./schemas.json"
	cfg.AppNeeds.API = "server/setup.go"
	cfg.AppNeeds.Blocks = "./blocks"

	fmt.Printf("📋 App:      %s v%s\n", cfg.Name, cfg.Version)
	fmt.Printf("📐 Schemas:  %s\n", cfg.AppNeeds.Schemas)
	fmt.Printf("🧱 Blocks:   %s\n", cfg.AppNeeds.Blocks)
	fmt.Printf("🐹 API:      %s\n", cfg.AppNeeds.API)

	// Check that the blocks folder actually exists
	if _, err := os.Stat(cfg.AppNeeds.Blocks); os.IsNotExist(err) {
		fmt.Println("🚨 blocks/ folder not found!")
		os.Exit(1)
	}
	fmt.Println("✅ blocks/ folder found!")

	// Quick JSON sanity check on schemas.json
	data, err := os.ReadFile(cfg.AppNeeds.Schemas)
	if err != nil {
		fmt.Printf("🚨 Could not read %s: %v\n", cfg.AppNeeds.Schemas, err)
		os.Exit(1)
	}

	var schema map[string]any
	if err := json.Unmarshal(data, &schema); err != nil {
		fmt.Printf("🚨 schemas.json is not valid JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ schemas.json is valid JSON!")

	fmt.Println("🎉 GoBlox server ready! 🐿️")
}