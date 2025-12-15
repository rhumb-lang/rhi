package loader

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// HashShelf calculates a deterministic SHA256 hash for a file or a directory (shelf).
// If path is a file, it hashes the content.
// If path is a directory, it hashes all files within it (recursively? or just .rh?),
// sorted by path, to ensure stability.
// For now, we'll assume it hashes all relevant files in the directory structure.
func HashShelf(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return hashFile(path)
	}

	// Directory hashing
	h := sha256.New()
	
	// Walk the directory deterministically
	// We use a file system walk
	var files []string
	err = filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		// Filter? For a "Shelf", maybe we only care about .rh files and resources?
		// But HashShelf implies integrity of the whole thing.
		// Let's ignore hidden files/directories (starting with .)
		if strings.HasPrefix(d.Name(), ".") {
			return nil
		}
		
		rel, _ := filepath.Rel(path, p)
		files = append(files, rel)
		return nil
	})
	if err != nil {
		return "", err
	}

	sort.Strings(files)

	for _, file := range files {
		fullPath := filepath.Join(path, file)
		
		// Hash the filename (relative) to detect structure changes
		h.Write([]byte(file))
		
		f, err := os.Open(fullPath)
		if err != nil {
			return "", err
		}
		
		// Hash content
		if _, err := io.Copy(h, f); err != nil {
			f.Close()
			return "", err
		}
		f.Close()
	}

	return "sha256:" + hex.EncodeToString(h.Sum(nil)), nil
}

func hashFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return "sha256:" + hex.EncodeToString(h.Sum(nil)), nil
}
