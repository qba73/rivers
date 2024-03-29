DROP TABLE IF EXISTS stations;

CREATE TABLE stations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name CHAR(50) NOT NULL,
    ref_id INT NOT NULL,
    long REAL,
    lat REAL,
    group_id INT);


INSERT INTO stations (name, ref_id, long, lat) VALUES
("Sandy Mills", 1041, -7.575758, 54.838318),
("Ballybofey", 1043, -7.790749, 54.799769),
("Glaslough", 3055, -6.894344, 54.323281),
("Cappog Bridge", 3058, -7.021297, 54.266809),
("Moyles Mill", 6011, -6.596077, 54.011574),
("Clarebane", 6012, -6.666056, 54.092856),
("Charleville Weir", 6013, -6.413996, 53.855843),
("Tallanstown Weir", 6014, -6.54957, 53.921092),
("Brewery Park", 6015, -6.416475, 53.9936),
("Mansfieldstown", 6021, -6.444491, 53.896604),
("Burley", 6025, -6.594341, 53.848038),
("Aclint", 6026, -6.640019, 53.924765),
("Ladyswell", 6036, -6.405590, 53.993758),
("Port Oriel", 6060, -6.221440, 53.798040),
("Dundalk Port", 6061, -6.385468, 54.007688),
("Tremblestown", 7001, -6.855643, 53.562065),
("Killyon", 7002, -6.970771, 53.487771),
("Castlerickard", 7003, -6.92185, 53.485524),
("Stramatt", 7004, -7.043613, 53.796298),
("Trim", 7005, -6.791844, 53.556405),
("Fyanstown", 7006, -6.802883, 53.725823),
("Boyne Aqueduct", 7007, -6.958844, 53.453137),
("Navan Weir", 7009, -6.672058, 53.643559),
("Liscartan", 7010, -6.720333, 53.663269),
("O'Dalys Bridge", 7011, -7.011203, 53.768982),
("Slane Castle", 7012, -6.562423, 53.707212),
("Virginia Hatchery", 7033, -7.078577, 53.83437),
("Blackcastle", 7037, -6.680586, 53.654697),
("Tinker's Bridge", 7049, -7.331587, 53.344534),
("Mornington Bridge", 7062, -6.254403, 53.719427),
("Virginia", 7081, -7.090788, 53.831615),
("Newfield", 7115, -6.340713, 53.723688),
("Broadmeadow", 8008, -6.231775, 53.474922),
("Duleek D/S", 8011, -6.407709, 53.656184),
("Leixlip", 9001, -6.490439, 53.368665),
("Waldron's Bridge", 9010, -6.266962, 53.305816),
("Botanic Gardens Backup", 9045, -6.276198, 53.375159),
("Tipper", 9107, -6.625515, 53.204874),
("Killashee", 9108, -6.672742, 53.194734),
("Haynestown", 9109, -6.593050, 53.216592),
("Bluebell", 9110, -6.676842, 53.208095),
("Hazelhatch", 9111, -6.529315, 53.330543),
("Arklow Town Bridge", 10042, -6.152069, 52.798265),
("Glenavon Park", 10047, -6.126190, 53.247336),
("Cherry Wood", 10048, -6.137616, 53.247002),
("Meadow Vale Park", 10049, -6.153028, 53.269589),
("Brides Glen", 10050, -6.146907, 53.238755),
("Arklow Harbour", 10060, -6.145231, 52.792047),
("Boleany", 11001, -6.270794, 52.64369),
("Scarawalsh", 12001, -6.550222, 52.548513),
("Enniscorthy", 12002, -6.566852, 52.502307),
("Tullow Town Bridge U/S", 12005, -6.738373, 52.802013),
("Tullowbeg", 12006, -6.738397, 52.796064),
("St. Johns Bridge", 12007, -6.573359, 52.493291),
("Rafter Bridge D/S", 12008, -6.564173, 52.500918),
("Rafter Bridge U/S", 12009, -6.564198, 52.501061),
("Edermine Bridge", 12061, -6.562578, 52.454094),
("Assaly", 12063, -6.43554, 52.279054),
("Ferrycarrig Bridge", 12064, -6.511452, 52.350837),
("Lady's Island", 13070, -6.383141, 52.211004),
("Sigginstown", 13072, -6.454423, 52.201860),
("Cull Pump House", 13081, -6.629581, 52.204544),
("Carlow", 14001, -6.938009, 52.834224),
("Borness", 14003, -7.308429, 53.132249),
("Clonbulloge", 14004, -7.086854, 53.258656),
("Portarlington", 14005, -7.192953, 53.161813),
("Pass Bridge", 14006, -7.070581, 53.145923),
("Derrybrock", 14007, -7.085032, 53.039052),
("Cushina", 14009, -7.174493, 53.194321),
("Rathangan", 14011, -6.992655, 53.220616),
("Ballinacarrig", 14013, -6.898154, 52.824264),
("Royal Oak", 14018, -6.981458, 52.700198),
("Levitstown", 14019, -6.949797, 52.935378),
("Barrow New Bridge", 14022, -6.931752, 52.846138),
("Graiguenamanagh U/S", 14029, -6.950884, 52.540754),
("Milford Lock", 14056, -6.963589, 52.778188),
("St. Mullins", 14067, -6.92429, 52.485473),
("Chapel Street", 14120, -7.337003, 53.119496),
("Manor Road", 14121, -7.338814, 53.115676),
("Turf Market", 14122, -6.955819, 52.540405),
("Coolroe", 14123, -6.991390, 52.547068),
("Annamult", 15001, -7.199625, 52.548458),
("John's Bridge  Nore", 15002, -7.250433, 52.653392),
("Dinin Bridge", 15003, -7.291781, 52.715345),
("Mcmahons Bridge", 15004, -7.379327, 52.867024),
("Durrow Foot Bridge", 15005, -7.397987, 52.847064),
("Brownsbarn", 15006, -7.0917, 52.50077),
("Kilbricken", 15007, -7.461606, 52.959252),
("Borris-In-Ossory", 15008, -7.644113, 52.943091),
("Callan", 15009, -7.388463, 52.545101),
("Ballyboodin", 15010, -7.454003, 52.846922),
("Mount Juliet", 15011, -7.189299, 52.531251),
("Blackfriar's Bridge", 15050, -7.259453, 52.654287),
("Sycamores", 15104, -7.256382, 52.665920),
("Archers Grove", 15105, -7.217501, 52.641541),
("Athlummon", 16001, -7.739544, 52.686101),
("Beakstown", 16002, -7.864804, 52.648839),
("Rathkennan", 16003, -7.924926, 52.629303),
("Thurles", 16004, -7.809622, 52.679064),
("Aughnagross", 16005, -8.014063, 52.523496),
("Ballinaclogh", 16006, -8.022296, 52.519469),
("Killardry", 16007, -7.975714, 52.417276),
("New Bridge  Suir", 16008, -7.99789, 52.45947),
("Caher Park", 16009,-7.922937, 52.357673),
("Anner", 16010, -7.628552, 52.382079),
("Clonmel", 16011, -7.694598, 52.351525),
("Tar Bridge", 16012, -7.842907, 52.273004),
("Fourmilewater FFWS Main Logger",16013, -7.756046, 52.273799),
("Clobanna", 16051, -7.791764, 52.715831),
("Fiddown", 16061, -7.315921, 52.328012),
("Carrick On Suir", 16062, -7.410374, 52.344104),
("Sheep's Bridge Weir",16115, -7.127922, 52.224183),
("Piltown", 16125, -7.324867, 52.349317),
("Tramore RD RDBT", 16128, -7.119044, 52.246922),
("John's Bridge", 16129, -7.110892, 52.256306),
("Small Bridge  Templemore", 16136, -7.837187, 52.792651),
("Newcastle Bridge", 16137, -7.810275, 52.274915),
("Ballydonagh", 16138, -7.794498, 52.296807),
("Knocklofyty", 16139, -7.789266, 52.336734),
("Ardfinnan Road", 16146, -7.746476, 52.341474),
("Joyce's Lane", 16147, -7.706706, 52.351555),
("Workhouse Bridge", 16148, -7.716815, 52.349656),
("Clonmel FFWS back-up", 16149, -7.694477, 52.351521),
("Caher Park FFWS back-up", 16150, -7.922937, 52.357673),
("Tar Bridge FFWS back-up", 16151, -7.842907, 52.273004),
("Fourmilewater FFWS back-up", 16152, -7.757987, 52.273744),
("New Bridge Suir FFWS back-up", 16153, -7.99789, 52.45947),
("Killardry FFWS back-up", 16154, -7.975714, 52.417276),
("Adelphi Quay", 16160, -7.102433, 52.259666),
("Dunmore East", 17061, -6.990829, 52.147523),
("Mogeely", 18001, -8.064303, 52.099557),
("Ballyduff", 18002, -8.051952, 52.144345),
("Killavullen", 18003, -8.515394, 52.149198),
("Ballynamona", 18004, -8.503225, 52.219141),
("Downing Bridge", 18005, -8.25896, 52.168534),
("Fr Murphy's Bridge", 18019, -8.889025, 52.120741),
("Glenavuddig Bridge", 18024, -8.405723, 52.248571),
("Glandalane", 18053, -8.220838, 52.149811),
("Mallow Railway Bridge", 18055, -8.656719, 52.131124),
("Mallow Town Bridge U/S", 18056, -8.641608, 52.13234),
("Mallow Town Bridge D/S", 18057, -8.641037, 52.132218),
("Youghal Quay", 18061, -7.847605, 51.957226),
("Castletownroche Weir", 18102, -8.460250, 52.173700),
("Castlelands", 18105, -8.622861, 52.134705),
("Fermoy Bridge U/S", 18106, -8.276634, 52.138568),
("Fermoy Bridge D/S", 18107, -8.275398, 52.139551),
("Araglin Bridge", 18108, -8.220864, 52.166895),
("Lombardstown Bridge", 18109, -8.783209, 52.122615),
("Kilbrin Road", 18110, -8.903999, 52.179203),
("Church Street", 18111, -8.910823, 52.178766),
("Keale Bridge", 18112, -9.028995, 52.089622),
("Ahane Bridge", 18113, -9.132805, 52.096076),
("Clashmorgan", 18114, -8.681001, 52.082987),
("Jordans Bridge", 18115, -8.625179, 52.078048),
("Fermoy Mill", 18117, -8.271953, 52.13972),
("Ballydahin", 18119, -8.654214, 52.131390),
("Nursetownbeg", 18120, -8.651831, 52.083151),
("Shronebeha", 18121, -8.888915, 52.110676),
("Gortageen", 18122, -9.030316, 52.089880),
("Greenane", 18123, -8.904016, 52.178989),
("Fermoy U/S Rowing", 18124, -8.281061, 52.138381),
("Ballea", 19001, -8.421563, 51.822081),
("Kilmona Bridge", 19044, -8.588573, 51.989534),
("Gothic Bridge", 19045, -8.558214, 51.928057),
("Ballyvourney", 19054,-9.161190, 51.939475),
("Ballymakera", 19055, -9.146320, 51.934349),
("Ballincolly", 19056, -8.456144, 51.922322),
("Glen Park", 19057, -8.452116, 51.912902),
("Blackpool Retail Park", 19058, -8.473423, 51.916579),
("Glennamought Br.", 19059, -8.477345, 51.929256),
("Ballycotton", 19068, -8.001473, 51.828183),
("Ringaskiddy NMCI", 19069, -8.305565, 51.834962),
("Inniscarra Headrace", 19094, -8.661623, 51.900081),
("Carrigadrohid Headrace", 19095, -8.864114, 51.896970),
("Macroom WWTP", 19100, -8.945926, 51.905393),
("Macroom Town Bridge", 19101, -8.962492, 51.906039),
("Waterworks Weir", 19102, -8.509944, 51.893964),
("Ovens Bridge", 19103, -8.655078, 51.880535),
("Morris's Bridge", 19104, -8.936169, 51.929307),
("Muskerry", 19105, -8.601826, 51.914717),
("Cooldaniel", 19106, -9.022294, 51.857668),
("Dripsey Bridge", 19107, -8.745839, 51.916024),
("Bawnafinny Bridge", 19108, -8.585341, 51.930029),
("Inniscarra Tailrace", 19109, -8.633395, 51.892008),
("Cooleen Bridge", 19110, -9.074104, 51.871030),
("Killaclug", 19111, -9.023169, 51.911606),
("Currach Club", 19160, -8.443842, 51.901557),
("Bandon", 20001, -8.731538, 51.746954),
("Curranure", 20002, -8.682672, 51.765232),
("Long Bridge Dunmanway", 20008, -9.097887, 51.724699),
("Ardcahan Bridge", 20015, -9.097987, 51.749136),
("Bealaboy Bridge", 20016, -9.075759, 51.709436),
("Clonakilty", 20019, -8.894235, 51.622634),
("Skibbereen", 20020, -9.264852, 51.551973),
("Sneem D/S", 21018, -9.899291, 51.841750),
("Bridge Street", 21020, -9.585794, 51.881086),
("Kenmare Pier", 21062, -9.588986, 51.872163),
("Riverville", 22003, -9.570732, 52.197621),
("Torc Weir", 22005, -9.50622, 52.000977),
("Flesk Bridge", 22006, -9.497947, 52.048034),
("White Bridge", 22009, -9.527219, 52.054921),
("Old Weir Bridge", 22016, -9.549443, 52.007436),
("Laune Bridge", 22035, -9.617092, 52.06155),
("Castlemaine", 22061, -9.70277, 52.167595),
("Tomies Pier", 22071, -9.605025, 52.056843),
("Bvm Park", 22082, -9.506491, 52.023362),
("Inch Bridge Galey", 23001, -9.534399, 52.468052),
("Listowel", 23002, -9.475689, 52.442655),
("Ballymullen", 23012, -9.691575, 52.260059),
("Sleveen Main Channel", 23030, -9.636918, 52.432616),
("Sleveen Back Channel", 23033, -9.637055, 52.432678),
("Sleveen Back Channel Rattoo Pump House", 23039, -9.636029, 52.438498),
("Sleeveen Back channel 100m U/S", 23040, -9.635125, 52.437969),
("Sleeveen Main Channel DS", 23041, -9.635651, 52.438723),
("Sleeveen Back Channel D/S", 23042, -9.636906, 52.441228),
("Lisloose", 23044, -9.699118, 52.288341),
("Ballyseedy", 23045, -9.641718, 52.256190),
("Spring Water Lane", 23046, -9.689855, 52.254659),
("Oakpark Road", 23047, -9.681367, 52.285828),
("Abbeydorney D/S", 23048, -9.688174, 52.347213),
("Manor West", 23049, -9.675723, 52.263356),
("Ferry Bridge Feale", 23061, -9.633367, 52.469001),
("Blennerville", 23062, -9.735782, 52.257899),
("Ballyard", 23063, -9.711965, 52.262768),
("Fenit", 23066, -9.863564, 52.270703),
("Moneycashen", 23068, -9.680923, 52.48275),
("Croom", 24001, -8.718469, 52.519224),
("Gray's Bridge", 24002, -8.619798, 52.512988),
("Bruree", 24004, -8.660844, 52.423225),
("Athlacca", 24005, -8.651166, 52.458766),
("Castleroberts", 24008, -8.767416, 52.543441),
("Adare Manor", 24009, -8.777368, 52.564413),
("Deel Bridge", 24011, -9.031113, 52.442145),
("Grange Bridge", 24012, -9.018854, 52.462813),
("Rathkeale", 24013, -8.943546, 52.52096),
("Riversfield Weir", 24034, -8.540769, 52.387574),
("Rossbrien Railway Bridge", 24047, -8.632614, 52.636086),
("Ferry Bridge  Maigue", 24061, -8.764818, 52.621226),
("Adare Quay", 24062, -8.798123, 52.568867),
("Limerick Dock", 24063, -8.644473, 52.658499),
("Foynes", 24064, -9.101269, 52.614159),
("Normoyle's Bridge", 24067, -8.825606, 52.559809),
("Islandmore Weir", 24082, -8.715255, 52.509116),
("Gortboy Hotel", 24100, -9.050871, 52.448666),
("Annacotty", 25001, -8.529147, 52.669276),
("Barrington's Bridge", 25002, -8.475033, 52.644946),
("Abington", 25003, -8.421221, 52.631868),
("New Bridge  Bilboa", 25004, -8.314669, 52.591644),
("Sunville", 25005, -8.329348, 52.581468),
("Ferbane", 25006, -7.828063, 53.269938),
("Moystown", 25011, -7.932055, 53.238253),
("Millbrook", 25014, -7.79795, 53.219546),
("Pollagh", 25015, -7.715862, 53.281349),
("Rahan  Clodiagh", 25016, -7.615676, 53.280799),
("Banagher", 25017, -7.993647, 53.193787),
("Conicar", 25019, -8.370801, 53.114751),
("Killeen", 25020, -8.303336, 53.149715),
("Croghan", 25021, -7.920462, 53.101742),
("Syngefield", 25022, -7.88158, 53.092805),
("Milltown", 25023, -7.897309, 52.96926),
("New Bridge Little Brosna", 25024, -7.975743, 53.131912),
("Ballyhooney", 25025, -8.205766, 53.01445),
("Gourdeen", 25027, -8.168575, 52.868412),
("Clarianna", 25029, -8.20771, 52.891623),
("Scarriff", 25030, -8.533125, 52.908365),
("Mullingar Pump Hse", 25050, -7.334809, 53.527883),
("Meelick Weir", 25056, -8.076068, 53.17608),
("Victoria Lock", 25058, -8.080451, 53.168772),
("Ball's Bridge", 25061, -8.61938, 52.666458),
("Tullamore", 25149, -7.501241, 53.273214),
("Culleen Fish Farm", 25213, -7.351569, 53.549125),
("Bracknagh Bridge", 25301, -7.504752, 53.180797),
("Waterpark Bridge", 25308, -8.464589, 52.695246),
("Clonsingle Bridge", 25309, -8.431994, 52.685081),
("Ballinamore", 26001, -8.366002, 53.489786),
("Rookwood", 26002, -8.292659, 53.563717),
("Bookala", 26004, -8.512284, 53.705593),
("Derrycahill", 26005, -8.262762, 53.431764),
("Willsbrook", 26006, -8.466028, 53.729357),
("Bellagill", 26007, -8.238554, 53.361687),
("Johnston's Bridge", 26008, -7.862746, 53.82777),
("Bellantra Bridge", 26009, -7.805157, 53.854401),
("Riverstown", 26010, -7.815074, 53.931213),
("Banada Bridge Lung", 26014, -8.557031, 53.897014),
("Corrascoffy", 26015, -7.918866, 53.886635),
("Bellavahan Bridge", 26018, -8.075036, 53.827979),
("Mullagh", 26019, -7.823850, 53.732942),
("Argar", 26020, -7.725933, 53.763855),
("Ballymahon", 26021, -7.757923, 53.562867),
("Kilmore", 26022, -7.872534, 53.710257),
("Camagh", 26025, -7.407165, 53.728937),
("Athlone", 26027, -7.940756, 53.421534),
("Shannonbridge", 26028, -8.049581, 53.279873),
("Blackrock Lock", 26074, -8.050638, 54.048462),
("Cuppanagh", 26075, -8.406391, 53.958279),
("Lough Rinn", 26079, -7.851744, 53.893989),
("Lough Derravaragh", 26082, -7.30328, 53.611316),
("Mount Nugent", 26083, -7.283102, 53.820952),
("Jamestown", 26085, -8.030334, 53.92306),
("Cuil Bridge", 26086, -8.434775, 53.929116),
("Lomcloon", 26087, -8.467149, 53.924673),
("Hodson's Bay", 26088, -7.986860, 53.467425),
("Drumsna", 26089, -8.008173, 53.923144),
("Derry Bay L. Ree", 26093, -7.849053, 53.556108),
("Ballinalack", 26104, -7.474758, 53.63141),
("Boyle Abbey Bridge", 26108, -8.296673, 53.972839),
("Ahascragh Pump Hse", 26140, -8.331003, 53.392636),
("Glen Lough Lower", 26305, -7.566716, 53.648283),
("Glen Lough Sluice", 26307, -7.565994, 53.647894),
("Carrick-On-Shannon", 26324, -8.095558, 53.943210),
("Athlone Weir U/S", 26333, -7.941778, 53.422891),
("Derryholmes  u/s", 26351, -8.001045, 53.248014),
("Derryholmes d/s", 26352, -7.994703, 53.243913),
("W.O.P.S Rail Bridge", 26353, -8.039517, 53.268136),
("Ballinasloe Town", 26354, -8.214933, 53.329500),
("Ballinasloe Old Channel", 26355, -8.218675, 53.330484),
("Raherbeg Rail Bridge", 26356, -8.086295, 53.274019),
("Deerpark Bridge", 26357, -8.261250, 53.338883),
("Bunowen Bridge", 26358, -8.253380, 53.350456),
("Scragh Bog", 26374, -7.366132, 53.585366),
("Inch Bridge Claureen", 27001, -9.036242, 52.825291),
("Ballycorey", 27002, -8.974654, 52.870043),
("Corrofin  Fergus", 27003, -9.061919, 52.943762),
("Carnelly", 27004, -8.936797, 52.808478),
("Owenogarney Railway Bridge", 27011, -8.771001, 52.732704),
("Victoria Bridge", 2702, -8.990842, 52.848558),
("Knox's Bridge", 27025, -8.972517, 52.848179),
("Tulla Road Bridge", 27026, -8.970894, 52.853952),
("Gaurus Bridge", 27028, -8.95047, 52.852006),
("Doora Bridge", 27060, -8.967172, 52.838778),
("Carrigaholt Pier", 27063, -9.698758, 52.600790),
("Clarecastle U/S", 27064, -8.962932, 52.817437),
("Clarecastle Barrage D/S", 27065, -8.962663, 52.817369),
("Ennis Bridge", 27066, -8.981789, 52.846674),
("Clarecastle Bridge", 27068, -8.961536, 52.815397),
("Shannon Airport", 27069, -8.917919, 52.678704),
("Ennistimon (New Location)", 28001, -9.293865, 52.938862),
("Doonbeg", 28002, -9.521525, 52.728373),
("Doonbeg Pier", 28060, -9.536603, 52.740479),
("Rathgorgin", 29001, -8.679536, 53.257849),
("Rahasane Turlough", 29002, -8.807869, 53.216947),
("Clarinbridge", 29004, -8.874783, 53.229342),
("Craughwell", 29007, -8.733144, 53.227709),
("Russaun", 29009, -8.772402, 53.053411),
("Aggard Bridge", 29010, -8.743236, 53.220817),
("Kilcolgan", 29011, -8.871329, 53.214109),
("Caherfinesker", 29014, -8.790411, 53.265585),
("Oranmore Bridge", 29015, -8.929459, 53.271762),
("Gortmackan", 29020, -8.642418, 53.183557),
("Ballycahalan", 29021, -8.752354, 53.099393),
("Kilcrimple", 29022, -8.739264, 53.045508),
("Killafeen", 29023, -8.76927, 53.023239),
("Poulataggle", 29024, -8.890516, 53.058588),
("Gort 1B", 29025, -8.890456, 53.058571),
("Cartronbower", 30001, -9.312859, 53.742967),
("Ower Bridge", 30002, -9.160029, 53.481485),
("Corrofin  Clare", 30004, -8.863839, 53.437557),
("Foxhill", 30005, -9.15424, 53.658173),
("Ballygaddy", 30007, -8.874377, 53.530827),
("Carrownagower", 30017, -9.289989, 53.573557),
("Cong Weir", 30031, -9.28876, 53.538599),
("Cregaree", 30034, -9.289413, 53.545681),
("Clooncormick", 30037, -9.117484, 53.653049),
("Wolfe Tone Bridge",30061, -9.055651, 53.269982),
("Caher Pier", 30081, -9.299349, 53.610987),
("Burriscara", 30082, -9.249069, 53.732912),
("Annaghdown Pier", 30083, -9.076388, 53.387032),
("Cong Pier", 30084, -9.275412, 53.530592),
("Angligham", 30089, -9.066216, 53.317909),
("Dangan", 30098, -9.075881, 53.295946),
("Galway Barrage", 30099, -9.056272, 53.277994),
("Oughterard", 30101, -9.321144, 53.431399),
("Rossaveel Pier", 31061, -9.562056, 53.266925),
("Shannagurraun", 31075, -9.308306, 53.274345),
("Derrinkee", 32013, -9.546362, 53.674948),
("Owenmore Bridge", 32014, -9.608129, 53.697039),
("Aasleagh Bridge", 32060, -9.671157, 53.617749),
("Rahans  Moy", 34001, -9.157683, 54.103927),
("Ballylahan", 34004, -9.102895, 53.93794),
("Scarrownageeragh", 34005, -9.060983, 53.922684),
("Ballycarroon", 34007, -9.344144, 54.08614),
("Curraughbonaun", 34009, -8.834877, 54.013081),
("Cloonacannana", 34010, -8.930502, 53.967417),
("Gneeve Bridge", 34011, -9.181784, 53.863947),
("Banada  Moy", 34013, -8.816832, 54.036519),
("Mill Bridge  Clydagh", 34014, -9.184442, 53.909069),
("Turlough", 34018, -9.208034, 53.885545),
("Enniscrone Pier", 34060, -9.098442, 54.220189),
("Ballina", 34061, -9.146781, 54.117184),
("Pollagh", 34071, -9.129626, 53.965284),
("Pontoon", 34081, -9.208076, 53.977665),
("Gortnaraby", 34082, -9.298326, 54.093363),
("Corryosla", 34083, -9.227889, 53.985389),
("Keenagh Deel Bridge", 34114, -9.514662, 54.081346),
("Mullenmore Spring", 34117, -9.309022, 54.090427),
("Richmond Bridge", 34118, -9.372725, 54.077808),
("Crossmolina", 34119, -9.319215, 54.100128),
("Crossmolina Weir U/S", 34120, -9.324923, 54.093316),
("Crossmolina Weir D/S", 34121, -9.323120, 54.093329),
("Knockglass House", 34122, -9.299753, 54.124702),
("Chapel View", 34123, -9.321135, 54.095990),
("Crossmolina U/S", 34124, -9.319344, 54.100031),
("Ballynacarrow", 35001, -8.54906, 54.14397),
("Billa Bridge", 35002, -8.553356, 54.179266),
("Ballygrania", 35003, -8.468385, 54.181719),
("Big Bridge", 35004, -8.511156, 54.059283),
("Ballysadare", 35005, -8.509294, 54.209196),
("Dromahair", 35011, -8.299487, 54.227059),
("New Bridge  Manorhamilton", 35028, -8.201256, 54.319812),
("Four Masters Bridge", 35029, -8.260686, 54.459445),
("Lareen", 35071, -8.240014, 54.452301),
("Templehouse Demesne", 35078, -8.582531, 54.111641),
("Ballynary", 35087, -8.309197, 54.061022),
("Butlers Bridge", 36010, -7.377088, 54.041879),
("Bellahillan", 36011, -7.457176, 53.962624),
("Sallaghan", 36012, -7.503405, 53.886691),
("Derreskit", 36013, -7.558754, 54.012758),
("Anlore", 36015, -7.177331, 54.177075),
("Ashfield", 36018, -7.121146, 54.072165),
("Belturbet", 36019, -7.450869, 54.097958),
("Killywillin", 36020, -7.691128, 54.080327),
("Kiltybardan", 36021, -7.860888, 54.057276),
("Aghacashlaun", 36022, -7.926414, 54.047317),
("Derryheen Bridge", 36023, -7.401796, 54.037962),
("Bellaheady", 36027, -7.618702, 54.089336),
("Aghoo", 36028, -7.797111, 54.028395),
("Tomkinroad", 36029, -7.519247, 54.107873),
("Kilconny", 36036, -7.449005, 54.102523),
("Urney Bridge", 36037, -7.405639, 54.048915),
("Gowly", 36071, -7.958086, 54.022519),
("Wood Island", 36073, -7.86659, 54.050742),
("Killykeen Forest Park", 36083, -7.470649, 54.007865),
("Innisconnell Pier", 36084, -7.458314, 54.016974),
("Ballinacur", 36091, -7.693641, 54.056571),
("Foalies Bridge", 36171, -7.436735, 54.138890),
("Clonconwal", 38001, -8.365278, 54.781706),
("Glenties", 38010, -8.285969, 54.793390),
("New Mills", 39001, -7.817926, 54.931113),
("Tullyarvan", 39003, -7.452395, 55.143612),
("Gartan Bridge", 39008, -7.894476, 55.000253),
("Aghawoney", 39009, -7.720693, 55.043786),
("Port Bridge", 39061, -7.714146, 54.948447),
("Ballyloskey", 40008, -7.262685, 55.247115),
("Malin Head", 40060, -7.334645, 55.371618);

