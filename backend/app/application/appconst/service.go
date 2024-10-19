package appconst

import (
	"github.com/BenMeredithConsult/locagri-apps/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri-apps/ent"
)

type service struct {
	repo  gateways.AppConstRepo
	cache gateways.CacheService
}

const (
	LANG        = "const_lang"
	COUNTRY     = "const_country"
	NATIONALITY = "const_nationality"
)

func NewAppConstService(repo gateways.AppConstRepo,
	cacheSrv gateways.CacheService) gateways.AppConstService {
	return &service{
		repo:  repo,
		cache: cacheSrv,
	}
}

func (s *service) GetLanguages() ([]*ent.Language, error) {
	langs := new([]*ent.Language)

	if ok := s.cache.Has(LANG); ok {
		// Fetch from cache
		if err := s.cache.Get(LANG, langs); err == nil {
			return *langs, nil
		}
	}
	langsFromDB, err := s.repo.SelectLanguages()
	if err != nil {
		return nil, err
	}
	go func() {
		_ = s.cache.Set(LANG, langsFromDB, 0)
	}()
	return langsFromDB, nil
}

func (s *service) GetCountries() ([]*ent.Country, error) {
	countries := new([]*ent.Country)
	if ok := s.cache.Has(COUNTRY); ok {
		// Fetch from cache
		if err := s.cache.Get(COUNTRY, countries); err == nil {
			return *countries, nil
		}
	}
	countriesFromDB, err := s.repo.SelectCountries()
	if err != nil {
		return nil, err
	}
	go func() {
		_ = s.cache.Set(COUNTRY, countriesFromDB, 0)
	}()
	return countriesFromDB, nil
}

func (s *service) GetNationalities() ([]*ent.Nationality, error) {
	nationalities := new([]*ent.Nationality)
	if ok := s.cache.Has(NATIONALITY); ok {
		if err := s.cache.Get(NATIONALITY, nationalities); err == nil {
			return *nationalities, nil
		}
	}
	nationalitiesFromDB, err := s.repo.SelectNationalities()
	if err != nil {

		return nil, err
	}
	go func() {
		_ = s.cache.Set(NATIONALITY, nationalitiesFromDB, 0)
	}()
	return nationalitiesFromDB, nil
}
