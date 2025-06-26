# README

## 📌 About

This GUI monitors CPU usage, RAM usage (with total RAM in bytes), and disk usage (with total disk space in bytes) on your computer.
The UI is still under development. If you'd like to contribute to this project, feel free to contact me 😊.

## 🚀 Getting Started

### ⚖️ Installation

Ensure you have the following installed:

- Go >= 1.20
- Node.js + npm
- Wails CLI v.2.10.1
- Gopsutil v3

> wails doctor

This checks all Wails dependencies.

### Before Running Projects

Before you running this project on development mode or production mode, make sure you've run this on frontend folder:

```bash
npm run build
```

After you run that command, make sure dist folder is in your frontend folder.

## ⚙️ Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect to this in your browser, and you can call your Go code from devtools.

Run in live dev mode with hot-reloading:

> wails dev

- Frontend will be served via Vite.
- Access Go methods in the browser via: http://localhost:34115
- You can test Go logic through browser devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## Application Configuration

Create your own config file at:

> ./config/config.yaml

Note: The config/ folder is ignored via .gitignore, so your local configs remain private.

🔑 Example `config/config.yaml`

```yaml
sub_config_a:
  field_a: "http://localhost:8080"
  field_b: 100
  field_c: 1.23
  field_d: true

sub_config_b:
  field_a: "http://localhost:8081"
  field_b: 200
  field_c: 2.34
  field_d: false

sub_config_c:
  field_a: "http://localhost:8082"
  field_b: 300
  field_c: 3.45
  field_d: true
```

You can define multiple sub-configs and load them using custom Go logic under `internal/config/config.go`.

## 🧱 Project Structure

This architecture follows a Clean Architecture pattern inspired by microservices to encourage separation of concerns and testability. Project structure is focused to reduce complexcity and enhance modularity

```plaintext
├── build/          # Build scripts & platform-specific assets
├── config/         # App config files (YAML/ENV) – .gitignored
├── interfaces/     # API, controller layer, data-binding exposed to UI
│   ├── api/        # Frontend-bound Go methods
├── frontend/       # React + TypeScript frontend (Vite-powered)
│   ├── src/        # Frontend-bound Go methods
│   │   ├── assets/ # Frontend-bound Go methods
├── internal/       # Business logic (models, services, utils)
│   ├── config/     # Load config from config.yaml or .env
│   ├── sensor/     # Get sensor data functions
├── main.go         # App entry point, Wails bootstrapper
├── go.mod          # Go module declarations
├── go.sum          # Go module version locks
├── wails.json      # Wails config (entry, frontend dir, build info)
└── README.md       # You're here!
```

## 👨‍💻 Author

##### 📧 mikoaf
