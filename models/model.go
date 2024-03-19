package models

type LevelIn struct{
	GeneName		string    `json:"genename"`
	CancerName 		string	  `json:"cancername"`
	Mutation		string	  `json:"mutation"`
	Drug 			string	  `json:"drug"`
}

