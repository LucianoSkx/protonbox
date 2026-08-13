package main

type Launcher struct {
	ID     string
	Name   Localized
	HasCmd bool // usa o placeholder %command% nas opções de inicialização
	Hint   Localized
}

func launchers() []Launcher {
	return []Launcher{
		{
			ID: "steam",
			Name: Localized{
				PT: "Steam",
				EN: "Steam",
			},
			HasCmd: true,
			Hint: Localized{
				PT: "Cole nas opções de inicialização do jogo no Steam.",
				EN: "Paste into the game's launch options in Steam.",
			},
		},
		{
			ID: "faugus",
			Name: Localized{
				PT: "Faugus Launcher",
				EN: "Faugus Launcher",
			},
			HasCmd: true,
			Hint: Localized{
				PT: "O Faugus Launcher usa o mesmo formato do Steam (opções de inicialização).",
				EN: "Faugus Launcher uses the same format as Steam (launch options).",
			},
		},
		{
			ID: "heroic",
			Name: Localized{
				PT: "Heroic",
				EN: "Heroic",
			},
			HasCmd: false,
			Hint: Localized{
				PT: "Cole em Launch Options do jogo, sem %command% — só variáveis de ambiente e wrappers.",
				EN: "Paste into the game's Launch Options, without %command% — environment variables and wrappers only.",
			},
		},
		{
			ID: "lutris",
			Name: Localized{
				PT: "Lutris",
				EN: "Lutris",
			},
			HasCmd: false,
			Hint: Localized{
				PT: "Use nas variáveis de ambiente ou nas opções de lançamento do jogo.",
				EN: "Use in the game's environment variables or launch options.",
			},
		},
		{
			ID: "bottles",
			Name: Localized{
				PT: "Bottles",
				EN: "Bottles",
			},
			HasCmd: false,
			Hint: Localized{
				PT: "Adicione nas variáveis de ambiente do game no Bottles.",
				EN: "Add to the game's environment variables in Bottles.",
			},
		},
	}
}
