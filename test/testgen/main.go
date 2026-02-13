package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LanguageConfig struct {
	OutputDir         string `json:"output_dir"`
	TemplateDir       string `json:"template_dir"`
	SourceDir         string `json:"source_dir,omitempty"`
	CreateMissingDirs bool   `json:"create_missing_dirs"`
}

type Config struct {
	DataPath  string                    `json:"data_path"`
	Languages map[string]LanguageConfig `json:"languages"`
}

func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	// Resolve paths relative to config file location
	configDir := filepath.Dir(path)
	// Get absolute path of config file directory to deal with relative paths correctly
	absConfigDir, err := filepath.Abs(configDir)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path of config file directory: %w", err)
	}

	resolve := func(p string) string {
		if p == "" {
			return ""
		}
		if filepath.IsAbs(p) {
			return p
		}
		return filepath.Join(absConfigDir, p)
	}

	cfg.DataPath = resolve(cfg.DataPath)
	for lang, langCfg := range cfg.Languages {
		langCfg.OutputDir = resolve(langCfg.OutputDir)
		langCfg.TemplateDir = resolve(langCfg.TemplateDir)
		langCfg.SourceDir = resolve(langCfg.SourceDir)
		cfg.Languages[lang] = langCfg
	}

	return &cfg, nil
}

func main() {
	configPath := flag.String("config", "config.json", "Path to the configuration JSON file")
	dryRun := flag.Bool("dry-run", false, "Log actions without writing files")
	check := flag.Bool("check", false, "Fail if generated files are not up to date; no writes")
	flag.Parse()

	if *dryRun && *check {
		fmt.Fprintln(os.Stderr, "Error: --dry-run and --check are mutually exclusive")
		os.Exit(1)
	}

	cfg, err := loadConfig(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	if cfg.DataPath == "" {
		fmt.Fprintln(os.Stderr, "Error: data_path is required in config")
		os.Exit(1)
	}

	data, err := LoadData(cfg.DataPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading data: %v\n", err)
		os.Exit(1)
	}

	if err := ValidateData(*data); err != nil {
		fmt.Fprintf(os.Stderr, "Error validating data: %v\n", err)
		os.Exit(1)
	}

	baseOpts := GenerationOptions{DryRun: *dryRun, CheckOnly: *check}

	for langName, langCfg := range cfg.Languages {
		generator, ok := generators[langName]
		if !ok {
			fmt.Fprintf(os.Stderr, "Warning: No generator found for language '%s'\n", langName)
			continue
		}

		fmt.Printf("=== Generating %s tests ===\n", strings.Title(langName))
		opts := baseOpts
		opts.CreateMissingDirs = langCfg.CreateMissingDirs
		if err := GenerateGeneric(generator, langName, *data, langCfg.TemplateDir, langCfg.OutputDir, langCfg.SourceDir, opts); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating %s tests: %v\n", langName, err)
			os.Exit(1)
		}
	}

	fmt.Println("Done.")
}
