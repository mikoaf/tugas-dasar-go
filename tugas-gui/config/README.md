# 🛠️ Project Configuration & Build

This directory contains essential configuration needed to build and run the project.

## 📁 Folder Structure

```bash
config/                     # Build-related scripts and assets (binary, packaging, etc.)
├── config.yaml             # Project-specific configuration file (YAML)
└── README.md               # Documentation for the config
```

## 📦 Description

- `config/config.yaml`: Used to store all configuration files. These can include environment variables, service endpoints, ports, hardware pin mapping, etc.

## ✅ Usage Notes

Make sure to define your config values inside `config/config.yaml` before building or running the app.

This folder is commonly ignored from version control (.gitignore) if it contains sensitive or environment-specific configurations.

## 💡 Tips

Use `NewConfig` struct from `config.go` to load a configuration file into your application.
