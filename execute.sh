#!/bin/bash

# Lê cada linha do arquivo users


for i in $(cat users); do
    # Extrai o IBM, EMAIL e PERMISSAO usando o delimitador ':'
    IBM=$(echo "$i" | cut -f1 -d:)
    EMAIL=$(echo "$i" | cut -f2 -d:)
    PERMISSAO=$(echo "$i" | cut -f3 -d:)

    # Define PERMISSAO como 0 se estiver vazia
    PERMISSAO=${PERMISSAO:-0}

    # Adiciona "000" à esquerda do IBM
    IBM="000$IBM"

    # Chama o script Setuser com o IBM modificado, EMAIL e PERMISSAO
    ./Setuser "$IBM" "$EMAIL" "$PERMISSAO"
done
