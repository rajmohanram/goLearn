package main

import (
  "fmt"
  "io/ioutil"
  "gopkg.in/yaml.v2" // Import YAML parsing library
)

// ParseYAMLFile parses a YAML file and prints its contents to stdout
func ParseYAMLFile(filePath string) error {
  // Read the YAML file content
  data, err := ioutil.ReadAll(filePath)
  if err != nil {
    return fmt.Errorf("error reading YAML file: %w", err)
  }

  // Create an empty interface to hold the parsed data
  var parsedData interface{}

  // Unmarshal the YAML content into the interface
  err = yaml.Unmarshal(data, &parsedData)
  if err != nil {
    return fmt.Errorf("error unmarshaling YAML data: %w", err)
  }

  // Print the parsed data (may need type assertion for specific use)
  fmt.Println("Parsed YAML data:")
  fmt.Printf("%v\n", parsedData)

  return nil
}

func main() {
  // Replace "path/to/your/file.yaml" with the actual file path
  err := ParseYAMLFile("config.yaml")
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  fmt.Println("Successfully parsed YAML file!")
}
