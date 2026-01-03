package utils

import (
	"sort"
	"strings"
	"web-lab/internal/entity"
)

type CategorizedGroups struct {
	Groups []AlphabetGroup   `json:"groups"`
	Other  []entity.Category `json:"other"` // Для цифр, символов и т.д.
}

type AlphabetGroup struct {
	Letter     string            `json:"letter"`
	Categories []entity.Category `json:"categories"`
}

func CategoriesGroupedByFirstLetter(categories []entity.Category) *CategorizedGroups {
	groups := make(map[string][]entity.Category)
	var otherCategories []entity.Category

	for _, category := range categories {
		if category.Name == "" {
			continue
		}

		// В категории убираем публикации-черновики
		var publishedCategories []entity.PublicationCategories
		for _, pc := range category.PublicationCategories {
			if !pc.Publication.IsDraft {
				publishedCategories = append(publishedCategories, pc)
			}
		}
		category.PublicationCategories = publishedCategories

		if len(publishedCategories) > 0 {
			// Группируем по первой букве (ex. A: cat1, cat2, ...)
			firstRune := []rune(strings.ToUpper(category.Name))[0]
			switch {
			case firstRune >= 'А' && firstRune <= 'Я':
				groups[string(firstRune)] = append(groups[string(firstRune)], category)
			case firstRune == 'Ё':
				groups["Ё"] = append(groups["Ё"], category)
			case firstRune >= 'A' && firstRune <= 'Z':
				groups[string(firstRune)] = append(groups[string(firstRune)], category)
			default:
				otherCategories = append(otherCategories, category)
			}
		}
	}

	// Преобразуем map в отсортированный слайс
	var resultGroups []AlphabetGroup
	var letters []string

	// Собираем буквы
	for letter := range groups {
		letters = append(letters, letter)
	}

	// Сортируем буквы
	sort.Strings(letters)

	// Формируем группы
	for _, letter := range letters {
		// Сортируем категории внутри группы по имени
		sort.Slice(groups[letter], func(i, j int) bool {
			return strings.ToLower(groups[letter][i].Name) <
				strings.ToLower(groups[letter][j].Name)
		})

		resultGroups = append(resultGroups, AlphabetGroup{
			Letter:     letter,
			Categories: groups[letter],
		})
	}

	return &CategorizedGroups{
		Groups: resultGroups,
		Other:  otherCategories,
	}
}
