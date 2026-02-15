package leekcode

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestK8sContextPruning(t *testing.T) {
	// 1. Definiujemy "stóg siana" - prawdziwe manifesty z dużą ilością boilerplate'u
	deploymentYaml := `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: api-gateway
  namespace: production
  generation: 12
  creationTimestamp: "2024-01-01T10:00:00Z"
  annotations:
    deployment.kubernetes.io/revision: "1"
    prometheus.io/scrape: "true"
spec:
  replicas: 3
  selector:
    matchLabels:
      app: api-gateway
  template:
    metadata:
      labels:
        app: api-gateway
    spec:
      containers:
      - name: gateway
        image: nginx:1.25.1
        ports:
        - containerPort: 8080
`

	serviceYaml := `
apiVersion: v1
kind: Service
metadata:
  name: api-service
  namespace: production
  labels:
    app: api-gateway
spec:
  type: ClusterIP
  selector:
    app: api-gateway
  ports:
  - port: 80
    targetPort: 8080
status:
  loadBalancer: {}
`

	ingressYaml := `
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: main-ingress
  namespace: production
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
    cert-manager.io/cluster-issuer: letsencrypt-prod
spec:
  rules:
  - host: api.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: api-service
            port:
              number: 80
`

	// 2. Inicjalizacja prunera z progiem 0.6 (usuń klucze obecne w >60% plików)
	pruner := NewPruner(0.6)

	// 3. Trening (wczytujemy wszystkie pliki do statystyk)
	yamls := []string{deploymentYaml, serviceYaml, ingressYaml}
	var docs []map[string]interface{}

	for _, y := range yamls {
		var doc map[string]interface{}
		err := yaml.Unmarshal([]byte(y), &doc)
		if err != nil {
			t.Fatalf("Błąd parsowania YAML: %v", err)
		}
		pruner.Train(doc)
		docs = append(docs, doc)
	}

	// 4. Debugowanie: Prunujemy Deployment i sprawdzamy co zostało
	t.Logf("--- STATYSTYKI KLUCZY (ILE RAZY WYSTĘPUJĄ) ---")
	for k, v := range pruner.Stats {
		t.Logf("Klucz [%s]: %d razy", k, v)
	}

	prunedDoc := pruner.Prune(docs[0]) // Prunujemy pierwszy dokument (Deployment)

	output, _ := yaml.Marshal(prunedDoc)
	fmt.Println("\n=== ORYGINALNY DEPLOYMENT (SKRÓCONY) ===")
	fmt.Println(deploymentYaml)

	fmt.Println("\n=== ZOPTYMALIZOWANY KONTEKST (PO PRUNINGU) ===")
	fmt.Println(string(output))

	// 5. Asercje dla naukowego potwierdzenia
	// apiVersion i metadata są w każdym pliku (3/3 = 100%), więc powinny zniknąć (przy progu 0.6)
	if _, ok := prunedDoc["apiVersion"]; ok {
		t.Error("Klucz 'apiVersion' powinien zostać usunięty!")
	}

	// 'replicas' jest tylko w Deployment (1/3 = 33%), więc musi zostać!
	spec := prunedDoc["spec"].(map[string]interface{})
	if _, ok := spec["replicas"]; !ok {
		t.Error("Klucz 'replicas' (unikalny sygnał) został niesłusznie usunięty!")
	}
}
