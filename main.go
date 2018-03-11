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

		b := tx.Bucket([]byte("games"))
		if b == nil {
			createdBucket, _ := tx.CreateBucket([]byte("games"))
			controllers.CreateTeamGame(createdBucket, "BabyFoot", 2, []string{
				"Matchs de qualification en 2 contre 2",
				"Pas de certificat médical sur cette épreuve",
				"Date limite d'inscription : Mardi 13 juin 2017",
				"Dates des qualifications : Lundi 12, Mardi 13, Mercredi 14, Jeudi 15 Juin 2017 de 12h15 à 13h45",
				"Date de la finale : Vendredi 16 Juin 2017 de 12h15 à 13h45",
			})
			controllers.CreateIndividualGame(createdBucket, "Badminton", []string{
				"Certificat médical et /ou licence obligatoire",
				"Date limite d'inscription : Jeudi 15 Juin 2017",
				"Date de l'épreuve : Samedi 17 Juin 2017 de 9h à 18h",
			})
			controllers.CreateTeamGame(createdBucket, "Badminton Double", 2, []string{
				"Certificat médical et /ou licence obligatoire",
				"Date limite d'inscription : Jeudi 15 Juin 2017",
				"Date de l'épreuve : Samedi 17 Juin 2017 de 9h à 18h",
			})
			controllers.CreateTeamGame(createdBucket, "Basketball", 5, []string{
				"Épreuve mixte mais aucune obligation de présence d’une fille dans le cinq entrant sur le terrain",
				"Certificat médical et /ou licence obligatoire",
				"Date limite d'inscription : Samedi 27 Mai 2017",
				"Dates des qualifications : Lundi 29, Mardi 30,  Mercredi 31 Mai, Jeudi 01, Vendredi 02, Mardi 06 juin de 17h30 à 21h30",
				"Date des phases finales : Mardi 06, Mercredi 07 juin de 18h à 21h30",
			})
			controllers.CreateTeamGame(createdBucket, "Beach Volley", 3, []string{
				"Tournoi mixte sur sable 3 x 3",
				"Certificat médical et /ou licence obligatoire",
				"Date limite d'inscription : Dimanche 4 Juin 2017",
				"Dates des qualifications : mardi 6, mercredi 7, jeudi 8, vendredi 9, lundi 12, mardi 13 de 18h-21h",
				"Date de la finale : jeudi 15 juin 2017 de 18h à 21h30",
			})
			controllers.CreateTeamGame(createdBucket, "Boules Carrées", 2, []string{
				"Épreuve disputée en doublettes formées (2 joueurs, pas forcément d'une même entreprise/école) avec des boules carrées",
				"Pas de certificat médical sur cette épreuve",
				"Date limite d'inscription : Jeudi 15 Juin 2017",
				"Dates des épreuves : Samedi 17 Juin de 9h-14h",
			})
			controllers.CreateIndividualGame(createdBucket, "Bowling", []string{
				"2 parties (pas de possibilité de rejouer) - Le meilleur score des 2 parties est retenu",
				"Pas de certificat médical",
				"Épreuve spéciale : 10 €",
				"Date limite d’inscription : Jeudi 27 mai 2017",
				"Quatre séances de Qualifications :",
				"	-   Lundi 29 mai,     de 17h30 à 20h",
				"	-   Mardi 30 mai,     de 17h30 à 20h",
				"	-   Jeudi 1 Juin,     de 17h30 à 20h",
				"	-   Mardi 6 Juin,     de 17h30 à 20h",
				"Date de la finale : Vendredi  9 Juin de 18h00 à 20h",
			})
			controllers.CreateTeamGame(createdBucket, "Course d'Obstacles", 2, []string{
				"Vous avez aimé la Sophia Color Race vous allez adorer la Sophia Défi Race. Après un warm-up Fitness, c’est un parcours de 5 km à parcourir à votre rythme, jalonné d’épreuves (sportives et de réflexion) et d’obstacles (plus ou moins propres, plus ou moins secs !) qui vous attend.",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : 13 Juin 2017",
				"Course : 15 Juin 2017 de 17h30 à 20h",
			})
			controllers.CreateTeamGame(createdBucket, "Course Orientation", 2, []string{
				"Un nouveau terrain de course : 2 parcours pédestres au choix",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Mardi 6 Juin 2017",
				"Attention, épreuve populaire qui peut être complète avant la fin des inscriptions",
				"Date de l'épreuve : Jeudi 8 Juin de 17h à 20h",
			})
			controllers.CreateIndividualGame(createdBucket, "Cross", []string{
				"Cross 2 distances à parcourir : 3 km ou 9 km",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Mardi 30 mai 2017",
				"Date de l'épreuve : Jeudi 1 Juin de 17h à 20h",
			})
			controllers.CreateIndividualGame(createdBucket, "Cross Marche", []string{
				"Marche (classique, active ou nordique) de 3 km",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Mardi 30 mai 2017",
				"Date de l'épreuve : Jeudi 1 Juin de 17h à 20h",
			})
			controllers.CreateIndividualGame(createdBucket, "Dance Party", []string{
				"Cours collectif, ouvert à tous (hommes & femmes). Nouveau concept super ludique et convivial, le KUDUROFIT® arrive sur la côte d’azur pour les Jeux de Sophia et sera animé  par Fabienne CAMARA.",
				"Cette année l’épreuve de DANSE vous apprendra à danser à 360° avec un groupe, un peu comme un Madison mais sur des rythmes Afro-Caraibéens. Nous terminerons cette Dance Party par des rythmes ZUMBA® avec des tubes de l’été: 1h30 de bonne humeur",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Mercredi 31 mai 2017",
				"Date de l'épreuve : Vendredi 02 juin de 18h à 20h",
			})
			controllers.CreateTeamGame(createdBucket, "Décathlon", 2, []string{
				"Épreuve « DECATHLON » organisée sur le magasin d’Antibes en partenariat avec les magasins de Cannes-Mandelieu, Grasse et Nice Lingostière.",
				"Le décathlon comprend les disciplines suivantes : Duathlon (enchaînement course à pied, vélo, course à pied) + Tir à l’arc + Fitness + Basketball + Baseball + Sports de raquettes + Épreuve à énigmes (découverte de l’enseigne et des produits)",
				"Certificat médical obligatoire",
				"Date limite d'inscription : Vendredi 16 Juin à 20h",
				"Date de l'épreuve : Samedi 17 juin, de 14h00 à 18h00",
			})
			controllers.CreateIndividualGame(createdBucket, "Echecs", []string{
				"Les echecs reviennent en force pour cette nouvelle édition !!!",
				"Pas de certificat médical pour cette épreuve",
				"Date limite d'inscription : Jeudi 15 Juin 2017",
				"Date de l'épreuve : Samedi 17 Juin de 13h30 à 18h",
			})
			controllers.CreateTeamGame(createdBucket, "Fléchettes", 2, []string{
				"But : être l'équipe à marquer le plus rapidement 301 ou 501 points, 3 coups par joueurs",
				"Pas de certificat médical",
				"Date limite d'inscription : Dimanche 11 juin 2017",
				"Date des qualifications : Mardi 13 juin de 18h à 21h",
				"Date des finales : Mercredi 14 juin de 18h à 21h",
			})
			controllers.CreateTeamGame(createdBucket, "Football à 5", 5, []string{
				"Equipe de 5 sur le terrain avec possibilité d'avoir 2 remplaçants max",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 9€ de supplément par personne",
				"Date limite d'inscription : Samedi 27 Mai 2017",
				"Date des qualifications : Du Lundi 29 Mai au Vendredi 2 Juin 2017 de 12h à 14h",
				"Date des finales :Dimanche 4 Juin 2017 de 10h à 14h.",
			})
			controllers.CreateTeamGame(createdBucket, "Football Féminin", 7, []string{
				"Équipes de 7 joueuses avec possibilité d'avoir 2 joueuses remplaçantes supplémentaires",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Jeudi 08 juin 2017",
				"Date des finales : Samedi 10 juin de 9h à 18h",
			})
			controllers.CreateTeamGame(createdBucket, "Football Masculin", 7, []string{
				"Equipes de 7 joueurs avec possibilité d'avoir 2 joueurs remplaçants supplémentaires (remplacements illimités)",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Lundi 29 Mai 2017",
				"Date des qualifications : ",
				"	-   Stade Chabert, 685 route de Biot, 06560 Valbonne (Mercredi 31 mai et Vendredi 2 Juin de 18-21h30)",
				"	-   Pierre bel (Mardi 6 Juin et Vendredi 9 Juin de 18-21h30)",
				"Date des finales : Dimanche 11 Juin de 9h à 15h (stade Pierre Bel à Biot)",
			})
			controllers.CreateIndividualGame(createdBucket, "Geocaching", []string{
				"Le Geocaching est une chasse au trésor des temps moderne basée sur la technologie GPS.",
				"Pas de certificat médical pour cette épreuve",
				"Date limite d'inscription : Vendredi 9 juin 2017",
				"Dates des événements  : L'épreuve dans son ensemble est accessible à tous du 29 mai au 22 juin. Des événements spécifiques auront lieu le Dimanche 11 juin 2017 de 12h à 17h",
				"Site spécifique à l'épreuve : Les événements auront lieu dans le Parc des Bouillides (l'emplacement GPS précis vous sera communiqué après votre inscription)",
			})
			controllers.CreateIndividualGame(createdBucket, "Golf", []string{
				"18 trous Stabl. NET & BRUT par équipes",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 63€ de supplément par personne",
				"Règlement et conditions assez complexes, voir directement sur le site : http://www.jeuxdesophia.com/jcms/rda_6289/fr/golf",
				"Date limite d'inscription : Mercredi 31 mai 2017",
				"Dates des épreuves  : Vendredi 02 juin de 14h à 18h et Samedi 03 juin de 8h à 18h ",
			})
			controllers.CreateTeamGame(createdBucket, "Handball", 6, []string{
				"Equipe mixte 5 joueurs de champs + 1 gardien - 1 fille sur le terrain tout le match ; sinon 4 joueurs de champs autorisés",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Vendredi 9 Juin 2017",
				"Date unique des finales ;: Dimanche 11 Juin de 9h à 18h",
			})
			controllers.CreateIndividualGame(createdBucket, "Jeux de l'Esprit", []string{
				"Notre épreuve vous propose des initiations comme des tournois : belote, awalé, jeu de Go, backgammon",
				"Pas de certificat médical et/ou licence obligatoire ",
				"Date limite d'inscription : Samedi 27 mai 2017",
				"Epreuves : Lundi 29 et Mercredi 31 Mai  de 18h à 21h",
				"Phase Finale : Vendredi 2 juin de 18h à 21h",
			})
			controllers.CreateIndividualGame(createdBucket, "Karting", []string{
				"Compétition de karting sur les hauteurs de la Côte d'Azur",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 14 € de supplément par personne",
				"Date limite d'inscription : Samedi 27 mai 2017",
				"Qualifications : Du lundi 29, mardi 30, mercredi 31 Mai et lundi 05, Mardi 06, Mercredi 07 juin 2017 de 18h à 21h",
				"Finale : Lundi 12 juin 2017 de 18h à 21h",
			})
			controllers.CreateTeamGame(createdBucket, "Kayak", 2, []string{
				"Kayak bi-places large et assez stable pour les novices",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 13 € de supplément par personne",
				"Date limite d'inscription : Jeudi 1 Juin 2017",
				"Epreuve : Samedi 3 juin 2017 de 9h à 13h",
			})
			controllers.CreateTeamGame(createdBucket, "Laser Quest", 3, []string{
				"Composition :  Par équipe de 3 personnes, parties de 7 minutes",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 13 € de supplément par personne",
				"Date limite d'inscription : Samedi 27 mai",
				"Date de qualifications : Lundi  29 mai, jeudi 1 Juin, mardi 6 Juin, jeudi 8 Juin et lundi 12 juin 2017 de 17h30 à 22h",
				"Finale : Vendredi 16 juin 2017 de 17h30 à 22h",
			})
			controllers.CreateIndividualGame(createdBucket, "Nautathlon", []string{
				"Epreuve sous la forme d'un mini Triathlon avec de la natation, du vélo (statique) et de la course à pied ",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Mercredi 1er juin 2016",
				"Date de l'épreuve : Samedi 4 juin 2016 de 17h à 21h",
			})
			controllers.CreateTeamGame(createdBucket, "Padel", 2, []string{
				"4 terrains, 2h de jeu & un maximum de sensations ...",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 10 € de supplément par personne",
				"Date limite d'inscription : Mercredi 31 Mai 2017",
				"Date de qualifications : Vendredi 2 Juin, lundi 5 juin, mardi 6 juin, mercredi 7 juin, jeudi 08 juin de 12h à 14h",
				"Date des finales : Lundi 12 juin de 12h à 14h",
			})
			controllers.CreateTeamGame(createdBucket, "Pétanque", 3, []string{
				"Épreuve disputée en triplettes formées (3 joueurs, pas forcément d'une même entreprise/école)",
				"Pas de certificat médical",
				"Date limite d'inscription : Jeudi 8 Juin 2017",
				"Date de l'épreuve : Samedi 10 Juin 2017 de 9h à 22h",
			})
			controllers.CreateIndividualGame(createdBucket, "Photomarathon", []string{
				"4 heures pour réaliser 3 photographies sur 3 thèmes imposés",
				"Pas de certificat médical",
				"Date limite d'inscription : Mardi 06 juin 2017",
				"Date des épreuves : Jeudi 8 juin 2017",
				"	-   de 10h à 12h à Antibes",
				"	-   de 12h30 h à 17h à Biot",
			})
			controllers.CreateIndividualGame(createdBucket, "Poker", []string{
				"Tournoi de Texas Hold’em No Limit",
				"Pas de certificat médical",
				"Date limite d'inscription : Dimanche 04 juin 2017",
				"Qualifications : Mardi 06, jeudi 08, lundi 12 Juin 2017 de 18h à 23h30",
				"Finale : Jeudi 15 juin 2017 de 18h00 à 23h30",
			})
			controllers.CreateTeamGame(createdBucket, "Rugby a 7", 7, []string{
				"Equipes de 7 joueurs + 2 remplaçants max",
				"Les équipes féminines et mixtes sont les bienvenues !",
				"Matchs de 2x7 min sur un terrain de 35 x 50 m",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Vendredi 9 juin 2017. Au-delà de cette date, il sera nécessaire d’avoir une confirmation des organisateurs, notamment pour les inscriptions au repas du midi.",
				"Date de l'épreuve : Dimanche 11 juin 2017 de 9h à 15h",
			})
			controllers.CreateIndividualGame(createdBucket, "Soirée de cloture", []string{
				"Soirée privée à l'AzurArena Antibes le jeudi 22 juin 2017.",
				"Entrée OFFERTE à TOUS les participants 2017 ainsi qu'à leurs accompagnants (nombre illimité).",
				"Les inscriptions à la soirée sont ouvertes. N'oubliez pas de vous désinscrire si vous découvrez que vous n'êtes plus disponible !",
				"Date limite d'inscription : Lundi 20 juin 2016 au soir sur le site des Jeux de Sophia",
				"Ouverture des portes : Jeudi 23 juin 2016 à 18h30",
			})
			controllers.CreateIndividualGame(createdBucket, "Sports Co Decouverte", []string{
				"4 sports proposés : un esprit loisir et non compétition",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : Mardi 6 juin 2017",
				"Date de la session : Jeudi 8 juin 2017 de 18 à 22h",
			})
			controllers.CreateTeamGame(createdBucket, "Tennis", 2, []string{
				"Format similaire à la Coupe Davis : Tournoi Equipes de 2 à votre convenance (H/H, F/F, mixte)",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 15 € de supplément par personne ",
				"Date limite des inscriptions : Dimanche 28 mai 2017",
				"Qualifications : Mardi 30 mai, Jeudi 1er Juin, Mardi 06 Juin et Jeudi 8 Juin 2017 de 18h à 22h",
				"Finale : Le Samedi 10 Juin de 13h à 20h",
				"Remise des prix : Mardi 13 juin 2017",
			})
			controllers.CreateIndividualGame(createdBucket, "Tennis de Table", []string{
				"Tournois individuels (Féminin, Homme non classé et classé)",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d'inscription : jeudi 1er juin 2017",
				"Date de l'épreuve : Samedi 3 juin 2017 de 9h à 18h",
				"	-   Tournoi non classé mixte de 9h à 13h (pointage à 8h30)",
				"	-   Tournoi non classé feminin de 14h à 18h (pointage à 13h30)",
				"	-   Tournoi classé mixte de 14h à 18h (pointage 13h30)",
			})
			controllers.CreateIndividualGame(createdBucket, "Tir à l'Arc", []string{
				"Session de tir à l'arc et Arc'trap (loisir)",
				"Épreuve spéciale : 11 € de supplément par personne",
				"Date limite d’inscription : Samedi 27 juin",
				"Dates des qualifications : 29, 30, 31 mai et 02, 12, 14 Juin 2017 de 12h15 à 18h15",
				"Finale : Vendredi 16 juin de 18h à 21h",
			})
			controllers.CreateTeamGame(createdBucket, "Tir à l'Arc par Equipe", 4, []string{
				"Session de tir à l'arc et Arc'trap (loisir)",
				"Les équipes sont constituées de 4 hommes ou 4 femmes. Pas de mixité.",
				"Les membres de l'équipe sont composés de 2 membres maximum d'une même entreprise.",
				"Épreuve spéciale : 11 € de supplément par personne / 2ème session Tir à l'arc 10 euros (paiement sur place)",
				"Date limite d’inscription : Samedi 27 juin",
				"Dates des qualifications : 29, 30, 31 mai et 02, 12, 14 Juin 2017 de 12h15 à 18h15",
				"Finale : Vendredi 16 juin de 18h à 21h",
			})
			controllers.CreateIndividualGame(createdBucket, "Tir Au Pistolet", []string{
				"Pistolet (automatique) ou Revolver (barillet) calibre 22 LR sur cible à 25 mètres",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 20 € de supplément par personne",
				"Date limite d’inscription : Samedi 27 mai 2017",
				"Dates des qualifications : Lundi 29, mardi 30, jeudi 01, vendredi 2, lundi 12 Juin 2017 de 18h à 20h30",
				"Finale : Mardi 13 Juin de 18h à 21h",
			})
			controllers.CreateTeamGame(createdBucket, "Ultimate Frisbee", 5, []string{
				"Equipes mixtes : 5 joueurs (minimum 1 fille) + 3 remplacants possibles",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d’inscription : Lundi 29 mai 2017",
				"Dates des qualifications : Mercredi 31 mai, Lundi 5 et Mercredi 7 Juin 2017 de 18h à 21h",
				"Finale : Dimanche 11 Juin 2017 de 9h à 18h",
			})
			controllers.CreateTeamGame(createdBucket, "Voile", 2, []string{
				"Régate se déroulant sur des catamarans de type Hoby Cat 16, avec équipage de 2 personnes",
				"Certificat médical et/ou licence obligatoire",
				"Épreuve spéciale : 27€ de supplément par personne ",
				"Date limite d’inscription : Jeudi 01 juin 2017",
				"Date de l'épreuve : Dimanche 04 juin 2017 de 9h à 18h",
			})
			controllers.CreateTeamGame(createdBucket, "VolleyBall", 6, []string{
				"Equipe de 6 joueurs (2 remplaçants acceptés)",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d’inscription : Lundi 29 mai 2017",
				"Dates des qualifications : Mardi 30 Mai et mercredi 31 Mai 2017 de 18h à 22h",
				"Date des finales : Mardi 06 Juin de 18h à 22h",
			})
			controllers.CreateIndividualGame(createdBucket, "VTT 15 km", []string{
				"Parcours compétition et parcours loisir encadrés - Catégorie Hommes et Femmes",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d’inscription : Dimanche  11 juin 2017",
				"Date de l'épreuve : Mardi 13 juin 2017 de 17h à 20h",
			})
			controllers.CreateTeamGame(createdBucket, "VTT Nocturne Relais", 2, []string{
				"Épreuve VTT en relai sur une boucle de 4 km",
				"Certificat médical et/ou licence obligatoire",
				"Date limite d’inscription : Mardi 6 juin 2017",
				"Date de l'épreuve : Jeudi 8 juin 2017 de 21h à 23h",
			})
			controllers.CreateIndividualGame(createdBucket, "Warm'Up des Jeux ", []string{
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
			})
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

	beego.Router("/addTeamToGame", &controllers.DataController{}, "post:AddTeamToGame")
	beego.Router("/removeTeamFromGame", &controllers.DataController{}, "post:RemoveTeamFromGame")
	beego.Router("/changeTeamName", &controllers.DataController{}, "post:ChangeTeamName")
	beego.Router("/changeManager", &controllers.DataController{}, "post:ChangeManager")
	beego.Router("/submitTeamComment", &controllers.DataController{}, "post:SubmitTeamComment")

	beego.Router("/addPlayerToTeam", &controllers.DataController{}, "post:AddPlayerToTeam")
	beego.Router("/removePlayerFromTeam", &controllers.DataController{}, "post:RemovePlayerFromTeam")
	beego.Router("/submitPlayerTeamComment", &controllers.DataController{}, "post:SubmitPlayerTeamComment")

	beego.Run()
}
