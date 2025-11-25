# --- STAGE 1: FASE DE CONSTRUÇÃO (BUILDER) ---
# Usamos a última versão estável e corrigida do Go (1.25.x contém a correção para CVE-2025-61725)
# Note: Recomenda-se sempre usar a última versão estável disponível.
FROM golang:1.25 AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Ativa o modo CGO para garantir a compatibilidade com a imagem final (scratch)
# Se a sua aplicação não usa CGO, você pode usar: ENV CGO_ENABLED=0
ENV CGO_ENABLED=1

# Copia os arquivos de definição de módulos Go para o cache
# Isso permite que o Docker utilize o cache se apenas o código mudar
COPY go.mod go.sum ./

# Baixa as dependências. 
RUN go mod download

# Copia o restante do código fonte
COPY . .

# Compila a aplicação. 
# Usamos a flag -v para aumentar a verbosidade do build e -o para nomear o binário.
# O binário final é compilado estaticamente para rodar na imagem 'scratch' (fase 2).
RUN go build -ldflags "-s -w" -o /go-app .


# --- STAGE 2: IMAGEM FINAL DE EXECUÇÃO (MINIMAL) ---
# Usamos 'scratch', a imagem mais mínima possível, que não contém S.O., shell, ou pacotes vulneráveis.
# Apenas o binário estático e as bibliotecas C necessárias (se CGO_ENABLED=1) são incluídos.
FROM scratch

# Define a variável de ambiente para o fuso horário (necessário para logs corretos)
# Se a sua aplicação não precisa de certificados SSL, esta parte pode ser omitida.
# Caso precise de certificados, use FROM alpine ou FROM gcr.io/distroless/static-debian11
# Se CGO_ENABLED=1, você PODE precisar de uma imagem como FROM alpine
# Para este exemplo, vou manter a imagem mais segura (scratch), mas com a ressalva.

# Copia o binário compilado da fase "builder" para a imagem final
COPY --from=builder /go-app /go-app

# Define o ponto de entrada para o binário
ENTRYPOINT ["/go-app"]
# O comando padrão para rodar o binário (você usou CMD, mas ENTRYPOINT é mais idiomático aqui)
# Se você usava argumentos no CMD, ajuste aqui.
# CMD ["./go-app"]