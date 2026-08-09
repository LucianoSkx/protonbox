<div align="center">
  <img src="assets/logo.png" alt="ProtonBox" width="400">
</div>

---

# Português 🇧🇷

Gerencie e combine comandos úteis do Proton (Proton padrão, Proton-GE e Proton-CachyOS) em um app gráfico simples. Clique em um comando para ver o que ele faz, marque vários para montar uma combinação pronta para colar nas opções de inicialização do Steam.

## Funcionalidades

- **44+ comandos** com descrição em Português e English
- **Combinação múltipla**: marque vários comandos e gere uma única linha pronta, por exemplo:
  `mangohud gamemoderun PROTON_LOG=1 %command%`
- Busca por comando, título, categoria e compatibilidade
- Tema claro / escuro / sistema
- Opção "copiar ao clicar"
- Seletor de idioma (PT/EN) no topo

## Como rodar

```bash
./protonbox
```

Ou compile do código-fonte:

```bash
go build -o protonbox .
```

## Dependências de compilação

Precisa de Go + `gcc` + bibliotecas do GLFW:

| Distro | Instalar |
|---|---|
| **Arch / CachyOS** | `sudo pacman -S base-devel go` |
| **Debian / Ubuntu** | `sudo apt install golang gcc libgl1-mesa-dev xorg-dev libwayland-dev` |
| **Fedora** | `sudo dnf install golang gcc mesa-libGL-devel libX11-devel libwayland-devel libXcursor-devel libXrandr-devel libXi-devel` |

O `build.sh` gera binário + AppImage sempre; o `.deb` precisa de `dpkg-dev` e o `.rpm` de `rpm-build` (ele pula o que não estiver instalado, sem falhar).

## Instalação

Escolha um dos pacotes na página de [Releases](https://github.com/LucianoSkx/protonbox/releases):

| Formato | Uso |
|---|---|
| **AppImage** | Qualquer distro: `chmod +x protonbox-x86_64.AppImage && ./protonbox-x86_64.AppImage` |
| **.deb** | Debian/Ubuntu: `sudo dpkg -i protonbox_0.1.0_amd64.deb` |
| **.rpm** | Fedora/openSUSE: `sudo rpm -i protonbox-0.1.0-1.x86_64.rpm` |
| **Binário** | Compile com `./build.sh` |

## Compilando os pacotes

Builds separados por formato:

```bash
./build-binary.sh     # binário
./build-appimage.sh   # AppImage (qualquer distro)
./build-deb.sh        # .deb (precisa dpkg-dev)
./build-rpm.sh        # .rpm (precisa rpm-build)
./build.sh            # todos de uma vez (pula o que faltar)
```

Artefatos em `dist/`. O `build.sh` chama os quatro — se alguma ferramenta faltar, ele avisa e continua com os demais.

## Licença

MIT — veja [LICENSE](LICENSE).

---

# English 🇺🇸

Manage and combine useful Proton launch commands (standard Proton, Proton-GE and Proton-CachyOS) in a simple GUI app. Click a command to see what it does, check several to build a combination ready to paste into Steam launch options.

## Features

- **44+ commands** with descriptions in English and Portuguese
- **Multiple combination**: check several commands and generate a single ready-to-paste line, e.g.:
  `mangohud gamemoderun PROTON_LOG=1 %command%`
- Search by command, title, category and compatibility
- Light / dark / system theme
- Optional "copy on click"
- Language selector (PT/EN) at the top

## Running

```bash
./protonbox
```

Or build from source:

```bash
go build -o protonbox .
```

## Build dependencies

You need Go + `gcc` + the GLFW libraries:

| Distro | Install |
|---|---|
| **Arch / CachyOS** | `sudo pacman -S base-devel go` |
| **Debian / Ubuntu** | `sudo apt install golang gcc libgl1-mesa-dev xorg-dev libwayland-dev` |
| **Fedora** | `sudo dnf install golang gcc mesa-libGL-devel libX11-devel libwayland-devel libXcursor-devel libXrandr-devel libXi-devel` |

`build.sh` always produces the binary + AppImage; `.deb` additionally needs `dpkg-dev` and `.rpm` needs `rpm-build` (missing tools are skipped gracefully).

## Installation

Pick a package from the [Releases](https://github.com/LucianoSkx/protonbox/releases) page:

| Format | Usage |
|---|---|
| **AppImage** | Any distro: `chmod +x protonbox-x86_64.AppImage && ./protonbox-x86_64.AppImage` |
| **.deb** | Debian/Ubuntu: `sudo dpkg -i protonbox_0.1.0_amd64.deb` |
| **.rpm** | Fedora/openSUSE: `sudo rpm -i protonbox-0.1.0-1.x86_64.rpm` |
| **Binary** | Build with `./build.sh` |

## Building packages

Separate scripts per format:

```bash
./build-binary.sh     # binary
./build-appimage.sh   # AppImage (any distro)
./build-deb.sh        # .deb (needs dpkg-dev)
./build-rpm.sh        # .rpm (needs rpm-build)
./build.sh            # everything at once (skips what's missing)
```

Artifacts land in `dist/`. `build.sh` runs all four — if a tool is missing it warns and continues with the rest.

## License

MIT — see [LICENSE](LICENSE).
