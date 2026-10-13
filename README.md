# Go Reloaded

Programme en Go qui lit un fichier texte, applique des corrections et écrit le résultat dans un autre fichier.

## Utilisation

```sh
go run . sample.txt result.txt
```

Le fichier de sortie est remplacé s'il existe déjà.

Le programme gère les conversions `(hex)` et `(bin)`, les changements de casse `(up)`, `(low)` et `(cap)`, la ponctuation, les apostrophes et le passage de `a` à `an`. Les marqueurs de casse acceptent un nombre de mots, par exemple `(up, 2)`.

## Tests

```sh
go test ./...
```
