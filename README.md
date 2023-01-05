# PIXELWAR

## How to add a new feature

1. Checkout to `main` branch: `git checkout main`
2. Pull latest modification: `git pull`
3. Create a new branch from `main` for developing your new feature: `git checkout -b <you_feature_name>`
4. Push when done: `git push`
5. Create a merge request from your newly dev branch to main in the Gitlab interface 


*to be written...*

## Contributors

## Usage

Front-end basique : `http://localhost:8080/canvas?placeID=place1`

## API

### Commande `/new_place`

- Requête : `POST`
- Objet `JSON` envoyé

| propriété  | type  | exemple de valeurs possibles |
|------------|-------|------------------------------|
| `height`   | `int` | `500`                        |
| `width`    | `int` | `500`                        |
| `cooldown` | `int` | `5` (en secondes)            |


- Code retour

| Code retour | Signification |
|-------------|---------------|
| `201`       | place créé    |
| `400`       | bad request   |

- Objet `JSON` renvoyé (si `201`)

| propriété  | type     | exemple de valeurs possibles |
|------------|----------|------------------------------|
| `place-id` | `string` | `"place1"`                   |

### Commande `/paint_pixel`

- Requête : `POST`
- Objet `JSON` envoyé

| propriété  | type     | exemple de valeurs possibles |
|------------|----------|------------------------------|
| `x`        | `int`    | `0`                          |
| `y`        | `int`    | `0`                          |
| `color`    | `string` | `"#000000"`                  |
| `place-id` | `string` | `"place1"`                   |
| `user-id`  | `string` | `"user1"`                    |


- Code retour

| Code retour | Signification |
|-------------|---------------|
| `200`       | pixel placé   |
| `400`       | bad request   |
| `425`       | too early     |

### Commande `/get_pixel`

- Requête : `POST`
- Objet `JSON` envoyé

| propriété  | type     | exemple de valeurs possibles |
|------------|----------|------------------------------|
| `x`        | `int`    | `0`                          |
| `y`        | `int`    | `0`                          |
| `place-id` | `string` | `"place1"`                   |

- Code retour

| Code retour | Signification |
|-------------|---------------|
| `200`       | OK            |
| `400`       | bad request   |

- Objet `JSON` renvoyé (si `201`)

| propriété | type     | exemple de valeurs possibles |
|-----------|----------|------------------------------|
| `color`   | `string` | `"#FFFFFF"`                  |

### Commande `/get_canvas`

- Requête : `POST`
- Objet `JSON` envoyé

| propriété    | type     | exemple de valeurs possibles |
|--------------|----------|------------------------------|
| `place-id`   | `string` | `"place1"`                   |
| `reset-diff` | `bool`   | `false`                      |

L'argument `reset-diff` à `true` permet de réinitialiser la valeur de `diff` d'une grille donnée.
La valeur de `diff` permet de stocker les pixels ayant été modifiés entre 2 requêtes `/get_diff`.

Rajouter cet attribut est utile au front-end pour afficher l'état de la grille dans un premier temps avec une requête
`/get_canvas` puis d'actualiser son état avec des requêtes `/get_diff`.

Cet argument permet à la première `/get_diff` de ne renvoyer que la différence depuis la requête `/get_canvas`, 
et non depuis toujours.

- Code retour

| Code retour | Signification |
|-------------|---------------|
| `200`       | OK            |
| `400`       | bad request   |

- Objet `JSON` renvoyé (si `201`)

| propriété | type         | exemple de valeurs possibles                                     |
|-----------|--------------|------------------------------------------------------------------|
| `height`  | `int`        | `500`                                                            |
| `width`   | `int`        | `500`                                                            |
| `grid`    | `[][]string` | `[[ "#FFFFFF", "#FFFFFF", … ], [ "#FFFFFF", "#FFFFFF", … ], … ]` |

### Commande `/get_diff`

Retourne les pixels qui différent depuis la dernière requête `/get_diff` sur une grille donnée.

- Requête : `POST`
- Objet `JSON` envoyé

| propriété  | type     | exemple de valeurs possibles |
|------------|----------|------------------------------|
| `place-id` | `string` | `"place1"`                   |

- Code retour

| Code retour | Signification |
|-------------|---------------|
| `200`       | OK            |
| `400`       | bad request   |

- Objet `JSON` renvoyé (si `201`)

| propriété | type         | exemple de valeurs possibles            |
|-----------|--------------|-----------------------------------------|
| `diff`    | `[]HexPixel` | `[{ x: 86, y: 962, c: "#FFFFFF" }, … ]` |

## Installation
...
