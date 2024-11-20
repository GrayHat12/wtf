package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type WikipediaTitleResponse struct {
	Pageid      int    `json:"id"`
	Key         string `json:"key"`
	Title       string `json:"title"`
	Excerpt     string `json:"excerpt"`
	Description string `json:"description"`
}

type wikipediaTitleResponseRaw struct {
	Pages []WikipediaTitleResponse `json:"pages"`
}

type wikipediaSummaryResponseRaw struct {
	Query struct {
		Pages map[string]struct {
			Pageid  int    `json:"pageid"`
			Title   string `json:"title"`
			Summary string `json:"extract"`
		} `json:"pages"`
	} `json:"query"`
}

type QueryResponse struct {
	Query       string
	WikiTitle   *WikipediaTitleResponse
	WikiSummary string
	Error       error
}

func fetchPage(query string, response *QueryResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	query_params := url.Values{}
	query_params.Add("q", query)
	query_params.Add("limit", "1")
	url := fmt.Sprintf("https://en.wikipedia.org/w/rest.php/v1/search/title?%s", query_params.Encode())

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		response.Error = err
		return
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		response.Error = err
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		response.Error = err
		return
	}

	var result wikipediaTitleResponseRaw
	err = json.Unmarshal(body, &result)
	if err != nil {
		response.Error = err
		return
	}
	if len(result.Pages) > 0 {
		response.WikiTitle = &result.Pages[0]
		// fmt.Fprintf(os.Stdout, "Fetched title info %+v\n", result.Pages[0])
	} else {
		response.Error = errors.New("Pretty absurd __init__")
	}
}

func fetchSummary(responses *[]QueryResponse) error {
	pageids := ""
	for _, resp := range *responses {
		if resp.WikiTitle == nil {
			// fmt.Fprintf(os.Stdout, "Skipping query %s because %+v\n", resp.Query, resp)
			continue
		}
		if len(pageids) > 0 {
			pageids = fmt.Sprintf("%s|%d", pageids, resp.WikiTitle.Pageid)
		} else {
			pageids = fmt.Sprintf("%d", resp.WikiTitle.Pageid)
		}
	}
	query_params := url.Values{}
	query_params.Add("action", "query")
	query_params.Add("prop", "extracts")
	query_params.Add("pageids", pageids)
	query_params.Add("format", "json")
	query_params.Add("exsentences", "3")
	query_params.Add("exintro", "true")
	query_params.Add("explaintext", "true")
	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?%s", query_params.Encode())

	// fmt.Fprintf(os.Stdout, "Making GET %s\n", url)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// // fmt.Fprintf(os.Stdout, "body %s\n", string(body))

	var result wikipediaSummaryResponseRaw
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	for index, resp := range *responses {
		if resp.WikiTitle == nil {
			// fmt.Fprintf(os.Stdout, "Skipping query(2) %s because %+v\n", resp.Query, resp)
			continue
		}
		if val, ok := result.Query.Pages[fmt.Sprintf("%d", resp.WikiTitle.Pageid)]; ok {
			(*responses)[index].WikiSummary = val.Summary
			// fmt.Fprintf(os.Stdout, "Setting wikiSummary %s %+v\n", val.Summary, (*responses)[index])
		} else {
			// fmt.Fprintf(os.Stdout, "Not found key %s %+v\n", fmt.Sprintf("%d", resp.WikiTitle.Pageid), result.Query.Pages)
		}
	}
	return nil
}

func Is(args []string) (result string) {
	result = ""
	queryResponses := make([]QueryResponse, len(args))
	waitGroup := sync.WaitGroup{}
	for index, argument := range args {
		queryResponses[index] = QueryResponse{
			Query: argument,
		}
		waitGroup.Add(1)
		go fetchPage(argument, &queryResponses[index], &waitGroup)
	}
	waitGroup.Wait()
	// fmt.Fprintf(os.Stdout, "Fetched titles %+v\n", queryResponses)
	err := fetchSummary(&queryResponses)
	// fmt.Fprintf(os.Stdout, "Fetched summary %+v\n", queryResponses)
	if err != nil {
		return fmt.Sprintf("Some error occured %s\n", err.Error())
	}

	for index, response := range queryResponses {
		if response.Error != nil {
			result += fmt.Sprintf("%d.> Error - %s\n\n", index+1, response.Error.Error())
		} else {
			result += fmt.Sprintf("%d.> %s or %s : %s\n\n", index+1, response.Query, response.WikiTitle.Title, strings.ReplaceAll(response.WikiSummary, "\n", " "))
		}
	}
	return result
}
