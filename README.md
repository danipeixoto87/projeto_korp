# Projeto Korp

Projeto de demonstração prática de uma aplicação HTTP desenvolvida em Go, containerizada com Docker e orquestrada com Docker Compose, utilizando Nginx como reverse proxy, Prometheus e Grafana para observabilidade e Ansible para provisionamento automatizado.

O projeto foi desenvolvido com foco em práticas de DevOps, automação, containerização, monitoramento e infraestrutura reproduzível.

---

## Arquitetura

```text
                         ┌─────────────────┐
                         │     Cliente     │
                         └────────┬────────┘
                                  │
                                  │ HTTP :80
                                  ▼
                         ┌─────────────────┐
                         │      Nginx      │
                         │  Reverse Proxy  │
                         └────────┬────────┘
                                  │
                                  │ :8080
                                  ▼
                         ┌─────────────────┐
                         │     Go API      │
                         │   Projeto Korp  │
                         └────────┬────────┘
                                  │
                    ┌─────────────┴─────────────┐
                    │                           │
                    │ /metrics                  │ HTTP
                    ▼                           │
             ┌───────────────┐                  │
             │   Prometheus  │                  │
             │     :9090     │                  │
             └───────┬───────┘                  │
                     │                          │
                     │ métricas                 │
                     ▼                          │
             ┌───────────────┐                  │
             │    Grafana    │                  │
             │     :3000     │                  │
             └───────────────┘                  │

             ┌─────────────────┐
             │     Ansible     │
             │  Provisioning   │
             └────────┬────────┘
                      │
                      ▼
             ┌─────────────────┐
             │ Docker Compose  │
             └─────────────────┘
```

Todos os containers utilizam a rede Docker `korp-network`.

---

## Tecnologias utilizadas

- **Go** — desenvolvimento da API HTTP
- **Docker** — containerização
- **Docker Compose** — orquestração dos serviços
- **Nginx** — reverse proxy
- **Prometheus** — coleta e armazenamento de métricas
- **Grafana** — visualização e monitoramento
- **Ansible** — provisionamento e automação
- **Git/GitHub** — versionamento do projeto

---

## Estrutura do projeto

```text
.
├── ansible
│   ├── inventory.ini
│   └── playbook.yml
│
├── app
│   ├── cmd
│   │   └── server
│   │       └── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── internal
│       ├── handlers
│       │   ├── health.go
│       │   └── projeto.go
│       ├── metrics
│       │   └── metrics.go
│       ├── middleware
│       │   └── metrics.go
│       ├── models
│       │   └── response.go
│       └── routes
│           └── routes.go
│
├── grafana
│   ├── dashboards
│   │   └── projeto-korp.json
│   └── provisioning
│       ├── dashboards
│       │   └── dashboard.yml
│       └── datasources
│           └── datasource.yml
│
├── nginx
│   └── http-server-projeto-korp.conf
│
├── prometheus
│   └── prometheus.yml
│
├── docker-compose.yml
└── README.md
```

---

## Aplicação Go

A aplicação fornece três endpoints principais.

### `/projeto-korp`

Retorna informações da aplicação em formato JSON.

Exemplo:

```json
{
  "nome": "Projeto Korp",
  "horario": "2026-08-11T14:45:33Z"
}
```

### `/health`

Endpoint utilizado para verificar a saúde da aplicação.

Resposta:

```text
OK
```

Esse endpoint também é utilizado pelo Docker Healthcheck.

### `/metrics`

Endpoint utilizado pelo Prometheus para coletar as métricas da aplicação.

---

## Métricas

A aplicação utiliza a biblioteca Prometheus para instrumentação.

São disponibilizadas as seguintes métricas principais.

### `http_requests_total`

Contador do total de requisições HTTP.

Labels:

```text
method
path
status
```

Exemplo:

```text
http_requests_total{
  method="GET",
  path="/projeto-korp",
  status="200"
}
```

### `http_request_duration_seconds`

Histograma utilizado para medir a duração das requisições HTTP.

Labels:

```text
method
path
```

---

## Docker

A aplicação Go utiliza um Dockerfile multi-stage.

### Build

A primeira etapa utiliza uma imagem Go para compilar a aplicação.

### Runtime

A segunda etapa utiliza uma imagem Alpine contendo o binário compilado e os componentes necessários para execução.

A aplicação é executada utilizando um usuário não-root dentro do container.

---

## Docker Compose

O ambiente é composto pelos seguintes serviços:

| Serviço | Função | Porta |
|---|---|---:|
| Go | API HTTP | 8080 |
| Nginx | Reverse Proxy | 80 |
| Prometheus | Monitoramento | 9090 |
| Grafana | Dashboard | 3000 |

A aplicação Go utiliza `expose` para disponibilizar a porta `8080` apenas dentro da rede Docker.

O acesso externo à aplicação é realizado através do Nginx.

Todos os serviços utilizam a rede Docker:

```text
korp-network
```

---

## Health Checks

A aplicação Go possui um Docker Healthcheck:

```text
GET http://localhost:8080/health
```

O Prometheus também possui Healthcheck:

```text
GET http://localhost:9090/-/healthy
```

Os serviços utilizam `depends_on` com condições de saúde quando necessário.

A inicialização segue o fluxo:

```text
Go
 │
 └── healthy
       │
       ├── Nginx
       │
       └── Prometheus
              │
              └── healthy
                    │
                    ▼
                  Grafana
```

---

## Nginx

O Nginx atua como reverse proxy para a aplicação Go.

O tráfego recebido na porta `80` é encaminhado para:

```text
http-server-projeto-korp:8080
```

Também são encaminhados headers relacionados à requisição original, incluindo:

```text
Host
X-Real-IP
X-Forwarded-For
X-Forwarded-Proto
```

---

## Prometheus

O Prometheus coleta as métricas da aplicação a cada 15 segundos.

Target configurado:

```text
http-server-projeto-korp:8080
```

As métricas são disponibilizadas através do endpoint:

```text
/metrics
```

---

## Grafana

O Grafana utiliza o Prometheus como datasource.

O datasource é provisionado automaticamente através de:

```text
grafana/provisioning/datasources/datasource.yml
```

O dashboard também é provisionado automaticamente através de:

```text
grafana/provisioning/dashboards/dashboard.yml
```

O dashboard do Projeto Korp apresenta:

- Total de requisições HTTP
- Taxa de requisições HTTP
- Disponibilidade da aplicação

O acesso anônimo está habilitado com permissão de visualização (`Viewer`).

---

## Deploy manual com Docker Compose

Clone o projeto:

```bash
git clone git@github.com:danipeixoto87/projeto_korp.git
cd projeto_korp
```

Crie a rede Docker:

```bash
docker network create --driver bridge korp-network
```

Suba o ambiente:

```bash
docker compose up -d --build
```

Verifique os containers:

```bash
docker compose ps
```

---

## Deploy automatizado com Ansible

O projeto possui um playbook Ansible responsável pelo provisionamento do ambiente.

O inventário utiliza conexão local:

```ini
[local]
localhost ansible_connection=local
```

Execute:

```bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml
```

O playbook realiza as seguintes etapas:

1. Instala Docker e Git.
2. Garante que o Docker esteja iniciado.
3. Verifica a versão do Docker.
4. Clona ou atualiza o projeto.
5. Cria a rede Docker `korp-network`.
6. Executa o Docker Compose.
7. Valida o endpoint `/projeto-korp`.

---

## Endpoints

Após a implantação:

### Aplicação

```text
http://localhost/projeto-korp
```

### Health Check

```text
http://localhost/health
```

### Métricas

```text
http://localhost/metrics
```

### Prometheus

```text
http://localhost:9090
```

### Grafana

```text
http://localhost:3000
```

---

## Comandos úteis

### Ver containers

```bash
docker compose ps
```

### Ver logs

```bash
docker compose logs
```

### Logs de um serviço

```bash
docker compose logs nginx
docker compose logs prometheus
docker compose logs grafana
docker compose logs http-server-projeto-korp
```

### Reiniciar o ambiente

```bash
docker compose restart
```

### Parar o ambiente

```bash
docker compose down
```

### Recriar as imagens

```bash
docker compose up -d --build
```

### Verificar a rede Docker

```bash
docker network inspect korp-network
```

---

## Validação

Após o deploy, a aplicação pode ser validada com:

```bash
curl http://localhost/projeto-korp
```

```bash
curl http://localhost/health
```

```bash
curl http://localhost/metrics
```

O estado dos containers pode ser verificado com:

```bash
docker compose ps
```

A aplicação Go e o Prometheus devem apresentar estado `healthy`.

---

## Objetivos do projeto

Este projeto foi desenvolvido para demonstrar, de forma prática, a integração entre diferentes componentes de um ambiente DevOps:

- Desenvolvimento de uma API HTTP em Go
- Containerização com Docker
- Orquestração com Docker Compose
- Reverse proxy com Nginx
- Instrumentação de aplicações
- Monitoramento com Prometheus
- Visualização de métricas com Grafana
- Health checks
- Automação e provisionamento com Ansible
- Versionamento utilizando Git
- Reprodutibilidade através de infraestrutura automatizada

---

## Autor

**Daniel Peixoto**

Projeto desenvolvido como laboratório prático de DevOps, automação, containerização e observabilidade.
