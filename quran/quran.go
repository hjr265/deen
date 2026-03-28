package quran

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hjr265/deen/alquran"
	"github.com/hjr265/deen/db"
	"github.com/hjr265/deen/model"
)

func Ayah(reference string, editions []string) ([]model.Ayah, error) {
	refs, err := expandReference(reference)
	if err != nil {
		return nil, err
	}

	store, err := db.OpenAyahs()
	if err != nil {
		return nil, err
	}
	defer store.Close()

	var result []model.Ayah
	for _, ref := range refs {
		ayahs, err := getAyah(store, ref, editions)
		if err != nil {
			return nil, err
		}
		result = append(result, ayahs...)
	}
	return result, nil
}

func expandReference(reference string) ([]string, error) {
	parts := strings.SplitN(reference, ":", 2)
	if len(parts) != 2 {
		return []string{reference}, nil
	}

	surah := parts[0]
	ayahPart := parts[1]

	rangeParts := strings.SplitN(ayahPart, "-", 2)
	if len(rangeParts) != 2 {
		return []string{reference}, nil
	}

	start, err := strconv.Atoi(rangeParts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid ayah range start: %s", rangeParts[0])
	}
	end, err := strconv.Atoi(rangeParts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid ayah range end: %s", rangeParts[1])
	}

	var refs []string
	for i := start; i <= end; i++ {
		refs = append(refs, fmt.Sprintf("%s:%d", surah, i))
	}
	return refs, nil
}

func toModel(a alquran.Ayah) model.Ayah {
	return model.Ayah{
		Text:          a.Text,
		NumberInSurah: a.NumberInSurah,
		SurahNumber:   a.Surah.Number,
		SurahName:     a.Surah.Name,
		SurahNameEn:   a.Surah.EnglishName,
		Edition:       a.Edition.Identifier,
		EditionNameEn: a.Edition.EnglishName,
	}
}

func getAyah(store *db.Ayahs, reference string, editions []string) ([]model.Ayah, error) {
	var cached []model.Ayah
	var missing []string
	for _, edition := range editions {
		ayah, err := store.Get(edition, reference)
		if err != nil {
			return nil, err
		}
		if ayah != nil {
			cached = append(cached, *ayah)
		} else {
			missing = append(missing, edition)
		}
	}

	if len(missing) == 0 {
		return cached, nil
	}

	fetched, err := alquran.GetAyah(reference, missing)
	if err != nil {
		return nil, err
	}
	for _, ayah := range fetched {
		m := toModel(ayah)
		if err := store.Put(reference, m); err != nil {
			return nil, err
		}
	}

	// Re-read from cache to pick up newly stored entries.
	cached = nil
	for _, edition := range editions {
		ayah, err := store.Get(edition, reference)
		if err != nil {
			return nil, err
		}
		if ayah != nil {
			cached = append(cached, *ayah)
		}
	}

	return cached, nil
}
