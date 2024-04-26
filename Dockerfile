
# Web frontend builder
FROM node:20 AS web-builder

WORKDIR /web/

COPY ./frontend/app/package.json ./frontend/app/package-lock.json ./

RUN npm install --frozen-lockfile

COPY ./frontend/app .

RUN npm run build

# Go backend builder
FROM golang:1.22 AS go-builder

WORKDIR /app/

COPY ./backend/go.mod ./backend/go.sum ./

RUN go mod download

COPY ./backend /app/

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/bin/backend

# Final image
FROM scratch AS final

WORKDIR /app/bin

COPY --from=go-builder /app/bin/backend /app/bin/backend
COPY --from=web-builder /web/dist /app/bin/public
COPY  ./docs /app/

ENTRYPOINT [ "./backend" ]