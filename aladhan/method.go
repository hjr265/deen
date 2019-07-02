package aladhan

type Method int

const (
	ShiaIthnaAnsari                          = 0
	UniversityOfIslamicSciencesKarachi       = 1
	IslamicSocietyOfNorthAmerica             = 2
	MuslimWorldLeague                        = 3
	UmmAlQuraUniversityMakkah                = 4
	EgyptianGeneralAuthorityOfSurvey         = 5
	InstituteOfGeophysicsUniversityOfTehran  = 7
	GulfRegion                               = 8
	Kuwait                                   = 9
	Qatar                                    = 10
	MajlisUgamaIslamSingapuraSingapore       = 11
	UnionOrganizationislamicDeFrance         = 12
	DiyanetİşleriBaşkanlığıTurkey            = 13
	SpiritualAdministrationOfMuslimsOfRussia = 14
)

var Methods = map[int]Method{
	0:  ShiaIthnaAnsari,
	1:  UniversityOfIslamicSciencesKarachi,
	2:  IslamicSocietyOfNorthAmerica,
	3:  MuslimWorldLeague,
	4:  UmmAlQuraUniversityMakkah,
	5:  EgyptianGeneralAuthorityOfSurvey,
	7:  InstituteOfGeophysicsUniversityOfTehran,
	8:  GulfRegion,
	9:  Kuwait,
	10: Qatar,
	11: MajlisUgamaIslamSingapuraSingapore,
	12: UnionOrganizationislamicDeFrance,
	13: DiyanetİşleriBaşkanlığıTurkey,
	14: SpiritualAdministrationOfMuslimsOfRussia,
}
