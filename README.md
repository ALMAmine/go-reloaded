# Go Reloaded

Programme en Go qui lit un fichier texte, applique des corrections et écrit le résultat dans un autre fichier. Il utilise uniquement la bibliothèque standard.

## Récupérer et compiler

```sh
git clone https://github.com/ALMAmine/go-reloaded.git
cd go-reloaded
go build -o go-reloaded .
```

## Utiliser

```sh
go run . sample.txt result.txt
```

Ou après compilation :

```sh
./go-reloaded sample.txt result.txt
```

Le premier argument est le fichier d'entrée, le second est le fichier de sortie. Ce dernier est remplacé s'il existe déjà. `sample.txt` et `result.txt` montrent un exemple complet.

Le programme gère les conversions `(hex)` et `(bin)`, les changements de casse `(up)`, `(low)` et `(cap)`, la ponctuation, les apostrophes et le passage de `a` à `an`. Les marqueurs de casse acceptent un nombre de mots, par exemple `(up, 2)`.

Exemple : `hello world (up, 2) !` devient `HELLO WORLD!`.

## Tests

```sh
go test ./...
```
