# WORKER-CLAIM-DECLARATION

Le but de ce projet est d'être appelé via les step transitions de la bdd claim dans les cas où on souhaite réaliser une déclaration de sinistre chez les clients utilisant les API VEOS et AD-HOC.

Ce projet fonctionne de manière très similaire au projet [worker-claim-document](https://github.com/ProovGroup/worker-claim-document) qui sert à envoyer des documents.

## Utilisation

Depuis la table steps_infos il y a deux colonnes à renseigner;

Dans la colonne `lambda_function_name` il faut inscrire le nom du worker "worker-claim-declaration".

Puis dans la colonne `parameters` il faut écrire au format JSON les paramètres requis (cf. [Paramètres](#paramètres)).

## Paramètres

Voici le format du JSON attendu à renseigner dans la colonne `parameters` de la table steps_infos:

```json
{
    "api": "VEOS"
}
```

Ces informations sont ensuite récupérées dans la lambda et utilisé afin de déterminé l'API à utilisé ainsi que les lignes dans les tables `prequalif` et `files` à récupérer afin de construire la requête de déclaration de sinistre.

Le champs "api" peut contenir les valeurs "VEOS", "ADHOC" ou "AD-HOC", la casse n'est pas imporante.

## Fonctionnement

En utilisant les steps transitions, une fois le rapport dans le bon état il est possible de déclencher le workflow.

Le workflow des steps transitions va ensuite appeler la lambda en lui passant en paramètre le proov_code ainsi que le contenu de la colonne `parameters` de la manière suivante:

```json
{
    "proov_code": "AA1234",
    "params": {
        "api": "VEOS"
    }
}
```

Ces informations sont récupérées dans la lambda et utilisées afin de déterminer l'API à utiliser ainsi que les lignes dans les tables `prequalif` et `files` à récupérer afin de construire la requête de déclaration de sinistre.

Une fois la requête envoyé un numéro de sinistre est reçu puis enregistré dans la colonne `sinistre_id` de la table `files` (sauf en cas d'échec).
