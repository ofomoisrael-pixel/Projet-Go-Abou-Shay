.PHONY: build run test clean help

# Variables
GO := go
BINARY := bin/projet-red.exe
MAIN := src/main.go

help:
	@echo "Projet RED - Commandes disponibles:"
	@echo ""
	@echo "  make build     - Compile le projet"
	@echo "  make run       - Exécute directement le code"
	@echo "  make clean     - Supprime les binaires"
	@echo "  make test      - Compile et vérifiie les erreurs"
	@echo "  make all       - Build + run"
	@echo ""

build:
	@echo "Compilation du Projet RED..."
	@mkdir -p bin
	@$(GO) build -o $(BINARY) $(MAIN)
	@echo "✅ Compilation réussie!"
	@echo "Exécutable: $(BINARY)"

run:
	@echo "Lancement du Projet RED..."
	@$(GO) run $(MAIN)

all: build run

clean:
	@echo "Nettoyage des fichiers compilés..."
	@rm -rf bin/
	@$(GO) clean
	@echo "✅ Nettoyage terminé"

test:
	@echo "Vérification du code..."
	@$(GO) build -o /dev/null $(MAIN) 2>&1 && echo "✅ Code valide" || echo "❌ Erreurs trouvées"

fmt:
	@echo "Formatage du code..."
	@$(GO) fmt ./...
	@echo "✅ Formatage terminé"

lint:
	@echo "Analyse statique du code..."
	@$(GO) vet ./...
	@echo "✅ Analyse terminée"
