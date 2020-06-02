package main

import (
	_ "cloud.google.com/go"
	_ "github.com/Azure/azure-sdk-for-go"
	_ "github.com/aliyun/alibaba-cloud-sdk-go"
	_ "github.com/aws/aws-sdk-go"
	_ "github.com/digitalocean/godo"

	_ "github.com/akamai/AkamaiOPEN-edgegrid-golang"
	_ "github.com/cloudflare/cloudflare-go"
	_ "github.com/fastly/go-fastly"

	_ "github.com/Medium/medium-sdk-go"
	_ "github.com/adlio/trello"
	_ "github.com/andygrunwald/go-jira"
	_ "github.com/dghubble/go-twitter"
	_ "github.com/google/go-github"
	_ "github.com/huandu/facebook"
	_ "github.com/jszwedko/go-circleci"
	_ "github.com/slack-go/slack"
	_ "github.com/stripe/stripe-go"
)
