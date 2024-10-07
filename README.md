# Application de Suivi des Dépenses

Cette application en ligne de commande, écrite en Go, vous permet de gérer vos dépenses personnelles. Elle offre des fonctionnalités pour ajouter, mettre à jour, supprimer et visualiser vos dépenses, ainsi que pour obtenir des résumés et exporter vos données.

## Fonctionnalités

- Ajouter une nouvelle dépense avec description, montant et catégorie
- Mettre à jour une dépense existante
- Supprimer une dépense
- Lister toutes les dépenses
- Afficher un résumé des dépenses totales
- Afficher un résumé des dépenses pour un mois spécifique
- Exporter les dépenses vers un fichier CSV

## Prérequis

- Go 1.16 ou version supérieure
- Package `github.com/urfave/cli/v2`

## Installation

1. Clonez ce dépôt ou téléchargez le fichier source.

2. Installez les dépendances nécessaires :
   ```
   go get github.com/urfave/cli/v2
   ```

3. Compilez l'application :
   ```
   go build -o expense-tracker
   ```

## Utilisation

Voici les commandes disponibles :

- Ajouter une dépense :
  ```
  ./expense-tracker add --description "Déjeuner" --amount 20 --category "Alimentation"
  ```

- Mettre à jour une dépense :
  ```
  ./expense-tracker update --id 1 --description "Dîner" --amount 25
  ```

- Supprimer une dépense :
  ```
  ./expense-tracker delete --id 1
  ```

- Lister toutes les dépenses :
  ```
  ./expense-tracker list
  ```

- Afficher un résumé des dépenses :
  ```
  ./expense-tracker summary
  ```

- Afficher un résumé pour un mois spécifique :
  ```
  ./expense-tracker summary --month 8
  ```

- Exporter les dépenses vers un fichier CSV :
  ```
  ./expense-tracker export --file mes_depenses.csv
  ```

## Exemples d'utilisation

```
$ ./expense-tracker add --description "Courses" --amount 50.75 --category "Alimentation"
Dépense ajoutée avec succès (ID : 1)

$ ./expense-tracker add --description "Cinéma" --amount 12 --category "Loisirs"
Dépense ajoutée avec succès (ID : 2)

$ ./expense-tracker list
ID      Date            Description     Montant     Catégorie
1       2024-08-06      Courses         $50.75      Alimentation
2       2024-08-06      Cinéma          $12.00      Loisirs

$ ./expense-tracker summary
Total des dépenses : $62.75

$ ./expense-tracker delete --id 2
Dépense supprimée avec succès (ID : 2)

$ ./expense-tracker summary
Total des dépenses : $50.75
```

## Stockage des données

Les dépenses sont stockées dans un fichier JSON nommé `expenses.json` dans le même répertoire que l'application. Ce fichier est créé automatiquement lors de la première utilisation.

## Contribution

Les contributions sont les bienvenues ! N'hésitez pas à ouvrir une issue ou à soumettre une pull request pour suggérer des améliorations ou signaler des bugs.

## Licence

Ce projet est sous licence MIT. Voir le fichier `LICENSE` pour plus de détails.