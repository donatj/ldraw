// Code generated by "colorgen"; DO NOT EDIT
package ldraw

import (
	"image/color"
)

// BrickColor is a brick color number as defined by http://www.ldraw.org/library/official/ldcfgalt.ldr
type BrickColor int

const (
	// Black as defined by:
	//    COLOUR Black                              CODE   0   VALUE #05131D   EDGE #3F474C
	Black BrickColor = 0
	// Blue as defined by:
	//    COLOUR Blue                               CODE   1   VALUE #0055BF   EDGE #0561BC
	Blue BrickColor = 1
	// Green as defined by:
	//    COLOUR Green                              CODE   2   VALUE #257A3E   EDGE #238841
	Green BrickColor = 2
	// DarkTurquoise as defined by:
	//    COLOUR Dark_Turquoise                     CODE   3   VALUE #00838F   EDGE #008FAB
	DarkTurquoise BrickColor = 3
	// Red as defined by:
	//    COLOUR Red                                CODE   4   VALUE #C91A09   EDGE #D51A09
	Red BrickColor = 4
	// DarkPink as defined by:
	//    COLOUR Dark_Pink                          CODE   5   VALUE #C870A0   EDGE #C26293
	DarkPink BrickColor = 5
	// Brown as defined by:
	//    COLOUR Brown                              CODE   6   VALUE #583927   EDGE #412D15
	Brown BrickColor = 6
	// LightGrey as defined by:
	//    COLOUR Light_Grey                         CODE   7   VALUE #9BA19D   EDGE #757B7C
	LightGrey BrickColor = 7
	// DarkGrey as defined by:
	//    COLOUR Dark_Grey                          CODE   8   VALUE #6D6E5C   EDGE #4D4B43
	DarkGrey BrickColor = 8
	// LightBlue as defined by:
	//    COLOUR Light_Blue                         CODE   9   VALUE #B4D2E3   EDGE #B4D2F3
	LightBlue BrickColor = 9
	// BrightGreen as defined by:
	//    COLOUR Bright_Green                       CODE  10   VALUE #4B9F4A   EDGE #4BA94A
	BrightGreen BrickColor = 10
	// LightTurquoise as defined by:
	//    COLOUR Light_Turquoise                    CODE  11   VALUE #55A5AF   EDGE #55A5BF
	LightTurquoise BrickColor = 11
	// Salmon as defined by:
	//    COLOUR Salmon                             CODE  12   VALUE #F2705E   EDGE #F2806E
	Salmon BrickColor = 12
	// Pink as defined by:
	//    COLOUR Pink                               CODE  13   VALUE #FC97AC   EDGE #FC97AC
	Pink BrickColor = 13
	// Yellow as defined by:
	//    COLOUR Yellow                             CODE  14   VALUE #F2CD37   EDGE #E8C51F
	Yellow BrickColor = 14
	// White as defined by:
	//    COLOUR White                              CODE  15   VALUE #FFFFFF   EDGE #A9A5A9
	White BrickColor = 15
	// LightGreen as defined by:
	//    COLOUR Light_Green                        CODE  17   VALUE #C2DAB8   EDGE #C2EAC8
	LightGreen BrickColor = 17
	// LightYellow as defined by:
	//    COLOUR Light_Yellow                       CODE  18   VALUE #FBE696   EDGE #FEE9A9
	LightYellow BrickColor = 18
	// Tan as defined by:
	//    COLOUR Tan                                CODE  19   VALUE #E4CD9E   EDGE #E2CBA0
	Tan BrickColor = 19
	// LightViolet as defined by:
	//    COLOUR Light_Violet                       CODE  20   VALUE #C9CAE2   EDGE #CACAEE
	LightViolet BrickColor = 20
	// Purple as defined by:
	//    COLOUR Purple                             CODE  22   VALUE #81007B   EDGE #91008B
	Purple BrickColor = 22
	// DarkBlueViolet as defined by:
	//    COLOUR Dark_Blue_Violet                   CODE  23   VALUE #2032B0   EDGE #2032C0
	DarkBlueViolet BrickColor = 23
	// Orange as defined by:
	//    COLOUR Orange                             CODE  25   VALUE #FE8A18   EDGE #FF8820
	Orange BrickColor = 25
	// Magenta as defined by:
	//    COLOUR Magenta                            CODE  26   VALUE #923978   EDGE #923988
	Magenta BrickColor = 26
	// Lime as defined by:
	//    COLOUR Lime                               CODE  27   VALUE #BBE90B   EDGE #C0ED00
	Lime BrickColor = 27
	// DarkTan as defined by:
	//    COLOUR Dark_Tan                           CODE  28   VALUE #958A73   EDGE #756A53
	DarkTan BrickColor = 28
	// BrightPink as defined by:
	//    COLOUR Bright_Pink                        CODE  29   VALUE #E4ADC8   EDGE #F4ADC8
	BrightPink BrickColor = 29
	// MediumLavender as defined by:
	//    COLOUR Medium_Lavender                    CODE  30   VALUE #AC78BA   EDGE #AF7CBE
	MediumLavender BrickColor = 30
	// Lavender as defined by:
	//    COLOUR Lavender                           CODE  31   VALUE #E1D5ED   EDGE #E5D9EF
	Lavender BrickColor = 31
	// VeryLightOrange as defined by:
	//    COLOUR Very_Light_Orange                  CODE  68   VALUE #F3CF9B   EDGE #F9DFAB
	VeryLightOrange BrickColor = 68
	// BrightReddishLilac as defined by:
	//    COLOUR Bright_Reddish_Lilac               CODE  69   VALUE #CD6298   EDGE #DD72A8
	BrightReddishLilac BrickColor = 69
	// ReddishBrown as defined by:
	//    COLOUR Reddish_Brown                      CODE  70   VALUE #582A12   EDGE #391A08
	ReddishBrown BrickColor = 70
	// LightBluishGrey as defined by:
	//    COLOUR Light_Bluish_Grey                  CODE  71   VALUE #A0A5A9   EDGE #777A85
	LightBluishGrey BrickColor = 71
	// DarkBluishGrey as defined by:
	//    COLOUR Dark_Bluish_Grey                   CODE  72   VALUE #6C6E68   EDGE #484A4B
	DarkBluishGrey BrickColor = 72
	// MediumBlue as defined by:
	//    COLOUR Medium_Blue                        CODE  73   VALUE #5C9DD1   EDGE #5A95DE
	MediumBlue BrickColor = 73
	// MediumGreen as defined by:
	//    COLOUR Medium_Green                       CODE  74   VALUE #73DCA1   EDGE #A1D390
	MediumGreen BrickColor = 74
	// LightPink as defined by:
	//    COLOUR Light_Pink                         CODE  77   VALUE #FECCCF   EDGE #FFCED2
	LightPink BrickColor = 77
	// LightFlesh as defined by:
	//    COLOUR Light_Flesh                        CODE  78   VALUE #F6D7B3   EDGE #F8D9B5
	LightFlesh BrickColor = 78
	// MediumDarkFlesh as defined by:
	//    COLOUR Medium_Dark_Flesh                  CODE  84   VALUE #CC702A   EDGE #CE732C
	MediumDarkFlesh BrickColor = 84
	// MediumLilac as defined by:
	//    COLOUR Medium_Lilac                       CODE  85   VALUE #3F3691   EDGE #4238A4
	MediumLilac BrickColor = 85
	// DarkFlesh as defined by:
	//    COLOUR Dark_Flesh                         CODE  86   VALUE #7C503A   EDGE #7E533C
	DarkFlesh BrickColor = 86
	// BlueViolet as defined by:
	//    COLOUR Blue_Violet                        CODE  89   VALUE #4C61DB   EDGE #4F61E6
	BlueViolet BrickColor = 89
	// Flesh as defined by:
	//    COLOUR Flesh                              CODE  92   VALUE #D09168   EDGE #D3946A
	Flesh BrickColor = 92
	// LightSalmon as defined by:
	//    COLOUR Light_Salmon                       CODE 100   VALUE #FEBABD   EDGE #FFBCBF
	LightSalmon BrickColor = 100
	// Violet as defined by:
	//    COLOUR Violet                             CODE 110   VALUE #4354A3   EDGE #4354B3
	Violet BrickColor = 110
	// MediumViolet as defined by:
	//    COLOUR Medium_Violet                      CODE 112   VALUE #6874CA   EDGE #6A78D4
	MediumViolet BrickColor = 112
	// MediumLime as defined by:
	//    COLOUR Medium_Lime                        CODE 115   VALUE #C7D23C   EDGE #C9D43E
	MediumLime BrickColor = 115
	// Aqua as defined by:
	//    COLOUR Aqua                               CODE 118   VALUE #B3D7D1   EDGE #B4DAD3
	Aqua BrickColor = 118
	// LightLime as defined by:
	//    COLOUR Light_Lime                         CODE 120   VALUE #D9E4A7   EDGE #DAE9A9
	LightLime BrickColor = 120
	// LightOrange as defined by:
	//    COLOUR Light_Orange                       CODE 125   VALUE #F9BA61   EDGE #FDBF5D
	LightOrange BrickColor = 125
	// VeryLightBluishGrey as defined by:
	//    COLOUR Very_Light_Bluish_Grey             CODE 151   VALUE #E6E3E0   EDGE #E9E6E5
	VeryLightBluishGrey BrickColor = 151
	// BrightLightOrange as defined by:
	//    COLOUR Bright_Light_Orange                CODE 191   VALUE #F8BB3D   EDGE #F8BB3D
	BrightLightOrange BrickColor = 191
	// BrightLightBlue as defined by:
	//    COLOUR Bright_Light_Blue                  CODE 212   VALUE #86C1E1   EDGE #89C5E6
	BrightLightBlue BrickColor = 212
	// Rust as defined by:
	//    COLOUR Rust                               CODE 216   VALUE #B31004   EDGE #B91205
	Rust BrickColor = 216
	// BrightLightYellow as defined by:
	//    COLOUR Bright_Light_Yellow                CODE 226   VALUE #FFF03A   EDGE #FFF13B
	BrightLightYellow BrickColor = 226
	// SkyBlue as defined by:
	//    COLOUR Sky_Blue                           CODE 232   VALUE #56BED6   EDGE #59C3DA
	SkyBlue BrickColor = 232
	// DarkBlue as defined by:
	//    COLOUR Dark_Blue                          CODE 272   VALUE #0D325B   EDGE #083469
	DarkBlue BrickColor = 272
	// DarkGreen as defined by:
	//    COLOUR Dark_Green                         CODE 288   VALUE #184632   EDGE #184A25
	DarkGreen BrickColor = 288
	// DarkBrown as defined by:
	//    COLOUR Dark_Brown                         CODE 308   VALUE #352100   EDGE #3C2807
	DarkBrown BrickColor = 308
	// MaerskBlue as defined by:
	//    COLOUR Maersk_Blue                        CODE 313   VALUE #54A9C8   EDGE #56ABCD
	MaerskBlue BrickColor = 313
	// DarkRed as defined by:
	//    COLOUR Dark_Red                           CODE 320   VALUE #720E0F   EDGE #790E0F
	DarkRed BrickColor = 320
	// DarkAzure as defined by:
	//    COLOUR Dark_Azure                         CODE 321   VALUE #1498D7   EDGE #088DCD
	DarkAzure BrickColor = 321
	// MediumAzure as defined by:
	//    COLOUR Medium_Azure                       CODE 322   VALUE #3EC2DD   EDGE #3AB2C5
	MediumAzure BrickColor = 322
	// LightAqua as defined by:
	//    COLOUR Light_Aqua                         CODE 323   VALUE #BDDCD8   EDGE #A0DEDA
	LightAqua BrickColor = 323
	// YellowishGreen as defined by:
	//    COLOUR Yellowish_Green                    CODE 326   VALUE #DFEEA5   EDGE #E5F6AB
	YellowishGreen BrickColor = 326
	// OliveGreen as defined by:
	//    COLOUR Olive_Green                        CODE 330   VALUE #9B9A5A   EDGE #9E9D5E
	OliveGreen BrickColor = 330
	// SandRed as defined by:
	//    COLOUR Sand_Red                           CODE 335   VALUE #D67572   EDGE #DA7876
	SandRed BrickColor = 335
	// MediumDarkPink as defined by:
	//    COLOUR Medium_Dark_Pink                   CODE 351   VALUE #F785B1   EDGE #FB89B8
	MediumDarkPink BrickColor = 351
	// EarthOrange as defined by:
	//    COLOUR Earth_Orange                       CODE 366   VALUE #FA9C1C   EDGE #FEA11E
	EarthOrange BrickColor = 366
	// SandPurple as defined by:
	//    COLOUR Sand_Purple                        CODE 373   VALUE #845E84   EDGE #8A648B
	SandPurple BrickColor = 373
	// SandGreen as defined by:
	//    COLOUR Sand_Green                         CODE 378   VALUE #A0BCAC   EDGE #A0BFAC
	SandGreen BrickColor = 378
	// SandBlue as defined by:
	//    COLOUR Sand_Blue                          CODE 379   VALUE #597184   EDGE #5B7488
	SandBlue BrickColor = 379
	// FabulandBrown as defined by:
	//    COLOUR Fabuland_Brown                     CODE 450   VALUE #B67B50   EDGE #B77B52
	FabulandBrown BrickColor = 450
	// MediumOrange as defined by:
	//    COLOUR Medium_Orange                      CODE 462   VALUE #FFA70B   EDGE #FFAA0F
	MediumOrange BrickColor = 462
	// DarkOrange as defined by:
	//    COLOUR Dark_Orange                        CODE 484   VALUE #A95500   EDGE #AD5906
	DarkOrange BrickColor = 484
	// VeryLightGrey as defined by:
	//    COLOUR Very_Light_Grey                    CODE 503   VALUE #E6E3DA   EDGE #E9E6DD
	VeryLightGrey BrickColor = 503
	// ReddishLilac as defined by:
	//    COLOUR Reddish_Lilac                      CODE 218   VALUE #8E5597   EDGE #9966A1
	ReddishLilac BrickColor = 218
	// FlamingoPink as defined by:
	//    COLOUR Flamingo_Pink                      CODE 295   VALUE #FF94C2   EDGE #FF9FC8
	FlamingoPink BrickColor = 295
	// Lilac as defined by:
	//    COLOUR Lilac                              CODE 219   VALUE #564E9D   EDGE #6760A7
	Lilac BrickColor = 219
	// DarkNougat as defined by:
	//    COLOUR Dark_Nougat                        CODE 128   VALUE #AD6140   EDGE #9C573A
	DarkNougat BrickColor = 128
	// TransClear as defined by:
	//    COLOUR Trans_Clear                        CODE  47   VALUE #FCFCFC   EDGE #A9ABB7   ALPHA 128
	TransClear BrickColor = 47
	// TransBlack as defined by:
	//    COLOUR Trans_Black                        CODE  40   VALUE #635F52   EDGE #171316   ALPHA 128
	TransBlack BrickColor = 40
	// TransRed as defined by:
	//    COLOUR Trans_Red                          CODE  36   VALUE #C91A09   EDGE #D91A09   ALPHA 128
	TransRed BrickColor = 36
	// TransNeonOrange as defined by:
	//    COLOUR Trans_Neon_Orange                  CODE  38   VALUE #FF800D   EDGE #FF7D10   ALPHA 128
	TransNeonOrange BrickColor = 38
	// TransOrange as defined by:
	//    COLOUR Trans_Orange                       CODE  57   VALUE #F08F1C   EDGE #ED8B1A   ALPHA 128
	TransOrange BrickColor = 57
	// TransNeonYellow as defined by:
	//    COLOUR Trans_Neon_Yellow                  CODE  54   VALUE #DAB000   EDGE #F5CD2F   ALPHA 128
	TransNeonYellow BrickColor = 54
	// TransYellow as defined by:
	//    COLOUR Trans_Yellow                       CODE  46   VALUE #F5CD2F   EDGE #B49819   ALPHA 128
	TransYellow BrickColor = 46
	// TransNeonGreen as defined by:
	//    COLOUR Trans_Neon_Green                   CODE  42   VALUE #C0FF00   EDGE #84C300   ALPHA 128
	TransNeonGreen BrickColor = 42
	// TransBrightGreen as defined by:
	//    COLOUR Trans_Bright_Green                 CODE  35   VALUE #56E646   EDGE #9DA86B   ALPHA 128
	TransBrightGreen BrickColor = 35
	// TransGreen as defined by:
	//    COLOUR Trans_Green                        CODE  34   VALUE #237841   EDGE #237841   ALPHA 128
	TransGreen BrickColor = 34
	// TransDarkBlue as defined by:
	//    COLOUR Trans_Dark_Blue                    CODE  33   VALUE #0020A0   EDGE #000064   ALPHA 128
	TransDarkBlue BrickColor = 33
	// TransMediumBlue as defined by:
	//    COLOUR Trans_Medium_Blue                  CODE  41   VALUE #559AB7   EDGE #196273   ALPHA 128
	TransMediumBlue BrickColor = 41
	// TransLightBlue as defined by:
	//    COLOUR Trans_Light_Blue                   CODE  43   VALUE #AEE9EF   EDGE #72B3B8   ALPHA 128
	TransLightBlue BrickColor = 43
	// TransVeryLightBlue as defined by:
	//    COLOUR Trans_Very_Light_Blue              CODE  39   VALUE #C1DFF0   EDGE #85A3B4   ALPHA 128
	TransVeryLightBlue BrickColor = 39
	// TransBrightReddishLilac as defined by:
	//    COLOUR Trans_Bright_Reddish_Lilac         CODE  44   VALUE #96709F   EDGE #5A3463   ALPHA 128
	TransBrightReddishLilac BrickColor = 44
	// TransPurple as defined by:
	//    COLOUR Trans_Purple                       CODE  52   VALUE #A5A5CB   EDGE #6D6E5C   ALPHA 128
	TransPurple BrickColor = 52
	// TransDarkPink as defined by:
	//    COLOUR Trans_Dark_Pink                    CODE  37   VALUE #DF6695   EDGE #A32A59   ALPHA 128
	TransDarkPink BrickColor = 37
	// TransPink as defined by:
	//    COLOUR Trans_Pink                         CODE  45   VALUE #FC97AC   EDGE #B8718C   ALPHA 128
	TransPink BrickColor = 45
	// TransLightGreen as defined by:
	//    COLOUR Trans_Light_Green                  CODE 285   VALUE #7DC291   EDGE #52805F   ALPHA 128
	TransLightGreen BrickColor = 285
	// TransFireYellow as defined by:
	//    COLOUR Trans_Fire_Yellow                  CODE 234   VALUE #FBE890   EDGE #BAAB6A   ALPHA 128
	TransFireYellow BrickColor = 234
	// TransLightBlueViolet as defined by:
	//    COLOUR Trans_Light_Blue_Violet            CODE 293   VALUE #6BABE4   EDGE #4D7BA3   ALPHA 128
	TransLightBlueViolet BrickColor = 293
	// TransBrightLightOrange as defined by:
	//    COLOUR Trans_Bright_Light_Orange          CODE 231   VALUE #FCB76D   EDGE #BD8951   ALPHA 128
	TransBrightLightOrange BrickColor = 231
	// TransReddishLilac as defined by:
	//    COLOUR Trans_Reddish_Lilac                CODE 284   VALUE #c281A5   EDGE #82566E   ALPHA 128
	TransReddishLilac BrickColor = 284
	// ChromeGold as defined by:
	//    COLOUR Chrome_Gold                        CODE 334   VALUE #BBA53D   EDGE #C2AB44
	ChromeGold BrickColor = 334
	// ChromeSilver as defined by:
	//    COLOUR Chrome_Silver                      CODE 383   VALUE #E0E0E0   EDGE #A4A4A4
	ChromeSilver BrickColor = 383
	// ChromeAntiqueBrass as defined by:
	//    COLOUR Chrome_Antique_Brass               CODE  60   VALUE #645A4C   EDGE #281E10
	ChromeAntiqueBrass BrickColor = 60
	// ChromeBlack as defined by:
	//    COLOUR Chrome_Black                       CODE  64   VALUE #1B2A34   EDGE #000000
	ChromeBlack BrickColor = 64
	// ChromeBlue as defined by:
	//    COLOUR Chrome_Blue                        CODE  61   VALUE #6C96BF   EDGE #202A68
	ChromeBlue BrickColor = 61
	// ChromeGreen as defined by:
	//    COLOUR Chrome_Green                       CODE  62   VALUE #3CB371   EDGE #007735
	ChromeGreen BrickColor = 62
	// ChromePink as defined by:
	//    COLOUR Chrome_Pink                        CODE  63   VALUE #AA4D8E   EDGE #6E1152
	ChromePink BrickColor = 63
	// PearlWhite as defined by:
	//    COLOUR Pearl_White                        CODE 183   VALUE #F2F3F2   EDGE #FFFFFF
	PearlWhite BrickColor = 183
	// PearlVeryLightGrey as defined by:
	//    COLOUR Pearl_Very_Light_Grey              CODE 150   VALUE #BBBDBC   EDGE #C3C7C5
	PearlVeryLightGrey BrickColor = 150
	// PearlLightGrey as defined by:
	//    COLOUR Pearl_Light_Grey                   CODE 135   VALUE #9CA3A8   EDGE #6C7378
	PearlLightGrey BrickColor = 135
	// FlatSilver as defined by:
	//    COLOUR Flat_Silver                        CODE 179   VALUE #898788   EDGE #696768
	FlatSilver BrickColor = 179
	// PearlDarkGrey as defined by:
	//    COLOUR Pearl_Dark_Grey                    CODE 148   VALUE #575857   EDGE #424342
	PearlDarkGrey BrickColor = 148
	// MetalBlue as defined by:
	//    COLOUR Metal_Blue                         CODE 137   VALUE #5677BA   EDGE #6F8ED4
	MetalBlue BrickColor = 137
	// PearlLightGold as defined by:
	//    COLOUR Pearl_Light_Gold                   CODE 142   VALUE #DCBE61   EDGE #DFBF64
	PearlLightGold BrickColor = 142
	// PearlGold as defined by:
	//    COLOUR Pearl_Gold                         CODE 297   VALUE #CC9C2B   EDGE #DDAE36
	PearlGold BrickColor = 297
	// FlatDarkGold as defined by:
	//    COLOUR Flat_Dark_Gold                     CODE 178   VALUE #B4883E   EDGE #B1843A
	FlatDarkGold BrickColor = 178
	// Copper as defined by:
	//    COLOUR Copper                             CODE 134   VALUE #964A27   EDGE #A25331
	Copper BrickColor = 134
	// ReddishGold as defined by:
	//    COLOUR Reddish_Gold                       CODE 189   VALUE #AC8247   EDGE #9B7540
	ReddishGold BrickColor = 189
	// MetallicSilver as defined by:
	//    COLOUR Metallic_Silver                    CODE  80   VALUE #A5A9B4   EDGE #A6AAB8
	MetallicSilver BrickColor = 80
	// MetallicGreen as defined by:
	//    COLOUR Metallic_Green                     CODE  81   VALUE #899B5F   EDGE #8EA15F
	MetallicGreen BrickColor = 81
	// MetallicGold as defined by:
	//    COLOUR Metallic_Gold                      CODE  82   VALUE #DBAC34   EDGE #CD9E56
	MetallicGold BrickColor = 82
	// MetallicBlack as defined by:
	//    COLOUR Metallic_Black                     CODE  83   VALUE #1A2831   EDGE #000000
	MetallicBlack BrickColor = 83
	// MetallicDarkGrey as defined by:
	//    COLOUR Metallic_Dark_Grey                 CODE  87   VALUE #6D6E5C   EDGE #5D5B53
	MetallicDarkGrey BrickColor = 87
	// MetallicCopper as defined by:
	//    COLOUR Metallic_Copper                    CODE 300   VALUE #C27F53   EDGE #AF724B
	MetallicCopper BrickColor = 300
	// MetallicBrightRed as defined by:
	//    COLOUR Metallic_Bright_Red                CODE 184   VALUE #D60026   EDGE #CC001C
	MetallicBrightRed BrickColor = 184
	// MetallicDarkGreen as defined by:
	//    COLOUR Metallic_Dark_Green                CODE 186   VALUE #008E3C   EDGE #008432
	MetallicDarkGreen BrickColor = 186
	// MilkyWhite as defined by:
	//    COLOUR Milky_White                        CODE  79   VALUE #FFFFFF   EDGE #C3C3C3   ALPHA 240
	MilkyWhite BrickColor = 79
	// GlowInDarkOpaque as defined by:
	//    COLOUR Glow_In_Dark_Opaque                CODE  21   VALUE #E0FFB0   EDGE #A4C374   ALPHA 240
	GlowInDarkOpaque BrickColor = 21
	// GlowInDarkTrans as defined by:
	//    COLOUR Glow_In_Dark_Trans                 CODE 294   VALUE #BDC6AD   EDGE #CFDBBF   ALPHA 240
	GlowInDarkTrans BrickColor = 294
	// GlowInDarkWhite as defined by:
	//    COLOUR Glow_In_Dark_White                 CODE 329   VALUE #F5F3D7   EDGE #B5B49F   ALPHA 240
	GlowInDarkWhite BrickColor = 329
	// GlitterTransDarkPink as defined by:
	//    COLOUR Glitter_Trans_Dark_Pink            CODE 114   VALUE #DF6695   EDGE #9A2A66   ALPHA 128
	GlitterTransDarkPink BrickColor = 114
	// GlitterTransClear as defined by:
	//    COLOUR Glitter_Trans_Clear                CODE 117   VALUE #FFFFFF   EDGE #C3C3C3   ALPHA 128
	GlitterTransClear BrickColor = 117
	// GlitterTransPurple as defined by:
	//    COLOUR Glitter_Trans_Purple               CODE 129   VALUE #640061   EDGE #280025   ALPHA 128
	GlitterTransPurple BrickColor = 129
	// GlitterTransLightBlue as defined by:
	//    COLOUR Glitter_Trans_Light_Blue           CODE 302   VALUE #AEE9EF   EDGE #72B3B0   ALPHA 128
	GlitterTransLightBlue BrickColor = 302
	// GlitterTransNeonGreen as defined by:
	//    COLOUR Glitter_Trans_Neon_Green           CODE 339   VALUE #C0FF00   EDGE #84C300   ALPHA 128
	GlitterTransNeonGreen BrickColor = 339
	// SpeckleBlackSilver as defined by:
	//    COLOUR Speckle_Black_Silver               CODE 132   VALUE #000000   EDGE #898788
	SpeckleBlackSilver BrickColor = 132
	// SpeckleBlackGold as defined by:
	//    COLOUR Speckle_Black_Gold                 CODE 133   VALUE #000000   EDGE #DBAC34
	SpeckleBlackGold BrickColor = 133
	// SpeckleBlackCopper as defined by:
	//    COLOUR Speckle_Black_Copper               CODE  75   VALUE #000000   EDGE #AB6038
	SpeckleBlackCopper BrickColor = 75
	// SpeckleDarkBluishGreySilver as defined by:
	//    COLOUR Speckle_Dark_Bluish_Grey_Silver    CODE  76   VALUE #635F61   EDGE #898788
	SpeckleDarkBluishGreySilver BrickColor = 76
	// RubberYellow as defined by:
	//    COLOUR Rubber_Yellow                      CODE  65   VALUE #F5CD2F   EDGE #B19705
	RubberYellow BrickColor = 65
	// RubberTransYellow as defined by:
	//    COLOUR Rubber_Trans_Yellow                CODE  66   VALUE #CAB000   EDGE #AD9600   ALPHA 128
	RubberTransYellow BrickColor = 66
	// RubberTransClear as defined by:
	//    COLOUR Rubber_Trans_Clear                 CODE  67   VALUE #FFFFFF   EDGE #C3C3C3   ALPHA 128
	RubberTransClear BrickColor = 67
	// RubberBlack as defined by:
	//    COLOUR Rubber_Black                       CODE 256   VALUE #212121   EDGE #595959
	RubberBlack BrickColor = 256
	// RubberBlue as defined by:
	//    COLOUR Rubber_Blue                        CODE 273   VALUE #0033B2   EDGE #001392
	RubberBlue BrickColor = 273
	// RubberRed as defined by:
	//    COLOUR Rubber_Red                         CODE 324   VALUE #C40026   EDGE #8A0000
	RubberRed BrickColor = 324
	// RubberOrange as defined by:
	//    COLOUR Rubber_Orange                      CODE 350   VALUE #D06610   EDGE #BB5C0E
	RubberOrange BrickColor = 350
	// RubberLightGrey as defined by:
	//    COLOUR Rubber_Light_Grey                  CODE 375   VALUE #C1C2C1   EDGE #696969
	RubberLightGrey BrickColor = 375
	// RubberDarkBlue as defined by:
	//    COLOUR Rubber_Dark_Blue                   CODE 406   VALUE #001D68   EDGE #000A48
	RubberDarkBlue BrickColor = 406
	// RubberPurple as defined by:
	//    COLOUR Rubber_Purple                      CODE 449   VALUE #81007B   EDGE #570052
	RubberPurple BrickColor = 449
	// RubberLime as defined by:
	//    COLOUR Rubber_Lime                        CODE 490   VALUE #D7F000   EDGE #639300
	RubberLime BrickColor = 490
	// RubberLightBluishGrey as defined by:
	//    COLOUR Rubber_Light_Bluish_Grey           CODE 496   VALUE #A3A2A4   EDGE #787878
	RubberLightBluishGrey BrickColor = 496
	// RubberFlatSilver as defined by:
	//    COLOUR Rubber_Flat_Silver                 CODE 504   VALUE #898788   EDGE #494748
	RubberFlatSilver BrickColor = 504
	// RubberWhite as defined by:
	//    COLOUR Rubber_White                       CODE 511   VALUE #FAFAFA   EDGE #676767
	RubberWhite BrickColor = 511
	// MainColour as defined by:
	//    COLOUR Main_Colour                        CODE  16   VALUE #FFFF80   EDGE #333333
	MainColour BrickColor = 16
	// EdgeColour as defined by:
	//    COLOUR Edge_Colour                        CODE  24   VALUE #7F7F7F   EDGE #333333
	EdgeColour BrickColor = 24
	// TransBlackIRLens as defined by:
	//    COLOUR Trans_Black_IR_Lens                CODE  32   VALUE #000000   EDGE #05131D   ALPHA 210
	TransBlackIRLens BrickColor = 32
	// Magnet as defined by:
	//    COLOUR Magnet                             CODE 493   VALUE #656761   EDGE #595959
	Magnet BrickColor = 493
	// ElectricContactAlloy as defined by:
	//    COLOUR Electric_Contact_Alloy             CODE 494   VALUE #D0D0D0   EDGE #6E6E6E
	ElectricContactAlloy BrickColor = 494
	// ElectricContactCopper as defined by:
	//    COLOUR Electric_Contact_Copper            CODE 495   VALUE #AE7A59   EDGE #723E1D
	ElectricContactCopper BrickColor = 495
)

// BrickColorMap connects BrickColor's to their coresponding color.Color
var BrickColorMap = map[BrickColor]color.Color{
	Black:                       color.RGBA{R: 0x5, G: 0x13, B: 0x1d, A: 0xff},
	Blue:                        color.RGBA{R: 0x0, G: 0x55, B: 0xbf, A: 0xff},
	Green:                       color.RGBA{R: 0x25, G: 0x7a, B: 0x3e, A: 0xff},
	DarkTurquoise:               color.RGBA{R: 0x0, G: 0x83, B: 0x8f, A: 0xff},
	Red:                         color.RGBA{R: 0xc9, G: 0x1a, B: 0x9, A: 0xff},
	DarkPink:                    color.RGBA{R: 0xc8, G: 0x70, B: 0xa0, A: 0xff},
	Brown:                       color.RGBA{R: 0x58, G: 0x39, B: 0x27, A: 0xff},
	LightGrey:                   color.RGBA{R: 0x9b, G: 0xa1, B: 0x9d, A: 0xff},
	DarkGrey:                    color.RGBA{R: 0x6d, G: 0x6e, B: 0x5c, A: 0xff},
	LightBlue:                   color.RGBA{R: 0xb4, G: 0xd2, B: 0xe3, A: 0xff},
	BrightGreen:                 color.RGBA{R: 0x4b, G: 0x9f, B: 0x4a, A: 0xff},
	LightTurquoise:              color.RGBA{R: 0x55, G: 0xa5, B: 0xaf, A: 0xff},
	Salmon:                      color.RGBA{R: 0xf2, G: 0x70, B: 0x5e, A: 0xff},
	Pink:                        color.RGBA{R: 0xfc, G: 0x97, B: 0xac, A: 0xff},
	Yellow:                      color.RGBA{R: 0xf2, G: 0xcd, B: 0x37, A: 0xff},
	White:                       color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
	LightGreen:                  color.RGBA{R: 0xc2, G: 0xda, B: 0xb8, A: 0xff},
	LightYellow:                 color.RGBA{R: 0xfb, G: 0xe6, B: 0x96, A: 0xff},
	Tan:                         color.RGBA{R: 0xe4, G: 0xcd, B: 0x9e, A: 0xff},
	LightViolet:                 color.RGBA{R: 0xc9, G: 0xca, B: 0xe2, A: 0xff},
	Purple:                      color.RGBA{R: 0x81, G: 0x0, B: 0x7b, A: 0xff},
	DarkBlueViolet:              color.RGBA{R: 0x20, G: 0x32, B: 0xb0, A: 0xff},
	Orange:                      color.RGBA{R: 0xfe, G: 0x8a, B: 0x18, A: 0xff},
	Magenta:                     color.RGBA{R: 0x92, G: 0x39, B: 0x78, A: 0xff},
	Lime:                        color.RGBA{R: 0xbb, G: 0xe9, B: 0xb, A: 0xff},
	DarkTan:                     color.RGBA{R: 0x95, G: 0x8a, B: 0x73, A: 0xff},
	BrightPink:                  color.RGBA{R: 0xe4, G: 0xad, B: 0xc8, A: 0xff},
	MediumLavender:              color.RGBA{R: 0xac, G: 0x78, B: 0xba, A: 0xff},
	Lavender:                    color.RGBA{R: 0xe1, G: 0xd5, B: 0xed, A: 0xff},
	VeryLightOrange:             color.RGBA{R: 0xf3, G: 0xcf, B: 0x9b, A: 0xff},
	BrightReddishLilac:          color.RGBA{R: 0xcd, G: 0x62, B: 0x98, A: 0xff},
	ReddishBrown:                color.RGBA{R: 0x58, G: 0x2a, B: 0x12, A: 0xff},
	LightBluishGrey:             color.RGBA{R: 0xa0, G: 0xa5, B: 0xa9, A: 0xff},
	DarkBluishGrey:              color.RGBA{R: 0x6c, G: 0x6e, B: 0x68, A: 0xff},
	MediumBlue:                  color.RGBA{R: 0x5c, G: 0x9d, B: 0xd1, A: 0xff},
	MediumGreen:                 color.RGBA{R: 0x73, G: 0xdc, B: 0xa1, A: 0xff},
	LightPink:                   color.RGBA{R: 0xfe, G: 0xcc, B: 0xcf, A: 0xff},
	LightFlesh:                  color.RGBA{R: 0xf6, G: 0xd7, B: 0xb3, A: 0xff},
	MediumDarkFlesh:             color.RGBA{R: 0xcc, G: 0x70, B: 0x2a, A: 0xff},
	MediumLilac:                 color.RGBA{R: 0x3f, G: 0x36, B: 0x91, A: 0xff},
	DarkFlesh:                   color.RGBA{R: 0x7c, G: 0x50, B: 0x3a, A: 0xff},
	BlueViolet:                  color.RGBA{R: 0x4c, G: 0x61, B: 0xdb, A: 0xff},
	Flesh:                       color.RGBA{R: 0xd0, G: 0x91, B: 0x68, A: 0xff},
	LightSalmon:                 color.RGBA{R: 0xfe, G: 0xba, B: 0xbd, A: 0xff},
	Violet:                      color.RGBA{R: 0x43, G: 0x54, B: 0xa3, A: 0xff},
	MediumViolet:                color.RGBA{R: 0x68, G: 0x74, B: 0xca, A: 0xff},
	MediumLime:                  color.RGBA{R: 0xc7, G: 0xd2, B: 0x3c, A: 0xff},
	Aqua:                        color.RGBA{R: 0xb3, G: 0xd7, B: 0xd1, A: 0xff},
	LightLime:                   color.RGBA{R: 0xd9, G: 0xe4, B: 0xa7, A: 0xff},
	LightOrange:                 color.RGBA{R: 0xf9, G: 0xba, B: 0x61, A: 0xff},
	VeryLightBluishGrey:         color.RGBA{R: 0xe6, G: 0xe3, B: 0xe0, A: 0xff},
	BrightLightOrange:           color.RGBA{R: 0xf8, G: 0xbb, B: 0x3d, A: 0xff},
	BrightLightBlue:             color.RGBA{R: 0x86, G: 0xc1, B: 0xe1, A: 0xff},
	Rust:                        color.RGBA{R: 0xb3, G: 0x10, B: 0x4, A: 0xff},
	BrightLightYellow:           color.RGBA{R: 0xff, G: 0xf0, B: 0x3a, A: 0xff},
	SkyBlue:                     color.RGBA{R: 0x56, G: 0xbe, B: 0xd6, A: 0xff},
	DarkBlue:                    color.RGBA{R: 0xd, G: 0x32, B: 0x5b, A: 0xff},
	DarkGreen:                   color.RGBA{R: 0x18, G: 0x46, B: 0x32, A: 0xff},
	DarkBrown:                   color.RGBA{R: 0x35, G: 0x21, B: 0x0, A: 0xff},
	MaerskBlue:                  color.RGBA{R: 0x54, G: 0xa9, B: 0xc8, A: 0xff},
	DarkRed:                     color.RGBA{R: 0x72, G: 0xe, B: 0xf, A: 0xff},
	DarkAzure:                   color.RGBA{R: 0x14, G: 0x98, B: 0xd7, A: 0xff},
	MediumAzure:                 color.RGBA{R: 0x3e, G: 0xc2, B: 0xdd, A: 0xff},
	LightAqua:                   color.RGBA{R: 0xbd, G: 0xdc, B: 0xd8, A: 0xff},
	YellowishGreen:              color.RGBA{R: 0xdf, G: 0xee, B: 0xa5, A: 0xff},
	OliveGreen:                  color.RGBA{R: 0x9b, G: 0x9a, B: 0x5a, A: 0xff},
	SandRed:                     color.RGBA{R: 0xd6, G: 0x75, B: 0x72, A: 0xff},
	MediumDarkPink:              color.RGBA{R: 0xf7, G: 0x85, B: 0xb1, A: 0xff},
	EarthOrange:                 color.RGBA{R: 0xfa, G: 0x9c, B: 0x1c, A: 0xff},
	SandPurple:                  color.RGBA{R: 0x84, G: 0x5e, B: 0x84, A: 0xff},
	SandGreen:                   color.RGBA{R: 0xa0, G: 0xbc, B: 0xac, A: 0xff},
	SandBlue:                    color.RGBA{R: 0x59, G: 0x71, B: 0x84, A: 0xff},
	FabulandBrown:               color.RGBA{R: 0xb6, G: 0x7b, B: 0x50, A: 0xff},
	MediumOrange:                color.RGBA{R: 0xff, G: 0xa7, B: 0xb, A: 0xff},
	DarkOrange:                  color.RGBA{R: 0xa9, G: 0x55, B: 0x0, A: 0xff},
	VeryLightGrey:               color.RGBA{R: 0xe6, G: 0xe3, B: 0xda, A: 0xff},
	ReddishLilac:                color.RGBA{R: 0x8e, G: 0x55, B: 0x97, A: 0xff},
	FlamingoPink:                color.RGBA{R: 0xff, G: 0x94, B: 0xc2, A: 0xff},
	Lilac:                       color.RGBA{R: 0x56, G: 0x4e, B: 0x9d, A: 0xff},
	DarkNougat:                  color.RGBA{R: 0xad, G: 0x61, B: 0x40, A: 0xff},
	TransClear:                  color.RGBA{R: 0xfc, G: 0xfc, B: 0xfc, A: 0x80},
	TransBlack:                  color.RGBA{R: 0x63, G: 0x5f, B: 0x52, A: 0x80},
	TransRed:                    color.RGBA{R: 0xc9, G: 0x1a, B: 0x9, A: 0x80},
	TransNeonOrange:             color.RGBA{R: 0xff, G: 0x80, B: 0xd, A: 0x80},
	TransOrange:                 color.RGBA{R: 0xf0, G: 0x8f, B: 0x1c, A: 0x80},
	TransNeonYellow:             color.RGBA{R: 0xda, G: 0xb0, B: 0x0, A: 0x80},
	TransYellow:                 color.RGBA{R: 0xf5, G: 0xcd, B: 0x2f, A: 0x80},
	TransNeonGreen:              color.RGBA{R: 0xc0, G: 0xff, B: 0x0, A: 0x80},
	TransBrightGreen:            color.RGBA{R: 0x56, G: 0xe6, B: 0x46, A: 0x80},
	TransGreen:                  color.RGBA{R: 0x23, G: 0x78, B: 0x41, A: 0x80},
	TransDarkBlue:               color.RGBA{R: 0x0, G: 0x20, B: 0xa0, A: 0x80},
	TransMediumBlue:             color.RGBA{R: 0x55, G: 0x9a, B: 0xb7, A: 0x80},
	TransLightBlue:              color.RGBA{R: 0xae, G: 0xe9, B: 0xef, A: 0x80},
	TransVeryLightBlue:          color.RGBA{R: 0xc1, G: 0xdf, B: 0xf0, A: 0x80},
	TransBrightReddishLilac:     color.RGBA{R: 0x96, G: 0x70, B: 0x9f, A: 0x80},
	TransPurple:                 color.RGBA{R: 0xa5, G: 0xa5, B: 0xcb, A: 0x80},
	TransDarkPink:               color.RGBA{R: 0xdf, G: 0x66, B: 0x95, A: 0x80},
	TransPink:                   color.RGBA{R: 0xfc, G: 0x97, B: 0xac, A: 0x80},
	TransLightGreen:             color.RGBA{R: 0x7d, G: 0xc2, B: 0x91, A: 0x80},
	TransFireYellow:             color.RGBA{R: 0xfb, G: 0xe8, B: 0x90, A: 0x80},
	TransLightBlueViolet:        color.RGBA{R: 0x6b, G: 0xab, B: 0xe4, A: 0x80},
	TransBrightLightOrange:      color.RGBA{R: 0xfc, G: 0xb7, B: 0x6d, A: 0x80},
	TransReddishLilac:           color.RGBA{R: 0xc2, G: 0x81, B: 0xa5, A: 0x80},
	ChromeGold:                  color.RGBA{R: 0xbb, G: 0xa5, B: 0x3d, A: 0xff},
	ChromeSilver:                color.RGBA{R: 0xe0, G: 0xe0, B: 0xe0, A: 0xff},
	ChromeAntiqueBrass:          color.RGBA{R: 0x64, G: 0x5a, B: 0x4c, A: 0xff},
	ChromeBlack:                 color.RGBA{R: 0x1b, G: 0x2a, B: 0x34, A: 0xff},
	ChromeBlue:                  color.RGBA{R: 0x6c, G: 0x96, B: 0xbf, A: 0xff},
	ChromeGreen:                 color.RGBA{R: 0x3c, G: 0xb3, B: 0x71, A: 0xff},
	ChromePink:                  color.RGBA{R: 0xaa, G: 0x4d, B: 0x8e, A: 0xff},
	PearlWhite:                  color.RGBA{R: 0xf2, G: 0xf3, B: 0xf2, A: 0xff},
	PearlVeryLightGrey:          color.RGBA{R: 0xbb, G: 0xbd, B: 0xbc, A: 0xff},
	PearlLightGrey:              color.RGBA{R: 0x9c, G: 0xa3, B: 0xa8, A: 0xff},
	FlatSilver:                  color.RGBA{R: 0x89, G: 0x87, B: 0x88, A: 0xff},
	PearlDarkGrey:               color.RGBA{R: 0x57, G: 0x58, B: 0x57, A: 0xff},
	MetalBlue:                   color.RGBA{R: 0x56, G: 0x77, B: 0xba, A: 0xff},
	PearlLightGold:              color.RGBA{R: 0xdc, G: 0xbe, B: 0x61, A: 0xff},
	PearlGold:                   color.RGBA{R: 0xcc, G: 0x9c, B: 0x2b, A: 0xff},
	FlatDarkGold:                color.RGBA{R: 0xb4, G: 0x88, B: 0x3e, A: 0xff},
	Copper:                      color.RGBA{R: 0x96, G: 0x4a, B: 0x27, A: 0xff},
	ReddishGold:                 color.RGBA{R: 0xac, G: 0x82, B: 0x47, A: 0xff},
	MetallicSilver:              color.RGBA{R: 0xa5, G: 0xa9, B: 0xb4, A: 0xff},
	MetallicGreen:               color.RGBA{R: 0x89, G: 0x9b, B: 0x5f, A: 0xff},
	MetallicGold:                color.RGBA{R: 0xdb, G: 0xac, B: 0x34, A: 0xff},
	MetallicBlack:               color.RGBA{R: 0x1a, G: 0x28, B: 0x31, A: 0xff},
	MetallicDarkGrey:            color.RGBA{R: 0x6d, G: 0x6e, B: 0x5c, A: 0xff},
	MetallicCopper:              color.RGBA{R: 0xc2, G: 0x7f, B: 0x53, A: 0xff},
	MetallicBrightRed:           color.RGBA{R: 0xd6, G: 0x0, B: 0x26, A: 0xff},
	MetallicDarkGreen:           color.RGBA{R: 0x0, G: 0x8e, B: 0x3c, A: 0xff},
	MilkyWhite:                  color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xf0},
	GlowInDarkOpaque:            color.RGBA{R: 0xe0, G: 0xff, B: 0xb0, A: 0xf0},
	GlowInDarkTrans:             color.RGBA{R: 0xbd, G: 0xc6, B: 0xad, A: 0xf0},
	GlowInDarkWhite:             color.RGBA{R: 0xf5, G: 0xf3, B: 0xd7, A: 0xf0},
	GlitterTransDarkPink:        color.RGBA{R: 0xdf, G: 0x66, B: 0x95, A: 0x80},
	GlitterTransClear:           color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x80},
	GlitterTransPurple:          color.RGBA{R: 0x64, G: 0x0, B: 0x61, A: 0x80},
	GlitterTransLightBlue:       color.RGBA{R: 0xae, G: 0xe9, B: 0xef, A: 0x80},
	GlitterTransNeonGreen:       color.RGBA{R: 0xc0, G: 0xff, B: 0x0, A: 0x80},
	SpeckleBlackSilver:          color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xff},
	SpeckleBlackGold:            color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xff},
	SpeckleBlackCopper:          color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xff},
	SpeckleDarkBluishGreySilver: color.RGBA{R: 0x63, G: 0x5f, B: 0x61, A: 0xff},
	RubberYellow:                color.RGBA{R: 0xf5, G: 0xcd, B: 0x2f, A: 0xff},
	RubberTransYellow:           color.RGBA{R: 0xca, G: 0xb0, B: 0x0, A: 0x80},
	RubberTransClear:            color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x80},
	RubberBlack:                 color.RGBA{R: 0x21, G: 0x21, B: 0x21, A: 0xff},
	RubberBlue:                  color.RGBA{R: 0x0, G: 0x33, B: 0xb2, A: 0xff},
	RubberRed:                   color.RGBA{R: 0xc4, G: 0x0, B: 0x26, A: 0xff},
	RubberOrange:                color.RGBA{R: 0xd0, G: 0x66, B: 0x10, A: 0xff},
	RubberLightGrey:             color.RGBA{R: 0xc1, G: 0xc2, B: 0xc1, A: 0xff},
	RubberDarkBlue:              color.RGBA{R: 0x0, G: 0x1d, B: 0x68, A: 0xff},
	RubberPurple:                color.RGBA{R: 0x81, G: 0x0, B: 0x7b, A: 0xff},
	RubberLime:                  color.RGBA{R: 0xd7, G: 0xf0, B: 0x0, A: 0xff},
	RubberLightBluishGrey:       color.RGBA{R: 0xa3, G: 0xa2, B: 0xa4, A: 0xff},
	RubberFlatSilver:            color.RGBA{R: 0x89, G: 0x87, B: 0x88, A: 0xff},
	RubberWhite:                 color.RGBA{R: 0xfa, G: 0xfa, B: 0xfa, A: 0xff},
	MainColour:                  color.RGBA{R: 0xff, G: 0xff, B: 0x80, A: 0xff},
	EdgeColour:                  color.RGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0xff},
	TransBlackIRLens:            color.RGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xd2},
	Magnet:                      color.RGBA{R: 0x65, G: 0x67, B: 0x61, A: 0xff},
	ElectricContactAlloy:        color.RGBA{R: 0xd0, G: 0xd0, B: 0xd0, A: 0xff},
	ElectricContactCopper:       color.RGBA{R: 0xae, G: 0x7a, B: 0x59, A: 0xff},
}
