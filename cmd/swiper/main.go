package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/onedusk/swiper/internal/batch"
	"github.com/onedusk/swiper/internal/config"
	"github.com/onedusk/swiper/internal/extractor"
	"github.com/onedusk/swiper/internal/scanner"
)

func main() {
	// Define command-line flags
	pdfFileFlag := flag.String("file", "", "Path to a single PDF file")
	inputDirFlag := flag.String("dir", "", "Directory containing PDF files to process")
	outputDirFlag := flag.String("output", "", "Directory to store extracted pages")
	processCountFlag := flag.Int("processes", 0, "Number of processes to use")
	configFlag := flag.String("config", "", "Path to a YAML configuration file")
	scanDirFlag := flag.String("scan", "", "Scan directory for PDF files and copy to pdf-docs")
	copyDirFlag := flag.String("copydir", "", "Directory to copy PDFs to (default: pdf-docs)")
	profileFlag := flag.String("profile", "", "Enable profiling (cpu or memory)")
	cacheFlag := flag.Bool("cache", false, "Enable result caching")
	benchmarkFlag := flag.Bool("benchmark", false, "Run in benchmark mode with detailed metrics")
	helpFlag := flag.Bool("help", false, "Prints help")
	flag.Parse()

	if *helpFlag {
		printHelp()
		os.Exit(0)
	}

	// Setup profiling if requested
	if *profileFlag == "cpu" {
		f, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal("Could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("Could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
		log.Println("CPU profiling enabled, output: cpu.prof")
	} else if *profileFlag == "mem" {
		defer func() {
			f, err := os.Create("mem.prof")
			if err != nil {
				log.Fatal("Could not create memory profile: ", err)
			}
			defer f.Close()
			runtime.GC()
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatal("Could not write memory profile: ", err)
			}
			log.Println("Memory profile written to mem.prof")
		}()
	}

	if *benchmarkFlag {
		log.Println("Running in benchmark mode with detailed metrics")
		runtime.GOMAXPROCS(runtime.NumCPU())
		log.Printf("Using %d CPU cores", runtime.NumCPU())
	}

	// Build configuration from all flags
	opts := &config.Options{
		PdfFile:      *pdfFileFlag,
		InputDir:     *inputDirFlag,
		OutputDir:    *outputDirFlag,
		ProcessCount: *processCountFlag,
		ScanDir:      *scanDirFlag,
		CopyDir:      *copyDirFlag,
		Profile:      *profileFlag,
		CacheResults: *cacheFlag,
	}

	// Load from config file if provided (CLI takes precedence)
	if *configFlag != "" {
		configOpts, err := config.LoadFromFile(*configFlag)
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}
		configOpts.Merge(opts)
		opts = configOpts
	}

	// Set defaults before validation
	opts.SetDefaults()

	// Validate configuration
	if err := opts.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	// Dispatch based on mode
	switch {
	case opts.ScanDir != "":
		runScanner(opts.ScanDir, opts.CopyDir)
	case opts.InputDir != "":
		runBatchProcessor(opts.InputDir, opts.OutputDir, opts.ProcessCount)
	default:
		runSingleExtraction(opts)
	}
}

func runSingleExtraction(opts *config.Options) {
	ext, err := extractor.New(opts.PdfFile, opts.OutputDir, opts.ProcessCount)
	if err != nil {
		log.Fatalf("Failed to initialize extractor: %v", err)
	}
	defer ext.Cleanup()

	if err := ext.ExtractPages(); err != nil {
		log.Fatalf("Extraction failed: %v", err)
	}
}

func runBatchProcessor(inputDir, outputDir string, processCount int) {
	processor, err := batch.New(inputDir, outputDir, processCount)
	if err != nil {
		log.Fatalf("Failed to initialize batch processor: %v", err)
	}

	if err := processor.ProcessAll(); err != nil {
		log.Fatalf("Batch processing failed: %v", err)
	}
}

func runScanner(scanDir, copyDir string) {
	scan, err := scanner.New(scanDir, copyDir)
	if err != nil {
		log.Fatalf("Failed to initialize scanner: %v", err)
	}

	if err := scan.ScanAndCopy(); err != nil {
		log.Fatalf("Scan failed: %v", err)
	}
}

func printHelp() {
	fmt.Println("PDF Tool - Extract pages from PDFs or scan/copy PDFs")
	fmt.Println("\nUsage:")
	fmt.Println("  Single PDF:    -file <pdf> [-output <dir>] [-processes <n>]")
	fmt.Println("  Batch mode:    -dir <directory> [-output <dir>] [-processes <n>]")
	fmt.Println("  Scan mode:     -scan <dir> [-copydir <dir>]")
	fmt.Println("\nPerformance Options:")
	fmt.Println("  -cache         Enable result caching for repeated extractions")
	fmt.Println("  -profile cpu   Enable CPU profiling")
	fmt.Println("  -profile mem   Enable memory profiling")
	fmt.Println("  -benchmark     Run with detailed performance metrics")
	fmt.Println("\nExamples:")
	fmt.Println("  ./swiper -file document.pdf")
	fmt.Println("  ./swiper -dir /path/to/pdfs -output extracted")
	fmt.Println("  ./swiper -scan . -copydir pdf-collection")
	fmt.Println("  ./swiper -file doc.pdf -profile cpu -benchmark")
	fmt.Println("\nFlags:")
	flag.PrintDefaults()
}