package main

type Localized struct {
	PT string
	EN string
}

type Command struct {
	Command     string
	CommandEN   string
	Title       Localized
	Category    Localized
	Compat      Localized
	Description Localized
}

func commands() []Command {
	return []Command{
		{
			Command: "PROTON_LOG=1 %command%",
			Title: Localized{
				PT: "Log detalhado do Proton",
				EN: "Detailed Proton log",
			},
			Category: Localized{
				PT: "Diagnóstico",
				EN: "Diagnostics",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Gera um arquivo de log detalhado (steam-<appid>.log) na sua pasta pessoal, para diagnosticar travamentos e erros de desempenho do jogo. O arquivo aparece depois de fechar o jogo.",
				EN: "Generates a detailed log file (steam-<appid>.log) in your home folder to help diagnose crashes and performance issues. The file appears after closing the game.",
			},
		},
		{
			Command: "PROTON_LOG=warn+pipewire,warn+mmdevapi %command%",
			Title: Localized{
				PT: "Log de problemas de áudio",
				EN: "Audio issue log",
			},
			Category: Localized{
				PT: "Diagnóstico",
				EN: "Diagnostics",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Variação do PROTON_LOG com canais de áudio, para diagnosticar problemas de som (clipping, estalo ou áudio ausente) no driver PipeWire. Não suba o nível para +pipewire completo: o tracing total atrapalha o timing do áudio e causa mais estalos.",
				EN: "PROTON_LOG variant with audio channels, to diagnose sound issues (clipping, crackling or missing audio) on the PipeWire driver. Don't go above warn+: full tracing perturbs audio timing and causes more crackling.",
			},
		},
		{
			Command: "PROTON_LOG_DIR=~/proton-logs %command%",
			Title: Localized{
				PT: "Pasta dos logs",
				EN: "Log folder",
			},
			Category: Localized{
				PT: "Diagnóstico",
				EN: "Diagnostics",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Redireciona os logs do Proton para a pasta indicada em vez da pasta pessoal. Útil para manter os logs organizados por jogo sem poluir o home.",
				EN: "Redirects Proton logs to the given folder instead of your home directory. Useful to keep logs organized per game.",
			},
		},
		{
			Command: "PROTON_DUMP_DEBUG_COMMANDS=1 %command%",
			Title: Localized{
				PT: "Scripts de depuração",
				EN: "Debug scripts",
			},
			Category: Localized{
				PT: "Diagnóstico",
				EN: "Diagnostics",
			},
			Compat: Localized{
				PT: "GE e CachyOS",
				EN: "GE and CachyOS",
			},
			Description: Localized{
				PT: "Ao rodar o jogo, o Proton grava scripts de depuração (os mesmos comandos usados para iniciar o jogo) em $PROTON_DEBUG_DIR/proton_$USER (padrão: /tmp). Ótimo para entender como o jogo foi lançado e reproduzir manualmente.",
				EN: "When running the game, Proton writes debug scripts (the same commands used to launch the game) to $PROTON_DEBUG_DIR/proton_$USER (default: /tmp). Great to understand how the game was launched and reproduce it manually.",
			},
		},
		{
			Command: "PROTON_CRASH_REPORT_DIR=~/crash-reports %command%",
			Title: Localized{
				PT: "Relatórios de crash",
				EN: "Crash reports",
			},
			Category: Localized{
				PT: "Diagnóstico",
				EN: "Diagnostics",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Grava logs de crash na pasta indicada. Atenção: não limpa logs antigos, então pode encher o disco se não for monitorado.",
				EN: "Writes crash logs to the given folder. Warning: it does not clean old logs, so it can fill your disk if not monitored.",
			},
		},
		{
			Command: "DXVK_HUD=fps %command%",
			Title: Localized{
				PT: "HUD de FPS (DXVK)",
				EN: "FPS HUD (DXVK)",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Mostra na tela apenas os quadros por segundo (FPS) do jogo. É o overlay mais leve, ideal para um teste rápido sem poluir a tela.",
				EN: "Shows only the frames per second (FPS) on screen. The lightest overlay, ideal for a quick test without cluttering the screen.",
			},
		},
		{
			Command: "DXVK_HUD=fps,gpu,api,devinfo %command%",
			Title: Localized{
				PT: "HUD completo (DXVK)",
				EN: "Full HUD (DXVK)",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Painel detalhado com FPS, uso da GPU, API gráfica em uso (Vulkan/D3D) e informações do driver. Combina as métricas: fps, gpu, api, devinfo, version, drawcalls, memory. Separe os itens com vírgula.",
				EN: "Detailed panel with FPS, GPU usage, graphics API in use (Vulkan/D3D) and driver info. Combine metrics: fps, gpu, api, devinfo, version, drawcalls, memory. Separate items with commas.",
			},
		},
		{
			Command: "DXVK_HUD=version %command%",
			Title: Localized{
				PT: "Versão do DXVK",
				EN: "DXVK version",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Exibe a versão do DXVK em uso no canto da tela. Útil para confirmar qual build de DXVK o Proton está carregando antes de reportar problemas.",
				EN: "Shows the DXVK version in use in the corner of the screen. Useful to confirm which DXVK build Proton is loading before reporting issues.",
			},
		},
		{
			Command: "mangohud %command%",
			Title: Localized{
				PT: "MangoHud",
				EN: "MangoHud",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos (exige mangohud instalado)",
				EN: "All (requires mangohud)",
			},
			Description: Localized{
				PT: "Painel avançado na tela com métricas completas: uso de CPU, GPU, memória RAM, temperaturas, frequências e FPS. Instale com: sudo pacman -S mangohud. Configurações em ~/.config/MangoHud/MangoHud.conf.",
				EN: "Advanced on-screen panel with full metrics: CPU/GPU usage, RAM, temperatures, frequencies and FPS. Install with: sudo pacman -S mangohud. Settings in ~/.config/MangoHud/MangoHud.conf.",
			},
		},
		{
			Command: "gamescope -w 1920 -h 1080 -r 60 %command%",
			Title: Localized{
				PT: "Gamescope",
				EN: "Gamescope",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos (exige gamescope instalado)",
				EN: "All (requires gamescope)",
			},
			Description: Localized{
				PT: "Compositor micro do Steam/Valve que isola o jogo e permite upscaling, limite de FPS, FSR e janela redimensionável. Ajuste -w/-h para a resolução interna do jogo e -r para o limite de FPS.",
				EN: "Valve's micro compositor that isolates the game and allows upscaling, FPS limit, FSR and a resizable window. Adjust -w/-h for the game's internal resolution and -r for the FPS limit.",
			},
		},
		{
			Command: "gamemoderun %command%",
			Title: Localized{
				PT: "GameMode",
				EN: "GameMode",
			},
			Category: Localized{
				PT: "Overlay e desempenho",
				EN: "Overlay & Performance",
			},
			Compat: Localized{
				PT: "Todos (exige gamemode instalado)",
				EN: "All (requires gamemode)",
			},
			Description: Localized{
				PT: "Ativa o GameMode (da Feral Interactive), que aplica otimizações temporárias de CPU/GPU enquanto o jogo roda. Instale com: sudo pacman -S gamemode lib32-gamemode.",
				EN: "Enables GameMode (by Feral Interactive), applying temporary CPU/GPU optimizations while the game runs. Install with: sudo pacman -S gamemode lib32-gamemode.",
			},
		},
		{
			Command: "WINE_FULLSCREEN_FSR=1 %command%",
			Title: Localized{
				PT: "FidelityFX Super Resolution (FSR)",
				EN: "FidelityFX Super Resolution (FSR)",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "Todos (Vulkan via DXVK)",
				EN: "All (Vulkan via DXVK)",
			},
			Description: Localized{
				PT: "Ativa o upscaling FSR da AMD no modo fullscreen: o jogo renderiza em resolução menor e o FSR sobe a imagem, ganhando FPS em placas mais fracas. Combine com WINE_FULLSCREEN_FSR_STRENGTH=2 para ajustar a nitidez (0 = máximo, 5 = mínimo). Só funciona em jogos Vulkan.",
				EN: "Enables AMD FSR upscaling in fullscreen: the game renders at a lower resolution and FSR upscales the image, gaining FPS on weaker GPUs. Combine with WINE_FULLSCREEN_FSR_STRENGTH=2 for sharpness (0 = max, 5 = min). Only works in Vulkan games.",
			},
		},
		{
			Command: "WINE_FULLSCREEN_INTEGER_SCALING=1 %command%",
			Title: Localized{
				PT: "Escala inteira (pixels nítidos)",
				EN: "Integer scaling (sharp pixels)",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Ativa a escala inteira em fullscreen: pixels nítidos e quadradinhos ao subir a resolução, sem blur. Útil em jogos antigos ou pixel art.",
				EN: "Enables integer scaling in fullscreen: sharp, blocky pixels when upscaling, no blur. Useful in old or pixel-art games.",
			},
		},
		{
			Command: "PROTON_FSR4_UPGRADE=1 %command%",
			Title: Localized{
				PT: "Upgrade FSR 3.1 para FSR 4",
				EN: "FSR 3.1 to FSR 4 upgrade",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "GE e CachyOS (GPU AMD)",
				EN: "GE and CachyOS (AMD GPU)",
			},
			Description: Localized{
				PT: "Baixa automaticamente a amdxcffx64.dll e atualiza jogos com FSR 3.1 para FSR 4, com melhor qualidade de imagem. Dá para fixar a versão: PROTON_FSR4_UPGRADE=\"4.0.2\". Só funciona em GPUs AMD compatíveis (RDNA3+).",
				EN: "Automatically downloads amdxcffx64.dll and upgrades games with FSR 3.1 to FSR 4, with better image quality. You can pin the version: PROTON_FSR4_UPGRADE=\"4.0.2\". Only works on compatible AMD GPUs (RDNA3+).",
			},
		},
		{
			Command: "PROTON_DLSS_UPGRADE=1 %command%",
			Title: Localized{
				PT: "Upgrade do DLSS",
				EN: "DLSS upgrade",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "GE e CachyOS (GPU NVIDIA)",
				EN: "GE and CachyOS (NVIDIA GPU)",
			},
			Description: Localized{
				PT: "Baixa automaticamente versões mais novas das DLLs nvngx_dlss e substitui as do jogo, melhorando qualidade e desempenho do DLSS. Fixe a versão com PROTON_DLSS_UPGRADE=\"310.2\".",
				EN: "Automatically downloads newer nvngx_dlss DLLs and replaces the game's, improving DLSS quality and performance. Pin the version with PROTON_DLSS_UPGRADE=\"310.2\".",
			},
		},
		{
			Command: "PROTON_XESS_UPGRADE=1 %command%",
			Title: Localized{
				PT: "Upgrade do XeSS",
				EN: "XeSS upgrade",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "GE e CachyOS (GPU Intel)",
				EN: "GE and CachyOS (Intel GPU)",
			},
			Description: Localized{
				PT: "Baixa automaticamente a DLL do XeSS (Intel) e atualiza jogos para a versão mais recente do upscaler.",
				EN: "Automatically downloads the XeSS (Intel) DLL and upgrades games to the latest upscaler version.",
			},
		},
		{
			Command: "PROTON_USE_WINED3D=1 %command%",
			Title: Localized{
				PT: "Forçar OpenGL (wined3d)",
				EN: "Force OpenGL (wined3d)",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Força o Proton a usar a tradução OpenGL do Wine (wined3d) em vez do Vulkan (DXVK). Geralmente perde desempenho; só útil para testes em placas muito antigas sem suporte Vulkan.",
				EN: "Forces Proton to use Wine's OpenGL translation (wined3d) instead of Vulkan (DXVK). Usually loses performance; only useful for testing on very old GPUs without Vulkan support.",
			},
		},
		{
			Command: "PROTON_DXVK_SAREK=1 %command%",
			Title: Localized{
				PT: "DXVK-Sarek (placas antigas)",
				EN: "DXVK-Sarek (old GPUs)",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Usa o fork dxvk-sarek do DXVK, feito para GPUs antigas que só suportam Vulkan 1.1/1.2 (em vez de 1.3). Usa o branch async e NÃO deve ser usado em jogos multiplayer ou com anti-cheat.",
				EN: "Uses the dxvk-sarek DXVK fork, made for old GPUs that only support Vulkan 1.1/1.2 (instead of 1.3). Uses the async branch and MUST NOT be used in multiplayer or anti-cheat games.",
			},
		},
		{
			Command: "PROTON_DXVK_LOWLATENCY=1 %command%",
			Title: Localized{
				PT: "DXVK de baixa latência",
				EN: "Low latency DXVK",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Usa o fork dxvk-low-latency, que adiciona frame pacing de baixa latência: melhora a responsividade do jogo (input lag) e a estabilidade da latência ao longo do tempo.",
				EN: "Uses the dxvk-low-latency fork, adding low-latency frame pacing: improves game responsiveness (input lag) and latency stability over time.",
			},
		},
		{
			Command: "DXVK_HDR=1 %command%",
			Title: Localized{
				PT: "HDR via DXVK",
				EN: "HDR via DXVK",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "CachyOS (exige monitor e compositor com HDR)",
				EN: "CachyOS (requires HDR monitor and compositor)",
			},
			Description: Localized{
				PT: "Ativa HDR nos jogos via DXVK. Em NVIDIA com drivers mais antigos, combine com ENABLE_HDR_WSI=1. Desative com DXVK_NO_HDR=1 se o HDR automático estiver causando problemas.",
				EN: "Enables HDR in games via DXVK. On NVIDIA with older drivers, combine with ENABLE_HDR_WSI=1. Disable with DXVK_NO_HDR=1 if automatic HDR is causing issues.",
			},
		},
		{
			Command: "PROTON_NO_D3D11=1 %command%",
			Title: Localized{
				PT: "Desabilitar D3D11",
				EN: "Disable D3D11",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Desativa a d3d11.dll. Só serve para jogos que rodam melhor caindo para o fallback D3D9. Se o jogo não tiver fallback, ele simplesmente não abre.",
				EN: "Disables d3d11.dll. Only useful for games that run better falling back to D3D9. If the game has no fallback, it simply won't start.",
			},
		},
		{
			Command: "PROTON_NO_ESYNC=1 %command%",
			Title: Localized{
				PT: "Desabilitar ESync",
				EN: "Disable ESync",
			},
			Category: Localized{
				PT: "Sincronização",
				EN: "Synchronization",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Desativa a sincronização por eventfd (ESync). Use como teste se o jogo estiver travando ou instável, para ver se a sincronização é a culpada.",
				EN: "Disables eventfd-based synchronization (ESync). Use as a test if the game is crashing or unstable, to see if sync is the culprit.",
			},
		},
		{
			Command: "PROTON_NO_FSYNC=1 %command%",
			Title: Localized{
				PT: "Desabilitar FSync",
				EN: "Disable FSync",
			},
			Category: Localized{
				PT: "Sincronização",
				EN: "Synchronization",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Desativa a sincronização por futex (FSync). Útil para testar instabilidade ou incompatibilidade de um jogo com o FSync. (Em sistemas sem suporte a FUTEX_WAIT_MULTIPLE, já é desativado sozinho.)",
				EN: "Disables futex-based synchronization (FSync). Useful to test instability or incompatibility of a game with FSync. (On systems without FUTEX_WAIT_MULTIPLE support, it's already disabled.)",
			},
		},
		{
			Command: "PROTON_ENABLE_NVAPI=1 %command%",
			Title: Localized{
				PT: "NVAPI da NVIDIA",
				EN: "NVIDIA NVAPI",
			},
			Category: Localized{
				PT: "GPU",
				EN: "GPU",
			},
			Compat: Localized{
				PT: "GE (GPU NVIDIA)",
				EN: "GE (NVIDIA GPU)",
			},
			Description: Localized{
				PT: "Habilita a biblioteca NVAPI da NVIDIA dentro do Proton, ativando funcionalidades proprietárias (como DLSS em alguns jogos e recursos específicos do driver). Desative com PROTON_DISABLE_NVAPI=1 se causar problemas.",
				EN: "Enables NVIDIA's NVAPI library inside Proton, activating proprietary features (like DLSS in some games and driver-specific features). Disable with PROTON_DISABLE_NVAPI=1 if it causes issues.",
			},
		},
		{
			Command: "DRI_PRIME=1 %command%",
			Title: Localized{
				PT: "Forçar GPU dedicada",
				EN: "Force discrete GPU",
			},
			Category: Localized{
				PT: "GPU",
				EN: "GPU",
			},
			Compat: Localized{
				PT: "Todos (laptops híbridos)",
				EN: "All (hybrid laptops)",
			},
			Description: Localized{
				PT: "Em laptops com GPU integrada + dedicada, força o jogo a rodar na placa de vídeo dedicada. Em alguns sistemas, use o valor com o índice da GPU: DRI_PRIME=pci-0000_01_00_0.",
				EN: "On laptops with integrated + discrete GPUs, forces the game to run on the dedicated GPU. On some systems, use the GPU index value: DRI_PRIME=pci-0000_01_00_0.",
			},
		},
		{
			Command: "PROTON_ENABLE_WAYLAND=1 %command%",
			Title: Localized{
				PT: "Driver nativo Wayland",
				EN: "Native Wayland driver",
			},
			Category: Localized{
				PT: "Wayland",
				EN: "Wayland",
			},
			Compat: Localized{
				PT: "GE e CachyOS (experimental)",
				EN: "GE and CachyOS (experimental)",
			},
			Description: Localized{
				PT: "Ativa o driver winewayland (janela nativa em Wayland, sem XWayland). É experimental: launchers Electron (Battle.net, EA App, Ubisoft Connect) podem abrir janela branca; nesse caso tente adicionar --in-process-gpu aos argumentos. Se houver problema com controle, use PROTON_USE_SDL=1.",
				EN: "Enables the winewayland driver (native Wayland window, no XWayland). It's experimental: Electron launchers (Battle.net, EA App, Ubisoft Connect) may show a white window; try adding --in-process-gpu to the arguments then. If you have controller issues, use PROTON_USE_SDL=1.",
			},
		},
		{
			Command: "PROTON_USE_PIPEWIRE=0 %command%",
			Title: Localized{
				PT: "Trocar driver de áudio",
				EN: "Switch audio driver",
			},
			Category: Localized{
				PT: "Áudio",
				EN: "Audio",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Desativa o driver de áudio winepipewire (ligado por padrão no proton-cachyos) e volta para o winepulse. Útil quando o jogo tem áudio estalando ou cortando. Também dá para escolher o driver com WINE_AUDIO_DRIVER.",
				EN: "Disables the winepipewire audio driver (default in proton-cachyos) and falls back to winepulse. Useful when a game's audio crackles or cuts out. You can also pick the driver with WINE_AUDIO_DRIVER.",
			},
		},
		{
			Command: "WINE_AUDIO_DRIVER=pulse %command%",
			Title: Localized{
				PT: "Escolher driver de áudio",
				EN: "Choose audio driver",
			},
			Category: Localized{
				PT: "Áudio",
				EN: "Audio",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Define qual driver de áudio o Wine usa. Padrão: pipewire,pulse,alsa. Exemplo: WINE_AUDIO_DRIVER=pulse força apenas o winepulse.drv.",
				EN: "Sets which audio driver Wine uses. Default: pipewire,pulse,alsa. Example: WINE_AUDIO_DRIVER=pulse forces winepulse.drv only.",
			},
		},
		{
			Command: "PROTON_USE_SDL=1 %command%",
			Title: Localized{
				PT: "Input via SDL",
				EN: "SDL input",
			},
			Category: Localized{
				PT: "Input",
				EN: "Input",
			},
			Compat: Localized{
				PT: "GE e CachyOS",
				EN: "GE and CachyOS",
			},
			Description: Localized{
				PT: "Usa o input do SDL em vez de HIDRAW/Steam Input. Resolve problemas de controle que não é detectado ou se comporta mal (comum com o driver Wayland).",
				EN: "Uses SDL input instead of HIDRAW/Steam Input. Fixes controllers that are not detected or behave badly (common with the Wayland driver).",
			},
		},
		{
			Command:   "HOST_LC_ALL=pt_BR.UTF-8 %command%",
			CommandEN: "HOST_LC_ALL=en_US.UTF-8 %command%",
			Title: Localized{
				PT: "Idioma do jogo",
				EN: "Game language",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Força um locale específico para o jogo, sobrescrevendo todos os outros ajustes de idioma. Troque pt_BR.UTF-8 pelo locale desejado (ex.: en_US.UTF-8) para jogos que pegam o idioma errado.",
				EN: "Forces a specific locale for the game, overriding all other language settings. Replace pt_BR.UTF-8 with the desired locale (e.g.: en_US.UTF-8) for games that pick the wrong language.",
			},
		},
		{
			Command: "PROTON_USE_OPTISCALER=1 %command%",
			Title: Localized{
				PT: "OptiScaler",
				EN: "OptiScaler",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Ativa a injeção automática do OptiScaler, que permite usar FSR/DLSS/XeSS em jogos sem suporte nativo. Configure via PROTON_OPTISCALER_CONFIG=\"Upscalers.Dx11Upscaler=fsr31;Upscalers.Dx12Upscaler=dlss\". Em desenvolvimento: nem todos os jogos funcionam.",
				EN: "Enables automatic OptiScaler injection, letting you use FSR/DLSS/XeSS in games without native support. Configure via PROTON_OPTISCALER_CONFIG=\"Upscalers.Dx11Upscaler=fsr31;Upscalers.Dx12Upscaler=dlss\". Work in progress: not all games work.",
			},
		},
		{
			Command: "LOW_LATENCY_LAYER=1 %command%",
			Title: Localized{
				PT: "AMD Anti-Lag 2 (qualquer GPU)",
				EN: "AMD Anti-Lag 2 (any GPU)",
			},
			Category: Localized{
				PT: "Latência",
				EN: "Latency",
			},
			Compat: Localized{
				PT: "Todos (GPU AMD/Intel; requer low_latency_layer)",
				EN: "All (AMD/Intel GPU; requires low_latency_layer)",
			},
			Description: Localized{
				PT: "Ativa o low_latency_layer, que expõe a extensão VK_AMD_anti_lag em GPUs AMD e Intel — o Anti-Lag 2 passa a funcionar em jogos Vulkan (CS2 nativo, e via dxvk-nvapi em jogos Proton com proton-cachyos/GE, que já embutem o layer). Desative com DISABLE_LOW_LATENCY_LAYER=1.",
				EN: "Enables low_latency_layer, which exposes the VK_AMD_anti_lag extension on AMD and Intel GPUs — Anti-Lag 2 now works in Vulkan games (native CS2, and via dxvk-nvapi in Proton games with proton-cachyos/GE, which already bundle the layer). Disable with DISABLE_LOW_LATENCY_LAYER=1.",
			},
		},
		{
			Command: "LOW_LATENCY_LAYER=1 LOW_LATENCY_LAYER_REFLEX=1 DXVK_CONFIG=\"dxgi.hideAmdGpu = True\" %command%",
			Title: Localized{
				PT: "NVIDIA Reflex em qualquer GPU",
				EN: "NVIDIA Reflex on any GPU",
			},
			Category: Localized{
				PT: "Latência",
				EN: "Latency",
			},
			Compat: Localized{
				PT: "Todos (requer low_latency_layer; jogos com Reflex)",
				EN: "All (requires low_latency_layer; Reflex games)",
			},
			Description: Localized{
				PT: "Faz o low_latency_layer expor VK_NV_low_latency2 em vez de anti-lag, ativando o Reflex (mesmo desempenho do Anti-Lag 2, mas com suporte em muito mais jogos). Se o menu Reflex não aparecer, tente também PROTON_FORCE_NVAPI=1 (atenção: quebra o upgrade FSR4).",
				EN: "Makes low_latency_layer expose VK_NV_low_latency2 instead of anti-lag, enabling Reflex (same performance as Anti-Lag 2, but supported in far more games). If the Reflex menu doesn't appear, also try PROTON_FORCE_NVAPI=1 (warning: breaks FSR4 upgrade).",
			},
		},
		{
			Command: "WINE_FULLSCREEN_FSR_STRENGTH=2 %command%",
			Title: Localized{
				PT: "Nitidez do FSR",
				EN: "FSR sharpness",
			},
			Category: Localized{
				PT: "Upscaling",
				EN: "Upscaling",
			},
			Compat: Localized{
				PT: "Todos (use junto com WINE_FULLSCREEN_FSR=1)",
				EN: "All (use with WINE_FULLSCREEN_FSR=1)",
			},
			Description: Localized{
				PT: "Ajusta a nitidez do FSR em fullscreen: 0 = nitidez máxima, 5 = mínima (2 é o recomendado pela AMD). Sem essa variável, o FSR usa o valor padrão.",
				EN: "Adjusts FSR sharpness in fullscreen: 0 = maximum sharpness, 5 = minimum (2 is AMD's recommendation). Without this variable, FSR uses the default value.",
			},
		},
		{
			Command: "VKD3D_CONFIG=dxr %command%",
			Title: Localized{
				PT: "Ray tracing (DXR) via vkd3d",
				EN: "Ray tracing (DXR) via vkd3d",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Todos (GPU com suporte a ray tracing)",
				EN: "All (ray tracing capable GPU)",
			},
			Description: Localized{
				PT: "Ativa o suporte a ray tracing (DXR) nos jogos Direct3D 12 que usam vkd3d-proton. Requer GPU com ray tracing habilitado no driver. Outras opções podem ser separadas por vírgula (ex.: dxr,force_bindless_texel_buffer).",
				EN: "Enables ray tracing (DXR) support in Direct3D 12 games using vkd3d-proton. Requires a ray tracing capable GPU with RT enabled in the driver. Other options can be comma separated (e.g.: dxr,force_bindless_texel_buffer).",
			},
		},
		{
			Command: "WINEDLLOVERRIDES=\"dinput8=n,b\" %command%",
			Title: Localized{
				PT: "Overrides de DLL (WINEDLLOVERRIDES)",
				EN: "DLL overrides (WINEDLLOVERRIDES)",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "Todos",
				EN: "All",
			},
			Description: Localized{
				PT: "Controla de onde cada DLL do jogo vem: n = nativa (a do jogo, usada por mods/wrappers como o dinput8 de jogos antigos), b = builtin (a do Proton). Ex.: dinput8=n,b força a DLL do próprio jogo. Separe várias DLLs com ponto-e-vírgula.",
				EN: "Controls where each game DLL comes from: n = native (the game's own, used by mods/wrappers like dinput8 in old games), b = builtin (Proton's). E.g.: dinput8=n,b forces the game's own DLL. Separate multiple DLLs with semicolons.",
			},
		},
		{
			Command: "PROTON_NO_D3D10=1 %command%",
			Title: Localized{
				PT: "Desabilitar D3D10",
				EN: "Disable D3D10",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "GE e CachyOS",
				EN: "GE and CachyOS",
			},
			Description: Localized{
				PT: "Desativa a d3d10.dll e a dxgi.dll, para jogos D3D10 que conseguem cair para D3D9 com mais desempenho. Se o jogo não tiver fallback, não abre.",
				EN: "Disables d3d10.dll and dxgi.dll, for D3D10 games that can fall back to D3D9 with better performance. If the game has no fallback, it won't start.",
			},
		},
		{
			Command: "PROTON_NO_D3D9=1 %command%",
			Title: Localized{
				PT: "Desabilitar D3D9",
				EN: "Disable D3D9",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "GE e CachyOS",
				EN: "GE and CachyOS",
			},
			Description: Localized{
				PT: "Desativa a d3d9.dll. Raro de precisar; útil quando o D3D9 via DXVK causa problemas e o jogo tem outro caminho de renderização.",
				EN: "Disables d3d9.dll. Rarely needed; useful when D3D9 via DXVK causes issues and the game has another rendering path.",
			},
		},
		{
			Command: "RADV_DEBUG=nofastclears %command%",
			Title: Localized{
				PT: "RADV: corrigir artefatos (nofastclears)",
				EN: "RADV: fix artifacts (nofastclears)",
			},
			Category: Localized{
				PT: "GPU",
				EN: "GPU",
			},
			Compat: Localized{
				PT: "Todos (GPU AMD / Mesa RADV)",
				EN: "All (AMD GPU / Mesa RADV)",
			},
			Description: Localized{
				PT: "Desativa os fast clears no driver RADV, corrigindo artefatos visuais (tela piscando, linhas estranhas) em alguns jogos AMD. Se o jogo sumir no HUD, é sintoma de fast clear.",
				EN: "Disables fast clears in the RADV driver, fixing visual artifacts (flickering, weird lines) in some AMD games. If a game disappears from the HUD, it's a fast clear symptom.",
			},
		},
		{
			Command: "DRI_CONFIG=\"radv_invariant_geom=true\" %command%",
			Title: Localized{
				PT: "RADV: geometria invariante (driconf)",
				EN: "RADV: invariant geometry (driconf)",
			},
			Category: Localized{
				PT: "GPU",
				EN: "GPU",
			},
			Compat: Localized{
				PT: "Todos (GPU AMD / Mesa RADV)",
				EN: "All (AMD GPU / Mesa RADV)",
			},
			Description: Localized{
				PT: "Ativa a geometria invariante do RADV via driconf (substitui a antiga RADV_DEBUG=invariantgeom, removida no Mesa 26). Corrige cintilação/sumindo de objetos em jogos mal otimizados.",
				EN: "Enables RADV invariant geometry via driconf (replaces the old RADV_DEBUG=invariantgeom, removed in Mesa 26). Fixes flickering/disappearing objects in poorly optimized games.",
			},
		},
		{
			Command: "MESA_VK_WSI_PRESENT_MODE=mailbox %command%",
			Title: Localized{
				PT: "Mesa: modo de apresentação Vulkan",
				EN: "Mesa: Vulkan present mode",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Todos (Mesa Vulkan)",
				EN: "All (Mesa Vulkan)",
			},
			Description: Localized{
				PT: "Força o modo de apresentação do driver Mesa para todos os jogos Vulkan: mailbox (sem vsync, baixa latência, sem tearing quando o compositor suporta) ou immediate (vsync desligado de vez). Opções: fifo (vsync), mailbox, immediate.",
				EN: "Forces the Mesa driver's present mode for all Vulkan games: mailbox (no vsync, low latency, no tearing when the compositor supports it) or immediate (vsync fully off). Options: fifo (vsync), mailbox, immediate.",
			},
		},
		{
			Command: "__GL_SHADER_DISK_CACHE_SKIP_CLEANUP=1 %command%",
			Title: Localized{
				PT: "NVIDIA: cache de shaders sem limpeza",
				EN: "NVIDIA: shader cache without cleanup",
			},
			Category: Localized{
				PT: "GPU",
				EN: "GPU",
			},
			Compat: Localized{
				PT: "NVIDIA",
				EN: "NVIDIA",
			},
			Description: Localized{
				PT: "Impede o driver NVIDIA de limpar o cache de shaders ao sair do jogo, evitando stutter recorrente na primeira execução. Recomendado pela wiki do CachyOS. Lembre de limpar o cache manualmente de tempos em tempos.",
				EN: "Prevents the NVIDIA driver from cleaning the shader cache on game exit, avoiding recurring stutter on first runs. Recommended by the CachyOS wiki. Remember to clean the cache manually from time to time.",
			},
		},
		{
			Command: "PROTON_USE_WOW64=1 %command%",
			Title: Localized{
				PT: "Prefixos WOW64",
				EN: "WOW64 prefixes",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "GE e CachyOS",
				EN: "GE and CachyOS",
			},
			Description: Localized{
				PT: "Usa o modo wow64 (novo modelo de prefixo do Wine, 32 e 64 bits sem camada de tradução separada). Mais moderno e em alguns casos com melhor desempenho, mas exige criar um prefixo novo (trocar a versão do Proton no jogo).",
				EN: "Uses wow64 mode (Wine's new prefix model, 32 and 64 bit without a separate translation layer). More modern and sometimes faster, but requires a fresh prefix (switch the game's Proton version).",
			},
		},
		{
			Command: "PROTON_DISCORD_BRIDGE=1 %command%",
			Title: Localized{
				PT: "Rich Presence no Discord",
				EN: "Discord Rich Presence",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "CachyOS",
				EN: "CachyOS",
			},
			Description: Localized{
				PT: "Ativa o rpc-bridge, que permite jogos rodando no Proton exibirem Rich Presence (\"jogando X\") no Discord.",
				EN: "Enables rpc-bridge, letting games running in Proton show Rich Presence (\"playing X\") on Discord.",
			},
		},
		{
			Command:   "LSFG_PROCESS=steam %command%",
			CommandEN: "LSFG_PROCESS=steam %command%",
			Title: Localized{
				PT: "lsfg-vk: frame generation",
				EN: "lsfg-vk: frame generation",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Vulkan (Proton/GE)",
				EN: "Vulkan (Proton/GE)",
			},
			Description: Localized{
				PT: "Ativa a layer de frame generation do lsfg-vk usando o profile \"steam\". Crie o profile e ajuste multiplicador, modo de desempenho, HDR etc. na interface gráfica do lsfg-vk (comando lsfg-vk-ui). Requer lsfg-vk instalado e Lossless Scaling na Steam.",
				EN: "Enables the lsfg-vk frame generation layer using the \"steam\" profile. Create the profile and tune multiplier, performance mode, HDR etc. in the lsfg-vk GUI (lsfg-vk-ui command). Requires lsfg-vk installed and Lossless Scaling on Steam.",
			},
		},
		{
			Command:   "PROTON_WAIT_ATTACH=1 %command%",
			CommandEN: "PROTON_WAIT_ATTACH=1 %command%",
			Title: Localized{
				PT: "Esperar depurador anexar",
				EN: "Wait for debugger attach",
			},
			Category: Localized{
				PT: "Diagnóstico",
				EN: "Diagnostics",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Espera um depurador anexar ao steam.exe antes de iniciar o processo do jogo. Para anexar ao jogo no início, configure o depurador para seguir processos filhos.",
				EN: "Waits for a debugger to attach to steam.exe before launching the game process. To attach to the game at startup, set debuggers to follow child processes.",
			},
		},
		{
			Command:   "PROTON_DXVK_D3D8=1 %command%",
			CommandEN: "PROTON_DXVK_D3D8=1 %command%",
			Title: Localized{
				PT: "D3D8 via DXVK",
				EN: "D3D8 via DXVK",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Usa o d3d8.dll do DXVK em vez do d3d8 do Wine para jogos Direct3D 8.",
				EN: "Use DXVK's d3d8.dll instead of Wine's d3d8 for Direct3D 8 games.",
			},
		},
		{
			Command:   "PROTON_NO_NTSYNC=1 %command%",
			CommandEN: "PROTON_NO_NTSYNC=1 %command%",
			Title: Localized{
				PT: "Desativar ntsync",
				EN: "Disable ntsync",
			},
			Category: Localized{
				PT: "Sincronização",
				EN: "Synchronization",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Desativa o ntsync (sincronização no estilo Windows NT). Use se algum jogo apresentar problema com o ntsync ativo.",
				EN: "Disables ntsync (Windows NT style synchronization). Use it if a game misbehaves with ntsync enabled.",
			},
		},
		{
			Command:   "PROTON_FORCE_LARGE_ADDRESS_AWARE=1 %command%",
			CommandEN: "PROTON_FORCE_LARGE_ADDRESS_AWARE=1 %command%",
			Title: Localized{
				PT: "Forçar LARGE_ADDRESS_AWARE",
				EN: "Force LARGE_ADDRESS_AWARE",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Força a flag LARGE_ADDRESS_AWARE em todos os executáveis, permitindo usar mais de 2 GB de RAM (já habilitado por padrão no Proton).",
				EN: "Forces the LARGE_ADDRESS_AWARE flag on all executables, allowing more than 2 GB of RAM (enabled by default in Proton).",
			},
		},
		{
			Command:   "PROTON_HEAP_DELAY_FREE=1 %command%",
			CommandEN: "PROTON_HEAP_DELAY_FREE=1 %command%",
			Title: Localized{
				PT: "Atrasar liberação de memória",
				EN: "Delay memory freeing",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Atrasa a liberação de parte da memória para contornar bugs de use-after-free em alguns jogos.",
				EN: "Delays freeing some memory to work around use-after-free bugs in some games.",
			},
		},
		{
			Command:   "PROTON_SET_GAME_DRIVE=1 %command%",
			CommandEN: "PROTON_SET_GAME_DRIVE=1 %command%",
			Title: Localized{
				PT: "Unidade S: do jogo",
				EN: "Game S: drive",
			},
			Category: Localized{
				PT: "Outros",
				EN: "Other",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Cria uma unidade S: apontando para a biblioteca Steam que contém o jogo. Útil para jogos que procuram arquivos em caminhos fixos.",
				EN: "Creates an S: drive pointing to the Steam Library which contains the game. Useful for games that look for files in fixed paths.",
			},
		},
		{
			Command:   "PROTON_OLD_GL_STRING=1 %command%",
			CommandEN: "PROTON_OLD_GL_STRING=1 %command%",
			Title: Localized{
				PT: "String GL limitada",
				EN: "Limit GL extension string",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Aplica overrides no driver para limitar o tamanho da string de extensões GL, para jogos antigos que travam com strings muito longas.",
				EN: "Sets driver overrides to limit the length of the GL extension string, for old games that crash on very long extension strings.",
			},
		},
		{
			Command:   "WINE_DO_NOT_CREATE_DXGI_DEVICE_MANAGER=1 %command%",
			CommandEN: "WINE_DO_NOT_CREATE_DXGI_DEVICE_MANAGER=1 %command%",
			Title: Localized{
				PT: "Sem DXGI device manager",
				EN: "Skip DXGI device manager",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Hack para contornar problemas de vídeo em alguns jogos causados por suporte incompleto a IMFDXGIDeviceManager.",
				EN: "Hack to work around video issues in some games due to incomplete IMFDXGIDeviceManager support.",
			},
		},
		{
			Command:   "WINE_DISABLE_VULKAN_OPWR=1 %command%",
			CommandEN: "WINE_DISABLE_VULKAN_OPWR=1 %command%",
			Title: Localized{
				PT: "Sem Vulkan OPWR",
				EN: "Disable Vulkan OPWR",
			},
			Category: Localized{
				PT: "Wayland",
				EN: "Wayland",
			},
			Compat: Localized{
				PT: "Wayland",
				EN: "Wayland",
			},
			Description: Localized{
				PT: "Desativa o render de janelas de outros processos via Vulkan (other process window rendering), que às vezes causa atraso de um frame no Wayland.",
				EN: "Disables Vulkan other process window rendering, which sometimes causes issues on Wayland due to blit being one frame behind.",
			},
		},
		{
			Command:   "PROTON_HIDE_NVIDIA_GPU=1 %command%",
			CommandEN: "PROTON_HIDE_NVIDIA_GPU=1 %command%",
			Title: Localized{
				PT: "Ocultar GPU NVIDIA",
				EN: "Hide NVIDIA GPU",
			},
			Category: Localized{
				PT: "GPU",
				EN: "GPU",
			},
			Compat: Localized{
				PT: "NVIDIA",
				EN: "NVIDIA",
			},
			Description: Localized{
				PT: "Faz a GPU NVIDIA ser sempre reportada como AMD. Alguns jogos exigem isso quando dependem de funcionalidade do driver NVIDIA que só existe no Windows.",
				EN: "Forces NVIDIA GPUs to always be reported as AMD GPUs. Some games require this if they depend on Windows-only NVIDIA driver functionality.",
			},
		},
		{
			Command:   "WINE_USE_KWIN_HACKS=1 %command%",
			CommandEN: "WINE_USE_KWIN_HACKS=1 %command%",
			Title: Localized{
				PT: "Hacks para KDE",
				EN: "KDE windowing hacks",
			},
			Category: Localized{
				PT: "Wayland",
				EN: "Wayland",
			},
			Compat: Localized{
				PT: "KDE (Wayland/X11)",
				EN: "KDE (Wayland/X11)",
			},
			Description: Localized{
				PT: "Ativa hacks específicos do KDE que melhoram a experiência com KDE mais antigo que 6.4 no Wayland e 6.6 no X11.",
				EN: "Enables KDE-specific windowing hacks that may improve experience with KDE older than 6.4 on Wayland and 6.6 on X11.",
			},
		},
		{
			Command:   "PROTON_USE_XALIA=1 %command%",
			CommandEN: "PROTON_USE_XALIA=1 %command%",
			Title: Localized{
				PT: "Xalia (UI de gamepad)",
				EN: "Xalia (gamepad UI)",
			},
			Category: Localized{
				PT: "Input",
				EN: "Input",
			},
			Compat: Localized{
				PT: "Proton padrão, GE e CachyOS",
				EN: "Standard Proton, GE and CachyOS",
			},
			Description: Localized{
				PT: "Ativa o Xalia, que adiciona UI de gamepad para algumas interfaces de teclado/mouse. Por padrão o Proton decide dinamicamente; use 0 para desativar.",
				EN: "Enables Xalia, which adds a gamepad UI for some keyboard/mouse interfaces. The default is dynamic; set to 0 to disable.",
			},
		},
		{
			Command:   "FNA3D_FORCE_DRIVER=D3D11 %command%",
			CommandEN: "FNA3D_FORCE_DRIVER=D3D11 %command%",
			Title: Localized{
				PT: "FNA: forçar D3D11",
				EN: "FNA: force D3D11",
			},
			Category: Localized{
				PT: "Renderização",
				EN: "Rendering",
			},
			Compat: Localized{
				PT: "Jogos FNA/XNA (ex.: Celeste)",
				EN: "FNA/XNA games (e.g. Celeste)",
			},
			Description: Localized{
				PT: "Força o FNA (framework de jogos estilo XNA) a usar D3D11 para renderização.",
				EN: "Forces FNA (XNA-style game framework) to use D3D11 for rendering.",
			},
		},
	}
}
