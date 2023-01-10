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

Pour y répondre, nous avons implémenté deux types d'agents :
- Agents `Manager` (que l'on pourrait assimiler aux *streamers* lors de l'édition 2022 du `r/Place`) dont le rôle est de donner les instructions aux troupes
- Agents `Worker`, (que l'on pourrait assimiler aux *viewers*) dont le rôle est de placer des pixels selon les ordres donnés par leur `Manager`

Les agents sont regroupés autour de centres d'intérêts, appelés `hobby`. Pour chaque `hobby` est associé un agent `Manager` et une multitude d'agents `Worker`.

À chaque `hobby` est associé une image pré-définie que les `Manager` cherchent à représenter sur le `canvas`, en donnant des instructions à leurs agents `Worker`.

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
**5.** Saisir `1` dans la case la `place-id` (en haut à droite) afin de renseigner l'id du `place` que l'on souhaite visualiser
**6.** Cliquer sur `start` pour débuter l'affichage

**En cas de problème, rafraîchir et saisir à nouveau le `place-id`**

> **Nota Bene :** en complément (ou en cas de problème), un front-end basique en pure javascript est disponible directement sur le *back-end*
>
>
> Pour y accéder, lancer le back-end et acccéder à l'url suivante depuis votre navigateur : http://localhost:5555/canvas?placeID=place1

#### II. *Back-end*

Pré-requis : `go`

Lien du dépôt Git : https://gitlab.utc.fr/pixelwar_ia04/pixelwar

**Installation du *Back-end* :**

Installation avec `go install` (usage limité) :
> `go install -v gitlab.utc.fr/pixelwar_ia04/pixelwar@latest`

**Attention :** l'installation avec `go install` nécessite que dossier `images` soit présent dans le repértoire depuis lequel vous souhaitez lancer l'éxecutable

Installation avec `git clone` (recommandé pour customisation des méta-paramètres dans `main.go`) :
> `git clone https://gitlab.utc.fr/pixelwar_ia04/pixelwar.git`

### 2.b Usage

#### 1. Lancement de la simulation

Si installation du back-end avec `go install` : `$GOPATH/bin/pixelwar`
> Attention : le dossier `images` doit être présent dans le repértoire depuis lequel vous lancer l'éxecutable

Si installation du back-end avec `git clone` : `go run main.go`

Lancer le *front-end* avec `npm start` puis accéder au *place* `1` tel que décrit dans "Lancement du *front-end*" (ou à défaut, accéder directement au *front-end* natif à http://localhost:5555/canvas?placeID=place1)

#### 2. Customisation des méta-paramètres de simulation

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
> Cela se fait soit via le champ `place-id` du front-end `React` comme expliqué plus haut, soit dans l'url du front-end basique natif au back-end (le lien est directement imprimé dans la console).

Il est également possible d'activer le mode `debug` afin d'afficher les logs, en passant les valeurs suivantes à `true`, dans `main.go` :
```
server.Debug = true
agent.Debug = true
```

#### 3. Bonus : reproduction de la Pixel War 2022

Nous avons implémenté un petit script permettant de mettre à l'épreuve notre serveur REST.
Ce script utilise l'historique des pixels placés lors de la *Pixel War* 2022 pour effectuer des requêtes sur le serveur.

1. Lancer le serveur : `go run cmd/launch_server/launch_server.go`
2. Lancer le script envoyant des requêtes à partir des données au format `csv` : `go run cmd/simulate_real_pixel_war/simulate_pixel_war.go` (avec `go run`, il faut lancer depuis le dossier racine `pixelwar` afin que le chemin vers le fichier `csv` soit correct)
3. Lancer le *front-end* avec `npm start` puis accéder au place `1` tel que décrit dans "Lancement du *front-end*" (ou à défaut, accéder directement à http://localhost:5555/canvas?placeID=place1)
4. Visualiser la grille se dessiner

Sur un laptop (branché sur secteur) équipé d'un processeur Ryzen 4700U, le serveur arrive à traiter environ 7500 requêtes `paint_pixel` par seconde.

> **Nota Bene :** le fichier `.csv` de l'historique des actions n'est pas complet (370.000 premières actions).
> Vous pouvez télécharger le fichier complet (10GB compressé pour un fichier `.csv` de 22GB) [sur cette page reddit](https://www.reddit.com/r/place/comments/txvk2d/rplace_datasets_april_fools_2022/).
> Attention, le fichier n'est pas trié selon le `timestamp`.

## 3. Architecture

L'architecture du projet est divisée en 3 grandes parties :
- Les agents, en interaction entre eux avec des `channels`, et effectuant des requêtes `http` pour interagir avec les pixels.
- L'API REST afin de gérer la grille et les requêtes pour la modifier.
- Le *Front-end* pour visualiser l'évolution de la grille

### 3.a Structuration des agents

Nous avons différencié 3 types d'agents :

#### Agent `Chat`

L’agent `Chat` fait office de serveur afin de permettre aux agents d’avoir connaissance les uns des autres.

Fonctionnement :
> **Goroutine :** écoute des requêtes des agents
> **Goroutine :** traitement des requêtes `FindWorker` des agents `Manager` afin de gagner en performance, en parallèlisant le traitement des requêtes afin de ne pas bloquer la goroutine écoutant les requêtes.

#### Agent `Worker`

L’agent `Worker` reçoit des ordres d'un agent `Manager` et s’occupe de placer des pixels.

Fonctionnement :
> 1 - S’enregistre sur le chat
> **Goroutine :** attend qu’on lui donne des pixels à placer
> **Goroutine :** place les pixels qu’il a à placer en respectant le cooldown
> -> synchronisation de la liste de pixels à placer entre les 2 goroutines

Les agents `Worker` effectuent des requêtes *http* `/paint_pixel` afin de placer des pixels sur la grille.

#### Agent `Manager`

L’agent `Manager` lit la grille, utilise des calques et donne des ordres aux agents `Worker`.

Fonctionnement :
> 1 - S’enregistre sur le chat
> 2 - Prend connaissance de ses workers
> 3 - Charge son calque
> **Goroutine :** regarde les éléments de son calque qui ne sont pas placés (ou qui ont été volés), envoie des instructions aux agents `Worker` en conséquence

L'image est dessinée à un endroit aléatoire du `canvas`.

Le `Manager` peut décider de dessiner son image à plusieurs endroits différents s'il a déjà complété son image à d'autres endroits, et également en fonction de son paramètre `conquestValue` (plus la valeur est élevée, plus l'agent `Manager` aura tendance à vouloir dessiner son image à plusieurs endoits différents. Si la valeur est `0`, il ne la dessinera qu'à un seul endroit puis la défendra).

Pour donner ses instructions, l'agent `Manager` compare l'état actuel de la grille avec les pixels correspondant à son image, puis attend le temps d'éxécution des instructions par les agents `Manager` (nombre d'instruction par `Worker` * `cooldown`) avant de recommencer.

Les agents `Manager` effectuent des requêtes *http* `/get_canvas` afin de prendre connaissance de la grille et ainsi donner les bons ordres aux agents `Worker` (voir section API REST).


### 3.b API REST

L'API REST a une double utilité :
- Permettre aux agents d'intéragir avec la grille
- Permettre au *Front-end* de connaître l'état de la grille

De ce fait, le seveur REST est totalement indépendant des agents, et pourrait donc fonctionner sur une machine à part. Ce choix a été fait dans une optique de réalisme vis à vis de la vraie *Pixel War*, mais rend le système moins performant que si toutes nos communications (hormis *Front-end*) étaient effectuées par *channel*.

L'idée était également de pouvoir intéragir avec la grille depuis le *front-end*.

Le serveur REST s'assure qu'aucune triche n'est effectuée, notamment au niveau du `cooldown`.

#### Commande `/new_place`

> Créer un nouveau place paramétré et retourne son id.

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

> Permet de peindre un pixel aux coordonnées données.

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

#### Commande `/canvas`

> Permet d'accéder au *front-end* natif.

- Requête : `GET`
- Attributs

| propriété  | type     | exemple de valeurs possibles |
|------------|----------|------------------------------|
| `place-id` | `string` | `"place1"`                   |

L'attribut `place-id` est à passer dans l'url.

Exemple : http://localhost:5555/canvas?placeID=place1

### 3.c Front-end

La principale fonctionnalité du Front-end est de proposer l'affichage d'un place ainsi que son évolution.

#### Technologie

Le *Front-end* de l'application est écrit en JavaScript (librairie React plus précisément).


Celui-ci propose d'une part 2 fonctionnalités d'interaction et d'autre part l'affichage d'un canvas.
Le canvas (ici nommé place), représente l'expression des agents dans un environnement en 2 dimensions.

#### Interactions

**La première interaction est le lancement de requêtes ayant pour objectif la récupération de l'état du place.**

Cette interaction necessite le renseignement de l'id du place dans l'input "place-id"

Après un clic sur le bouton `Start`, 2 requêtes sont alors envoyées :
- `/get_canvas` est effectuée jusqu'à première réponse du serveur, cette requête permet d'initialiser le canvas du Front-end afin de récupérer l'état intégral du place voulu au moment t.
- `/get_diff` est effectué toutes les 0.1s, il permet de récupérer non pas le place en entier, mais seulement la différence de pixels depuis le dernier  `/get_diff` (ou `/get_canvas`). Cette méthode permet d'optimiser les temps de calculs serveurs et donc les temps de réponses aux requêtes.

**La seconde interaction est l'affectation d'une couleur à un pixel du canvas**

Cette interaction s'effectue en renseignant les coordonnées et la couleur du pixel dans l'onglet Personnalisation.

#### Affichage

L'affichage du place s'effectue à l'aide d'une balise canvas.
Les requêtes `/get_diff` renvoient les coordonnées des requêtes modifiées, il est donc aisé de mettre à jour l'état du canvas avec `ctx.fillStyle` et `ctx.fillRect` dès la reception de la réponse serveur.

Nous avons également rendu possible le zoom molette et le déplacement sur celui-ci afin de mieux apprécier les activités des agents.

## 4. Critique du projet

Nous sommes heureux d'avoir pu aboutir à un projet fonctionnel implémentant un premier comportement des agents. Nous avons énormément appris grâce à ce projet que ça soit sur les systèmes multi-agents ou les technologies `Go` et `REST`.

Néanmoins, par manque de temps, nous n'avons pas développé toutes nos idées et émettons quelques critiques à l'égart de notre projet :

- Il serait mieux de pouvoir lancer la simulation depuis le *front-end*. Cependant cela nécessiterait des changements dans le code que nous n'avons pas eu le temps d'implémenter.

- Nous aurions également pu rajouter des stratégies d'abandon, si un groupe n'arrive pas à remplir son dessin après un certain temps, il abandonne cet emplacement et essaye autre part.

- L'emplacement du dessin est aléatoire, cela est logique au début car tous les agents sont lancés en même temps et regardent donc le même canvas vide. Cependant, comme notre système support la création d'agents managers et workers au milieu de l'exécution, il serait mieux si les managers vérifient si l'emplacement voulu est bien vide, ou capable d'être volé à un autre groupe.

- Il est également possible de donner plusieurs hobby à chaque agent `Worker`, cependant ils n'en ont qu'un seul actuellement car ils sont créés en même temps que leur manager.

- Chaque `Manager` n'a qu'un seul dessin lié à son hobby, qu'il reproduit lorsqu'il essaye de prendre plus de place. Il serait possible d'avoir plusieurs dessins par hobby, comme dans la vraie Pixelwar.

- La représentation de la grille sur le *front-end* utilise un *canvas* pour représenter les pixels, ce qui peut donner un rendu un peu flou lorsque l'on zoom.