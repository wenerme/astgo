package xmlgen

import "encoding/xml"

type Docs struct {
	XMLName    xml.Name `xml:"docs"`
	Text       string   `xml:",chardata"`
	Xi         string   `xml:"xi,attr"`
	ConfigInfo []struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name,attr"`
		Language    string `xml:"language,attr"`
		Synopsis    string `xml:"synopsis"`
		Description struct {
			Text string `xml:",chardata"`
			Para []struct {
				Text     string   `xml:",chardata"`
				Emphasis []string `xml:"emphasis"`
				Literal  []string `xml:"literal"`
				Filename string   `xml:"filename"`
			} `xml:"para"`
			Warning struct {
				Text string `xml:",chardata"`
				Para struct {
					Text     string `xml:",chardata"`
					Emphasis string `xml:"emphasis"`
				} `xml:"para"`
			} `xml:"warning"`
			Note struct {
				Text string `xml:",chardata"`
				Para string `xml:"para"`
			} `xml:"note"`
			Enumlist []struct {
				Text string `xml:",chardata"`
				Enum []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					Para string `xml:"para"`
				} `xml:"enum"`
			} `xml:"enumlist"`
		} `xml:"description"`
		ConfigFile struct {
			Text         string `xml:",chardata"`
			Name         string `xml:"name,attr"`
			ConfigObject []struct {
				Text         string `xml:",chardata"`
				Name         string `xml:"name,attr"`
				Synopsis     string `xml:"synopsis"`
				ConfigOption []struct {
					Text        string `xml:",chardata"`
					Name        string `xml:"name,attr"`
					Default     string `xml:"default,attr"`
					Synopsis    string `xml:"synopsis"`
					Description struct {
						Text string `xml:",chardata"`
						Para []struct {
							Text        string   `xml:",chardata"`
							Literal     []string `xml:"literal"`
							Filename    []string `xml:"filename"`
							Replaceable []string `xml:"replaceable"`
							Variable    string   `xml:"variable"`
							Emphasis    []string `xml:"emphasis"`
							Astcli      string   `xml:"astcli"`
						} `xml:"para"`
						Enumlist struct {
							Text string `xml:",chardata"`
							Enum []struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
								Para []struct {
									Text        string   `xml:",chardata"`
									Literal     []string `xml:"literal"`
									Replaceable string   `xml:"replaceable"`
									Emphasis    string   `xml:"emphasis"`
								} `xml:"para"`
								Enumlist struct {
									Text string `xml:",chardata"`
									Enum []struct {
										Text string `xml:",chardata"`
										Name string `xml:"name,attr"`
										Para string `xml:"para"`
									} `xml:"enum"`
								} `xml:"enumlist"`
								Note struct {
									Text string `xml:",chardata"`
									Para struct {
										Text    string `xml:",chardata"`
										Literal string `xml:"literal"`
									} `xml:"para"`
								} `xml:"note"`
							} `xml:"enum"`
						} `xml:"enumlist"`
						Note []struct {
							Text string `xml:",chardata"`
							Para []struct {
								Text        string   `xml:",chardata"`
								Variable    string   `xml:"variable"`
								Literal     []string `xml:"literal"`
								Replaceable []string `xml:"replaceable"`
								Filename    string   `xml:"filename"`
							} `xml:"para"`
						} `xml:"note"`
						Include struct {
							Text     string `xml:",chardata"`
							Xpointer string `xml:"xpointer,attr"`
						} `xml:"include"`
						Warning struct {
							Text string `xml:",chardata"`
							Para struct {
								Text    string `xml:",chardata"`
								Literal string `xml:"literal"`
							} `xml:"para"`
						} `xml:"warning"`
						Variablelist struct {
							Text     string `xml:",chardata"`
							Variable []struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
								Para struct {
									Text     string `xml:",chardata"`
									Variable string `xml:"variable"`
								} `xml:"para"`
							} `xml:"variable"`
						} `xml:"variablelist"`
						Example []struct {
							Text  string `xml:",chardata"`
							Title string `xml:"title,attr"`
						} `xml:"example"`
					} `xml:"description"`
					SeeAlso struct {
						Text string `xml:",chardata"`
						Ref  []struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"ref"`
					} `xml:"see-also"`
					Syntax struct {
						Text      string `xml:",chardata"`
						Argsep    string `xml:"argsep,attr"`
						Parameter []struct {
							Text       string `xml:",chardata"`
							Name       string `xml:"name,attr"`
							Required   string `xml:"required,attr"`
							Para       string `xml:"para"`
							Optionlist struct {
								Text   string `xml:",chardata"`
								Option []struct {
									Text string `xml:",chardata"`
									Name string `xml:"name,attr"`
									Para string `xml:"para"`
								} `xml:"option"`
							} `xml:"optionlist"`
							Argument struct {
								Text     string `xml:",chardata"`
								Name     string `xml:"name,attr"`
								Required string `xml:"required,attr"`
								Para     string `xml:"para"`
							} `xml:"argument"`
						} `xml:"parameter"`
					} `xml:"syntax"`
				} `xml:"configOption"`
				Description struct {
					Text    string `xml:",chardata"`
					Include struct {
						Text     string `xml:",chardata"`
						Xpointer string `xml:"xpointer,attr"`
					} `xml:"include"`
					Para []struct {
						Text        string   `xml:",chardata"`
						Literal     []string `xml:"literal"`
						Replaceable []string `xml:"replaceable"`
						Emphasis    []string `xml:"emphasis"`
						Filename    string   `xml:"filename"`
					} `xml:"para"`
					Enumlist struct {
						Text string `xml:",chardata"`
						Enum []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
							Para string `xml:"para"`
						} `xml:"enum"`
					} `xml:"enumlist"`
					Note struct {
						Text string `xml:",chardata"`
						Para struct {
							Text     string `xml:",chardata"`
							Emphasis string `xml:"emphasis"`
							Filename string `xml:"filename"`
						} `xml:"para"`
					} `xml:"note"`
				} `xml:"description"`
			} `xml:"configObject"`
		} `xml:"configFile"`
	} `xml:"configInfo"`
	Manager []struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name,attr"`
		Language string `xml:"language,attr"`
		Module   string `xml:"module,attr"`
		Synopsis string `xml:"synopsis"`
		Syntax   struct {
			Text    string `xml:",chardata"`
			Include []struct {
				Text     string `xml:",chardata"`
				Xpointer string `xml:"xpointer,attr"`
			} `xml:"include"`
			Parameter []struct {
				Text     string `xml:",chardata"`
				Name     string `xml:"name,attr"`
				Required string `xml:"required,attr"`
				Default  string `xml:"default,attr"`
				Para     []struct {
					Text        string   `xml:",chardata"`
					Replaceable []string `xml:"replaceable"`
					Literal     []string `xml:"literal"`
					Variable    string   `xml:"variable"`
					Filename    string   `xml:"filename"`
				} `xml:"para"`
				Enumlist []struct {
					Text string `xml:",chardata"`
					Enum []struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
						Para struct {
							Text        string `xml:",chardata"`
							Replaceable string `xml:"replaceable"`
						} `xml:"para"`
						Note struct {
							Text string `xml:",chardata"`
							Para struct {
								Text    string `xml:",chardata"`
								Literal string `xml:"literal"`
							} `xml:"para"`
						} `xml:"note"`
						Enumlist struct {
							Text string `xml:",chardata"`
							Enum []struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
								Para []struct {
									Text    string `xml:",chardata"`
									Literal string `xml:"literal"`
								} `xml:"para"`
								Include struct {
									Text     string `xml:",chardata"`
									Xpointer string `xml:"xpointer,attr"`
								} `xml:"include"`
							} `xml:"enum"`
						} `xml:"enumlist"`
					} `xml:"enum"`
				} `xml:"enumlist"`
				Warning struct {
					Text string `xml:",chardata"`
					Para struct {
						Text     string   `xml:",chardata"`
						Variable []string `xml:"variable"`
					} `xml:"para"`
				} `xml:"warning"`
				Include struct {
					Text     string `xml:",chardata"`
					Xpointer string `xml:"xpointer,attr"`
				} `xml:"include"`
			} `xml:"parameter"`
		} `xml:"syntax"`
		Description struct {
			Text string `xml:",chardata"`
			Para []struct {
				Text        string   `xml:",chardata"`
				Variable    string   `xml:"variable"`
				Literal     []string `xml:"literal"`
				Replaceable []string `xml:"replaceable"`
			} `xml:"para"`
			Note struct {
				Text string `xml:",chardata"`
				Para struct {
					Text        string   `xml:",chardata"`
					Literal     []string `xml:"literal"`
					Replaceable []string `xml:"replaceable"`
				} `xml:"para"`
			} `xml:"note"`
			Variablelist struct {
				Text     string `xml:",chardata"`
				Variable struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					Para string `xml:"para"`
				} `xml:"variable"`
			} `xml:"variablelist"`
			Warning struct {
				Text string `xml:",chardata"`
				Para struct {
					Text    string `xml:",chardata"`
					Literal string `xml:"literal"`
				} `xml:"para"`
			} `xml:"warning"`
		} `xml:"description"`
		SeeAlso struct {
			Text string `xml:",chardata"`
			Ref  []struct {
				Text   string `xml:",chardata"`
				Type   string `xml:"type,attr"`
				Module string `xml:"module,attr"`
			} `xml:"ref"`
		} `xml:"see-also"`
		Responses struct {
			Text         string `xml:",chardata"`
			ListElements struct {
				Text    string `xml:",chardata"`
				Include []struct {
					Text     string `xml:",chardata"`
					Xpointer string `xml:"xpointer,attr"`
				} `xml:"include"`
				ManagerEvent struct {
					Text                 string `xml:",chardata"`
					Language             string `xml:"language,attr"`
					Name                 string `xml:"name,attr"`
					ManagerEventInstance struct {
						Text     string `xml:",chardata"`
						Class    string `xml:"class,attr"`
						Synopsis string `xml:"synopsis"`
						Syntax   struct {
							Text            string `xml:",chardata"`
							ChannelSnapshot string `xml:"channel_snapshot"`
						} `xml:"syntax"`
					} `xml:"managerEventInstance"`
				} `xml:"managerEvent"`
			} `xml:"list-elements"`
			Include struct {
				Text     string `xml:",chardata"`
				Xpointer string `xml:"xpointer,attr"`
			} `xml:"include"`
			ManagerEvent struct {
				Text                 string `xml:",chardata"`
				Name                 string `xml:"name,attr"`
				Language             string `xml:"language,attr"`
				ManagerEventInstance struct {
					Text     string `xml:",chardata"`
					Class    string `xml:"class,attr"`
					Synopsis string `xml:"synopsis"`
					Syntax   struct {
						Text      string `xml:",chardata"`
						Parameter []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
							Para string `xml:"para"`
						} `xml:"parameter"`
						BridgeSnapshot string `xml:"bridge_snapshot"`
					} `xml:"syntax"`
				} `xml:"managerEventInstance"`
			} `xml:"managerEvent"`
		} `xml:"responses"`
	} `xml:"manager"`
	Function []struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name,attr"`
		Language string `xml:"language,attr"`
		Module   string `xml:"module,attr"`
		Synopsis string `xml:"synopsis"`
		Syntax   struct {
			Text      string `xml:",chardata"`
			Argsep    string `xml:"argsep,attr"`
			Parameter []struct {
				Text     string `xml:",chardata"`
				Name     string `xml:"name,attr"`
				Required string `xml:"required,attr"`
				Argsep   string `xml:"argsep,attr"`
				Multiple string `xml:"multiple,attr"`
				Default  string `xml:"default,attr"`
				Para     []struct {
					Text        string   `xml:",chardata"`
					Literal     []string `xml:"literal"`
					Replaceable []string `xml:"replaceable"`
					Filename    string   `xml:"filename"`
					Emphasis    string   `xml:"emphasis"`
				} `xml:"para"`
				Enumlist struct {
					Text string `xml:",chardata"`
					Enum []struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
						Para []struct {
							Text        string   `xml:",chardata"`
							Replaceable []string `xml:"replaceable"`
							Literal     []string `xml:"literal"`
							Include     struct {
								Text     string `xml:",chardata"`
								Xpointer string `xml:"xpointer,attr"`
							} `xml:"include"`
							Filename string   `xml:"filename"`
							Emphasis []string `xml:"emphasis"`
						} `xml:"para"`
						Enumlist struct {
							Text string `xml:",chardata"`
							Enum []struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
								Para struct {
									Text    string `xml:",chardata"`
									Literal string `xml:"literal"`
								} `xml:"para"`
							} `xml:"enum"`
						} `xml:"enumlist"`
						Warning struct {
							Text string `xml:",chardata"`
							Para string `xml:"para"`
						} `xml:"warning"`
						Note struct {
							Text string `xml:",chardata"`
							Para string `xml:"para"`
						} `xml:"note"`
					} `xml:"enum"`
					ConfigOptionToEnum struct {
						Text    string `xml:",chardata"`
						Include struct {
							Text     string `xml:",chardata"`
							Xpointer string `xml:"xpointer,attr"`
						} `xml:"include"`
					} `xml:"configOptionToEnum"`
				} `xml:"enumlist"`
				Optionlist struct {
					Text   string `xml:",chardata"`
					Option []struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
						Para []struct {
							Text        string   `xml:",chardata"`
							Literal     []string `xml:"literal"`
							Replaceable []string `xml:"replaceable"`
						} `xml:"para"`
						Argument []struct {
							Text     string `xml:",chardata"`
							Name     string `xml:"name,attr"`
							Required string `xml:"required,attr"`
							Para     string `xml:"para"`
						} `xml:"argument"`
					} `xml:"option"`
				} `xml:"optionlist"`
				Argument []struct {
					Text     string `xml:",chardata"`
					Name     string `xml:"name,attr"`
					Required string `xml:"required,attr"`
				} `xml:"argument"`
				Include struct {
					Text     string `xml:",chardata"`
					Xpointer string `xml:"xpointer,attr"`
				} `xml:"include"`
				Warning struct {
					Text string `xml:",chardata"`
					Para struct {
						Text     string   `xml:",chardata"`
						Variable []string `xml:"variable"`
					} `xml:"para"`
				} `xml:"warning"`
				Note struct {
					Text string `xml:",chardata"`
					Para string `xml:"para"`
				} `xml:"note"`
			} `xml:"parameter"`
		} `xml:"syntax"`
		Description struct {
			Text string `xml:",chardata"`
			Para []struct {
				Text        string   `xml:",chardata"`
				Replaceable []string `xml:"replaceable"`
				Literal     []string `xml:"literal"`
				Variable    []string `xml:"variable"`
				Filename    []string `xml:"filename"`
				Emphasis    string   `xml:"emphasis"`
			} `xml:"para"`
			Warning struct {
				Text string `xml:",chardata"`
				Para struct {
					Text     string   `xml:",chardata"`
					Emphasis []string `xml:"emphasis"`
					Literal  string   `xml:"literal"`
				} `xml:"para"`
			} `xml:"warning"`
			Example []struct {
				Text     string `xml:",chardata"`
				Title    string `xml:"title,attr"`
				Language string `xml:"language,attr"`
			} `xml:"example"`
			Note []struct {
				Text string `xml:",chardata"`
				Para []struct {
					Text        string   `xml:",chardata"`
					Literal     []string `xml:"literal"`
					Replaceable []string `xml:"replaceable"`
					Emphasis    []string `xml:"emphasis"`
				} `xml:"para"`
			} `xml:"note"`
			Enumlist []struct {
				Text string `xml:",chardata"`
				Enum []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					Para string `xml:"para"`
				} `xml:"enum"`
			} `xml:"enumlist"`
			Include struct {
				Text     string `xml:",chardata"`
				Xpointer string `xml:"xpointer,attr"`
			} `xml:"include"`
			Variablelist struct {
				Text     string `xml:",chardata"`
				Variable struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"name,attr"`
					Value []struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
					} `xml:"value"`
					Para string `xml:"para"`
				} `xml:"variable"`
			} `xml:"variablelist"`
		} `xml:"description"`
		SeeAlso struct {
			Text string `xml:",chardata"`
			Ref  []struct {
				Text   string `xml:",chardata"`
				Type   string `xml:"type,attr"`
				Module string `xml:"module,attr"`
			} `xml:"ref"`
		} `xml:"see-also"`
	} `xml:"function"`
	Info []struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name,attr"`
		Language string `xml:"language,attr"`
		Tech     string `xml:"tech,attr"`
		Enumlist struct {
			Text string `xml:",chardata"`
			Enum []struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
				Para []struct {
					Text        string   `xml:",chardata"`
					Replaceable string   `xml:"replaceable"`
					Literal     []string `xml:"literal"`
				} `xml:"para"`
				Parameter []struct {
					Text     string `xml:",chardata"`
					Name     string `xml:"name,attr"`
					Required string `xml:"required,attr"`
					Para     struct {
						Text        string `xml:",chardata"`
						Replaceable string `xml:"replaceable"`
						Literal     string `xml:"literal"`
					} `xml:"para"`
					Enumlist struct {
						Text string `xml:",chardata"`
						Enum []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
							Para []struct {
								Text    string   `xml:",chardata"`
								Literal []string `xml:"literal"`
							} `xml:"para"`
							Enumlist struct {
								Text string `xml:",chardata"`
								Enum []struct {
									Text string `xml:",chardata"`
									Name string `xml:"name,attr"`
									Para struct {
										Text    string `xml:",chardata"`
										Literal string `xml:"literal"`
									} `xml:"para"`
								} `xml:"enum"`
							} `xml:"enumlist"`
							Note struct {
								Text string `xml:",chardata"`
								Para struct {
									Text    string `xml:",chardata"`
									Literal string `xml:"literal"`
								} `xml:"para"`
							} `xml:"note"`
						} `xml:"enum"`
					} `xml:"enumlist"`
				} `xml:"parameter"`
				Enumlist struct {
					Text string `xml:",chardata"`
					Enum []struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
						Para string `xml:"para"`
					} `xml:"enum"`
				} `xml:"enumlist"`
			} `xml:"enum"`
		} `xml:"enumlist"`
		Example struct {
			Text  string `xml:",chardata"`
			Title string `xml:"title,attr"`
		} `xml:"example"`
		Para struct {
			Text    string   `xml:",chardata"`
			Literal []string `xml:"literal"`
		} `xml:"para"`
	} `xml:"info"`
	Application []struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name,attr"`
		Language string `xml:"language,attr"`
		Module   string `xml:"module,attr"`
		Synopsis string `xml:"synopsis"`
		Syntax   struct {
			Text      string `xml:",chardata"`
			Argsep    string `xml:"argsep,attr"`
			Parameter []struct {
				Text      string `xml:",chardata"`
				Name      string `xml:"name,attr"`
				Required  string `xml:"required,attr"`
				Argsep    string `xml:"argsep,attr"`
				Multiple  string `xml:"multiple,attr"`
				Hasparams string `xml:"hasparams,attr"`
				Para      []struct {
					Text        string   `xml:",chardata"`
					Literal     []string `xml:"literal"`
					Directory   string   `xml:"directory"`
					Filename    string   `xml:"filename"`
					Variable    string   `xml:"variable"`
					Replaceable []string `xml:"replaceable"`
					Emphasis    string   `xml:"emphasis"`
				} `xml:"para"`
				Enumlist struct {
					Text string `xml:",chardata"`
					Enum []struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
						Para string `xml:"para"`
					} `xml:"enum"`
				} `xml:"enumlist"`
				Optionlist struct {
					Text   string `xml:",chardata"`
					Option []struct {
						Text      string `xml:",chardata"`
						Name      string `xml:"name,attr"`
						Implies   string `xml:"implies,attr"`
						Argsep    string `xml:"argsep,attr"`
						Hasparams string `xml:"hasparams,attr"`
						Para      []struct {
							Text        string   `xml:",chardata"`
							Literal     []string `xml:"literal"`
							Replaceable []string `xml:"replaceable"`
							Variable    []string `xml:"variable"`
							Emphasis    []string `xml:"emphasis"`
							Filename    string   `xml:"filename"`
							Directory   string   `xml:"directory"`
						} `xml:"para"`
						Argument []struct {
							Text      string `xml:",chardata"`
							Name      string `xml:"name,attr"`
							Required  string `xml:"required,attr"`
							Hasparams string `xml:"hasparams,attr"`
							Argsep    string `xml:"argsep,attr"`
							Multiple  string `xml:"multiple,attr"`
							Para      []struct {
								Text        string   `xml:",chardata"`
								Replaceable string   `xml:"replaceable"`
								Variable    string   `xml:"variable"`
								Literal     []string `xml:"literal"`
								Filename    string   `xml:"filename"`
								Emphasis    string   `xml:"emphasis"`
							} `xml:"para"`
							Argument []struct {
								Text     string `xml:",chardata"`
								Name     string `xml:"name,attr"`
								Multiple string `xml:"multiple,attr"`
								Required string `xml:"required,attr"`
							} `xml:"argument"`
						} `xml:"argument"`
						Enumlist struct {
							Text string `xml:",chardata"`
							Enum []struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
								Para string `xml:"para"`
							} `xml:"enum"`
						} `xml:"enumlist"`
						Warning struct {
							Text string `xml:",chardata"`
							Para struct {
								Text     string   `xml:",chardata"`
								Variable []string `xml:"variable"`
							} `xml:"para"`
						} `xml:"warning"`
						Note []struct {
							Text string `xml:",chardata"`
							Para struct {
								Text        string   `xml:",chardata"`
								Literal     []string `xml:"literal"`
								Replaceable string   `xml:"replaceable"`
								Variable    string   `xml:"variable"`
							} `xml:"para"`
						} `xml:"note"`
						Variablelist struct {
							Text     string `xml:",chardata"`
							Variable []struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
								Para struct {
									Text        string   `xml:",chardata"`
									Replaceable []string `xml:"replaceable"`
								} `xml:"para"`
								Value []struct {
									Text    string `xml:",chardata"`
									Name    string `xml:"name,attr"`
									Default string `xml:"default,attr"`
								} `xml:"value"`
							} `xml:"variable"`
						} `xml:"variablelist"`
					} `xml:"option"`
				} `xml:"optionlist"`
				Argument []struct {
					Text      string `xml:",chardata"`
					Name      string `xml:"name,attr"`
					Argsep    string `xml:"argsep,attr"`
					Required  string `xml:"required,attr"`
					Multiple  string `xml:"multiple,attr"`
					Hasparams string `xml:"hasparams,attr"`
					Argument  []struct {
						Text     string `xml:",chardata"`
						Name     string `xml:"name,attr"`
						Required string `xml:"required,attr"`
						Multiple string `xml:"multiple,attr"`
					} `xml:"argument"`
					Para []struct {
						Text        string   `xml:",chardata"`
						Literal     string   `xml:"literal"`
						Replaceable []string `xml:"replaceable"`
						Filename    string   `xml:"filename"`
					} `xml:"para"`
				} `xml:"argument"`
				Warning struct {
					Text string `xml:",chardata"`
					Para struct {
						Text     string   `xml:",chardata"`
						Variable []string `xml:"variable"`
					} `xml:"para"`
				} `xml:"warning"`
				Note struct {
					Text string `xml:",chardata"`
					Para struct {
						Text        string   `xml:",chardata"`
						Replaceable []string `xml:"replaceable"`
						Emphasis    string   `xml:"emphasis"`
						Literal     string   `xml:"literal"`
					} `xml:"para"`
				} `xml:"note"`
				Include struct {
					Text     string `xml:",chardata"`
					Xpointer string `xml:"xpointer,attr"`
				} `xml:"include"`
			} `xml:"parameter"`
			Include []struct {
				Text     string `xml:",chardata"`
				Xpointer string `xml:"xpointer,attr"`
			} `xml:"include"`
		} `xml:"syntax"`
		Description struct {
			Text string `xml:",chardata"`
			Para []struct {
				Text        string   `xml:",chardata"`
				Replaceable []string `xml:"replaceable"`
				Literal     []string `xml:"literal"`
				Emphasis    string   `xml:"emphasis"`
				Filename    []string `xml:"filename"`
				Variable    []string `xml:"variable"`
				Directory   string   `xml:"directory"`
				Astcli      string   `xml:"astcli"`
			} `xml:"para"`
			Note []struct {
				Text string `xml:",chardata"`
				Para []struct {
					Text        string   `xml:",chardata"`
					Literal     []string `xml:"literal"`
					Replaceable []string `xml:"replaceable"`
					Variable    []string `xml:"variable"`
					Filename    string   `xml:"filename"`
				} `xml:"para"`
			} `xml:"note"`
			Enumlist struct {
				Text string `xml:",chardata"`
				Enum []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					Para []struct {
						Text        string   `xml:",chardata"`
						Literal     []string `xml:"literal"`
						Replaceable string   `xml:"replaceable"`
						Directory   string   `xml:"directory"`
						Filename    string   `xml:"filename"`
					} `xml:"para"`
				} `xml:"enum"`
			} `xml:"enumlist"`
			Variablelist []struct {
				Text     string `xml:",chardata"`
				Variable []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					Para []struct {
						Text     string   `xml:",chardata"`
						Literal  []string `xml:"literal"`
						Emphasis string   `xml:"emphasis"`
						Variable string   `xml:"variable"`
					} `xml:"para"`
					Value []struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
					} `xml:"value"`
				} `xml:"variable"`
			} `xml:"variablelist"`
			Example []struct {
				Text  string `xml:",chardata"`
				Title string `xml:"title,attr"`
			} `xml:"example"`
			Warning []struct {
				Text string `xml:",chardata"`
				Para struct {
					Text     string   `xml:",chardata"`
					Literal  []string `xml:"literal"`
					Variable []string `xml:"variable"`
				} `xml:"para"`
			} `xml:"warning"`
			Include struct {
				Text     string `xml:",chardata"`
				Xpointer string `xml:"xpointer,attr"`
			} `xml:"include"`
		} `xml:"description"`
		SeeAlso struct {
			Text string `xml:",chardata"`
			Ref  []struct {
				Text   string `xml:",chardata"`
				Type   string `xml:"type,attr"`
				Module string `xml:"module,attr"`
			} `xml:"ref"`
		} `xml:"see-also"`
	} `xml:"application"`
	ManagerEvent []struct {
		Text                 string `xml:",chardata"`
		Language             string `xml:"language,attr"`
		Name                 string `xml:"name,attr"`
		ManagerEventInstance struct {
			Text     string `xml:",chardata"`
			Class    string `xml:"class,attr"`
			Synopsis string `xml:"synopsis"`
			Syntax   struct {
				Text            string `xml:",chardata"`
				ChannelSnapshot []struct {
					Text   string `xml:",chardata"`
					Prefix string `xml:"prefix,attr"`
				} `xml:"channel_snapshot"`
				Parameter []struct {
					Text     string `xml:",chardata"`
					Name     string `xml:"name,attr"`
					Required string `xml:"required,attr"`
					Multiple string `xml:"multiple,attr"`
					Para     []struct {
						Text        string   `xml:",chardata"`
						Variable    string   `xml:"variable"`
						Literal     []string `xml:"literal"`
						Replaceable []string `xml:"replaceable"`
						Include     struct {
							Text     string `xml:",chardata"`
							Xpointer string `xml:"xpointer,attr"`
						} `xml:"include"`
					} `xml:"para"`
					Note struct {
						Text string `xml:",chardata"`
						Para struct {
							Text        string   `xml:",chardata"`
							Literal     []string `xml:"literal"`
							Filename    string   `xml:"filename"`
							Replaceable string   `xml:"replaceable"`
						} `xml:"para"`
					} `xml:"note"`
					Enumlist struct {
						Text string `xml:",chardata"`
						Enum []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
							Para []struct {
								Text    string `xml:",chardata"`
								Literal string `xml:"literal"`
							} `xml:"para"`
							Note struct {
								Text string `xml:",chardata"`
								Para struct {
									Text        string   `xml:",chardata"`
									Filename    string   `xml:"filename"`
									Literal     []string `xml:"literal"`
									Replaceable string   `xml:"replaceable"`
								} `xml:"para"`
							} `xml:"note"`
						} `xml:"enum"`
					} `xml:"enumlist"`
				} `xml:"parameter"`
				Include []struct {
					Text     string `xml:",chardata"`
					Xpointer string `xml:"xpointer,attr"`
				} `xml:"include"`
				BridgeSnapshot []struct {
					Text   string `xml:",chardata"`
					Prefix string `xml:"prefix,attr"`
				} `xml:"bridge_snapshot"`
			} `xml:"syntax"`
			SeeAlso struct {
				Text string `xml:",chardata"`
				Ref  []struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"ref"`
			} `xml:"see-also"`
			Description struct {
				Text string `xml:",chardata"`
				Para []struct {
					Text        string   `xml:",chardata"`
					Literal     []string `xml:"literal"`
					Replaceable string   `xml:"replaceable"`
					Filename    string   `xml:"filename"`
				} `xml:"para"`
				Note struct {
					Text string `xml:",chardata"`
					Para struct {
						Text     string `xml:",chardata"`
						Filename string `xml:"filename"`
					} `xml:"para"`
				} `xml:"note"`
			} `xml:"description"`
		} `xml:"managerEventInstance"`
	} `xml:"managerEvent"`
	Agi []struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name,attr"`
		Language string `xml:"language,attr"`
		Synopsis string `xml:"synopsis"`
		Syntax   struct {
			Text      string `xml:",chardata"`
			Parameter []struct {
				Text     string `xml:",chardata"`
				Name     string `xml:"name,attr"`
				Required string `xml:"required,attr"`
				Para     struct {
					Text        string   `xml:",chardata"`
					Literal     string   `xml:"literal"`
					Replaceable []string `xml:"replaceable"`
					Filename    string   `xml:"filename"`
				} `xml:"para"`
				Enumlist struct {
					Text string `xml:",chardata"`
					Enum []struct {
						Text      string `xml:",chardata"`
						Name      string `xml:"name,attr"`
						Parameter struct {
							Text     string `xml:",chardata"`
							Name     string `xml:"name,attr"`
							Literal  string `xml:"literal,attr"`
							Required string `xml:"required,attr"`
						} `xml:"parameter"`
					} `xml:"enum"`
				} `xml:"enumlist"`
			} `xml:"parameter"`
		} `xml:"syntax"`
		Description struct {
			Text string `xml:",chardata"`
			Para []struct {
				Text        string   `xml:",chardata"`
				Literal     []string `xml:"literal"`
				Replaceable []string `xml:"replaceable"`
			} `xml:"para"`
			Enumlist struct {
				Text string `xml:",chardata"`
				Enum []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					Para string `xml:"para"`
				} `xml:"enum"`
			} `xml:"enumlist"`
			Variablelist struct {
				Text     string `xml:",chardata"`
				Variable []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					Para struct {
						Text    string `xml:",chardata"`
						Literal string `xml:"literal"`
					} `xml:"para"`
					Value []struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
					} `xml:"value"`
				} `xml:"variable"`
			} `xml:"variablelist"`
		} `xml:"description"`
		SeeAlso struct {
			Text string `xml:",chardata"`
			Ref  []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"ref"`
		} `xml:"see-also"`
	} `xml:"agi"`
}
