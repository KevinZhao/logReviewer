package util

func BuildCodeAnalysisPrompt(codes []string) (string, []Message) {

	var messages []Message

	// System prompt preparation
	systemPrompt :=
		`
Human: You are a AI trained software developer expert, please follow below instruction and help me to anlyze my source code.
1. please list how many lines in source code file.
2. please extract public function name of the source code file. 
3. please write major action the public function did in the summary.

Please out put json format as below:\n
{
	"FileName": ""
}
`
	var userPrompt string

	userPrompt = "Here is some source code file, please help me to analyze them"

	for _, code := range codes {

		userPrompt = userPrompt + code + "/n/n"
	}

	assistantPrompt := `{
		"FileName":`

	messages = []Message{
		{Role: "user", Content: userPrompt},
		{Role: "assistant", Content: assistantPrompt}}

	return systemPrompt, messages
}
