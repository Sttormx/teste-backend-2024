# Teste iCasei: Backend Júnior/Pleno
Esta aplicação é um sistema de cadastro de produtos que mantém e sincroniza dados em duas bases de dados distintas usando mensageria (Kafka).

## Setup
Para fazer o setup de todos os serviços e dependências do teste, utilize o **Docker Compose**.
```
docker compose up --build -d
```

A aplicação tem um *uptime* de cerca de **20 segundos**. 

Acesse o Kafkdrop pela URL http://localhost:19000/ e **verifique** se os tópicos do Kafka estão criados. Se sim, você já pode fazer o uso da aplicação.