package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type PokemonResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      *string `json:"back_default"`
		BackFemale       *string `json:"back_female"`
		BackShiny        *string `json:"back_shiny"`
		BackShinyFemale  *string `json:"back_shiny_female"`
		FrontDefault     *string `json:"front_default"`
		FrontFemale      *string `json:"front_female"`
		FrontShiny       *string `json:"front_shiny"`
		FrontShinyFemale *string `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault *string `json:"front_default"`
				FrontFemale  *string `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     *string `json:"front_default"`
				FrontFemale      *string `json:"front_female"`
				FrontShiny       *string `json:"front_shiny"`
				FrontShinyFemale *string `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault *string `json:"front_default"`
				FrontShiny   *string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      *string `json:"back_default"`
				BackFemale       *string `json:"back_female"`
				BackShiny        *string `json:"back_shiny"`
				BackShinyFemale  *string `json:"back_shiny_female"`
				FrontDefault     *string `json:"front_default"`
				FrontFemale      *string `json:"front_female"`
				FrontShiny       *string `json:"front_shiny"`
				FrontShinyFemale *string `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			Generation1 struct {
				RedBlue struct {
					BackDefault  *string `json:"back_default"`
					BackGray     *string `json:"back_gray"`
					FrontDefault *string `json:"front_default"`
					FrontGray    *string `json:"front_gray"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault  *string `json:"back_default"`
					BackGray     *string `json:"back_gray"`
					FrontDefault *string `json:"front_default"`
					FrontGray    *string `json:"front_gray"`
				} `json:"yellow"`
			} `json:"generation-i"`
			Generation2 struct {
				Crystal struct {
					BackDefault  *string `json:"back_default"`
					BackShiny    *string `json:"back_shiny"`
					FrontDefault *string `json:"front_default"`
					FrontShiny   *string `json:"front_shiny"`
				} `json:"crystal"`
				Gold struct {
					BackDefault  *string `json:"back_default"`
					BackShiny    *string `json:"back_shiny"`
					FrontDefault *string `json:"front_default"`
					FrontShiny   *string `json:"front_shiny"`
				} `json:"gold"`
				Silver struct {
					BackDefault  *string `json:"back_default"`
					BackShiny    *string `json:"back_shiny"`
					FrontDefault *string `json:"front_default"`
					FrontShiny   *string `json:"front_shiny"`
				} `json:"silver"`
			} `json:"generation-ii"`
			Generation3 struct {
				Emerald struct {
					FrontDefault *string `json:"front_default"`
					FrontShiny   *string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  *string `json:"back_default"`
					BackShiny    *string `json:"back_shiny"`
					FrontDefault *string `json:"front_default"`
					FrontShiny   *string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  *string `json:"back_default"`
					BackShiny    *string `json:"back_shiny"`
					FrontDefault *string `json:"front_default"`
					FrontShiny   *string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			Generation4 struct {
				DiamondPearl struct {
					BackDefault      *string `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        *string `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      *string `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        *string `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      *string `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        *string `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			Generation5 struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      *string `json:"back_default"`
						BackFemale       *string `json:"back_female"`
						BackShiny        *string `json:"back_shiny"`
						BackShinyFemale  *string `json:"back_shiny_female"`
						FrontDefault     *string `json:"front_default"`
						FrontFemale      *string `json:"front_female"`
						FrontShiny       *string `json:"front_shiny"`
						FrontShinyFemale *string `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      *string `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        *string `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			Generation6 struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			Generation7 struct {
				Icons struct {
					FrontDefault *string `json:"front_default"`
					FrontFemale  *string `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     *string `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       *string `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			Generation8 struct {
				Icons struct {
					FrontDefault *string `json:"front_default"`
					FrontFemale  *string `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Cries struct {
		Latest *string `json:"latest"`
		Legacy *string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Types []struct {
			Slot int `json:"slot"`
			Type struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
}

func (c Client) GetPokemon(pokemonName string) (PokemonResponse, error) {
	fullURL := baseURL + "/pokemon/" + pokemonName

	value, exists := c.cache.Get(fullURL)
	if exists {
		result := PokemonResponse{}
		err := json.Unmarshal(value, &result)
		if err != nil {
			return PokemonResponse{}, err
		}
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	pokemon := PokemonResponse{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return PokemonResponse{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemon, nil
}
