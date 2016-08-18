package rating

const (
	// FactorType - What type of bug is this? For example is it a crashing issue, a problem with localization or a matter of visual polish?
	FactorType = "type"
	// FactorLikelihood - How likely are users to experience the bug? For example, does everyone run into the issue or do only a few users run into it?
	FactorLikelihood = "likelihood"
	// FactorImpact - Of the people who experience the bug, how badly does it affect their experience with the product?
	FactorImpact = "impact"
)

const (
	// TypeCrash - Bug causes crash or data loss. Asserts in the Debug release
	TypeCrash = "crash"
	// TypeMajorUsability - Impairs usability in key scenarios
	TypeMajorUsability = "major-usability"
	// TypeMinorUsability - Impairs usability in secondary scenarios
	TypeMinorUsability = "minor-usability"
	// TypeBalancing - Enables degenerate usage strategies that harm the experience
	TypeBalancing = "balancing"
	// TypeVisual - Aesthetic issues
	TypeVisual = "visual"
	// TypeLocalization - Wording and translations
	TypeLocalization = "localization"
	// TypeDocumentation - A documentation issue
	TypeDocumentation = "documentation"
)

const (
	// ImpactBlocking - Blocking further progress on the daily workflow
	ImpactBlocking = "blocking"
	// ImpactReturn - A User would return the product. Cannot RTM. The Team would hold the release for this bug
	ImpactReturn = "return"
	// ImpactReject - A User would likely not purchase the product. Will show up in review. Clearly a noticeable issue
	ImpactReject = "reject"
	// ImpactPain - Users won’t like this once they notice it. A moderate number of users won’t buy
	ImpactPain = "pain"
	// ImpactNuisance - Not a big deal but noticeable. Extremely unlikely to affect sales
	ImpactNuisance = "nuisance"
)

const (
	// LikelihoodAll - Will affect all user
	LikelihoodAll = "all"
	// LikelihoodMost - Will affect most users
	LikelihoodMost = "most"
	// LikelihoodAverage - Will affect average number of users
	LikelihoodAverage = "average"
	// LikelihoodFew - Will only affect a few users
	LikelihoodFew = "few"
	// LikelihoodNone - Will affect almost no one
	LikelihoodNone = "none"
)

var listType = []string{TypeBalancing, TypeCrash, TypeDocumentation, TypeLocalization, TypeMajorUsability, TypeMinorUsability, TypeVisual}
var listImpact = []string{ImpactBlocking, ImpactNuisance, ImpactPain, ImpactReject, ImpactReturn}
var listLikelihood = []string{LikelihoodAll, LikelihoodAverage, LikelihoodFew, LikelihoodMost, LikelihoodNone}

// KnownFactor checks if given factor is available
func KnownFactor(factor string) bool {
	return factor == FactorType || factor == FactorLikelihood || factor == FactorImpact
}

// AvailableRatings return list of available ratings for factor
func AvailableRatings(factor string) []string {
	switch factor {
	case FactorType:
		return listType
	case FactorImpact:
		return listImpact
	case FactorLikelihood:
		return listLikelihood
	default:
		return []string{}
	}
}

// AllowedRating checks if given rating is available for factor
func AllowedRating(factor string, rating string) bool {
	list := AvailableRatings(factor)

	for _, item := range list {
		if rating == item {
			return true
		}
	}

	return false
}
