# 🧠 MindMapVault

MindMapVault is a terminal + browser-based **knowledge graph tool** where you can create and manage node-based concepts and relationships — just like mind maps, dependency graphs, or semantic networks.

Built using **Node.js + Express + TypeScript** with **Mermaid.js** for real-time graph rendering.

---

## ✨ Features

- ✅ Add **nodes** with custom labels
- ✅ Add **edges** between existing nodes
- ✅ Visualize entire map in-browser using Mermaid.js
- ✅ Fully in-memory: No DB needed
- ✅ RESTful API + CLI compatible
- ✅ Modular TypeScript structure
- ✅ Toggle between **form** + **live graph**

---

## 🚀 Run the Server

```bash
npm run dev
```

Then open in browser:  
👉 [http://localhost:4000/view](http://localhost:4000/view)

---

## 📡 API Usage

### ➕ Add a Node

```bash
curl -X POST http://localhost:4000/nodes   -H "Content-Type: application/json"   -d '{"id":"react","label":"ReactJS"}'
```

### 🔗 Add an Edge

```bash
curl -X POST http://localhost:4000/edges   -H "Content-Type: application/json"   -d '{"from":"react","to":"jsx"}'
```

### 🧠 Get Full Map

```bash
curl http://localhost:4000/map
```

---

## 🌐 Web UI Guide

Open [http://localhost:4000/view](http://localhost:4000/view)

1. **Add Node**  
   - Type: `Add Node`  
   - Node ID: `react`  
   - Label: `ReactJS`

2. **Add Edge**  
   - Type: `Add Edge`  
   - From: `react`  
   - To: `jsx`

3. Graph will re-render live.  
4. Mermaid-style diagram will show all connections.

---

## 🧠 Technology Stack

| Layer         | Tech                            |
|---------------|----------------------------------|
| Server        | Node.js + Express + TypeScript   |
| API Structure | REST (Modular routes)            |
| Storage       | In-memory GraphStore class       |
| Frontend      | HTML + JS (no framework)         |
| Graph         | Mermaid.js                       |

---

## 📌 Notes

- No database required (yet).
- Graph is stored in memory only. Restarting the server resets it.
- Built for fast prototyping and graph visualization.

---

## 🛠 Commands Summary

```bash
npm run dev         # Start server with ts-node-dev
npm run build       # Build project to dist/
npm start           # Run built version from dist/
```

---



## 🔗 Inspiration

Inspired by:
- Mermaid.js
- Graph theory
- React DevTools
- Prompt chaining for LLMs

---

## 👨‍💻 Author

**Eddy Mouity** – future framework builder, system engineer & AI-native dev.
