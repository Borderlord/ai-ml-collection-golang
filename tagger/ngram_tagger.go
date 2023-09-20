
package tagger

import (
	"container/list"
	"github.com/adrian3ka/go-learn-ai/helper"
	"strings"
)

type UnigramTagger struct {
	mapTag        map[string]string
	backoffTagger Tagger
}

type UnigramTaggerConfig struct {
	BackoffTagger Tagger
}

func NewUnigramTagger(cfg UnigramTaggerConfig) *UnigramTagger {
	u := UnigramTagger{
		backoffTagger: cfg.BackoffTagger,
	}
	u.mapTag = make(map[string]string)
	return &u
}

func (u *UnigramTagger) Predict(text string) ([][2]string, error) {
	splitedStrings := strings.Split(text, " ")
	var tuple [][2]string

	for _, splitedString := range splitedStrings {
		var selectedTag *string

		if helper.IsLetter(splitedString) {
			if val, exists := u.mapTag[splitedString]; exists {
				selectedTag = &val
			}
		}

		if selectedTag == nil && u.backoffTagger != nil {
			predictedValue, err := u.backoffTagger.Predict(splitedString)

			if err != nil {
				return nil, err
			}

			selectedTag = &predictedValue[0][1]
		}

		if selectedTag == nil {
			x := ""
			selectedTag = &x
		}

		tuple = append(tuple, [2]string{
			splitedString,
			*selectedTag,