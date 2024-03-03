package robloxapi

type RblxChallengeMetadata struct {
	UserID                           string `json:"userId"`
	ChallengeID                      string `json:"challengeId"`
	ShouldShowRememberDeviceCheckbox bool   `json:"shouldShowRememberDeviceCheckbox"`
	RememberDevice                   bool   `json:"rememberDevice"`
	SessionCookie                    string `json:"sessionCookie"`
	VerificationToken                string `json:"verificationToken"`
	ActionType                       string `json:"actionType"`
	RequestPath                      string `json:"requestPath"`
	RequestMethod                    string `json:"requestMethod"`
	SharedParameters                 struct {
		ShouldAnalyze         bool   `json:"shouldAnalyze"`
		GenericChallengeID    string `json:"genericChallengeId"`
		UseContinueMode       bool   `json:"useContinueMode"`
		RenderNativeChallenge bool   `json:"renderNativeChallenge"`
	} `json:"sharedParameters"`
}

type SolvedChallengeMetadata struct {
	VerificationToken string `json:"verificationToken"`
	RememberDevice    bool   `json:"rememberDevice"`
	ChallengeID       string `json:"challengeId"`
	ActionType        string `json:"actionType"`
}

func NewSolvedChallengeMetadata(verificationToken string, rememberDevice bool, challengeID string, actionType string) *SolvedChallengeMetadata {
	return &SolvedChallengeMetadata{
		verificationToken,
		rememberDevice,
		challengeID,
		actionType,
	}
}
