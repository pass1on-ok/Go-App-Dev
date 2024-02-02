package api

type Titan struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Image     string   `json:"img"`
	Height    int      `json:"height"`
	Abilities []string `json:"abilities"`
	Ally      string   `json:"ally"`
}

var Titans = []Titan{
	{1, "Attack Titan", "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTSVa1iO1y5RXcPf8EYNyZPXOnK5nNio7FFDg&usqp=CAU", 15, []string{"Future memory inheritance"}, "Eldia"},
	{2, "Founding Titan", "https://i.pinimg.com/originals/d8/4f/27/d84f27082ec276814a3de844bcbba617.png", 350, []string{"Titan creation", "Titan behavioral control", "Memory and body manipulation of subjects of Ymir", "Telepathic communication with subjects of Ymir"}, "Eldia"},
	{3, "War Hammer Titan", "https://pm1.aminoapps.com/6939/c90bc9298e330dc5b4cbda349012951e122638ear1-554-554v2_00.jpg", 15, []string{"Structural hardening", "Remote operation"}, "Eldia"},
	{4, "Colossal Titan", "https://image.civitai.com/xG1nkqKTMzGDvpLrqFT7WA/cc5b3634-8505-4c45-1555-538bdc484900/width=450/02887-3684040636.jpeg", 60, []string{"Explosive transformation", "Immense size and strength", "Steam emission"}, "Eldia"},
	{5, "Armored Titan", "https://qph.cf2.quoracdn.net/main-qimg-4aa8d12a9e448f6d1d6a054eb5ba9c41-lq", 15, []string{"Armored skin", "Hardening"}, "Marley"},
	{6, "Female Titan", "https://i1.sndcdn.com/artworks-g1ylJr6dmpsKcSUS-WOgD5g-t500x500.jpg", 14, []string{"Versatility", "Titan", "Attraction", "Crystallization"}, "Marley"},
	{7, "Beast Titan", "https://www.looper.com/img/gallery/zekes-beast-titan-powers-from-attack-on-titan-explained/l-intro-1620834206.jpg", 17, []string{"Powerful and accurate throwing", "Hardening", "Speech", "Can turn subjects of Ymir into titans that he can crudely control (users with royal blood only)"}, "None"},
	{8, "Jaw Titan", "https://mangainsider.com/wp-content/uploads/2022/02/Jaw-Titan-Guide.png", 5, []string{"Powerful jaw strength", "Hardened claws", "Great speed and agility"}, "Marley"},
	{9, "Cart Titan", "https://i.pinimg.com/736x/3f/2d/57/3f2d5757638fd15ef4d0d9dca90ba196.jpg", 4, []string{"Quadrupedal form", "High endurance", "Great speed", "Speech"}, "Marley"},
}
