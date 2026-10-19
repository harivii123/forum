package database

// posts without an image
var seedTextPosts = []seedPost{
	{
		user:       "loud_librarian",
		title:      "Welcome to Literary Lions! Introduce yourself here",
		daysAgo:    60,
		categories: []string{"General Discussion"},
		body: `Hello and welcome, fellow lions!

This is the place to say hi. Tell us:
- What you're reading right now
- Your all-time favourite book
- One book you think is overrated (be brave)

A few house rules: argue about books as much as you like, but be kind to the people. Spoiler warnings in titles, please. And if you can't find a category for your post, message me and I'll consider adding one.

I'll start: reading Piranesi for the third time, favourite is Earthsea, overrated is... The Alchemist. There, I said it.`,
		comments: []seedComment{
			{"maya_reads", "Hi all! Reading Hello Beautiful, favourite is Little Women, overrated: anything with 'girl' in the title from 2014-2018."},
			{"paperbag_critic", "Reading everything. Favourite: undecided after 30 years. Overrated: most of what you all love. Nice to meet you."},
			{"loud_librarian", "@paperbag_critic every forum needs one. Welcome 😄"},
			{"old_man_oskar", "Good evening. I'm Oskar, 74, from Turku. Reading Dickens again. Favourite: The Brothers Karamazov. Overrated: Ulysses, and I'm too old to pretend otherwise."},
			{"dusk_reader", "@old_man_oskar the Ulysses opinion from someone who's actually read it carries weight."},
			{"starry_sofi", "Hiii ✨ Sofi here. Currently: Fourth Wing reread. Favourite: Howl's Moving Castle. Overrated: I can't say it, people will hunt me down."},
			{"theo_logian", "Theo, theology student. Reading Conversations with God for... research. Favourite: Gilead by Marilynne Robinson. Overrated: The Da Vinci Code, for professional reasons."},
			{"lurker_lena", "Hi. Been lurking for a while. Reading cosy mysteries. Favourite is Anne of Green Gables. Not brave enough for the third one yet."},
			{"kirjatoukka_kaisa", "Moi! Kaisa from Tampere. Reading Anne Carson. Favourite: Muumipappa ja meri (Moominpappa at Sea). Overrated: Seitsemän veljestä, sorry Finland."},
			{"tuomas_reads", "@kirjatoukka_kaisa a Finn calling Seitsemän veljestä overrated in public. Respect."},
		},
	},
	{
		user:       "paperbag_critic",
		title:      "Unpopular opinion: you do NOT have to finish every book you start",
		daysAgo:    48,
		categories: []string{"Unpopular Opinions", "Reading Habits"},
		body: `Saw someone in another forum proudly say they've "never DNF'd a book in their life" like it's a moral achievement.

Life is short. There are more good books than you will ever have time to read. If a book isn't working for you after 50-100 pages, put it down. Not finishing is not failing. You're not getting graded.

I DNF about 30% of what I start. Fight me.`,
		comments: []seedComment{
			{"silver_sage", "Nancy Pearl's rule of 50: give a book 50 pages. If you're over 50 years old, subtract your age from 100. At my age I can quit after 25 pages with a clear conscience."},
			{"dostoevsky_defender", "Strongly disagree. Some of the greatest books are difficult for the first hundred pages. If I'd quit The Brothers Karamazov at page 80 I would have missed the best book I've ever read."},
			{"paperbag_critic", "@dostoevsky_defender and if you'd quit 20 mediocre books at page 80, you'd have had time for 5 more Karamazovs."},
			{"dostoevsky_defender", "@paperbag_critic there aren't 5 more Karamazovs. That's my point."},
			{"maya_reads", "I used to force myself to finish everything and it put me in a two-year reading slump. DNFing saved my reading life."},
			{"meadow_marginalia", `I'm somewhere in the middle. I quit books that bore me, but not books that challenge me. Those are different feelings and it took me years to tell them apart.

Boredom feels like nothing's happening. Challenge feels like something's happening that I don't understand yet.`},
			{"paperbag_critic", "@meadow_marginalia okay, that distinction is genuinely useful. Annoyingly wise."},
			{"quiet_chapter", "I DNF and then feel guilty for months and sometimes go back. Is that a third category?"},
			{"loud_librarian", "@quiet_chapter that's the 'paused' shelf. It's a real thing. Libraries approve."},
			{"old_man_oskar", "At 74 I finish only what I love. Wish I'd started earlier."},
		},
	},
	{
		user:       "maya_reads",
		title:      "What are you reading this week?",
		daysAgo:    4,
		categories: []string{"General Discussion", "Reading Habits"},
		body: `Weekly check-in! What's on your nightstand, in your bag, in your ears?

Me: halfway through Katabasis because @hiding_behind_pages wouldn't stop talking about it, and an audiobook of Pride and Prejudice for walks.`,
		comments: []seedComment{
			{"hiding_behind_pages", "My work here is done. Rereading The Secret History in the meantime."},
			{"skyline_reader", "Klara and the Sun by Ishiguro. On the train, on the platform, in the queue for coffee."},
			{"horror_hank", "The Shining. Again. It's October, it's tradition."},
			{"cozy_ella", "Matilda with my class. The trunchbull chapters are going down VERY well."},
			{"finn_in_berlin", "Still Middletide in German. Page 210. I now know the German for 'mudflat' (Watt) and will never forget it."},
			{"abbey_road_reader", "Paul McCartney's The Lyrics. Two volumes, I'm treating it like a devotional."},
			{"wattpad_wren", "Three different fanfics and one actual published novel, I'm a well-rounded person."},
			{"austen_annotator", "Marking 64 essays on Macbeth. Does that count?"},
			{"maya_reads", "@austen_annotator it counts and you deserve a medal."},
		},
	},
	{
		user:       "quiet_chapter",
		title:      "Do audiobooks count as reading? (please be nice)",
		daysAgo:    44,
		categories: []string{"Reading Habits", "General Discussion", "Unpopular Opinions"},
		body: `Someone at my book club laughed when I said I "read" a book I'd actually listened to on my walks. It stung more than I want to admit.

I listen to about 40 books a year and read maybe 15 on paper. I remember the audiobooks just as well, sometimes better because a good narrator makes the characters stick.

So, honestly: does it count?`,
		comments: []seedComment{
			{"loud_librarian", "It counts. Libraries lend audiobooks. Blind readers read audiobooks. Oral storytelling is literally older than writing. Your book club friend is wrong."},
			{"paperbag_critic", `I'm going to be the annoying one. It counts as experiencing the book. It doesn't count as reading, because reading is a specific activity with your eyes and your own pace.

Not better or worse, just different. Why do people get so offended about the word?`},
			{"quiet_chapter", "@paperbag_critic because when people say 'that's not reading' they usually mean 'that's not as good as reading', even if you don't."},
			{"paperbag_critic", "@quiet_chapter fair. That's on the people, not the word, but fair."},
			{"dostoevsky_defender", "With dense books I think something gets lost. You can't easily go back three sentences, slow down, underline. For Crime and Punishment I'd want paper."},
			{"mum_reads_aloud", "Audiobooks are the only reason I've finished any adult book since my second child was born. They count. They count SO much."},
			{"silver_sage", "Homer was performed aloud for centuries before anyone wrote him down. I believe the Iliad still counts."},
			{"skyline_reader", "I do both. Paper for the commute, audio for running. My brain doesn't care which, it still hands me the story."},
			{"lurker_lena", "This thread made me feel better about my 'reading' too. Thank you."},
		},
	},
	{
		user:       "lens_and_ink",
		title:      "Recommend me Finnish literature (in English translation)",
		daysAgo:    35,
		categories: []string{"Recommendations", "Finnish Literature", "Translated Fiction", "Tove Jansson"},
		body: `Heading to Helsinki and Lapland in winter for a shoot. Ritual: read the country's literature before I go.

All I've read so far is the Kalevala (selections) and the Moomins as a kid. Help me, Finnish lions! Novels, short stories, anything translated.`,
		comments: []seedComment{
			{"kirjatoukka_kaisa", `Here's my starter pack:
- Tove Jansson, The Summer Book. Short, perfect, a grandmother and granddaughter on an island. Technically Finland-Swedish but very Finnish in feeling.
- Väinö Linna, The Unknown Soldier. THE Finnish war novel.
- Sofi Oksanen, Purge. Dark, Estonia and Finland, unforgettable.
- Rosa Liksom, Compartment No. 6. A train across the Soviet Union. Became a film too.
- Arto Paasilinna, The Year of the Hare. Funny and strange, very Finnish humour.`},
			{"tuomas_reads", "Add Aleksis Kivi's Seven Brothers despite what Kaisa said in the welcome thread 😄 It's the foundational Finnish novel."},
			{"kirjatoukka_kaisa", "@tuomas_reads foundational, yes. Fun to read, no. I stand by it."},
			{"old_man_oskar", "The Summer Book. Tove Jansson understood children and old people better than anyone. I've given it as a gift twenty times."},
			{"finn_in_berlin", "Monika Fagerholm's The American Girl if you want something weird and dreamy."},
			{"lens_and_ink", "This is amazing. Taking the whole list. Starting with The Summer Book on the flight."},
			{"silver_sage", "The Summer Book is a minor miracle. You'll understand Finnish silence much better after reading it."},
		},
	},
	{
		user:       "maya_reads",
		title:      "Book club vote for November: Katabasis vs. Crime and Punishment",
		daysAgo:    10,
		categories: []string{"Book Club", "Fantasy", "Classics", "R.F. Kuang", "Fyodor Dostoevsky"},
		body: `Both books have had big threads here recently and both have devoted fans, so let's settle it with a vote.

Option A: Katabasis by R.F. Kuang. New, dark academia, Hell as a university.
Option B: Crime and Punishment by Dostoevsky. A classic, a murder, a fever dream.

Fun fact: they actually pair well. Both are about clever young people who think their intelligence lets them bend the rules.

Comment A or B! Discussion starts the first Sunday of November.`,
		comments: []seedComment{
			{"hiding_behind_pages", "A. Obviously A."},
			{"dusk_reader", "B!!! Please. I need people to talk about Sonya with."},
			{"dostoevsky_defender", "B, and I'll host a translation comparison session for free."},
			{"paperbag_critic", "@dostoevsky_defender that is a threat, not an offer."},
			{"velvet_vellum", "A. It's November, we need some fresh dread."},
			{"silver_sage", "B. But the pairing Maya points out is excellent. Alice Law and Raskolnikov would hate each other."},
			{"skyline_reader", "A. Crime and Punishment needs more than a month for most people."},
			{"dusk_reader", "@skyline_reader I read it in a month!"},
			{"skyline_reader", "@dusk_reader you are 19 and have no job. Respectfully."},
			{"maya_reads", "Current count: A 4, B 4. This is going to come down to the wire 😬"},
		},
	},
	{
		user:       "bluewing_books",
		title:      "Books that made you cry in public",
		daysAgo:    29,
		categories: []string{"General Discussion", "Recommendations", "Kazuo Ishiguro"},
		body: `Mine: Never Let Me Go by Kazuo Ishiguro, on a bus, at the part with the cassette tape. A stranger handed me a tissue without saying a word. Ishiguro would have been proud.

Where did you break down, and which book did it?`,
		comments: []seedComment{
			{"maya_reads", "Hello Beautiful, plane, seat 14C. The man in 14B will never recover."},
			{"skyline_reader", "The Remains of the Day on a train. The last chapter on the pier. Ishiguro gets me every time too."},
			{"horror_hank", "Pet Sematary in a library. Not because it was scary. Because of the cat. And the son. Mostly the son."},
			{"cozy_ella", "Charlotte's Web, in front of 28 eight-year-olds. They were very kind about it."},
			{"old_man_oskar", "A Grief Observed by C.S. Lewis, in a hospital waiting room, the year my wife was ill. I don't mind admitting it."},
			{"bluewing_books", "@old_man_oskar sending you so much warmth."},
			{"paperbag_critic", "I don't cry at books. ...Okay, The Iron Giant. That's a film. Leave me alone."},
			{"starry_sofi", "A Little Life on the metro and people looked SO concerned."},
			{"silver_sage", "@starry_sofi that book should come with a public health warning."},
		},
	},
	{
		user:       "meadow_marginalia",
		title:      "Annotating systems: show me yours",
		daysAgo:    24,
		categories: []string{"Reading Habits", "Editions & Collecting"},
		body: `My system:
- Pencil underline = good sentence
- Pencil bracket in margin = important for the plot or a theme
- Tiny star = I want to quote this
- Question mark = I disagree or don't understand
- Small tabs: yellow for character, green for setting, pink for "OH NO"

People who write in pen in books: I respect you, but I fear you.`,
		comments: []seedComment{
			{"starry_sofi", "Glitter gel pens and a full colour-coded tab rainbow. My books look like a craft store exploded."},
			{"austen_annotator", "@starry_sofi I've seen your Instagram. They're actually beautiful. I'm still shaken by the gel pens."},
			{"old_man_oskar", "I write the date and where I was when I finished, on the last page. 60 years of my life is inside my books."},
			{"dostoevsky_defender", "Pen. Blue. In the margins. Arguing with the translator."},
			{"meadow_marginalia", "@dostoevsky_defender of course you do."},
			{"lurker_lena", "I have never written in a book. It feels wrong. But these replies are making me want to try."},
			{"paperbag_critic", "I write 'no.' next to sentences I disagree with. Some books have 'no.' on every page."},
			{"velvet_vellum", "Library books: nothing, I am not a monster. Own books: chaos."},
		},
	},
	{
		user:       "silver_sage",
		title:      "The best opening lines in literature",
		daysAgo:    16,
		categories: []string{"Quotes", "Classics", "General Discussion"},
		body: `A good first line is a contract with the reader. Let's collect the best.

I'll start with one that never gets old:
"It was a bright cold day in April, and the clocks were striking thirteen." — Orwell, Nineteen Eighty-Four.

One sentence, and you already know something is wrong with the world.`,
		comments: []seedComment{
			{"austen_annotator", "\"It is a truth universally acknowledged, that a single man in possession of a good fortune, must be in want of a wife.\" Irony in the very first sentence."},
			{"dusk_reader", "\"Happy families are all alike; every unhappy family is unhappy in its own way.\" Tolstoy wins."},
			{"dostoevsky_defender", "@dusk_reader the Russians are dominating as usual."},
			{"moonlit_myths", "\"In a hole in the ground there lived a hobbit.\" Simple. Perfect. Tolkien wrote it on the back of an exam he was marking."},
			{"horror_hank", "\"The terror, which would not end for another twenty-eight years — if it ever did end — began, so far as I know or can tell, with a boat made from a sheet of newspaper floating down a gutter swollen with rain.\" It. Long but iconic."},
			{"paperbag_critic", "\"Call me Ishmael.\" Overrated. Three words, everyone loses their minds. The rest of Moby-Dick is about rope."},
			{"silver_sage", "@paperbag_critic the rope is a metaphor. The whale is a metaphor. Everything is a metaphor, dear."},
			{"saltwater_pages", "\"Mrs Dalloway said she would buy the flowers herself.\" So quiet, so much in it."},
			{"skyline_reader", "\"The past is a foreign country; they do things differently there.\" The Go-Between. Better than most whole novels."},
		},
	},
	{
		user:       "lurker_lena",
		title:      "I've been in a reading slump for three months. Help?",
		daysAgo:    2,
		categories: []string{"Reading Habits", "Recommendations", "Agatha Christie"},
		body: `I used to read every night. Since summer I start books, read five pages, then pick up my phone. Nothing grabs me. I feel like I've forgotten how to concentrate.

What actually helped you get out of a slump? Specific books welcome, but tips too.`,
		comments: []seedComment{
			{"loud_librarian", "Go short and go fun. Novellas, short story collections, graphic novels. Rebuild the habit before you rebuild the attention span."},
			{"maya_reads", "Reread a childhood favourite! Comfort and zero risk. Anne of Green Gables is right there in your intro post 😉"},
			{"quiet_chapter", "Switch format. Audiobooks got me out of my last slump; my eyes were just tired of screens."},
			{"paperbag_critic", "Put your phone in another room. Unglamorous, but it's the actual answer."},
			{"lurker_lena", "@paperbag_critic ...okay that one hurt because it's true."},
			{"noir_nights", `Read something with a hook so strong you can't put it down. A thriller with short chapters. Agatha Christie is perfect: And Then There Were None is short, every chapter ends on a cliffhanger, and you'll want to know who did it.

Slumps are often about books that ask too much of you at the wrong time. Pick one that gives instead.`},
			{"silver_sage", "Be patient with yourself. Reading comes in seasons. It always comes back."},
			{"lurker_lena", "Thank you all. Started And Then There Were None last night. Read 60 pages. 60!!"},
			{"noir_nights", "@lurker_lena told you. Christie never misses."},
		},
	},
}
