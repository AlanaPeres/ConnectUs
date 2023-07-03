package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	/* Este é um exemplo de um "import blank" (importação em branco).
	O _ antes do pacote significa que apenas a função de inicialização do pacote será executada,
	mas nenhuma referência direta ao pacote será feita no código. Nesse caso específico,
	estamos importando o pacote github.com/go-sql-driver/mysql,
	que é um driver para o banco de dados MySQL. A importação em branco é usada quando queremos
	executar o código de inicialização de um pacote, mas não precisamos referenciar
	diretamente as funções ou tipos definidos nele.*/)

// Conectar abre a conexão com o banco de dados retornar o objeto *sql.DB representando essa conexão.
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
