# Govinci

**Govinci** is a fully idiomatic Go framework for building native mobile apps using a declarative, functional DSL. Designed entirely in Go — Govinci offers a new approach to mobile development where UI, logic, and state management are written in pure Go, and rendered natively on Android and iOS.

---

## ✨ Features

- **Declarative Syntax** – Compose views with pure functions and fluent props
- **Native Rendering** – Output native components on Android/iOS 
- **Component-Based** – Build custom views by composing smaller ones
- **Styling System** – Functional styling with support for themes and inheritance
- **State Management** – Built-in state system inspired by hooks
- **Event Handling** – Built-in callback registry for interactions
- **Theming & Tokens** – Define centralized visual identity and reusable design primitives
- **Bridge-Free Events** – Events and hardware calls require no manual bridge setup
- **App Config Injection** – Provide global config for name, author, version, locale

---

## 📦 Example

A **simple social network profile screen**, broken into components:

### `main.go`
```go
package main

import (
    "fmt"
    "govinci/core"
)

func main() {
    ctx := core.NewContext().WithConfig(&core.AppConfig{
        Name: "LetsBe Social",
        Author: "Ismael GraHms",
        Version: "0.1.0",
        Locale: "en-MZ"
    })

    App(ctx).Render(ctx)
    
}
```

### `app.go`
```go
func App(ctx *core.Context) core.View {
    return core.SafeArea(
        core.Scroll(
            core.Column(
                ProfileHeader(),
                core.Spacer(16),
                ProfileStats(),
                core.Spacer(12),
                PostList(),
            ),
        ),
    )
}
```

### `profile.go`
```go
func ProfileHeader() core.View {
    return core.Column(
        core.Image("https://example.com/avatar.jpg", core.UseStyle(core.Style{BorderRadius: 40})),
        core.Text("Ismael GraHms", core.FontSize(20), core.FontWeight(core.Bold)),
        core.Text("Software Engineer • Maputo"),
    )
}

func ProfileStats() core.View {
    return core.Row(
        Stat("Posts", "128"),
        core.Spacer(12),
        Stat("Followers", "1.2k"),
        core.Spacer(12),
        Stat("Following", "180"),
    )
}

func Stat(label, value string) core.View {
    return core.Column(
        core.Text(value, core.FontWeight(core.Bold)),
        core.Text(label, core.TextColor("#888")),
    )
}
```

### `posts.go`
```go
func PostList() core.View {
    return core.Column(
        Post("Enjoying the Govinci project! 🚀"),
        core.Spacer(8),
        Post("Working on UI DSLs in Go is pure joy.", "#golang #ux #native"),
    )
}

func Post(content string, tags ...string) core.View {
    full := content
    if len(tags) > 0 {
        full += "\n" + tags[0]
    }
    return core.Card(
        core.Column(
            core.Text(full),
            core.Spacer(4),
            core.Row(
                core.Button("Like"),
                core.Spacer(4),
                core.Button("Comment"),
            ),
        ),
    )
}
```

---

## 📐 Architecture

- `core/` – core abstractions: Node, View, Context, State
- `android/` – native renderer for Android (Kotlin)
- `ios/` – native renderer for iOS (Swift or Kotlin Multiplatform)
- `examples/` – declarative UI demos in Go

---

## 📱 Renderers

Renderers are responsible for turning the abstract `Node` tree into real UI elements.

- Native Android using `FrameLayout`, `TextView`, `Button`, etc.
- Native iOS via `UIView`, `UILabel`, etc. (coming soon)
- HTML (optional, for export and dev tools)


---

## 🛠 Dev Experience

- Hot reload (planned)
- Custom DSLs and style tokens
- Testing helpers for views and events
- Code generation for component scaffolds (planned)

---

## 📃 License

MIT License © 2025 Ismael GraHms

