package geo

// Comuna representa una de las 346 comunas oficiales de Chile según SUBDERE e INE.
type Comuna struct {
	Cut           string `json:"cut"`            // Código Único Territorial oficial del INE de 5 dígitos (ej: "13120", "16101")
	Name          string `json:"name"`           // Nombre oficial de la comuna
	ProvinciaCode string `json:"provincia_code"` // Código oficial de la provincia de 3 dígitos (ej: "131", "161")
	RegionCode    string `json:"region_code"`    // Código oficial de la región de 2 dígitos (ej: "13", "16")
}

// comunas contiene el listado maestro inmutable de las 346 comunas oficiales de Chile.
var comunas = []Comuna{
	{
		Cut:           "01101",
		Name:          "Iquique",
		ProvinciaCode: "011",
		RegionCode:    "01",
	},
	{
		Cut:           "01107",
		Name:          "Alto Hospicio",
		ProvinciaCode: "011",
		RegionCode:    "01",
	},
	{
		Cut:           "01401",
		Name:          "Pozo Almonte",
		ProvinciaCode: "014",
		RegionCode:    "01",
	},
	{
		Cut:           "01402",
		Name:          "Camiña",
		ProvinciaCode: "014",
		RegionCode:    "01",
	},
	{
		Cut:           "01403",
		Name:          "Colchane",
		ProvinciaCode: "014",
		RegionCode:    "01",
	},
	{
		Cut:           "01404",
		Name:          "Huara",
		ProvinciaCode: "014",
		RegionCode:    "01",
	},
	{
		Cut:           "01405",
		Name:          "Pica",
		ProvinciaCode: "014",
		RegionCode:    "01",
	},
	{
		Cut:           "02101",
		Name:          "Antofagasta",
		ProvinciaCode: "021",
		RegionCode:    "02",
	},
	{
		Cut:           "02102",
		Name:          "Mejillones",
		ProvinciaCode: "021",
		RegionCode:    "02",
	},
	{
		Cut:           "02103",
		Name:          "Sierra Gorda",
		ProvinciaCode: "021",
		RegionCode:    "02",
	},
	{
		Cut:           "02104",
		Name:          "Taltal",
		ProvinciaCode: "021",
		RegionCode:    "02",
	},
	{
		Cut:           "02201",
		Name:          "Calama",
		ProvinciaCode: "022",
		RegionCode:    "02",
	},
	{
		Cut:           "02202",
		Name:          "Ollagüe",
		ProvinciaCode: "022",
		RegionCode:    "02",
	},
	{
		Cut:           "02203",
		Name:          "San Pedro de Atacama",
		ProvinciaCode: "022",
		RegionCode:    "02",
	},
	{
		Cut:           "02301",
		Name:          "Tocopilla",
		ProvinciaCode: "023",
		RegionCode:    "02",
	},
	{
		Cut:           "02302",
		Name:          "María Elena",
		ProvinciaCode: "023",
		RegionCode:    "02",
	},
	{
		Cut:           "03101",
		Name:          "Copiapó",
		ProvinciaCode: "031",
		RegionCode:    "03",
	},
	{
		Cut:           "03102",
		Name:          "Caldera",
		ProvinciaCode: "031",
		RegionCode:    "03",
	},
	{
		Cut:           "03103",
		Name:          "Tierra Amarilla",
		ProvinciaCode: "031",
		RegionCode:    "03",
	},
	{
		Cut:           "03201",
		Name:          "Chañaral",
		ProvinciaCode: "032",
		RegionCode:    "03",
	},
	{
		Cut:           "03202",
		Name:          "Diego de Almagro",
		ProvinciaCode: "032",
		RegionCode:    "03",
	},
	{
		Cut:           "03301",
		Name:          "Vallenar",
		ProvinciaCode: "033",
		RegionCode:    "03",
	},
	{
		Cut:           "03302",
		Name:          "Alto del Carmen",
		ProvinciaCode: "033",
		RegionCode:    "03",
	},
	{
		Cut:           "03303",
		Name:          "Freirina",
		ProvinciaCode: "033",
		RegionCode:    "03",
	},
	{
		Cut:           "03304",
		Name:          "Huasco",
		ProvinciaCode: "033",
		RegionCode:    "03",
	},
	{
		Cut:           "04101",
		Name:          "La Serena",
		ProvinciaCode: "041",
		RegionCode:    "04",
	},
	{
		Cut:           "04102",
		Name:          "Coquimbo",
		ProvinciaCode: "041",
		RegionCode:    "04",
	},
	{
		Cut:           "04103",
		Name:          "Andacollo",
		ProvinciaCode: "041",
		RegionCode:    "04",
	},
	{
		Cut:           "04104",
		Name:          "La Higuera",
		ProvinciaCode: "041",
		RegionCode:    "04",
	},
	{
		Cut:           "04105",
		Name:          "Paiguano",
		ProvinciaCode: "041",
		RegionCode:    "04",
	},
	{
		Cut:           "04106",
		Name:          "Vicuña",
		ProvinciaCode: "041",
		RegionCode:    "04",
	},
	{
		Cut:           "04201",
		Name:          "Illapel",
		ProvinciaCode: "042",
		RegionCode:    "04",
	},
	{
		Cut:           "04202",
		Name:          "Canela",
		ProvinciaCode: "042",
		RegionCode:    "04",
	},
	{
		Cut:           "04203",
		Name:          "Los Vilos",
		ProvinciaCode: "042",
		RegionCode:    "04",
	},
	{
		Cut:           "04204",
		Name:          "Salamanca",
		ProvinciaCode: "042",
		RegionCode:    "04",
	},
	{
		Cut:           "04301",
		Name:          "Ovalle",
		ProvinciaCode: "043",
		RegionCode:    "04",
	},
	{
		Cut:           "04302",
		Name:          "Combarbalá",
		ProvinciaCode: "043",
		RegionCode:    "04",
	},
	{
		Cut:           "04303",
		Name:          "Monte Patria",
		ProvinciaCode: "043",
		RegionCode:    "04",
	},
	{
		Cut:           "04304",
		Name:          "Punitaqui",
		ProvinciaCode: "043",
		RegionCode:    "04",
	},
	{
		Cut:           "04305",
		Name:          "Río Hurtado",
		ProvinciaCode: "043",
		RegionCode:    "04",
	},
	{
		Cut:           "05101",
		Name:          "Valparaíso",
		ProvinciaCode: "051",
		RegionCode:    "05",
	},
	{
		Cut:           "05102",
		Name:          "Casablanca",
		ProvinciaCode: "051",
		RegionCode:    "05",
	},
	{
		Cut:           "05103",
		Name:          "Concón",
		ProvinciaCode: "051",
		RegionCode:    "05",
	},
	{
		Cut:           "05104",
		Name:          "Juan Fernández",
		ProvinciaCode: "051",
		RegionCode:    "05",
	},
	{
		Cut:           "05105",
		Name:          "Puchuncaví",
		ProvinciaCode: "051",
		RegionCode:    "05",
	},
	{
		Cut:           "05107",
		Name:          "Quintero",
		ProvinciaCode: "051",
		RegionCode:    "05",
	},
	{
		Cut:           "05109",
		Name:          "Viña del Mar",
		ProvinciaCode: "051",
		RegionCode:    "05",
	},
	{
		Cut:           "05201",
		Name:          "Isla de Pascua",
		ProvinciaCode: "052",
		RegionCode:    "05",
	},
	{
		Cut:           "05301",
		Name:          "Los Andes",
		ProvinciaCode: "053",
		RegionCode:    "05",
	},
	{
		Cut:           "05302",
		Name:          "Calle Larga",
		ProvinciaCode: "053",
		RegionCode:    "05",
	},
	{
		Cut:           "05303",
		Name:          "Rinconada",
		ProvinciaCode: "053",
		RegionCode:    "05",
	},
	{
		Cut:           "05304",
		Name:          "San Esteban",
		ProvinciaCode: "053",
		RegionCode:    "05",
	},
	{
		Cut:           "05401",
		Name:          "La Ligua",
		ProvinciaCode: "054",
		RegionCode:    "05",
	},
	{
		Cut:           "05402",
		Name:          "Cabildo",
		ProvinciaCode: "054",
		RegionCode:    "05",
	},
	{
		Cut:           "05403",
		Name:          "Papudo",
		ProvinciaCode: "054",
		RegionCode:    "05",
	},
	{
		Cut:           "05404",
		Name:          "Petorca",
		ProvinciaCode: "054",
		RegionCode:    "05",
	},
	{
		Cut:           "05405",
		Name:          "Zapallar",
		ProvinciaCode: "054",
		RegionCode:    "05",
	},
	{
		Cut:           "05501",
		Name:          "Quillota",
		ProvinciaCode: "055",
		RegionCode:    "05",
	},
	{
		Cut:           "05502",
		Name:          "Calera",
		ProvinciaCode: "055",
		RegionCode:    "05",
	},
	{
		Cut:           "05503",
		Name:          "Hijuelas",
		ProvinciaCode: "055",
		RegionCode:    "05",
	},
	{
		Cut:           "05504",
		Name:          "La Cruz",
		ProvinciaCode: "055",
		RegionCode:    "05",
	},
	{
		Cut:           "05506",
		Name:          "Nogales",
		ProvinciaCode: "055",
		RegionCode:    "05",
	},
	{
		Cut:           "05601",
		Name:          "San Antonio",
		ProvinciaCode: "056",
		RegionCode:    "05",
	},
	{
		Cut:           "05602",
		Name:          "Algarrobo",
		ProvinciaCode: "056",
		RegionCode:    "05",
	},
	{
		Cut:           "05603",
		Name:          "Cartagena",
		ProvinciaCode: "056",
		RegionCode:    "05",
	},
	{
		Cut:           "05604",
		Name:          "El Quisco",
		ProvinciaCode: "056",
		RegionCode:    "05",
	},
	{
		Cut:           "05605",
		Name:          "El Tabo",
		ProvinciaCode: "056",
		RegionCode:    "05",
	},
	{
		Cut:           "05606",
		Name:          "Santo Domingo",
		ProvinciaCode: "056",
		RegionCode:    "05",
	},
	{
		Cut:           "05701",
		Name:          "San Felipe",
		ProvinciaCode: "057",
		RegionCode:    "05",
	},
	{
		Cut:           "05702",
		Name:          "Catemu",
		ProvinciaCode: "057",
		RegionCode:    "05",
	},
	{
		Cut:           "05703",
		Name:          "Llaillay",
		ProvinciaCode: "057",
		RegionCode:    "05",
	},
	{
		Cut:           "05704",
		Name:          "Panquehue",
		ProvinciaCode: "057",
		RegionCode:    "05",
	},
	{
		Cut:           "05705",
		Name:          "Putaendo",
		ProvinciaCode: "057",
		RegionCode:    "05",
	},
	{
		Cut:           "05706",
		Name:          "Santa María",
		ProvinciaCode: "057",
		RegionCode:    "05",
	},
	{
		Cut:           "05801",
		Name:          "Quilpué",
		ProvinciaCode: "058",
		RegionCode:    "05",
	},
	{
		Cut:           "05802",
		Name:          "Limache",
		ProvinciaCode: "058",
		RegionCode:    "05",
	},
	{
		Cut:           "05803",
		Name:          "Olmué",
		ProvinciaCode: "058",
		RegionCode:    "05",
	},
	{
		Cut:           "05804",
		Name:          "Villa Alemana",
		ProvinciaCode: "058",
		RegionCode:    "05",
	},
	{
		Cut:           "06101",
		Name:          "Rancagua",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06102",
		Name:          "Codegua",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06103",
		Name:          "Coinco",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06104",
		Name:          "Coltauco",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06105",
		Name:          "Doñihue",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06106",
		Name:          "Graneros",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06107",
		Name:          "Las Cabras",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06108",
		Name:          "Machalí",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06109",
		Name:          "Malloa",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06110",
		Name:          "Mostazal",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06111",
		Name:          "Olivar",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06112",
		Name:          "Peumo",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06113",
		Name:          "Pichidegua",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06114",
		Name:          "Quinta de Tilcoco",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06115",
		Name:          "Rengo",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06116",
		Name:          "Requínoa",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06117",
		Name:          "San Vicente",
		ProvinciaCode: "061",
		RegionCode:    "06",
	},
	{
		Cut:           "06201",
		Name:          "Pichilemu",
		ProvinciaCode: "062",
		RegionCode:    "06",
	},
	{
		Cut:           "06202",
		Name:          "La Estrella",
		ProvinciaCode: "062",
		RegionCode:    "06",
	},
	{
		Cut:           "06203",
		Name:          "Litueche",
		ProvinciaCode: "062",
		RegionCode:    "06",
	},
	{
		Cut:           "06204",
		Name:          "Marchihue",
		ProvinciaCode: "062",
		RegionCode:    "06",
	},
	{
		Cut:           "06205",
		Name:          "Navidad",
		ProvinciaCode: "062",
		RegionCode:    "06",
	},
	{
		Cut:           "06206",
		Name:          "Paredones",
		ProvinciaCode: "062",
		RegionCode:    "06",
	},
	{
		Cut:           "06301",
		Name:          "San Fernando",
		ProvinciaCode: "063",
		RegionCode:    "06",
	},
	{
		Cut:           "06302",
		Name:          "Chépica",
		ProvinciaCode: "063",
		RegionCode:    "06",
	},
	{
		Cut:           "06303",
		Name:          "Chimbarongo",
		ProvinciaCode: "063",
		RegionCode:    "06",
	},
	{
		Cut:           "06304",
		Name:          "Lolol",
		ProvinciaCode: "063",
		RegionCode:    "06",
	},
	{
		Cut:           "06305",
		Name:          "Nancagua",
		ProvinciaCode: "063",
		RegionCode:    "06",
	},
	{
		Cut:           "06306",
		Name:          "Palmilla",
		ProvinciaCode: "063",
		RegionCode:    "06",
	},
	{
		Cut:           "06307",
		Name:          "Peralillo",
		ProvinciaCode: "063",
		RegionCode:    "06",
	},
	{
		Cut:           "06308",
		Name:          "Placilla",
		ProvinciaCode: "063",
		RegionCode:    "06",
	},
	{
		Cut:           "06309",
		Name:          "Pumanque",
		ProvinciaCode: "063",
		RegionCode:    "06",
	},
	{
		Cut:           "06310",
		Name:          "Santa Cruz",
		ProvinciaCode: "063",
		RegionCode:    "06",
	},
	{
		Cut:           "07101",
		Name:          "Talca",
		ProvinciaCode: "071",
		RegionCode:    "07",
	},
	{
		Cut:           "07102",
		Name:          "Constitución",
		ProvinciaCode: "071",
		RegionCode:    "07",
	},
	{
		Cut:           "07103",
		Name:          "Curepto",
		ProvinciaCode: "071",
		RegionCode:    "07",
	},
	{
		Cut:           "07104",
		Name:          "Empedrado",
		ProvinciaCode: "071",
		RegionCode:    "07",
	},
	{
		Cut:           "07105",
		Name:          "Maule",
		ProvinciaCode: "071",
		RegionCode:    "07",
	},
	{
		Cut:           "07106",
		Name:          "Pelarco",
		ProvinciaCode: "071",
		RegionCode:    "07",
	},
	{
		Cut:           "07107",
		Name:          "Pencahue",
		ProvinciaCode: "071",
		RegionCode:    "07",
	},
	{
		Cut:           "07108",
		Name:          "Río Claro",
		ProvinciaCode: "071",
		RegionCode:    "07",
	},
	{
		Cut:           "07109",
		Name:          "San Clemente",
		ProvinciaCode: "071",
		RegionCode:    "07",
	},
	{
		Cut:           "07110",
		Name:          "San Rafael",
		ProvinciaCode: "071",
		RegionCode:    "07",
	},
	{
		Cut:           "07201",
		Name:          "Cauquenes",
		ProvinciaCode: "072",
		RegionCode:    "07",
	},
	{
		Cut:           "07202",
		Name:          "Chanco",
		ProvinciaCode: "072",
		RegionCode:    "07",
	},
	{
		Cut:           "07203",
		Name:          "Pelluhue",
		ProvinciaCode: "072",
		RegionCode:    "07",
	},
	{
		Cut:           "07301",
		Name:          "Curicó",
		ProvinciaCode: "073",
		RegionCode:    "07",
	},
	{
		Cut:           "07302",
		Name:          "Hualañé",
		ProvinciaCode: "073",
		RegionCode:    "07",
	},
	{
		Cut:           "07303",
		Name:          "Licantén",
		ProvinciaCode: "073",
		RegionCode:    "07",
	},
	{
		Cut:           "07304",
		Name:          "Molina",
		ProvinciaCode: "073",
		RegionCode:    "07",
	},
	{
		Cut:           "07305",
		Name:          "Rauco",
		ProvinciaCode: "073",
		RegionCode:    "07",
	},
	{
		Cut:           "07306",
		Name:          "Romeral",
		ProvinciaCode: "073",
		RegionCode:    "07",
	},
	{
		Cut:           "07307",
		Name:          "Sagrada Familia",
		ProvinciaCode: "073",
		RegionCode:    "07",
	},
	{
		Cut:           "07308",
		Name:          "Teno",
		ProvinciaCode: "073",
		RegionCode:    "07",
	},
	{
		Cut:           "07309",
		Name:          "Vichuquén",
		ProvinciaCode: "073",
		RegionCode:    "07",
	},
	{
		Cut:           "07401",
		Name:          "Linares",
		ProvinciaCode: "074",
		RegionCode:    "07",
	},
	{
		Cut:           "07402",
		Name:          "Colbún",
		ProvinciaCode: "074",
		RegionCode:    "07",
	},
	{
		Cut:           "07403",
		Name:          "Longaví",
		ProvinciaCode: "074",
		RegionCode:    "07",
	},
	{
		Cut:           "07404",
		Name:          "Parral",
		ProvinciaCode: "074",
		RegionCode:    "07",
	},
	{
		Cut:           "07405",
		Name:          "Retiro",
		ProvinciaCode: "074",
		RegionCode:    "07",
	},
	{
		Cut:           "07406",
		Name:          "San Javier",
		ProvinciaCode: "074",
		RegionCode:    "07",
	},
	{
		Cut:           "07407",
		Name:          "Villa Alegre",
		ProvinciaCode: "074",
		RegionCode:    "07",
	},
	{
		Cut:           "07408",
		Name:          "Yerbas Buenas",
		ProvinciaCode: "074",
		RegionCode:    "07",
	},
	{
		Cut:           "08101",
		Name:          "Concepción",
		ProvinciaCode: "081",
		RegionCode:    "08",
	},
	{
		Cut:           "08102",
		Name:          "Coronel",
		ProvinciaCode: "081",
		RegionCode:    "08",
	},
	{
		Cut:           "08103",
		Name:          "Chiguayante",
		ProvinciaCode: "081",
		RegionCode:    "08",
	},
	{
		Cut:           "08104",
		Name:          "Florida",
		ProvinciaCode: "081",
		RegionCode:    "08",
	},
	{
		Cut:           "08105",
		Name:          "Hualqui",
		ProvinciaCode: "081",
		RegionCode:    "08",
	},
	{
		Cut:           "08106",
		Name:          "Lota",
		ProvinciaCode: "081",
		RegionCode:    "08",
	},
	{
		Cut:           "08107",
		Name:          "Penco",
		ProvinciaCode: "081",
		RegionCode:    "08",
	},
	{
		Cut:           "08108",
		Name:          "San Pedro de la Paz",
		ProvinciaCode: "081",
		RegionCode:    "08",
	},
	{
		Cut:           "08109",
		Name:          "Santa Juana",
		ProvinciaCode: "081",
		RegionCode:    "08",
	},
	{
		Cut:           "08110",
		Name:          "Talcahuano",
		ProvinciaCode: "081",
		RegionCode:    "08",
	},
	{
		Cut:           "08111",
		Name:          "Tomé",
		ProvinciaCode: "081",
		RegionCode:    "08",
	},
	{
		Cut:           "08112",
		Name:          "Hualpén",
		ProvinciaCode: "081",
		RegionCode:    "08",
	},
	{
		Cut:           "08201",
		Name:          "Lebu",
		ProvinciaCode: "082",
		RegionCode:    "08",
	},
	{
		Cut:           "08202",
		Name:          "Arauco",
		ProvinciaCode: "082",
		RegionCode:    "08",
	},
	{
		Cut:           "08203",
		Name:          "Cañete",
		ProvinciaCode: "082",
		RegionCode:    "08",
	},
	{
		Cut:           "08204",
		Name:          "Contulmo",
		ProvinciaCode: "082",
		RegionCode:    "08",
	},
	{
		Cut:           "08205",
		Name:          "Curanilahue",
		ProvinciaCode: "082",
		RegionCode:    "08",
	},
	{
		Cut:           "08206",
		Name:          "Los Alamos",
		ProvinciaCode: "082",
		RegionCode:    "08",
	},
	{
		Cut:           "08207",
		Name:          "Tirúa",
		ProvinciaCode: "082",
		RegionCode:    "08",
	},
	{
		Cut:           "08301",
		Name:          "Los Angeles",
		ProvinciaCode: "083",
		RegionCode:    "08",
	},
	{
		Cut:           "08302",
		Name:          "Antuco",
		ProvinciaCode: "083",
		RegionCode:    "08",
	},
	{
		Cut:           "08303",
		Name:          "Cabrero",
		ProvinciaCode: "083",
		RegionCode:    "08",
	},
	{
		Cut:           "08304",
		Name:          "Laja",
		ProvinciaCode: "083",
		RegionCode:    "08",
	},
	{
		Cut:           "08305",
		Name:          "Mulchén",
		ProvinciaCode: "083",
		RegionCode:    "08",
	},
	{
		Cut:           "08306",
		Name:          "Nacimiento",
		ProvinciaCode: "083",
		RegionCode:    "08",
	},
	{
		Cut:           "08307",
		Name:          "Negrete",
		ProvinciaCode: "083",
		RegionCode:    "08",
	},
	{
		Cut:           "08308",
		Name:          "Quilaco",
		ProvinciaCode: "083",
		RegionCode:    "08",
	},
	{
		Cut:           "08309",
		Name:          "Quilleco",
		ProvinciaCode: "083",
		RegionCode:    "08",
	},
	{
		Cut:           "08310",
		Name:          "San Rosendo",
		ProvinciaCode: "083",
		RegionCode:    "08",
	},
	{
		Cut:           "08311",
		Name:          "Santa Bárbara",
		ProvinciaCode: "083",
		RegionCode:    "08",
	},
	{
		Cut:           "08312",
		Name:          "Tucapel",
		ProvinciaCode: "083",
		RegionCode:    "08",
	},
	{
		Cut:           "08313",
		Name:          "Yumbel",
		ProvinciaCode: "083",
		RegionCode:    "08",
	},
	{
		Cut:           "08314",
		Name:          "Alto Biobío",
		ProvinciaCode: "083",
		RegionCode:    "08",
	},
	{
		Cut:           "09101",
		Name:          "Temuco",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09102",
		Name:          "Carahue",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09103",
		Name:          "Cunco",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09104",
		Name:          "Curarrehue",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09105",
		Name:          "Freire",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09106",
		Name:          "Galvarino",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09107",
		Name:          "Gorbea",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09108",
		Name:          "Lautaro",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09109",
		Name:          "Loncoche",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09110",
		Name:          "Melipeuco",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09111",
		Name:          "Nueva Imperial",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09112",
		Name:          "Padre Las Casas",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09113",
		Name:          "Perquenco",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09114",
		Name:          "Pitrufquén",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09115",
		Name:          "Pucón",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09116",
		Name:          "Saavedra",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09117",
		Name:          "Teodoro Schmidt",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09118",
		Name:          "Toltén",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09119",
		Name:          "Vilcún",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09120",
		Name:          "Villarrica",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09121",
		Name:          "Cholchol",
		ProvinciaCode: "091",
		RegionCode:    "09",
	},
	{
		Cut:           "09201",
		Name:          "Angol",
		ProvinciaCode: "092",
		RegionCode:    "09",
	},
	{
		Cut:           "09202",
		Name:          "Collipulli",
		ProvinciaCode: "092",
		RegionCode:    "09",
	},
	{
		Cut:           "09203",
		Name:          "Curacautín",
		ProvinciaCode: "092",
		RegionCode:    "09",
	},
	{
		Cut:           "09204",
		Name:          "Ercilla",
		ProvinciaCode: "092",
		RegionCode:    "09",
	},
	{
		Cut:           "09205",
		Name:          "Lonquimay",
		ProvinciaCode: "092",
		RegionCode:    "09",
	},
	{
		Cut:           "09206",
		Name:          "Los Sauces",
		ProvinciaCode: "092",
		RegionCode:    "09",
	},
	{
		Cut:           "09207",
		Name:          "Lumaco",
		ProvinciaCode: "092",
		RegionCode:    "09",
	},
	{
		Cut:           "09208",
		Name:          "Purén",
		ProvinciaCode: "092",
		RegionCode:    "09",
	},
	{
		Cut:           "09209",
		Name:          "Renaico",
		ProvinciaCode: "092",
		RegionCode:    "09",
	},
	{
		Cut:           "09210",
		Name:          "Traiguén",
		ProvinciaCode: "092",
		RegionCode:    "09",
	},
	{
		Cut:           "09211",
		Name:          "Victoria",
		ProvinciaCode: "092",
		RegionCode:    "09",
	},
	{
		Cut:           "10101",
		Name:          "Puerto Montt",
		ProvinciaCode: "101",
		RegionCode:    "10",
	},
	{
		Cut:           "10102",
		Name:          "Calbuco",
		ProvinciaCode: "101",
		RegionCode:    "10",
	},
	{
		Cut:           "10103",
		Name:          "Cochamó",
		ProvinciaCode: "101",
		RegionCode:    "10",
	},
	{
		Cut:           "10104",
		Name:          "Fresia",
		ProvinciaCode: "101",
		RegionCode:    "10",
	},
	{
		Cut:           "10105",
		Name:          "Frutillar",
		ProvinciaCode: "101",
		RegionCode:    "10",
	},
	{
		Cut:           "10106",
		Name:          "Los Muermos",
		ProvinciaCode: "101",
		RegionCode:    "10",
	},
	{
		Cut:           "10107",
		Name:          "Llanquihue",
		ProvinciaCode: "101",
		RegionCode:    "10",
	},
	{
		Cut:           "10108",
		Name:          "Maullín",
		ProvinciaCode: "101",
		RegionCode:    "10",
	},
	{
		Cut:           "10109",
		Name:          "Puerto Varas",
		ProvinciaCode: "101",
		RegionCode:    "10",
	},
	{
		Cut:           "10201",
		Name:          "Castro",
		ProvinciaCode: "102",
		RegionCode:    "10",
	},
	{
		Cut:           "10202",
		Name:          "Ancud",
		ProvinciaCode: "102",
		RegionCode:    "10",
	},
	{
		Cut:           "10203",
		Name:          "Chonchi",
		ProvinciaCode: "102",
		RegionCode:    "10",
	},
	{
		Cut:           "10204",
		Name:          "Curaco de Vélez",
		ProvinciaCode: "102",
		RegionCode:    "10",
	},
	{
		Cut:           "10205",
		Name:          "Dalcahue",
		ProvinciaCode: "102",
		RegionCode:    "10",
	},
	{
		Cut:           "10206",
		Name:          "Puqueldón",
		ProvinciaCode: "102",
		RegionCode:    "10",
	},
	{
		Cut:           "10207",
		Name:          "Queilén",
		ProvinciaCode: "102",
		RegionCode:    "10",
	},
	{
		Cut:           "10208",
		Name:          "Quellón",
		ProvinciaCode: "102",
		RegionCode:    "10",
	},
	{
		Cut:           "10209",
		Name:          "Quemchi",
		ProvinciaCode: "102",
		RegionCode:    "10",
	},
	{
		Cut:           "10210",
		Name:          "Quinchao",
		ProvinciaCode: "102",
		RegionCode:    "10",
	},
	{
		Cut:           "10301",
		Name:          "Osorno",
		ProvinciaCode: "103",
		RegionCode:    "10",
	},
	{
		Cut:           "10302",
		Name:          "Puerto Octay",
		ProvinciaCode: "103",
		RegionCode:    "10",
	},
	{
		Cut:           "10303",
		Name:          "Purranque",
		ProvinciaCode: "103",
		RegionCode:    "10",
	},
	{
		Cut:           "10304",
		Name:          "Puyehue",
		ProvinciaCode: "103",
		RegionCode:    "10",
	},
	{
		Cut:           "10305",
		Name:          "Río Negro",
		ProvinciaCode: "103",
		RegionCode:    "10",
	},
	{
		Cut:           "10306",
		Name:          "San Juan de la Costa",
		ProvinciaCode: "103",
		RegionCode:    "10",
	},
	{
		Cut:           "10307",
		Name:          "San Pablo",
		ProvinciaCode: "103",
		RegionCode:    "10",
	},
	{
		Cut:           "10401",
		Name:          "Chaitén",
		ProvinciaCode: "104",
		RegionCode:    "10",
	},
	{
		Cut:           "10402",
		Name:          "Futaleufú",
		ProvinciaCode: "104",
		RegionCode:    "10",
	},
	{
		Cut:           "10403",
		Name:          "Hualaihué",
		ProvinciaCode: "104",
		RegionCode:    "10",
	},
	{
		Cut:           "10404",
		Name:          "Palena",
		ProvinciaCode: "104",
		RegionCode:    "10",
	},
	{
		Cut:           "11101",
		Name:          "Coyhaique",
		ProvinciaCode: "111",
		RegionCode:    "11",
	},
	{
		Cut:           "11102",
		Name:          "Lago Verde",
		ProvinciaCode: "111",
		RegionCode:    "11",
	},
	{
		Cut:           "11201",
		Name:          "Aysén",
		ProvinciaCode: "112",
		RegionCode:    "11",
	},
	{
		Cut:           "11202",
		Name:          "Cisnes",
		ProvinciaCode: "112",
		RegionCode:    "11",
	},
	{
		Cut:           "11203",
		Name:          "Guaitecas",
		ProvinciaCode: "112",
		RegionCode:    "11",
	},
	{
		Cut:           "11301",
		Name:          "Cochrane",
		ProvinciaCode: "113",
		RegionCode:    "11",
	},
	{
		Cut:           "11302",
		Name:          "O'Higgins",
		ProvinciaCode: "113",
		RegionCode:    "11",
	},
	{
		Cut:           "11303",
		Name:          "Tortel",
		ProvinciaCode: "113",
		RegionCode:    "11",
	},
	{
		Cut:           "11401",
		Name:          "Chile Chico",
		ProvinciaCode: "114",
		RegionCode:    "11",
	},
	{
		Cut:           "11402",
		Name:          "Río Ibáñez",
		ProvinciaCode: "114",
		RegionCode:    "11",
	},
	{
		Cut:           "12101",
		Name:          "Punta Arenas",
		ProvinciaCode: "121",
		RegionCode:    "12",
	},
	{
		Cut:           "12102",
		Name:          "Laguna Blanca",
		ProvinciaCode: "121",
		RegionCode:    "12",
	},
	{
		Cut:           "12103",
		Name:          "Río Verde",
		ProvinciaCode: "121",
		RegionCode:    "12",
	},
	{
		Cut:           "12104",
		Name:          "San Gregorio",
		ProvinciaCode: "121",
		RegionCode:    "12",
	},
	{
		Cut:           "12201",
		Name:          "Cabo de Hornos",
		ProvinciaCode: "122",
		RegionCode:    "12",
	},
	{
		Cut:           "12202",
		Name:          "Antártica",
		ProvinciaCode: "122",
		RegionCode:    "12",
	},
	{
		Cut:           "12301",
		Name:          "Porvenir",
		ProvinciaCode: "123",
		RegionCode:    "12",
	},
	{
		Cut:           "12302",
		Name:          "Primavera",
		ProvinciaCode: "123",
		RegionCode:    "12",
	},
	{
		Cut:           "12303",
		Name:          "Timaukel",
		ProvinciaCode: "123",
		RegionCode:    "12",
	},
	{
		Cut:           "12401",
		Name:          "Natales",
		ProvinciaCode: "124",
		RegionCode:    "12",
	},
	{
		Cut:           "12402",
		Name:          "Torres del Paine",
		ProvinciaCode: "124",
		RegionCode:    "12",
	},
	{
		Cut:           "13101",
		Name:          "Santiago",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13102",
		Name:          "Cerrillos",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13103",
		Name:          "Cerro Navia",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13104",
		Name:          "Conchalí",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13105",
		Name:          "El Bosque",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13106",
		Name:          "Estación Central",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13107",
		Name:          "Huechuraba",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13108",
		Name:          "Independencia",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13109",
		Name:          "La Cisterna",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13110",
		Name:          "La Florida",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13111",
		Name:          "La Granja",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13112",
		Name:          "La Pintana",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13113",
		Name:          "La Reina",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13114",
		Name:          "Las Condes",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13115",
		Name:          "Lo Barnechea",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13116",
		Name:          "Lo Espejo",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13117",
		Name:          "Lo Prado",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13118",
		Name:          "Macul",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13119",
		Name:          "Maipú",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13120",
		Name:          "Ñuñoa",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13121",
		Name:          "Pedro Aguirre Cerda",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13122",
		Name:          "Peñalolén",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13123",
		Name:          "Providencia",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13124",
		Name:          "Pudahuel",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13125",
		Name:          "Quilicura",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13126",
		Name:          "Quinta Normal",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13127",
		Name:          "Recoleta",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13128",
		Name:          "Renca",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13129",
		Name:          "San Joaquín",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13130",
		Name:          "San Miguel",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13131",
		Name:          "San Ramón",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13132",
		Name:          "Vitacura",
		ProvinciaCode: "131",
		RegionCode:    "13",
	},
	{
		Cut:           "13201",
		Name:          "Puente Alto",
		ProvinciaCode: "132",
		RegionCode:    "13",
	},
	{
		Cut:           "13202",
		Name:          "Pirque",
		ProvinciaCode: "132",
		RegionCode:    "13",
	},
	{
		Cut:           "13203",
		Name:          "San José de Maipo",
		ProvinciaCode: "132",
		RegionCode:    "13",
	},
	{
		Cut:           "13301",
		Name:          "Colina",
		ProvinciaCode: "133",
		RegionCode:    "13",
	},
	{
		Cut:           "13302",
		Name:          "Lampa",
		ProvinciaCode: "133",
		RegionCode:    "13",
	},
	{
		Cut:           "13303",
		Name:          "Tiltil",
		ProvinciaCode: "133",
		RegionCode:    "13",
	},
	{
		Cut:           "13401",
		Name:          "San Bernardo",
		ProvinciaCode: "134",
		RegionCode:    "13",
	},
	{
		Cut:           "13402",
		Name:          "Buin",
		ProvinciaCode: "134",
		RegionCode:    "13",
	},
	{
		Cut:           "13403",
		Name:          "Calera de Tango",
		ProvinciaCode: "134",
		RegionCode:    "13",
	},
	{
		Cut:           "13404",
		Name:          "Paine",
		ProvinciaCode: "134",
		RegionCode:    "13",
	},
	{
		Cut:           "13501",
		Name:          "Melipilla",
		ProvinciaCode: "135",
		RegionCode:    "13",
	},
	{
		Cut:           "13502",
		Name:          "Alhué",
		ProvinciaCode: "135",
		RegionCode:    "13",
	},
	{
		Cut:           "13503",
		Name:          "Curacaví",
		ProvinciaCode: "135",
		RegionCode:    "13",
	},
	{
		Cut:           "13504",
		Name:          "María Pinto",
		ProvinciaCode: "135",
		RegionCode:    "13",
	},
	{
		Cut:           "13505",
		Name:          "San Pedro",
		ProvinciaCode: "135",
		RegionCode:    "13",
	},
	{
		Cut:           "13601",
		Name:          "Talagante",
		ProvinciaCode: "136",
		RegionCode:    "13",
	},
	{
		Cut:           "13602",
		Name:          "El Monte",
		ProvinciaCode: "136",
		RegionCode:    "13",
	},
	{
		Cut:           "13603",
		Name:          "Isla de Maipo",
		ProvinciaCode: "136",
		RegionCode:    "13",
	},
	{
		Cut:           "13604",
		Name:          "Padre Hurtado",
		ProvinciaCode: "136",
		RegionCode:    "13",
	},
	{
		Cut:           "13605",
		Name:          "Peñaflor",
		ProvinciaCode: "136",
		RegionCode:    "13",
	},
	{
		Cut:           "14101",
		Name:          "Valdivia",
		ProvinciaCode: "141",
		RegionCode:    "14",
	},
	{
		Cut:           "14102",
		Name:          "Corral",
		ProvinciaCode: "141",
		RegionCode:    "14",
	},
	{
		Cut:           "14103",
		Name:          "Lanco",
		ProvinciaCode: "141",
		RegionCode:    "14",
	},
	{
		Cut:           "14104",
		Name:          "Los Lagos",
		ProvinciaCode: "141",
		RegionCode:    "14",
	},
	{
		Cut:           "14105",
		Name:          "Máfil",
		ProvinciaCode: "141",
		RegionCode:    "14",
	},
	{
		Cut:           "14106",
		Name:          "Mariquina",
		ProvinciaCode: "141",
		RegionCode:    "14",
	},
	{
		Cut:           "14107",
		Name:          "Paillaco",
		ProvinciaCode: "141",
		RegionCode:    "14",
	},
	{
		Cut:           "14108",
		Name:          "Panguipulli",
		ProvinciaCode: "141",
		RegionCode:    "14",
	},
	{
		Cut:           "14201",
		Name:          "La Unión",
		ProvinciaCode: "142",
		RegionCode:    "14",
	},
	{
		Cut:           "14202",
		Name:          "Futrono",
		ProvinciaCode: "142",
		RegionCode:    "14",
	},
	{
		Cut:           "14203",
		Name:          "Lago Ranco",
		ProvinciaCode: "142",
		RegionCode:    "14",
	},
	{
		Cut:           "14204",
		Name:          "Río Bueno",
		ProvinciaCode: "142",
		RegionCode:    "14",
	},
	{
		Cut:           "15101",
		Name:          "Arica",
		ProvinciaCode: "151",
		RegionCode:    "15",
	},
	{
		Cut:           "15102",
		Name:          "Camarones",
		ProvinciaCode: "151",
		RegionCode:    "15",
	},
	{
		Cut:           "15201",
		Name:          "Putre",
		ProvinciaCode: "152",
		RegionCode:    "15",
	},
	{
		Cut:           "15202",
		Name:          "General Lagos",
		ProvinciaCode: "152",
		RegionCode:    "15",
	},
	{
		Cut:           "16101",
		Name:          "Chillán",
		ProvinciaCode: "161",
		RegionCode:    "16",
	},
	{
		Cut:           "16102",
		Name:          "Bulnes",
		ProvinciaCode: "161",
		RegionCode:    "16",
	},
	{
		Cut:           "16103",
		Name:          "Chillán Viejo",
		ProvinciaCode: "161",
		RegionCode:    "16",
	},
	{
		Cut:           "16104",
		Name:          "El Carmen",
		ProvinciaCode: "161",
		RegionCode:    "16",
	},
	{
		Cut:           "16105",
		Name:          "Pemuco",
		ProvinciaCode: "161",
		RegionCode:    "16",
	},
	{
		Cut:           "16106",
		Name:          "Pinto",
		ProvinciaCode: "161",
		RegionCode:    "16",
	},
	{
		Cut:           "16107",
		Name:          "Quillón",
		ProvinciaCode: "161",
		RegionCode:    "16",
	},
	{
		Cut:           "16108",
		Name:          "San Ignacio",
		ProvinciaCode: "161",
		RegionCode:    "16",
	},
	{
		Cut:           "16109",
		Name:          "Yungay",
		ProvinciaCode: "161",
		RegionCode:    "16",
	},
	{
		Cut:           "16201",
		Name:          "Quirihue",
		ProvinciaCode: "162",
		RegionCode:    "16",
	},
	{
		Cut:           "16202",
		Name:          "Cobquecura",
		ProvinciaCode: "162",
		RegionCode:    "16",
	},
	{
		Cut:           "16203",
		Name:          "Coelemu",
		ProvinciaCode: "162",
		RegionCode:    "16",
	},
	{
		Cut:           "16204",
		Name:          "Ninhue",
		ProvinciaCode: "162",
		RegionCode:    "16",
	},
	{
		Cut:           "16205",
		Name:          "Portezuelo",
		ProvinciaCode: "162",
		RegionCode:    "16",
	},
	{
		Cut:           "16206",
		Name:          "Ranquil",
		ProvinciaCode: "162",
		RegionCode:    "16",
	},
	{
		Cut:           "16207",
		Name:          "Treguaco",
		ProvinciaCode: "162",
		RegionCode:    "16",
	},
	{
		Cut:           "16301",
		Name:          "San Carlos",
		ProvinciaCode: "163",
		RegionCode:    "16",
	},
	{
		Cut:           "16302",
		Name:          "Coihueco",
		ProvinciaCode: "163",
		RegionCode:    "16",
	},
	{
		Cut:           "16303",
		Name:          "Ñiquén",
		ProvinciaCode: "163",
		RegionCode:    "16",
	},
	{
		Cut:           "16304",
		Name:          "San Fabián",
		ProvinciaCode: "163",
		RegionCode:    "16",
	},
	{
		Cut:           "16305",
		Name:          "San Nicolás",
		ProvinciaCode: "163",
		RegionCode:    "16",
	},
}

// GetComunas retorna una copia de las 346 comunas oficiales de Chile.
func GetComunas() []Comuna {
	result := make([]Comuna, len(comunas))
	copy(result, comunas)
	return result
}

// GetComunasByRegion retorna las comunas pertenecientes a una región específica.
// Acepta código de región (ej: "13", "16"), código ISO (ej: "CL-RM", "CL-NB") o nombre de región.
// Si la región no existe o no tiene comunas, retorna un slice vacío.
func GetComunasByRegion(regionQuery string) []Comuna {
	reg, ok := FindRegion(regionQuery)
	if !ok || reg == nil {
		return []Comuna{}
	}

	var result []Comuna
	for i := range comunas {
		if comunas[i].RegionCode == reg.Code {
			result = append(result, comunas[i])
		}
	}
	return result
}

// GetComunasByProvincia retorna las comunas pertenecientes a una provincia específica.
// Acepta código de provincia (ej: "131", "161") o nombre de provincia (ej: "santiago", "diguillin").
// Si la provincia no existe o no tiene comunas, retorna un slice vacío.
func GetComunasByProvincia(provinciaQuery string) []Comuna {
	prov, ok := FindProvincia(provinciaQuery)
	if !ok || prov == nil {
		return []Comuna{}
	}

	var result []Comuna
	for i := range comunas {
		if comunas[i].ProvinciaCode == prov.Code {
			result = append(result, comunas[i])
		}
	}
	return result
}

// FindComuna busca una comuna por su Código Único Territorial (CUT de 5 dígitos, ej: "13120", "16101")
// o por su nombre oficial (tolerante a mayúsculas, tildes y eñes, ej: "nunoa", "vina del mar", "aysen").
// Retorna un puntero a una copia de Comuna y true si fue encontrada, o nil y false si no existe coincidencia.
func FindComuna(query string) (*Comuna, bool) {
	return searchEntity(
		query,
		comunas,
		func(c Comuna) []string {
			return []string{c.Cut}
		},
		func(c Comuna) string {
			return c.Name
		},
	)
}

// SearchComunas busca y retorna todas las comunas que coincidan con el término de búsqueda,
// ordenadas por relevancia (coincidencias exactas primero, luego prefijos y finalmente subcadenas).
// Ideal para cajas de autocompletado y filtros de búsqueda interactivos en UI (ej: "San" -> San Bernardo, San Miguel...).
func SearchComunas(query string) []Comuna {
	return searchAllEntities(
		query,
		comunas,
		func(c Comuna) []string {
			return []string{c.Cut}
		},
		func(c Comuna) string {
			return c.Name
		},
	)
}

// GetRegionByComuna resuelve la región oficial a la que pertenece una comuna dada.
// Acepta código CUT (ej: "13120", "16101") o nombre de comuna (ej: "nunoa", "chillan").
// Retorna un puntero a Region y true si fue resuelta, o nil y false si la comuna no existe.
func GetRegionByComuna(comunaQuery string) (*Region, bool) {
	comuna, ok := FindComuna(comunaQuery)
	if !ok || comuna == nil {
		return nil, false
	}
	return FindRegion(comuna.RegionCode)
}

// GetProvinciaByComuna resuelve la provincia oficial a la que pertenece una comuna dada.
// Acepta código CUT (ej: "13120", "16101") o nombre de comuna (ej: "nunoa", "chillan").
// Retorna un puntero a Provincia y true si fue resuelta, o nil y false si la comuna no existe.
func GetProvinciaByComuna(comunaQuery string) (*Provincia, bool) {
	comuna, ok := FindComuna(comunaQuery)
	if !ok || comuna == nil {
		return nil, false
	}
	return FindProvincia(comuna.ProvinciaCode)
}
