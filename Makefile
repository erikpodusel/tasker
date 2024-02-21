export PORT=:8080

run:
	npx tailwindcss build -i tailwind.css -o static/style.css
	PORT=$(PORT) go run cmd/main.go
