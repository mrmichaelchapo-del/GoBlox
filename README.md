# GoBlox 🐿️

> A block-based desktop app platform — XML blocks, a Go backend, and a WebGL/Electron frontend.

![status](https://img.shields.io/badge/status-experimental-yellow)
![license](https://img.shields.io/badge/license-Apache--2.0-blue)
![blocks](https://img.shields.io/badge/blocks-16-brightgreen)
![backend](https://img.shields.io/badge/backend-Go-00ADD8)

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