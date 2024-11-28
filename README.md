
# SOPlanning 1.52.01 - Remote Code Execution (Authenticated)

## Description

Ce projet est un exploit en Go qui cible une vulnérabilité de type *Remote Code Execution (RCE)* authentifiée dans SOPlanning version 1.52.01. Cette vulnérabilité permet à un attaquant disposant d'identifiants valides d'exécuter des commandes arbitraires sur le serveur via l'upload d'un shell PHP.

## Fonctionnalités

- Exploitation de la vulnérabilité RCE via l'upload d'un fichier malveillant.
- Shell interactif pour l'exécution de commandes sur le serveur cible.
- Génération aléatoire des noms de fichiers et identifiants pour éviter la détection.

## Prérequis

- Langage Go installé ([Go](https://golang.org/dl/)).
- Accès au serveur SOPlanning version 1.52.01 avec des identifiants valides.

## Installation

1. Clonez le dépôt ou copiez le code source dans un fichier Go.
2. Compilez le code source avec Go :
   ```bash
   go build soplanning_rce.go
   ```

## Utilisation

1. Exécutez l'exploit avec la commande suivante :
   ```bash
   ./soplanning_rce
   ```
2. Entrez les informations nécessaires :
   - **URL de la cible** : L'adresse du serveur SOPlanning (par exemple, `http://localhost/soplanning`).
   - **Nom d'utilisateur** : Identifiant d'un utilisateur authentifié.
   - **Mot de passe** : Mot de passe de l'utilisateur.

3. Si l'exploit réussit, une URL vers le shell web est générée. Vous pouvez exécuter des commandes via cette URL ou utiliser le shell interactif intégré.

## Exemple de Session

```plaintext
Target URL (e.g., http://localhost:8080): http://localhost/soplanning
Username: admin
Password: admin123

[+] Uploaded ==> 200 OK
[+] Exploit completed.
Access webshell here: http://localhost/soplanning/upload/files/abc123/xyz.php?cmd=<command>
Do you want an interactive shell? (yes/no): yes
soplaning:~$ whoami
www-data
soplaning:~$ pwd
/var/www/html/soplanning/upload/files/abc123
```

## Avertissement

Ce script est fourni à des fins éducatives et de tests de sécurité uniquement. L'utilisation non autorisée de cet outil pour accéder à des systèmes sans consentement explicite est illégale et peut entraîner des sanctions sévères. L'auteur décline toute responsabilité en cas d'utilisation abusive.
