# GoBlox 🐿️

> A block-based desktop app platform — XML blocks, a Go backend, and a WebGL/Electron frontend.

![status](https://img.shields.io/badge/status-experimental-yellow)
![version](https://img.shields.io/badge/version-0.1.0-blue)
![license](https://img.shields.io/badge/license-Apache--2.0-blue)
![blocks](https://img.shields.io/badge/blocks-16-brightgreen)
![categories](https://img.shields.io/badge/categories-5-success)
![backend](https://img.shields.io/badge/backend-Go-00ADD8)
![desktop](https://img.shields.io/badge/desktop-Electron-47848F)
![renderer](https://img.shields.io/badge/renderer-WebGL-990000)
![audio](https://img.shields.io/badge/audio-Web%20Audio-FF6F00)
![manifest](https://img.shields.io/badge/manifest-JSON5-8A2BE2)
![blocks](https://img.shields.io/badge/blocks-XML-orange)
![schema](https://img.shields.io/badge/schema-JSON%20Schema-4B8BBE)
![PRs](https://img.shields.io/badge/PRs-welcome-brightgreen)
![made with](https://img.shields.io/badge/made%20with-%E2%9D%A4%EF%B8%8F-red)
![gopher](https://img.shields.io/badge/powered%20by-Go%20gopher%20%F0%9F%90%B9-00ADD8)

<p align="center">
  <img src="https://img.shields.io/badge/%F0%9F%90%BF-GoBlox-ffb6c1?style=for-the-badge" alt="GoBlox" />
  <img src="https://img.shields.io/badge/XML-Blocks-8A2BE2?style=for-the-badge" alt="XML Blocks" />
  <img src="https://img.shields.io/badge/WebGL-Renderer-990000?style=for-the-badge" alt="WebGL" />
  <img src="https://img.shields.io/badge/Go-Backend-00ADD8?style=for-the-badge" alt="Go Backend" />
</p>

---

## What is GoBlox?

GoBlox is a desktop programming environment where apps are built from
**XML blocks** — like building with Lego, but for real software. 🧱

Each block declares what it does, how it renders (WebGL), and which
category it belongs to. A **Go backend** validates and serves blocks,
while an **Electron desktop app** renders the result.

And yes — there's a Hamster. 🐹

---

## Project Structure

```
GoBlox 🐿️/
├── package.json5        # App manifest (JSON5)
├── schemas.json         # Merged block + config schemas (JSON)
├── NOTICE               # Apache 2.0 notice
├── LICENSE              # Apache License 2.0
├── README.md
├── CONTRIBUTING.md
├── blocks/              # Block XML definitions
│   ├── motion/          # 5 blocks
│   ├── looks/           # 2 blocks
│   ├── sound/           # 2 blocks
│   ├── opers/           # 6 blocks
│   └── events/          # 1 block
├── server/              # Go backend 🐹
└── ...                  # Electron desktop app (WebGL)
```

---

## The Block Library 🧱

| Category | Blocks | What they do |
|---|---|---|
| 🏃 **Motion** | `Movesteps`, `Turnblocks`, `Measureblocks`, `gotox`, `gotoy` | Move & turn the Hamster |
| 👁️ **Looks** | `show`, `hideblocks` | Show / hide the sprite |
| 🔊 **Sound** | `playsound`, `pausesound` | Play & pause audio |
| 🧮 **Operators** | `join`, `join3`, `struct`, `add`, `multiply`, `subtract` | Math & text |
| 🎬 **Events** | `whenclicked` | Start when ▶️ is clicked |

**16 blocks and counting!** 🎉

---

## Block Anatomy

Every block follows the same shape:

```xml
<block type="motion">
  <text>Move 10 steps</text>

  <!-- REQUIRED: the trust anchor 🔍 -->
  <mode chip="web" define=".motion"></mode>

  <behavior>
    <action>move</action>
    <target>hamster</target>
    <steps>10</steps>
    <renderer>webgl</renderer>
    <axis>x</axis>
  </behavior>

  <params>
    <param name="steps" type="number" value="10" min="1" max="1000" />
    <param name="axis" type="string" value="x" enum="x,y" />
  </params>
</block>
```

### The Trust Rule 🔍

`<mode>` is **required** on every block. It declares the block's
category and runtime target. **No valid `<mode>` = block is not loaded.**

```xml
<mode chip="web" define=".motion"></mode>
```

---

## How It Works

```
package.json5  ──►  appNeeds.schemas  ──►  schemas.json
                                               │
                                               ▼
                                    validates blocks/*/*.xml
                                               │
                                               ▼
                              Electron (WebGL) renders the Hamster 🐹
                                    + Go backend serves data 🐹
```

1. 📋 `package.json5` declares what the app needs
2. 📐 `schemas.json` defines what a valid block looks like
3. 🧱 Blocks in `blocks/` are validated against those schemas
4. 🖥️ Electron reads the blocks and renders them with **WebGL**
5. 🐹 The Go server handles validation, data, and logic

---

## Tech Stack

| Layer | Tech |
|---|---|
| Manifest | JSON5 |
| Schemas | JSON Schema (draft 2020-12) |
| Blocks | XML |
| Backend | Go 🐹 |
| Desktop | Electron 🪟 |
| Rendering | WebGL 🎨 |
| Audio | Web Audio API 🎹 |

---

## Getting Started

*(coming soon!)*

---

## Contributing

Want to help build GoBlox? Read [CONTRIBUTING.md](CONTRIBUTING.md) —
we're friendly here! 🐿️💕

---

## License

Licensed under the **Apache License 2.0**.
See [LICENSE](LICENSE) and [NOTICE](NOTICE).

---

<p align="center">Made with 💖 by MrMichaelChapoDel</p>