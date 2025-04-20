````markdown
# Todo App

A simple task-tracking application built with Go, Fiber and HTMX, demonstrating OOP principles and clean separation of concerns.

---

## 📋 Features

- Create, list and delete todos via HTMX‑powered UI
- In‑memory repository (swappable for a real database)
- Fiber HTTP server with clean handler–service–repository layers

---

## ⚙️ Tech Stack

- **Language:** Go
- **Web Framework:** Fiber
- **Frontend:** HTMX + Go templates
- **Styling:** CSS (static)

---

### Installation

1. Clone the repo
   ```bash
   git clone https://github.com/your‑username/todo-app.git
   cd todo-app
   ```
````

2. Install dependencies

   ```bash
   go mod download
   ```

3. Start the server

   ```bash
   go run cmd/server/main.go
   ```

4. Open your browser at `http://localhost:3000`

---

## 📌 API Endpoints

| Method | Route        | Description         |
| ------ | ------------ | ------------------- |
| GET    | `/todos`     | List all todos      |
| POST   | `/todos`     | Create a new todo   |
| DELETE | `/todos/:id` | Delete a todo by ID |

---

## 🤝 Contributing

1. Fork the repo
2. Create a feature branch (`git checkout -b feature/XYZ`)
3. Commit your changes (`git commit -m "…"`), push, and open a PR
4. Ensure tests pass and code is linted

---

## 📜 License

This project is licensed under the MIT License.

```

Let me know if you’d like any changes or additions!
```
