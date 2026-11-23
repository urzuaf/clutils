import type { Region, Provincia, Comuna } from './types.js';

export const REGIONES: readonly Region[] = [
  {
    "code": "15",
    "isoCode": "CL-AP",
    "name": "Arica y Parinacota",
    "abbreviation": "AyP",
    "capital": "Arica",
    "order": 1
  },
  {
    "code": "01",
    "isoCode": "CL-TA",
    "name": "Tarapacá",
    "abbreviation": "TPCA",
    "capital": "Iquique",
    "order": 2
  },
  {
    "code": "02",
    "isoCode": "CL-AN",
    "name": "Antofagasta",
    "abbreviation": "ANTOF",
    "capital": "Antofagasta",
    "order": 3
  },
  {
    "code": "03",
    "isoCode": "CL-AT",
    "name": "Atacama",
    "abbreviation": "ATCMA",
    "capital": "Copiapó",
    "order": 4
  },
  {
    "code": "04",
    "isoCode": "CL-CO",
    "name": "Coquimbo",
    "abbreviation": "COQ",
    "capital": "La Serena",
    "order": 5
  },
  {
    "code": "05",
    "isoCode": "CL-VS",
    "name": "Valparaíso",
    "abbreviation": "VALPO",
    "capital": "Valparaíso",
    "order": 6
  },
  {
    "code": "13",
    "isoCode": "CL-RM",
    "name": "Metropolitana de Santiago",
    "abbreviation": "RM",
    "capital": "Santiago",
    "order": 7
  },
  {
    "code": "06",
    "isoCode": "CL-LI",
    "name": "Libertador General Bernardo O'Higgins",
    "abbreviation": "LGBO",
    "capital": "Rancagua",
    "order": 8
  },
  {
    "code": "07",
    "isoCode": "CL-ML",
    "name": "Maule",
    "abbreviation": "MAULE",
    "capital": "Talca",
    "order": 9
  },
  {
    "code": "16",
    "isoCode": "CL-NB",
    "name": "Ñuble",
    "abbreviation": "NUBLE",
    "capital": "Chillán",
    "order": 10
  },
  {
    "code": "08",
    "isoCode": "CL-BI",
    "name": "Biobío",
    "abbreviation": "BBIO",
    "capital": "Concepción",
    "order": 11
  },
  {
    "code": "09",
    "isoCode": "CL-AR",
    "name": "La Araucanía",
    "abbreviation": "ARAUC",
    "capital": "Temuco",
    "order": 12
  },
  {
    "code": "14",
    "isoCode": "CL-LR",
    "name": "Los Ríos",
    "abbreviation": "RIOS",
    "capital": "Valdivia",
    "order": 13
  },
  {
    "code": "10",
    "isoCode": "CL-LL",
    "name": "Los Lagos",
    "abbreviation": "LAGOS",
    "capital": "Puerto Montt",
    "order": 14
  },
  {
    "code": "11",
    "isoCode": "CL-AI",
    "name": "Aysén del General Carlos Ibáñez del Campo",
    "abbreviation": "AYSEN",
    "capital": "Coyhaique",
    "order": 15
  },
  {
    "code": "12",
    "isoCode": "CL-MA",
    "name": "Magallanes y de la Antártica Chilena",
    "abbreviation": "MAG",
    "capital": "Punta Arenas",
    "order": 16
  }
] as const;

export const PROVINCIAS: readonly Provincia[] = [
  {
    "code": "011",
    "name": "Iquique",
    "regionCode": "01"
  },
  {
    "code": "014",
    "name": "Tamarugal",
    "regionCode": "01"
  },
  {
    "code": "021",
    "name": "Antofagasta",
    "regionCode": "02"
  },
  {
    "code": "022",
    "name": "El Loa",
    "regionCode": "02"
  },
  {
    "code": "023",
    "name": "Tocopilla",
    "regionCode": "02"
  },
  {
    "code": "031",
    "name": "Copiapó",
    "regionCode": "03"
  },
  {
    "code": "032",
    "name": "Chañaral",
    "regionCode": "03"
  },
  {
    "code": "033",
    "name": "Huasco",
    "regionCode": "03"
  },
  {
    "code": "041",
    "name": "Elqui",
    "regionCode": "04"
  },
  {
    "code": "042",
    "name": "Choapa",
    "regionCode": "04"
  },
  {
    "code": "043",
    "name": "Limarí",
    "regionCode": "04"
  },
  {
    "code": "051",
    "name": "Valparaíso",
    "regionCode": "05"
  },
  {
    "code": "052",
    "name": "Isla de Pascua",
    "regionCode": "05"
  },
  {
    "code": "053",
    "name": "Los Andes",
    "regionCode": "05"
  },
  {
    "code": "054",
    "name": "Petorca",
    "regionCode": "05"
  },
  {
    "code": "055",
    "name": "Quillota",
    "regionCode": "05"
  },
  {
    "code": "056",
    "name": "San Antonio",
    "regionCode": "05"
  },
  {
    "code": "057",
    "name": "San Felipe de Aconcagua",
    "regionCode": "05"
  },
  {
    "code": "058",
    "name": "Marga Marga",
    "regionCode": "05"
  },
  {
    "code": "061",
    "name": "Cachapoal",
    "regionCode": "06"
  },
  {
    "code": "062",
    "name": "Cardenal Caro",
    "regionCode": "06"
  },
  {
    "code": "063",
    "name": "Colchagua",
    "regionCode": "06"
  },
  {
    "code": "071",
    "name": "Talca",
    "regionCode": "07"
  },
  {
    "code": "072",
    "name": "Cauquenes",
    "regionCode": "07"
  },
  {
    "code": "073",
    "name": "Curicó",
    "regionCode": "07"
  },
  {
    "code": "074",
    "name": "Linares",
    "regionCode": "07"
  },
  {
    "code": "081",
    "name": "Concepción",
    "regionCode": "08"
  },
  {
    "code": "082",
    "name": "Arauco",
    "regionCode": "08"
  },
  {
    "code": "083",
    "name": "Biobío",
    "regionCode": "08"
  },
  {
    "code": "091",
    "name": "Cautín",
    "regionCode": "09"
  },
  {
    "code": "092",
    "name": "Malleco",
    "regionCode": "09"
  },
  {
    "code": "101",
    "name": "Llanquihue",
    "regionCode": "10"
  },
  {
    "code": "102",
    "name": "Chiloé",
    "regionCode": "10"
  },
  {
    "code": "103",
    "name": "Osorno",
    "regionCode": "10"
  },
  {
    "code": "104",
    "name": "Palena",
    "regionCode": "10"
  },
  {
    "code": "111",
    "name": "Coyhaique",
    "regionCode": "11"
  },
  {
    "code": "112",
    "name": "Aysén",
    "regionCode": "11"
  },
  {
    "code": "113",
    "name": "Capitán Prat",
    "regionCode": "11"
  },
  {
    "code": "114",
    "name": "General Carrera",
    "regionCode": "11"
  },
  {
    "code": "121",
    "name": "Magallanes",
    "regionCode": "12"
  },
  {
    "code": "122",
    "name": "Antártica Chilena",
    "regionCode": "12"
  },
  {
    "code": "123",
    "name": "Tierra del Fuego",
    "regionCode": "12"
  },
  {
    "code": "124",
    "name": "Última Esperanza",
    "regionCode": "12"
  },
  {
    "code": "131",
    "name": "Santiago",
    "regionCode": "13"
  },
  {
    "code": "132",
    "name": "Cordillera",
    "regionCode": "13"
  },
  {
    "code": "133",
    "name": "Chacabuco",
    "regionCode": "13"
  },
  {
    "code": "134",
    "name": "Maipo",
    "regionCode": "13"
  },
  {
    "code": "135",
    "name": "Melipilla",
    "regionCode": "13"
  },
  {
    "code": "136",
    "name": "Talagante",
    "regionCode": "13"
  },
  {
    "code": "141",
    "name": "Valdivia",
    "regionCode": "14"
  },
  {
    "code": "142",
    "name": "Ranco",
    "regionCode": "14"
  },
  {
    "code": "151",
    "name": "Arica",
    "regionCode": "15"
  },
  {
    "code": "152",
    "name": "Parinacota",
    "regionCode": "15"
  },
  {
    "code": "161",
    "name": "Diguillín",
    "regionCode": "16"
  },
  {
    "code": "162",
    "name": "Itata",
    "regionCode": "16"
  },
  {
    "code": "163",
    "name": "Punilla",
    "regionCode": "16"
  }
] as const;

export const COMUNAS: readonly Comuna[] = [
  {
    "cut": "01101",
    "name": "Iquique",
    "provinciaCode": "011",
    "regionCode": "01"
  },
  {
    "cut": "01107",
    "name": "Alto Hospicio",
    "provinciaCode": "011",
    "regionCode": "01"
  },
  {
    "cut": "01401",
    "name": "Pozo Almonte",
    "provinciaCode": "014",
    "regionCode": "01"
  },
  {
    "cut": "01402",
    "name": "Camiña",
    "provinciaCode": "014",
    "regionCode": "01"
  },
  {
    "cut": "01403",
    "name": "Colchane",
    "provinciaCode": "014",
    "regionCode": "01"
  },
  {
    "cut": "01404",
    "name": "Huara",
    "provinciaCode": "014",
    "regionCode": "01"
  },
  {
    "cut": "01405",
    "name": "Pica",
    "provinciaCode": "014",
    "regionCode": "01"
  },
  {
    "cut": "02101",
    "name": "Antofagasta",
    "provinciaCode": "021",
    "regionCode": "02"
  },
  {
    "cut": "02102",
    "name": "Mejillones",
    "provinciaCode": "021",
    "regionCode": "02"
  },
  {
    "cut": "02103",
    "name": "Sierra Gorda",
    "provinciaCode": "021",
    "regionCode": "02"
  },
  {
    "cut": "02104",
    "name": "Taltal",
    "provinciaCode": "021",
    "regionCode": "02"
  },
  {
    "cut": "02201",
    "name": "Calama",
    "provinciaCode": "022",
    "regionCode": "02"
  },
  {
    "cut": "02202",
    "name": "Ollagüe",
    "provinciaCode": "022",
    "regionCode": "02"
  },
  {
    "cut": "02203",
    "name": "San Pedro de Atacama",
    "provinciaCode": "022",
    "regionCode": "02"
  },
  {
    "cut": "02301",
    "name": "Tocopilla",
    "provinciaCode": "023",
    "regionCode": "02"
  },
  {
    "cut": "02302",
    "name": "María Elena",
    "provinciaCode": "023",
    "regionCode": "02"
  },
  {
    "cut": "03101",
    "name": "Copiapó",
    "provinciaCode": "031",
    "regionCode": "03"
  },
  {
    "cut": "03102",
    "name": "Caldera",
    "provinciaCode": "031",
    "regionCode": "03"
  },
  {
    "cut": "03103",
    "name": "Tierra Amarilla",
    "provinciaCode": "031",
    "regionCode": "03"
  },
  {
    "cut": "03201",
    "name": "Chañaral",
    "provinciaCode": "032",
    "regionCode": "03"
  },
  {
    "cut": "03202",
    "name": "Diego de Almagro",
    "provinciaCode": "032",
    "regionCode": "03"
  },
  {
    "cut": "03301",
    "name": "Vallenar",
    "provinciaCode": "033",
    "regionCode": "03"
  },
  {
    "cut": "03302",
    "name": "Alto del Carmen",
    "provinciaCode": "033",
    "regionCode": "03"
  },
  {
    "cut": "03303",
    "name": "Freirina",
    "provinciaCode": "033",
    "regionCode": "03"
  },
  {
    "cut": "03304",
    "name": "Huasco",
    "provinciaCode": "033",
    "regionCode": "03"
  },
  {
    "cut": "04101",
    "name": "La Serena",
    "provinciaCode": "041",
    "regionCode": "04"
  },
  {
    "cut": "04102",
    "name": "Coquimbo",
    "provinciaCode": "041",
    "regionCode": "04"
  },
  {
    "cut": "04103",
    "name": "Andacollo",
    "provinciaCode": "041",
    "regionCode": "04"
  },
  {
    "cut": "04104",
    "name": "La Higuera",
    "provinciaCode": "041",
    "regionCode": "04"
  },
  {
    "cut": "04105",
    "name": "Paiguano",
    "provinciaCode": "041",
    "regionCode": "04"
  },
  {
    "cut": "04106",
    "name": "Vicuña",
    "provinciaCode": "041",
    "regionCode": "04"
  },
  {
    "cut": "04201",
    "name": "Illapel",
    "provinciaCode": "042",
    "regionCode": "04"
  },
  {
    "cut": "04202",
    "name": "Canela",
    "provinciaCode": "042",
    "regionCode": "04"
  },
  {
    "cut": "04203",
    "name": "Los Vilos",
    "provinciaCode": "042",
    "regionCode": "04"
  },
  {
    "cut": "04204",
    "name": "Salamanca",
    "provinciaCode": "042",
    "regionCode": "04"
  },
  {
    "cut": "04301",
    "name": "Ovalle",
    "provinciaCode": "043",
    "regionCode": "04"
  },
  {
    "cut": "04302",
    "name": "Combarbalá",
    "provinciaCode": "043",
    "regionCode": "04"
  },
  {
    "cut": "04303",
    "name": "Monte Patria",
    "provinciaCode": "043",
    "regionCode": "04"
  },
  {
    "cut": "04304",
    "name": "Punitaqui",
    "provinciaCode": "043",
    "regionCode": "04"
  },
  {
    "cut": "04305",
    "name": "Río Hurtado",
    "provinciaCode": "043",
    "regionCode": "04"
  },
  {
    "cut": "05101",
    "name": "Valparaíso",
    "provinciaCode": "051",
    "regionCode": "05"
  },
  {
    "cut": "05102",
    "name": "Casablanca",
    "provinciaCode": "051",
    "regionCode": "05"
  },
  {
    "cut": "05103",
    "name": "Concón",
    "provinciaCode": "051",
    "regionCode": "05"
  },
  {
    "cut": "05104",
    "name": "Juan Fernández",
    "provinciaCode": "051",
    "regionCode": "05"
  },
  {
    "cut": "05105",
    "name": "Puchuncaví",
    "provinciaCode": "051",
    "regionCode": "05"
  },
  {
    "cut": "05107",
    "name": "Quintero",
    "provinciaCode": "051",
    "regionCode": "05"
  },
  {
    "cut": "05109",
    "name": "Viña del Mar",
    "provinciaCode": "051",
    "regionCode": "05"
  },
  {
    "cut": "05201",
    "name": "Isla de Pascua",
    "provinciaCode": "052",
    "regionCode": "05"
  },
  {
    "cut": "05301",
    "name": "Los Andes",
    "provinciaCode": "053",
    "regionCode": "05"
  },
  {
    "cut": "05302",
    "name": "Calle Larga",
    "provinciaCode": "053",
    "regionCode": "05"
  },
  {
    "cut": "05303",
    "name": "Rinconada",
    "provinciaCode": "053",
    "regionCode": "05"
  },
  {
    "cut": "05304",
    "name": "San Esteban",
    "provinciaCode": "053",
    "regionCode": "05"
  },
  {
    "cut": "05401",
    "name": "La Ligua",
    "provinciaCode": "054",
    "regionCode": "05"
  },
  {
    "cut": "05402",
    "name": "Cabildo",
    "provinciaCode": "054",
    "regionCode": "05"
  },
  {
    "cut": "05403",
    "name": "Papudo",
    "provinciaCode": "054",
    "regionCode": "05"
  },
  {
    "cut": "05404",
    "name": "Petorca",
    "provinciaCode": "054",
    "regionCode": "05"
  },
  {
    "cut": "05405",
    "name": "Zapallar",
    "provinciaCode": "054",
    "regionCode": "05"
  },
  {
    "cut": "05501",
    "name": "Quillota",
    "provinciaCode": "055",
    "regionCode": "05"
  },
  {
    "cut": "05502",
    "name": "Calera",
    "provinciaCode": "055",
    "regionCode": "05"
  },
  {
    "cut": "05503",
    "name": "Hijuelas",
    "provinciaCode": "055",
    "regionCode": "05"
  },
  {
    "cut": "05504",
    "name": "La Cruz",
    "provinciaCode": "055",
    "regionCode": "05"
  },
  {
    "cut": "05506",
    "name": "Nogales",
    "provinciaCode": "055",
    "regionCode": "05"
  },
  {
    "cut": "05601",
    "name": "San Antonio",
    "provinciaCode": "056",
    "regionCode": "05"
  },
  {
    "cut": "05602",
    "name": "Algarrobo",
    "provinciaCode": "056",
    "regionCode": "05"
  },
  {
    "cut": "05603",
    "name": "Cartagena",
    "provinciaCode": "056",
    "regionCode": "05"
  },
  {
    "cut": "05604",
    "name": "El Quisco",
    "provinciaCode": "056",
    "regionCode": "05"
  },
  {
    "cut": "05605",
    "name": "El Tabo",
    "provinciaCode": "056",
    "regionCode": "05"
  },
  {
    "cut": "05606",
    "name": "Santo Domingo",
    "provinciaCode": "056",
    "regionCode": "05"
  },
  {
    "cut": "05701",
    "name": "San Felipe",
    "provinciaCode": "057",
    "regionCode": "05"
  },
  {
    "cut": "05702",
    "name": "Catemu",
    "provinciaCode": "057",
    "regionCode": "05"
  },
  {
    "cut": "05703",
    "name": "Llaillay",
    "provinciaCode": "057",
    "regionCode": "05"
  },
  {
    "cut": "05704",
    "name": "Panquehue",
    "provinciaCode": "057",
    "regionCode": "05"
  },
  {
    "cut": "05705",
    "name": "Putaendo",
    "provinciaCode": "057",
    "regionCode": "05"
  },
  {
    "cut": "05706",
    "name": "Santa María",
    "provinciaCode": "057",
    "regionCode": "05"
  },
  {
    "cut": "05801",
    "name": "Quilpué",
    "provinciaCode": "058",
    "regionCode": "05"
  },
  {
    "cut": "05802",
    "name": "Limache",
    "provinciaCode": "058",
    "regionCode": "05"
  },
  {
    "cut": "05803",
    "name": "Olmué",
    "provinciaCode": "058",
    "regionCode": "05"
  },
  {
    "cut": "05804",
    "name": "Villa Alemana",
    "provinciaCode": "058",
    "regionCode": "05"
  },
  {
    "cut": "06101",
    "name": "Rancagua",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06102",
    "name": "Codegua",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06103",
    "name": "Coinco",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06104",
    "name": "Coltauco",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06105",
    "name": "Doñihue",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06106",
    "name": "Graneros",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06107",
    "name": "Las Cabras",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06108",
    "name": "Machalí",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06109",
    "name": "Malloa",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06110",
    "name": "Mostazal",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06111",
    "name": "Olivar",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06112",
    "name": "Peumo",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06113",
    "name": "Pichidegua",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06114",
    "name": "Quinta de Tilcoco",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06115",
    "name": "Rengo",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06116",
    "name": "Requínoa",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06117",
    "name": "San Vicente",
    "provinciaCode": "061",
    "regionCode": "06"
  },
  {
    "cut": "06201",
    "name": "Pichilemu",
    "provinciaCode": "062",
    "regionCode": "06"
  },
  {
    "cut": "06202",
    "name": "La Estrella",
    "provinciaCode": "062",
    "regionCode": "06"
  },
  {
    "cut": "06203",
    "name": "Litueche",
    "provinciaCode": "062",
    "regionCode": "06"
  },
  {
    "cut": "06204",
    "name": "Marchihue",
    "provinciaCode": "062",
    "regionCode": "06"
  },
  {
    "cut": "06205",
    "name": "Navidad",
    "provinciaCode": "062",
    "regionCode": "06"
  },
  {
    "cut": "06206",
    "name": "Paredones",
    "provinciaCode": "062",
    "regionCode": "06"
  },
  {
    "cut": "06301",
    "name": "San Fernando",
    "provinciaCode": "063",
    "regionCode": "06"
  },
  {
    "cut": "06302",
    "name": "Chépica",
    "provinciaCode": "063",
    "regionCode": "06"
  },
  {
    "cut": "06303",
    "name": "Chimbarongo",
    "provinciaCode": "063",
    "regionCode": "06"
  },
  {
    "cut": "06304",
    "name": "Lolol",
    "provinciaCode": "063",
    "regionCode": "06"
  },
  {
    "cut": "06305",
    "name": "Nancagua",
    "provinciaCode": "063",
    "regionCode": "06"
  },
  {
    "cut": "06306",
    "name": "Palmilla",
    "provinciaCode": "063",
    "regionCode": "06"
  },
  {
    "cut": "06307",
    "name": "Peralillo",
    "provinciaCode": "063",
    "regionCode": "06"
  },
  {
    "cut": "06308",
    "name": "Placilla",
    "provinciaCode": "063",
    "regionCode": "06"
  },
  {
    "cut": "06309",
    "name": "Pumanque",
    "provinciaCode": "063",
    "regionCode": "06"
  },
  {
    "cut": "06310",
    "name": "Santa Cruz",
    "provinciaCode": "063",
    "regionCode": "06"
  },
  {
    "cut": "07101",
    "name": "Talca",
    "provinciaCode": "071",
    "regionCode": "07"
  },
  {
    "cut": "07102",
    "name": "Constitución",
    "provinciaCode": "071",
    "regionCode": "07"
  },
  {
    "cut": "07103",
    "name": "Curepto",
    "provinciaCode": "071",
    "regionCode": "07"
  },
  {
    "cut": "07104",
    "name": "Empedrado",
    "provinciaCode": "071",
    "regionCode": "07"
  },
  {
    "cut": "07105",
    "name": "Maule",
    "provinciaCode": "071",
    "regionCode": "07"
  },
  {
    "cut": "07106",
    "name": "Pelarco",
    "provinciaCode": "071",
    "regionCode": "07"
  },
  {
    "cut": "07107",
    "name": "Pencahue",
    "provinciaCode": "071",
    "regionCode": "07"
  },
  {
    "cut": "07108",
    "name": "Río Claro",
    "provinciaCode": "071",
    "regionCode": "07"
  },
  {
    "cut": "07109",
    "name": "San Clemente",
    "provinciaCode": "071",
    "regionCode": "07"
  },
  {
    "cut": "07110",
    "name": "San Rafael",
    "provinciaCode": "071",
    "regionCode": "07"
  },
  {
    "cut": "07201",
    "name": "Cauquenes",
    "provinciaCode": "072",
    "regionCode": "07"
  },
  {
    "cut": "07202",
    "name": "Chanco",
    "provinciaCode": "072",
    "regionCode": "07"
  },
  {
    "cut": "07203",
    "name": "Pelluhue",
    "provinciaCode": "072",
    "regionCode": "07"
  },
  {
    "cut": "07301",
    "name": "Curicó",
    "provinciaCode": "073",
    "regionCode": "07"
  },
  {
    "cut": "07302",
    "name": "Hualañé",
    "provinciaCode": "073",
    "regionCode": "07"
  },
  {
    "cut": "07303",
    "name": "Licantén",
    "provinciaCode": "073",
    "regionCode": "07"
  },
  {
    "cut": "07304",
    "name": "Molina",
    "provinciaCode": "073",
    "regionCode": "07"
  },
  {
    "cut": "07305",
    "name": "Rauco",
    "provinciaCode": "073",
    "regionCode": "07"
  },
  {
    "cut": "07306",
    "name": "Romeral",
    "provinciaCode": "073",
    "regionCode": "07"
  },
  {
    "cut": "07307",
    "name": "Sagrada Familia",
    "provinciaCode": "073",
    "regionCode": "07"
  },
  {
    "cut": "07308",
    "name": "Teno",
    "provinciaCode": "073",
    "regionCode": "07"
  },
  {
    "cut": "07309",
    "name": "Vichuquén",
    "provinciaCode": "073",
    "regionCode": "07"
  },
  {
    "cut": "07401",
    "name": "Linares",
    "provinciaCode": "074",
    "regionCode": "07"
  },
  {
    "cut": "07402",
    "name": "Colbún",
    "provinciaCode": "074",
    "regionCode": "07"
  },
  {
    "cut": "07403",
    "name": "Longaví",
    "provinciaCode": "074",
    "regionCode": "07"
  },
  {
    "cut": "07404",
    "name": "Parral",
    "provinciaCode": "074",
    "regionCode": "07"
  },
  {
    "cut": "07405",
    "name": "Retiro",
    "provinciaCode": "074",
    "regionCode": "07"
  },
  {
    "cut": "07406",
    "name": "San Javier",
    "provinciaCode": "074",
    "regionCode": "07"
  },
  {
    "cut": "07407",
    "name": "Villa Alegre",
    "provinciaCode": "074",
    "regionCode": "07"
  },
  {
    "cut": "07408",
    "name": "Yerbas Buenas",
    "provinciaCode": "074",
    "regionCode": "07"
  },
  {
    "cut": "08101",
    "name": "Concepción",
    "provinciaCode": "081",
    "regionCode": "08"
  },
  {
    "cut": "08102",
    "name": "Coronel",
    "provinciaCode": "081",
    "regionCode": "08"
  },
  {
    "cut": "08103",
    "name": "Chiguayante",
    "provinciaCode": "081",
    "regionCode": "08"
  },
  {
    "cut": "08104",
    "name": "Florida",
    "provinciaCode": "081",
    "regionCode": "08"
  },
  {
    "cut": "08105",
    "name": "Hualqui",
    "provinciaCode": "081",
    "regionCode": "08"
  },
  {
    "cut": "08106",
    "name": "Lota",
    "provinciaCode": "081",
    "regionCode": "08"
  },
  {
    "cut": "08107",
    "name": "Penco",
    "provinciaCode": "081",
    "regionCode": "08"
  },
  {
    "cut": "08108",
    "name": "San Pedro de la Paz",
    "provinciaCode": "081",
    "regionCode": "08"
  },
  {
    "cut": "08109",
    "name": "Santa Juana",
    "provinciaCode": "081",
    "regionCode": "08"
  },
  {
    "cut": "08110",
    "name": "Talcahuano",
    "provinciaCode": "081",
    "regionCode": "08"
  },
  {
    "cut": "08111",
    "name": "Tomé",
    "provinciaCode": "081",
    "regionCode": "08"
  },
  {
    "cut": "08112",
    "name": "Hualpén",
    "provinciaCode": "081",
    "regionCode": "08"
  },
  {
    "cut": "08201",
    "name": "Lebu",
    "provinciaCode": "082",
    "regionCode": "08"
  },
  {
    "cut": "08202",
    "name": "Arauco",
    "provinciaCode": "082",
    "regionCode": "08"
  },
  {
    "cut": "08203",
    "name": "Cañete",
    "provinciaCode": "082",
    "regionCode": "08"
  },
  {
    "cut": "08204",
    "name": "Contulmo",
    "provinciaCode": "082",
    "regionCode": "08"
  },
  {
    "cut": "08205",
    "name": "Curanilahue",
    "provinciaCode": "082",
    "regionCode": "08"
  },
  {
    "cut": "08206",
    "name": "Los Alamos",
    "provinciaCode": "082",
    "regionCode": "08"
  },
  {
    "cut": "08207",
    "name": "Tirúa",
    "provinciaCode": "082",
    "regionCode": "08"
  },
  {
    "cut": "08301",
    "name": "Los Angeles",
    "provinciaCode": "083",
    "regionCode": "08"
  },
  {
    "cut": "08302",
    "name": "Antuco",
    "provinciaCode": "083",
    "regionCode": "08"
  },
  {
    "cut": "08303",
    "name": "Cabrero",
    "provinciaCode": "083",
    "regionCode": "08"
  },
  {
    "cut": "08304",
    "name": "Laja",
    "provinciaCode": "083",
    "regionCode": "08"
  },
  {
    "cut": "08305",
    "name": "Mulchén",
    "provinciaCode": "083",
    "regionCode": "08"
  },
  {
    "cut": "08306",
    "name": "Nacimiento",
    "provinciaCode": "083",
    "regionCode": "08"
  },
  {
    "cut": "08307",
    "name": "Negrete",
    "provinciaCode": "083",
    "regionCode": "08"
  },
  {
    "cut": "08308",
    "name": "Quilaco",
    "provinciaCode": "083",
    "regionCode": "08"
  },
  {
    "cut": "08309",
    "name": "Quilleco",
    "provinciaCode": "083",
    "regionCode": "08"
  },
  {
    "cut": "08310",
    "name": "San Rosendo",
    "provinciaCode": "083",
    "regionCode": "08"
  },
  {
    "cut": "08311",
    "name": "Santa Bárbara",
    "provinciaCode": "083",
    "regionCode": "08"
  },
  {
    "cut": "08312",
    "name": "Tucapel",
    "provinciaCode": "083",
    "regionCode": "08"
  },
  {
    "cut": "08313",
    "name": "Yumbel",
    "provinciaCode": "083",
    "regionCode": "08"
  },
  {
    "cut": "08314",
    "name": "Alto Biobío",
    "provinciaCode": "083",
    "regionCode": "08"
  },
  {
    "cut": "09101",
    "name": "Temuco",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09102",
    "name": "Carahue",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09103",
    "name": "Cunco",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09104",
    "name": "Curarrehue",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09105",
    "name": "Freire",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09106",
    "name": "Galvarino",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09107",
    "name": "Gorbea",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09108",
    "name": "Lautaro",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09109",
    "name": "Loncoche",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09110",
    "name": "Melipeuco",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09111",
    "name": "Nueva Imperial",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09112",
    "name": "Padre Las Casas",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09113",
    "name": "Perquenco",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09114",
    "name": "Pitrufquén",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09115",
    "name": "Pucón",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09116",
    "name": "Saavedra",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09117",
    "name": "Teodoro Schmidt",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09118",
    "name": "Toltén",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09119",
    "name": "Vilcún",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09120",
    "name": "Villarrica",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09121",
    "name": "Cholchol",
    "provinciaCode": "091",
    "regionCode": "09"
  },
  {
    "cut": "09201",
    "name": "Angol",
    "provinciaCode": "092",
    "regionCode": "09"
  },
  {
    "cut": "09202",
    "name": "Collipulli",
    "provinciaCode": "092",
    "regionCode": "09"
  },
  {
    "cut": "09203",
    "name": "Curacautín",
    "provinciaCode": "092",
    "regionCode": "09"
  },
  {
    "cut": "09204",
    "name": "Ercilla",
    "provinciaCode": "092",
    "regionCode": "09"
  },
  {
    "cut": "09205",
    "name": "Lonquimay",
    "provinciaCode": "092",
    "regionCode": "09"
  },
  {
    "cut": "09206",
    "name": "Los Sauces",
    "provinciaCode": "092",
    "regionCode": "09"
  },
  {
    "cut": "09207",
    "name": "Lumaco",
    "provinciaCode": "092",
    "regionCode": "09"
  },
  {
    "cut": "09208",
    "name": "Purén",
    "provinciaCode": "092",
    "regionCode": "09"
  },
  {
    "cut": "09209",
    "name": "Renaico",
    "provinciaCode": "092",
    "regionCode": "09"
  },
  {
    "cut": "09210",
    "name": "Traiguén",
    "provinciaCode": "092",
    "regionCode": "09"
  },
  {
    "cut": "09211",
    "name": "Victoria",
    "provinciaCode": "092",
    "regionCode": "09"
  },
  {
    "cut": "10101",
    "name": "Puerto Montt",
    "provinciaCode": "101",
    "regionCode": "10"
  },
  {
    "cut": "10102",
    "name": "Calbuco",
    "provinciaCode": "101",
    "regionCode": "10"
  },
  {
    "cut": "10103",
    "name": "Cochamó",
    "provinciaCode": "101",
    "regionCode": "10"
  },
  {
    "cut": "10104",
    "name": "Fresia",
    "provinciaCode": "101",
    "regionCode": "10"
  },
  {
    "cut": "10105",
    "name": "Frutillar",
    "provinciaCode": "101",
    "regionCode": "10"
  },
  {
    "cut": "10106",
    "name": "Los Muermos",
    "provinciaCode": "101",
    "regionCode": "10"
  },
  {
    "cut": "10107",
    "name": "Llanquihue",
    "provinciaCode": "101",
    "regionCode": "10"
  },
  {
    "cut": "10108",
    "name": "Maullín",
    "provinciaCode": "101",
    "regionCode": "10"
  },
  {
    "cut": "10109",
    "name": "Puerto Varas",
    "provinciaCode": "101",
    "regionCode": "10"
  },
  {
    "cut": "10201",
    "name": "Castro",
    "provinciaCode": "102",
    "regionCode": "10"
  },
  {
    "cut": "10202",
    "name": "Ancud",
    "provinciaCode": "102",
    "regionCode": "10"
  },
  {
    "cut": "10203",
    "name": "Chonchi",
    "provinciaCode": "102",
    "regionCode": "10"
  },
  {
    "cut": "10204",
    "name": "Curaco de Vélez",
    "provinciaCode": "102",
    "regionCode": "10"
  },
  {
    "cut": "10205",
    "name": "Dalcahue",
    "provinciaCode": "102",
    "regionCode": "10"
  },
  {
    "cut": "10206",
    "name": "Puqueldón",
    "provinciaCode": "102",
    "regionCode": "10"
  },
  {
    "cut": "10207",
    "name": "Queilén",
    "provinciaCode": "102",
    "regionCode": "10"
  },
  {
    "cut": "10208",
    "name": "Quellón",
    "provinciaCode": "102",
    "regionCode": "10"
  },
  {
    "cut": "10209",
    "name": "Quemchi",
    "provinciaCode": "102",
    "regionCode": "10"
  },
  {
    "cut": "10210",
    "name": "Quinchao",
    "provinciaCode": "102",
    "regionCode": "10"
  },
  {
    "cut": "10301",
    "name": "Osorno",
    "provinciaCode": "103",
    "regionCode": "10"
  },
  {
    "cut": "10302",
    "name": "Puerto Octay",
    "provinciaCode": "103",
    "regionCode": "10"
  },
  {
    "cut": "10303",
    "name": "Purranque",
    "provinciaCode": "103",
    "regionCode": "10"
  },
  {
    "cut": "10304",
    "name": "Puyehue",
    "provinciaCode": "103",
    "regionCode": "10"
  },
  {
    "cut": "10305",
    "name": "Río Negro",
    "provinciaCode": "103",
    "regionCode": "10"
  },
  {
    "cut": "10306",
    "name": "San Juan de la Costa",
    "provinciaCode": "103",
    "regionCode": "10"
  },
  {
    "cut": "10307",
    "name": "San Pablo",
    "provinciaCode": "103",
    "regionCode": "10"
  },
  {
    "cut": "10401",
    "name": "Chaitén",
    "provinciaCode": "104",
    "regionCode": "10"
  },
  {
    "cut": "10402",
    "name": "Futaleufú",
    "provinciaCode": "104",
    "regionCode": "10"
  },
  {
    "cut": "10403",
    "name": "Hualaihué",
    "provinciaCode": "104",
    "regionCode": "10"
  },
  {
    "cut": "10404",
    "name": "Palena",
    "provinciaCode": "104",
    "regionCode": "10"
  },
  {
    "cut": "11101",
    "name": "Coyhaique",
    "provinciaCode": "111",
    "regionCode": "11"
  },
  {
    "cut": "11102",
    "name": "Lago Verde",
    "provinciaCode": "111",
    "regionCode": "11"
  },
  {
    "cut": "11201",
    "name": "Aysén",
    "provinciaCode": "112",
    "regionCode": "11"
  },
  {
    "cut": "11202",
    "name": "Cisnes",
    "provinciaCode": "112",
    "regionCode": "11"
  },
  {
    "cut": "11203",
    "name": "Guaitecas",
    "provinciaCode": "112",
    "regionCode": "11"
  },
  {
    "cut": "11301",
    "name": "Cochrane",
    "provinciaCode": "113",
    "regionCode": "11"
  },
  {
    "cut": "11302",
    "name": "O'Higgins",
    "provinciaCode": "113",
    "regionCode": "11"
  },
  {
    "cut": "11303",
    "name": "Tortel",
    "provinciaCode": "113",
    "regionCode": "11"
  },
  {
    "cut": "11401",
    "name": "Chile Chico",
    "provinciaCode": "114",
    "regionCode": "11"
  },
  {
    "cut": "11402",
    "name": "Río Ibáñez",
    "provinciaCode": "114",
    "regionCode": "11"
  },
  {
    "cut": "12101",
    "name": "Punta Arenas",
    "provinciaCode": "121",
    "regionCode": "12"
  },
  {
    "cut": "12102",
    "name": "Laguna Blanca",
    "provinciaCode": "121",
    "regionCode": "12"
  },
  {
    "cut": "12103",
    "name": "Río Verde",
    "provinciaCode": "121",
    "regionCode": "12"
  },
  {
    "cut": "12104",
    "name": "San Gregorio",
    "provinciaCode": "121",
    "regionCode": "12"
  },
  {
    "cut": "12201",
    "name": "Cabo de Hornos",
    "provinciaCode": "122",
    "regionCode": "12"
  },
  {
    "cut": "12202",
    "name": "Antártica",
    "provinciaCode": "122",
    "regionCode": "12"
  },
  {
    "cut": "12301",
    "name": "Porvenir",
    "provinciaCode": "123",
    "regionCode": "12"
  },
  {
    "cut": "12302",
    "name": "Primavera",
    "provinciaCode": "123",
    "regionCode": "12"
  },
  {
    "cut": "12303",
    "name": "Timaukel",
    "provinciaCode": "123",
    "regionCode": "12"
  },
  {
    "cut": "12401",
    "name": "Natales",
    "provinciaCode": "124",
    "regionCode": "12"
  },
  {
    "cut": "12402",
    "name": "Torres del Paine",
    "provinciaCode": "124",
    "regionCode": "12"
  },
  {
    "cut": "13101",
    "name": "Santiago",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13102",
    "name": "Cerrillos",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13103",
    "name": "Cerro Navia",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13104",
    "name": "Conchalí",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13105",
    "name": "El Bosque",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13106",
    "name": "Estación Central",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13107",
    "name": "Huechuraba",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13108",
    "name": "Independencia",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13109",
    "name": "La Cisterna",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13110",
    "name": "La Florida",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13111",
    "name": "La Granja",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13112",
    "name": "La Pintana",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13113",
    "name": "La Reina",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13114",
    "name": "Las Condes",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13115",
    "name": "Lo Barnechea",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13116",
    "name": "Lo Espejo",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13117",
    "name": "Lo Prado",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13118",
    "name": "Macul",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13119",
    "name": "Maipú",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13120",
    "name": "Ñuñoa",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13121",
    "name": "Pedro Aguirre Cerda",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13122",
    "name": "Peñalolén",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13123",
    "name": "Providencia",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13124",
    "name": "Pudahuel",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13125",
    "name": "Quilicura",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13126",
    "name": "Quinta Normal",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13127",
    "name": "Recoleta",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13128",
    "name": "Renca",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13129",
    "name": "San Joaquín",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13130",
    "name": "San Miguel",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13131",
    "name": "San Ramón",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13132",
    "name": "Vitacura",
    "provinciaCode": "131",
    "regionCode": "13"
  },
  {
    "cut": "13201",
    "name": "Puente Alto",
    "provinciaCode": "132",
    "regionCode": "13"
  },
  {
    "cut": "13202",
    "name": "Pirque",
    "provinciaCode": "132",
    "regionCode": "13"
  },
  {
    "cut": "13203",
    "name": "San José de Maipo",
    "provinciaCode": "132",
    "regionCode": "13"
  },
  {
    "cut": "13301",
    "name": "Colina",
    "provinciaCode": "133",
    "regionCode": "13"
  },
  {
    "cut": "13302",
    "name": "Lampa",
    "provinciaCode": "133",
    "regionCode": "13"
  },
  {
    "cut": "13303",
    "name": "Tiltil",
    "provinciaCode": "133",
    "regionCode": "13"
  },
  {
    "cut": "13401",
    "name": "San Bernardo",
    "provinciaCode": "134",
    "regionCode": "13"
  },
  {
    "cut": "13402",
    "name": "Buin",
    "provinciaCode": "134",
    "regionCode": "13"
  },
  {
    "cut": "13403",
    "name": "Calera de Tango",
    "provinciaCode": "134",
    "regionCode": "13"
  },
  {
    "cut": "13404",
    "name": "Paine",
    "provinciaCode": "134",
    "regionCode": "13"
  },
  {
    "cut": "13501",
    "name": "Melipilla",
    "provinciaCode": "135",
    "regionCode": "13"
  },
  {
    "cut": "13502",
    "name": "Alhué",
    "provinciaCode": "135",
    "regionCode": "13"
  },
  {
    "cut": "13503",
    "name": "Curacaví",
    "provinciaCode": "135",
    "regionCode": "13"
  },
  {
    "cut": "13504",
    "name": "María Pinto",
    "provinciaCode": "135",
    "regionCode": "13"
  },
  {
    "cut": "13505",
    "name": "San Pedro",
    "provinciaCode": "135",
    "regionCode": "13"
  },
  {
    "cut": "13601",
    "name": "Talagante",
    "provinciaCode": "136",
    "regionCode": "13"
  },
  {
    "cut": "13602",
    "name": "El Monte",
    "provinciaCode": "136",
    "regionCode": "13"
  },
  {
    "cut": "13603",
    "name": "Isla de Maipo",
    "provinciaCode": "136",
    "regionCode": "13"
  },
  {
    "cut": "13604",
    "name": "Padre Hurtado",
    "provinciaCode": "136",
    "regionCode": "13"
  },
  {
    "cut": "13605",
    "name": "Peñaflor",
    "provinciaCode": "136",
    "regionCode": "13"
  },
  {
    "cut": "14101",
    "name": "Valdivia",
    "provinciaCode": "141",
    "regionCode": "14"
  },
  {
    "cut": "14102",
    "name": "Corral",
    "provinciaCode": "141",
    "regionCode": "14"
  },
  {
    "cut": "14103",
    "name": "Lanco",
    "provinciaCode": "141",
    "regionCode": "14"
  },
  {
    "cut": "14104",
    "name": "Los Lagos",
    "provinciaCode": "141",
    "regionCode": "14"
  },
  {
    "cut": "14105",
    "name": "Máfil",
    "provinciaCode": "141",
    "regionCode": "14"
  },
  {
    "cut": "14106",
    "name": "Mariquina",
    "provinciaCode": "141",
    "regionCode": "14"
  },
  {
    "cut": "14107",
    "name": "Paillaco",
    "provinciaCode": "141",
    "regionCode": "14"
  },
  {
    "cut": "14108",
    "name": "Panguipulli",
    "provinciaCode": "141",
    "regionCode": "14"
  },
  {
    "cut": "14201",
    "name": "La Unión",
    "provinciaCode": "142",
    "regionCode": "14"
  },
  {
    "cut": "14202",
    "name": "Futrono",
    "provinciaCode": "142",
    "regionCode": "14"
  },
  {
    "cut": "14203",
    "name": "Lago Ranco",
    "provinciaCode": "142",
    "regionCode": "14"
  },
  {
    "cut": "14204",
    "name": "Río Bueno",
    "provinciaCode": "142",
    "regionCode": "14"
  },
  {
    "cut": "15101",
    "name": "Arica",
    "provinciaCode": "151",
    "regionCode": "15"
  },
  {
    "cut": "15102",
    "name": "Camarones",
    "provinciaCode": "151",
    "regionCode": "15"
  },
  {
    "cut": "15201",
    "name": "Putre",
    "provinciaCode": "152",
    "regionCode": "15"
  },
  {
    "cut": "15202",
    "name": "General Lagos",
    "provinciaCode": "152",
    "regionCode": "15"
  },
  {
    "cut": "16101",
    "name": "Chillán",
    "provinciaCode": "161",
    "regionCode": "16"
  },
  {
    "cut": "16102",
    "name": "Bulnes",
    "provinciaCode": "161",
    "regionCode": "16"
  },
  {
    "cut": "16103",
    "name": "Chillán Viejo",
    "provinciaCode": "161",
    "regionCode": "16"
  },
  {
    "cut": "16104",
    "name": "El Carmen",
    "provinciaCode": "161",
    "regionCode": "16"
  },
  {
    "cut": "16105",
    "name": "Pemuco",
    "provinciaCode": "161",
    "regionCode": "16"
  },
  {
    "cut": "16106",
    "name": "Pinto",
    "provinciaCode": "161",
    "regionCode": "16"
  },
  {
    "cut": "16107",
    "name": "Quillón",
    "provinciaCode": "161",
    "regionCode": "16"
  },
  {
    "cut": "16108",
    "name": "San Ignacio",
    "provinciaCode": "161",
    "regionCode": "16"
  },
  {
    "cut": "16109",
    "name": "Yungay",
    "provinciaCode": "161",
    "regionCode": "16"
  },
  {
    "cut": "16201",
    "name": "Quirihue",
    "provinciaCode": "162",
    "regionCode": "16"
  },
  {
    "cut": "16202",
    "name": "Cobquecura",
    "provinciaCode": "162",
    "regionCode": "16"
  },
  {
    "cut": "16203",
    "name": "Coelemu",
    "provinciaCode": "162",
    "regionCode": "16"
  },
  {
    "cut": "16204",
    "name": "Ninhue",
    "provinciaCode": "162",
    "regionCode": "16"
  },
  {
    "cut": "16205",
    "name": "Portezuelo",
    "provinciaCode": "162",
    "regionCode": "16"
  },
  {
    "cut": "16206",
    "name": "Ranquil",
    "provinciaCode": "162",
    "regionCode": "16"
  },
  {
    "cut": "16207",
    "name": "Treguaco",
    "provinciaCode": "162",
    "regionCode": "16"
  },
  {
    "cut": "16301",
    "name": "San Carlos",
    "provinciaCode": "163",
    "regionCode": "16"
  },
  {
    "cut": "16302",
    "name": "Coihueco",
    "provinciaCode": "163",
    "regionCode": "16"
  },
  {
    "cut": "16303",
    "name": "Ñiquén",
    "provinciaCode": "163",
    "regionCode": "16"
  },
  {
    "cut": "16304",
    "name": "San Fabián",
    "provinciaCode": "163",
    "regionCode": "16"
  },
  {
    "cut": "16305",
    "name": "San Nicolás",
    "provinciaCode": "163",
    "regionCode": "16"
  }
] as const;
