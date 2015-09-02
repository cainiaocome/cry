package key

import "math/big"
import "crypto/rsa"

func fromBase10(base10 string) *big.Int {
	i, ok := new(big.Int).SetString(base10, 10)
	if !ok {
		panic("bad number: " + base10)
	}
	return i
}

func Get_private_key() *rsa.PrivateKey {
	var rsaPrivateKey = &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: fromBase10("742596489872619689003227781924897539868688609047468702525209250059018349118389298815846328379526657777284569882574931670408865793671926101906406357244415325669747810879914345861694120238158467002148817916598134538837664117756594776702980408254468991776478873073136119188995730278433997827284476357432491783744573279510530354016204234535067958977860061231889963337113890489078916930189594835718697512928758389069267708543662349987349308431252604186707091056865363506381722331830583733762118286444436263378179708963233006140723868711207752555530406130706858354823387819554535420546871257482659815943216407817182626673809341543144645605081275153202081559629458114460205553528275914528019975088717791642618205795352667516686827772600851486845474664342166536819279711136286321009702755762003640560120544116540922111777888643336463777723953698276867736158632981875702775129549954661941864198129745337811195089485381995135370572547095077914053279054418857926585146770387171919288872311873783165658141850343157061723281167760496486415803583583427735628673709040709161593987196945345645948450820761218094839037013667707250613315779810515698324434816002291888995992817612562912535212784628864514844499568630806488437049862298347561269798505899"),
			E: 65537,
		},
		D: fromBase10("492046516818263118466899101730147484150904112911581677634881237832260735316936315380290931990798253108604032030010778762950775853185275355528719594478519546473118981452618833163618523449990561508282533805762759850909037098334988299408378842920034727984094405651783511234602432020705805203927985501648014948946520204811706068681106380894537835400057572958707625584237464249023482425095490421305284640720987122470252685406847087114159081413966833037944297559308152803219956547595146232488822902922771026095144023402480938884308619539865978832169734443534878374936381220747908824591419875126791011296430604230604934057216512722794489604777200820133187131970111994662305189887068982847433109068282160111689904893940590468800058780843602656852323578616344073357964611268126723405228156788388572359269656953929038959905961730345894442088178804963635430099148297467469517939423655041491126967248441474736643765948594772352025675315932130816055241073905921708094942132627718183124201315799847499762727862163164166618065789039714455003022500650953257790461197383027949781227156711711017081403064063061434837702137699438976743327779825292724936585830454398484497716400155771826555339171393129403267344752954465432764342926482120449842166975073"),
		Primes: []*big.Int{
			fromBase10("29144735541159982576210737943486578746927653929409341948807760455888500319655846583613526543759088936328164718708197479626221436506207959999687339017846532863226430783435047979896674290460682541332811798357763084942198203064486291032950459407640780625200943860282895930890096604944789050884872902247039088835300777511779840347327239276224087069066647242497194991289735598592857944457216003395604550609069227203007801323683821230640784033379721756907520077695057777812908826145820387306653824264865599485528547127235675706278137348672750242691083829312967772701001294288593602799678214991334323385709077444191813743963"),
			fromBase10("25479609819203176257893197087755946503135416335436505658251345117100417336522902405237925146028017784647679695915746987711349246908397459130221578136189446826765611992025774693995763716934645412348196846782295763567409511541643962514102051054692972116415606530980695658084528204833811392913161813771815497916048898967804530606675443843604163384774231837012499850611114795805643240311104213476826240264938687248879014840940836543952199782424367993388619425729922045965197841788284700239628734446824419069797460044552132458518970529549497866655785870774076196254598404499090690135115474604366359352286248173006456908273"),
		},
	}

    return rsaPrivateKey
}