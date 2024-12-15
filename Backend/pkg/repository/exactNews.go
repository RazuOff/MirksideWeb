package repository

import (
	"errors"
	"mirkside/pkg/model"
)

var exactNews = []model.ExactNews{
	{
		ID: 1,
		ExactNewsBlock: []model.ExactNewsBlock{
			{Text: "На севере деревни исчез местный крестьянин. Люди говорят, что его видели в лесу в ночное время. На севере деревни исчез местный крестьянин. Люди говорят, что его видели в лесу в ночное время. На севере деревни исчез местный крестьянин. Люди говорят, что его видели в лесу в ночное время. На севере деревни исчез местный крестьянин. Люди говорят, что его видели в лесу в ночное время.", ImgUrl: "/img/Image for news.png"},
			{Text: "Когда они пошли искать его, следы обрываются. Когда они пошли искать его, следы обрываются. Когда они пошли искать его, следы обрываются. Когда они пошли искать его, следы обрываются.Когда они пошли искать его, следы обрываются."},
			{ImgUrl: "/img/Image for news.png"},
		},
	},
	{
		ID: 2,
		ExactNewsBlock: []model.ExactNewsBlock{
			{Text: "Жители деревни утверждают, что старый монастырь на холме скрывает темные силы."},
			{Text: "Несколько людей слышали странные звуки ночью, до того как исчезли без следа."},
			{ImgUrl: "/img/main.png"},
		},
	},
	{
		ID: 3,
		ExactNewsBlock: []model.ExactNewsBlock{
			{Text: "Кто-то заметил тень женщины на кладбище, которая уходит вглубь леса."},
			{Text: "Селяне боятся, что это предвестие чьей-то смерти."},
			{ImgUrl: "path/to/image3.jpg"},
		},
	},
	{
		ID: 4,
		ExactNewsBlock: []model.ExactNewsBlock{
			{Text: "Ночью в лесу неподалеку от деревни видят странные огни."},
			{Text: "Кто-то утверждает, что это души погибших, ищущие покой."},
			{ImgUrl: "path/to/image4.jpg"},
		},
	},
	{
		ID: 5,
		ExactNewsBlock: []model.ExactNewsBlock{
			{Text: "Мужчина был найден в лесу с пустыми глазами и без памяти."},
			{Text: "Он утверждает, что его преследуют темные существа, и он слышит их шепот."},
			{ImgUrl: "path/to/image5.jpg"},
		},
	},
	{
		ID: 6,
		ExactNewsBlock: []model.ExactNewsBlock{
			{Text: "Сельские старожилы предупреждают, чтобы никто не подходил к старому колодцу."},
			{Text: "Говорят, что тот, кто посмотрит в его темные воды, станет свидетелем ужасных видений."},
			{ImgUrl: "path/to/image6.jpg"},
		},
	},
	{
		ID: 7,
		ExactNewsBlock: []model.ExactNewsBlock{
			{Text: "В небе упала звезда, но вместо огня на земле осталась лишь темная, холодная тень."},
			{Text: "Несколько людей утверждают, что она перемещается ночью."},
			{ImgUrl: "path/to/image7.jpg"},
		},
	},
	{
		ID: 8,
		ExactNewsBlock: []model.ExactNewsBlock{
			{Text: "В одном из старых домов было найдено письмо, датированное 100 лет назад."},
			{Text: "Оно предупреждает о гибели всей деревни, если не будет завершен древний обряд."},
			{ImgUrl: "path/to/image8.jpg"},
		},
	},
	{
		ID: 9,
		ExactNewsBlock: []model.ExactNewsBlock{
			{Text: "Селяне услышали странные звуки из старых подземных катакомб под деревней."},
			{Text: "Некоторые утверждают, что это крики тех, кто не смог выбраться из темных лабиринтов."},
			{ImgUrl: "path/to/image9.jpg"},
		},
	},
	{
		ID: 10,
		ExactNewsBlock: []model.ExactNewsBlock{
			{Text: "Женщина из соседней деревни утверждает, что встречала духа леса."},
			{Text: "Он предложил ей сделку, но взамен попросил душу ближайшего жителя."},
			{ImgUrl: "path/to/image10.jpg"},
		},
	},
}

func GetExactNewsPageById(id int) ([]model.ExactNewsBlock, error) {
	if id > len(exactNews) || id <= 0 {
		return []model.ExactNewsBlock{}, errors.New("id do not exists")
	}

	return exactNews[id-1].ExactNewsBlock, nil
}
