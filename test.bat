@echo off
REM Script de test du Projet RED pour Windows

echo ===========================================
echo   TEST DU PROJET RED
echo ===========================================
echo.

REM Vérifier Go
echo 1. Verification de Go...
where go >nul 2>nul
if %ERRORLEVEL% EQU 0 (
    echo    OK - Go est installe
    go version
) else (
    echo    ERREUR - Go n'est pas installe
    exit /b 1
)

echo.
echo 2. Compilation du projet...
go build -o bin/projet-red.exe src/main.go
if %ERRORLEVEL% EQU 0 (
    echo    OK - Compilation reussie
) else (
    echo    ERREUR - Erreur de compilation
    exit /b 1
)

echo.
echo 3. Verification de l'executable...
if exist "bin\projet-red.exe" (
    echo    OK - Executable cree
    dir /S bin\projet-red.exe
) else (
    echo    ERREUR - Executable non trouve
    exit /b 1
)

echo.
echo ===========================================
echo   TOUS LES TESTS SONT PASSES!
echo ===========================================
echo.
echo Pour lancer le jeu, executez :
echo   go run src/main.go
echo   OU
echo   bin\projet-red.exe
pause
