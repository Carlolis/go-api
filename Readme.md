Test Technique - Développeur Backend

L’objectif de ce test est d’écrire un micro service exposantune API RESTful en
Golang.

L'api est exposé sur le port 8010.

Le serveur doit permettre les opérations suivantes sur un élément
simple que nous appellerons “document”, possédant 3 attributs: un ID, un nom
et une description:

- Création => Méthode "POST" sur /document, avec en body le document sans l'id.

- Récupération => Méthode "GET" sur /document pour récupérer tout les documents. Et Méthode "GET" sur /document:{id} pour récupérer un document selon son id.

- Suppression => Méthode "DELETE" sur /document/{id}.
  Les documents seront stockés enmémoire pour cet exercice.

Des tests unitaires sont attendus: c'est fait

Une image Docker empaquetant votre service est un plus :

Pour lancer le docker :

docker image load < go-test.tar
docker run -p 8010:8010 go-charles-ilieff

Travailler dans un repository git local est un plus: c'est fait.
