package metrics

import (
	"encoding/json"
)

// ExecutionSuccessRate oblicza wskaźnik sukcesu wykonania zadań w środowisku (np. K8s).
//
// Opis: Określa, jaki ułamek wygenerowanych przez model poleceń MCP
// faktycznie doprowadził do oczekiwanego stanu końcowego.
//
// Formuła:
// $$\text{ESR} = \frac{E_{success}}{E_{total}}$$
// Gdzie $E_{success}$ to liczba zadań, które zmieniły stan klastra zgodnie z intencją,
// a $E_{total}$ to całkowita liczba przeprowadzonych eksperymentów.
func ExecutionSuccessRate(successfulExecutions, totalExecutions int) float64 {
	if totalExecutions == 0 {
		return 0.0
	}
	return float64(successfulExecutions) / float64(totalExecutions)
}

// SyntaxValidationRate oblicza odsetek poprawnych składniowo odpowiedzi.
//
// Opis: Mierzy, jak często model potrafi wygenerować poprawny format danych
// (np. JSON dla argumentów MCP lub YAML dla manifestu K8s), który
// przechodzi przez parser bez błędu rzutowania (unmarshal error).
//
// Formuła:
// $$\text{SVR} = \frac{V_{valid}}{V_{total}}$$
// Gdzie $V_{valid}$ to liczba odpowiedzi z poprawną składnią, a $V_{total}$ to liczba wszystkich odpowiedzi.
func SyntaxValidationRate(validResponses, totalResponses int) float64 {
	if totalResponses == 0 {
		return 0.0
	}
	return float64(validResponses) / float64(totalResponses)
}

// ToolSelectionAccuracy oblicza dokładność wyboru odpowiedniego narzędzia MCP.
//
// Opis: Ocenia zdolność modelu do wybrania właściwej funkcji (narzędzia) z dostępnej puli.
// Nawet jeśli argumenty są błędne, weryfikujemy tu sam fakt podjęcia właściwej decyzji
// (np. użycie `get_pod_logs` zamiast `delete_pod`).
//
// Formuła:
// $$\text{TSA} = \frac{T_{correct}}{T_{total}}$$
func ToolSelectionAccuracy(correctSelections, totalSelections int) float64 {
	if totalSelections == 0 {
		return 0.0
	}
	return float64(correctSelections) / float64(totalSelections)
}

// Tokenizer reprezentuje funkcję do zliczania tokenów (np. owrapowaną bibliotekę tiktoken lub model HF).
type Tokenizer func(text string) int

// CalculateTokenEfficiency kompleksowo oblicza TE (Token Efficiency) na podstawie surowej odpowiedzi modelu.
//
// Opis: Ważna metryka przy porównywaniu gadatliwości modeli. Automatyzuje ona
// proces ekstrakcji, minifikacji JSON-a oraz ponownej tokenizacji zminifikowanego
// ładunku (payload), minimalizując w ten sposób szum i formatowanie.
//
// Formuła:
// $$\text{TE} = \frac{\text{Tokens}_{payload}}{\text{Tokens}_{total}}$$
//
// Academic Methodology:
// To eliminate subjective bias in evaluating model verbosity, the Token Efficiency (TE) metric was calculated deterministically.
// We define the 'useful payload' strictly as the machine-actionable JSON parameters required by the Model Context Protocol (MCP).
// For each successful tool call, the extracted JSON object was structurally minified (stripping whitespace and line breaks) to establish the theoretical minimum information entropy.
// This minified string was then re-tokenized using the native tokenizer of the respective model (e.g., tiktoken for OpenAI models, HuggingFace tokenizers for Qwen).
// The TE is the ratio of these payload tokens to the total completion tokens reported by the API.
// If a model failed to produce valid, parsable JSON, its TE for that task was recorded as 0, penalizing the generation of non-actionable text.
func CalculateTokenEfficiency(rawJSON string, totalCompletionTokens int, tokenize Tokenizer) float64 {
	if totalCompletionTokens == 0 {
		return 0.0
	}

	// 1. Ekstrakcja i minifikacja do postaci "Machine-Actionable Payload"
	minifiedPayload, err := ExtractMachineActionablePayload(rawJSON)
	if err != nil {
		// Jeśli JSON nie jest poprawny (np. model zhalucynował zły format), TE wynosi 0.
		return 0.0
	}

	// 2. Re-tokenizacja zminifikowanego tekstu natywnym tokenizerem
	payloadTokens := tokenize(minifiedPayload)

	// 3. Obliczenie efektywności
	return float64(payloadTokens) / float64(totalCompletionTokens)
}

// ExtractMachineActionablePayload wyciąga i minifikuje parametry JSON niezbędne dla MCP.
//
// Opis: Realizuje założenie minimalnej entropii informacji (theoretical minimum information entropy).
// Pomyślne zminifikowanie JSON-a gwarantuje brak sztucznego zawyżania użytecznych tokenów
// przez formatowanie (spacje, taby, znaki nowej linii).
//
// Formuła minimalizacji szumu:
// $$ P_{minified} = \min_{whitespace} \text{JSON}(P_{raw}) $$
func ExtractMachineActionablePayload(rawJSON string) (string, error) {
	var parsed interface{}
	// json.Unmarshal używa wbudowanego parsera do zdekodowania JSON-a
	err := json.Unmarshal([]byte(rawJSON), &parsed)
	if err != nil {
		// Zgodnie z metodologią, jeśli nie jest poprawnym JSON-em, traktujemy to jako błąd (TE = 0)
		return "", err
	}
	// json.Marshal domyślnie minifikuje strukturę (usuwa spacje i taby)
	minified, err := json.Marshal(parsed)
	if err != nil {
		return "", err
	}
	return string(minified), nil
}

// ContextHallucinationRate oblicza wskaźnik halucynacji argumentów.
//
// Opis: W kontekście dużego szumu informacyjnego i optymalizacji K8s MCP kluczowe
// jest, aby model nie zmyślał parametrów, które nie istnieją w kontekście
// (np. nieistniejąca nazwa Poda wyciągnięta z wag zamiast z RAG-a).
// Metryka ta sprawdza jaki ułamek argumentów narzędzia nie występował w dokumencie.
//
// Formuła:
// $$\text{CHR} = \frac{A_{hallucinated}}{A_{total}}$$
// Gdzie $A_{hallucinated}$ to liczba argumentów nieistniejących w kontekście,
// a $A_{total}$ to całkowita liczba wygenerowanych argumentów.
func ContextHallucinationRate(hallucinatedArgs, totalArgs int) float64 {
	if totalArgs == 0 {
		return 0.0
	}
	return float64(hallucinatedArgs) / float64(totalArgs)
}

// SchemaComplianceRate oblicza wskaźnik pełnej zgodności ze schematem narzędzia.
//
// Opis: SVR mierzy tylko poprawność składniową JSON-a (czy to w ogóle JSON).
// SCR (Schema Compliance Rate) określa odsetek odpowiedzi, w których
// payload w 100% pasuje do oczekiwanego schematu JSON narzędzia MCP (JSON Schema),
// posiadając wszystkie wymagane pola i poprawne typy bez dodatkowych halucynacji.
//
// Formuła:
// $$\text{SCR} = \frac{C_{schema\_valid}}{C_{total}}$$
func SchemaComplianceRate(schemaValidResponses, totalResponses int) float64 {
	if totalResponses == 0 {
		return 0.0
	}
	return float64(schemaValidResponses) / float64(totalResponses)
}

// ContextDensityScore (CDS) mierzy stopień wykorzystania dostarczonego kontekstu RAG.
//
// Opis: Weryfikuje, czy model potrafi wyekstrahować kluczowe informacje z dużych manifestów K8s.
// Wysoki CDS przy niskim CHR (Hallucination) dowodzi wyższej inteligencji analitycznej modelu.
//
// Formuła:
// $$ \text{CDS} = \frac{T_{relevant}}{T_{context}} $$
// Gdzie $T_{relevant}$ to liczba tokenów z kontekstu (np. nazwy podów, selektory)
// faktycznie użytych w poprawnym wywołaniu MCP, a $T_{context}$ to całkowita długość okna kontekstowego.
func ContextDensityScore(relevantTokens, contextWindowTokens int) float64 {
	if contextWindowTokens == 0 {
		return 0.0
	}
	return float64(relevantTokens) / float64(contextWindowTokens)
}

// LatencyToActionEfficiency (LAE) definiuje sprawność operacyjną modelu.
//
// Opis: Kluczowa metryka dla IEEE Access. Pokazuje relację sukcesu wykonania (ESR)
// do czasu odpowiedzi. Pozwala wykazać przewagę modeli Qwen/DeepSeek (niski latency)
// nad modelami takimi jak Anthropic/Vertex w zadaniach Real-time DevOps.
//
// Formuła:
// $$ \text{LAE} = \frac{\text{ESR}}{L_{avg}} $$
// Gdzie $L_{avg}$ to średni czas do wykonania akcji (Latency) wyrażony w sekundach.
func LatencyToActionEfficiency(esr float64, avgLatencySeconds float64) float64 {
	if avgLatencySeconds <= 0 {
		return 0.0
	}
	return esr / avgLatencySeconds
}
