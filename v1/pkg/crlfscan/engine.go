package crlfscan

import "fmt"
import "strings"
import "regexp"
import "net/url"
import "github.com/machinexa2/gobasic"
import "github.com/machinexa2/gobasic/pathfuncs"

var re = regexp.MustCompile(`([^&]+)=([^&]+)`)
func recursivePathGenerator(inputDir string, suffix string) []string{
        split := strings.Split(inputDir, "/")
        result := make([]string, 0, len(split)-1)
        for i := 2; i < len(split); i++ {
                replaced := strings.Join(split[:i], "/") + "/" + suffix
                result = append(result, replaced)
        }
        return result
}

func expandParameters(query string) ([]string, []string) {
	// url.ParseQuery isn't used as map isnt sorted, and range doesnt print in order
	var parameter, value []string
	match := re.FindAllStringSubmatch(query, -1)
	for i := 0; i < len(match); i++ {
		group := match[i][1:]
		parameter = append(parameter, group[0])
		value = append(value, group[1])
	}
	return parameter, value
}

func parameterReplacement(parameter []string, value []string, replace_string string, selected_parameters_to_replace []string) [][]string{
	var returner_list [][]string
	var length int = len(parameter)
	for counter := 0; counter != length; counter++ {
		var temporary_counter []string
		holder := value[counter]
		for index := 0; index < length; index++ {
			value[counter] = replace_string
			/** Related to skipping, selected_parameters_to_replace is too **/
			if gobasic.InArray(parameter[index], selected_parameters_to_replace) {
				temporary_counter = append(temporary_counter, parameter[index] + "=" + value[index])
			}
		}
		returner_list = append(returner_list, temporary_counter)
		value[counter] = holder
	}
	var returnable_list [][]string
	// just eliminating unnecessary parameters (optional)
	for i := 0; i < len(returner_list); i++ {
		contains := false
		individual := returner_list[i]
		for j := 0; j < len(individual); j++ {
			var unit_individual string = individual[j]
			if strings.Split(unit_individual, "=")[1] == replace_string {
				contains = true
			}
		}
		if contains == true {
			returnable_list = append(returnable_list, individual)
		}
	}
	return returnable_list
}

func queryUrlsGenerator(base_url string, parameters [][]string) []string{
	var new_urls []string
	base_url = pathfunctions.Ender(base_url, "?")
	for i := 0; i < len(parameters); i++ {
		individual := parameters[i]
		joined := strings.Join(individual, "&")
		new_urls = append(new_urls, base_url + joined)
	}
	return new_urls
}

func queryGenerator(single_url string, payloads []string) []string{
	var payloads_list []string
	var base_url string
	var query string
	parsed_url, _ := url.Parse(single_url)
	base_url = parsed_url.Scheme + "://" + parsed_url.Host + parsed_url.Path
	query = parsed_url.RawQuery
	if len(query) > 500 {
		return payloads_list
	}
	parameters, values := expandParameters(query)
	/** Skipping not implemented and wont be I guess **/
	selected_parameters := parameters // skipping related variable
	for i := 0; i < len(payloads); i++ {
		payload := payloads[i]
		replaced_query := parameterReplacement(parameters, values, payload, selected_parameters)
		payloaded_urls := queryUrlsGenerator(base_url, replaced_query)
		for j := 0; j < len(payloaded_urls); j++ {
			var payloaded_url string = payloaded_urls[j]
			payloads_list = append(payloads_list, payloaded_url)
		}
	}
	 return payloads_list

}

func pathGenerator(single_url string, payloads []string) []string{
	var payloads_list []string
	var base_url string
	parsed_url, _ := url.Parse(single_url)
	base_url = parsed_url.Scheme + "://" + parsed_url.Host
	var count int = strings.Count(parsed_url.Path, "/")
	if count > 1 {
		for i := 0; i < len(payloads); i++ {
			var payloaded_directories []string = recursivePathGenerator(parsed_url.Path, payloads[i])
			for j := 0; j < len(payloaded_directories); j++ {
				payloaded_url := pathfunctions.Ender(base_url, "/") + pathfunctions.Unstarter(payloaded_directories[j], "/")
				payloads_list = append(payloads_list, payloaded_url)
			}
		}
	} else {
		payloads_list = domainGenerator(single_url, payloads)
		return payloads_list
	}
	return payloads_list
}

func domainGenerator(single_url string, payloads []string) []string{
	var payloads_list []string
	var base_url string
	parsed_url, _ := url.Parse(single_url)
	base_url = parsed_url.Scheme + "://" + parsed_url.Host
	for i := 0; i < len(payloads); i++ {
		var payloaded_url string
		payloaded_url = pathfunctions.Ender(base_url, "/") + pathfunctions.Unstarter(payloads[i], "/")
		payloads_list = append(payloads_list, payloaded_url)
	}
	return payloads_list
}
func urlsGenerator(single_url string, payloads []string) []string{
	var mega_urls []string
	single_url = pathfunctions.Urler(single_url)
	parsed_url, err := url.Parse(single_url)
	if err != nil {
		fmt.Printf("Skipping %s due to error\n", single_url);
		fmt.Println(err);
		return mega_urls
	}
	if parsed_url.RawQuery != "" {
		temporary_urls := queryGenerator(single_url, payloads)
		for single := 0; single < len(temporary_urls); single++ {
			mega_urls = append(mega_urls, temporary_urls[single])
		}
	} else if parsed_url.Path != "" {
		temporary_urls := pathGenerator(single_url, payloads)
		for single := 0; single < len(temporary_urls); single++ {
			mega_urls = append(mega_urls, temporary_urls[single])
		}
	} else if parsed_url.Host != "" {
		temporary_urls := domainGenerator(single_url, payloads)
		for single := 0; single < len(temporary_urls); single++ {
			mega_urls = append(mega_urls, temporary_urls[single])
		}
	}
	return mega_urls
}
func Generator(targets chan string, payloads []string, queue chan string){
	for target := range targets {
		payloaded_targets := urlsGenerator(target, payloads)
		if len(payloaded_targets) > 0 {
			for t := 0; t < len(payloaded_targets); t++ {
				queue <- payloaded_targets[t]
			}
		}
	}
	close(queue)
}
