package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuarioStruct struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(rw http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		rw.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuarioStruct
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		rw.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		rw.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		rw.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		rw.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		rw.Write([]byte("Erro ao obter o ID inserido"))
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))
}

func BuscarUsuarios(rw http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		rw.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		rw.Write([]byte("Erro ao buscar os usuários"))
		return
	}
	defer linhas.Close()

	var usuarios []usuarioStruct
	for linhas.Next() {
		var usuario usuarioStruct

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			rw.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	rw.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(rw).Encode(usuarios); erro != nil {
		rw.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

func BuscarUsuario(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		rw.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		rw.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		rw.Write([]byte("Erro ao buscar o usuário"))
		return
	}
	defer linha.Close()

	var usuario usuarioStruct
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			rw.Write([]byte("Erro ao escanear o usuário"))
			return
		}
	}

	rw.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(rw).Encode(usuario); erro != nil {
		rw.Write([]byte("Erro ao converter o usuário para JSON"))
		return
	}
}

func AtualizarUsuario(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		rw.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		rw.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var usuario usuarioStruct
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		rw.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		rw.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		rw.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		rw.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(rw http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		rw.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		rw.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		rw.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		rw.Write([]byte("Erro ao deletar o usuário"))
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
