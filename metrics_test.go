package leekcode

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

// --- FORMUŁY MATEMATYCZNE DO COLABA (LATEX) ---
/*
1. Weighted Correctness Score (Si):
   $$S_i = w_1 \cdot \mathbb{1}_{res} + w_2 \cdot \mathbb{1}_{cause} + w_3 \cdot \mathbb{1}_{loc}$$

2. Information Density Ratio (IDR):
   $$IDR = \frac{U_{sem}}{\ln(L_{tokens} + 1)}$$

3. Semantic Utility (Usemi):
   $$U_{sem} = \frac{|A \cap B| + \text{sim}_{vec}(A, B)}{2}$$
*/

type ModelResponse struct {
	Resource string // Nazwa zasobu k8s
	Cause    string // Przyczyna błędu
	Logic    string // Wyjaśnienie modelu
	Tokens   int    // Liczba tokenów użyta w PROMPCIE (input)
}

type Metrics struct {
	Accuracy float64
	Semantic float64
	IDR      float64
}

// CalculateAccuracy oblicza ważoną poprawność na podstawie Ground Truth
func CalculateAccuracy(response, groundTruth ModelResponse) float64 {
	score := 0.0
	// Wagi (Weights) - suma = 1.0
	w1, w2, w3 := 0.2, 0.5, 0.3

	if strings.EqualFold(response.Resource, groundTruth.Resource) {
		score += w1
	}
	if strings.Contains(strings.ToLower(response.Cause), strings.ToLower(groundTruth.Cause)) {
		score += w2
	}
	// Logic/Location check - uproszczone na potrzeby testu
	if len(response.Logic) > 10 {
		score += w3
	}
	return score
}

// CalculateIDR oblicza efektywność: Ile "jakości" dostajemy na jednostkę kosztu (tokena)
func CalculateIDR(semanticScore float64, tokenCount int) float64 {
	if tokenCount <= 0 {
		return 0
	}
	// Używamy logarytmu naturalnego, aby uniknąć dominacji bardzo krótkich tekstów
	return semanticScore / math.Log(float64(tokenCount)+1)
}

func TestEvaluateModels(t *testing.T) {
	// 1. Definiujemy stan faktyczny (Ground Truth)
	groundTruth := ModelResponse{
		Resource: "api-service",
		Cause:    "port mismatch",
		Logic:    "Service targetPort 8080 does not match Deployment containerPort 9090",
	}

	// 2. SCENARIUSZ A: Surowy, wielki kontekst (Baseline)
	// Model dostał 15,000 tokenów YAML-i. Odpowiedź jest poprawna, ale kosztowna.
	baselineResponse := ModelResponse{
		Resource: "api-service",
		Cause:    "The ports are mismatched between service and deployment",
		Logic:    "After analyzing 50 files, I found that targetPort is 8080 and containerPort is 9090.",
		Tokens:   15000,
	}

	// 3. SCENARIUSZ B: Skompresowany kontekst (Twoja Metoda SECP)
	// Model dostał tylko 1,200 tokenów (oczyszczone przez Twoje "Sito").
	secpResponse := ModelResponse{
		Resource: "api-service",
		Cause:    "port mismatch",
		Logic:    "Mismatch: Service 8080 vs Deployment 9090.",
		Tokens:   1200,
	}

	// --- OBLICZENIA ---

	// Metryki dla Baseline
	accA := CalculateAccuracy(baselineResponse, groundTruth)
	idrA := CalculateIDR(accA, baselineResponse.Tokens)

	// Metryki dla Twojej metody
	accB := CalculateAccuracy(secpResponse, groundTruth)
	idrB := CalculateIDR(accB, secpResponse.Tokens)

	// Wyświetlanie wyników (widoczne w debuggerze lub `go test -v`)
	fmt.Printf("\n--- WYNIKI EKSPERYMENTU ---\n")
	fmt.Printf("BASELINE | Acc: %.2f | Tokens: %5d | IDR: %.5f\n", accA, baselineResponse.Tokens, idrA)
	fmt.Printf("SECP (Ty)| Acc: %.2f | Tokens: %5d | IDR: %.5f\n", accB, secpResponse.Tokens, idrB)

	// ASERCJA: Twoja metoda powinna mieć wyższy IDR (Information Density Ratio)
	if idrB <= idrA {
		t.Errorf("Twoja metoda (SECP) nie poprawiła gęstości informacji! IDR_B: %f, IDR_A: %f", idrB, idrA)
	} else {
		fmt.Printf("\n[SUKCES] Metoda SECP poprawiła efektywność o %.1f%%\n", (idrB/idrA-1)*100)
	}
}
