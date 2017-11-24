//https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
//https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/

package chatbot
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var reflections map[string]string //map[string] of type string

type response struct {  //response structure with regex and array of answers
	rex     *regexp.Regexp
	answers []string
} //struct

func newResponse(pattern string, answers []string) response { //response function to return response
	response := response{}
	rex := regexp.MustCompile(pattern)
	response.rex = rex
	response.answers = answers
	return response
} 

func buildResponseList() []response { //buildResponse
	allResponses := []response{}
	file, err := os.Open("./data/patterns.dat") //opens patterns 
	if err != nil {  //if statement for error        
		panic(err) 
	} //buildResponse

	defer file.Close() //close file

	scanner := bufio.NewScanner(file) //scanner NewScanner

	for scanner.Scan() { //
		patternStr := scanner.Text()
		scanner.Scan() 
		answersAsStr := scanner.Text()

		answerList := strings.Split(answersAsStr, ";")     //In my patterns the various the bot responses to one input are seperated by ";"
		resp := newResponse("(?i)"+patternStr, answerList) //this regex will allow for any case (upper&lower) entered by the user
		allResponses = append(allResponses, resp)
	} //scanner

	return allResponses
}

func getRandomAnswer(answers []string) string {
	rand.Seed(time.Now().UnixNano()) // seed to make it return different values.
	index := rand.Intn(len(answers)) // Intn generates a number between 0 and num - 1
	return answers[index]            // can be any element
}

func subWords(original string) string {
	if reflections == nil { // map hasn't been made yet
		reflections = map[string]string{ // will only happen once.
			"am":     "are",
			"was":    "were",
			"i":      "you",
			"i'd":    "you would",
			"i've":   "you have",
			"i'll":   "you will",
			"my":     "your",
			"are":    "am",
			"you've": "I have",
			"you'll": "I will",
			"your":   "my",
			"yours":  "mine",
			"you":    "me",
			"me":     "you",
		}
	}

	words := strings.Split(original, " ")

	for index, word := range words { //change word if its in map
		val, ok := reflections[word]
		if ok {
			words[index] = val // swap with value 
		}
	}
	return strings.Join(words, " ")
}
func Ask(userInput string) string {
	responses := buildResponseList()
	for _, resp := range responses { // look at every single response/pattern/answers
		if resp.rex.MatchString(userInput) {
			match := resp.rex.FindStringSubmatch(userInput)
			captured := match[1]
			captured = subWords(captured)

			formatAnswer := getRandomAnswer(resp.answers) // get random element.

			if strings.Contains(formatAnswer, "%s") { //format answer
				formatAnswer = fmt.Sprintf(formatAnswer, captured)
			}
			return formatAnswer // returns formatted answer
		} 
	}
	return "No matches found" 
}
