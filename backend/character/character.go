package character

type Character struct {
	ID      int
	Name    string
	Avatar  string
	GroupID int
}

var AllCharacters []Character = []Character{
	{Name: "Bylad", ID: 1, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63734_thumb.jpg", GroupID: 1},
	{Name: "Sothis", ID: 2, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63736_thumb.jpg", GroupID: 1},
	{Name: "Rhea", ID: 3, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63737_thumb.jpg", GroupID: 1},
	{Name: "Seteth", ID: 4, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63738_thumb.jpg", GroupID: 1},
	{Name: "Flayn", ID: 5, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63739_thumb.jpg", GroupID: 1},
	{Name: "Hanneman", ID: 6, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63740_thumb.jpg", GroupID: 1},
	{Name: "Manuela", ID: 7, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63741_thumb.jpg", GroupID: 1},
	{Name: "Catherine", ID: 8, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63742_thumb.jpg", GroupID: 1},
	{Name: "Shamir", ID: 9, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63743_thumb.jpg", GroupID: 1},
	{Name: "Cyril", ID: 10, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63744_thumb.jpg", GroupID: 1},
	{Name: "Jeralt", ID: 11, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63745_thumb.jpg", GroupID: 1},
	{Name: "Edelgard", ID: 12, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63746_thumb.jpg", GroupID: 1},
	{Name: "Hubert", ID: 13, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63747_thumb.jpg", GroupID: 1},
	{Name: "Ferdinand", ID: 14, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63748_thumb.jpg", GroupID: 1},
	{Name: "Linhardt", ID: 15, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63749_thumb.jpg", GroupID: 1},
	{Name: "Caspar", ID: 16, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63750_thumb.jpg", GroupID: 1},
	{Name: "Bernadetta", ID: 17, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63751_thumb.jpg", GroupID: 1},
	{Name: "Dorothea", ID: 18, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63752_thumb.jpg", GroupID: 1},
	{Name: "Petra", ID: 19, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63753_thumb.jpg", GroupID: 1},
	{Name: "Jeritza", ID: 20, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63754_thumb.jpg", GroupID: 1},
	{Name: "Dimitri", ID: 21, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63755_thumb.jpg", GroupID: 1},
	{Name: "Dedue", ID: 22, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63756_thumb.jpg", GroupID: 1},
	{Name: "Felix", ID: 23, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63757_thumb.jpg", GroupID: 1},
	{Name: "Sylvain", ID: 24, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63758_thumb.jpg", GroupID: 1},
	{Name: "Ashe", ID: 25, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63759_thumb.jpg", GroupID: 1},
	{Name: "Annette", ID: 26, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63760_thumb.jpg", GroupID: 1},
	{Name: "Mercedes", ID: 27, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63761_thumb.jpg", GroupID: 1},
	{Name: "Ingrid", ID: 28, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63762_thumb.jpg", GroupID: 1},
	{Name: "Claude", ID: 29, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63763_thumb.jpg", GroupID: 1},
	{Name: "Lorenz", ID: 30, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63764_thumb.jpg", GroupID: 1},
	{Name: "Raphael", ID: 31, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63765_thumb.jpg", GroupID: 1},
	{Name: "Ignatz", ID: 32, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63766_thumb.jpg", GroupID: 1},
	{Name: "Hilda", ID: 33, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63767_thumb.jpg", GroupID: 1},
	{Name: "Marianne", ID: 34, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63768_thumb.jpg", GroupID: 1},
	{Name: "Lysithea", ID: 35, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63769_thumb.jpg", GroupID: 1},
	{Name: "Leonie", ID: 36, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63770_thumb.jpg", GroupID: 1},
	{Name: "Bylass", ID: 37, Avatar: "https://em-uploads.s3.amazonaws.com/anonymousicons/63771_thumb.jpg", GroupID: 1},
	{Name: "Albedo", ID: 38, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Albedo_Thumb.png", GroupID: 2},
	{Name: "Amber", ID: 39, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Amber_Thumb.png", GroupID: 2},
	{Name: "Barbara", ID: 40, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Barbara_Thumb.png", GroupID: 2},
	{Name: "Beidou", ID: 41, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Beidou_Thumb.png", GroupID: 2},
	{Name: "Bennett", ID: 42, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Bennett_Thumb.png", GroupID: 2},
	{Name: "Chongyun", ID: 43, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Chongyun_Thumb.png", GroupID: 2},
	{Name: "Diluc", ID: 44, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Diluc_Thumb.png", GroupID: 2},
	{Name: "Ayaka", ID: 45, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Ayaka_Thumb.png", GroupID: 2},
	{Name: "Dainsleif", ID: 46, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Dainsleif_Thumb.png", GroupID: 2},
	{Name: "Diona", ID: 47, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Diona_Thumb.png", GroupID: 2},
	{Name: "Fischl", ID: 48, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Fischl_Thumb.png", GroupID: 2},
	{Name: "Ganyu", ID: 49, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Ganyu_Thumb.png", GroupID: 2},
	{Name: "Jean", ID: 50, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Jean_Thumb.png", GroupID: 2},
	{Name: "Kaeya", ID: 51, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Kaeya_Thumb.png", GroupID: 2},
	{Name: "Keqing", ID: 52, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Keqing_Thumb.png", GroupID: 2},
	{Name: "Klee", ID: 53, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Klee_Thumb.png", GroupID: 2},
	{Name: "Lisa", ID: 54, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Lisa_Thumb.png", GroupID: 2},
	{Name: "Mona", ID: 55, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Mona_Thumb.png", GroupID: 2},
	{Name: "Ningguang", ID: 56, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Ningguang_Thumb.png", GroupID: 2},
	{Name: "Noelle", ID: 57, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Noelle_Thumb.png", GroupID: 2},
	{Name: "Qiqi", ID: 58, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Qiqi_Thumb.png", GroupID: 2},
	{Name: "Razor", ID: 59, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Razor_Thumb.png", GroupID: 2},
	{Name: "Scaramouche", ID: 60, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Scaramouche_Thumb.png", GroupID: 2},
	{Name: "Sucrose", ID: 61, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Sucrose_Thumb.png", GroupID: 2},
	{Name: "Tartaglia", ID: 62, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Tartaglia_Thumb.png", GroupID: 2},
	{Name: "Traveler", ID: 63, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Traveler_Thumb.png", GroupID: 2},
	{Name: "Venti", ID: 64, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Venti_Thumb.png", GroupID: 2},
	{Name: "Xiangling", ID: 65, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Xiangling_Thumb.png", GroupID: 2},
	{Name: "Xiao", ID: 66, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Xiao_Thumb.png", GroupID: 2},
	{Name: "Xingqiu", ID: 67, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Xingqiu_Thumb.png", GroupID: 2},
	{Name: "Xinyan", ID: 68, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Xinyan_Thumb.png", GroupID: 2},
	{Name: "Zhongli", ID: 69, Avatar: "https://jaketheduck.com/games/images/mafia-GI-icons/Character_Zhongli_Thumb.png", GroupID: 2},
}

func SetupCharacters() {
	return
	/*
		Conn, err := database.GetDBConn()
		defer Conn.Close()
		if err != nil {
			log.Printf("Getting DB for Character Selected failed %+v\n", err)
			return
		}
		rows, err := Conn.Query("SELECT id, name, avatar, group_id FROM characters")
		if err != nil {
			log.Printf("Character SELECT failed %+v\n", err)
			return
		}
		characters := make([]Character, 0)
		for rows.Next() {
			c := Character{}
			rows.Scan(&c.ID, &c.Name, &c.Avatar, &c.GroupID)
			characters = append(characters, c)
		}
		AllCharacters = characters
		log.Printf("Characters Loaded: %+v\n", AllCharacters)*/
}
