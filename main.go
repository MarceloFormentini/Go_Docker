package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" //Driver para PostgreSQL
)

// Struct é uma estrutura de dados, como se fosse uma classe
type Produto struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Preco     float64 `json:"preco"`
}

func main() {
	// Configuração da conexão com o PostgreSQL
	connStr := "host=postgres port=5432 user=postgres password=postgres dbname=db_imersao_full_cycle sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Rota para listar produtos
	http.HandleFunc("/produtos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		produtos, err := listarProdutos(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Erro ao listar produtos: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(produtos)
	})

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}

func listarProdutos(db *sql.DB) ([]Produto, error) {
	rows, err := db.Query("SELECT id, nome, descricao, preco FROM produto")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produtos []Produto
	for rows.Next() {
		var produto Produto
		err := rows.Scan(&produto.ID, &produto.Nome, &produto.Descricao, &produto.Preco)
		if err != nil {
			return nil, err
		}
		produtos = append(produtos, produto)
	}

	// if err := rows.Err(); err != nil {
	// 	return nil, err
	// }

	return produtos, nil
}
