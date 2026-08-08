# ProtonBox

Gerencie e combine comandos úteis do Proton (Proton padrão, Proton-GE e Proton-CachyOS) em um app gráfico simples. Clique em um comando para ver o que ele faz, marque vários para montar uma combinação pronta para colar nas opções de inicialização do Steam.

Manage and combine useful Proton launch commands (standard Proton, Proton-GE and Proton-CachyOS) in a simple GUI app. Click a command to see what it does, check several to build a combination ready to paste into Steam launch options.

## Funcionalidades / Features

- 44+ comandos com descrição em **Português** e **English**
- Copiar comando individual ou **combinar vários** em um único comando (`mangohud gamemoderun PROTON_LOG=1 %command%`)
- Busca por nome/categoria
- Tema claro/escuro/sistema
- Botão "Copiar ao clicar" opcional

## Como rodar / Running

```bash
./protonbox
```

Ou compile: `go build -o protonbox .`

## Downloads / Releases

- **AppImage** (Linux, qualquer distro)
- **.deb** (Debian/Ubuntu)
- **.rpm** (Fedora/openSUSE)
- **Binário** (Linux x86_64)

Disponíveis na página de [Releases](https://github.com/LucianoSkx/protonbox/releases).

## Builds locais / Building

```bash
./build.sh
```

Gera `protonbox` (binário), `ProtonBox-x86_64.AppImage`, `protonbox_<versão>_amd64.deb` e `protonbox-<versão>.x86_64.rpm`.

## Licença / License

MIT — veja [LICENSE](LICENSE).
