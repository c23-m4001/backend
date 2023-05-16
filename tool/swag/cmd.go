package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/swaggo/swag"
	"github.com/swaggo/swag/format"
	"github.com/swaggo/swag/gen"
)

const (
	searchDirFlag         = "dir"
	excludeFlag           = "exclude"
	generalInfoFlag       = "generalInfo"
	propertyStrategyFlag  = "propertyStrategy"
	outputFlag            = "output"
	outputTypesFlag       = "outputTypes"
	parseVendorFlag       = "parseVendor"
	parseDependencyFlag   = "parseDependency"
	markdownFilesFlag     = "markdownFiles"
	codeExampleFilesFlag  = "codeExampleFiles"
	parseInternalFlag     = "parseInternal"
	generatedTimeFlag     = "generatedTime"
	requiredByDefaultFlag = "requiredByDefault"
	parseDepthFlag        = "parseDepth"
	instanceNameFlag      = "instanceName"
	overridesFileFlag     = "overridesFile"
	parseGoListFlag       = "parseGoList"
	quietFlag             = "quiet"
	tagsFlag              = "tags"
)

func newInitCommand() *cobra.Command {
	var (
		quite             bool
		generalInfo       string
		searchDir         string
		exclude           string
		propertyStrategy  string
		output            string
		outputTypes       string
		parseVendor       bool
		parseDependency   bool
		markdownFiles     string
		codeExampleFiles  string
		parseInternal     bool
		generatedTime     bool
		parseDepth        int
		requiredByDefault bool
		instanceName      string
		overridesFile     string
		parseGoList       bool
		tags              string
	)

	cmd := &cobra.Command{
		Use:     "init",
		Aliases: []string{"i"},
		Short:   "Create docs.go",
		Long:    "Create docs.go",
		RunE: func(cmd *cobra.Command, args []string) error {
			strategy := propertyStrategy

			switch strategy {
			case swag.CamelCase, swag.SnakeCase, swag.PascalCase:
			default:
				return fmt.Errorf("not supported %s propertyStrategy", strategy)
			}

			outputTypes := strings.Split(outputTypes, ",")
			if len(outputTypes) == 0 {
				return fmt.Errorf("no output types specified")
			}
			logger := log.New(os.Stdout, "", log.LstdFlags)
			if quite {
				logger = log.New(io.Discard, "", log.LstdFlags)
			}

			return New().Build(&Config{
				SearchDir:           searchDir,
				Excludes:            exclude,
				MainAPIFile:         generalInfo,
				PropNamingStrategy:  strategy,
				OutputDir:           output,
				OutputTypes:         outputTypes,
				ParseVendor:         parseVendor,
				ParseDependency:     parseDependency,
				MarkdownFilesDir:    markdownFiles,
				ParseInternal:       parseInternal,
				GeneratedTime:       generatedTime,
				RequiredByDefault:   requiredByDefault,
				CodeExampleFilesDir: codeExampleFiles,
				ParseDepth:          parseDepth,
				InstanceName:        instanceName,
				OverridesFile:       overridesFile,
				ParseGoList:         parseGoList,
				Tags:                tags,
				Debugger:            logger,
			})
		},
	}

	cmd.Flags().BoolVarP(
		&quite,
		quietFlag,
		"q",
		false,
		"Make the logger quite.",
	)
	cmd.Flags().StringVarP(
		&generalInfo,
		generalInfoFlag,
		"g",
		"main.go",
		"Go file path in which 'swagger general API Info' is written",
	)
	cmd.Flags().StringVarP(
		&searchDir,
		searchDirFlag,
		"d",
		"./",
		"Directories you want to parse,comma separated and general-info file must be in the first one",
	)
	cmd.Flags().StringVar(
		&exclude,
		excludeFlag,
		"",
		"Exclude directories and files when searching, comma separated",
	)
	cmd.Flags().StringVarP(
		&propertyStrategy,
		propertyStrategyFlag,
		"p",
		swag.CamelCase,
		"Property Naming Strategy like "+swag.SnakeCase+","+swag.CamelCase+","+swag.PascalCase,
	)
	cmd.Flags().StringVarP(
		&output,
		outputFlag,
		"o",
		"./docs",
		"Output directory for all the generated files(swagger.json, swagger.yaml and docs.go)",
	)
	cmd.Flags().StringVar(
		&outputTypes,
		outputTypesFlag,
		"go,json,yaml",
		"Output types of generated files (docs.go, swagger.json, swagger.yaml) like go,json,yaml",
	)
	cmd.Flags().BoolVar(
		&parseVendor,
		parseVendorFlag,
		false,
		"Parse go files in 'vendor' folder, disabled by default",
	)
	cmd.Flags().BoolVar(
		&parseDependency,
		parseDependencyFlag,
		false,
		"Parse go files inside dependency folder, disabled by default",
	)
	cmd.Flags().StringVar(
		&markdownFiles,
		markdownFilesFlag,
		"",
		"Parse folder containing markdown files to use as description, disabled by default",
	)
	cmd.Flags().StringVar(
		&codeExampleFiles,
		codeExampleFilesFlag,
		"",
		"Parse folder containing code example files to use for the x-codeSamples extension, disabled by default",
	)
	cmd.Flags().BoolVar(
		&parseInternal,
		parseInternalFlag,
		false,
		"Parse go files in internal packages, disabled by default",
	)
	cmd.Flags().BoolVar(
		&generatedTime,
		generatedTimeFlag,
		false,
		"Generate timestamp at the top of docs.go, disabled by default",
	)
	cmd.Flags().IntVar(
		&parseDepth,
		parseDepthFlag,
		100,
		"Dependency parse depth",
	)
	cmd.Flags().BoolVar(
		&requiredByDefault,
		requiredByDefaultFlag,
		false,
		"Set validation required for all fields by default",
	)
	cmd.Flags().StringVar(
		&instanceName,
		instanceNameFlag,
		"",
		"This parameter can be used to name different swagger document instances. It is optional.",
	)
	cmd.Flags().StringVar(
		&overridesFile,
		overridesFileFlag,
		gen.DefaultOverridesFile,
		"File to read global type overrides from.",
	)
	cmd.Flags().BoolVar(
		&parseGoList,
		parseGoListFlag,
		true,
		"Parse dependency via 'go list'",
	)
	cmd.Flags().StringVarP(
		&tags,
		tagsFlag,
		"t",
		"",
		"A comma-separated list of tags to filter the APIs for which the documentation is generated.Special case if the tag is prefixed with the '!' character then the APIs with that tag will be excluded",
	)

	return cmd
}

func newFmtCommand() *cobra.Command {
	var (
		searchDir   string
		exclude     string
		generalInfo string
	)

	cmd := &cobra.Command{
		Use:     "fmt",
		Aliases: []string{"f"},
		Short:   "format swag comments",
		Long:    "format swag comments",
		RunE: func(cmd *cobra.Command, args []string) error {
			return format.New().Build(&format.Config{
				SearchDir: searchDir,
				Excludes:  exclude,
				MainFile:  generalInfo,
			})
		},
	}

	cmd.Flags().StringVarP(
		&searchDir,
		searchDirFlag,
		"d",
		"./",
		"Directories you want to parse,comma separated and general-info file must be in the first one",
	)
	cmd.Flags().StringVar(
		&exclude,
		excludeFlag,
		"",
		"Exclude directories and files when searching, comma separated",
	)
	cmd.Flags().StringVarP(
		&generalInfo,
		generalInfoFlag,
		"g",
		"main.go",
		"Go file path in which 'swagger general API Info' is written",
	)

	return cmd
}

func main() {
	app := &cobra.Command{
		Use:     "swag",
		Version: swag.Version,
		Short:   "Automatically generate RESTful API documentation with Swagger 2.0 for Go.",
		Long:    "Automatically generate RESTful API documentation with Swagger 2.0 for Go.",
		Run:     func(cmd *cobra.Command, args []string) {},
	}

	app.AddCommand(
		newInitCommand(),
		newFmtCommand(),
	)

	if err := app.Execute(); err != nil {
		log.Fatal(err)
	}
}
