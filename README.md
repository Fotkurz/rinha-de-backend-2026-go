# Detectação de Fraude

Projeto escrito para participar da [Rinha de backend 2026](https://github.com/zanfranceschi/rinha-de-backend-2026/)
totalmente em Golang.

Sinta-se livre para clonar e usar o código como quiser.

## Requisitos

- Golang 1.26.1


## Estrutura

```
cmd/api
    main.go # lança a aplicação
internal/
    handler/ # handlers http da aplicação
        dto/ # modelos de requisição e resposta http
    service/ # regras de negócio
    domain/ # modelos do domínio
    repository/ # camada de persistência
pkg/ # lib criadas pro projeto
    vector/ # lib com funções reusáveis para vetores
assets/ # recursos adicionais    
