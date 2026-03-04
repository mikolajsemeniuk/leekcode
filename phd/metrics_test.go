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
func TestEvaluationMetrics(t *testing.T) {
	// 1. Symulacja danych z eksperymentu
	expectedToolName := "get_pod_logs"

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
			CompletionTokens: 45, // Model wygenerował w sumie 45 tokenów
			TotalTokens:      195,
		},
	}

	// --- SVR: Syntax Validation Rate ---
	isValidJSON := true
	var dummy map[string]interface{}
	if err := json.Unmarshal([]byte(mockResponse.Choices[0].Message.ToolCalls[0].Arguments), &dummy); err != nil {
		isValidJSON = false
	}

	validCount := 0
	if isValidJSON {
		validCount = 1
	}
	svr := SyntaxValidationRate(validCount, 1) // 1 poprawny na 1 test
	if svr != 1.0 {
		t.Errorf("Oczekiwano SVR = 1.0, otrzymano %v", svr)
	}

	// --- TSA: Tool Selection Accuracy ---
	correctToolCount := 0
	if mockResponse.Choices[0].Message.ToolCalls[0].Name == expectedToolName {
		correctToolCount = 1
	}
	tsa := ToolSelectionAccuracy(correctToolCount, 1) // 1 poprawny wybór na 1 test
	if tsa != 1.0 {
		t.Errorf("Oczekiwano TSA = 1.0, otrzymano %v", tsa)
	}

	// --- ESR: Execution Success Rate ---
	// Zgodnie z Twoją uwagą, przekazujemy teraz dwie liczby.
	// Symulujemy paczkę 50 zadań wysłanych do klastra K8s, z czego 38 zakończyło się sukcesem.
	successfulExecutions := 38
	totalExecutions := 50
	esr := ExecutionSuccessRate(successfulExecutions, totalExecutions)
	expectedESR := 0.76 // 38 / 50
	if math.Abs(esr-expectedESR) > 0.001 {
		t.Errorf("Oczekiwano ESR = %v, otrzymano %v", expectedESR, esr)
	}

	// --- TE: Token Efficiency ---
	// 1. Ekstrakcja i Normalizacja Danych (Minifikacja)
	// Zgodnie z metodologią naukową, definiujemy "użyteczny token" algorytmicznie.
	rawArgs := mockResponse.Choices[0].Message.ToolCalls[0].Arguments
	
	// 2. Pobranie Całkowitej Liczby Tokenów
	mockResponse.Usage.CompletionTokens = 68
	
	// 3. Re-tokenizacja (Zdefiniowanie tokenizera dla modelu)
	tkm, err := tiktoken.EncodingForModel("gpt-4o")
	if err != nil {
		t.Fatalf("Błąd podczas ładowania tokenizera: %v", err)
	}

	realTokenizer := func(text string) int {
		tokens := tkm.Encode(text, nil, nil)
		return len(tokens)
	}
	
	te := CalculateTokenEfficiency(rawArgs, mockResponse.Usage.CompletionTokens, realTokenizer)
	expectedTE := 11.0 / 68.0 // ~0.1617
	if math.Abs(te-expectedTE) > 0.001 {
		t.Errorf("Oczekiwano TE = %v, otrzymano %v", expectedTE, te)
	}

	// Dowód dla recenzenta w konsoli
	minifiedPayload, _ := ExtractMachineActionablePayload(rawArgs)
	t.Logf("Raw JSON: %s", rawArgs)
	t.Logf("Minified JSON (Machine-Actionable Payload): %s", minifiedPayload)

	// --- SCR: Schema Compliance Rate ---
	// Symulacja: Oczekiwaliśmy kluczy "pod_name" oraz "namespace". Model podał je prawidłowo.
	// Ale w drugim teście model podał nieprawidłowy klucz "pod" zamiast "pod_name".
	// Testujemy 2 próbki - 1 prawidłowa, 1 nieprawidłowa.
	schemaValidResponses := 1
	totalSchemaResponses := 2
	scr := SchemaComplianceRate(schemaValidResponses, totalSchemaResponses)
	expectedSCR := 0.50
	if math.Abs(scr-expectedSCR) > 0.001 {
		t.Errorf("Oczekiwano SCR = %v, otrzymano %v", expectedSCR, scr)
	}

	// --- CHR: Context Hallucination Rate ---
	// Symulacja: Dostarczono z RAG-a nazwy podów ["nginx", "redis"].
	// W 10 odpowiedziach modelu, model wymyślił "mysql" jako parametr w 2 przypadkach.
	// Zatem 2 na 10 wygenerowanych argumentów to halucynacje.
	hallucinatedArgs := 2
	totalArgs := 10
	chr := ContextHallucinationRate(hallucinatedArgs, totalArgs)
	expectedCHR := 0.20
	if math.Abs(chr-expectedCHR) > 0.001 {
		t.Errorf("Oczekiwano CHR = %v, otrzymano %v", expectedCHR, chr)
	}

	// Wypisanie logów do konsoli (go test -v)
	t.Logf("--- Wyniki Ewaluacji (Baza testowa: %d zadań) ---", totalExecutions)
	t.Logf("Syntax Validation Rate (SVR): %.2f%%", svr*100)
	t.Logf("Schema Compliance Rate (SCR): %.2f%%", scr*100)
	t.Logf("Tool Selection Accuracy (TSA): %.2f%%", tsa*100)
	t.Logf("Execution Success Rate (ESR): %.2f%%", esr*100)
	t.Logf("Token Efficiency (TE): %.2f%%", te*100)
	t.Logf("Context Hallucination Rate (CHR): %.2f%%", chr*100)
}
