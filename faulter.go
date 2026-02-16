package leekcode

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	// 1. Parsowanie argumentów
	inPath := flag.String("in", "", "Ścieżka do pliku wejściowego YAML")
	outPath := flag.String("out", "", "Ścieżka do zapisu zepsutego YAML")
	seed := flag.Int64("seed", 42, "Ziarno dla generatora liczb losowych (dla determinizmu)")
	flag.Parse()

	if *inPath == "" || *outPath == "" {
		log.Fatal("Musisz podać argumenty -in oraz -out")
	}

	// 2. Inicjalizacja deterministycznego RNG
	// To jest kluczowe dla artykułu naukowego!
	rng := rand.New(rand.NewSource(*seed))

	// 3. Wczytanie pliku
	file, err := os.Open(*inPath)
	if err != nil {
		log.Fatalf("Błąd otwarcia pliku: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var docs []map[string]interface{}

	// 4. Parsowanie wielu dokumentów (obsługa separatora "---")
	for {
		var doc map[string]interface{}
		err := decoder.Decode(&doc)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Błąd dekodowania YAML: %v", err)
		}
		docs = append(docs, doc)
	}

	// 5. Wstrzykiwanie błędu (Fault Injection Logic)
	injected := false
	services := make([]int, 0)

	// Znajdź indeksy wszystkich serwisów
	for i, doc := range docs {
		if kind, ok := doc["kind"].(string); ok && kind == "Service" {
			services = append(services, i)
		}
	}

	if len(services) > 0 {
		// Wybierz losowy serwis w sposób deterministyczny
		targetIdx := services[rng.Intn(len(services))]
		targetDoc := docs[targetIdx]

		// Próba wejścia w spec -> selector
		if spec, ok := targetDoc["spec"].(map[string]interface{}); ok {
			if selector, ok := spec["selector"].(map[string]interface{}); ok {
				// Wybierz pierwszy klucz z selectora i zepsuj go
				for k, v := range selector {
					oldVal := fmt.Sprintf("%v", v)
					newVal := oldVal + "-BROKEN"
					selector[k] = newVal

					fmt.Printf("[SEED %d] Wstrzyknięto błąd w Service: %s\n", *seed, targetDoc["metadata"].(map[string]interface{})["name"])
					fmt.Printf("Zmiana: %s: %s -> %s\n", k, oldVal, newVal)
					injected = true
					break // Psujemy tylko jeden selektor w jednym serwisie
				}
			}
		}
	} else {
		fmt.Println("Ostrzeżenie: Nie znaleziono obiektów Service w pliku wejściowym.")
	}

	if !injected {
		fmt.Println("Ostrzeżenie: Nie udało się wstrzyknąć błędu (brak odpowiedniej struktury).")
	}

	// 6. Zapis wyniku
	outFile, err := os.Create(*outPath)
	if err != nil {
		log.Fatalf("Błąd tworzenia pliku wyjściowego: %v", err)
	}
	defer outFile.Close()

	encoder := yaml.NewEncoder(outFile)
	encoder.SetIndent(2)
	for _, doc := range docs {
		if err := encoder.Encode(doc); err != nil {
			log.Fatalf("Błąd zapisu YAML: %v", err)
		}
	}
	encoder.Close()
	fmt.Printf("Zapisano zepsuty manifest do: %s\n", *outPath)
}
