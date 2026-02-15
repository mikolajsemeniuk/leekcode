package leekcode

// ImportanceMap przechowuje globalną statystykę wystąpień kluczy
type ImportanceMap map[string]int

// Pruner statystyczny
type Pruner struct {
	Threshold float64
	Stats     ImportanceMap
	TotalDocs int
}

func NewPruner(threshold float64) *Pruner {
	return &Pruner{
		Threshold: threshold,
		Stats:     make(ImportanceMap),
	}
}

// Train analizuje pliki, aby nauczyć się, co jest "szumem" (częste klucze)
func (p *Pruner) Train(data map[string]interface{}) {
	p.TotalDocs++
	p.countKeys(data)
}

func (p *Pruner) countKeys(data map[string]interface{}) {
	for k, v := range data {
		p.Stats[k]++
		if subMap, ok := v.(map[string]interface{}); ok {
			p.countKeys(subMap)
		}
	}
}

// Prune usuwa klucze, które są zbyt powszechne (niska entropia)
func (p *Pruner) Prune(data map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range data {
		// Uproszczona miara entropii: jeśli klucz występuje w > 80% plików, uznajemy go za szum
		score := float64(p.Stats[k]) / float64(p.TotalDocs)

		if score > p.Threshold && k != "name" && k != "kind" {
			// Pomijamy "szum", ale zachowujemy kluczowe identyfikatory
			continue
		}

		if subMap, ok := v.(map[string]interface{}); ok {
			result[k] = p.Prune(subMap)
		} else {
			result[k] = v
		}
	}
	return result
}
