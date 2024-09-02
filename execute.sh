#!/bin/bash

# Lê cada linha do arquivo users
for i in $(cat users); do
    # Extrai o IBM e o EMAIL usando o delimitador ':'
    IBM=$(echo "$i" | cut -f1 -d:)
    EMAIL=$(echo "$i" | cut -f2 -d:)

    # Adiciona "000" à esquerda do IBM
    IBM="000$IBM"

    # Chama o script Setusers com o IBM modificado e o EMAIL
    ./Setuser "$IBM" "$EMAIL"
done

