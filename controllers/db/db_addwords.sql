use hangman;

insert into wordlist
	(wordlist_name, word_num, word, message, creation_date)
values
	('default', 0, 'doggo', 'Doggos make life worth living',now()),
    ('default', 1, 'avocado', 'House prices are proportional to avocado consumption',now()),
    ('default', 2, 'brooklyn', 'Cooking raw with the Brooklyn boy',now()),
    ('default', 3, 'guatemala', 'We can get away-ay, maybe to Guatemala…',now()),
    ('default', 4, 'football', 'Sissoko…kicks it up towards Llorente…Dele Alli…they\'re slipping…they\'re sliding…IT\'S IN LUCAS MOURA WITH THE HATTRICK GOAL',now()),
    ('default', 5, 'decorum', 'My friend, you would not tell with such high zest, To children ardent for some desperate glory, The old lie: Dulce et decorum est, Pro patria mori',now()),
    ('default', 6, 'hedonistic', 'The only way to get rid of temptation is to yield to it. Resist it, and your soul grows sick with longing for the things it has forbidden to itself',now()),
    ('default', 7, 'england', 'It\'s coming home…it\'s coming home…it\'s coming…FOOTBALL\'S COMING HOME',now()),
    ('default', 8, 'climbing', 'Do not go gentle into that good night, Forearms should burn and rave at close of day; Climb, climb, to ever greater heights',now())
    ;
    
    
select * from wordlist;

#delete from wordlist where id_num > 0;