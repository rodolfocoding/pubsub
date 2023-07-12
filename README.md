# Publish/Subscribe Library

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Uma biblioteca simples de Publish/Subscribe em Go, que suporta roteamento básico de mensagens com base nos nomes das filas, assim como o registro de todas as mensagens publicadas em um sistema de arquivos.

## Funcionalidades

- Roteamento de mensagens com base nos nomes das filas
- Registro de todas as mensagens publicadas em um arquivo de log
- Suporte a assinantes múltiplos por fila
- Operações de publicação, assinatura e cancelamento de assinatura

## Requisitos

- Go 1.15 ou superior

## Instalação

Para utilizar esta biblioteca em seu projeto, é necessário ter o Go instalado. Em seguida, você pode instalar a biblioteca utilizando o comando `go get`:

```shell
go get -u github.com/rodolfocoding/pubsub
```
