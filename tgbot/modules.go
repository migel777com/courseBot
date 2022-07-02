package tgbot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleModule(chatID, module int64, message string, server Server) {

	switch module {
	case 1:
		ExecuteModule1(chatID, message, server)
	case 2:
		ExecuteModule2(chatID, message, server)
	case 3:
		ExecuteModule3(chatID, message, server)
	case 5:
		ExecuteModule5(chatID, message, server)
	case 6:
		ExecuteModule6(chatID, message, server)
	case 7:
		ExecuteModule7(chatID, message, server)
	case 8:
		ExecuteModule8(chatID, message, server)
	case 9:
		ExecuteModule9(chatID, message, server)
	case 10:
		ExecuteModule10(chatID, message, server)
	case 11:
		ExecuteModule11(chatID, message, server)
	case 12:
		ExecuteModule12(chatID, message, server)
	case 13:
		ExecuteModule13(chatID, message, server)
	case 14:
		ExecuteModule14(chatID, message, server)
	case 16:
		ExecuteModule16(chatID, message, server)
	case 17:
		ExecuteModule17(chatID, message, server)
	default:
		msg := tgbotapi.NewMessage(chatID, "Курс пройден")
		msg.ReplyMarkup = mainMenu
		_, _ = server.Bot.Send(msg)
	}

}

func HandleStart(chatID, module int64, server Server) {

	switch module {
	case 2:
		StartModule2(chatID, server)
	case 3:
		StartModule3(chatID, server)
	case 4:
		StartModule4(chatID, server)
	case 5:
		StartModule5(chatID, server)
	case 6:
		StartModule6(chatID, server)
	case 7:
		StartModule7(chatID, server)
	case 8:
		StartModule8(chatID, server)
	case 9:
		StartModule9(chatID, server)
	case 10:
		StartModule10(chatID, server)
	case 11:
		StartModule11(chatID, server)
	case 12:
		StartModule12(chatID, server)
	case 13:
		StartModule13(chatID, server)
	case 14:
		StartModule14(chatID, server)
	case 15:
		StartModule15(chatID, server)
	case 16:
		StartModule16(chatID, server)
	case 17:
		StartModule17(chatID, server)
	}
}

func HandleGet(chatID, module int64, server Server) {

	switch module {
	case 1:
		StartModule1(chatID, server)
		ExecuteModule1(chatID, " ", server)
	case 2:
		StartModule2(chatID, server)
		ExecuteModule2(chatID, " ", server)
	case 3:
		StartModule3(chatID, server)
		ExecuteModule3(chatID, " ", server)
	case 4:
		GetModule4(chatID, server)
	case 5:
		StartModule5(chatID, server)
		ExecuteModule5(chatID, " ", server)
	case 6:
		StartModule6(chatID, server)
		ExecuteModule6(chatID, " ", server)
	case 7:
		StartModule7(chatID, server)
		ExecuteModule7(chatID, " ", server)
	case 8:
		StartModule8(chatID, server)
		ExecuteModule8(chatID, " ", server)
	case 9:
		StartModule9(chatID, server)
		ExecuteModule9(chatID, " ", server)
	case 10:
		StartModule10(chatID, server)
		ExecuteModule10(chatID, " ", server)
	case 11:
		StartModule11(chatID, server)
		ExecuteModule11(chatID, " ", server)
	case 12:
		StartModule12(chatID, server)
		ExecuteModule12(chatID, " ", server)
	case 13:
		StartModule13(chatID, server)
		ExecuteModule13(chatID, " ", server)
	case 14:
		StartModule14(chatID, server)
		ExecuteModule14(chatID, " ", server)
	case 15:
		GetModule15(chatID, server)
	case 16:
		StartModule16(chatID, server)
		ExecuteModule16(chatID, " ", server)
	case 17:
		StartModule17(chatID, server)
		ExecuteModule17(chatID, " ", server)
	}
}

func StartModule(chatID int64, server Server) {
	text := "Модуль 1. Эффективности можно научиться"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = ""
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = ""
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Домашнее задание: "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "1-й вопрос: “Считаете ли вы себя эффективным руководителем?”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule(chatID int64, message string, server Server) {
	ans1, ans2 := server.Db.GetTwoAnswers(chatID, 1)

	fmt.Println(ans1, ans2)

	if ans1 == "empty" && ans2 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans1_1", message)

		text := "2-ой вопрос: “Что бы вы хотели улучшить в себе или команде?”"
		msg := tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

	} else if ans1 != "empty" && ans2 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans1_2", message)

		msg := tgbotapi.NewMessage(chatID, "Вы выполнили все задания")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule1(chatID int64, server Server) {
	text := "Модуль 1. Эффективности можно научиться"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "“Работа управляющего состоит в том, чтобы быть эффективным. Чем бы он " +
		"ни занимался — бизнесом или работал в больнице, в правительственном " +
		"учреждении или в профсоюзном комитете, университете или в армейском " +
		"подразделении — от него требуется правильное выполнение задач, то есть " +
		"ожидают проявления эффективности”."
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Эффективность — особая технология, применяемая работником " +
		"умственного труда в рамках организации. Долгое время наиболее важным в " +
		"деятельности организации считалась результативность работника " +
		"физического труда по выполнению поставленных задач (без оценки " +
		"адекватности самой задачи). Работники интеллектуального труда были в " +
		"меньшинстве. Сейчас положение дел кардинально изменилось. " +
		"Управляющий — человек, который в силу занимаемой им должности или " +
		"имеющихся знаний отвечает за деятельность, которая непосредственно " +
		"влияет на способность организации функционировать и добиватьс " +
		"результатов. К сожалению, многих талантливых и обладающих творческим " +
		"воображением руководителей нельзя назвать управляющими. Такие люди " +
		"часто не понимают, что благодаря одним лишь способностям, без должной " +
		"целеустремленности невозможно добиться эффективности. Существует и " +
		"другая крайность — высокоэффективные сотрудники, не имеющие особых " +
		"талантов, которые медленно и целеустремленно движутся к намеченной цели."
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Домашнее задание: "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "1-й вопрос: “Считаете ли вы себя эффективным руководителем?”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule1(chatID int64, message string, server Server) {
	ans1, ans2 := server.Db.GetTwoAnswers(chatID, 1)

	fmt.Println(ans1, ans2)

	if ans1 == "empty" && ans2 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans1_1", message)

		text := "2-ой вопрос: “Что бы вы хотели улучшить в себе или команде?”"
		msg := tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

	} else if ans1 != "empty" && ans2 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans1_2", message)

		msg := tgbotapi.NewMessage(chatID, "Вы выполнили все задания")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)
	} else if ans1 != "empty" && ans2 != "empty" {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)

		text := "2-ой вопрос: “Что бы вы хотели улучшить в себе или команде?”"
		msg = tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

		msg = tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans2)
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule2(chatID int64, server Server) {
	text := "Модуль 2. Сущность работы управляющих"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Чтобы понять специфику работы управляющего и ее отличие от деятельности " +
		"любого другого интеллектуального работника, давайте рассмотрим пример работы врача. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "“Во время осмотра пациента доктор, как правило, концентрирует на нем все свое " +
		"внимание. В это время доктор как бы остается один на один со стоящей перед ним задачей и ему " +
		"практически ничего не мешает. Задача, которую должен решить доктор, вполне очевидна. Подходы к " +
		"ее решению зависят от истории болезни пациента. Иными словами, доктор имеет перед собой заданную " +
		"цель: восстановить здоровье больного или по крайней мере улучшить его самочувствие. В целом врачам не " +
		"свойственны ни самоорганизация, ни организация своего труда. Тем не менее для них обычно не составляет " +
		"труда быть эффективными”. \n\n" +
		"Мотивация работника, занимающегося интеллектуальной деятельностью, зависит от его эффективности, от его " +
		"способностей достигать поставленных целей. Если его труд лишен эффективности, то очень скоро его желание " +
		"работать и приносить конкретную пользу исчезает Управляющий какой-либо организации находится совершенно в " +
		"иной позиции. Он постоянно сталкивается с четырьмя главными проблемами, на решение которых он не может повлиять. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "\x31\xE2\x83\xA3 Управляющий не принадлежит самому себе. Каждый волен занимать его время. Он не может уподобиться доктору " +
		"и приказать медсестре не пускать в кабинет пациентов в течение ближайшего часа. \n" +
		"\x32\xE2\x83\xA3 Управляющий вынужден постоянно находиться во «включенном» состоянии, до тех пор, пока не изменит реальность, " +
		"в которой живет и работает. Еще раз воспользуемся аналогией с врачом. Врачи зависят от хода событий. Спрашивая на " +
		"пациента «На что жалуетесь?», они задают рамки для будущей задачи и слышат важную для ее решения информацию. " +
		"Управляющие же должны не плыть по течению, а самостоятельно сформулировать критерии важности для своей работы \n" +
		"\x33\xE2\x83\xA3 Эффективность управляющего зависит от людей, работающих в смежных областях, и от его непосредственного начальства. " +
		"Иными словами, степень эффективности усилий управляющего будет определяться тем, в какой степени плодами его работы " +
		"смогут воспользоваться его коллеги и начальники. \n" +
		"\x34\xE2\x83\xA3 Управляющий воспринимает происходящее, руководствуясь интересами компании которые не всегда могут быть объективными. " +
		"Порой они даже входят в конфликт с окружением компании \n" +
		"Развитие любой организации можно сравнить с развитием биологических организмов. У обычной амебы каждая частичка " +
		"организма находится в постоянном и прямом контакте с окружающей средой, Именно поэтому она не нуждается в специальных " +
		"органах познания окружающего мира или сохранения своей целостности. Но такому сложному и крупному живому организму, " +
		"как человек, необходим скелет, обеспечивающий целостность его тела. Человек нуждается во многих специальных органах, " +
		"приводящих в движение процессы, без которых его функционирование невозможно, – дыхание, пищеварение и т. д. Самое " +
		"главное – человеку необходимы мозг и ряд сложных нервных систем. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "“Организация отличается от человека тем, что она вносит свой вклад в развитие окружающей среды. " +
		"Но по мере роста организации и ее видимых достижений все больше внимания, энергии и способностей управляющего " +
		"направляется на внутренние события в ущерб выполнению стоящих перед ним задач и достижению реальной эффективности " +
		"для внешнего мира”. \n\n" +
		"Автомобилю «Эдсел» компании «Форд» по данным проведенных ранее исследований было суждено " +
		"стать победителем в своей категории. Однако ни одно статистическое исследование не могло учесть " +
		"качественногофактора – к моменту выхода автомобиля на рынок потребители начали руководствоваться не ценой, а " +
		"вкусовыми предпочтениями. Результаты оказались катастрофическими. \n\n" +
		"Эта проблема усугубляется развитием компьютерной " +
		"техники. Компьютеры могут производить массу вычислений, однако адекватных данных во внешних событиях обычно бывает " +
		"слишком мало или они поступают для анализа слишком поздно. Многие события во внешнем мире нельзя назвать полноценными " +
		"«фактами», которые могут подвергаться машинной обработке. \n\n" +
		"Управляющий не может по своему усмотрению изменить эти обстоятельства – они являются необходимой основой его " +
		"существования. Вместе с тем, он должен осознать, что не достигнет никаких результатов, пока сознательно не " +
		"научится быть эффективным. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Домашнее задание: "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "1-й вопрос: “Какие выводы сделали для себя из этого модуля?”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule2(chatID int64, message string, server Server) {
	ans1, ans2 := server.Db.GetTwoAnswers(chatID, 2)

	fmt.Println(ans1, ans2)

	if ans1 == "empty" && ans2 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans2_1", message)

		text := "2-ой вопрос: “В чем отличие руководителя от врача?”"
		msg := tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

	} else if ans1 != "empty" && ans2 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans2_2", message)

		msg := tgbotapi.NewMessage(chatID, "Вы выполнили все задания")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)
	} else if ans1 != "empty" && ans2 != "empty" {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)

		text := "2-ой вопрос: “В чем отличие руководителя от врача?”"
		msg = tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

		msg = tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans2)
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule3(chatID int64, server Server) {
	text := "Модуль 3. Можно ли научиться эффективности?"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Эффективных управляющих объединяет их умение добиваться положительного результата во " +
		"всем, за что бы они ни брались. Манера решения ими задач практически не зависит от места их " +
		"работы или характера деятельности. Следовательно, эффективность – это набор практических методов, " +
		"которым всегда можно научиться. Но для того, чтобы они вошли в привычку, необходима практика, практика и практика. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Основные элементы повышения эффективности труда управляющего: \n" +
		"\x31\xE2\x83\xA3 Эффективные управляющие должны знать, на что они расходуют свое время \n" +
		"\x32\xE2\x83\xA3 Эффективные управляющие должны концентрироваться на достижениях, " +
		"выходящих за рамки их организаций. Главное – это не процесс и методы " +
		"работы, а понимание того, каких конечных результатов необходимо достичь. \n" +
		"\x33\xE2\x83\xA3 Эффективные управляющие должны строить свою деятельность на " +
		"сильных сторонах, как своих, так и руководителей, коллег и подчиненных. " +
		"Необходимо опираться на то, что надежно. Нельзя начинать работу с решения нереальных на данных момент задач. \n" +
		"\x34\xE2\x83\xA3 Эффективные управляющие концентрируются на нескольких важнейших " +
		"участках, где исполнение поставленных заданий принесет наиболее " +
		"ощутимые результаты. По сути дела, они должны заниматься именно " +
		"приоритетными вещами и не тратить сил и времени впустую. \n" +
		"\x35\xE2\x83\xA3 Эффективные управляющие должны системно принимать эффективные " +
		"решения. Решений должно быть немного, но все они должны быть " +
		"фундаментальными и основанными на стратегии, а не сиюминутных тактических соображениях. \n\n" +
		"Эффективность – это что-то вроде привычки, набора практических методов, которым всегда можно " +
		"научиться. Эти методы обманчиво просты, но применять их на практике невероятно сложно. Их надо заучивать как " +
		"таблицу умножения, зазубривать до отвращения, до выработки условного рефлекса, до тех пор, пока они не станут " +
		"частью вашего «я»."
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Домашнее задание: "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "1-й вопрос: “На что вы расходуете свое время? Насколько это эффективно для развития компании?”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule3(chatID int64, message string, server Server) {
	ans1, ans2 := server.Db.GetTwoAnswers(chatID, 3)

	fmt.Println(ans1, ans2)

	if ans1 == "empty" && ans2 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans3_1", message)

		text := "2-ой вопрос: “Каких результатов хочет добиться ваша компания? Насколько четко вы знаете путь к их достижению?”"
		msg := tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

	} else if ans1 != "empty" && ans2 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans3_2", message)

		msg := tgbotapi.NewMessage(chatID, "Вы выполнили все задания")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)
	} else if ans1 != "empty" && ans2 != "empty" {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)

		text := "2-ой вопрос: “Каких результатов хочет добиться ваша компания? Насколько четко вы знаете путь к их достижению?”"
		msg = tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

		msg = tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans2)
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule4(chatID int64, server Server) {
	text := "Модуль 4. Знайте свое время"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "И, наконец, он должен заняться укрупнением времени – выполнить творческое задание урывками невозможно, " +
		"для них нужны большие блоки времени. \n" +
		"Три метода оптимизации времени – регистрация, управление и укрупнение. Что это за методы и как их использовать " +
		"вы узнаете в следующих модулях. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	msg = tgbotapi.NewMessage(chatID, "В этом модуле нет заданий")
	msg.ReplyMarkup = NextModule
	_, _ = server.Bot.Send(msg)
}

func GetModule4(chatID int64, server Server) {
	text := "Модуль 4. Знайте свое время"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "И, наконец, он должен заняться укрупнением времени – выполнить творческое задание урывками невозможно, " +
		"для них нужны большие блоки времени. \n" +
		"Три метода оптимизации времени – регистрация, управление и укрупнение. Что это за методы и как их использовать " +
		"вы узнаете в следующих модулях. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)
}

func StartModule5(chatID int64, server Server) {
	text := "Модуль 5. Что означает регистрация времени?"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Любой управляющий независимо от занимаемой должности время от времени вынужден тратить большое " +
		"количество времени впустую. Жизнь управляющего сопровождают «поглотители времени». При этом время " +
		"представляет собой самый дефицитный ресурс. Поэтому, если мы полагаемся при планировании времени лишь " +
		"на собственную память, то не замечаем, на что расходуется наше время. \n\n" +
		"“Президент одной компании был убежден, что делил свое время приблизительно на три части. Фактические же " +
		"записи расхода времени за шесть недель показали, что на указанные им три блока деятельности он практически " +
		"не затратил времени (хотя услужливая память внушила ему, что он это сделал). Согласно записям по факту, " +
		"большую часть времени он потратил на выполнение диспетчерских функций над операциями, которые шли своим " +
		"чередом, и его вмешательство никоим образом не ускоряло этот процесс. Ему потребовалось несколько раз детально " +
		"фиксировать все, что он делал, прежде чем он убедился, что при анализе использования времени он должен полагаться " +
		"не на память, а на конкретные записи”. \n\n" +
		"Кроме того, управленческий труд связан с обязательным расходованием " +
		"времени на обязательные, но непродуктивные занятия. Поэтому, для того чтобы избежать погружения в рутину, " +
		"управляющий (и все остальные работники интеллектуального труда) должен начать распоряжаться своим временем " +
		"укрупненными блоками. Дробление времени никогда не приведет к желаемому результату. \n\n" +
		"“На написание отчета может потребоваться шесть – восемь часов. Вряд ли будет возможным выполнить эту работу " +
		"за семь часов, уделяя ей два раза по пятнадцать минут в течение трех недель. Чаще всего в результате такой " +
		"работы появляются лишь разрозненные листки, на которых записаны какието неясные мысли. Если же запереться в " +
		"кабинете, отключить телефон и целиком уйти в работу, то вполне возможно за пять или шесть часов написать то, " +
		"что я называю «нулевым вариантом», то есть непосредственным предшественником первого. С этого момента можно " +
		"выполнять работу небольшими порциями, переписывать, исправлять и редактировать раздел за разделом, абзац за " +
		"абзацем и предложение за предложением”. \n\n" +
		"Запись фактического расходования времени можно считать первым шагом на пути повышения эффективности " +
		"управленческого труда. \n\n" +
		"Руководители современных организаций должны выделять значительную часть времени на встречу с работниками и " +
		"обсуждение всех проблем. Умелый управляющий, ставящий перед собой задачу принять оптимальное решение по " +
		"тому или иному кадровому вопросу, обязательно отводит на него несколько часов. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Домашнее задание: "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "1-й вопрос: “Как вы теперь собираетесь планировать свое время?”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule5(chatID int64, message string, server Server) {
	ans1, ans2 := server.Db.GetTwoAnswers(chatID, 5)

	fmt.Println(ans1, ans2)

	if ans1 == "empty" && ans2 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans5_1", message)

		text := "2-ой вопрос: “Какие полезные идеи вы взяли для себя из этого модуля?”"
		msg := tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

	} else if ans1 != "empty" && ans2 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans5_2", message)

		msg := tgbotapi.NewMessage(chatID, "Вы выполнили все задания")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)
	} else if ans1 != "empty" && ans2 != "empty" {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)

		text := "2-ой вопрос: “Какие полезные идеи вы взяли для себя из этого модуля?”"
		msg = tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

		msg = tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans2)
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule6(chatID int64, server Server) {
	text := "Модуль 6. Что означает управление временем?"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Полезным будет выявить участки разбазаривания времени с целью их устранения. Для этого нужно ответить " +
		"на ряд диагностических вопросов. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "\x31\xE2\x83\xA3 Какие виды деятельности не приносят результата, но отнимают время?"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule6(chatID int64, message string, server Server) {
	ans1, ans2, ans3 := server.Db.GetTreeAnswers(chatID, 6)

	fmt.Println(ans1, ans2)

	if ans1 == "empty" && ans2 == "empty" && ans3 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans6_1", message)

		text := "Что можно с этим сделать?"
		msg := tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

	} else if ans1 != "empty" && ans2 == "empty" && ans3 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans6_2", message)

		text := "Что произойдет, если этого не делать вообще?"
		msg := tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

	} else if ans1 != "empty" && ans2 != "empty" && ans3 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans6_3", message)

		text := "\x32\xE2\x83\xA3 Какой из видов моей деятельности может выполнить кто-нибудь другой с " +
			"не меньшим (а, может, и большим) успехом? Здесь нужно понимать, что речь идет не о простом «делегировании» " +
			"своей ответственности другим, а о высвобождении своего времени для концентрации сил и внимания на самом " +
			"важном, и одновременном предоставлении сотрудникам возможности проявить свои способности. \n" +
			"\x33\xE2\x83\xA3 Мешает ли моя манера работы распределению времени коллег и подчиненных? Для этого достаточно " +
			"время от времени задавать им прямой вопрос «Делаю ли я что-то, что поглощает ваше время и не повышает " +
			"эффективность вашей работы?» "
		msg := tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

		text = "“Многие управляющие прекрасно видят все участки своей деятельности, которые практически вхолостую " +
			"потребляют их время, и, тем не менее, боятся отказаться от них, ведь они опасаются по ошибке исключить " +
			"что-то важное из своей работы. Дело, однако, в том, что подобную ошибку можно довольно быстро исправить. " +
			"Поспешное сокращение каких-то важных видов деятельности быстро дает о себе знать, что побуждает управляющих " +
			"принять корректирующие действия”. "
		msg = tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

		msg = tgbotapi.NewMessage(chatID, "В этом модуле нет заданий")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)
	} else if ans1 != "empty" && ans2 != "empty" && ans3 != "empty" {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)

		text := "Что можно с этим сделать?"
		msg = tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

		msg = tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans2)
		_, _ = server.Bot.Send(msg)

		text = "Что произойдет, если этого не делать вообще?"
		msg = tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

		msg = tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans3)
		_, _ = server.Bot.Send(msg)

		text = "\x32\xE2\x83\xA3 Какой из видов моей деятельности может выполнить кто-нибудь другой с " +
			"не меньшим (а, может, и большим) успехом? Здесь нужно понимать, что речь идет не о простом «делегировании» " +
			"своей ответственности другим, а о высвобождении своего времени для концентрации сил и внимания на самом " +
			"важном, и одновременном предоставлении сотрудникам возможности проявить свои способности. \n" +
			"\x33\xE2\x83\xA3 Мешает ли моя манера работы распределению времени коллег и подчиненных? Для этого достаточно " +
			"время от времени задавать им прямой вопрос «Делаю ли я что-то, что поглощает ваше время и не повышает " +
			"эффективность вашей работы?» "
		msg = tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

		text = "“Многие управляющие прекрасно видят все участки своей деятельности, которые практически вхолостую " +
			"потребляют их время, и, тем не менее, боятся отказаться от них, ведь они опасаются по ошибке исключить " +
			"что-то важное из своей работы. Дело, однако, в том, что подобную ошибку можно довольно быстро исправить. " +
			"Поспешное сокращение каких-то важных видов деятельности быстро дает о себе знать, что побуждает управляющих " +
			"принять корректирующие действия”. "
		msg = tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule7(chatID int64, server Server) {
	text := "Модуль 7. Что означает укрупнение времени? "
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Управляющие, которые ведут запись своего рабочего времени и анализируют его, вполне могут определить, " +
		"располагают ли они временем для выполнения важнейших заданий. Конечная стадия управления временем состоит " +
		"в его укрупнении. Суть укрупнения времени заключается в том, что за счёт эффективного тайм менеджмента, " +
		"вы можете потратить больше времени на результативную работу. А это становится возможным, когда у вас уже " +
		"есть понимание ваших временных затрат. \n\n" +
		"Время – это наиболее дефицитный ресурс, и если им не управлять, то все остальное также останется неуправляемым. \n\n" +
		"“Одним из самых умелых распорядителей своим временем, каких мне довелось встречать, был президент крупного банка, " +
		"консультировавшийся у меня по вопросам управления. В течение двух лет я регулярно встречался с ним один раз в " +
		"месяц, и встречи продолжались не более полутора часов. На каждой из них мы обсуждали только один вопрос. Как " +
		"только проходили положенные полтора часа, он незаметно подходил к дверям, вежливо намекая, что его лимит " +
		"времени исчерпан. В ответ на мой вопрос, «Почему вы постоянно стремитесь уместить наши занятия именно в " +
		"полуторачасовые рамки?», он ответил: «Я пришел к выводу, что могу концентрировать свое внимание на какой-то " +
		"одной теме в течение примерно полутора часов. Если же я уделяю выполнению задания больше времени, мое внимание " +
		"начинает рассеиваться. Но я также обнаружил, что более или менее важное задание можно выполнить не менее чем за " +
		"полтора часа». Итак, каждый месяц в определенный день я находился у него в кабинете в течение полутора часов, и " +
		"наши беседы ни разу не были прерваны телефонными звонками или визитами секретарши, обычно спешащей объявить о " +
		"прибытии важного человека \n" +
		"Излишне говорить, что этот руководитель буквально за считанные часы хорошо спланированных встреч достиг " +
		"гораздо больше, чем многие другие не менее способные работники за ежедневные встречи со мной”. \n\n" +
		"Чем более высокую ступень на иерархической лестнице занимает управляющий, тем выше доля времени, которое ему " +
		"неподконтрольно и которое он не может использовать на благо своей организации. Чем крупнее организация, тем " +
		"больше требуется времени на управление и сохранение ее целостности и тем меньше его остается для того, чтобы " +
		"эта организация функционировала и что-то производила. \n\n" +
		"Именно поэтому хороший руководитель всегда стремится к укрупнению своего времени. Он понимает, что для " +
		"решения творческих задач ему необходимы крупные блоки времени, потому что мелкие отрезки времени невозможно " +
		"использовать эффективно. Установлено, что даже четверть рабочего дня, сгруппированная крупными временными " +
		"блоками, может быть достаточна для выполнения важных работ. Наоборот, три четверти рабочего времени могут быть " +
		"затрачены бездарно, если они разбиты по пятнадцать – тридцать минут. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Домашнее задание: "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "“На что вы бы теперь хотели укрупнить в своем графике?”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule7(chatID int64, message string, server Server) {
	ans1 := server.Db.GetAnswer(chatID, 7)

	fmt.Println(ans1)

	if ans1 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans7_1", message)

		msg := tgbotapi.NewMessage(chatID, "Вы выполнили все задания")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule8(chatID int64, server Server) {
	text := "Модуль 8. Собственные обязательства управляющего"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Задать себе вопрос «Чем я могу помочь своей организации?» – значит начать " +
		"поиск неиспользованных резервов на своем рабочем месте. Часто оказывается, что то, что мы привыкли " +
		"называть образцовым выполнением\nсвоих обязанностей – лишь малая часть того, чего можно было бы достичь на " +
		"рабочем месте. \n\n" +
		"Управляющие, не задающиеся вопросом о целях компании и потенциале своего вклада, часто сужают диапазон " +
		"своих достижений. Ориентация на вклад – ориентация на эффективность"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Домашнее задание: "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "“А как вы помогаете своей организации?”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule8(chatID int64, message string, server Server) {
	ans1 := server.Db.GetAnswer(chatID, 8)

	fmt.Println(ans1)

	if ans1 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans8_1", message)

		msg := tgbotapi.NewMessage(chatID, "Вы выполнили все задания")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule9(chatID int64, server Server) {
	text := "Модуль 9. Повышение эффективности работы специалиста"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Нацеленность на свой вклад, на достижение качественных результатов в работе особенно важна для работника " +
		"интеллектуального труда. Выбор цели влияет на его способность добиться результата. \n\n" +
		"Необходимо создавать условия, в которых специалист способен повысить свою собственную эффективность и " +
		"эффективность специальности \n\n" +
		"Общаясь с сотрудниками, эффективный управляющий мысленно спрашивает их: «Какого вклада вы от меня ждете, " +
		"чтобы внести в нашу организацию свой собственный вклад? Как, когда и в какой форме вам это нужно?»"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Домашнее задание: "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "“Считаете ли эффективными ваших сотрудников? Что вы бы хотели улучшить?”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule9(chatID int64, message string, server Server) {
	ans1 := server.Db.GetAnswer(chatID, 9)

	fmt.Println(ans1)

	if ans1 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans9_1", message)

		msg := tgbotapi.NewMessage(chatID, "Вы выполнили все задания")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule10(chatID int64, server Server) {
	text := "Модуль 10. Правильные взаимоотношения"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Сконцентрированность на вкладе предполагает четыре главных условия эффективных взаимоотношений: \n" +
		"\xF0\x9F\x94\xB6 Коммуникация \n" +
		"\xF0\x9F\x94\xB6 Коллективная деятельность \n" +
		"\xF0\x9F\x94\xB6 Саморазвитие \n" +
		"\xF0\x9F\x94\xB6 Развитие других"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "“Коммуникация является важнейшим параметром работы любой крупной современной организации. \n" +
		"Сконцентрированность на вкладе улучшает коммуникацию (то есть, взаимопонимание) и делает возможной эффективную " +
		"коллективную деятельность. \n" +
		"Индивидуальное саморазвитие в значительной мере зависит от сконцентрированности на вкладе в общее дело. \n" +
		"Управляющий, сконцентрированный на личном вкладе, стимулирует развитие других – подчиненных, коллег, начальников”."
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Домашнее задание: "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "“С чего вы хотели бы начать свое саморазвитие?”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule10(chatID int64, message string, server Server) {
	ans1 := server.Db.GetAnswer(chatID, 10)

	fmt.Println(ans1)

	if ans1 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans10_1", message)

		msg := tgbotapi.NewMessage(chatID, "Вы выполнили все задания")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)

	} else {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule11(chatID int64, server Server) {
	text := "Модуль 11. Ставка на сильные качества"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Проблема выявления сильных сторон и положительных качеств изначально встает перед управляющим при подборе " +
		"кадров. Эффективный руководитель заполняет вакансии и продвигает работников по службе, исходя из их качеств. " +
		"Принимая кадровые решения, он делает упор на наличие достоинств, а не отсутствие недостатков у работников. \n\n" +
		"Управляющий, которого заботят слабые стороны работника, а не то, в чем он силен и на что способен – сам по сути " +
		"человек слабый. Правильнее будет действовать по принципу эффективного использования сильных сторон кандидатов. \n\n" +
		"Желание застраховаться от тех или иных слабостей мешает организации достичь целей своей деятельности. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Каким образом эффективные управляющие могут подбирать работников, " +
		"исходя из их сильных качеств, не приспосабливая рабочие места под особенности личности? \n" +
		"\x31\xE2\x83\xA3 Рабочие места – дело рук человека со всеми присущими ему недостатками. Эффективный управляющий " +
		"всегда будет остерегаться «невозможной» работы, с которой не может справиться нормальный человек. \n" +
		"\x32\xE2\x83\xA3 Необходимо придать ответственность каждой должности и быть требовательным по отношению к работнику. " +
		"Работа должна выявлять все способности человека и добиваться того, чтобы эти способности сказывались на " +
		"результатах. \n" +
		"\x33\xE2\x83\xA3 Эффективный управляющий знает, что начинать работу с людьми нужно с раскрытия и правильного использования " +
		"их потенциала, а не с раздачи поручений выполнять стандартные обязанности. При найме человека эффективный " +
		"управляющий должен анализировать возможности каждого кандидата. \n" +
		"\x34\xE2\x83\xA3 Эффективный управляющий знает, что для продуктивного использования сильных качеств работника часто бывает " +
		"необходимым мириться с его слабостями. \n\n" +
		"Каждый начальник отвечает за работу своих подчиненных. Он обладает полномочиями влиять на карьеру других. " +
		"Поэтому продуктивное использование сильных сторон человека становится чем-то большим, нежели основа эффективности. " +
		"Это – обязанность руководителя. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Домашнее задание: "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "1-й вопрос: “А какие у вас есть сильные стороны?”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule11(chatID int64, message string, server Server) {
	ans1, ans2 := server.Db.GetTwoAnswers(chatID, 11)

	fmt.Println(ans1, ans2)

	if ans1 == "empty" && ans2 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans11_1", message)

		text := "2-ой вопрос: “Как часто вы на них опираетесь?”"
		msg := tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

	} else if ans1 != "empty" && ans2 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans11_2", message)

		msg := tgbotapi.NewMessage(chatID, "Вы выполнили все задания")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)
	} else if ans1 != "empty" && ans2 != "empty" {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)

		text := "2-ой вопрос: “Как часто вы на них опираетесь?”"
		msg = tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

		msg = tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans2)
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule12(chatID int64, server Server) {
	text := "Модуль 12. Повышение собственной эффективности"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Эффективных управляющих беспокоят их ограничения, но они способны видеть и то, что им под силу выполнить. " +
		"В то время, как другие жалуются, что у них не получается то одно, то другое, эффективный управляющий предпочитает " +
		"не тратить время и делать то, что он умеет лучше всего. \n" +
		"Изменение человеческой натуры подчиненных не входит в задачи управляющего. Его задача заключается в том, " +
		"чтобы оценить все имеющиеся в коллективе умения и силы, а затем максимально эффективно использовать их в " +
		"интересах организации. \n" +
		"Эффективный управляющий начинает день с ответа на вопрос: «Что я могу сделать?» "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Домашнее задание: "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "“А что вы можете сделать сегодня?”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule12(chatID int64, message string, server Server) {
	ans1 := server.Db.GetAnswer(chatID, 12)

	fmt.Println(ans1)

	if ans1 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans12_1", message)

		msg := tgbotapi.NewMessage(chatID, "Вы выполнили все задания")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)

	} else {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule13(chatID int64, server Server) {
	text := "Модуль 13. Всему свое время"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Секрет эффективности заключается в сосредоточенности и целенаправленности. \n\n" +
		"“Чем сильнее управляющий концентрируется на том, что обеспечивает успех его организации, тем больше он " +
		"нуждается в укрупненных отрезках времени на выполнение своих функций. Чем больше он ориентируется не на процесс " +
		"работы как таковой, а на достижение результатов, тем сильнее необходимость в приложении усилий, которые, в " +
		"свою очередь, требуют весомого количества времени. В противном случае эти усилия не будут плодотворными. Вместе " +
		"с тем, для того, чтобы иметь в своем распоряжении по-настоящему продуктивное время (несколько часов или недель)," +
		" необходимы дисциплина и железная решимость сказать «нет». \n\n" +
		"Секрет людей, которым удается успешно решать множество сложных задач, кроется в том, что в определенный период " +
		"времени они концентрируются на чем-то одном. Делая все последовательно, они тратят гораздо меньше времени, " +
		"чем большинство из нас. \n\n" +
		"Необходимо регулярно пересматривать использованные в прошлом методы и отказываться " +
		"от них, если они перестали быть эффективными. В силах каждого управляющего ограничить свою зависимость от " +
		"прошлого путем сокращения «перешедших по наследству» бесперспективных видов деятельности и задач. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Домашнее задание: "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "1-й вопрос: “Как часто вы концентрируетесь на чем-то одном?”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule13(chatID int64, message string, server Server) {
	ans1, ans2 := server.Db.GetTwoAnswers(chatID, 13)

	fmt.Println(ans1, ans2)

	if ans1 == "empty" && ans2 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans13_1", message)

		text := "2-ой вопрос: “Вы больше уделяете внимание процессу или результату?”"
		msg := tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

	} else if ans1 != "empty" && ans2 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans13_2", message)

		msg := tgbotapi.NewMessage(chatID, "Вы выполнили все задания")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)
	} else if ans1 != "empty" && ans2 != "empty" {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)

		text := "2-ой вопрос: “Как часто вы на них опираетесь?”"
		msg = tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

		msg = tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans2)
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule14(chatID int64, server Server) {
	text := "Модуль 14. Приоритеты"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Правила определения приоритетов: \n" +
		"\xF0\x9F\x94\xB6 Ориентируйтесь на будущее, а не на прошлое. \n" +
		"\xF0\x9F\x94\xB6 Концентрируйтесь на возможностях, а не на проблемах. \n" +
		"\xF0\x9F\x94\xB6 Выбирайте свое собственное направление, не плывите по течению вместе с другими. \n\n" +
		"Систематическое избавление от старого является единственным средством внедрения нового. \n\n" +
		"Приоритеты так же, как и задачи, необходимо время от времени пересматривать, чтобы в один прекрасный день " +
		"не обнаружить, что вы движетесь не тем курсом. \n\n" +
		"Эффективный руководитель не позволит себе отвлечься от решения той задачи, на которой он сконцентрирован " +
		"в данный момент. \n\n" +
		"Решив одну задачу, он оценивает ситуацию, и принимается за решение следующей, которая теперь " +
		"приобретает первостепенное значение. \n\n" +
		"Концентрация – единственная возможность управляющего стать подлинным хозяином времени и " +
		"обстоятельств, а не их рабом. \n\n" +
		"Ставьте себе высокие цели, которые позволяют круто изменить ситуацию, вместо надежных и легко достижимых"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Домашнее задание: "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "“Какие у вас цели на будущее? Постарайтесь описать их как можно детальнее”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule14(chatID int64, message string, server Server) {
	ans1 := server.Db.GetAnswer(chatID, 14)

	fmt.Println(ans1)

	if ans1 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans14_1", message)

		msg := tgbotapi.NewMessage(chatID, "Вы выполнили все задания")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)

	} else {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule15(chatID int64, server Server) {
	text := "Модуль 15. Элементы принятия решении"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Принятие решений – процесс, занимающий незначительную часть времени управляющего. Однако это типично " +
		"управленческая задача и то, как управляющий справляется с ней, сильно влияет на его эффективность \n\n" +
		"Эффективный управляющий не стремится принимать множество решений. Он концентрируется только на самых важных, " +
		"и не ставит себе цель решить «текущие задачи». \n\n" +
		"Самое сложное в принятии решений – выбор верного компромисса. Необходимо научиться отличать нужный " +
		"компромисс от ненужного. \n\n" +
		"Элементы процесса принятия решений: \n" +
		"\x31\xE2\x83\xA3 Четкое осознание того, что данная проблема носит общий характер и может быть решена только посредством " +
		"установления правила или принципа. \n" +
		"\x32\xE2\x83\xA3 Определение требований, которым должно отвечать решение проблемы, то есть определение ограничений. \n" +
		"\x33\xE2\x83\xA3 Продумывание того, что считать «правильным», и понимание того, на какие компромиссы, уступки или " +
		"корректировки можно пойти. Важно отметить, что мы говорим о «верном», а не о «приемлемом»! \n" +
		"\x34\xE2\x83\xA3 Пуск решения в действие. \n" +
		"\x35\xE2\x83\xA3 Обратная связь, которая проверяет актуальность и эффективность решения в реальных условиях. \n\n" +
		"Как отличить правильный компромисс от неправильного? Имеются два вида компромиссов. Первый выражается старой " +
		"пословицей: «Половина буханки хлеба лучше, чем совсем ничего». Второй вид иллюстрируется Соломоновым решением в " +
		"притче о младенце – (см. врез), когда лучше не иметь ничего, чем половину. В первом случае пограничные условия " +
		"все еще удовлетворяются: хлеб – это еда, и половина буханки хлеба вполне удовлетворяет условиям. В то же " +
		"время половина младенца не может удовлетворить пограничные условия. Ведь половина ребенка не может существовать, " +
		"это будет просто половина трупа. \n\n" +
		"“Хорошим примером могли бы служить действия президента Кеннеди. После прихода к власти на Кубе Фиделя Кастро " +
		"администрация США поддержала действия его Фиделя Кастро и снабдила их оружием (это выступление закончилось " +
		"военным поражением мятежников). Кеннеди извлек правильные уроки. Во время Карибского кризиса он четко определил " +
		"условия для правильного компромисса (с одной стороны, он отказался от своего изначального требования о проведении " +
		"на Кубе наземной инспекции), а с другой не отступил от важнейшего условия (демонтирования и вывоза советских ракет). \n\n" +
		"Перед тем, как принять ответственное решение, опытный управляющий исходит из того, что проблема носит типичный " +
		"характер. Он также признает, что событие, привлекшее его внимание, на самом деле является симптомом. Он всегда " +
		"пытается выявить сущность проблемы и не останавливается на лечении одного лишь симптома. \n\n" +
		"Если событие по-настоящему уникально, опытный управляющий всегда будет пытаться разглядеть в нем предвестника " +
		"новой базовой проблемы, потому что уникальное на первый взгляд событие может оказаться первым проявлением " +
		"типичной ситуации”. \n\n" +
		"К царю Соломону пришло две женщины, каждая из которых называла себя матерью одного и того же ребенка. " +
		"Царь приказал рассечь ребенка пополам и отдать по его половине каждой из женщин. Одна женщина согласилась " +
		"с этим решением, а другая – настоящая его мать – попросила царя, чтобы тот не убивал младенца, а отдал его " +
		"другой участнице спора живым. \n\n" +
		"После того, как решение принято, не стоит пренебрегать обратной связью. Для ее обеспечения необходима " +
		"систематизированная и организованная информация. Если мы не научим себя контролировать и проверять результаты, " +
		"рано или поздно это придет к снижению или полной потере эффективности. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	msg = tgbotapi.NewMessage(chatID, "В этом модуле нет заданий")
	msg.ReplyMarkup = NextModule
	_, _ = server.Bot.Send(msg)
}

func GetModule15(chatID int64, server Server) {
	text := "Модуль 15. Элементы принятия решении"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Принятие решений – процесс, занимающий незначительную часть времени управляющего. Однако это типично " +
		"управленческая задача и то, как управляющий справляется с ней, сильно влияет на его эффективность \n\n" +
		"Эффективный управляющий не стремится принимать множество решений. Он концентрируется только на самых важных, " +
		"и не ставит себе цель решить «текущие задачи». \n\n" +
		"Самое сложное в принятии решений – выбор верного компромисса. Необходимо научиться отличать нужный " +
		"компромисс от ненужного. \n\n" +
		"Элементы процесса принятия решений: \n" +
		"\x31\xE2\x83\xA3 Четкое осознание того, что данная проблема носит общий характер и может быть решена только посредством " +
		"установления правила или принципа. \n" +
		"\x32\xE2\x83\xA3 Определение требований, которым должно отвечать решение проблемы, то есть определение ограничений. \n" +
		"\x33\xE2\x83\xA3 Продумывание того, что считать «правильным», и понимание того, на какие компромиссы, уступки или " +
		"корректировки можно пойти. Важно отметить, что мы говорим о «верном», а не о «приемлемом»! \n" +
		"\x34\xE2\x83\xA3 Пуск решения в действие. \n" +
		"\x35\xE2\x83\xA3 Обратная связь, которая проверяет актуальность и эффективность решения в реальных условиях. \n\n" +
		"Как отличить правильный компромисс от неправильного? Имеются два вида компромиссов. Первый выражается старой " +
		"пословицей: «Половина буханки хлеба лучше, чем совсем ничего». Второй вид иллюстрируется Соломоновым решением в " +
		"притче о младенце – (см. врез), когда лучше не иметь ничего, чем половину. В первом случае пограничные условия " +
		"все еще удовлетворяются: хлеб – это еда, и половина буханки хлеба вполне удовлетворяет условиям. В то же " +
		"время половина младенца не может удовлетворить пограничные условия. Ведь половина ребенка не может существовать, " +
		"это будет просто половина трупа. \n\n" +
		"“Хорошим примером могли бы служить действия президента Кеннеди. После прихода к власти на Кубе Фиделя Кастро " +
		"администрация США поддержала действия его Фиделя Кастро и снабдила их оружием (это выступление закончилось " +
		"военным поражением мятежников). Кеннеди извлек правильные уроки. Во время Карибского кризиса он четко определил " +
		"условия для правильного компромисса (с одной стороны, он отказался от своего изначального требования о проведении " +
		"на Кубе наземной инспекции), а с другой не отступил от важнейшего условия (демонтирования и вывоза советских ракет). \n\n" +
		"Перед тем, как принять ответственное решение, опытный управляющий исходит из того, что проблема носит типичный " +
		"характер. Он также признает, что событие, привлекшее его внимание, на самом деле является симптомом. Он всегда " +
		"пытается выявить сущность проблемы и не останавливается на лечении одного лишь симптома. \n\n" +
		"Если событие по-настоящему уникально, опытный управляющий всегда будет пытаться разглядеть в нем предвестника " +
		"новой базовой проблемы, потому что уникальное на первый взгляд событие может оказаться первым проявлением " +
		"типичной ситуации”. \n\n" +
		"К царю Соломону пришло две женщины, каждая из которых называла себя матерью одного и того же ребенка. " +
		"Царь приказал рассечь ребенка пополам и отдать по его половине каждой из женщин. Одна женщина согласилась " +
		"с этим решением, а другая – настоящая его мать – попросила царя, чтобы тот не убивал младенца, а отдал его " +
		"другой участнице спора живым. \n\n" +
		"После того, как решение принято, не стоит пренебрегать обратной связью. Для ее обеспечения необходима " +
		"систематизированная и организованная информация. Если мы не научим себя контролировать и проверять результаты, " +
		"рано или поздно это придет к снижению или полной потере эффективности. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func StartModule16(chatID int64, server Server) {
	text := "Модуль 16. Эффективные решения"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Решение представляет собой суждение. Это выбор оптимального варианта из имеющихся альтернатив. " +
		"В основе него лежит выбор между правильным и неправильным. В лучшем случае решение – это действительно " +
		"выбор между верным и неверным действием, но гораздо чаще это выбор между двумя направлениями движения, " +
		"верность которых нельзя доказать. \n\n" +
		"Наиболее действенным методом нахождения приемлемого измерения является личное участие в обратной связи, только " +
		"связь эта должна осуществляться до принятия решения. \n\n" +
		"“Решения, принимаемые управляющими в современных организациях, как правило, не могут основываться на всеобщем " +
		"одобрении. Качественные решения вырабатываются только в результате столкновения взглядов, высказывания разных " +
		"точек зрения и выбора одного суждения из многих возможных. Таким образом, первое правило принятия решения " +
		"должно звучать так: «Если не имеется предварительных разногласий, невозможно выработать оптимальное решение». \n\n" +
		"Наше видение сужается, если нет альтернатив. \n\n" +
		"Если разногласий нет, их можно спровоцировать. Каковы достоинства спровоцированных разногласий? \n" +
		"\x31\xE2\x83\xA3 Только таким образом лицо, ответственное за принятие решений, может избежать участи пленника своей " +
		"организации. Каждый в организации пытается навязать ему свое мнение, и сделать так, чтобы прошло решение, " +
		"нужное именно ему. \n" +
		"\x32\xE2\x83\xA3 Разногласия могут обеспечить альтернативы предлагаемому решению. Решение без альтернативы – " +
		"ход отчаявшегося игрока. \n\n" +
		"Прежде чем принять окончательное решение, опытный руководитель организует обмен мнениями. Такой метод дает ему " +
		"возможность выбора оптимального решения из нескольких возможных, а также замены одобренного решения, которое " +
		"впоследствии оказалось непригодным. \n\n" +
		"“Управляющим платят не за то, чтобы они делали то, что им по душе. Им платят, чтобы они способствовали правильному " +
		"ходу производственного процесса. Больше всего от них ждут выработки эффективных решений”. \n\n" +
		"В век всеобщей компьютеризации стоит задуматься, способна ли машина принимать решения. Компьютер может оказать " +
		"влияние на стратегические решения. Разумеется, он не в состоянии их принять. Все, на что он способен – разработать " +
		"заключения, которые следуют из определенных допущений, или определить, какие допущения лежат в основе " +
		"предполагаемого хода действий. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Домашнее задание: "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "“А как вы теперь будете принимать свои решения после прочтения двух последних модулей?”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule16(chatID int64, message string, server Server) {
	ans1 := server.Db.GetAnswer(chatID, 16)

	fmt.Println(ans1)

	if ans1 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans16_1", message)

		msg := tgbotapi.NewMessage(chatID, "Вы выполнили все задания")
		msg.ReplyMarkup = NextModule
		_, _ = server.Bot.Send(msg)

	} else {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)
	}
}

func StartModule17(chatID int64, server Server) {
	text := "Модуль 17. Заключение"
	msg := tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "Управляющему платят за его эффективность. Он просто обязан быть эффективным в той организации, в которой " +
		"работает. Эффективности можно и нужно учиться. При этом, управляющий должен научиться быть эффективным самому, " +
		"а, помимо этого, сделать эффективной всю свою организацию. \n\n" +
		"Прежде всего, управляющий должен заняться регистрированием своего времени, его анализом и устранением его " +
		"непроизводительного расходования. Разобравшись с этим процедурным вопросом, он должен заняться оценкой своего " +
		"личного вклада в организацию и понять, в какой степени он ориентируется не на отдельные задачи, а на достижение " +
		"общих целей организации. \n\n" +
		"Активировав свои сильные стороны, управляющий может проделать то же самое в отношении своих подчиненных, а затем " +
		"пойти по широкому и четко очерченному пути принятия правильных решений, которые и помогут ему достичь высот " +
		"производительности и эффективности – как своей личной, так и всей компании. "
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

	text = "“Понравился ли вам курс? Что именно вам понравилось?”"
	msg = tgbotapi.NewMessage(chatID, text)
	_, _ = server.Bot.Send(msg)

}

func ExecuteModule17(chatID int64, message string, server Server) {
	ans1, ans2, ans3 := server.Db.GetTreeAnswers(chatID, 17)

	fmt.Println(ans1, ans2)

	if ans1 == "empty" && ans2 == "empty" && ans3 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans17_1", message)

		text := "“Есть ли у вас замечания или предложения?”"
		msg := tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

	} else if ans1 != "empty" && ans2 == "empty" && ans3 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans17_2", message)

		text := "“Как бы вы оценили эффективность курса для вашего саморазвития (от 1 до 5)”"
		msg := tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

	} else if ans1 != "empty" && ans2 != "empty" && ans3 == "empty" {
		server.Db.AnswerQuestion(chatID, "ans17_3", message)
		server.Db.UpdateUserModule(chatID, 18)

		text := "Спасибо за прохождение нашего курса! Рекомендуйте его своим друзьям и приглашайте их на нашу платформу " +
			"Cursus.asia"
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = mainMenu
		_, _ = server.Bot.Send(msg)
	} else if ans1 != "empty" && ans2 != "empty" && ans3 != "empty" {
		msg := tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans1)
		_, _ = server.Bot.Send(msg)

		text := "“Есть ли у вас замечания или предложения?”"
		msg = tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

		msg = tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans2)
		_, _ = server.Bot.Send(msg)

		text = "“Как бы вы оценили эффективность курса для вашего саморазвития (от 1 до 5)”"
		msg = tgbotapi.NewMessage(chatID, text)
		_, _ = server.Bot.Send(msg)

		msg = tgbotapi.NewMessage(chatID, "Ваш ответ: "+ans3)
		_, _ = server.Bot.Send(msg)
	}
}
