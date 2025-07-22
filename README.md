# Client Go Wazuh

Ce projet est une bibliothèque client Go conçue pour interagir avec
l'API Wazuh et récupérer les alertes directement depuis l'Indexeur
Wazuh (Elasticsearch/OpenSearch).

## Fonctionnalités

- **Authentification:** Gère l'authentification et le rafraîchissement
des tokens de l'API Wazuh.
- **Version de l'API:** Permet de récupérer la version de l'API Wazuh.
- **Statut des agents:** Fournit des informations sur le statut des
agents (actifs, déconnectés, total, synchronisés, non synchronisés).
- **Récupération des alertes:** Récupère les alertes de l'Indexeur
Wazuh avec support de la pagination.

---

## Installation

Pour utiliser cette bibliothèque dans votre projet Go, exécutez la
commande suivante :

```bash
go get github.com/socme-project/wazuh-go
```

---

## Utilisation

Voici un exemple d'utilisation de la bibliothèque pour rafraîchir
un token et récupérer des alertes :

```go
package main

import (
    "fmt"
    "os"

    wazuhapi "github.com/socme-project/wazuh-go"
)

func main() {
    // Initialisez la structure WazuhAPI avec vos détails de connexion.
    // Il est fortement recommandé d'utiliser des variables d'environnement
    // ou un système de gestion de secrets pour les identifiants en production.
    wazuh := wazuhapi.WazuhAPI{
        Host:     "10.8.178.20", // Remplacez par l'IP/nom d'hôte de votre API Wazuh
        Port:     "55000",       // Port de l'API Wazuh
        Username: "admin",
        Password: "HMthisismys3cr3tP5ssword34a;", // Mot de passe de l'utilisateur de l'API Wazuh
        Indexer: wazuhapi.Indexer{
            Username: "admin",
            Password: "HMthisismys3cr3tP5ssword34a;", // Mot de passe de l'utilisateur de l'Indexeur
            Host:     "10.8.178.20",                   // Remplacez par l'IP/nom d'hôte de votre Indexeur
            Port:     "9200",                          // Port de l'Indexeur
        },
        Insecure: true, // true pour ignorer la vérification du certificat TLS (non recommandé en production)
    }

    // Rafraîchir le token d'authentification
    err := wazuh.RefreshToken()
    if err != nil {
        fmt.Printf("Erreur lors du rafraîchissement du token : %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Token API Wazuh rafraîchi avec succès.")

    // Récupérer les alertes
    lastAlertId := 0 // Utilisez 0 pour la première requête, puis le dernier ID d'alerte pour la pagination.
    alerts, newLastAlertId, err := wazuh.GetAlerts(lastAlertId)
    if err != nil {
        fmt.Printf("Erreur lors de la récupération des alertes : %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("Nombre d'alertes récupérées : %d\n", len(alerts))
    fmt.Printf("Dernier ID d'alerte pour la prochaine requête : %d\n", newLastAlertId)

    // Exemple de récupération du statut des agents
    agentsStatus, err := wazuh.GetAgents()
    if err != nil {
        fmt.Printf("Erreur lors de la récupération du statut des agents : %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("Statut des agents : Actifs=%d, Déconnectés=%d, Total=%d\n",
        agentsStatus.Active, agentsStatus.Disconnected, agentsStatus.Total)

    // Exemple de récupération de la version de l'API
    apiVersion, err := wazuh.GetApiVersion()
    if err != nil {
        fmt.Printf("Erreur lors de la récupération de la version de l'API : %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("Version de l'API Wazuh : %s\n", apiVersion)
}
```

---

## Contributions

Les contributions sont les bienvenues ! N'hésitez pas à
ouvrir une issue ou à soumettre une PR.
