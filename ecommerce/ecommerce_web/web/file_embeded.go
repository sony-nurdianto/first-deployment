package web

import "embed"

//go:embed static/*
var TemplateFs embed.FS
