package database

var seedPosts = append(seedBookPosts, seedTextPosts...)

// posts with a book cover or other picture
var seedBookPosts = []seedPost{
	{
		user:       "saltwater_pages",
		title:      "Autobiography of Red ruined every other novel for me (in a good way)",
		image:      seedCovers + "autobiography-of-red.jpg",
		daysAgo:    58,
		categories: []string{"Book Reviews", "Poetry", "Literary Fiction", "Anne Carson"},
		body: `A novel in verse about Geryon, the red winged monster that Herakles kills in the old myth, except here Geryon is a shy gay teenager who takes photographs and Herakles is the careless boy who breaks his heart.

I didn't think a book could be this strange and this tender at the same time. Carson starts with a dry little essay about the Greek poet Stesichorus, then some fragment "translations", and then suddenly you are inside the most aching coming-of-age story I have read in years.

The volcano chapters near the end. The tango bar in Buenos Aires. "He was a monster and he was red." I read the last forty pages twice in a row.

5/5. If you think you don't like poetry, this is the one to try.`,
		comments: []seedComment{
			{"bluewing_books", "The photograph titles in the last section! Every single one is its own poem."},
			{"paperbag_critic", "I tried. I really tried. I got to the part with the fragments and felt like I was being tested on homework I was never assigned. Maybe I'm not the target audience."},
			{"saltwater_pages", "@paperbag_critic skip the front matter and start with the romance itself, then go back. Honestly that's how half of my friends read it."},
			{"silver_sage", `I taught this book in a seminar the year it came out and it split the room exactly like this thread is splitting. The students who expected a novel were annoyed; the ones who expected poetry were annoyed too. Then about week three something clicked for most of them.

What Carson does is take a myth told from the hero's point of view and turn the camera around. Geryon's interior life is all the original text never gave him. That's the whole trick, and it is a profound one.`},
			{"paperbag_critic", "Fine. I'll give it another go. If I end up crying on a train I'm blaming every one of you."},
			{"dusk_reader", "Just ordered it because of this post. My TBR is crying."},
			{"kirjatoukka_kaisa", "There is a Finnish translation too, Punainen omaelämäkerta. Surprisingly good, I'll post about it at some point."},
		},
	},
	{
		user:       "silver_sage",
		title:      "Eros the Bittersweet: the best book about desire is an essay about Sappho",
		image:      seedCovers + "eros-the-bittersweet.jpg",
		daysAgo:    55,
		categories: []string{"Essays", "Non-fiction", "Poetry", "Anne Carson"},
		body: `Carson's first book, originally her doctoral work, and it still reads like a love letter to the gap between wanting and having.

Her central idea: eros is lack. Sappho called it "glukupikron", sweetbitter, and Carson argues that desire needs an obstacle, a third point in a triangle, to exist at all. The lover, the beloved, and whatever stands between them.

It's academic in places but never dry. Chapter by chapter she walks through Sappho, the Greek novelists, Plato's Phaedrus and the invention of the alphabet, and you come out the other side seeing every love story you have ever read differently.

Recommended reading order for the brave: this first, then Autobiography of Red.`,
		comments: []seedComment{
			{"saltwater_pages", "The triangle idea genuinely changed how I read romance novels. Every obstacle in a love story is load bearing."},
			{"rosie_rides_west", "So what you're telling me is that the miscommunication trope is actually ancient Greek philosophy. I feel vindicated."},
			{"silver_sage", "@rosie_rides_west in a manner of speaking, yes. Sappho would have written an excellent enemies-to-lovers."},
			{"lurker_lena", "Is this readable without knowing any Greek?"},
			{"saltwater_pages", "@lurker_lena completely. She explains everything as she goes."},
			{"theo_logian", "Her chapter on the reach of desire toward something that can't be possessed reads almost like mystical theology. I quoted it in an essay on Augustine and my professor loved it."},
		},
	},
	{
		user:       "kirjatoukka_kaisa",
		title:      "Punainen omaelämäkerta: how does the Finnish translation hold up?",
		image:      seedCovers + "punainen-omaelamakerta.jpg",
		daysAgo:    52,
		categories: []string{"Translated Fiction", "Poetry", "Finnish Literature", "Anne Carson"},
		body: `Promised @saltwater_pages I'd post about this. Autobiography of Red in Finnish!

I read both versions side by side over a week. Short verdict: the translation is brave. Carson's English is full of flat, almost deadpan lines that suddenly turn into something enormous, and Finnish handles the deadpan part beautifully. Our language is good at saying huge things in a very quiet voice.

What gets lost: some of the wordplay, obviously. What's gained: the volcano parts sound even more ancient somehow.

Also look at that cover. Red on grey. It looks like lava under ash.

Has anyone else read a favourite book in two languages? Did it feel like the same book?`,
		comments: []seedComment{
			{"finn_in_berlin", "This is exactly what I'm doing with German right now and the answer is no, it never feels like the same book. Sometimes it feels like a cousin of the book."},
			{"saltwater_pages", "Thank you for posting! That cover is gorgeous. Wish I could read Finnish."},
			{"tuomas_reads", "Saw this at Akateeminen and almost bought it purely for the cover. Now I have an actual reason."},
			{"bluewing_books", `I read Pedro Páramo in Spanish after loving it in English and it was a completely different experience. The English felt like a ghost story; the Spanish felt like a song. Neither was wrong.

I think translation is a kind of reading that someone else does for you, and then you read their reading.`},
			{"kirjatoukka_kaisa", "@bluewing_books that's a beautiful way to put it. Stealing that."},
		},
	},
	{
		user:       "tuomas_reads",
		title:      "Alvar Aalto Library in Vyborg: Saving a Modern Masterpiece",
		image:      seedCovers + "alvar-aalto-library-vyborg.jpg",
		daysAgo:    49,
		categories: []string{"Non-fiction", "Architecture & Design", "Book Reviews"},
		body: `A bit off the usual path for this forum, but I suspect some of you will appreciate a book about a library.

The Viipuri (Vyborg) library was finished in 1935, survived the war on the "wrong" side of the new border, was neglected for decades and then restored by a joint Finnish-Russian committee over roughly twenty years. This book documents the whole restoration: the undulating wooden ceiling in the lecture hall, the round skylights, the sunken reading room.

It's heavy on photographs and technical drawings, so not a cover-to-cover read, but the chapters about tracking down original materials and craft methods are fascinating.

The cover photo of the entrance canopy alone makes me want to visit.`,
		comments: []seedComment{
			{"lens_and_ink", "I photographed the reading room a few years ago. The light from those skylights is unreal, it's like being inside a lantern. Great pick."},
			{"loud_librarian", "As a librarian I am obligated to say: more posts about libraries please."},
			{"old_man_oskar", "My father visited Viipuri before the war and talked about this building for the rest of his life. Thank you for this."},
			{"tuomas_reads", "@old_man_oskar that's wonderful. Did he ever describe it? I'd love to hear."},
			{"old_man_oskar", "He said it was the first building where he understood that a ceiling could be beautiful. He was a carpenter. The wooden waves in the lecture hall stayed with him."},
			{"paperbag_critic", "A coffee table book about a building is not a book. Fight me."},
			{"tuomas_reads", "@paperbag_critic it has more text than half the novels you've reviewed here."},
			{"loud_librarian", "Okay okay, let's keep it friendly :)"},
		},
	},
	{
		user:       "theo_logian",
		title:      "Maan päällä ja taivaassa: Bible stories that actually respect kids",
		image:      seedCovers + "maan-paalla-ja-taivaassa.jpg",
		daysAgo:    46,
		categories: []string{"Religion & Spirituality", "Children's Books", "Finnish Literature", "Mauri Vihko"},
		body: `Mauri Vihko's "Maan päällä ja taivaassa – Juttuja Raamatusta" (On Earth and in Heaven – Stories from the Bible).

I picked this up for my younger cousin and ended up reading half of it myself. Most children's Bible adaptations either sand every story down to "be nice" or go so heavy on moralising that kids switch off. This one tells the stories plainly, keeps the strange and difficult bits in, and leaves room for questions.

The language is clear modern Finnish, chapters are short enough for bedtime, and the tone is warm without being preachy.

Worth a look whether or not you're religious, if you want kids to know the stories behind half of Western art and literature.`,
		comments: []seedComment{
			{"kirjatoukka_kaisa", "My grandmother had a really old fashioned illustrated Bible storybook and I was terrified of the flood picture for years. A gentler version sounds good."},
			{"cozy_ella", "As a teacher I agree with the point about cultural literacy. Kids encounter these stories in paintings, films and other books constantly."},
			{"paperbag_critic", "Is it possible to recommend religious books to kids without it being indoctrination though? Genuine question."},
			{"theo_logian", `@paperbag_critic genuine answer: I think it depends on whether the book tells the child what to believe or tells them what the story says. This one is much more the second.

A kid who knows the story of Jonah can still decide they don't believe it. A kid who has never heard of it just misses the reference in Pinocchio.`},
			{"mum_reads_aloud", "Saving this. My eldest has started asking about church after a school visit and I have no idea where to begin."},
		},
	},
	{
		user:       "paperbag_critic",
		title:      "ARC review: Yesteryear by Caro Claire Burke (no major spoilers)",
		image:      seedCovers + "yesteryear.jpg",
		daysAgo:    43,
		categories: []string{"Book Reviews", "Upcoming Releases", "Satire", "Mystery & Thriller", "Caro Claire Burke"},
		body: `Got an advance copy (hence "FINAL COVER TO COME" in giant letters on the front, very glamorous).

The pitch: a tradwife influencer whose whole brand is the old-fashioned homestead life suddenly finds herself actually living it, with none of the ring lights.

What works: the satire is sharp. The first act skewers influencer culture so precisely it made me wince at my own screen time. What doesn't: the middle section drags and the tone wobbles between dark comedy and genuine horror without always landing either.

3.5 stars, rounding down to 3 because I am who I am.`,
		comments: []seedComment{
			{"starry_sofi", "Three stars from you is basically five from anyone else 😂 adding it."},
			{"rosie_rides_west", "The premise is so good. The yellow cover is very loud though."},
			{"paperbag_critic", "@rosie_rides_west the yellow is temporary. The publisher literally says so on the cover."},
			{"quiet_chapter", "How did you get an ARC? Asking for a me."},
			{"paperbag_critic", "@quiet_chapter years of writing mean but honest reviews. Apparently publishers like that. Who knew."},
			{"maya_reads", "This might be a great book club pick when it comes out. Lots to argue about."},
		},
	},
	{
		user:       "theo_logian",
		title:      "Conversations with God, Book 3: profound or just very confident?",
		image:      seedCovers + "conversations-with-god-3.jpg",
		daysAgo:    40,
		categories: []string{"Religion & Spirituality", "Non-fiction", "Unpopular Opinions", "Neale Donald Walsch"},
		body: `A friend lent me the whole trilogy and I just finished the third one. I'm a theology student so I went in very sceptical.

For anyone who doesn't know: Walsch writes questions to God and then writes down the answers he says he received. Book 3 goes the furthest: highly evolved beings, the nature of time, death, the soul.

Some passages genuinely moved me, especially the parts about fear versus love as the root of every choice. Others read like a self-help seminar with a very impressive guest speaker.

I'm curious how others read these. As spiritual texts? As philosophy? As fiction?`,
		comments: []seedComment{
			{"paperbag_critic", "As fiction. Next question."},
			{"lurker_lena", "This book helped me through a really bad year after my dad died. I don't care how it was written, it made me less afraid."},
			{"paperbag_critic", "@lurker_lena I'm sorry about your dad, really. But a book helping you doesn't make its claims true. Lots of things are comforting."},
			{"lurker_lena", "@paperbag_critic I didn't say it was true. I said it helped. Those are allowed to be different things."},
			{"theo_logian", `I think this exchange is actually the whole question in miniature.

Religious texts have always done two jobs at once: making claims about reality and helping people live. You can reject the first and still recognise the second. I'm not convinced by Walsch's method, but I can't dismiss what @lurker_lena describes either. That's also how a lot of people relate to the Psalms.`},
			{"noir_nights", "The 'highly evolved beings' chapter is where it lost me. That's sci-fi with fewer spaceships."},
			{"silver_sage", "Read it the way you'd read Kahlil Gibran: as wisdom literature with a strong authorial voice. It holds up better that way than as revelation."},
			{"paperbag_critic", "Fair. @lurker_lena that came out harsher than I meant. Glad it helped you."},
			{"lurker_lena", "Thanks. Appreciated."},
		},
	},
	{
		user:       "maya_reads",
		title:      "Hello Beautiful made me ugly cry on a plane",
		image:      seedCovers + "hello-beautiful.jpg",
		daysAgo:    38,
		categories: []string{"Book Reviews", "Literary Fiction", "Ann Napolitano"},
		body: `Ann Napolitano's Hello Beautiful. William Waters grows up in a house where nobody looks at him, then finds himself pulled into the loud, loving Padavano family with four sisters, and marries one of them.

It's a quiet homage to Little Women and once you notice it you can't stop seeing it. Julia, Sylvie, Cecelia and Emeline map loosely onto the March sisters but each of them becomes her own person.

The middle third is the heaviest portrayal of depression I've read in mainstream fiction, and it's handled with real care. Then the last section... I'll just say I had to apologise to the man in seat 14B.

Has anyone read Dear Edward too? Is it as good?`,
		comments: []seedComment{
			{"bluewing_books", "Sylvie's chapters. Oh my heart."},
			{"paperbag_critic", "I found it sentimental. Every character is either a saint or tragically misunderstood, and the Little Women thing felt like a crutch."},
			{"maya_reads", "@paperbag_critic sentimental isn't automatically bad! Some books are allowed to want you to feel things."},
			{"paperbag_critic", "@maya_reads wanting me to feel things and engineering a feeling are different. The last chapter felt engineered."},
			{"rosie_rides_west", "Respectfully, it's a book about a family. What did you expect, a heist?"},
			{"quiet_chapter", "Dear Edward is just as sad but in a different way. Bring tissues for that one too."},
			{"austen_annotator", `I'm on the fence here. I loved William's storyline and how carefully the book treats mental illness. But I agree a little with @paperbag_critic that Julia gets flattened into the "difficult sister" for too long. She deserved more interiority earlier on.

Still a 4 for me.`},
			{"maya_reads", "@austen_annotator ok that's a fair point about Julia. I'll allow it."},
		},
	},
	{
		user:       "wattpad_wren",
		title:      "Let's talk honestly about After by Anna Todd",
		image:      seedCovers + "after.jpg",
		daysAgo:    36,
		categories: []string{"Romance", "Young Adult", "Unpopular Opinions", "Adaptations", "Anna Todd"},
		body: `I read After on Wattpad in 2014 when it was still a One Direction fanfic and Hardin was literally Harry. It was the first time I saw a story written on a phone, by someone like me, become a real book and then films.

Rereading it now as an adult, I see everything people criticise: Hardin is controlling, Tessa forgives way too much, the relationship is a red flag parade.

But I also think people forget how important it was. It proved that fanfiction readers are a massive audience and that publishing had been ignoring them.

Can a book be both important and kind of bad?`,
		comments: []seedComment{
			{"starry_sofi", "Yes!! I loved it at 15 and I'd be horrified if my little sister dated a Hardin. Both things are true."},
			{"austen_annotator", "I've confiscated this from more than one student, so I have opinions. The problem isn't that it's messy; it's that the book frames possessive behaviour as proof of love and never really questions it."},
			{"wattpad_wren", "@austen_annotator but plenty of classics do that too. Heathcliff digs up Cathy's grave and he's on the syllabus."},
			{"austen_annotator", "@wattpad_wren Wuthering Heights knows its characters are monstrous. That's the difference. After wants you to swoon."},
			{"rosie_rides_west", `I'm going to defend it a bit. Romance readers are not stupid. We know the difference between a fantasy and a relationship advice manual. Teenagers reading about a toxic boyfriend isn't the same as teenagers wanting one.

The panic about "what girls read" is very old and very boring.`},
			{"austen_annotator", "@rosie_rides_west I'm not panicking, I'm teaching critical reading. There's a difference too."},
			{"paperbag_critic", "I read the whole series in one hate-read weekend. Important? Maybe. Good? No. Entertaining? ...unfortunately yes."},
			{"wattpad_wren", "@paperbag_critic the hate-read is a form of love and you know it."},
			{"velvet_vellum", "Honestly this thread is better than the book."},
		},
	},
	{
		user:       "hiding_behind_pages",
		title:      "Katabasis: two PhD students walk into Hell to save their advisor",
		image:      seedCovers + "katabasis.jpg",
		daysAgo:    34,
		categories: []string{"Book Reviews", "Fantasy", "Dark Academia", "R.F. Kuang"},
		body: `R.F. Kuang's Katabasis. Alice Law and her rival Peter Murdoch are Cambridge grad students in magick, and when their advisor Professor Grimes dies in a lab accident (partly Alice's fault), they decide to go down to Hell and bring him back. Because he's the only one who can sign off their theses.

That's it. That's the premise. And it is so, so academic in the best possible way. Hell is structured like a university. There are logic puzzles, paradoxes, Dante and Orpheus references everywhere.

It's also the most honest book I've read about what academia does to people. Alice's relationship to her own ambition hurt to read.

Not for everyone: it's dense, and if you hate footnote-y digressions you'll suffer. For me: 5/5, already rereading.`,
		comments: []seedComment{
			{"velvet_vellum", "That cover is a work of art. The staircases! The cat!"},
			{"hiding_behind_pages", "@velvet_vellum the cat is plot relevant, which is all I will say."},
			{"dusk_reader", "Is it like The Secret History? I want more books like that."},
			{"hiding_behind_pages", "@dusk_reader same mood, totally different plot. Less murder, more existential dread."},
			{"paperbag_critic", "Kuang always writes the same protagonist: brilliant, miserable, self-destructive. Babel, Yellowface, now this. I'm tired."},
			{"hiding_behind_pages", "@paperbag_critic she writes about ambition. Of course the protagonists are ambitious. Should they be lazy and cheerful?"},
			{"moonlit_myths", "The way she uses the Greek idea of katabasis, the hero's descent to the underworld, and then twists it is so smart. Orpheus is basically the cautionary tale the whole book argues with."},
			{"skyline_reader", "Read it on the commute. Missed my stop twice. Strong recommend, weak recommend for your punctuality."},
			{"maya_reads", "Adding to the book club shortlist!"},
		},
	},
	{
		user:       "dusk_reader",
		title:      "Finished Crime and Punishment at 3am and I need to talk to someone",
		image:      seedCovers + "crime-and-punishment.jpg",
		daysAgo:    32,
		categories: []string{"Classics", "Translated Fiction", "Book Reviews", "Fyodor Dostoevsky"},
		body: `I just finished it. The Macmillan Collector's Library edition, tiny print, gold edges, I've been carrying it everywhere for a month.

Raskolnikov spends 500 pages trying to convince himself that some people are allowed to step over the moral line for a greater purpose, and the entire book is his body and mind refusing to agree with him.

Porfiry Petrovich is the best detective in literature and he barely does any detecting. He just... waits. Talks. Lets Raskolnikov collapse.

And Sonya. The scene where she reads him the story of Lazarus. I had to put the book down and stare at the wall.

19 years old and I think I just had a religious experience in my bedroom. Please tell me this is normal.`,
		comments: []seedComment{
			{"dostoevsky_defender", "Completely normal. Which translation though? This matters more than you think."},
			{"dusk_reader", "@dostoevsky_defender the Collector's Library one uses Constance Garnett I think?"},
			{"dostoevsky_defender", `@dusk_reader ah. Garnett. Okay. It's not your fault.

Garnett translated dozens of Russian novels at incredible speed and smoothed everything into the same polite Victorian English. Dostoevsky's prose is feverish, repetitive, ugly on purpose. Pevear & Volokhonsky keep the roughness. Read it again in P&V in a few years and it'll feel like a different, angrier book.`},
			{"silver_sage", "I'd gently push back. P&V are accurate but sometimes so literal that the English becomes stilted. Garnett has carried millions of readers to Dostoevsky. Including, apparently, this one at 3am."},
			{"dostoevsky_defender", "@silver_sage 'stilted' is a feature when the original is stilted! You can't sand down the author to make him comfortable."},
			{"silver_sage", "@dostoevsky_defender and you can't hold a 19-year-old's first great reading experience hostage over translation politics."},
			{"dusk_reader", "I'm just happy you're all as intense about this as I am."},
			{"old_man_oskar", "Read it at 17 in 1969. Reread it every decade. Each time Raskolnikov is younger and Porfiry is wiser. Welcome to the club, young one."},
			{"noir_nights", "Porfiry walked so Columbo could run."},
		},
	},
	{
		user:       "finn_in_berlin",
		title:      "Learning German with 'Was die Gezeiten verbergen' (Middletide)",
		image:      seedCovers + "was-die-gezeiten-verbergen.jpg",
		daysAgo:    30,
		categories: []string{"Mystery & Thriller", "Translated Fiction", "Reading Habits", "Sarah Crouch"},
		body: `This is Sarah Crouch's Middletide in German translation. I picked it at the Hugendubel near my flat because the cover is a sunset over water and I am weak.

My German is B1 on a good day. Rules I set myself:
- No dictionary for the first read of each chapter, just guessing
- Then a second pass with a dictionary for words I saw at least three times
- Absolutely no switching to English

It's a slow-burn mystery set in a small coastal town in the Pacific Northwest: a man comes home after failing as a writer, a woman is found dead, and the town has secrets. Very atmospheric. The tide and the mud flats are basically a character.

Three weeks in, 140 pages. Any tips from people who learned a language through novels?`,
		comments: []seedComment{
			{"kirjatoukka_kaisa", "Mahtavaa! I learned English through Harry Potter, reading each book in Finnish first and then English. Familiar plot helps a lot."},
			{"lens_and_ink", "Your three-sightings rule is smart. Otherwise you spend more time in the dictionary than in the book."},
			{"bluewing_books", "Read the English original after! It's fascinating to see what the translator did."},
			{"finn_in_berlin", "@bluewing_books planning to. Mostly to find out if I understood the ending correctly 😅"},
			{"quiet_chapter", "Also try audiobooks in German while following along in the text. Hearing it helps the pronunciation stick."},
			{"noir_nights", "Is it actually a good mystery or just a good mood? Asking as a mystery snob."},
			{"finn_in_berlin", "@noir_nights so far more mood than puzzle. Ask me again in 200 pages."},
		},
	},
	{
		user:       "old_man_oskar",
		title:      "My 60th annual reading of A Christmas Carol",
		image:      seedCovers + "a-christmas-carol.jpg",
		daysAgo:    28,
		categories: []string{"Classics", "Reading Habits", "Charles Dickens"},
		body: `I first read A Christmas Carol at 14. My teacher gave me this exact kind of cheap little paperback. Every December since, I have read it again, the last few years aloud to my grandchildren.

This year I start early because my eyes are slower than they used to be.

Things that change as you get older:
At 14, I loved the ghosts.
At 30, I was Bob Cratchit, tired and underpaid.
At 50, I recognised a bit of Scrooge in myself and didn't like it.
At 74, I mostly think about the Ghost of Christmas Yet to Come, and what people will say about me.

It's a short book. It's the perfect book. Dickens wrote it in six weeks to pay his debts. Not every great thing needs to take a long time.`,
		comments: []seedComment{
			{"loud_librarian", "Oskar this is the most beautiful post on the forum. Thank you."},
			{"mum_reads_aloud", "I'm stealing this tradition for my kids. Maybe the Muppets version first though."},
			{"old_man_oskar", "@mum_reads_aloud the Muppet one is very faithful to the book! Michael Caine is an excellent Scrooge. Good choice."},
			{"dusk_reader", "Your list of ages made me tear up. I'm going to start my own list this year."},
			{"paperbag_critic", "I've always found it a bit saccharine, honestly. Tiny Tim exists purely to make you cry."},
			{"silver_sage", "@paperbag_critic Tiny Tim is a blunt instrument, but Dickens was writing against real child poverty and wanted his readers angry. The sentiment is a weapon, not decoration."},
			{"maya_reads", "60 years! You're the reason I love this forum."},
			{"old_man_oskar", "You are all very kind. Now go read it, it only takes an evening."},
		},
	},
	{
		user:       "abbey_road_reader",
		title:      "John & Paul: A Love Story in Songs is the best Beatles book in years",
		image:      seedCovers + "john-and-paul.jpg",
		daysAgo:    26,
		categories: []string{"Biography & Memoir", "Music", "Non-fiction", "Book Reviews", "Ian Leslie"},
		body: `I've read probably forty Beatles books. Most of them are either gossip or discography. Ian Leslie's John & Paul is neither: it's the story of one friendship, told through the songs they wrote to and about each other.

Leslie's argument is that we've been sold a cartoon: John the edgy genius, Paul the cute craftsman. Instead he shows two people who were constantly in conversation through their music, even after the split. "How Do You Sleep?", "Dear Friend", "Here Today". It's devastating when you hear them as replies.

Every chapter is anchored on a song, which makes it a perfect book to read with headphones on.

Five stars, four copies of Revolver, zero regrets.`,
		comments: []seedComment{
			{"skyline_reader", "Listening to Here Today right after reading the last chapter was a mistake. On a train. In public."},
			{"paperbag_critic", "A 'love story'? It's a business partnership between two men who wrote some songs. The title is clickbait."},
			{"abbey_road_reader", "@paperbag_critic Leslie addresses that literally in chapter one. It's not about romance; it's about love between friends, which we oddly don't have good language for."},
			{"paperbag_critic", "@abbey_road_reader fine, but then call it that. Words matter."},
			{"silver_sage", "Some of the most important literature ever written is about intense friendship. Montaigne on La Boétie, Tennyson's In Memoriam. Love is the correct word."},
			{"old_man_oskar", "I saw them play in Stockholm in 1963. I was 11 and it was the loudest thing I've ever heard. Buying this book today."},
			{"abbey_road_reader", "@old_man_oskar ...you saw them LIVE?? I need you to write a post about this immediately."},
		},
	},
	{
		user:       "cozy_ella",
		title:      "We read all 15 Oz books aloud in class this year",
		image:      seedCovers + "oz-books.jpg",
		daysAgo:    25,
		categories: []string{"Children's Books", "Classics", "Fantasy", "L. Frank Baum"},
		body: `Year 3 project: reading the entire Oz series by L. Frank Baum aloud, one chapter at the end of each day. We used this all-in-one edition with the original illustrations.

Most people only know The Wonderful Wizard of Oz (and mostly through the film). The sequels are genuinely bonkers: Ozma, Tik-Tok the mechanical man, the Patchwork Girl, the Nome King, a whole book where Dorothy's chickens become important.

My class's ranking:
1. Ozma of Oz (Billina the hen is their hero)
2. The Patchwork Girl of Oz
3. The Wonderful Wizard of Oz
15. The Road to Oz ("nothing happens, miss")

The kids now wear paper crowns to reading time. I have no regrets.`,
		comments: []seedComment{
			{"mum_reads_aloud", "Please be my children's teacher."},
			{"moonlit_myths", "Ozma of Oz is genuinely one of the great children's fantasy books. The Wheelers terrified me as a kid."},
			{"cozy_ella", "@moonlit_myths the Wheelers had half my class hiding under their desks. In the best way."},
			{"old_man_oskar", "My grandmother read me these in Swedish in the 1950s. I still remember the Patchwork Girl. Wonderful that children still love them."},
			{"austen_annotator", "Worth knowing that some of the later books have outdated and unpleasant passages. Did you edit as you read?"},
			{"cozy_ella", "@austen_annotator yes, a few. And we talked about why when it came up. They handled it really thoughtfully."},
			{"paperbag_critic", "I can't believe The Road to Oz came last. Okay no, I can. 'Nothing happens, miss' is the most accurate review on this site."},
		},
	},
	{
		user:       "bluewing_books",
		title:      "The Emperor of Gladness: Ocean Vuong's novel about care",
		image:      seedCovers + "the-emperor-of-gladness.jpg",
		daysAgo:    23,
		categories: []string{"Literary Fiction", "Book Reviews", "Ocean Vuong"},
		body: `Hai is nineteen, standing on a bridge in the rain in a town called East Gladness, Connecticut, about to jump. An old Lithuanian widow named Grazina, who has dementia, shouts at him from across the river and talks him down. He ends up becoming her caretaker.

That's the start. What follows is one of the warmest books about poverty I've read. Hai works at a fast-casual chain restaurant and the coworker chapters are so funny and so tired and so real.

Vuong is a poet, and you can tell, but this is much more of a novel than On Earth We're Briefly Gorgeous. Fewer fireworks, more people.

The relationship between Hai and Grazina is the heart of it. Two people the world has quietly thrown away, taking care of each other.

5/5. Read it slowly.`,
		comments: []seedComment{
			{"maya_reads", "Grazina is one of my favourite characters of the year. The scenes where she thinks it's still the war... ugh."},
			{"paperbag_critic", "Beautiful sentences, yes. But nearly 400 pages and I couldn't tell you what happened in the middle third. It's vibes."},
			{"bluewing_books", "@paperbag_critic the middle third is the restaurant! It's about work and routine and how care happens inside boredom. The 'nothing happening' is the point."},
			{"paperbag_critic", "@bluewing_books 'the boredom is the point' is the most literary fiction sentence ever written."},
			{"saltwater_pages", "Honestly it's also a good one to read in chunks. Sentence by sentence it's stunning."},
			{"lurker_lena", "Is it very sad? I need to know before I start."},
			{"bluewing_books", "@lurker_lena sad in places but ultimately hopeful. Much gentler than the opening suggests."},
		},
	},
	{
		user:       "meadow_marginalia",
		title:      "\"Whatever our souls are made of, his and mine are the same\" — romantic or terrifying?",
		image:      seedCovers + "wuthering-heights-art.jpg",
		daysAgo:    21,
		categories: []string{"Quotes", "Classics", "Romance", "Unpopular Opinions", "Emily Brontë"},
		body: `Found this illustration of Catherine Earnshaw on an art print site and it's been my phone background for a week.

The line is from Wuthering Heights, Cathy speaking to Nelly about Heathcliff: "Whatever our souls are made of, his and mine are the same."

It's on wedding invitations, tattoos, Pinterest boards everywhere. But in context she says it right before explaining why she'll marry Edgar instead, and the rest of the book is two people destroying everyone around them, including each other's children.

So: is this the most romantic sentence in English literature, or a warning label?`,
		comments: []seedComment{
			{"velvet_vellum", "Warning label. Absolutely. People who put this on their wedding invitations have not finished the book."},
			{"rosie_rides_west", "It can be both! A love can be overwhelming and destructive. That's what makes it romantic in the old sense, not the Hallmark sense."},
			{"austen_annotator", "This is exactly what I meant in the After thread, @wattpad_wren. Brontë shows the damage. Heathcliff's cruelty to Isabella isn't framed as passion."},
			{"wattpad_wren", "@austen_annotator and yet here we are, with the quote on a pretty painting and teenagers swooning over Heathcliff for 175 years. The reader decides what's romantic, not the author."},
			{"silver_sage", `Remember that Emily Brontë was writing about something closer to spiritual identity than romance. Cathy doesn't say "I love him"; she says "I am Heathcliff." It's an annihilation of self. That's more mysticism than courtship.

Whether that's beautiful or horrifying is, I think, the question the book refuses to answer for you.`},
			{"meadow_marginalia", "@silver_sage 'I am Heathcliff' is so much creepier when you put it that way. I love it."},
			{"paperbag_critic", "Hot take: nobody in Wuthering Heights is likeable and that's why it's great."},
			{"velvet_vellum", "@paperbag_critic for once we fully agree."},
			{"starry_sofi", "The illustration is gorgeous though, where's it from?"},
			{"meadow_marginalia", "@starry_sofi an independent illustrator, I'll DM you the link!"},
		},
	},
	{
		user:       "austen_annotator",
		title:      "Sense and Sensibility: Team Elinor or Team Marianne?",
		image:      seedCovers + "sense-and-sensibility.jpg",
		daysAgo:    20,
		categories: []string{"Classics", "Romance", "Editions & Collecting", "Jane Austen"},
		body: `Treated myself to this floral edition with the sprayed edges for my annual Austen reread. It's even prettier in person.

Every time I teach Sense and Sensibility the class splits:
Team Elinor: restraint, responsibility, quietly carrying the whole family while her heart breaks.
Team Marianne: honesty, passion, refusing to pretend.

Austen seems to side with Elinor... but then she gives Marianne the most growth. And Colonel Brandon, who I will defend forever.

Where do you stand? And which adaptation is the best? (The answer is 1995, but I'll hear you out.)`,
		comments: []seedComment{
			{"meadow_marginalia", "Elinor. Forever. The scene where she finally learns Edward isn't married and runs out of the room crying is one of the best moments in all of Austen."},
			{"starry_sofi", "Marianne!! She's 17 and feels everything and I relate too much."},
			{"rosie_rides_west", "Colonel Brandon deserved a better proposal scene. We barely see him win her over."},
			{"austen_annotator", "@rosie_rides_west that's the most common complaint, and it's fair. Austen hurries it in a single paragraph."},
			{"paperbag_critic", "Edward Ferrars is the most boring man in English literature and Elinor could do better."},
			{"meadow_marginalia", "@paperbag_critic Edward is boring ON PURPOSE. He's the anti-Willoughby! Boring is safe!"},
			{"paperbag_critic", "@meadow_marginalia 'boring is safe' is not the romance you think it is."},
			{"velvet_vellum", "1995. Emma Thompson wrote that script. There is no debate."},
			{"old_man_oskar", "My late wife was a Marianne and I was an Elinor. It worked for 48 years."},
			{"maya_reads", "@old_man_oskar okay that's the best answer in this thread."},
		},
	},
	{
		user:       "lurker_lena",
		title:      "Help me find a book: the library catalogue only shows a blank cover",
		image:      seedCovers + "no-cover.jpg",
		daysAgo:    19,
		categories: []string{"Help Me Find a Book", "Children's Books", "General Discussion"},
		body: `Bit of a weird one. I'm trying to find a children's book I read around 2005. The library database found a match for my vague search, but it only shows this blank placeholder cover and no description, so I can't tell if it's the right one.

What I remember:
- A girl who finds a door at the back of a wardrobe-ish cupboard in her grandmother's house
- It leads to a world where it's always autumn
- There's a talking fox who lies a lot
- Blue cover with silver lettering (I think)

It is NOT Narnia, I checked. Any ideas?`,
		comments: []seedComment{
			{"loud_librarian", "Ooh I love these. Do you remember if it was a series? And roughly how long? Picture book or chapter book?"},
			{"lurker_lena", "@loud_librarian chapter book, maybe 200 pages. I think there was a sequel but I never read it."},
			{"cozy_ella", "A lying fox and eternal autumn... it sounds a bit like the Coraline vibe but cosier. Was there anything scary?"},
			{"lurker_lena", "@cozy_ella a little! The grandmother's house had a clock that ran backwards."},
			{"moonlit_myths", "The trickster fox makes me think the author was drawing on European folklore. Reynard the Fox maybe? Could be a translated book."},
			{"loud_librarian", "I'm going to check our children's catalogue tomorrow with those keywords. Give me a day!"},
			{"lurker_lena", "This is why I finally made an account. You're all amazing."},
		},
	},
	{
		user:       "horror_hank",
		title:      "Cujo: does the dog die? (spoilers, obviously)",
		image:      seedCovers + "cujo.jpg",
		daysAgo:    18,
		categories: []string{"Horror", "Book Reviews", "Stephen King"},
		body: `SPOILERS BELOW.

Yes, the dog dies. I knew that going in. What I didn't know was that King would also kill Tad, the four year old, which the film famously changed.

Cujo is a St. Bernard who chases a rabbit into a cave, gets bitten by rabid bats, and slowly turns into a monster through no fault of his own. King even gives us chapters from Cujo's point of view, confused and sick, still trying to be a Good Dog.

Then Donna and Tad are trapped in a broken Pinto in the summer heat for days. It's almost unbearably tense.

King has said he barely remembers writing it because of his drinking at the time. You can feel the desperation on the page.

4/5. Not his best, but maybe his cruellest.`,
		comments: []seedComment{
			{"noir_nights", "Cujo's POV chapters are what make this book. You're terrified of him and heartbroken for him at the same time."},
			{"cozy_ella", "I knew I shouldn't have clicked. Why did I click."},
			{"horror_hank", "@cozy_ella the title literally asks if the dog dies, this one's on you 😅"},
			{"paperbag_critic", "The ending is just cruel for the sake of cruelty. King killing a child to make a point about fate feels cheap."},
			{"noir_nights", `@paperbag_critic I strongly disagree. The whole book is built on the idea that the universe doesn't care. Every bad thing in Cujo happens because of tiny accidents: a rabbit, a bat, a broken car part, a mechanic who's away.

If Tad survives, the book is lying to you. The film made it comforting and that's why the film is forgettable.`},
			{"horror_hank", "@noir_nights well said. It's the most nihilistic thing he wrote in that era."},
			{"paperbag_critic", "@noir_nights an argument I can respect, even if I still hated it."},
			{"quiet_chapter", "Read this once at 16 and still check under the car before I get in. Thanks King."},
		},
	},
	{
		user:       "bluewing_books",
		title:      "Learning Dutch, preordered a novel, the shop shows 'Omslag (nog) niet beschikbaar'",
		image:      seedCovers + "omslag-niet-beschikbaar.jpg",
		daysAgo:    17,
		categories: []string{"Upcoming Releases", "Translated Fiction", "Reading Habits"},
		body: `"Cover (not yet) available." That's all I have, which is somehow very exciting.

I'm six months into Dutch lessons and decided my reward for passing A2 is preordering a real adult novel, a Dutch translation of a book I already know well in English, so I have a fighting chance.

Inspired partly by @finn_in_berlin doing the same with German.

Question for multilingual people: did reading a book you already knew help, or did it make you lazy and just fill in from memory?`,
		comments: []seedComment{
			{"finn_in_berlin", "Honoured to be an inspiration 😄 For me knowing the plot helps a lot in the beginning, but you have to be strict about not skimming."},
			{"kirjatoukka_kaisa", "It helped me. You know who everyone is, so your brain has energy for the grammar."},
			{"lens_and_ink", "Children's books are great as a step in between. Simple sentences, real language."},
			{"paperbag_critic", "The blank cover posts are becoming a genre on this forum."},
			{"loud_librarian", "@paperbag_critic honestly, I support it. Mystery covers are fun."},
			{"bluewing_books", "Update when it arrives with the real cover, promise!"},
		},
	},
	{
		user:       "meadow_marginalia",
		title:      "Northanger Abbey is Austen's funniest book and nobody talks about it",
		image:      seedCovers + "northanger-abbey.jpg",
		daysAgo:    15,
		categories: []string{"Classics", "Satire", "Unpopular Opinions", "Jane Austen"},
		body: `Catherine Morland is seventeen and has read far too many gothic novels. She visits an old abbey and immediately decides her host must have murdered his wife. He hasn't. (Well. He's still awful, just not in that way.)

Austen is openly making fun of the gothic craze and of Catherine, but she's also defending novel-reading in one of the best passages she ever wrote, the bit where she breaks the fourth wall to scold authors who sneer at novels.

And Henry Tilney is the funniest hero Austen ever gave us. He teases Catherine constantly and it's genuinely charming, not mean.

It was her first finished novel, published after her death, and it gets ignored because it's short and silly. I think it's her most fun.`,
		comments: []seedComment{
			{"velvet_vellum", "As a gothic novel enjoyer, being lovingly roasted by Austen is an honour."},
			{"austen_annotator", "Henry Tilney's speech about muslin! The man knows fabric! Underrated king."},
			{"silver_sage", "It's also a sharp book about money, actually. General Tilney's interest in Catherine is entirely about her supposed fortune. The comedy hides a pretty cynical engine."},
			{"paperbag_critic", "It's fine. It's a parody of books nobody reads anymore, so half the jokes need footnotes."},
			{"meadow_marginalia", "@paperbag_critic The Mysteries of Udolpho is still in print and I've read it. Twice. The jokes land."},
			{"paperbag_critic", "@meadow_marginalia you're the exception that proves my rule."},
			{"starry_sofi", "Catherine Morland is literally a BookTok girlie who reads too many dark romances. Austen predicted us."},
		},
	},
	{
		user:       "old_man_oskar",
		title:      "Oliver Twist: rereading a book I first read when it was a mere 124 years old",
		image:      seedCovers + "oliver-twist.jpg",
		daysAgo:    14,
		categories: []string{"Classics", "Historical Fiction", "Editions & Collecting", "Charles Dickens"},
		body: `My granddaughter gave me this little Macmillan Collector's edition for my birthday. The book is as small as my hand and I've needed a magnifying glass, but what a joy.

First read it in 1962 at a school library in Turku. What I remember from then: the workhouse, "Please, sir, I want some more," and being scared of Bill Sikes.

What I noticed now: Oliver is almost a bystander in his own story. The real heart of the book is Nancy. She's the one who makes the hardest choice, and she pays for it.

Also: Dickens is furious. The early chapters about the workhouse are a political pamphlet with a plot attached. And it worked; people were outraged.

One uncomfortable thing: the portrayal of Fagin is clearly antisemitic, and Dickens himself later revised some of it after a reader wrote to him. Worth knowing going in.`,
		comments: []seedComment{
			{"austen_annotator", "Thank you for flagging the Fagin issue. That correspondence with Eliza Davis is a great teaching moment about authors listening to readers."},
			{"dusk_reader", "Nancy deserved so much better. Sikes is one of the scariest villains I've read."},
			{"velvet_vellum", "Those Collector's Library editions are dangerously collectible. I have eleven."},
			{"old_man_oskar", "@velvet_vellum my granddaughter says I'm now required to collect them too. Apparently it's 'a whole thing'."},
			{"mum_reads_aloud", "The Artful Dodger is the reason my son's school project was about pickpockets. The teacher was confused."},
			{"silver_sage", "Lovely post, Oskar. Dickens is a novelist you need to reread at different ages; you've proved it twice now on this forum."},
		},
	},
	{
		user:       "moonlit_myths",
		title:      "English Fairy Tales: the original versions are wonderfully weird",
		image:      seedCovers + "english-fairy-tales.jpg",
		daysAgo:    13,
		categories: []string{"Fairy Tales & Folklore", "Classics", "Short Stories", "Children's Books"},
		body: `This Wordsworth Classics edition collects Flora Annie Steel's retellings with the Arthur Rackham illustrations. The cover alone, those swirling fairy women, is pure Rackham.

Things I learned:
- Jack and the Beanstalk has several versions and in some of them Jack is just a thief with no moral justification
- Tom Tit Tot is the English Rumpelstiltskin and honestly better
- "Mr Fox" is a full-on Bluebeard-style horror story about a bride discovering her fiancé's chamber of murdered women
- Molly Whuppie outsmarts a giant three times and is the best fairy tale heroine I've ever met

These are not cosy stories. They're violent, funny, unfair, and much stranger than the versions we grew up with.

Anyone else collect fairy tale anthologies?`,
		comments: []seedComment{
			{"velvet_vellum", "Mr Fox is SO creepy. 'Be bold, be bold, but not too bold.' Chills."},
			{"cozy_ella", "I read Molly Whuppie to my class and now half of them want to be giant-tricksters when they grow up."},
			{"horror_hank", "Fairy tales are just horror for children and I've been saying it for years."},
			{"kirjatoukka_kaisa", "Finnish folk tales are similarly brutal. Collected by Eero Salmelainen in the 1800s, and very dark."},
			{"moonlit_myths", "@kirjatoukka_kaisa ooh, is there an English translation? I need it."},
			{"kirjatoukka_kaisa", "@moonlit_myths partly! Some are in collections of Finnish folk tales in English. I'll dig up titles."},
			{"paperbag_critic", "Rackham's illustrations are the best part of every book they're in. Fight me on that one too."},
			{"moonlit_myths", "@paperbag_critic no fight. That's just true."},
		},
	},
	{
		user:       "moonlit_myths",
		title:      "The Witcher reading order debate (again): start with Blood of Elves or The Last Wish?",
		image:      seedCovers + "the-witcher-blood-of-elves.jpg",
		daysAgo:    12,
		categories: []string{"Fantasy", "Translated Fiction", "Adaptations", "Andrzej Sapkowski"},
		body: `A friend just started The Witcher with Blood of Elves because it says "the first novel" and the cover is gorgeous (that stone rose!). She's confused about who Yennefer is, why Geralt has a child surprise thing going on, and who Dandelion is.

This is why you start with the short story collections: The Last Wish, then Sword of Destiny. The novels assume you've read them.

Blood of Elves itself is wonderful once you have that context: Ciri training at Kaer Morhen, the political chess game, Triss and Yennefer teaching her magic. It's slower and more political than people expect after the games.

Settle it once and for all: short stories first, yes?`,
		comments: []seedComment{
			{"hiding_behind_pages", "Short stories first. This is not up for debate."},
			{"skyline_reader", "I started with Blood of Elves and was fine. Confused but fine. Then went back to the stories and they hit harder knowing what's coming."},
			{"moonlit_myths", "@skyline_reader that's... actually a valid way to do it. I hate that you have a point."},
			{"noir_nights", "Came in from the games, read in publication order, never looked back. Sapkowski's dialogue is so much funnier than people expect."},
			{"paperbag_critic", "The books are slow, the politics are confusing and Geralt spends half of them brooding. The games did it better."},
			{"moonlit_myths", "@paperbag_critic the games are literally sequels to the books. They couldn't exist without them."},
			{"paperbag_critic", "@moonlit_myths and sequels can be better than the original. See: The Empire Strikes Back."},
			{"hiding_behind_pages", "Comparing a 500 page novel to a 150 hour game is an apples to... very large oranges situation."},
			{"moonlit_myths", `For the record, here's what I tell people:
1. The Last Wish
2. Sword of Destiny
3. Blood of Elves
4. Time of Contempt
5. Baptism of Fire
6. The Tower of the Swallow
7. The Lady of the Lake
8. Season of Storms (whenever, it's a side story)

Thread closed. (It is not closed.)`},
		},
	},
	{
		user:       "mum_reads_aloud",
		title:      "All Fall Down: we have read this 400 times and my baby still laughs",
		image:      seedCovers + "all-fall-down.jpg",
		daysAgo:    11,
		categories: []string{"Children's Books", "Reading Habits", "Helen Oxenbury"},
		body: `Helen Oxenbury's All Fall Down. Four babies. They play, they sing, they bounce, they all fall down. That's the entire book.

My 10-month-old has heard this probably 400 times and every single time we get to "all fall down" she laughs like it's the funniest thing that's ever happened.

What makes Oxenbury so good is how she draws babies. They're chubby and wobbly and completely real, diverse without making a big deal out of it, and they look like they're having the time of their lives.

Board book parents: what's the one your kid demands every night?`,
		comments: []seedComment{
			{"cozy_ella", "Oxenbury is a legend. We're Going on a Bear Hunt with Michael Rosen is in every classroom I've worked in."},
			{"old_man_oskar", "With my grandchildren it was The Very Hungry Caterpillar. Every night for a year. I can still recite it."},
			{"mum_reads_aloud", "@old_man_oskar the caterpillar is on our shelf too. The holes in the pages have been 'tested' by small fingers a lot."},
			{"starry_sofi", "This is the most wholesome post on the forum and I needed it today."},
			{"paperbag_critic", "Four babies fall down. Five stars. No notes."},
			{"maya_reads", "@paperbag_critic is this the first five star review you've ever given??"},
			{"paperbag_critic", "@maya_reads babies are an easy audience. The book respects its reader."},
		},
	},
	{
		user:       "kirjatoukka_kaisa",
		title:      "Erottua: ten writers, ten short stories, one theme",
		image:      seedCovers + "erottua.jpg",
		daysAgo:    9,
		categories: []string{"Short Stories", "Finnish Literature", "Book Reviews", "Elina Koskinen"},
		body: `"Erottua – novelleja kymmeneltä kirjoittajalta" (roughly: "Standing out / Having separated – short stories by ten writers"), edited by Elina Koskinen.

The title is a nice Finnish double meaning: "erottua" means to stand out, but it sounds like separating or parting. The stories play with both, from breakups to people who don't fit in.

As with any anthology, uneven, but three of the ten stories really stuck with me, and it's a great way to discover new Finnish voices. Smaller publishers like this are where a lot of the interesting work is happening right now.

The cover looks like an infrared photo of a garden, which suits the "seeing something differently" theme.

Does anyone else read anthologies cover to cover, or do you skip around?`,
		comments: []seedComment{
			{"finn_in_berlin", "Skip around, always. I read the shortest one first as a taste test."},
			{"tuomas_reads", "Bought this at a book fair because of the cover. Will report back."},
			{"bluewing_books", "Cover to cover! The editor put them in that order for a reason. It's like an album."},
			{"kirjatoukka_kaisa", "@bluewing_books that's how I feel too. The ordering here definitely matters, the last story echoes the first."},
			{"lurker_lena", "Short stories are perfect for my attention span this year. Wish I could read Finnish."},
		},
	},
	{
		user:       "lens_and_ink",
		title:      "The Mires by Tina Makereti: a small town and a swamp full of secrets",
		image:      seedCovers + "the-mires.jpg",
		daysAgo:    7,
		categories: []string{"Literary Fiction", "Book Reviews", "Translated Fiction", "Tina Makereti"},
		body: `I shot a project in Aotearoa New Zealand last year and asked a bookseller in Wellington for something local and recent. She handed me Tina Makereti's The Mires without a word.

Set in a small town beside a wetland, told through several neighbours who don't really know each other: a Māori family with deep roots in the land, a refugee mother and her son who've just arrived, and a man slowly being pulled into online rage.

The swamp itself, the mire, is the soul of the book. Old, patient, holding everything that's been thrown into it.

It's about belonging and who gets to belong. Some te reo Māori phrases are woven in without translation, which I loved. It made me slow down and look things up, like a guest should.

That cover with the kingfisher is also stunning.`,
		comments: []seedComment{
			{"bluewing_books", "This sounds like exactly my kind of book. Adding immediately."},
			{"austen_annotator", "Small correction for the categories: it's written in English, so not translated fiction, but I understand the impulse with the te reo parts!"},
			{"lens_and_ink", "@austen_annotator ha, fair. I tagged it that way because so much of it felt like moving between languages. I'll leave it as a happy accident."},
			{"paperbag_critic", "Multi-POV community novels always have one character who's clearly the author's favourite and one who's just there to be wrong. Is that the case here?"},
			{"lens_and_ink", "@paperbag_critic surprisingly not. Even the man going down the rabbit hole is written with real compassion, which is the brave part of the book."},
			{"saltwater_pages", "The wetland as a character is my weakness. Sold."},
		},
	},
	{
		user:       "velvet_vellum",
		title:      "Show me your favourite map endpapers (mine: Mississippi river)",
		image:      seedCovers + "mississippi-map.jpg",
		daysAgo:    5,
		categories: []string{"Editions & Collecting", "Classics", "Mark Twain"},
		body: `Finally got a special edition of Adventures of Huckleberry Finn and the endpapers are this gorgeous blue map of the Mississippi with little illustrations of steamboats and towns. Illinois, Missouri, Arkansas, all the way down.

I love a book with a map inside. It makes the book feel like a place you can go to.

Huck and Jim's raft trip down the river is one of the great journeys in literature, and seeing it laid out like this made me want to reread it with a pencil and trace their route.

Please share your best book maps! Tolkien doesn't count, everyone knows Tolkien.`,
		comments: []seedComment{
			{"moonlit_myths", "Earthsea! Le Guin's map of the archipelago is the most beautiful fantasy map ever. (Not Tolkien, so it counts.)"},
			{"cozy_ella", "The Hundred Acre Wood map from Winnie-the-Pooh. Nothing will ever beat it."},
			{"old_man_oskar", "Moominvalley in Tove Jansson's books. My children wore out our copy studying it."},
			{"austen_annotator", "Worth noting with Huck Finn that it's a book many schools now teach with a lot of context because of its language. Still one of the most important American novels."},
			{"velvet_vellum", "@austen_annotator definitely. Twain was satirising racism, but the book uses the language of its time. It needs that context."},
			{"skyline_reader", "Treasure Island. The original map with X marks the spot. The book was literally born from a map Stevenson drew with his stepson."},
			{"velvet_vellum", "@skyline_reader I didn't know that! That's the best origin story."},
		},
	},
}
