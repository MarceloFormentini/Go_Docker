# a partir da imagem Golang (Alpine uma versão mais enxuta)
FROM golang:1.24-alpine

# pasta onde o projeto vai ficar
WORKDIR /app

# copiar todos os arquivos do projeto dentro dp contêiner
COPY . .

# comando para manter o contêiner executando
CMD ["tail", "-f", "/dev/null"]