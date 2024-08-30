Groupie Tracker

Groupie Tracker est une application développée en Go pour suivre les informations sur les groupes de musique via une API donnée. Cette application permet de manipuler les données fournies par l'API afin de créer un site web convivial affichant les informations sur les groupes, leurs concerts à venir et passés, ainsi que des visualisations de données.
Structure du projet

Le projet est organisé comme suit :

    clients: Ce répertoire contient les clients HTTP utilisés pour communiquer avec l'API, ainsi que les fichiers HTML pour les différentes pages du site web.

        style: Ce sous-répertoire contient les fichiers CSS utilisés pour styliser les pages web.
            artist.css: Fichier CSS pour la mise en forme spécifique à la page des artistes.
            error.css: Fichier CSS pour la mise en forme spécifique aux pages d'erreur.
            style.css: Fichier CSS principal pour la mise en forme générale du site.

        artist.html: Page affichant les informations sur les artistes.

        error400.html: Page d'erreur 400 (Bad Request).

        error404.html: Page d'erreur 404 (Not Found).

        error500.html: Page d'erreur 500 (Internal Server Error).

        index.html: Page d'accueil du site web.

    constants: Ce répertoire contient les constantes utilisées dans le projet.
        api.go: Ce fichier contient les constantes pour les URL de l'API utilisées dans le projet.
        models.go: Ce fichier contient les modèles de données utilisés dans le projet.

    loaders: Les modules de chargement des données depuis l'API résident dans ce répertoire. Ces modules sont responsables de la récupération des données depuis l'API et de leur formatage pour être utilisées dans l'application.

        artist.go: Ce fichier contient le module de chargement des données pour les artistes. Il récupère les données sur un artiste spécifique à partir de l'API en utilisant son ID, puis formate ces données pour être utilisées dans l'application.

        index.go: Ce fichier contient le module de chargement des données pour l'index. Il récupère les données sur tous les artistes à partir de l'API, puis formate ces données pour être utilisées dans l'application.

        loaders.go: Ce fichier contient des fonctions utilitaires pour le chargement des données depuis l'API.

    routes: Ce répertoire contient les définitions des routes de l'application web. Les routes sont utilisées pour gérer les différentes requêtes HTTP entrantes et pour rediriger ces requêtes vers les contrôleurs appropriés.

        errs.go: Ce fichier contient les définitions des routes pour gérer les erreurs HTTP. Il spécifie les fonctions de gestion pour les erreurs 400, 404 et 500, qui affichent les pages d'erreur correspondantes.

        index.go: Ce fichier contient les définitions des routes pour gérer la page d'accueil de l'application. Il spécifie les fonctions de gestion pour la page d'accueil, qui affiche les informations générales sur les artistes et les événements à venir.

        artists.go: Ce fichier contient les définitions des routes pour gérer les requêtes liées aux artistes. Il spécifie les endpoints et les fonctions de gestion pour récupérer, créer, mettre à jour et supprimer des données sur les artistes.

        init.go: Ce fichier contient la fonction Init() qui initialise toutes les routes de l'application.

    test: Les tests unitaires et d'intégration sont regroupés dans ce répertoire. Ces tests sont utilisés pour vérifier le bon fonctionnement des différentes parties de l'application, y compris les clients HTTP, les loaders et les routes.

    go.mod: Le fichier de configuration Go modules se trouve ici. Ce fichier est utilisé pour déclarer les dépendances du projet et pour gérer les versions des packages Go utilisés.

    main.go: Le point d'entrée de l'application. Ce fichier contient le code principal de l'application, y compris l'initialisation du serveur web et la gestion des requêtes HTTP entrantes.

Installation

    Assurez-vous d'avoir Go installé sur votre système.
    Clonez ce dépôt dans votre espace de travail local.
    Exécutez go mod tidy pour installer les dépendances.
    Lancez l'application en exécutant go run main.go.

Utilisation

Une fois l'application lancée, vous pouvez accéder au site web à l'adresse suivante : http://localhost:3001.
Contribuer

Les contributions sont les bienvenues ! Pour contribuer à ce projet, veuillez suivre les étapes suivantes :

    Fork le projet.
    Créez une branche pour votre fonctionnalité (git checkout -b fonctionnalité/ma-fonctionnalité).
    Committez vos modifications (git commit -am 'Ajouter une nouvelle fonctionnalité').
    Poussez vers la branche (git push origin fonctionnalité/ma-fonctionnalité).
    Créez une nouvelle pull request.

Auteurs

    Baran   AKSOY
    Martin  DEPREAU
    Toni    DO CARMO
    Oumar   NDIAYE
    Jason   ZACKO