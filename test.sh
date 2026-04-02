#!/bin/bash
# Script de test du Projet RED

echo "=========================================="
echo "  TEST DU PROJET RED"
echo "=========================================="
echo ""

# Vérifier Go
echo "1. Vérification de Go..."
if command -v go &> /dev/null; then
    echo "   ✅ Go est installé"
    go version
else
    echo "   ❌ Go n'est pas installé"
    exit 1
fi

echo ""
echo "2. Compilation du projet..."
go build -o bin/projet-red.exe src/main.go
if [ $? -eq 0 ]; then
    echo "   ✅ Compilation réussie"
else
    echo "   ❌ Erreur de compilation"
    exit 1
fi

echo ""
echo "3. Vérification de l'exécutable..."
if [ -f "bin/projet-red.exe" ]; then
    echo "   ✅ Exécutable créé"
    ls -lh bin/projet-red.exe
else
    echo "   ❌ Exécutable non trouvé"
    exit 1
fi

echo ""
echo "=========================================="
echo "  ✅ TOUS LES TESTS SONT PASSES!"
echo "=========================================="
echo ""
echo "Pour lancer le jeu, exécutez :"
echo "  go run src/main.go"
echo "  OU"
echo "  ./bin/projet-red.exe"
