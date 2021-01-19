# quotation

Serviço criado para consumir dados da [**tabela de cotações**](https://www.infomoney.com.br/ferramentas/cotacoes-opcoes-de-acoes/) cedida por **InfoMoney**. Atualmente os dados são exportados apenas para o formato `.csv`.

## Como usar?

Faça download [clicando aqui](https://github.com/thalysonalexr/quotation/archive/main.zip).

Ou, faça o clone do repositório:
```bash
$ git clone https://github.com/thalysonalexr/quotation.git
```

### [**Instale o Go para seu OS!**](https://golang.org/dl/)

Faça download e sincronize o go modules:
```bash
$ go mod download
$ go mod vendor
$ go mod tidy
```

e então execute:
```bash
# executando em modo de trabalhos (jobs)
$ go run .
# executando ui http
$ go run http/main.go
# executando ui cmd
$ go run cmd/main.go
```

Ou você pode gerar o executável do projeto para o seu OS caso prefira, fazendo:
```bash
# gerando executável para modo de trabalhos (jobs)
$ GOOS={seu-sistema} go build . -o quotation-jobs
# gerando executável para ui http
$ GOOS={seu-sistema} go run http/main.go -o quotation-http
# gerando executável para ui cmd
$ GOOS={seu-sistema} go run cmd/main.go -o quotation-cmd
```

Você também pode passar a arquitetura para gerar o executável, por exemplo `GOARCH=ppc64`.

## Sobre as varias variáveis (rsrs engenheiros)
```bash
# configurações do trabalho
CRON_DESCRIPTOR=# aqui você deve colocar uma notação cron para realizar download do csv x vezes #
PATH_SAVE_FILES=# você pode passar outro diretório aqui para armazenar os downloads, por padrão é {diretorio-projeto}/tmp #

# configurações de redis
REDIS_HOST=0.0.0.0 # host do redis, você pode utilizar docker! #
REDIS_PORT=6379 # porta de serviço do redis #
REDIS_ALLOW_EMPTY_PASSWORD=no
REDIS_REPLICATION_MODE=master
REDIS_PASSWORD=development

# server (optional)
SERVER_PORT=8999 # porta em que será exposto o servidor http para requisições #
```

## Quer utilizar com docker?

Muito simples, está tudo configurado! Você precisa apenas instalar o docker e docker-compose em sua máquina. E depois:

```bash
$ docker-compose up # pronto!
```

## Interfaces

Foram criadas 3 interfaces para consumo dos dados da **InfoMoney**. Estas são:

### UI Http
Básicamente temos um servidor http, onde você poderá passar ou não um `path` para a rota `GET` `/download-quotation`.

### UI CMD
Você pode executar o download manualmente via linha de comando, executando:
```bash
$ go run cmd/main.go /meu/diretorio/pra/salvar/os/arquivos
# ou, no caso de um executavel
$ ./quotation-cmd /meu/diretorio/pra/salvar/os/arquivos
```

### Jobs
Você pode executar o download manualmente via linha de comando, executando:
```bash
$ go run . /meu/diretorio/pra/salvar/os/arquivos
# ou, no caso de um executavel
$ ./quotation-jobs /meu/diretorio/pra/salvar/os/arquivos
```

Desenvolvido com :heartbeat: por [**Thalyson Rodrigues**](https://www.linkedin.com/in/thalysonrodrigues/)
