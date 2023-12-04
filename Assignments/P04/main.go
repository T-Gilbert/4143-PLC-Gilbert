package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
  "path/filepath"
  "sync"
  "time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
  for _, url := range urls {
    filename := generateFilename(url)
    err := downloadImage(url, filename)
    if err != nil {
      fmt.Printf("Error downloading %s: %v\n", url, err)
    }
  }
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
  var wg sync.WaitGroup
  errors := make(chan error, len(urls))

  for _, url := range urls {
    wg.Add(1)
    go func(url string) {
      defer wg.Done()
      filename := generateFilename(url)
      err := downloadImage(url, filename)
      if err != nil {
        errors <- fmt.Errorf("Error downloading %s: %v", url, err)
      }
    }(url)
  }

  go func() {
    wg.Wait()
    close(errors)
  }()

  for err := range errors {
    fmt.Println(err)
  }
}

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
  response, err := http.Get(url)
  if err != nil {
    return err
  }
  defer response.Body.Close()

  if response.StatusCode != http.StatusOK {
    return fmt.Errorf("HTTP error: %s", response.Status)
  }

  file, err := os.Create(filename)
  if err != nil {
    return err
  }
  defer file.Close()

  _, err = io.Copy(file, response.Body)
  if err != nil {
    return err
  }

  fmt.Printf("Downloaded: %s\n", filename)
  return nil
}

// Helper function to generate a unique filename based on the URL.
func generateFilename(url string) string {
  ext := filepath.Ext(url)
  return fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
}

func main() {
  urls := []string{
    "https://images.unsplash.com/photo-1640092256249-7e5089c74f30?q=80&w=1587&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    "https://images.unsplash.com/photo-1640580410941-cc7b71039148?q=80&w=1548&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    "https://plus.unsplash.com/premium_photo-1666433656515-77386ea16d5a?q=80&w=1587&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    "https://images.unsplash.com/photo-1700939931739-3aa4e97f566b?q=80&w=1587&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    "https://images.unsplash.com/photo-1700463108499-1f910f2c10ee?q=80&w=1635&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    // Add more image URLs
  }

  // Sequential download
  start := time.Now()
  downloadImagesSequential(urls)
  fmt.Printf("Sequential download took: %v\n", time.Since(start))

  // Concurrent download
  start = time.Now()
  downloadImagesConcurrent(urls)
  fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}

