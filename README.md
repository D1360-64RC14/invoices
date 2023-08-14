# invoices

Aprendizado de Go HTML templates com gin.

## Rotas

### `/` (root)

Página inicial contendo cards com as páginas criadas utilizando templates.

### `/purchase-invoice`

Um Invoice simples, com breve resumo da compra.

### `/detailed-purchase-invoice`

Um Invoice com informações detalhadas de compra.

## Repositório

Este repositório segue o padrão [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

- Principal comando em `cmd/back-end` (função `main` principal);
- Código principal em `internal`;
- Templates em `web/templates`.

## Execução

```bash
go run ./cmd/back-end
```