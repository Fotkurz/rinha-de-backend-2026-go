# Detectação de Fraude

Projeto escrito para participar da [Rinha de backend 2026](https://github.com/zanfranceschi/rinha-de-backend-2026/)
totalmente em Golang.

Sinta-se livre para clonar e usar o código como quiser.

## Requisitos

- Golang 1.26.1


## Estrutura

```
cmd/api
    main.go # lança server http
cmd/preprocessor
    main.go # pre processamento do references.json.gz
internal/
    handler/ # handlers http da aplicação
        dto/ # modelos de requisição e resposta http
    service/ # regras de negócio
    domain/ # modelos do domínio
    repository/ # camada de persistência
pkg/ # lib criadas pro projeto
    vector/ # lib com funções reusáveis para vetores
    env/ # lib para ler env vars
assets/ # recursos adicionais    
```

## Requisitos

### Funcionais

- API que determina se uma transação é uma fraude ou não, baseado em dados de referência pré-determinados e pré-compilados.
- API que determina o estado da aplicação.

### Não funcionais

- Load Balancer.
- Número mínimo de replicas (2).

