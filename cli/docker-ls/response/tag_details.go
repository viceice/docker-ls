package response

import (
	"git.mayflower.de/vaillant-team/docker-ls/lib"
)

type TagDetailsL0 struct {
	RepositoryName string `yaml:"repository,omitempty"`
	TagName        string `yaml:"tagName"`
	ContentDigest  string `yaml:"digest"`
}

func NewTagDetailsL0(tag lib.TagDetails, includeRepository bool) *TagDetailsL0 {
	tagDetails := TagDetailsL0{
		TagName:       tag.TagName(),
		ContentDigest: tag.ContentDigest(),
	}

	if includeRepository {
		tagDetails.RepositoryName = tag.RepositoryName()
	}

	return &tagDetails
}