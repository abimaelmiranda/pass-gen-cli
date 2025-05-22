# pass-gen-cli

Um gerador rápido de senhas seguras feito em Go.

## Características
- Gera senhas aleatórias e seguras
- Permite especificar o tamanho da senha (até 255 caracteres)
- Suporta letras maiúsculas, minúsculas, números e símbolos

## Instalação

Clone o repositório e compile o projeto:

```sh
git clone https://github.com/seu-usuario/pass-gen-cli.git
cd pass-gen-cli
go build -o pass-gen-cli
```

## Uso

Execute o comando abaixo para gerar uma senha (o padrão é 10 caracteres):

```sh
./pass-gen-cli --len 16
```

## Exemplo de saída

```
$ ./pass-gen-cli --len 16
A8@kz!2pQw#1Lm$N
```

## Licença

Este projeto está licenciado sob a licença MIT.
