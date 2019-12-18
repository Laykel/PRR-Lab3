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

### Packages et fichiers

- Le package `main` contient le point d'entrée principal du programme. Il lance un serveur pour répondre aux "pings" des autres processus, ainsi que pour lire les messages entrants. L'algorithme est ensuite initié et une élection lancée.
- Le package `network` contient les fichiers suivants.
    - Le fichier `protocol.go` contient les types utilisés pour la communication réseau ainsi que des valeurs importantes.
    - Le fichier `connection.go` contient le serveur de réception et la fonction d'envoi de requêtes.

### Problèmes connus

Malheureusement, ayant d'autres priorités, nous n'avons pas pu régler tous les problèmes que nous rencontrons avec
le programme. En effet, nous avons mal conçu certains aspects du laboratoire, ce qui résulte en des problèmes de
concurrence : plusieurs envois simultanés semblent poser problème au décodeur Gob, erreur **"extra data in buffer"**.

Néanmoins, avec un seul processus, l'algorithme semble suivre son cours comme voulu, ce qui nous laisse penser que son
implémentation au moins, est juste. Plus nous ajoutons de processus, plus les problèmes précédemment cités surviennent.
Par exemple, avec deux processus, il arrive souvent que le programme fonctionne et que le processus de plus grande
aptitude soit correctement choisi.
