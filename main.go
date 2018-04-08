package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
	"github.com/thoratou/organize-jds/controllers"
)

func main() {
	//open settings
	settings, err := controllers.DeserializeCompanyFromJSONFile("./settings.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	controllers.SetSettings(settings)

	//open and feed DB if required
	db, err := bolt.Open("bolt.db", 0600, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatal(err)
		return
	}

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("users"))
		tx.CreateBucketIfNotExists([]byte("teams"))
		tx.CreateBucketIfNotExists([]byte("players"))
		tx.CreateBucketIfNotExists([]byte("forumids"))

		b := tx.Bucket([]byte("games"))
		if b == nil {
			createdBucket, _ := tx.CreateBucket([]byte("games"))
			controllers.CreateTeamGame(createdBucket, "BabyFoot", 2, []string{
				"Matchs de qualification en 2 contre 2",
				"Pas de certificat médical sur cette épreuve",
				"Date limite d'inscription : Samedi 26 mai 2018",
				"Dates des qualifications : Lundi 28, Mardi 29, Mercredi 30, Jeudi 31 mai 2018 de 12h15 à 13h45",
				"Date de la finale : Vendredi 1er Juin 2018 de 12h15 à 13h45",
			})
			controllers.CreateIndividualGame(createdBucket, "Badminton", []string{
				"Certificat médical et /ou licence obligatoire",
				"Date limite d'inscription : Vendredi 1 juin 2018",
				"Date de l'épreuve : Dimanche 3 juin 2018 de 9h à 18h",
			})
			controllers.CreateTeamGame(createdBucket, "Badminton Double", 2, []string{
				"Certificat médical et /ou licence obligatoire",
				"Date limite d'inscription : Vendredi 1 juin 2018",
				"Date de l'épreuve : Dimanche 3 juin 2018 de 9h à 18h",
			})
			controllers.CreateTeamGame(createdBucket, "Basketball", 5, []string{
				"Épreuve mixte mais aucune obligation de présence d’une fille dans le cinq entrant sur le terrain",
				"Certificat médical et /ou licence obligatoire",
				"Date limite d'inscription : Non Communiquée",
				"Dates des qualifications : Lundi 14 Mai, Mardi 15 Mai,  Mercredi 16 Mai, Jeudi 17 Mai et Vendredi 18 Mai de 17h30 à 21h30",
				"Date des phases finales : Mardi 22 Mai (demi-finale), Mercredi 23 Mai de 17h30 à 21h30",
			})
			controllers.CreateTeamGame(createdBucket, "Beach Volley", 3, []string{
				"Tournoi mixte sur sable 3 x 3",
				"Certificat médical et /ou licence obligatoire",
				"Date limite d'inscription : Dimanche 20 mai 2018",
				"Dates des qualifications : Mardi 22, Mercredi 23, Jeudi 24, Vendredi 25, Lundi 28 et Mardi 29 mai de 18h à 21h",
				"Date de la finale : Jeudi 31 mai 2018 de 18h à 21h30",
				"Lieu : Route des 3 moulins, Antibes",
				"Le nombre d'équipes est limité à 96",
			})
			controllers.CreateIndividualGame(createdBucket, "Belote / Jeux de l'Esprit", []string{
				"Notre épreuve vous propose des initiations comme des tournois : belote, awalé, jeu de Go, backgammon",
				"Pas de certificat médical et/ou licence obligatoire ",
				"Date limite d'inscription : Samedi 12 Mai 2018",
				"Epreuves : Lundi 14 Mai et Mercredi 16 Mai de 18h à 21h",
				"Phase Finale : Vendredi 18 Mai de 18h à 21h",
			})
			controllers.CreateTeamGame(createdBucket, "Boules Carrées", 2, []string{
				"Épreuve disputée en doublettes formées (2 joueurs, pas forcément d'une même entreprise/école) avec des boules carrées",
				"Pas de certificat médical sur cette épreuve",
				"Date limite d'inscription : Jeudi 7 Juin 2018",
				"Dates des épreuves : Samedi 9 Juin de 9h-17h",
			})
			controllers.CreateIndividualGame(createdBucket, "Bowling", []string{
				"2 parties (pas de possibilité de rejouer) - Le meilleur score des 2 parties est retenu",
				"Pas de certificat médical",
				"Épreuve spéciale : 10 €",
				"Date limite d’inscription : Vendredi 11 Mai 2018",
				"Quatre séances de Qualifications :",
				"	-   Lundi 14 Mai,     de 17h30 à 20h",
				"	-   Mardi 15 Mai,     de 17h30 à 20h",
				"	-   Jeudi 16 Mai,     de 17h30 à 20h",
				"	-   Mardi 17 Mai,     de 17h30 à 20h",
				"	-   Lundi 4 Juin,     de 17h30 à 20h",
				"	-   Mardi 5 Juin,     de 17h30 à 20h",
				"Date de la finale : Vendredi 8 Juin de 18h00 à 20h",
			})
			controllers.CreateTeamGame(createdBucket, "Course d'Obstacles", 2, []string{
				"Vous avez aimé la Sophia Color Race vous allez adorer la Sophia Défi Race. Après un warm-up Fitness, c’est un parcours de 5 km à parcourir à votre rythme, jalonné d’épreuves (sportives et de réflexion) et d’obstacles (plus ou moins propres, plus ou moins secs !) qui vous attend.",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Mardi 29 Mai 2018",
				"Course : Jeudi 31 Mai 2018 de 17h30 à 20h30",
			})
			controllers.CreateTeamGame(createdBucket, "Course Orientation", 2, []string{
				"Un nouveau terrain de course : 2 parcours pédestres au choix",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Mardi 5 Juin 2018",
				"Attention, épreuve populaire qui peut être complète avant la fin des inscriptions",
				"Date de l'épreuve : Jeudi 7 Juin de 17h à 20h",
			})
			controllers.CreateIndividualGame(createdBucket, "Cross", []string{
				"Cross 2 distances à parcourir : 3 km ou 9 km",
				"Attention, l’épreuve se déroule en même temps que le trail, il n’est donc pas possible de participer aux deux épreuves.",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Samedi 12 mai 2018",
				"Date de l'épreuve : Mardi 15 Mai de 17h à 20h",
			})
			controllers.CreateIndividualGame(createdBucket, "Cross Marche", []string{
				"Marche (classique, active ou nordique) de 3 km",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Non Communiquée",
				"Date de l'épreuve : Non Communiquée",
			})
			controllers.CreateIndividualGame(createdBucket, "Dance Party", []string{
				"Cours collectif, ouvert à tous (hommes & femmes). Nouveau concept super ludique et convivial, le KUDUROFIT® arrive sur la côte d’azur pour les Jeux de Sophia et sera animé  par Fabienne CAMARA.",
				"Cette année l’épreuve de DANSE vous apprendra à danser à 360° avec un groupe, un peu comme un Madison mais sur des rythmes Afro-Caraibéens. Nous terminerons cette Dance Party par des rythmes ZUMBA® avec des tubes de l’été: 1h30 de bonne humeur",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Non Communiquée",
				"Date de l'épreuve : Non Communiquée",
			})
			controllers.CreateTeamGame(createdBucket, "Décathlon", 2, []string{
				"Épreuve « DECATHLON » organisée sur le magasin d’Antibes en partenariat avec les magasins de Cannes-Mandelieu, Grasse et Nice Lingostière.",
				"Le décathlon comprend les disciplines suivantes : Duathlon (enchaînement course à pied, vélo, course à pied) + Tir à l’arc + Fitness + Basketball + Baseball + Sports de raquettes + Épreuve à énigmes (découverte de l’enseigne et des produits)",
				"Certificat médical obligatoire",
				"Date limite d'inscription : Vendredi 8 Juin à 20h",
				"Date de l'épreuve : Samedi 9 juin, de 16h00 à 20h00",
			})
			controllers.CreateIndividualGame(createdBucket, "E-Sport Test", []string{
				"20 équipes Counter Strike de 5 personnes, parties de 60 minutes (maximum)",
				"Pas de certificat médical pour cette épreuve",
				"Date limite d'inscription : Jeudi 7 Juin 2018",
				"Date de l'épreuve : Samedi 9 Juin de 09h00 à 18h00",
			})
			controllers.CreateIndividualGame(createdBucket, "Echecs", []string{
				"Les echecs reviennent en force pour cette nouvelle édition !!!",
				"Pas de certificat médical pour cette épreuve",
				"Date limite d'inscription : Jeudi 31 Mai 2018",
				"Date de l'épreuve : Samedi 2 Juin de 13h30 à 18h",
			})
			controllers.CreateIndividualGame(createdBucket, "Escape Game", []string{
				"Par équipe de 4 ou 5 personnes, parties de 90 minutes",
				"Pas de certificat médical pour cette épreuve",
				"Date limite d'inscription : Non Communiquée",
				"Date de l'épreuve (Qualification) :",
				"	-   Mardi 22 Mai,         de 17h30 à 19h00",
				"	-   Mardi 22 Mai,         de 18h30 à 20h00",
				"	-   Jeudi 24 Mai,         de 17h30 à 19h00",
				"	-   Jeudi 24 Mai,         de 18h30 à 20h00",
				"	-   Lundi 28 Mai,         de 17h30 à 19h00",
				"	-   Lundi 28 Mai,         de 18h30 à 20h00",
				"	-   Mercredi 30 Mai,      de 17h30 à 19h00",
				"	-   Mercredi 30 Mai,      de 18h30 à 20h00",
				"	-   Vendredi 01 Juin,     de 17h30 à 19h00",
				"	-   Vendredi 01 Juin,     de 18h30 à 20h00",
				"	-   Mardi 05 Juin,        de 17h30 à 19h00",
				"	-   Mardi 05 Juin,        de 18h30 à 20h00",
				"Date de la finale : Lundi 11 Juin 18h00 en centre-ville Cannes",
			})
			controllers.CreateTeamGame(createdBucket, "Fléchettes", 2, []string{
				"But : être l'équipe à marquer le plus rapidement 301 ou 501 points, 3 coups par joueurs",
				"Pas de certificat médical",
				"Date limite d'inscription : Lundi 21 mai 2018",
				"Date des qualifications : Mercredi 23 et mardi 29 mai 2018 de 18h à 21h",
				"Date des finales : Mercredi 30 mai de 18h à 21h",
			})
			controllers.CreateTeamGame(createdBucket, "Football à 5", 5, []string{
				"Equipe de 5 sur le terrain avec possibilité d'avoir 2 remplaçants max",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 9€ de supplément par personne",
				"Date limite d'inscription : Samedi 12 Mai 2018",
				"Date des qualifications : Du Lundi 14 Mai au Vendredi 18 Mai 2018 de 12h à 14h",
				"Date des finales :Dimanche 20 Mai 2018 de 10h à 14h.",
			})
			controllers.CreateTeamGame(createdBucket, "Football Féminin", 7, []string{
				"Équipes de 7 joueuses avec possibilité d'avoir 2 joueuses remplaçantes supplémentaires",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Mardi 15 mai 2018",
				"Date des finales : Jeudi 17 mai 2018 de 18h à 22h",
			})
			controllers.CreateTeamGame(createdBucket, "Football Masculin", 7, []string{
				"Equipes de 7 joueurs avec possibilité d'avoir 2 joueurs remplaçants supplémentaires (remplacements illimités)",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Mercredi 16 Mai 2018",
				"Date des qualifications : ",
				"	-   Stade Chabert (685 route de Biot, 06560 Valbonne) : vendredi 18 et mercredi 23 mai 2018 de 18h à 21h30 (session à 18h et à 19h30).",
				"	-   Pierre bel (1 chemin des Combes, 06410 Biot) : samedi 19 de 14h à 20h (session à 14h et à 15h30) et lundi 28 mai 2018 de 18h à 21h30 (session à 18h et à 19h30)",
				"Date de la finales : Samedi 2 juin 2018 de 9h à 15h (stade Pierre Bel)",
			})
			controllers.CreateIndividualGame(createdBucket, "Geocaching", []string{
				"Le Geocaching est une chasse au trésor des temps moderne basée sur la technologie GPS.",
				"Au programme de cette année : la course aux objets voyageurs ! ",
				"<a href=\"https://www.geocaching-jds.fr\">https://www.geocaching-jds.fr</a>",
				"Pas de certificat médical pour cette épreuve",
				"Date limite d'inscription : Samedi 2 juin 2018",
				"Dates des événements  : ",
				"Permanence physique en même temps que la remise des tshirts les mardi 15 et jeudi 17 mai, place Bermond pour répondre à toutes vos questions et expliquer les régles de la course aux objets voyageurs.",
				"L'épreuve dans son ensemble est accessible à tous du 19 mai au 8 juin.",
				"Un événement dédié avec des activités surprise aura lieu le Dimanche 3 juin 2018 de 12h30 à 16h dans un lieu restant à définir.",
				"Site spécifique à l'épreuve : <a href=\"https://www.geocaching-jds.fr\">https://www.geocaching-jds.fr</a>",
			})
			controllers.CreateIndividualGame(createdBucket, "Golf", []string{
				"18 trous Stabl. NET & BRUT par équipes",
				"Classement en NET & BRUT / Classement Société : total des trois meilleurs scores en NET en excluant les débutants / Index maxi 54",
				"Certificat médical ET licence obligatoire",
				"Épreuve spéciale : 63€ de supplément par personne",
				"Règlement et conditions assez complexes, voir directement sur le site : http://www.jeuxdesophia.com/jcms/rda_6289/fr/golf",
				"Date limite d'inscription : Mercredi 16 mai 2018",
				"Dates des épreuves : ",
				"   -   Vendredi 18 mai 2018 : 10 départs de 3 joueurs (toutes les 9 minutes) de 14h à 18h.",
				"   -   Samedi 19 mai 2018 : 27 départs de 3 joueurs (toutes les 9 minutes) de 8h à 18h.",
			})
			controllers.CreateTeamGame(createdBucket, "Handball", 6, []string{
				"Equipe mixte 5 joueurs de champs + 1 gardien - 1 fille sur le terrain tout le match ; sinon 4 joueurs de champs autorisés",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Vendredi 8 Juin 2018",
				"Date unique des finales : Dimanche 10 Juin de 9h à 18h",
			})
			controllers.CreateIndividualGame(createdBucket, "Karting", []string{
				"Compétition de karting sur les hauteurs de la Côte d'Azur",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 14 € de supplément par personne",
				"Date limite d'inscription : Samedi 12 mai 2018",
				"Qualifications : Lundi 14, Mardi 15, Mercredi 16, Lundi 21, Mardi 22, Mercredi 23 Mai 2018 de 18h à 21h",
				"Finale : Lundi 28 juin 2018 de 18h à 21h",
			})
			controllers.CreateTeamGame(createdBucket, "Kayak", 2, []string{
				"Kayak bi-places large et assez stable pour les novices",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 14 € de supplément par personne",
				"Date limite d'inscription : Jeudi 24 Mai 2018",
				"Epreuve : Samedi 26 mai 2018 de 9h à 13h",
			})
			controllers.CreateTeamGame(createdBucket, "Laser Quest", 3, []string{
				"Composition :  Par équipe de 3 personnes, parties de 7 minutes",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 13 € de supplément par personne",
				"Date limite d'inscription : Samedi 12 mai 2018",
				"Date de qualifications : Lundi 14, mercredi 16, lundi 21, mercredi 23, lundi 28 et mercredi 30 mai 2018 de 17h30 à 22h00",
				"Date de la finale : Lundi 4 juin 2018 de 17h30 à 22h00",
			})
			/*controllers.CreateIndividualGame(createdBucket, "Nautathlon", []string{
				"Epreuve sous la forme d'un mini Triathlon avec de la natation, du vélo (statique) et de la course à pied ",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Mercredi 1er juin 2016",
				"Date de l'épreuve : Samedi 4 juin 2016 de 17h à 21h",
			})*/
			controllers.CreateTeamGame(createdBucket, "Padel", 2, []string{
				"4 terrains, 2h de jeu & un maximum de sensations ...",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 10 € de supplément par personne",
				"Date limite d'inscription : Mardi 15 Mai 2018",
				"Date de qualifications : Jeudi 17 Mai, Vendredi 18 Mai, Mardi 22 Mai, Mercredi 23 Mai et Jeudi 24 Mai de 12h à 14h",
				"Date des finales : Lundi 28 Mai de 12h à 14h",
			})
			controllers.CreateTeamGame(createdBucket, "Pétanque", 3, []string{
				"Épreuve disputée en triplettes formées (3 joueurs, pas forcément d'une même entreprise/école)",
				"Pas de certificat médical",
				"Date limite d'inscription : Non Communiquée",
				"Date de l'épreuve : Non Communiquée",
			})
			controllers.CreateIndividualGame(createdBucket, "Photomarathon", []string{
				"4 heures pour réaliser 3 photographies sur 3 thèmes imposés",
				"Pas de certificat médical",
				"Date limite d'inscription : Lundi 4 juin 2018",
				"Date des épreuves : Mercredi 6 juin 2018 de 18h à 20h à Valbonne Sophia Antipolis (Chemin La source, Les Bouillides)",
			})
			controllers.CreateIndividualGame(createdBucket, "Poker", []string{
				"Tournoi de Texas Hold’em No Limit",
				"Pas de certificat médical",
				"Date limite d'inscription : Lundi 21 mai 2018",
				"Qualifications : Mercredi 23, lundi 28, jeudi 31 mai 2018 (demi-finale) de 18h à 23h30",
				"Finale : Vendredi 8 juin 2018 de 18 à 23h30",
			})
			controllers.CreateTeamGame(createdBucket, "Rugby / Tag Rugby", 7, []string{
				"Journée découverte rugby: Rugby à VII (règles loisir) et tag-rugby (sans contact)",
				"Equipes de 7 joueurs + 2 remplaçants max",
				"Les équipes féminines et mixtes sont les bienvenues !",
				"Matchs de 2x7 min sur un terrain de 35 x 50 m",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Mardi 5 Juin 2018. Au delà de cette date, il sera nécessaire d’avoir une confirmation des organisateurs, notamment pour les inscriptions au repas du midi.",
				"Date de l'épreuve : Dimanche 10 Juin 2018 de 9h à 15h au stade de la Fontonne, Antibes",
			})
			controllers.CreateIndividualGame(createdBucket, "Soirée de cloture", []string{
				"<b>Informations 2017, en attente pour 2018 : </b>Soirée privée à l'AzurArena Antibes le jeudi 22 juin 2017.",
				"Entrée OFFERTE à TOUS les participants 2017 ainsi qu'à leurs accompagnants (nombre illimité).",
				"Les inscriptions à la soirée sont ouvertes. N'oubliez pas de vous désinscrire si vous découvrez que vous n'êtes plus disponible !",
				"Date limite d'inscription : Non Communiquée",
				"Ouverture des portes : Non Communiquée",
			})
			controllers.CreateIndividualGame(createdBucket, "Sports Co Decouverte", []string{
				"4 sports proposés : un esprit loisir et non compétition",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Mardi 22 mai 2018",
				"Date de la session : Jeudi 24 mai 2018 de 18 à 22h",
			})
			controllers.CreateIndividualGame(createdBucket, "Squash", []string{
				"Equipe de 2 personnes",
				"1 rencontre entre équipes = 30 minutes",
				"Épreuve spéciale : 5 € de supplément par personne ",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Non Communiquée",
				"Date de l'épreuve : ",
				"	-   Mardi 29 Mai,         de 18h30 à 21h30",
				"	-   Jeudi 31 Mai,         de 18h30 à 21h30",
				"	-   Lundi 4 Juin,         de 18h30 à 21h30",
				"Date de la finale : Jeudi 7 Juin 18h00 à 23h00",
			})
			controllers.CreateTeamGame(createdBucket, "Tennis", 2, []string{
				"Format similaire à la Coupe Davis : Tournoi Equipes de 2 à votre convenance (H/H, F/F, mixte)",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 15 € de supplément par personne ",
				"Date limite des inscriptions : Mardi 22 Mai 2018",
				"Qualifications : Jeudi 24 Mai, Lundi 28 Mai, Lundi 4 Juin et Mardi 5 Juin de 18h à 22h",
				"Finale : Le Samedi 9 Juin de 13h à 20h",
				"Remise des prix : Mardi 12 juin 2018",
			})
			controllers.CreateIndividualGame(createdBucket, "Tennis de Table", []string{
				"Tournois individuels (Féminin, Homme non classé et classé)",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Vendredi 25 Mai 2018",
				"Date de l'épreuve : Dimanche 27 Mai 2018 de 9h à 18h",
				"	-   Tournoi non classé mixte de 9h à 13h (pointage à 8h30)",
				"	-   Tournoi non classé feminin de 14h à 18h (pointage à 13h30)",
				"	-   Tournoi classé mixte de 14h à 18h (pointage 13h30)",
			})
			controllers.CreateIndividualGame(createdBucket, "Tir à l'Arc / Arc'trap", []string{
				"Session de tir à l'arc et Arc'trap (loisir)",
				"Épreuve spéciale : 12 € de supplément par personne",
				"Date limite d’inscription : Lundi 28 mai 2018",
				"Dates des qualifications : 14, 15, 18, 22, 23, 24, 25, 28, 29 mai 2018 de 12h15 à 18h15",
				"Finale : Jeudi 31 mai de 18h15 à 21h",
			})
			controllers.CreateTeamGame(createdBucket, "Tir à l'Arc par Equipe", 4, []string{
				"Session de tir à l'arc et Arc'trap (loisir)",
				"Les équipes sont constituées de 4 hommes ou 4 femmes. Pas de mixité.",
				"Les membres de l'équipe sont composés de 2 membres maximum d'une même entreprise.",
				"Épreuve spéciale : 12 € de supplément par personne / 2ème session Tir à l'arc 10 euros (paiement sur place)",
				"Date limite d’inscription : Lundi 28 mai 2018",
				"Dates des qualifications : 14, 15, 18, 22, 23, 24, 25, 28, 29 mai 2018 de 12h15 à 18h15",
				"Finale : Jeudi 31 mai de 18h15 à 21h",
			})
			controllers.CreateIndividualGame(createdBucket, "Tir Au Pistolet", []string{
				"Pistolet (automatique) ou Revolver (barillet) calibre 22 LR sur cible à 25 mètres",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 20 € de supplément par personne",
				"Date limite d’inscription : Non Communiquée",
				"Dates des qualifications : Non Communiquée",
				"Finale : Non Communiquée",
			})
			controllers.CreateIndividualGame(createdBucket, "Trail", []string{
				"Course à pied nature (Trail) de 12 km pour 350 m de dénivelé positif.",
				"Attention, l’épreuve se déroule en même temps que le cross (3/9 km), il n’est donc pas possible de participer aux deux épreuves.",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d’inscription : Samedi 12 mai 2018",
				"Date de l'épreuve : Mardi 15 mai 2018 de 17h à 20h",
			})
			controllers.CreateTeamGame(createdBucket, "Ultimate Frisbee", 5, []string{
				"Equipes mixtes : 5 joueurs (minimum 1 fille) + 3 remplacants possibles",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d’inscription : Lundi 21 mai 2018",
				"Dates des qualifications :Mercredi 23 et mardi 29 mai 2018 de 18h à 21h",
				"Finale : Dimanche 3 Juin 2018 de 9h à 18h",
			})
			controllers.CreateTeamGame(createdBucket, "Voile", 2, []string{
				"Régate se déroulant sur des catamarans de type Hoby Cat 16, avec équipage de 2 personnes",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 27€ de supplément par personne ",
				"Date limite d’inscription : Vendredi 1er juin 2018",
				"Date de l'épreuve : Dimanche 03 juin 2018 de 9h à 18h",
			})
			controllers.CreateTeamGame(createdBucket, "VolleyBall", 6, []string{
				"Equipe de 6 joueurs (2 remplaçants acceptés)",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d’inscription : Samedi 12 Mai 2018",
				"Dates des qualifications : Lundi 14 Mai, Mercredi 16 Mai et Lundi 21 Mai 2018 de 18h à 22h",
				"Date des finales : Mercredi 23 Mai 2018 de 18h à 22h",
			})
			controllers.CreateIndividualGame(createdBucket, "VTT 15 km", []string{
				"Parcours compétition et parcours loisir encadrés - Catégorie Hommes et Femmes",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d’inscription : Dimanche 20 mai 2018",
				"Date de l'épreuve : Mardi 22 mai 2018 de 17h à 20h",
			})
			controllers.CreateTeamGame(createdBucket, "VTT Nocturne Relais", 2, []string{
				"Épreuve VTT en relai sur une boucle de 4 km",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d’inscription : Dimanche 27 mai 2018",
				"Date de l'épreuve : Mardi 29 mai 2018 de 21h à 23h",
			})
			/*controllers.CreateIndividualGame(createdBucket, "Warm'Up des Jeux ", []string{
				"Warm'Up tous les mardis et jeudis sur les lieux de vie de Sophia",
				"Animations gratuites en musique (badminton, mini tennis de table, tir à l'arc, mini golf, football, boules carrées ...) et distribution de flyers avec planning des épreuves.",
				"Dates de l'evenement :",
				"	-   Mardi 2 mai - 12:00 à 13:30	Rendez-vous sur la place Bermond (Sophia - au dessus du CIV) entre 12:00 et 13:30 pour mettre l'ambiance et tester les produits mis à disposition par notre partenaire Décathlon Antibes",
				"	-   Jeudi 4 mai - 12:00 à 13:30	Rendez-vous sur la place Méjane, quartier Garbejaire. On vous attend nombreux !",
				"	-   Mardi 9 mai - 12:00 à 13:30	Spécial étudiants : Campus Sophia Tech (Polytech)",
				"	-   Jeudi 11 mai - 12:00 à 13:30	Espace Saint Philippe",
				"	-   Mardi 16 mai - 12:00 à 13:30	Place Sophie Laffitte",
				"	-   Jeudi 18 mai - 12:00 à 13:30	Place Bermond",
				"	-   Mardi 23 mai - 12:00 à 13:30	Espace Saint Philippe",
				"	-   Jeudi 25 mai	Férié",
				"	-   Mardi 30 mai - 12:00 à 13:30	 Place Méjane, quartier Garbejaire",
			})*/
		}
		return nil
	})

	controllers.SetDB(db)

	defer db.Close()

	rand.Seed(time.Now().UnixNano())

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/signin", &controllers.HomeController{}, "post:SignInQuery")
	beego.Router("/signup", &controllers.HomeController{}, "post:SignUpQuery")

	beego.Router("/games/*", &controllers.DataController{}, "get:Get")

	beego.Router("/submitPlayerData", &controllers.DataController{}, "post:SubmitPlayerData")

	beego.Router("/addPlayerToGame", &controllers.DataController{}, "post:AddPlayerToGame")
	beego.Router("/removePlayerFromGame", &controllers.DataController{}, "post:RemovePlayerFromGame")
	beego.Router("/submitPlayerGameComment", &controllers.DataController{}, "post:SubmitPlayerGameComment")

	beego.Router("/submitGameNewPost", &controllers.DataController{}, "post:SubmitGameNewPost")
	beego.Router("/submitGameModifyPost", &controllers.DataController{}, "post:SubmitGameModifyPost")
	beego.Router("/submitGameDeletePost", &controllers.DataController{}, "post:SubmitGameDeletePost")
	beego.Router("/restoreGamePost", &controllers.DataController{}, "post:RestoreGamePost")

	beego.Router("/addTeamToGame", &controllers.DataController{}, "post:AddTeamToGame")
	beego.Router("/removeTeamFromGame", &controllers.DataController{}, "post:RemoveTeamFromGame")
	beego.Router("/changeTeamName", &controllers.DataController{}, "post:ChangeTeamName")
	beego.Router("/changeManager", &controllers.DataController{}, "post:ChangeManager")
	beego.Router("/submitTeamComment", &controllers.DataController{}, "post:SubmitTeamComment")

	beego.Router("/submitTeamNewPost", &controllers.DataController{}, "post:SubmitTeamNewPost")
	beego.Router("/submitTeamModifyPost", &controllers.DataController{}, "post:SubmitTeamModifyPost")
	beego.Router("/submitTeamDeletePost", &controllers.DataController{}, "post:SubmitTeamDeletePost")
	beego.Router("/restoreTeamPost", &controllers.DataController{}, "post:RestoreTeamPost")

	beego.Router("/addPlayerToTeam", &controllers.DataController{}, "post:AddPlayerToTeam")
	beego.Router("/removePlayerFromTeam", &controllers.DataController{}, "post:RemovePlayerFromTeam")
	beego.Router("/submitPlayerTeamComment", &controllers.DataController{}, "post:SubmitPlayerTeamComment")

	beego.Run()
}
