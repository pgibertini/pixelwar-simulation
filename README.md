# PIXELWAR IA04

## Contributeurs

- Thomas Delplanque
- Valentin Fagnet
- Pierre Gibertini
- Amaury Michel

## 1. Fonctionnement du projet

### 1.a Présentation de la "Pixel War"

Le projet s'inspire de la "Pixel War" : la guerre des pixels apparue sur le `r/Place` créé par le site communautaire de discussion `Reddit`.

Points clés du `r/Place` :
- Projet collaboratif apparu lors du 1er avril 2017
- Toile numérique où chaque utilisateur peut changer la couleur d’un pixel
- Deux éditions : 2017 (1000x1000 pixels) et 2022 (2000x2000 pixels)
- Créations artistiques et collaboratives
- Pourquoi une guerre ? ⇒ France vs. coalition Espagne - Etats-Unis - Canada
- Défense drapeau et attaque territoire ennemi

**Résultat final du `r/Place` 2022 :**

![Résultat final du r/Place 2022](https://placedata.reddit.com/data/final_place.png)

### 1.b Objectif du projet

Ce projet a pour objectif de répondre à la problématique suivante :

**Comment simuler la Pixel War ?**

Pour y répondre, nous avons implémenter deux types d'agents :
- Agents `Manager` (que l'on pourrait assimiler aux *streamers* lors de l'édition 2022 du `r/Place`) dont le rôle est de donner les instructions aux troupes
- Agents `Worker`, (que l'on pourrait assimiler aux *viewers*) dont le rôle est de placer des pixels selon les ordre donné par leur `Manager`

Les agents sont regroupés autour de centres d'intérêts, appelé `hobby`. Pour chaque `hobby` est associé un agent `Manager` et une multitude d'agents `Worker`.

À chaque `hobby` est associé une image pré-définie que les `Manager` cherchent à représenter sur le `canvas`, en donnat des instructions à leur `Worker`.

L'architecture est détaillée dans la section 3.

## 2. Utilisation

Dans cette section, nous détaillons comment installer et utiliser le *front-end* ainsi que le *back-end*.

**Le *front-end* permet uniquement la visualisation, le lancement de la simulation doit se faire depuis le *back-end*.**

### 2.a Installation

#### I. *Front-end*

Le *front-end* a été développé en `React`. Un front-end simplifié, pur `javascript` et natif au *back-end* est également disponible (voir **Nota Bene**).

Pré-requis : `npm`

Lien du dépôt Git : https://gitlab.utc.fr/pixelwar_ia04/Frontend_IA04

**Installation du *Front-end* :**

**1.** `git clone https://gitlab.utc.fr/pixelwar_ia04/Frontend_IA04.git`
**2.** `npm install`

**Lancement du *Front-end* :**

**3.** `npm start`

Cela devrait ouvrir une page dans votre navigateur à sur l'url http://localhost:3000/

**4.** Lancer le *back-end*
**5.** Dans l'onglet `personnalisation`, saisir `1` dans la case la `place-id` (case la plus, à gauche) afin de renseigner l'id du `place` que l'on souhaite visualiser
**6.** Cliquer sur `start` pour débuter l'affichage

> **Nota Bene :** en complément (ou en cas de problème), un front-end basique en pure javascript est disponible directement sur le *back-end*
> 
> Pour y accéder, lancer le back-end est acccéder à l'url suivante depuis votre navigateur : http://localhost:5555/canvas?placeID=place1

#### II. *Back-end*

Pré-requis : `go`

Lien du dépôt Git : https://gitlab.utc.fr/pixelwar_ia04/pixelwar

**Installation du *Back-end* :**

Installation avec `go install` (usage limité) :
> `go install -v gitlab.utc.fr/pixelwar_ia04/pixelwar@latest`

Installation avec `git clone` :
> `git clone https://gitlab.utc.fr/pixelwar_ia04/pixelwar.git`

### 2.b Usage

#### 1. Lancement de la simulation

Si installation du back-end avec `go install` : `$GOPATH/bin/pixelwar`
Si installation du back-end avec `git clone` : `go run main.go`

Lancer le *front-end* avec `npm start` puis accéder au place `1` tel que décrit dans "Lancement du *front-end*" (ou à défaut, accéder directement à http://localhost:5555/canvas?placeID=place1)

#### 2. Customisation des paramètres de simulation

Le script `main.go` peut être modifié afin de changer les paramètres de simulation. Pour cela, changer les paramètres de la fonction `LaunchPixelWar`. Des exemples de paramètrages sont disponibles en tant que code commenté.

Voici les différents paramètres et leur signification :

| paramètre            | type     | exemples de valeurs possibles           | signification                                                                                                                                                                                        |
|----------------------|----------|-----------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `nbWorkerPerManager` | `string` | `"random"`, `"proportional"` ou `"100"` | nombre de *worker* par *manager*. peut être renseigné par une valeur numérique, ou `"proportional"` (nombre de worker proportionnel à la taille de l'image), ou `"random"` (aléatoire entre 5 et 50) |
| `cooldown`           | `int`    | `1` (en secondes)                       | temps minimal entre chaque placement de pixel par un même agent                                                                                                                                      |
| `size`               | `int`    | `500`                                   | hauteur et largeur du *canvas*                                                                                                                                                                       |
| `conquestValue`      | `int`    | `5`                                     | appétence des managers à conquérir des territoires en plaçant leur image à divers endroit différents                                                                                                 |

Voici différentes configurations de simulation pré-établies :
```go
go launcher.LaunchPixelWar("proportional", 1, 500, 0)
go launcher.LaunchPixelWar("random", 1, 500, 5)
go launcher.LaunchPixelWar("50", 1, 200, 10)
go launcher.LaunchPixelWar("50", 1, 500, 10)
```

> **Nota Bene :** rien n'empêche de lancer ces 4 (ou plus) *pixel wars* en même temps. Leur visualisation sera accessible en spécifiant les différents `place-id` dans le front-end.
>
> Cela se fait soit via l'onglet "personnalisation" du front-end `React` comme expliqué plus haut, soit dans l'url du front-end basique natif au back-end (le lien est directement imprimé dans la console). 

Il est également possible d'activer le mode `debug` afin d'afficher les logs, en passant les valeurs suivantes à `true`, dans `main.go` :
```
server.Debug = true
agent.Debug = true
```

#### 3. Bonus : reproduction de la Pixel War 2022

Nous avons implémenté un petit script permettant de mettre à l'épreuve notre serveur REST. 
Ce script utilise l'historique des pixels placés lors de la *Pixel War* 2022 pour effectuer des requêtes sur le serveur.

1. Lancer le serveur : `go run cmd/launch_server/launch_server.go`
2. Lancer le script envoyant des requêtes à partir des données au format `csv` : `go run cmd/simulate_real_pixel_war/simulate_pixel_war.go`
3. Lancer le *front-end* avec `npm start` puis accéder au place `1` tel que décrit dans "Lancement du *front-end*" (ou à défaut, accéder directement à http://localhost:5555/canvas?placeID=place1)
4. Visualiser la grille se dessiner

Sur un laptop (branché sur secteur) équipé d'un processeur Ryzen 4700U, le serveur arrive à traiter environ 7500 requêtes `paint_pixel` par seconde.

> **Nota Bene :** le fichier `csv` de l'historique des actions n'est pas complet (370.000 premières actions).
> Vous pouvez télécharger le fichier complet (10GB compressé pour un fichier csv de 22GB) [sur cette page reddit](https://www.reddit.com/r/place/comments/txvk2d/rplace_datasets_april_fools_2022/).
> Attention, le fichier n'est pas trié selon le `timestamp`.

## 3. Architecture

### 3.a Structuration des agents

### 3.b API REST

#### Commande `/new_place`

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

#### Commande `/paint_pixel`

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

#### Commande `/get_pixel`

> Retourne la couleur d'un pixel pour des coordonnées données.

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

#### Commande `/get_canvas`

> Retourne l'entièreté de la grille.

- Requête : `POST`
- Objet `JSON` envoyé

| propriété    | type     | exemple de valeurs possibles |
|--------------|----------|------------------------------|
| `place-id`   | `string` | `"place1"`                   |
| `reset-diff` | `bool`   | `false`                      |

> L'argument `reset-diff` à `true` permet de réinitialiser la valeur de `diff` d'une grille donnée.
La valeur de `diff` permet de stocker les pixels ayant été modifiés entre 2 requêtes `/get_diff`.

> Rajouter cet attribut est utile au front-end pour afficher l'état de la grille dans un premier temps avec une requête
`/get_canvas` puis d'actualiser son état avec des requêtes `/get_diff`.

> Cet argument permet à la première `/get_diff` de ne renvoyer que la différence depuis la requête `/get_canvas`,
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

#### Commande `/get_diff`

> Retourne les pixels qui différent depuis la dernière requête `/get_diff` sur une grille donnée.

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



## 4. Critique du projet

- Il serait mieux de pouvoir lancer tout depuis le *front*. Cependant cela nécessiterait des changements dans le code que nous n'avons pas eu le temps d'implémenter.

- Nous aurions également pu rajouter des stratégies d'abandon, si un groupe n'arrive pas à remplir son dessin après un certain temps, il abandonne cet emplacement et essaye autre part.

- L'emplacement du dessin est aléatoire, cela est logique au début car tous les agents sont lancés en même temps et regardent donc le même canvas vide. Cependant, comme notre système support la création d'agents managers et workers au milieu de l'exécution, il serait mieux si les managers vérifient si l'emplacement voulu est bien vide, ou capable d'être volé à un autre groupe.

- Il est également possible de donner plusieurs hobby à chaque, cependant ils n'en ont qu'un seul actuellement car ils sont créés en même temps que leur manager.

- Chaque manager n'a qu'un seul dessin lié à son hobby, qu'il reproduit lorsqu'il essaye de prendre plus de place. Il serait possible d'avoir plusieurs dessin par hobby, comme dans la vraie Pixelwar.

- La représentation de la grille sur le *front-end* utilise un *canvas* pour représenter les pixels, ce qui peut donner un rendu un peu flou lorsque l'on zoom.