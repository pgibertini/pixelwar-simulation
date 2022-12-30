# PIXELWAR

## How to add a new feature

1. Checkout to `main` branch: `git checkout main`
2. Pull latest modification: `git pull`
3. Create a new branch from `main` for developing your new feature: `git checkout -b <you_feature_name>`
4. Push when done: `git push`
5. Create a merge request from your newly dev branch to main in the Gitlab interface 


*to be written...*

## Contributors
## API

### Commande `/new_place`

- Requête : `POST`
- Objet `JSON` envoyé

| propriété | type  | exemple de valeurs possibles |
|-----------|-------|------------------------------|
| `height`  | `int` | `500`                        |
| `width`   | `int` | `500`                        |


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
| `x`        | `int`    | `0`                          |
| `color`    | `string` | `"000000"`                   |
| `place-id` | `string` | `place1`                     |


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
| `x`        | `int`    | `0`                          |
| `place-id` | `string` | `place1`                     |


- Code retour

| Code retour | Signification |
|-------------|---------------|
| `200`       | pixel placé   |
| `400`       | bad request   |

- Objet `JSON` renvoyé (si `201`)

| propriété | type     | exemple de valeurs possibles |
|-----------|----------|------------------------------|
| `color`   | `string` | `"000000"`                   |

## Usage
## Installation
...
