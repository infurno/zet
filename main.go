package main

import (
    "io/ioutil"
    "log"
    "os"
    "os/exec"
    "path/filepath"

    "gopkg.in/yaml.v2"
)

type Config struct {
    TemplatePath string `yaml:"templatePath"`
}

func main() {
    configPath := filepath.Join(os.Getenv("HOME"), ".config", "zet.yml")

    // Read the config file
    data, err := ioutil.ReadFile(configPath)
    if err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }

    // Parse the config file
    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        log.Fatalf("Error parsing config file: %v", err)
    }

    // Read the template content
    templateContent, err := ioutil.ReadFile(config.TemplatePath)
    if err != nil {
        log.Fatalf("Error reading template file: %v", err)
    }

    // Create a temporary file for the new document
    tmpFile, err := ioutil.TempFile("", "zet-*.md")
    if err != nil {
        log.Fatalf("Error creating temporary file: %v", err)
    }
    defer os.Remove(tmpFile.Name())

    // Write the template content to the temporary file
    if _, err := tmpFile.Write(templateContent); err != nil {
        log.Fatalf("Error writing to temporary file: %v", err)
    }
    tmpFile.Close()

    // Open the temporary file in nvim
    cmd := exec.Command("nvim", tmpFile.Name())
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err = cmd.Run()
    if err != nil {
        log.Fatalf("Error opening file in nvim: %v", err)
    }
}

