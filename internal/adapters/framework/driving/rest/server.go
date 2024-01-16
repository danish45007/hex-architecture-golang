package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/danish45007/hex-architecture-golang/internal/ports"
)

// Adapter represents the REST adapter.
type Adapter struct {
	api ports.APIPort
}

// NewAdapter creates a new instance of Adapter.
func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

// Run starts the REST server.
func (resta *Adapter) Run() {
	http.HandleFunc("/addition", func(w http.ResponseWriter, r *http.Request) {
		_, _ = resta.GetAddition(w, r)
	})
	http.HandleFunc("/subtraction", func(w http.ResponseWriter, r *http.Request) {
		_, _ = resta.GetSubtraction(w, r)
	})
	http.HandleFunc("/multiplication", func(w http.ResponseWriter, r *http.Request) {
		_, _ = resta.GetMultiplication(w, r)
	})
	http.HandleFunc("/division", func(w http.ResponseWriter, r *http.Request) {
		_, _ = resta.GetDivision(w, r)
	})

	var port int = 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("HTTP server listening on port %d\n", port)
	http.ListenAndServe(addr, nil)
}

// HandleAddition handles the HTTP request for the addition operation.
func (resta *Adapter) GetAddition(w http.ResponseWriter, r *http.Request) (int32, error) {
	aStr := r.FormValue("a")
	bStr := r.FormValue("b")

	a, err := strconv.ParseInt(aStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid value for 'a'", http.StatusBadRequest)
		return 0, err
	}

	b, err := strconv.ParseInt(bStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid value for 'b'", http.StatusBadRequest)
		return 0, err
	}

	result, err := resta.api.GetAddition(int32(a), int32(b))
	if err != nil {
		http.Error(w, "Failed to perform addition", http.StatusInternalServerError)
		return 0, err
	}

	fmt.Fprintf(w, "Result of addition: %d", result)
	return result, nil
}

func (resta *Adapter) GetSubtraction(w http.ResponseWriter, r *http.Request) (int32, error) {
	aStr := r.FormValue("a")
	bStr := r.FormValue("b")

	a, err := strconv.ParseInt(aStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid value for 'a'", http.StatusBadRequest)
		return 0, err
	}

	b, err := strconv.ParseInt(bStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid value for 'b'", http.StatusBadRequest)
		return 0, err
	}

	result, err := resta.api.GetSubtraction(int32(a), int32(b))
	if err != nil {
		http.Error(w, "Failed to perform subtraction", http.StatusInternalServerError)
		return 0, err
	}

	fmt.Fprintf(w, "Result of subtraction: %d", result)
	return result, nil
}

func (resta *Adapter) GetMultiplication(w http.ResponseWriter, r *http.Request) (int32, error) {
	aStr := r.FormValue("a")
	bStr := r.FormValue("b")

	a, err := strconv.ParseInt(aStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid value for 'a'", http.StatusBadRequest)
		return 0, err
	}

	b, err := strconv.ParseInt(bStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid value for 'b'", http.StatusBadRequest)
		return 0, err
	}

	result, err := resta.api.GetMultiplication(int32(a), int32(b))
	if err != nil {
		http.Error(w, "Failed to perform multiplication", http.StatusInternalServerError)
		return 0, err
	}

	fmt.Fprintf(w, "Result of multiplication: %d", result)
	return result, nil
}

func (resta *Adapter) GetDivision(w http.ResponseWriter, r *http.Request) (int32, error) {
	aStr := r.FormValue("a")
	bStr := r.FormValue("b")

	a, err := strconv.ParseInt(aStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid value for 'a'", http.StatusBadRequest)
		return 0, err
	}

	b, err := strconv.ParseInt(bStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid value for 'b'", http.StatusBadRequest)
		return 0, err
	}

	result, err := resta.api.GetDivision(int32(a), int32(b))
	if err != nil {
		http.Error(w, "Failed to perform division", http.StatusInternalServerError)
		return 0, err
	}

	fmt.Fprintf(w, "Result of division: %d", result)
	return result, nil
}
