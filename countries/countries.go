//////////////////////////////////////////////////////////////////////
// countries.go
//
// @usage
// 
//     1. Import this package.
//
//         --------------------------------------------------
//         import myCountries "countries"
//         --------------------------------------------------
//
//     2. Open  a database using a connector.
//
//         --------------------------------------------------
//         db, err := sql.Open("driver-name", "database=test1")
//         if err != nil {
//             // Error Handling
//         }
//         --------------------------------------------------
//
//     3. Initilize this package.
//
//         --------------------------------------------------
//         myCountries.Init(db)
//         --------------------------------------------------
//
//     4. Now, you can get contents as you want.
//
//         4-1. When you would like to get all countries which status is active.
//
//             --------------------------------------------------
//             langCode := "en"
//             allCountries := myCountries.GetOnlyActive(langCode)
//             --------------------------------------------------
//
//         4-2. When you would like to get countries in Africa which status is active.
//         
//             --------------------------------------------------
//             langCode := "en"
//             allCountries := myCountries.GetAfricaOnlyActive(langCode)
//             --------------------------------------------------
//
//
// MIT License
//
// Copyright (c) 2019 noknow.info
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A 
// PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTW//ARE.
//////////////////////////////////////////////////////////////////////
package countries

import (
    "database/sql"
    "log"
    "strconv"
    _ "github.com/go-sql-driver/mysql"
)

const (
    TABLE_NAME = "countries"
)

var (
    db *sql.DB
)

type Columns struct {
    CountryCode string
    Ar string
    De string
    En string
    Es string
    Fr string
    Ja string
    Pt string
    Ru string
    ZhCn string
    ZhTw string
    Continent int
    Status int
    Name string
}

func Init(mydb *sql.DB) {
    db = mydb
    query := "CREATE TABLE IF NOT EXISTS " + TABLE_NAME +
            "(country_code VARCHAR(2) NOT NULL COMMENT 'Country code of 2 digits'," +
            "ar VARCHAR(255) NOT NULL COMMENT 'Arabic'," +
            "de VARCHAR(255) NOT NULL COMMENT 'German'," +
            "en VARCHAR(255) NOT NULL COMMENT 'English'," +
            "es VARCHAR(255) NOT NULL COMMENT 'Spanish'," +
            "fr VARCHAR(255) NOT NULL COMMENT 'French'," +
            "ja VARCHAR(255) NOT NULL COMMENT 'Japanese'," +
            "pt VARCHAR(255) NOT NULL COMMENT 'Portuguese'," +
            "ru VARCHAR(255) NOT NULL COMMENT 'Russian'," +
            "zh_cn VARCHAR(255) NOT NULL COMMENT 'Chinese (Simplified Chinese)'," +
            "zh_tw VARCHAR(255) NOT NULL COMMENT 'Chinese (Traditional Chinese)'," +
            "continent TINYINT(1) UNSIGNED NOT NULL COMMENT '1: Africa, 2: Asia, 3: Europe, 4: North America, 5: South America, 6: Australia / Oceania, 7: Antarctica'," +
            "status TINYINT(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '0 inactive, 1: active'," +
            "PRIMARY KEY(country_code)" +
            ") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='countriess table'"
    _, err := db.Exec(query)
    if err != nil {
        db.Close()
        log.Fatalf("[FATAL] db.Exec() error: %s\n", err)
        return
    }

    insertQuery := "INSERT IGNORE INTO " + TABLE_NAME +
            " (country_code,ar,de,en,es,fr,ja,pt,ru,zh_cn,zh_tw,continent,status)" +
            " VALUES" +
            " ('AD','أندورا','Andorra','Andorra','Andorra','Andorre','アンドラ','Andorra','Андорра','安道尔','安道尔',3,1)," +
            " ('AE','الإمارات العربية المتحدة','Vereinigte Arabische Emirate','United Arab Emirates','Emiratos Árabes Unidos','Émirats arabes unis','アラブ首長国連邦','Emirados Árabes Unidos','ОАЭ','阿联酋','阿联酋',2,1)," +
            " ('AF','أفغانستان','Afghanistan','Afghanistan','Afganistán','Afghanistan','アフガニスタン','Afeganistão','Афганистан','阿富汗','阿富汗',2,1)," +
            " ('AG','أنتيغوا وباربودا','Antigua und Barbuda','Antigua and Barbuda','Antigua y Barbuda','Antigua-et-Barbuda','アンティグア・バーブーダ','Antígua e Barbuda','Антигуа и Барбуда','安地卡及巴布達','安地卡及巴布達',4,1)," +
            " ('AI','أنغويلا','Anguilla','Anguilla','Anguila','Anguilla','アンギラ','Anguilla','Ангилья','安圭拉','安圭拉',4,1)," +
            " ('AL','ألبانيا','Albanien','Albania','Albania','Albanie','アルバニア','Albânia','Албания','阿尔巴尼亚','阿尔巴尼亚',3,1)," +
            " ('AM','أرمينيا','Armenien','Armenia','Armenia','Arménie','アルメニア','Armênia','Армения','亞美尼亞','亞美尼亞',3,1)," +
            " ('AO','أنغولا','Angola','Angola','Angola','Angola','アンゴラ','Angola','Ангола','安哥拉','安哥拉',1,1)," +
            " ('AQ','القارة القطبية الجنوبية','Antarktika','Antarctica','Antártida','Antarctique','南極','Antártica','Антарктида','南极洲','南极洲',7,1)," +
            " ('AR','الأرجنتين','Argentinien','Argentina','Argentina','Argentine','アルゼンチン','Argentina','Аргентина','阿根廷','阿根廷',5,1)," +
            " ('AS','ساموا الأمريكية','Amerikanisch-Samoa','American Samoa','Samoa Americana','Samoa américaines','アメリカ領サモア','Samoa Americana','Американское Самоа','美属萨摩亚','美属萨摩亚',6,1)," +
            " ('AT','النمسا','Österreich','Austria','Austria','Autriche','オーストリア','Áustria','Австрия','奥地利','奥地利',3,1)," +
            " ('AU','أستراليا','Australien','Australia','Australia','Australie','オーストラリア','Austrália','Австралия','澳大利亚','澳大利亚',5,1)," +
            " ('AW','أروبا','Aruba','Aruba','Aruba','Aruba','アルバ','Aruba','Аруба','阿鲁巴','阿鲁巴',4,1)," +
            " ('AX','جزر أولاند','Åland','Åland Islands','Åland','les Åland','オーランド諸島','Ilhas Aland','Аландские острова','奥兰','奥兰',1,1)," +
            " ('AZ','أذربيجان','Aserbaidschan','Azerbaijan','Azerbaiyán','Azerbaïdjan','アゼルバイジャン','Azerbaijão','Азербайджан','阿塞拜疆','阿塞拜疆',3,1)," +
            " ('BA','البوسنة والهرسك','Bosnien und Herzegowina','Bosnia and Herzegovina','Bosnia y Herzegovina','Bosnie-Herzégovine','ボスニア・ヘルツェゴビナ','Bósnia e Herzegovina','Босния и Герцеговина','波斯尼亚和黑塞哥维那','波斯尼亚和黑塞哥维那',3,1)," +
            " ('BB','باربادوس','Barbados','Barbados','Barbados','Barbade','バルバドス','Barbados','Барбадос','巴巴多斯','巴巴多斯',4,1)," +
            " ('BD','بنغلاديش','Bangladesch','Bangladesh','Bangladés','Bangladesh','バングラデシュ','Bangladesh','Бангладеш','孟加拉国','孟加拉国',2,1)," +
            " ('BE','بلجيكا','Belgien','Belgium','Bélgica','Belgique','ベルギー','Bélgica','Бельгия','比利時','比利時',3,1)," +
            " ('BF','بوركينا فاسو','Burkina Faso','Burkina Faso','Burkina Faso','Burkina Faso','ブルキナファソ','Burkina Faso','Буркина-Фасо','布吉納法索','布吉納法索',1,1)," +
            " ('BG','بلغاريا','Bulgarien','Bulgaria','Bulgaria','Bulgarie','ブルガリア','Bulgária','Болгария','保加利亚','保加利亚',3,1)," +
            " ('BH','البحرين','Bahrain','Bahrain','Baréin','Bahreïn','バーレーン','Barém','Бахрейн','巴林','巴林',2,1)," +
            " ('BI','بوروندي','Burundi','Burundi','Burundi','Burundi','ブルンジ','Burundi','Бурунди','布隆迪','布隆迪',1,1)," +
            " ('BJ','بنين','Benin','Benin','Benín','Bénin','ベナン','Benin','Бенин','贝宁','贝宁',1,1)," +
            " ('BL','سان بارتيلمي','Saint-Barthélemy','Saint Barthélemy','San Bartolomé','Saint-Barthélemy','サン・バルテルミー','São Bartolomeu','Сен-Бартелеми','圣巴泰勒米','圣巴泰勒米',4,1)," +
            " ('BM','برمودا','Bermuda','Bermuda','Bermudas','Bermudes','バミューダ','Bermudas','Бермуды','百慕大','百慕大',4,1)," +
            " ('BN','بروناي','Brunei Darussalam','Brunei Darussalam','Brunéi','Brunei','ブルネイ・ダルサラーム','Brunei Darussalam','Бруней','文莱','文莱',2,1)," +
            " ('BO','بوليفيا','Bolivien','Bolivia','Bolivia','Bolivie','ボリビア多民族国','Bolívia','Боливия','玻利维亚','玻利维亚',5,1)," +
            " ('BQ','الجزر الكاريبية الهولندية','Bonaire','Bonaire','Bonaire','Pays-Bas caribéens','ボネール','Bonaire','Бонэйр, Синт-Эстатиус и Саба','荷兰加勒比区','荷兰加勒比区',4,1)," +
            " ('BR','البرازيل','Brasilien','Brazil','Brasil','Brésil','ブラジル','Brasil','Бразилия','巴西','巴西',5,1)," +
            " ('BS','باهاماس','Bahamas','Bahamas','Bahamas','Bahamas','バハマ','Bahamas','Багамские Острова','巴哈马','巴哈马',4,1)," +
            " ('BT','بوتان','Bhutan','Bhutan','Bután','Bhoutan','ブータン','Butão','Бутан','不丹','不丹',2,1)," +
            " ('BV','جزيرة بوفيه','Bouvetinsel','Bouvet Island','Isla Bouvet','Île Bouvet','ブーベ島','Ilha Bouvet','Остров Буве','布韦岛','布韦岛',7,1)," +
            " ('BW','بوتسوانا','Botswana','Botswana','Botsuana','Botswana','ボツワナ','Botsuana','Ботсвана','博茨瓦纳','博茨瓦纳',1,1)," +
            " ('BY','روسيا البيضاء','Belarus','Belarus','Bielorrusia','Biélorussie','ベラルーシ','Bielorrússia','Белоруссия','白俄羅斯','白俄羅斯',3,1)," +
            " ('BZ','بليز','Belize','Belize','Belice','Belize','ベリーズ','Belize','Белиз','伯利兹','伯利兹',4,1)," +
            " ('CA','كندا','Kanada','Canada','Canadá','Canada','カナダ','Canadá','Канада','加拿大','加拿大',4,1)," +
            " ('CC','جزر كوكوس','Kokosinseln','Cocos (Keeling) Islands','Islas Cocos','Îles Cocos','ココス (キーリング) 諸島','Ilhas Cocos','Кокосовые острова','科科斯 (基林) 群島','科科斯 (基林) 群島',1,1)," +
            " ('CD','جمهورية الكونغو الديمقراطية','Kongo, Demokratische Republik','Congo, Democratic Republic of the','República Democrática del Congo','République démocratique du Congo','コンゴ民主共和国','Congo, República Democrática do','ДР Конго','刚果 (金)','刚果 (金)',1,1)," +
            " ('CF','جمهورية أفريقيا الوسطى','Zentralafrikanische Republik','Central African Republic','República Centroafricana','République centrafricaine','中央アフリカ共和国','República Centro-Africana','ЦАР','中非','中非',1,1)," +
            " ('CG','جمهورية الكونغو','Kongo, Republik','Congo','República del Congo','République du Congo','コンゴ共和国','Congo','Республика Конго','刚果 (布)','',1,1)," +
            " ('CH','سويسرا','Schweiz','Switzerland','Suiza','Suisse','スイス','Suíço','Швейцария','瑞士','瑞士',3,1)," +
            " ('CI','ساحل العاج','Côte d''Ivoire','Côte d''Ivoire','Costa de Marfil','Côte d''Ivoire','コートジボワール','Costa do Marfim','Кот-д’Ивуар','科特迪瓦','科特迪瓦',1,1)," +
            " ('CK','جزر كوك','Cookinseln','Cook Islands','Islas Cook','Îles Cook','クック諸島','Ilhas Cook','Острова Кука','庫克群島','庫克群島',6,1)," +
            " ('CL','تشيلي','Chile','Chile','Chile','Chili','チリ','Chile','Чили','智利','智利',5,1)," +
            " ('CM','الكاميرون','Kamerun','Cameroon','Camerún','Cameroun','カメルーン','Camarões','Камерун','喀麦隆','喀麦隆',1,1)," +
            " ('CN','الصين','China','China','China','Chine','中華人民共和国','China','Китай','中国','中国',2,1)," +
            " ('CO','كولومبيا','Kolumbien','Colombia','Colombia','Colombie','コロンビア','Colômbia','Колумбия','哥伦比亚','哥伦比亚',5,1)," +
            " ('CR','كوستاريكا','Costa Rica','Costa Rica','Costa Rica','Costa Rica','コスタリカ','Costa Rica','Коста-Рика','哥斯达黎加','哥斯达黎加',4,1)," +
            " ('CU','كوبا','Kuba','Cuba','Cuba','Cuba','キューバ','Cuba','Куба','古巴','古巴',4,1)," +
            " ('CV','الرأس الأخضر','Kap Verde','Cabo Verde','Cabo Verde','Cap-Vert','カーボベルデ','Cabo Verde','Кабо-Верде','佛得角','佛得角',1,1)," +
            " ('CW','كوراساو','Curaçao','Curaçao','Curazao','Curaçao','キュラソー','Curaçao','Кюрасао','库拉索','库拉索',4,1)," +
            " ('CX','جزيرة عيد الميلاد','Weihnachtsinsel','Christmas Island','Isla de Navidad','Île Christmas','クリスマス島','Ilha do Natal','Остров Рождества','圣诞岛','圣诞岛',1,1)," +
            " ('CY','قبرص','Zypern','Cyprus','Chipre','Chypre','キプロス','Chipre','Кипр','賽普勒斯','賽普勒斯',3,1)," +
            " ('CZ','جمهورية التشيك','Tschechien','Czech','República Checa','Tchéquie','チェコ','Tcheca','Чехия','捷克','捷克',3,1)," +
            " ('DE','ألمانيا','Deutschland','Germany','Alemania','Allemagne','ドイツ','Alemanha','Германия','德國','德國',3,1)," +
            " ('DJ','جيبوتي','Dschibuti','Djibouti','Yibuti','Djibouti','ジブチ','Djibuti','Джибути','吉布提','吉布提',1,1)," +
            " ('DK','الدنمارك','Dänemark','Denmark','Dinamarca','Danemark','デンマーク','Dinamarca','Дания','丹麥','丹麥',3,1)," +
            " ('DM','دومينيكا','Dominica','Dominica','Dominica','Dominique','ドミニカ国','Dominica','Доминика','多米尼克','多米尼克',4,1)," +
            " ('DO','جمهورية الدومينيكان','Dominikanische Republik','Dominican Republic','República Dominicana','République dominicaine','ドミニカ共和国','República Dominicana','Доминиканская Республика','多米尼加','多米尼加',4,1)," +
            " ('DZ','الجزائر','Algerien','Algeria','Argelia','Algérie','アルジェリア','Argélia','Алжир','阿尔及利亚','阿尔及利亚',1,1)," +
            " ('EC','الإكوادور','Ecuador','Ecuador','Ecuador','Équateur','エクアドル','Equador','Эквадор','厄瓜多尔','厄瓜多尔',5,1)," +
            " ('EE','إستونيا','Estland','Estonia','Estonia','Estonie','エストニア','Estônia','Эстония','爱沙尼亚','爱沙尼亚',3,1)," +
            " ('EG','مصر','Ägypten','Egypt','Egipto','Égypte','エジプト','Egito','Египет','埃及','埃及',2,1)," +
            " ('EH','الصحراء الغربية','Westsahara','Western Sahara','República Árabe Saharaui Democrática','République arabe sahraouie démocratique','西サハラ','Saara Ocidental','САДР','西撒哈拉','西撒哈拉',1,1)," +
            " ('ER','إريتريا','Eritrea','Eritrea','Eritrea','Érythrée','エリトリア','Eritreia','Эритрея','厄立特里亚','厄立特里亚',1,1)," +
            " ('ES','إسبانيا','Spanien','Spain','España','Espagne','スペイン','Espanha','Испания','西班牙','西班牙',3,1)," +
            " ('ET','إثيوبيا','Äthiopien','Ethiopia','Etiopía','Éthiopie','エチオピア','Etiópia','Эфиопия','衣索比亞','衣索比亞',1,1)," +
            " ('FI','فنلندا','Finnland','Finland','Finlandia','Finlande','フィンランド','Finlândia','Финляндия','芬兰','芬兰',3,1)," +
            " ('FJ','فيجي','Fidschi','Fiji','Fiyi','Fidji','フィジー','Fiji','Фиджи','斐济','斐济',6,1)," +
            " ('FK','جزر فوكلاند','Falklandinseln','Falkland Islands (Malvinas)','Islas Malvinas','Malouines','フォークランド (マルビナス) 諸島','Ilhas Falkland (Malvinas)','Фолклендские острова','福克蘭群島','福克蘭群島',5,1)," +
            " ('FM','ولايات ميكرونيسيا المتحدة','Mikronesien','Micronesia (Federated States of)','Micronesia','États fédérés de Micronésie','ミクロネシア連邦','Micronésia (Estados Federados da)','Микронезия','密克羅尼西亞聯邦','密克羅尼西亞聯邦',6,1)," +
            " ('FO','جزر فارو','Färöer','Faroe Islands','Islas Feroe','Îles Féroé','フェロー諸島','ilhas Faroe','Фареры','法罗群岛','法罗群岛',1,1)," +
            " ('FR','فرنسا','Frankreich','France','Francia','France','フランス','França','Франция','法国','法国',3,1)," +
            " ('GA','الغابون','Gabun','Gabon','Gabón','Gabon','ガボン','Gabão','Габон','加彭','加彭',1,1)," +
            " ('GB','المملكة المتحدة','Vereinigtes Königreich Großbritannien und Nordirland','United Kingdom','Reino Unido','Royaume-Uni','イギリス','Reino Unido','Великобритания','英國','英國',3,1)," +
            " ('GD','غرينادا','Grenada','Grenada','Granada','Grenade','グレナダ','Granada','Гренада','格瑞那達','格瑞那達',4,1)," +
            " ('GE','جورجيا','Georgien','Georgia','Georgia','Géorgie','ジョージア','Geórgia','Грузия','格鲁吉亚','格鲁吉亚',3,1)," +
            " ('GF','غويانا الفرنسية','Französisch-Guayana','French Guiana','Guayana Francesa','Guyane','フランス領ギアナ','Guiana Francesa','Гвиана','法属圭亚那','法属圭亚那',5,1)," +
            " ('GG','غيرنزي','Guernsey','Guernsey','Guernsey','Guernesey','ガーンジー','Guernsey','Гернси','根西','根西',1,1)," +
            " ('GH','غانا','Ghana','Ghana','Ghana','Ghana','ガーナ','Gana','Гана','加纳','加纳',1,1)," +
            " ('GI','جبل طارق','Gibraltar','Gibraltar','Gibraltar','Gibraltar','ジブラルタル','Gibraltar','Гибралтар','直布罗陀','直布罗陀',1,1)," +
            " ('GL','جرينلاند','Grönland','Greenland','Groenlandia','Groenland','グリーンランド','Gronelândia','Гренландия','格陵兰','格陵兰',4,1)," +
            " ('GM','غامبيا','Gambia','Gambia','Gambia','Gambie','ガンビア','Gâmbia','Гамбия','冈比亚','冈比亚',1,1)," +
            " ('GN','غينيا','Guinea','Guinea','Guinea','Guinée','ギニア','Guiné','Гвинея','几内亚','几内亚',1,1)," +
            " ('GP','غوادلوب','Guadeloupe','Guadeloupe','Guadalupe','Guadeloupe','グアドループ','Guadalupe','Гваделупа','瓜德罗普','瓜德罗普',4,1)," +
            " ('GQ','غينيا الاستوائية','Äquatorialguinea','Equatorial Guinea','Guinea Ecuatorial','Guinée équatoriale','赤道ギニア','Guiné Equatorial','Экваториальная Гвинея','赤道几内亚','赤道几内亚',1,1)," +
            " ('GR','اليونان','Griechenland','Greece','Grecia','Grèce','ギリシャ','Grécia','Греция','希臘','希臘',3,1)," +
            " ('GS','جورجيا الجنوبية وجزر ساندويتش الجنوبية','Südgeorgien und die Südlichen Sandwichinseln','South Georgia and the South Sandwich Islands','Islas Georgias del Sur y Sandwich del Sur','Géorgie du Sud-et-les îles Sandwich du Sud','サウスジョージア・サウスサンドウィッチ諸島','Ilhas Geórgia do Sul e Sandwich do Sul','Южная Георгия и Южные Сандвичевы Острова','南乔治亚和南桑威奇群岛','南乔治亚和南桑威奇群岛',5,1)," +
            " ('GT','غواتيمالا','Guatemala','Guatemala','Guatemala','Guatemala','グアテマラ','Guatemala','Гватемала','危地马拉','危地马拉',4,1)," +
            " ('GU','غوام','Guam','Guam','Guam','Guam','グアム','Guam','Гуам','關島','關島',6,1)," +
            " ('GW','غينيا بيساو','Guinea-Bissau','Guinea-Bissau','Guinea-Bisáu','Guinée-Bissau','ギニアビサウ','Guiné-Bissau','Гвинея-Бисау','几内亚比绍','几内亚比绍',1,1)," +
            " ('GY','غيانا','Guyana','Guyana','Guyana','Guyana','ガイアナ','Guiana','Гайана','圭亚那','圭亚那',5,1)," +
            " ('HK','هونغ كونغ','Hongkong','Hong Kong','Hong Kong','Hong Kong','香港','Hong Kong','Гонконг','香港','香港',2,1)," +
            " ('HM','جزيرة هيرد وجزر ماكدونالد','Heard und McDonaldinseln','Heard Island and McDonald Islands','Islas Heard y McDonald','Îles Heard-et-MacDonald','ハード島とマクドナルド諸島','Ilha Heard e Ilhas McDonald','Херд и Макдональд','赫德岛和麦克唐纳群岛','赫德岛和麦克唐纳群岛',1,1)," +
            " ('HN','هندوراس','Honduras','Honduras','Honduras','Honduras','ホンジュラス','Honduras','Гондурас','洪都拉斯','洪都拉斯',4,1)," +
            " ('HR','كرواتيا','Kroatien','Croatia','Croacia','Croatie','クロアチア','Croácia','Хорватия','克罗地亚','克罗地亚',3,1)," +
            " ('HT','هايتي','Haiti','Haiti','Haití','Haïti','ハイチ','Haiti','Гаити','海地','海地',4,1)," +
            " ('HU','المجر','Ungarn','Hungary','Hungría','Hongrie','ハンガリー','Hungria','Венгрия','匈牙利','匈牙利',3,1)," +
            " ('ID','إندونيسيا','Indonesien','Indonesia','Indonesia','Indonésie','インドネシア','Indonésia','Индонезия','印尼','印尼',2,1)," +
            " ('IE','أيرلندا','Irland','Ireland','Irlanda','Irlande','アイルランド','Irlanda','Ирландия','爱尔兰','爱尔兰',3,1)," +
            " ('IL','إسرائيل','Israel','Israel','Israel','Israël','イスラエル','Israel','Израиль','以色列','以色列',2,1)," +
            " ('IM','جزيرة مان','Insel Man','Isle of Man','Isla de Man','Île de Man','マン島','Ilha de Man','Остров Мэн','马恩岛','马恩岛',1,1)," +
            " ('IN','الهند','Indien','India','India','Inde','インド','Índia','Индия','印度','印度',2,1)," +
            " ('IO','إقليم المحيط الهندي البريطاني','Britisches Territorium im Indischen Ozean','British Indian Ocean Territory','Territorio Británico del Océano Índico','Territoire britannique de l''océan Indien','イギリス領インド洋地域','Território Britânico do Oceano Índico','Британская территория в Индийском океане','英屬印度洋領地','英屬印度洋領地',1,1)," +
            " ('IQ','العراق','Irak','Iraq','Irak','Irak','イラク','Iraque','Ирак','伊拉克','伊拉克',2,1)," +
            " ('IR','إيران','Iran, Islamische Republik','Iran (Islamic Republic of)','Irán','Iran','イラン・イスラム共和国','Irã (Republic Islâmica do Irã)','Иран','伊朗','伊朗',2,1)," +
            " ('IS','آيسلندا','Island','Iceland','Islandia','Islande','アイスランド','Islândia','Исландия','冰島','冰島',3,1)," +
            " ('IT','إيطاليا','Italien','Italy','Italia','Italie','イタリア','Itália','Италия','義大利','義大利',3,1)," +
            " ('JE','جيرزي','Jersey','Jersey','Jersey','Jersey','ジャージー','Jersey','Джерси','澤西','澤西',1,1)," +
            " ('JM','جامايكا','Jamaika','Jamaica','Jamaica','Jamaïque','ジャマイカ','Jamaica','Ямайка','牙买加','牙买加',4,1)," +
            " ('JO','الأردن','Jordanien','Jordan','Jordania','Jordanie','ヨルダン','Jordânia','Иордания','约旦','约旦',2,1)," +
            " ('JP','اليابان','Japan','Japan','Japón','Japon','日本','Japão','Япония','日本','日本',2,1)," +
            " ('KE','كينيا','Kenia','Kenya','Kenia','Kenya','ケニア','Quênia','Кения','肯尼亚','肯尼亚',1,1)," +
            " ('KG','قيرغيزستان','Kirgisistan','Kyrgyzstan','Kirguistán','Kirghizistan','キルギス','Quirguistão','Киргизия','吉尔吉斯斯坦','吉尔吉斯斯坦',2,1)," +
            " ('KH','كمبوديا','Kambodscha','Cambodia','Camboya','Cambodge','カンボジア','Camboja','Камбоджа','柬埔寨','柬埔寨',2,1)," +
            " ('KI','كيريباتي','Kiribati','Kiribati','Kiribati','Kiribati','キリバス','Kiribati','Кирибати','基里巴斯','基里巴斯',6,1)," +
            " ('KM','جزر القمر','Komoren','Comoros','Comoras','Comores','コモロ','Comores','Коморы','科摩罗','科摩罗',1,1)," +
            " ('KN','سانت كيتس ونيفيس','St. Kitts und Nevis','Saint Kitts and Nevis','San Cristóbal y Nieves','Saint-Christophe-et-Niévès','セントクリストファー・ネイビス','São Cristóvão e Nevis','Сент-Китс и Невис','圣基茨和尼维斯','圣基茨和尼维斯',4,1)," +
            " ('KP','كوريا الشمالية','Nordkorea','North Korea','Corea del Norte','Corée du Nord','朝鮮民主主義人民共和国','Coreia do Norte','КНДР (Корейская Народно-Демократическая Республика)','朝鲜','朝鲜',2,1)," +
            " ('KR','كوريا الجنوبية','Südkorea','South Korea','Corea del Sur','Corée du Sud','大韓民国','Coreia do Sul','Республика Корея','韩国','韩国',2,1)," +
            " ('KW','الكويت','Kuwait','Kuwait','Kuwait','Koweït','クウェート','Kuwait','Кувейт','科威特','科威特',2,1)," +
            " ('KY','جزر كايمان','Kaimaninseln','Cayman Islands','Islas Caimán','Îles Caïmans','ケイマン諸島','Ilhas Cayman','Острова Кайман','开曼群岛','开曼群岛',4,1)," +
            " ('KZ','كازاخستان','Kasachstan','Kazakhstan','Kazajistán','Kazakhstan','カザフスタン','Cazaquistão','Казахстан','哈萨克斯坦','哈萨克斯坦',3,1)," +
            " ('LA','لاوس','Laos','Laos','Laos','Laos','ラオス人民民主共和国','Laos','Лаос','老挝','老挝',2,1)," +
            " ('LB','لبنان','Libanon','Lebanon','Líbano','Liban','レバノン','Líbano','Ливан','黎巴嫩','黎巴嫩',2,1)," +
            " ('LC','سانت لوسيا','St. Lucia','Saint Lucia','Santa Lucía','Sainte-Lucie','セントルシア','Santa Lúcia','Сент-Люсия','圣卢西亚','圣卢西亚',4,1)," +
            " ('LI','ليختنشتاين','Liechtenstein','Liechtenstein','Liechtenstein','Liechtenstein','リヒテンシュタイン','Liechtenstein','Лихтенштейн','列支敦斯登','列支敦斯登',1,1)," +
            " ('LK','سريلانكا','Sri Lanka','Sri Lanka','Sri Lanka','Sri Lanka','スリランカ','Sri Lanka','Шри-Ланка','斯里蘭卡','斯里蘭卡',2,1)," +
            " ('LR','ليبيريا','Liberia','Liberia','Liberia','Liberia','リベリア','Libéria','Либерия','利比里亚','利比里亚',1,1)," +
            " ('LS','ليسوتو','Lesotho','Lesotho','Lesoto','Lesotho','レソト','Lesoto','Лесото','賴索托','賴索托',1,1)," +
            " ('LT','ليتوانيا','Litauen','Lithuania','Lituania','Lituanie','リトアニア','Lituânia','Литва','立陶宛','立陶宛',3,1)," +
            " ('LU','لوكسمبورغ','Luxemburg','Luxembourg','Luxemburgo','Luxembourg','ルクセンブルク','Luxemburgo','Люксембург','卢森堡','卢森堡',3,1)," +
            " ('LV','لاتفيا','Lettland','Latvia','Letonia','Lettonie','ラトビア','Letônia','Латвия','拉脫維亞','拉脫維亞',3,1)," +
            " ('LY','ليبيا','Libyen','Libya','Libia','Libye','リビア','Líbia','Ливия','利比亞','利比亞',1,1)," +
            " ('MA','المغرب','Marokko','Morocco','Marruecos','Maroc','モロッコ','Marrocos','Марокко','摩洛哥','摩洛哥',1,1)," +
            " ('MC','موناكو','Monaco','Monaco','Mónaco','Monaco','モナコ','Mônaco','Монако','摩納哥','摩納哥',3,1)," +
            " ('MD','مولدوفا','Moldawien','Moldova, Republic of','Moldavia','Moldavie','モルドバ共和国','Moldávia, República da','Молдавия','摩尔多瓦','摩尔多瓦',3,1)," +
            " ('ME','الجبل الأسود','Montenegro','Montenegro','Montenegro','Monténégro','モンテネグロ','Montenegro','Черногория','蒙特內哥羅','蒙特內哥羅',3,1)," +
            " ('MF','تجمع سان مارتين','Saint-Martin','Saint Martin (French part)','San Martín','Saint-Martin','サン・マルタン (フランス領)','São Martinho (parte francesa)','Сен-Мартен','法属圣马丁','法属圣马丁',4,1)," +
            " ('MG','مدغشقر','Madagaskar','Madagascar','Madagascar','Madagascar','マダガスカル','Madagáscar','Мадагаскар','马达加斯加','马达加斯加',1,1)," +
            " ('MH','جزر مارشال','Marshallinseln','Marshall Islands','Islas Marshall','Îles Marshall','マーシャル諸島','Ilhas Marshall','Маршалловы Острова','马绍尔群岛','马绍尔群岛',6,1)," +
            " ('MK','مقدونيا','Nordmazedonien','North Macedonia','Macedonia del Norte','Macédoine du Nord','北マケドニア','Macedônia do Norte','Северная Македония','北馬其頓','北馬其頓',3,1)," +
            " ('ML','مالي','Mali','Mali','Malí','Mali','マリ','Mali','Мали','马里','马里',1,1)," +
            " ('MM','ميانمار','Myanmar','Myanmar','Birmania','Birmanie','ミャンマー','Myanmar','Мьянма','緬甸','緬甸',2,1)," +
            " ('MN','منغوليا','Mongolei','Mongolia','Mongolia','Mongolie','モンゴル','Mongólia','Монголия','蒙古國','蒙古國',2,1)," +
            " ('MO','ماكاو','Macau','Macao','Macao','Macao','マカオ','Macau','Макао','澳門','澳門',1,1)," +
            " ('MP','جزر ماريانا الشمالية','Nördliche Marianen','Northern Mariana Islands','Islas Marianas del Norte','Îles Mariannes du Nord','北マリアナ諸島','Ilhas Marianas do Norte','Северные Марианские Острова','北馬里亞納群島','北馬里亞納群島',6,1)," +
            " ('MQ','مارتينيك','Martinique','Martinique','Martinica','Martinique','マルティニーク','Martinica','Мартиника','马提尼克','马提尼克',4,1)," +
            " ('MR','موريتانيا','Mauretanien','Mauritania','Mauritania','Mauritanie','モーリタニア','Mauritânia','Мавритания','毛里塔尼亚','毛里塔尼亚',1,1)," +
            " ('MS','مونتسرات','Montserrat','Montserrat','Montserrat','Montserrat','モントセラト','Montserrat','Монтсеррат','蒙特塞拉特','蒙特塞拉特',4,1)," +
            " ('MT','مالطا','Malta','Malta','Malta','Malte','マルタ','Malta','Мальта','馬爾他','馬爾他',3,1)," +
            " ('MU','موريشيوس','Mauritius','Mauritius','Mauricio','Maurice','モーリシャス','Maurícia','Маврикий','模里西斯','模里西斯',1,1)," +
            " ('MV','جزر المالديف','Malediven','Maldives','Maldivas','Maldives','モルディブ','Maldivas','Мальдивы','馬爾地夫','馬爾地夫',2,1)," +
            " ('MW','مالاوي','Malawi','Malawi','Malaui','Malawi','マラウイ','Malawi','Малави','马拉维','马拉维',1,1)," +
            " ('MX','المكسيك','Mexiko','Mexico','México','Mexique','メキシコ','México','Мексика','墨西哥','墨西哥',4,1)," +
            " ('MY','ماليزيا','Malaysia','Malaysia','Malasia','Malaisie','マレーシア','Malásia','Малайзия','马来西亚','马来西亚',2,1)," +
            " ('MZ','موزمبيق','Mosambik','Mozambique','Mozambique','Mozambique','モザンビーク','Moçambique','Мозамбик','莫桑比克','莫桑比克',1,1)," +
            " ('NA','ناميبيا','Namibia','Namibia','Namibia','Namibie','ナミビア','Namíbia','Намибия','纳米比亚','纳米比亚',1,1)," +
            " ('NC','كاليدونيا الجديدة','Neukaledonien','New Caledonia','Nueva Caledonia','Nouvelle-Calédonie','ニューカレドニア','Nova Caledônia','Новая Каледония','新喀里多尼亞','新喀里多尼亞',6,1)," +
            " ('NE','النيجر','Niger','Niger','Níger','Niger','ニジェール','Níger','Нигер','尼日尔','尼日尔',1,1)," +
            " ('NF','جزيرة نورفولك','Norfolkinsel','Norfolk Island','Isla Norfolk','Île Norfolk','ノーフォーク島','Ilha Norfolk','Остров Норфолк','诺福克岛','诺福克岛',6,1)," +
            " ('NG','نيجيريا','Nigeria','Nigeria','Nigeria','Nigeria','ナイジェリア','Nigéria','Нигерия','奈及利亞','奈及利亞',1,1)," +
            " ('NI','نيكاراغوا','Nicaragua','Nicaragua','Nicaragua','Nicaragua','ニカラグア','Nicarágua','Никарагуа','尼加拉瓜','尼加拉瓜',4,1)," +
            " ('NL','هولندا','Niederlande','Netherlands','Países Bajos','Pays-Bas','オランダ','Países Baixos','Нидерланды','荷蘭','荷蘭',3,1)," +
            " ('NO','النرويج','Norwegen','Norway','Noruega','Norvège','ノルウェー','Noruega','Норвегия','挪威','挪威',3,1)," +
            " ('NP','نيبال','Nepal','Nepal','Nepal','Népal','ネパール','Nepal','Непал','尼泊尔','尼泊尔',2,1)," +
            " ('NR','ناورو','Nauru','Nauru','Nauru','Nauru','ナウル','Nauru','Науру','瑙鲁','瑙鲁',6,1)," +
            " ('NU','نييوي','Niue','Niue','Niue','Niue','ニウエ','Niue','Ниуэ','纽埃','纽埃',6,1)," +
            " ('NZ','نيوزيلندا','Neuseeland','New Zealand','Nueva Zelanda','Nouvelle-Zélande','ニュージーランド','Nova Zelândia','Новая Зеландия','新西蘭','新西蘭',6,1)," +
            " ('OM','عمان','Oman','Oman','Omán','Oman','オマーン','Omã','Оман','阿曼','阿曼',2,1)," +
            " ('PA','بنما','Panama','Panama','Panamá','Panama','パナマ','Panamá','Панама','巴拿马','巴拿马',4,1)," +
            " ('PE','بيرو','Peru','Peru','Perú','Pérou','ペルー','Peru','Перу','秘魯','秘魯',5,1)," +
            " ('PF','بولينزيا الفرنسية','Französisch-Polynesien','French Polynesia','Polinesia Francesa','Polynésie française','フランス領ポリネシア','Polinésia Francesa','Французская Полинезия','法屬玻里尼西亞','法屬玻里尼西亞',6,1)," +
            " ('PG','بابوا غينيا الجديدة','Papua-Neuguinea','Papua New Guinea','Papúa Nueva Guinea','Papouasie-Nouvelle-Guinée','パプアニューギニア','Papua Nova Guiné','Папуа — Новая Гвинея','巴布亚新几内亚','巴布亚新几内亚',6,1)," +
            " ('PH','الفلبين','Philippinen','Philippines','Filipinas','Philippines','フィリピン','Filipinos','Филиппины','菲律賓','菲律賓',2,1)," +
            " ('PK','باكستان','Pakistan','Pakistan','Pakistán','Pakistan','パキスタン','Paquistão','Пакистан','巴基斯坦','巴基斯坦',2,1)," +
            " ('PL','بولندا','Polen','Poland','Polonia','Pologne','ポーランド','Polônia','Польша','波蘭','波蘭',3,1)," +
            " ('PM','سان بيير وميكلون','Saint-Pierre und Miquelon','Saint Pierre and Miquelon','San Pedro y Miquelón','Saint-Pierre-et-Miquelon','サンピエール島・ミクロン島','São Pedro e Miquelon','Сен-Пьер и Микелон','圣皮埃尔和密克隆','圣皮埃尔和密克隆',4,1)," +
            " ('PN','جزر بيتكيرن','Pitcairninseln','Pitcairn','Islas Pitcairn','Îles Pitcairn','ピトケアン','Pitcairn','Острова Питкэрн','皮特凯恩群岛','皮特凯恩群岛',1,1)," +
            " ('PR','بورتوريكو','Puerto Rico','Puerto Rico','Puerto Rico','Porto Rico','プエルトリコ','Porto Rico','Пуэрто-Рико','波多黎各','波多黎各',4,1)," +
            " ('PS','فلسطين','Staat Palästina','Palestine, State of','Palestina','Palestine','パレスチナ','Palestina','Государство Палестина','巴勒斯坦','巴勒斯坦',2,1)," +
            " ('PT','البرتغال','Portugal','Portugal','Portugal','Portugal','ポルトガル','Portugal','Португалия','葡萄牙','葡萄牙',3,1)," +
            " ('PW','بالاو','Palau','Palau','Palaos','Palaos','パラオ','Palau','Палау','帛琉','帛琉',6,1)," +
            " ('PY','باراغواي','Paraguay','Paraguay','Paraguay','Paraguay','パラグアイ','Paraguai','Парагвай','巴拉圭','巴拉圭',5,1)," +
            " ('QA','قطر','Katar','Qatar','Catar','Qatar','カタール','Catar','Катар','卡塔尔','卡塔尔',2,1)," +
            " ('RE','لا ريونيون','Réunion','Reunion','Reunión','La Réunion','レユニオン','Reunião','Реюньон','留尼汪','留尼汪',1,1)," +
            " ('RO','رومانيا','Rumänien','Romania','Rumania','Roumanie','ルーマニア','Romênia','Румыния','羅馬尼亞','羅馬尼亞',3,1)," +
            " ('RS','صربيا','Serbien','Serbia','Serbia','Serbie','セルビア','Sérvio','Сербия','塞爾維亞','塞爾維亞',3,1)," +
            " ('RU','روسيا','Russische Föderation','Russia','Rusia','Russie','ロシア','Rússia','Россия','俄羅斯','俄羅斯',2,1)," +
            " ('RW','رواندا','Ruanda','Rwanda','Ruanda','Rwanda','ルワンダ','Ruanda','Руанда','卢旺达','卢旺达',1,1)," +
            " ('SA','السعودية','Saudi-Arabien','Saudi Arabia','Arabia Saudita','Arabie saoudite','サウジアラビア','Arábia Saudita','Саудовская Аравия','沙烏地阿拉伯','沙烏地阿拉伯',2,1)," +
            " ('SB','جزر سليمان','Salomonen','Solomon Islands','Islas Salomón','Salomon','ソロモン諸島','Ilhas Salomão','Соломоновы Острова','所罗门群岛','所罗门群岛',6,1)," +
            " ('SC','سيشل','Seychellen','Seychelles','Seychelles','Seychelles','セーシェル','Seychelles','Сейшельские Острова','塞舌尔','塞舌尔',1,1)," +
            " ('SD','السودان','Sudan','Sudan','Sudán','Soudan','スーダン','Sudão','Судан','苏丹','苏丹',1,1)," +
            " ('SE','السويد','Schweden','Sweden','Suecia','Suède','スウェーデン','Suécia','Швеция','瑞典','瑞典',3,1)," +
            " ('SG','سنغافورة','Singapur','Singapore','Singapur','Singapour','シンガポール','Cingapura','Сингапур','新加坡','新加坡',2,1)," +
            " ('SH','سانت هيلانة وأسينشين وتريستان دا كونا','St. Helena','Saint Helena, Ascension and Tristan da Cunha','Santa Elena, Ascensión y Tristán de Acuña','Sainte-Hélène, Ascension et Tristan da Cunha','セントヘレナ・アセンションおよびトリスタンダクーニャ','Santa Helena, Ascensão e Tristão da Cunha','Острова Святой Елены, Вознесения и Тристан-да-Кунья','圣赫勒拿、阿森松和特里斯坦-达库尼亚','圣赫勒拿、阿森松和特里斯坦-达库尼亚',1,1)," +
            " ('SI','سلوفينيا','Slowenien','Slovenia','Eslovenia','Slovénie','スロベニア','Eslovênia','Словения','斯洛維尼亞','斯洛維尼亞',3,1)," +
            " ('SJ','سفالبارد ويان ماين','Svalbard und Jan Mayen','Svalbard and Jan Mayen','Svalbard y Jan Mayen','Svalbard et ile Jan Mayen','スヴァールバル諸島およびヤンマイエン島','Svalbard e Jan Mayen','Шпицберген и Ян-Майен','斯瓦尔巴和扬马延','斯瓦尔巴和扬马延',1,1)," +
            " ('SK','سلوفاكيا','Slowakei','Slovakia','Eslovaquia','Slovaquie','スロバキア','Eslováquia','Словакия','斯洛伐克','斯洛伐克',3,1)," +
            " ('SL','سيراليون','Sierra Leone','Sierra Leone','Sierra Leona','Sierra Leone','シエラレオネ','Serra Leoa','Сьерра-Леоне','塞拉利昂','塞拉利昂',1,1)," +
            " ('SM','سان مارينو','San Marino','San Marino','San Marino','Saint-Marin','サンマリノ','San Marino','Сан-Марино','圣马力诺','圣马力诺',3,1)," +
            " ('SN','السنغال','Senegal','Senegal','Senegal','Sénégal','セネガル','Senegal','Сенегал','塞内加尔','塞内加尔',1,1)," +
            " ('SO','الصومال','Somalia','Somalia','Somalia','Somalie','ソマリア','Somália','Сомали','索馬利亞','索馬利亞',1,1)," +
            " ('SR','سورينام','Suriname','Suriname','Surinam','Suriname','スリナム','Suriname','Суринам','苏里南','苏里南',5,1)," +
            " ('SS','جنوب السودان','Südsudan','South Sudan','Sudán del Sur','Soudan du Sud','南スーダン','Sudão do Sul','Южный Судан','南蘇丹','南蘇丹',1,1)," +
            " ('ST','ساو تومي وبرينسيب','São Tomé und Príncipe','Sao Tome and Principe','Santo Tomé y Príncipe','Sao Tomé-et-Principe','サントメ・プリンシペ','São Tomé e Príncipe','Сан-Томе и Принсипи','聖多美和普林西比','聖多美和普林西比',1,1)," +
            " ('SV','السلفادور','El Salvador','El Salvador','El Salvador','Salvador','エルサルバドル','El Salvador','Сальвадор','薩爾瓦多','薩爾瓦多',4,1)," +
            " ('SX','سينت مارتن','Sint Maarten','Sint Maarten (Dutch part)','San Martín','Saint-Martin','シント・マールテン (オランダ領)','São Martinho (parte holandesa)','Синт-Мартен','聖馬丁','聖馬丁',4,1)," +
            " ('SY','سوريا','Syrien','Syria','Siria','Syrie','シリア・アラブ共和国','Sírio','Сирия','叙利亚','叙利亚',2,1)," +
            " ('SZ','سوازيلاند','Swasiland','Eswatini','Suazilandia','Swaziland','エスワティニ','Eswatini','Эсватини','斯威士兰','斯威士兰',1,1)," +
            " ('TC','جزر توركس وكايكوس','Turks- und Caicosinseln','Turks and Caicos Islands','Islas Turcas y Caicos','Îles Turques-et-Caïques','タークス・カイコス諸島','Ilhas Turks e Caicos','Теркс и Кайкос','特克斯和凯科斯群岛','特克斯和凯科斯群岛',4,1)," +
            " ('TD','تشاد','Tschad','Chad','Chad','Tchad','チャド','Chade','Чад','乍得','乍得',1,1)," +
            " ('TF','أراض فرنسية جنوبية وأنتارتيكية','Französische Süd- und Antarktisgebiete','French Southern Territories','Tierras Australes y Antárticas Francesas','Terres australes et antarctiques françaises','フランス領南方・南極地域','Territórios Franceses do Sul','Французские Южные и Антарктические Территории','法属南方和南极洲领地','法属南方和南极洲领地',7,1)," +
            " ('TG','توغو','Togo','Togo','Togo','Togo','トーゴ','Ir','Того','多哥','多哥',1,1)," +
            " ('TH','تايلاند','Thailand','Thailand','Tailandia','Thaïlande','タイ','Tailândia','Таиланд','泰國','泰國',2,1)," +
            " ('TJ','طاجيكستان','Tadschikistan','Tajikistan','Tayikistán','Tadjikistan','タジキスタン','Tajiquistão','Таджикистан','塔吉克斯坦','塔吉克斯坦',2,1)," +
            " ('TK','توكيلاو','Tokelau','Tokelau','Tokelau','Tokelau','トケラウ','Tokelau','Токелау','托克勞','托克勞',6,1)," +
            " ('TL','تيمور الشرقية','Timor-Leste','Timor-Leste','Timor Oriental','Timor oriental','東ティモール','Timor-Leste','Восточный Тимор','东帝汶','东帝汶',2,1)," +
            " ('TM','تركمانستان','Turkmenistan','Turkmenistan','Turkmenistán','Turkménistan','トルクメニスタン','Turquemenistão','Туркмения','土库曼斯坦','土库曼斯坦',2,1)," +
            " ('TN','تونس','Tunesien','Tunisia','Túnez','Tunisie','チュニジア','Tunísia','Тунис','突尼西亞','突尼西亞',1,1)," +
            " ('TO','تونغا','Tonga','Tonga','Tonga','Tonga','トンガ','Tonga','Тонга','汤加','汤加',6,1)," +
            " ('TR','تركيا','Türkei','Turkey','Turquía','Turquie','トルコ','Turquia','Турция','土耳其','土耳其',2,1)," +
            " ('TT','ترينيداد وتوباغو','Trinidad und Tobago','Trinidad and Tobago','Trinidad y Tobago','Trinité-et-Tobago','トリニダード・トバゴ','Trindade e Tobago','Тринидад и Тобаго','千里達及托巴哥','千里達及托巴哥',4,1)," +
            " ('TV','توفالو','Tuvalu','Tuvalu','Tuvalu','Tuvalu','ツバル','Tuvalu','Тувалу','图瓦卢','图瓦卢',6,1)," +
            " ('TW','تايوان','Taiwan','Taiwan','Taiwán','Taïwan','台湾','Taiwan','Китайская Республика','台湾','台湾',2,1)," +
            " ('TZ','تنزانيا','Tansania, Vereinigte Republik','Tanzania, United Republic of','Tanzania','Tanzanie','タンザニア','Tanzânia, República Unida da','Танзания','坦桑尼亚','坦桑尼亚',1,1)," +
            " ('UA','أوكرانيا','Ukraine','Ukraine','Ucrania','Ukraine','ウクライナ','Ucrânia','Украина','烏克蘭','烏克蘭',3,1)," +
            " ('UG','أوغندا','Uganda','Uganda','Uganda','Ouganda','ウガンダ','Uganda','Уганда','乌干达','乌干达',1,1)," +
            " ('UM','جزر الولايات المتحدة الصغيرة النائية','United States Minor Outlying Islands','United States Minor Outlying Islands','Islas ultramarinas de Estados Unidos','Îles mineures éloignées des États-Unis','合衆国領有小離島','Ilhas Menores Distantes dos Estados Unidos','Внешние малые острова (США)','美國本土外小島嶼','美國本土外小島嶼',1,1)," +
            " ('US','الولايات المتحدة','Vereinigte Staaten von Amerika','United States of America','Estados Unidos','États-Unis','アメリカ合衆国','Estados Unidos da America','США','美國','美國',4,1)," +
            " ('UY','الأوروغواي','Uruguay','Uruguay','Uruguay','Uruguay','ウルグアイ','Uruguai','Уругвай','乌拉圭','乌拉圭',5,1)," +
            " ('UZ','أوزبكستان','Usbekistan','Uzbekistan','Uzbekistán','Ouzbékistan','ウズベキスタン','Usbequistão','Узбекистан','乌兹别克斯坦','乌兹别克斯坦',2,1)," +
            " ('VA','الفاتيكان','Vatikanstadt','Holy See','Ciudad del Vaticano','Saint-Siège','バチカン市国','Cidade do Vaticano','Ватикан','梵蒂冈','梵蒂冈',3,1)," +
            " ('VC','سانت فينسنت والغرينادين','St. Vincent und die Grenadinen','Saint Vincent and the Grenadines','San Vicente y las Granadinas','Saint-Vincent-et-les-Grenadines','セントビンセントおよびグレナディーン諸島','São Vicente e Granadinas','Сент-Винсент и Гренадины','圣文森特和格林纳丁斯','圣文森特和格林纳丁斯',4,1)," +
            " ('VE','فنزويلا','Venezuela','Venezuela (Bolivarian Republic of)','Venezuela','Venezuela','ベネズエラ・ボリバル共和国','Venezuela','Венесуэла','委內瑞拉','委內瑞拉',5,1)," +
            " ('VG','جزر العذراء البريطانية','Britische Jungferninseln','Virgin Islands (British)','Islas Vírgenes Británicas','Îles Vierges britanniques','イギリス領ヴァージン諸島','Ilhas Virgens Britânicas','Виргинские Острова (Великобритания)','英屬維爾京群島','英屬維爾京群島',4,1)," +
            " ('VI','جزر العذراء الأمريكية','Amerikanische Jungferninseln','Virgin Islands (U.S.)','Islas Vírgenes de los Estados Unidos','Îles Vierges des États-Unis','アメリカ領ヴァージン諸島','Ilhas Virgens (EUA)','Виргинские Острова (США)','美屬維爾京群島','美屬維爾京群島',4,1)," +
            " ('VN','فيتنام','Vietnam','Vietnam','Vietnam','Viêtnam','ベトナム','Vietnã','Вьетнам','越南','越南',2,1)," +
            " ('VU','فانواتو','Vanuatu','Vanuatu','Vanuatu','Vanuatu','バヌアツ','Vanuatu','Вануату','瓦努阿圖','瓦努阿圖',6,1)," +
            " ('WF','والس وفوتونا','Wallis und Futuna','Wallis and Futuna','Wallis y Futuna','Wallis-et-Futuna','ウォリス・フツナ','Wallis e Futuna','Уоллис и Футуна','瓦利斯和富圖納','瓦利斯和富圖納',6,1)," +
            " ('WS','ساموا','Samoa','Samoa','Samoa','Samoa','サモア','Samoa','Самоа','萨摩亚','萨摩亚',6,1)," +
            " ('YE','اليمن','Jemen','Yemen','Yemen','Yémen','イエメン','Iémen','Йемен','葉門','葉門',2,1)," +
            " ('YT','مايوت','Mayotte','Mayotte','Mayotte','Mayotte','マヨット','Mayotte','Майотта','马约特','马约特',1,1)," +
            " ('ZA','جنوب أفريقيا','Südafrika','South Africa','Sudáfrica','Afrique du Sud','南アフリカ','África do Sul','ЮАР','南非','南非',1,1)," +
            " ('ZM','زامبيا','Sambia','Zambia','Zambia','Zambie','ザンビア','Zâmbia','Замбия','尚比亞','尚比亞',1,1)," +
            " ('ZW','زيمبابوي','Simbabwe','Zimbabwe','Zimbabue','Zimbabwe','ジンバブエ','Zimbábue','Зимбабве','辛巴威','辛巴威',1,1)"
    _, err = db.Exec(insertQuery)
    if err != nil {
        db.Close()
        log.Fatalf("[FATAL] db.Exec() error: %s\n", err)
        return
    }
}


//////////////////////////////////////////////////////////////////////
// Select
//////////////////////////////////////////////////////////////////////
func Select(columns Columns, langCode string, orderby string, orderDesc bool, limit int, offset int) []Columns {
    var result []Columns
    whereFlag := false
    bufferQuery := make([]byte, 0)
    query := "SELECT * FROM " + TABLE_NAME
    bufferQuery = append(bufferQuery, query...)

    if columns.CountryCode != "" {
        query = " WHERE country_code='" + columns.CountryCode + "'"
        bufferQuery = append(bufferQuery, query...)
        whereFlag = true
    }

    if columns.Ar != "" {
        if whereFlag {
            query = " AND ar='" + columns.Ar + "'"
        } else {
            query = " WHERE ar='" + columns.Ar + "'"
            whereFlag = true
        }
        bufferQuery = append(bufferQuery, query...)
    }

    if columns.De != "" {
        if whereFlag {
            query = " AND de='" + columns.De + "'"
        } else {
            query = " WHERE de='" + columns.De + "'"
            whereFlag = true
        }
        bufferQuery = append(bufferQuery, query...)
    }

    if columns.En != "" {
        if whereFlag {
            query = " AND en='" + columns.En + "'"
        } else {
            query = " WHERE en='" + columns.En + "'"
            whereFlag = true
        }
        bufferQuery = append(bufferQuery, query...)
    }

    if columns.Es != "" {
        if whereFlag {
            query = " AND es='" + columns.Es + "'"
        } else {
            query = " WHERE es='" + columns.Es + "'"
            whereFlag = true
        }
        bufferQuery = append(bufferQuery, query...)
    }

    if columns.Fr != "" {
        if whereFlag {
            query = " AND fr='" + columns.Fr + "'"
        } else {
            query = " WHERE fr='" + columns.Fr + "'"
            whereFlag = true
        }
        bufferQuery = append(bufferQuery, query...)
    }

    if columns.Ja != "" {
        if whereFlag {
            query = " AND ja='" + columns.Ja + "'"
        } else {
            query = " WHERE ja='" + columns.Ja + "'"
            whereFlag = true
        }
        bufferQuery = append(bufferQuery, query...)
    }

    if columns.Pt != "" {
        if whereFlag {
            query = " AND pt='" + columns.Pt + "'"
        } else {
            query = " WHERE pt='" + columns.Pt + "'"
            whereFlag = true
        }
        bufferQuery = append(bufferQuery, query...)
    }

    if columns.Ru != "" {
        if whereFlag {
            query = " AND ru='" + columns.Ru + "'"
        } else {
            query = " WHERE ru='" + columns.Ru + "'"
            whereFlag = true
        }
        bufferQuery = append(bufferQuery, query...)
    }

    if columns.ZhCn != "" {
        if whereFlag {
            query = " AND zh_cn='" + columns.ZhCn + "'"
        } else {
            query = " WHERE zh_cn='" + columns.ZhCn + "'"
            whereFlag = true
        }
        bufferQuery = append(bufferQuery, query...)
    }

    if columns.ZhTw != "" {
        if whereFlag {
            query = " AND zh_tw='" + columns.ZhTw + "'"
        } else {
            query = " WHERE zh_tw='" + columns.ZhTw + "'"
            whereFlag = true
        }
        bufferQuery = append(bufferQuery, query...)
    }

    if columns.Continent != 0 {
        if whereFlag {
            query = " AND continent='" + strconv.Itoa(columns.Continent)
        } else {
            query = " WHERE continent=" + strconv.Itoa(columns.Continent)
            whereFlag = true
        }
        bufferQuery = append(bufferQuery, query...)
    }

    if columns.Status == 0 && columns.Status == 1 {
        if whereFlag {
            query = " AND status='" + strconv.Itoa(columns.Status)
        } else {
            query = " WHERE status=" + strconv.Itoa(columns.Status)
            whereFlag = true
        }
        bufferQuery = append(bufferQuery, query...)
    }

    if orderDesc {
        query = " ORDER BY " + orderby + " DESC"
    } else {
        query = " ORDER BY " + orderby + " ASC"
    }
    bufferQuery = append(bufferQuery, query...)

    if limit > 0 && offset > 0 {
        query = " LIMIT " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)
    } else if limit > 0 {
        query = " LIMIT " + strconv.Itoa(limit)
    } else if offset > 0 {
        query = " LIMIT " + strconv.Itoa(offset)
    } else {
        query = ""
    }
    bufferQuery = append(bufferQuery, query...)

    rows, err := db.Query(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] [posts] db.Query() error: %s\n", err)
        return result
    }

    defer rows.Close()
    for rows.Next() {
        var countryCode string
        var ar string
        var de string
        var en string
        var es string
        var fr string
        var ja string
        var pt string
        var ru string
        var zhCn string
        var zhTw string
        var continent int
        var status int
        if err := rows.Scan(&countryCode, &ar, &de, &en, &es, &fr, &ja, &pt, &ru, &zhCn, &zhTw, &continent, &status); err != nil {
            log.Printf("[ERROR] [it] rows.Scan() error: %s\n", err)
            return result
        }
        columns := Columns{
            CountryCode: countryCode,
            Ar: ar,
            De: de,
            En: en,
            Es: es,
            Fr: fr,
            Ja: ja,
            Pt: pt,
            Ru: ru,
            ZhCn: zhCn,
            ZhTw: zhTw,
            Continent: status,
            Status: status,
        }
        if langCode == "ar" {
            columns.Name = ar
        } else if langCode == "de" {
            columns.Name = de
        } else if langCode == "es" {
            columns.Name = es
        } else if langCode == "fr" {
            columns.Name = fr
        }else if langCode == "ja" {
            columns.Name = ja
        }else if langCode == "pt" {
            columns.Name = pt
        }else if langCode == "ru" {
            columns.Name = ru
        }else if langCode == "zh_cn" {
            columns.Name = zhCn
        }else if langCode == "zh_tw" {
            columns.Name = zhCn
        }else {
            columns.Name = en
        }
        result = append(result, columns)
    }
    return result
}


//////////////////////////////////////////////////////////////////////
// Get only active countries
//////////////////////////////////////////////////////////////////////
func GetOnlyActive(langCode string) []Columns {
    columns := Columns{
        Status: 1,
    }
    return Select(columns, langCode, "country_code", false, 0, 0)
}


//////////////////////////////////////////////////////////////////////
// Get only active countries in Africa
//////////////////////////////////////////////////////////////////////
func GetAfricaOnlyActive(langCode string) []Columns {
    columns := Columns{
        Continent: 1,
        Status: 1,
    }
    return Select(columns, langCode, "country_code", false, 0, 0)
}


//////////////////////////////////////////////////////////////////////
// Get only active countries in Asia
//////////////////////////////////////////////////////////////////////
func GetAsiaOnlyActive(langCode string) []Columns {
    columns := Columns{
        Continent: 2,
        Status: 1,
    }
    return Select(columns, langCode, "country_code", false, 0, 0)
}


//////////////////////////////////////////////////////////////////////
// Get only active countries in Europe
//////////////////////////////////////////////////////////////////////
func GetEuropeOnlyActive(langCode string) []Columns {
    columns := Columns{
        Continent: 3,
        Status: 1,
    }
    return Select(columns, langCode, "country_code", false, 0, 0)
}


//////////////////////////////////////////////////////////////////////
// Get only active countries in North America
//////////////////////////////////////////////////////////////////////
func GetNorthAmericaOnlyActive(langCode string) []Columns {
    columns := Columns{
        Continent: 4,
        Status: 1,
    }
    return Select(columns, langCode, "country_code", false, 0, 0)
}


//////////////////////////////////////////////////////////////////////
// Get only active countries in South America
//////////////////////////////////////////////////////////////////////
func GetSouthAmericaOnlyActive(langCode string) []Columns {
    columns := Columns{
        Continent: 5,
        Status: 1,
    }
    return Select(columns, langCode, "country_code", false, 0, 0)
}


//////////////////////////////////////////////////////////////////////
// Get only active countries in Australia / Oceania
//////////////////////////////////////////////////////////////////////
func GetAustraliaOceaniaOnlyActive(langCode string) []Columns {
    columns := Columns{
        Continent: 6,
        Status: 1,
    }
    return Select(columns, langCode, "country_code", false, 0, 0)
}


//////////////////////////////////////////////////////////////////////
// Get only active countries in Antarctica
//////////////////////////////////////////////////////////////////////
func GetAntarcticaOnlyActive(langCode string) []Columns {
    columns := Columns{
        Continent: 7,
        Status: 1,
    }
    return Select(columns, langCode, "country_code", false, 0, 0)
}

