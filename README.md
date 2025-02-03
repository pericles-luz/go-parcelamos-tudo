# go-parcelamos-tudo
Consumer para API da parcelamos tudo

Implementado usando um outro pacote meu para facilitar a implementação de consumers para APIs RESTful, `github.com/pericles-luz/go-rest`. Mas pode ser substituído por outro que atenda à interface

Sinta-se a vontade para contribuir com o projeto, implementando as funcionalidades que achar necessárias.

## Instalação

```bash
go get -u github.com/pericles-luz/go-parcelamos-tudo
```

Funcionalidades implementadas:

### Autenticação

Chama a API da parcelamos tudo para autenticar o usuário e obter o token de acesso.
Pela documentação, a autenticação é feita através de um POST na rota `/auth/login` com o corpo da requisição contendo `grant_type`, `client_id`, `client_secret` e `scopes`. Necessita também dos headers `api-version` com o valor `1`, `Content-Type` e `Accept` com o valor `application/json`.

### Criar assinatura

Deve ser usada para registrar a adesão de um cliente a um plano existente.

Pela documentação, a criação de uma assinatura é feita através de um POST na rota `/subscriptions` com o corpo da requisição contendo `id_plan`, `id_card`, `charge_type`, `external_reerence_id`, `start_date`, `cycles` e `customer`.

- `id_plan` é o id do plano que o cliente deseja assinar. Este plano já deve esxtar previamente cadastrado na API da parcelamos tudo.
- `id_card` é o id do cartão de crédito que o cliente deseja usar para pagar a assinatura. Este cartão já deve estar previamente cadastrado na API da parcelamos tudo. Mas só é necessário se o `charge_type` for `credit_card`.
- `charge_type` é o tipo de cobrança que o cliente deseja. Pode ser `credit_card` ou `pix`.
- `customer` é um objeto contendo `name`, `email`, `document` e `ip`.

Necessita também dos headers `api-version` com o valor `1`, `Content-Type` e `Accept` com o valor `application/json`.

Importante salvar o `id` da assinatura para futuras consultas e cancelamentos

### Remover assinatura

Deve ser usada para cancelar a assinatura de um cliente.

Pela documentação, a remoção de uma assinatura é feita através de um DELETE na rota `/subscriptions/{id}` onde `{id}` é o id da assinatura que se deseja cancelar.

Necessita também dos headers `api-version` com o valor `1`, `Content-Type` e `Accept` com o valor `application/json`.

### Criar Plano

Deve ser usada para cadastrar um novo plano na API da parcelamos tudo.

Pela documentação, a criação de um plano é feita através de um POST na rota `/plans` com o corpo da requisição contendo `name`, `description`, `external_reference_id`, `currency`, `amount`, `period`, `days_until_due`.

- `name` é o nome do plano
- `description` é a descrição do plano
- `external_reference_id` é o id usado no seu sistema para identificar o plano
- `currency` é a moeda do plano e só pode ser `BRL`
- `amount` é o valor do plano em centavos(`int32`)
- `period` é o período de cobrança do plano. Pode ser `monthly`, `weekly` ou `yearly`
- `days_until_due` é o número de dias até o vencimento da cobrança