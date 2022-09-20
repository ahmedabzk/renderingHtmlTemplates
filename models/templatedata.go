package models

// TemplateDAta holds data from the handlers to the template
type TemplateData struct{
	StringMap 	map[string]string
	IntMap 		map[string]int
	FloatMap 	map[string]float32
	Data 		map[string]interface{}
	CSRFToken 	string
	Flash 		string
	Warning 	string
	Error 		string
}