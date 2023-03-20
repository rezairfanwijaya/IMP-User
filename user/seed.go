package user

import "imp/helper"

func GenerateSeedUser() []User {
	agungPassword, _ := helper.HashingPassword("12345")
	bagusPassword, _ := helper.HashingPassword("12345")
	cacaPassword, _ := helper.HashingPassword("12345")
	dadanPassword, _ := helper.HashingPassword("12345")
	evanPassword, _ := helper.HashingPassword("12345")
	farizPassword, _ := helper.HashingPassword("12345")
	galuhPassword, _ := helper.HashingPassword("12345")
	hanaPassword, _ := helper.HashingPassword("12345")
	irenePassword, _ := helper.HashingPassword("12345")
	jazkiaPassword, _ := helper.HashingPassword("12345")
	kevinPassword, _ := helper.HashingPassword("12345")
	lalaPassword, _ := helper.HashingPassword("12345")
	marioPassword, _ := helper.HashingPassword("12345")
	nauvalPassword, _ := helper.HashingPassword("12345")
	opikPassword, _ := helper.HashingPassword("12345")
	pernandesPassword, _ := helper.HashingPassword("12345")
	qiantisPassword, _ := helper.HashingPassword("12345")
	raisaPassword, _ := helper.HashingPassword("12345")
	samuelPassword, _ := helper.HashingPassword("12345")
	teguhPassword, _ := helper.HashingPassword("12345")

	users := []User{
		{
			Username: "agung",
			Password: agungPassword,
			FullName: "agung permana",
		}, {
			Username: "bagus",
			Password: bagusPassword,
			FullName: "bagus pratama",
		}, {
			Username: "caca",
			Password: cacaPassword,
			FullName: "caca pertiwi",
		}, {
			Username: "dadan",
			Password: dadanPassword,
			FullName: "dadan permadana",
		}, {
			Username: "evan",
			Password: evanPassword,
			FullName: "evan evin",
		}, {
			Username: "fariz",
			Password: farizPassword,
			FullName: "fariz galung",
		}, {
			Username: "galuh",
			Password: galuhPassword,
			FullName: "galuh permata",
		}, {
			Username: "hana",
			Password: hanaPassword,
			FullName: "hana haho",
		}, {
			Username: "irene",
			Password: irenePassword,
			FullName: "irene oktavia",
		}, {
			Username: "jazkia",
			Password: jazkiaPassword,
			FullName: "jazkia kia",
		}, {
			Username: "kevin",
			Password: kevinPassword,
			FullName: "kevin chan",
		}, {
			Username: "lala",
			Password: lalaPassword,
			FullName: "lala lili",
		}, {
			Username: "mario",
			Password: marioPassword,
			FullName: "mario utama",
		}, {
			Username: "nauval",
			Password: nauvalPassword,
			FullName: "nauval iza",
		}, {
			Username: "opik",
			Password: opikPassword,
			FullName: "opik kresnantya",
		}, {
			Username: "pernandes",
			Password: pernandesPassword,
			FullName: "pernandes diaz",
		}, {
			Username: "qianti",
			Password: qiantisPassword,
			FullName: "qianti putri",
		}, {
			Username: "raisa",
			Password: raisaPassword,
			FullName: "raisa putri",
		}, {
			Username: "samuel",
			Password: samuelPassword,
			FullName: "samuel putra",
		}, {
			Username: "teguh",
			Password: teguhPassword,
			FullName: "teguh samudra",
		},
	}

	return users
}
