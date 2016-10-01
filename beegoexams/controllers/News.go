package controllers

import "github.com/astaxie/beego"
import "encoding/json"

type News struct {
	beego.Controller
}


type NewsResult struct {
	News []DetailsResult
}

type DetailsResult struct {
	ShortDescription string
	LongDescription string
	URLFullImages []string
	URLMediumImages []string
	URLSmallImages []string
	ExtraInfo string
	Title string
	SutTitle string
	Date string
	CreatedBy string
	ModifyBy string
}



// Use Get rather than Post so that we can simulate easier in the browser
func (this *News) Get() {
	var result NewsResult

    new1 := DetailsResult{}
    new1.ShortDescription = "If confirmed, the claim of direct Russian involvement in the bombing that killed at least 20 people in Syria would have far-reaching consequences"
	new1.LongDescription = "US defence officials now believe that Russian planes dropped the bombs that destroyed a UN aid convoy and killed at least 20 people, the Guardian has learned. The claim of direct Russian involvement in the bombing, if confirmed, would have far-reaching consequences. Earlier on Tuesday, Ban Ki-moon used his farewell address to the UN general assembly to denounce it as a “sickening, savage and apparently deliberate attack”, describing the bombers at “cowards”, and UN officials have said it is a potential war crime. Ban Ki-moon condemns ‘apparently deliberate’ Syria aid convoy attack Read more The outgoing secretary general told world leaders in New York that the UN had been forced to suspend aid convoys in Syria because of Monday’s attack on Syrian Red Crescent trucks that were carrying UN food supplies to a rural area west of Aleppo city. Victims of the attack included the local director of the Syrian Arab Red Crescent, Omar Barakat. Ban hailed the dead aid workers as heroes and said “those who bombed them were cowards” before calling for accountability for crimes committed in the war. “Just when you think it cannot get any worse, the bar of depravity sinks lower,” he said. Aid officials said the convoy was hit from the air while food and medicine were being unloaded at a warehouse in opposition-controlled Urem al-Kubra. Reuters news agency quoted two US officials as saying two Russian Sukhoi SU-24 warplanes were in the sky above the aid convoy at the precise time it was struck. The White House and state department said they could not confirm the allegations, while the Russian foreign ministry rejected them with “resentment and indignation”. Previously, US officials had said that they would hold Moscow responsible for the attack, even if it was carried out by Syrian aircraft, as Russia had taken responsibility for the regime’s compliance with the ceasefire as part of the 9 September agreement. Syria aid convoy attack: ‘the bombardment was continuous’ Read more But Moscow has not conceded that the convoy was hit by an airstrike, claiming instead that the 18 lorries had “caught fire”. In a separate statement on Tuesday, the country’s defence ministry said that the aid convoy had been accompanied by a militants’ pickup truck armed with a heavy mortar, Russian news agencies reported. The US officials said there was no doubt the convoy was destroyed in an airstrike and that western coalition forces had no role in it. “There are only three parties that fly in Syria: the coalition, the Russians and the Syrian regime. It was not the coalition. We don’t fly over Aleppo. We have no reason to. We strike only Isis, and Isis is not there. We would leave it to the Russians and the Syrian regime to explain their actions,” said Capt Jeff Davis, a Pentagon spokesman. In a meeting with John Kerry, Sergei Lavrov, the Russian foreign minister, admitted that the Russian military had been monitoring the convoy – apparent drone surveillance footage of its progress had been live-streamed on a defence ministry website . But he claimed the Russians had “lost track of it when it entered rebel territory”, according to diplomatic sources. Moscow had launched an investigation, Lavrov told the other foreign ministers. Later on Tuesday, however, the Russian foreign ministry put out an angry denunciation of allegations against Moscow and Damascus."
	new1.URLFullImages = []string{"https://i.guim.co.uk/img/media/70f503852730d0fc45a2802e67651bb6e6510608/0_308_5068_3040/master/5068.jpg?w=620&q=55&auto=format&usm=12&fit=max&s=7d4dfa33260042be482b560ea00baf7f","https://interactive.guim.co.uk/embed/2016/09/syria-map/images/UremSyriaMap_ai2html-medium.png"}
	new1.URLMediumImages = []string{""}
	new1.URLSmallImages = []string{"https://ci6.googleusercontent.com/proxy/SOikPyVERFWh0Xk1Qzc7Yz6Yk9leihnlGslag3-eZQdmuQ-81sKetfKt40QjviABJFkVgkXTrTBPThr6PgRIOFPGIgkph4aZb7x2MRtyTwBoXK9NIYeUmFyYk4xysebhFH2ExtK4XpdNJRcD59np=s0-d-e1-ft#https://media.guim.co.uk/70f503852730d0fc45a2802e67651bb6e6510608/0_308_5068_3040/500.jpg"}
	new1.ExtraInfo = "The scene the morning after a convoy delivering aid was hit by a deadly airstrike in Syria. Photograph: Omar Haj Kadour/AFP/Getty Images "
	new1.Title = "Russian planes dropped bombs that destroyed UN convoy, US officials say"
	new1.SutTitle = ""
	new1.Date = "Wednesday 21 September 2016 10.32 BST"
	new1.CreatedBy = "The Guardian"
	new1.ModifyBy = "The Guardian"

    new2 := DetailsResult{}
    new2.ShortDescription = "All party members will have to abide by social media code of conduct banning harassment, hateful language, bullying and discrimination or risk expulsion"
	new2.LongDescription = "Labour is to try to stop a “tsunami of abuse” by making all existing and new members sign a pledge about online behaviour or face being barred from the party.Labour at impasse over bid to attract MPs back to frontbenchRead moreThe national executive committee (NEC) agreed to toughen up Labour’s stance on internet abuse during a crucial meeting on Tuesday, which comes as the party’s acrimonious leadership battle draws to a close.During a session that stretched over eight hours, Jeremy Corbyn expressed disappointment and sadness about the way in which a flood of Labour MPs resigned from the frontbench in protest at his performance earlier in the summer.The leader agreed to enter talks with senior colleagues – likely to include deputy leader, Tom Watson, chief whip Rosie Winterton and the chair of the parliamentary Labour party, John Cryer – about new ways to form his shadow cabinet. But he refused to sign up to a proposal by Watson, for MPs to be allowed to elect frontbenchers as a means of brokering peace within the parliamentary party.A motion to place a deadline of Saturday 24 September on the talks, so a new system is in place when the result of the leadership contest is announced, was also narrowly defeated by 16 to 15 votes.However the group unanimously agreed to a new statement on social media behaviour that will be included as a separate item in the terms and conditions of any membership.Party members will have to explicitly promise “to act within the spirit and rules of the Labour party in my conduct both on and offline, with members and non-members”.The statement they will have to sign adds: “I stand against all forms of abuse. I understand that if found to be in breach of the Labour party policy on online and offline abuse, I will be subject to the rules and procedures of the Labour party.”Punishments could include being suspended from the party or eventually being expelled.The pledge will be linked to a social media code of conduct, drawn up by the party, which warns that “harassment, intimidation, hateful language and bullying” will not be tolerated. It also lists discrimination on the basis of gender, race, religion, age, sexual orientation, gender identity or disability.“Abusing someone online is just as serious as doing so face to face. We stand against all forms of abuse and will take action against those who commit it,” it adds, warning against the use of anonymous accounts, trolling, sexualised language, or the publication of private information.The initiative comes after an MP, Ruth Smeeth, said that she had been subjected to 25,000 incidents of abuse, including a string of antisemitic attacks. The politician, who has been called a “CIA/MI5/Mossad informant” and a “fucking traitor,” said she believed that abuse had become normal for many colleagues.AdvertisementShe also revealed that she had received police protection after facing the vitriol, much of which occurred after she walked out of the launch of Shami Chakrabarti’s inquiry into antisemitism for the Labour party. “I think this is a great step forward,” said Smeeth, about the pledge. “MPs have been subject to a tsunami of abuse. It is unfortunate that we’ve got to this place but given the changing nature of social media I welcome the proposal.”She praised Watson and Yvette Cooper MP, who chairs the Reclaim the Internet campaign, for drawing up the proposal, which was presented to the NEC but will need to be rubber stamped by Labour conference.Cooper is expected to say on Saturday at a women’s conference prior to the main meeting in Liverpool that the party needs to go further ."
	new2.URLFullImages = []string{"https://i.guim.co.uk/img/media/499f5759c36d450d10c15db1065bce6661629d70/0_223_2589_1553/master/2589.jpg?w=620&q=55&auto=format&usm=12&fit=max&s=0d47efa3babf1f51c98d739aecdf21fa"}
	new2.URLMediumImages = []string{"https://i.guim.co.uk/img/media/499f5759c36d450d10c15db1065bce6661629d70/0_223_2589_1553/master/2589.jpg?w=620&q=55&auto=format&usm=12&fit=max&s=0d47efa3babf1f51c98d739aecdf21fa,https://i.guim.co.uk/img/media/499f5759c36d450d10c15db1065bce6661629d70/0_223_2589_1553/master/2589.jpg?w=620&q=55&auto=format&usm=12&fit=max&s=0d47efa3babf1f51c98d739aecdf21fa,https://i.guim.co.uk/img/media/499f5759c36d450d10c15db1065bce6661629d70/0_223_2589_1553/master/2589.jpg?w=620&q=55&auto=format&usm=12&fit=max&s=0d47efa3babf1f51c98d739aecdf21fa"}
	new2.URLSmallImages = []string{"https://i.guim.co.uk/img/media/499f5759c36d450d10c15db1065bce6661629d70/0_223_2589_1553/master/2589.jpg?w=620&q=55&auto=format&usm=12&fit=max&s=0d47efa3babf1f51c98d739aecdf21fa"}
	new2.ExtraInfo = ""
	new2.Title = "Labour introduces tougher policy on 'tsunami of online abuse'"
	new2.SutTitle = ""
	new2.Date = "Wednesday 21 Sep 2016"
	new2.CreatedBy = "The Guardian"
	new2.ModifyBy = "The Guardian"

    new3 := DetailsResult{}
    new3.ShortDescription = ""
	new3.LongDescription = "A disabled woman who suffers crippling pain claims she was forced to stand on a train for more than an hour after two men 'refused' to move from her reserved seats. Cat Lee, from Hebden Bridge in West Yorkshire, booked the seats on the Virgin East Coast service from King's Cross to Skipton yesterday. When the 43-year-old and her friend boarded the train, the men apparently refused to move - despite her telling them about her disability - so she posted their pictures on social media. When Cat Lee asked to sit in her reserved seats, the men apparently refused to move - despite her telling them about her disability - so she posted their pictures on social media +3 When Cat Lee asked to sit in her reserved seats, the men apparently refused to move - despite her telling them about her disability - so she posted their pictures on social media 'Their wives and mothers would be proud': Cat Lee was furious with the pair on the train +3 'Their wives and mothers would be proud': Cat Lee was furious with the pair on the train Ms Lee says she suffered nerve damage and crippling groin pain following an operation, which left her unable to work. DO YOU KNOW THE MEN ON THE TRAIN? Email amie.gordon@mailonline.co.uk or phone 020 361 53 729 The mother-of-one claims she took a friend for support on the three hour journey but said both were left to stand on the busy train for more than an hour before she managed to find another seat. Ms Lee took a photo of the two men and posted it on Facebook. Read more: http://www.dailymail.co.uk/news/article-3800246/I-m-sure-wives-mothers-proud-Disabled-woman-shames-fellow-train-passengers-Facebook-refused-reserved-seats.html#ixzz4KuQvwNPZ Follow us: @MailOnline on Twitter | DailyMail on Facebook"
	new3.URLFullImages = []string{"http://i.dailymail.co.uk/i/pix/2016/09/21/15/38A5B4B800000578-3800246-image-m-3_1474467685217.jpg"}
	new3.URLMediumImages = []string{"http://i.dailymail.co.uk/i/pix/2016/09/21/15/38A5B4B800000578-3800246-image-m-3_1474467685217.jpg","http://i.dailymail.co.uk/i/pix/2016/09/21/15/38A5B4B800000578-3800246-_Their_wives_and_mothers_would_be_proud_Cat_Lee_was_furious_with-m-25_1474468162892.jpg","http://i.dailymail.co.uk/i/pix/2016/09/21/15/38A5B13900000578-3800246-Cat_Lee_from_Hebden_Bridge_in_West_Yorkshire_booked_the_seats_on-m-12_1474468043452.jpg"}
	new3.URLSmallImages = []string{"http://i.dailymail.co.uk/i/pix/2016/09/21/15/38A5B4B800000578-3800246-image-m-3_1474467685217.jpg"}
	new3.ExtraInfo = ""
	new3.Title = "'I'm sure their wives and mothers would be proud': Disabled woman accuses fellow train passengers on Facebook after they 'refused to move from her reserved seats' "
	new3.SutTitle = ""
	new3.Date = " 14:28 GMT, 21 September 2016 | Updated: 16:06 GMT, 21 September 2016"
	new3.CreatedBy = "Amie Gordon For Mailonline"
	new3.ModifyBy = "Amie Gordon For Mailonline"
	
	new4 := DetailsResult{}
    new4.ShortDescription = "he 120-hour Certificate in TEFL offered by the Centro Cultural Colombo Americano aims to provide the highest standard in the teaching of the English language. It is designed for those who have little or no experience teaching English as a Foreign Language, or for those who are experienced, but lack the academic background and credentials necessary to advance. The certificate aims to equip students with a strong foundation of language rules, language learning theories, teaching techniques, and materials development with strong emphasis on communicative methodologies."
	new4.LongDescription = "The 120-hour Certificate in TEFL offered by the Centro Cultural Colombo Americano aims to provide the highest standard in the teaching of the English language. It is designed for those who have little or no experience teaching English as a Foreign Language, or for those who are experienced, but lack the academic background and credentials necessary to advance. The certificate aims to equip students with a strong foundation of language rules, language learning theories, teaching techniques, and materials development with strong emphasis on communicative methodologies. ..................................................................................................................................................................... ENTRY REQUIREMENTS The Diploma course TEFL has the following entry requirements : Minimum age 18 on entry to the course. Present an exam certificate showing at least a C1 level of english. In case the candidates do not have a certificate, they have to take an exam at the Colombo Americano. The willingness to actively participate in the course and work cooperatively with group members. Dedication to the demanding schedule of classes and practical training required for the course. * We will accept applicants who score lower on the English language proficiency test with the condition that they successfully complete intensive English training before re-sitting the exam. For more information contact with: Deybis Castillo, academic assistant Tel. 687 5800 Ext. 122 dcastillo@colomboamericano.edu.co"
	new4.URLFullImages = []string{"https://www.colomboamericano.edu.co/archivos/6b115b64.jpg"}
	new4.URLMediumImages = []string{"https://www.colomboamericano.edu.co/archivos/6b115b64.jpg"}
	new4.URLSmallImages = []string{"https://www.colomboamericano.edu.co/archivos/6b115b64.jpg"}
	new4.ExtraInfo = ""
	new4.Title = "TEFL CERTIFICATE DATES 2016"
	new4.SutTitle = ""
	new4.Date = ""
	new4.CreatedBy = "Administrative assistant"
	new4.ModifyBy = "Administrative assistant"
	
	new5 := DetailsResult{}
    new5.ShortDescription = ""
	new5.LongDescription = ""
	new5.URLFullImages = []string{""}
	new5.URLMediumImages = []string{""}
	new5.URLSmallImages = []string{"https://www.colomboamericano.edu.co/archivos/eda2e83a.jpg"}
	new5.ExtraInfo = ""
	new5.Title = "Noticia Vacia"
	new5.SutTitle = ""
	new5.Date = ""
	new5.CreatedBy = ""
	new5.ModifyBy = ""


    aSlice := []DetailsResult{new1,new2,new3,new4,new5}
    
   
	// info log 
	beego.Info(aSlice)
	result.News = aSlice
	
	//this.Data["json"] = result
	//	this.ServeJSON()
	
	slice := []interface{}{new1,new2,new3,new4,new5}
 
   s, _ := json.Marshal(slice) 
   this.Ctx.ResponseWriter.Write(s) 
}