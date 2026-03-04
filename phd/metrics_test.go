package metrics

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/pkoukk/tiktoken-go"
)

// Symulacja struktur odpowiedzi z popularnych SDK (np. go-openai)
type MockToolCall struct {
	Name      string
	Arguments string // Zazwyczaj JSON w formie stringa
}

type MockMessage struct {
	Role      string
	Content   string
	ToolCalls []MockToolCall
}

type MockUsage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

type MockChatCompletionResponse struct {
	Choices []struct {
		Message MockMessage
	}
	Usage MockUsage
}

// TestEvaluationMetrics symuluje proces ewaluacji modelu LLM.
// TestEvaluationMetrics przeprowadza kompleksową ewaluację modelu,
// łącząc klasyczne metryki dokładności z innowacyjnymi metrykami wydajnościowymi (CDS, LAE).
func TestEvaluationMetrics(t *testing.T) {
	// --- 1. Konfiguracja i Symulacja Danych ---
	expectedToolName := "get_pod_logs"
	contextWindowTokens := 8192 // Rozmiar okna kontekstowego (np. duży RAG z YAMLami)
	avgLatency := 0.450         // Symulowane opóźnienie: 450ms (atut mniejszych modeli)

	mockResponse := MockChatCompletionResponse{
		Choices: []struct{ Message MockMessage }{
			{
				Message: MockMessage{
					Role:    "assistant",
					Content: "Jasne, chętnie pomogę! Aby pobrać logi z tego poda, musimy użyć narzędzia MCP z odpowiednimi argumentami. Oto one:\n\nJSON\n{\n   \"pod_name\": \"nginx\",\n   \"namespace\": \"default\"\n}\nMam nadzieję, że to rozwiąże Twój problem z Kubernetesem!",
					ToolCalls: []MockToolCall{
						{
							Name:      "get_pod_logs",
							Arguments: "{\n   \"pod_name\": \"nginx\",\n   \"namespace\": \"default\"\n}",
						},
					},
				},
			},
		},
		Usage: MockUsage{
			PromptTokens:     150,
			CompletionTokens: 68,
			TotalTokens:      218,
		},
	}

	rawArgs := mockResponse.Choices[0].Message.ToolCalls[0].Arguments

	// --- SVR: Syntax Validation Rate ---
	var dummy map[string]interface{}
	isValidJSON := json.Unmarshal([]byte(rawArgs), &dummy) == nil
	validCount := 0
	if isValidJSON {
		validCount = 1
	}
	svr := SyntaxValidationRate(validCount, 1)

	// --- TSA: Tool Selection Accuracy ---
	correctToolCount := 0
	if mockResponse.Choices[0].Message.ToolCalls[0].Name == expectedToolName {
		correctToolCount = 1
	}
	tsa := ToolSelectionAccuracy(correctToolCount, 1)

	// --- ESR: Execution Success Rate ---
	successfulExecutions := 38
	totalExecutions := 50
	esr := ExecutionSuccessRate(successfulExecutions, totalExecutions)

	// --- TE: Token Efficiency ---
	tkm, err := tiktoken.EncodingForModel("gpt-4o")
	if err != nil {
		t.Fatalf("Błąd ładowania tokenizera: %v", err)
	}
	realTokenizer := func(text string) int {
		tokens := tkm.Encode(text, nil, nil)
		return len(tokens)
	}
	te := CalculateTokenEfficiency(rawArgs, mockResponse.Usage.CompletionTokens, realTokenizer)

	// --- SCR: Schema Compliance Rate ---
	schemaValidResponses := 1
	totalSchemaResponses := 2
	scr := SchemaComplianceRate(schemaValidResponses, totalSchemaResponses)

	// --- CHR: Context Hallucination Rate ---
	hallucinatedArgs := 2
	totalArgs := 10
	chr := ContextHallucinationRate(hallucinatedArgs, totalArgs)

	// --- NOWA METRYKA: CDS (Context Density Score) ---
	// Obliczamy ile "gęstej" informacji (payload) model wygenerował w stosunku do całego kontekstu RAG.
	payloadTokens := realTokenizer(rawArgs)
	cds := ContextDensityScore(payloadTokens, contextWindowTokens)

	// --- NOWA METRYKA: LAE (Latency-to-Action Efficiency) ---
	// Obliczamy sprawność operacyjną: sukcesy na sekundę opóźnienia.
	lae := LatencyToActionEfficiency(esr, avgLatency)

	// --- Weryfikacja Poprawności (Asercje) ---
	if math.Abs(esr-0.76) > 0.001 {
		t.Errorf("Błąd ESR: oczekiwano 0.76, otrzymano %v", esr)
	}
	if lae <= 0 {
		t.Errorf("Błąd LAE: wartość musi być większa od 0 przy pomyślnej egzekucji")
	}

	// --- Raport końcowy dla recenzenta (go test -v) ---
	t.Logf("\n--- [PEŁNY RAPORT EWALUACJI NAUKOWEJ] ---")
	t.Logf("1. Syntax Validation Rate (SVR):    %.2f%%", svr*100)
	t.Logf("2. Schema Compliance Rate (SCR):    %.2f%%", scr*100)
	t.Logf("3. Tool Selection Accuracy (TSA):   %.2f%%", tsa*100)
	t.Logf("4. Execution Success Rate (ESR):    %.2f%%", esr*100)
	t.Logf("5. Token Efficiency (TE):           %.2f%%", te*100)
	t.Logf("6. Context Hallucination Rate (CHR):%.2f%%", chr*100)
	t.Logf("-----------------------------------------")
	t.Logf("7. Context Density Score (CDS):     %.6f (Payload/Context)", cds)
	t.Logf("8. Latency-to-Action Efficiency:    %.2f actions/sec", lae)
	t.Logf("-----------------------------------------")

	// Logika dowodowa: mniejszy model wygrywa przez LAE
	if lae > 1.5 && te > 0.15 {
		t.Log("Wniosek: Model wykazuje wysoką efektywność brzegową (Edge Efficiency).")
	}
}
