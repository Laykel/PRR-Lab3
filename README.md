# PRR - Laboratoire 3 : Chang et Roberts avec pannes

_Auteurs : Jael Dubey, Luc Wachter_

## Description

Ce laboratoire a pour but l'implémentation de l'algorithme de Chang et Roberts afin de permettre aux processus lancés
de déterminer lequel d'entre eux est le plus apte à une tâche donnée. En outre, un élément de tolérance aux pannes
doit être ajouté, de sorte qu'un processus qui n'arrive pas à contacter le suivant essaie avec celui d'après et ainsi
de suite.

Le logiciel, s'assure que tous les processus (nombre configuré dans le fichier json de configuration) ont été lancés.
Ensuite, une élection est lancée par chaque processus, afin de déterminer celui avec la plus haute aptitude. Une fois
ce travail terminé, les processus interrogent périodiquement l'élu, afin de s'assurer qu'il fonctionne encore. Si ce
n'est pas le cas, une nouvelle élection est démarrée.

## Utilisation

### Ajustement des paramètres dans le fichier

1. Ouvrir le fichier `main/parameters.json`.
2. Adapter les valeurs pour le nombre de processus, et les détails de chaque processus.

### Lancer les processus

1. Se positionner dans le dossier `PRR-Lab3/` (à la racine des packages).
2. Ouvrir autant de terminaux que précisé dans le fichier de paramètres.
3. Lancer dans chaque terminal le programme, en précisant l'id du processus : `go run main/main.go 0`.
4. `go run main/main.go 1`, etc.

---

Le programme attend alors que tous les processus soient lancés, puis lance une élection.

## Implémentation


### Packages et fichiers

- Le package `main` contient le point d'entrée principal du programme. Il lance le serveur TCP dans une go routine afin d'écouter les connections entrantes, puis lance le processus client dans une autre go routine (en leur passant à chacun les channels nécessaires à la communication inter-processus.)
- Le package `network` contient les fichiers suivants.
    - Le fichier `protocol.go` contient les types utilisés pour la communication réseau ainsi que des valeurs importantes et les fonctions d'encodage en bytes et de décodage.
    - Le fichier `connection.go` contient le serveur de réception TCP et la fonction d'envoi de requêtes.

### Problèmes connus

- Nous avons effectué des tests uniquement sur le package `network`, malheureusement.
