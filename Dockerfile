FROM golang:1.18.4-alpine3.16
RUN cd src/
WORKDIR /src
COPY . .
ENV URL_SHORTENER_PROJECT_ SPREADSHEET_ID=1P7Me-PLTskt4v-LLDxnCu92X30Z0p7M4DhzWLS6hDF4
RUN go mod download
CMD ["go", "run", "main.go"]