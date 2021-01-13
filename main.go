// Cache Push step keeps the project's cache in sync with the project's current state based on the defined files to be cached and ignored.
//
// Files to be cached are described by a path and an optional descriptor file path.
// Files to be cached can be referred by direct file path while multiple files can be selected by referring the container directory.
// Optional indicator represents a files, based on which the step synchronizes the given file(s).
// Syntax: file/path/to/cache, dir/to/cache, file/path/to/cache -> based/on/this/file, dir/to/cache -> based/on/this/file
//
// Ignore items are used to ignore certain file(s) from a directory to be cached or to mark that certain file(s) not relevant in cache synchronization.
// Syntax: not/relevant/file/or/pattern, !file/or/pattern/to/remove/from/cache
package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bitrise-io/go-utils/log"
)

const (
	cacheInfoFilePath = "/tmp/cache-info.json"
	cacheArchivePath  = "/tmp/cache-archive.tar"
	stackVersionsPath = "/tmp/archive_info.json"
	stepID            = "cache-push"
)

func logErrorfAndExit(format string, args ...interface{}) {
	log.Errorf(format, args...)
	os.Exit(1)
}

func main() {
	stepStartedAt := time.Now()

	configs, err := ParseConfig()
	if err != nil {
		logErrorfAndExit(err.Error())
	}

	configs.Print()
	fmt.Println()

	log.SetEnableDebugLog(configs.DebugMode)

	// Cleaning paths
	startTime := time.Now()

	log.Infof("Cleaning paths")

	pathToIndicatorPath := parseIncludeList(strings.Split(configs.Paths, "\n"))
	if len(pathToIndicatorPath) == 0 {
		log.Warnf("No path to cache, skip caching...")
		os.Exit(0)
	}

	pathToIndicatorPath, err = normalizeIndicatorByPath(pathToIndicatorPath)
	if err != nil {
		logErrorfAndExit("Failed to parse include list: %s", err)
	}

	excludeByPattern := parseIgnoreList(strings.Split(configs.IgnoredPaths, "\n"))
	excludeByPattern, err = normalizeExcludeByPattern(excludeByPattern)
	if err != nil {
		logErrorfAndExit("Failed to parse ignore list: %s", err)
	}

	pathToIndicatorPath = interleave(pathToIndicatorPath, excludeByPattern)

	log.Donef("Done in %s\n", time.Since(startTime))

	if len(pathToIndicatorPath) == 0 {
		log.Warnf("No path to cache, skip caching...")
		os.Exit(0)
	}

	// Check previous cache
	startTime = time.Now()

	log.Infof("Checking previous cache status")

	log.Infof("Reading Cache Descriptor")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	prevDescriptor, err := readCacheDescriptor(cacheInfoFilePath)
	if err != nil {
		logErrorfAndExit("Failed to read previous cache descriptor: %s", err)
	}

	if prevDescriptor != nil {
		log.Printf("Previous cache info found at: %s", cacheInfoFilePath)
	} else {
		log.Printf("No previous cache info found")
	}

	log.Donef("Creating Cache Descriptor")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	curDescriptor, err := cacheDescriptor(pathToIndicatorPath, ChangeIndicator(configs.FingerprintMethodID))
	if err != nil {
		logErrorfAndExit("Failed to create current cache descriptor: %s", err)
	}

	log.Donef("Done in %s\n", time.Since(startTime))

	// Checking file changes
	if prevDescriptor != nil {
		startTime = time.Now()

		log.Infof("Checking for file changes")

		logDebugPaths := func(paths []string) {
			for _, pth := range paths {
				log.Debugf("- %s", pth)
			}
		}

		log.Donef("Comparing previous descriptor with current descriptor")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")
		log.Infof("..")

		result := compare(prevDescriptor, curDescriptor)

		log.Warnf("%d files needs to be removed", len(result.removed))
		logDebugPaths(result.removed)
		log.Warnf("%d files has changed", len(result.changed))
		logDebugPaths(result.changed)
		log.Warnf("%d files added", len(result.added))
		logDebugPaths(result.added)
		log.Debugf("%d ignored files removed", len(result.removedIgnored))
		logDebugPaths(result.removedIgnored)
		log.Debugf("%d files did not change", len(result.matching))
		logDebugPaths(result.matching)
		log.Debugf("%d ignored files added", len(result.addedIgnored))
		logDebugPaths(result.addedIgnored)

		if result.hasChanges() {
			log.Donef("File changes found in %s\n", time.Since(startTime))
		} else {
			log.Donef("No files found in %s\n", time.Since(startTime))
			log.Printf("Total time: %s", time.Since(stepStartedAt))
			os.Exit(0)
		}
	}

	// Generate cache archive
	startTime = time.Now()

	log.Infof("Generating cache archive")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	archive, err := NewArchive(cacheArchivePath, configs.CompressArchive == "true")
	if err != nil {
		logErrorfAndExit("Failed to create archive: %s", err)
	}

	log.Donef("Getting Stack Version Data")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	stackData, err := stackVersionData(configs.StackID)
	if err != nil {
		logErrorfAndExit("Failed to get stack version info: %s", err)
	}

	log.Donef("Writing archive data to disk")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	// This is the first file written, to speed up reading it in subsequent builds
	if err = archive.writeData(stackData, stackVersionsPath); err != nil {
		logErrorfAndExit("Failed to write cache info to archive, error: %s", err)
	}

	log.Donef("Pouplating Archive")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	if err := archive.Write(pathToIndicatorPath); err != nil {
		logErrorfAndExit("Failed to populate archive: %s", err)
	}

	log.Donef("Writing header to archive")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	if err := archive.WriteHeader(curDescriptor, cacheInfoFilePath); err != nil {
		logErrorfAndExit("Failed to write archive header: %s", err)
	}

	log.Donef("Closing Archive")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	if err := archive.Close(); err != nil {
		logErrorfAndExit("Failed to close archive: %s", err)
	}

	log.Donef("Done in %s\n", time.Since(startTime))

	// Upload cache archive
	startTime = time.Now()

	log.Infof("Uploading cache archive")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")
	log.Infof("..")

	if err := uploadArchive(cacheArchivePath, configs.CacheAPIURL, configs.BuildSlug); err != nil {
		logErrorfAndExit("Failed to upload archive: %s", err)
	}
	log.Donef("Done in %s\n", time.Since(startTime))
	log.Donef("Total time: %s", time.Since(stepStartedAt))
}
