# 🛡️ Go Backup Tool

[![Thai](https://img.shields.io/badge/lang-th-red.svg)](READMEth.md)
[![English](https://img.shields.io/badge/lang-en-blue.svg)](README.md)


A lightweight, high-performance command-line utility for automated file backups, built with Go. This tool streamlines the process of zipping folders while maintaining data integrity and keeping your storage organized.

## ✨ Key Features

- ⚡ **Lightning Fast:** Optimized for speed using Go's efficient file handling.
- 🧹 **Smart Filter:** Automatically skips unnecessary files/folders like `.git`, `node_modules`, and binaries to keep backups lean.
- 📍 **Custom Destination:** Choose exactly where you want to save your backup files (e.g., External Drive, Desktop, or Local folders).
- 📅 **Dynamic Naming:** Automatically generates filenames based on the source folder and the exact timestamp.
- 🔒 **Integrity Verification:** Every backup is finalized with a **SHA-256 Checksum** to ensure your data is perfectly preserved.

## 🚀 Getting Started

### Prerequisites
- [Go](https://go.dev/doc/install) (version 1.16 or higher) installed on your machine.

### Installation & Run

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/punnapob101/go-backup-tool.git](https://github.com/punnapob101/go-backup-tool.git)
   cd go-backup-tool
