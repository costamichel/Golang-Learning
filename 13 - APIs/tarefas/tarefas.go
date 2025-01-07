package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Tarefa struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

var (
	tarefas   []Tarefa
	nextID    int = 1
	tarefasMu sync.Mutex
)

func listarTarefas(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(tarefas)
}

func adicionarTarefa(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var novaTarefa Tarefa
	err := json.NewDecoder(r.Body).Decode(&novaTarefa)
	if err != nil {
		http.Error(w, "Error when decoding JSON", http.StatusBadRequest)
		return
	}

	tarefasMu.Lock()
	fmt.Println(novaTarefa)
	novaTarefa.ID = nextID
	nextID++
	tarefas = append(tarefas, novaTarefa)
	tarefasMu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novaTarefa)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		listarTarefas(w, r)
	case "POST":
		adicionarTarefa(w, r)
	case "DELETE":
		atualizarOuRemoverTarefa(w, r)
	case "PUT":
		atualizarOuRemoverTarefa(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func atualizarOuRemoverTarefa(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" && r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	strId := r.URL.Path[len("/"):]
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	idx_found := -1
	for idx, tarefa := range tarefas {
		if tarefa.ID == id {
			idx_found = idx
			break
		}
	}

	if idx_found < 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "PUT":
		var dados Tarefa
		err := json.NewDecoder(r.Body).Decode(&dados)
		if err != nil {
			http.Error(w, "Error when decoding json", http.StatusBadRequest)
			return
		}

		if dados.Nome == "" {
			http.Error(w, "Field 'nome' couldn't be blank", http.StatusBadRequest)
			return
		}

		tarefasMu.Lock()
		defer tarefasMu.Unlock()
		tarefas[idx_found].Nome = dados.Nome
		var tarefaAtualizada *Tarefa
		tarefaAtualizada = &tarefas[idx_found]

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(tarefaAtualizada)

	case "DELETE":
		tarefasMu.Lock()
		defer tarefasMu.Unlock()
		tarefas = append(tarefas[:idx_found], tarefas[idx_found+1:]...)
		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error when starting server")
	}
}
