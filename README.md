# 🎮 Minecraft Events Community

Une application web moderne pour gérer et suivre les événements de la communauté Minecraft. Construite avec Go, Echo, Templ et TailwindCSS 4.

![Go Version](https://img.shields.io/badge/Go-1.25.5-00ADD8?logo=go)
![Echo](https://img.shields.io/badge/Echo-v4.13.4-00ADD8)
![Templ](https://img.shields.io/badge/Templ-v0.3.960-FF6B6B)
![TailwindCSS](https://img.shields.io/badge/TailwindCSS-4-38B2AC)

## ✨ Fonctionnalités

- 📅 **Gestion complète des événements** (CRUD)

  - Créer de nouveaux événements
  - Modifier les événements existants
  - Supprimer des événements avec confirmation
  - Affichage en grille de cards élégantes

- 🎨 **Interface moderne et responsive**

  - Design moderne avec TailwindCSS 4
  - Police Google Inter intégrée
  - Logo personnalisé dans la navigation
  - Animations et transitions fluides
  - Timeline visuelle des événements

- 🔗 **Intégration Discord**

  - Liens directs vers les serveurs Discord
  - Boutons d'action pour chaque événement

- 💾 **Base de données SQLite**

  - Persistance des données avec GORM
  - Migration automatique des schémas
  - Données de démonstration incluses

- 📦 **Assets embarqués**
  - Fichiers statiques intégrés avec `embed.FS`
  - Déploiement simplifié (un seul binaire)

## 🚀 Installation

### Prérequis

- Go 1.25.5 ou supérieur
- [Templ CLI](https://templ.guide/quick-start/installation) pour la génération des templates

```bash
# Installer Templ
go install github.com/a-h/templ/cmd/templ@latest
```

### Installation du projet

```bash
# Cloner le dépôt
git clone <votre-repo>
cd jevent

# Installer les dépendances
go mod download

# Générer les templates Templ
go tool templ generate

# Compiler le projet
go build -o minecraftevent
```

## 🎯 Utilisation

### Démarrer le serveur

```bash
# Mode développement
go run .

# Ou exécuter le binaire compilé
./minecraftevent
```

Le serveur démarre sur **http://localhost:3000**

### Routes disponibles

| Méthode | Route         | Description                             |
| ------- | ------------- | --------------------------------------- |
| GET     | `/`           | Liste tous les événements               |
| GET     | `/create`     | Formulaire de création d'événement      |
| POST    | `/create`     | Créer un nouvel événement               |
| GET     | `/edit/:id`   | Formulaire d'édition d'événement        |
| POST    | `/edit/:id`   | Mettre à jour un événement              |
| DELETE  | `/delete/:id` | Supprimer un événement                  |
| GET     | `/assets/*`   | Fichiers statiques (CSS, fonts, images) |

## 📁 Structure du projet

```
jevent/
├── assets/
│   ├── fonts/              # Polices Google Inter
│   │   ├── Inter-Regular.woff2
│   │   └── Inter-Bold.woff2
│   ├── img/                # Images et logo
│   │   └── logo-minecraft.png
│   └── styles.css          # Styles CSS personnalisés
├── database/
│   └── database.go         # Configuration et initialisation DB
├── handlers/
│   └── handlers.go         # Gestionnaires de routes
├── models/
│   └── event.go            # Modèle de données Event
├── templates/
│   ├── create.templ        # Template de création
│   ├── edit.templ          # Template d'édition
│   ├── index.templ         # Page d'accueil et liste
│   └── layout.templ        # Layout principal
├── main.go                 # Point d'entrée de l'application
├── go.mod                  # Dépendances Go
└── README.md               # Ce fichier
```

## 🛠️ Technologies utilisées

### Backend

- **[Go](https://golang.org/)** - Langage de programmation
- **[Echo](https://echo.labstack.com/)** - Framework web performant
- **[GORM](https://gorm.io/)** - ORM pour Go
- **[SQLite](https://www.sqlite.org/)** - Base de données embarquée

### Frontend

- **[Templ](https://templ.guide/)** - Moteur de templates type-safe pour Go
- **[TailwindCSS 4](https://tailwindcss.com/)** - Framework CSS moderne
- **[Google Fonts - Inter](https://fonts.google.com/specimen/Inter)** - Typographie

### Assets

- **embed.FS** - Système de fichiers embarqué de Go
- Logo Minecraft personnalisé

## 📝 Modèle de données

### Event

```go
type Event struct {
    ID          uint
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   *time.Time
    Title       string      // Titre de l'événement
    Organizer   string      // Organisateur
    Description string      // Description
    StartDate   time.Time   // Date de début
    EndDate     *time.Time  // Date de fin (optionnel)
    ImageURL    string      // URL de l'image
    DiscordURL  string      // Lien Discord
}
```

## 🎨 Fonctionnalités de l'interface

### Page d'accueil

- Grille responsive de cards d'événements
- Timeline visuelle avec numérotation
- Affichage des images ou icône par défaut
- Boutons d'action (Modifier, Supprimer, Discord)

### Création/Édition

- Formulaire intuitif avec validation
- Tous les champs du modèle Event
- Support des dates de début et fin
- Pré-remplissage en mode édition

### Suppression

- Confirmation JavaScript avant suppression
- Suppression AJAX sans rechargement de page
- Feedback visuel immédiat

## 🔧 Développement

### Générer les templates après modification

```bash
templ generate
```

### Mode watch (développement)

```bash
# Terminal 1 : Watch des templates
templ generate --watch

# Terminal 2 : Watch du serveur
# Si air est installé (hot reload)
air # ou go run .
```

### Build pour production

```bash
# Générer les templates
go tool templ generate

# Compiler avec optimisations
go build -ldflags="-s -w" -o minecraftevent

# Le binaire contient tous les assets
./minecraftevent
```

## 🚀 Déploiement

L'application génère un seul binaire autonome incluant :

- Le code compilé
- Tous les assets (CSS, fonts, images)
- Les templates compilés

Pour déployer :

1. Compiler le projet
2. Copier le binaire sur le serveur
3. Lancer `./jevent`

La base de données SQLite (`events.db`) sera créée automatiquement au premier démarrage.

## 📄 Licence

Ce projet est sous licence MIT.

## 👤 Auteur

Créé avec ❤️ pour la communauté Minecraft

---

**Built with Echo + Templ + TailwindCSS 4** 🚀
