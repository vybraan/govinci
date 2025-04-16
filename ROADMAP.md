# Govinci Roadmap

> A native UI runtime for Go — declarative, composable, and portable.

---

## ✅ Done

### 🎯 Core Engine
- [x] Declarative component system in idiomatic Go
- [x] `View` interface & `ComponentFunc` for composable UIs
- [x] Reconciler with diffing and patching system
- [x] `NewState` for local stateful logic
- [x] `If`, `Match`, `For` for conditional and iterative rendering
- [x] Style system with reusable `StyleProps`
- [x] `Box`, `Text`, `Image`, `Button`, `Input`, `Spacer`, etc.
- [x] Hooks: `UseInterval`, `UseTimeout`, `UseMemo`, `UseReducer`
- [x] WASM runtime (`govinci-runtime.js`) with event bridge

### 🧪 Layout & Styling
- [x] `FlexRow`, `FlexColumn`, `Gap`, `Align`, `Justify`, `ZIndex`
- [x] `PositionSticky`, `Absolute`, `Relative`, etc.
- [x] Responsive layouts via style merging
- [x] Shadow, border, radius, hover styles

### 🧠 Developer Experience
- [x] One-line hot reload with patch diffing
- [x] Error boundaries and safe rendering fallback
- [x] Component-scoped rendering paths with `data-node-path`
- [x] Internal path-based rendering IDs for patches
- [x] Logging and inspection of patches


### 🌐 Web Support
- [x] WASM compilation with `main.wasm`
- [x] HTML + JS runtime to mount and patch views
- [x] WebAssembly event bridge to Go via `window.GoInvokeCallback`

---

## 🧩 In Progress

### 🔧 Core Abstractions
- [ ] Children-aware `Context` to preserve subcomponent state
- [ ] Navigation system (`Push`, `Pop`, `Reset`) stack-safe
- [ ] Theming system (`Theme{}` with ColorPalette, Typography)

### 📱 Native Runtime Bridges 
- [ ] **Android Runtime** (Go → JSON → Swift/Java UI Renderer)

### 🧰 UI DSL
- [ ] Real-world design system demo (Google-like, Apple-like, Flat)
- [ ] Components like `Tabs`, `Modal`, `Snackbar`, `Avatar`, `Badge`
- [ ] Forms with validation
- [ ] Keyboard-aware scroll area for mobile

### 📦 Packaging
- [ ] `govinci build --target=wasm`
- [ ] `govinci build --target=android`
- [ ] `govinci build --target=ios`

---

## 🔜 Planned


### Native Bridge (Planned for Android/iOS)

- [ ] Camera: `CameraView`, capture event
- [ ] Keystore (Secure): `Keystore.Save()`, `Keystore.Get()`
- [ ] Device Storage (Plain): `DeviceStorage.Set()`, `DeviceStorage.Get()`
- [ ] Bluetooth: `Scan`, `Connect`, `Send`
- [ ] Location / GPS
- [ ] FaceID / Biometric authentication
- [ ] Contacts access

### 📱 Native Runtime Bridges

- [ ] **iOS Runtime** (via Swift bridge)
- [ ] **Still thinking** for desktop

### 🛠️ DevTools
- [ ] State inspector overlay (similar to React DevTools)
- [ ] Visual patch viewer
- [ ] Hot module replacement

### 🧪 Testing & Perf
- [ ] Benchmark diff/patch engine
- [ ] Snapshot testing for views
- [ ] Latency profiling in runtime patching

### 🧬 Extensions
- [ ] Animations & transitions
- [ ] Router-style navigation for web
- [ ] Accessibility & keyboard navigation

---

## 🌍 Vision

> Build **native experiences** using only Go.

---

## 💬 Contribute

Open to collaborators and contributors.
Join the discussion or check the GitHub repo.

---

