package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var httpMethods = map[string]struct{}{
	"get":     {},
	"put":     {},
	"post":    {},
	"delete":  {},
	"options": {},
	"head":    {},
	"patch":   {},
	"trace":   {},
}

func main() {
	input := flag.String("input", "docs/openapi.yaml", "input openapi yaml path")
	output := flag.String("output", "", "output openapi yaml path")
	flag.Parse()

	if *output == "" {
		*output = *input
	}

	if err := process(*input, *output); err != nil {
		fmt.Fprintf(os.Stderr, "openapi postprocess failed: %v\n", err)
		os.Exit(1)
	}
}

func process(input, output string) error {
	content, err := os.ReadFile(input)
	if err != nil {
		return err
	}

	var doc yaml.Node
	if err := yaml.Unmarshal(content, &doc); err != nil {
		return err
	}

	if len(doc.Content) == 0 {
		return fmt.Errorf("empty openapi document")
	}

	root := doc.Content[0]
	if root.Kind != yaml.MappingNode {
		return fmt.Errorf("invalid openapi document root kind: %d", root.Kind)
	}

	components := ensureMappingValue(root, "components")
	securitySchemes := ensureMappingValue(components, "securitySchemes")
	setMapValue(securitySchemes, "BearerAuth", bearerSchemeNode())
	setMapValue(root, "security", bearerSecurityNode())

	paths := findMapValue(root, "paths")
	if paths != nil && paths.Kind == yaml.MappingNode {
		for index := 0; index < len(paths.Content); index += 2 {
			pathItem := paths.Content[index+1]
			if pathItem.Kind != yaml.MappingNode {
				continue
			}

			for itemIndex := 0; itemIndex < len(pathItem.Content); itemIndex += 2 {
				methodNode := pathItem.Content[itemIndex]
				operationNode := pathItem.Content[itemIndex+1]
				if _, ok := httpMethods[methodNode.Value]; !ok || operationNode.Kind != yaml.MappingNode {
					continue
				}
				setMapValue(operationNode, "security", bearerSecurityNode())
			}
		}
	}

	var buffer bytes.Buffer
	encoder := yaml.NewEncoder(&buffer)
	encoder.SetIndent(2)
	if err := encoder.Encode(&doc); err != nil {
		return err
	}
	if err := encoder.Close(); err != nil {
		return err
	}

	return os.WriteFile(output, buffer.Bytes(), 0644)
}

func ensureMappingValue(node *yaml.Node, key string) *yaml.Node {
	if value := findMapValue(node, key); value != nil {
		if value.Kind != yaml.MappingNode {
			*value = *newMappingNode()
		}
		return value
	}

	value := newMappingNode()
	node.Content = append(node.Content, newStringNode(key), value)
	return value
}

func findMapValue(node *yaml.Node, key string) *yaml.Node {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil
	}

	for index := 0; index < len(node.Content); index += 2 {
		if node.Content[index].Value == key {
			return node.Content[index+1]
		}
	}

	return nil
}

func setMapValue(node *yaml.Node, key string, value *yaml.Node) {
	if node.Kind != yaml.MappingNode {
		node.Kind = yaml.MappingNode
		node.Tag = "!!map"
		node.Content = nil
	}

	for index := 0; index < len(node.Content); index += 2 {
		if node.Content[index].Value == key {
			node.Content[index+1] = value
			return
		}
	}

	node.Content = append(node.Content, newStringNode(key), value)
}

func bearerSchemeNode() *yaml.Node {
	node := newMappingNode()
	setMapValue(node, "type", newStringNode("http"))
	setMapValue(node, "scheme", newStringNode("bearer"))
	setMapValue(node, "bearerFormat", newStringNode("JWT"))
	return node
}

func bearerSecurityNode() *yaml.Node {
	schemeList := newSequenceNode()
	securityItem := newMappingNode()
	setMapValue(securityItem, "BearerAuth", newSequenceNode())
	schemeList.Content = append(schemeList.Content, securityItem)
	return schemeList
}

func newMappingNode() *yaml.Node {
	return &yaml.Node{Kind: yaml.MappingNode, Tag: "!!map"}
}

func newSequenceNode() *yaml.Node {
	return &yaml.Node{Kind: yaml.SequenceNode, Tag: "!!seq"}
}

func newStringNode(value string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: value}
}
