package models

// var SIMPLE_FRONTEND_HUH_FORM = huh.NewForm(
// 	huh.NewGroup(
// 		huh.NewMultiSelect[int]().
// 			Title("Choose one or more deployment strategies!").
// 			Options(
// 				utils.GetDeploymentStrategyHuhOptions(projectType)...,
// 			).
// 			Value(&projectDeploymentStrategy),
// 	),
// 	huh.NewGroup(
// 		huh.NewConfirm().
// 			Title("Shall we create the project? Here are the selected configuration:").
// 			Description(
// 				utils.FormatProjectConfig(
// 					projectName,
// 					projectDir,
// 					projectType,
// 					projectDeploymentStrategy,
// 				)).
// 			Value(&confirm),
// 	),
// )
