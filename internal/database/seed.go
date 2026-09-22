package database

import (
	"database/sql"
	"fmt"
)

func Seed(db *sql.DB) error {
	steps := []struct {
		name string
		fn   func(*sql.DB) error
	}{
		{"users", insertDataUsers},
		{"categories", insertDataCategories},
		{"posts", insertDataPosts},
		{"category_post", insertCategoryPostConnections},
		{"comments", insertDataComments},
		{"post votes", insertDataPostVotes},
		{"comment votes", insertDataCommentVotes},
	}
	for _, s := range steps {
		if err := s.fn(db); err != nil {
			return fmt.Errorf("seeding %s: %w", s.name, err)
		}
	}
	return nil
}

func insertDataUsers(db *sql.DB) error {
	sql := `INSERT OR IGNORE INTO user (username, email, password_hash) VALUES
            ('bookworm_bella', 'bella@lions.test', '$2a$10$fLFb.iqCyA31Gk0zbKeP8.rNcTHY1lXYSfesw38kL3TXLOzGI19w.'),
            ('inkwell_ivan',   'ivan@lions.test',  '$2a$10$fLFb.iqCyA31Gk0zbKeP8.rNcTHY1lXYSfesw38kL3TXLOzGI19w.'),
            ('proust_fan',     'proust@lions.test','$2a$10$fLFb.iqCyA31Gk0zbKeP8.rNcTHY1lXYSfesw38kL3TXLOzGI19w.'),
            ('marginalia_mo',  'mo@lions.test',    '$2a$10$fLFb.iqCyA31Gk0zbKeP8.rNcTHY1lXYSfesw38kL3TXLOzGI19w.'),
            ('dogeared_dana',  'dana@lions.test',  '$2a$10$fLFb.iqCyA31Gk0zbKeP8.rNcTHY1lXYSfesw38kL3TXLOzGI19w.'),
            ('sci_fi_sam',     'sam@lions.test',   '$2a$10$fLFb.iqCyA31Gk0zbKeP8.rNcTHY1lXYSfesw38kL3TXLOzGI19w.'),
            ('quiet_quill',    'quill@lions.test', '$2a$10$fLFb.iqCyA31Gk0zbKeP8.rNcTHY1lXYSfesw38kL3TXLOzGI19w.'),
            ('leo_the_lion',   'leo@lions.test',   '$2a$10$fLFb.iqCyA31Gk0zbKeP8.rNcTHY1lXYSfesw38kL3TXLOzGI19w.');`
	_, err := db.Exec(sql)
	return err
}

func insertDataCategories(db *sql.DB) error {
	sql := `INSERT OR IGNORE INTO category (name, type) VALUES
            ('General Discussion', 'genres'),
            ('Book Reviews',       'genres'),
            ('Classics',           'genres'),
            ('Science Fiction',    'genres'),
            ('Fantasy',            'genres'),
            ('Mystery & Thriller', 'genres'),
            ('Character Studies',  'genres'),
            ('Author Interviews',  'genres'),
            ('Recommendations',    'genres'),
            ('Marcel Proust',      'authors'),
            ('Ursula K. Le Guin',  'authors'),
            ('Emily Bronte',       'authors'),
            ('Agatha Christie',    'authors'),
            ('J.R.R. Tolkien',     'authors'),
            ('Kazuo Ishiguro',     'authors'),
            ('Gillian Flynn',      'authors'),
            ('Andy Weir',          'authors'),
            ('Neal Stephenson',    'authors'),
            ('George Eliot',       'authors'),
            ('Brandon Sanderson',  'authors'),
            ('N.K. Jemisin',       'authors'),
            ('Susanna Clarke',     'authors');`
	_, err := db.Exec(sql)
	return err
}

func insertDataPosts(db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM post").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	sql := `INSERT INTO post (user_id, title, body, created_at) VALUES
            (8, 'Welcome to the Literary Lions forum!',
                'No more sticky notes falling out of books. Introduce yourself here and tell us what you are reading this month.',
                datetime('now', '-30 days')),
            (3, 'Is ''In Search of Lost Time'' worth the commitment?',
                'I am 200 pages into Swann''s Way and loving the prose, but seven volumes is a lot. Has anyone here finished the whole thing?',
                datetime('now', '-28 days')),
            (6, 'Review: The Left Hand of Darkness',
                'Le Guin''s world-building is subtle but relentless. The ice crossing is one of the best sequences I have read in any genre. 5/5.',
                datetime('now', '-25 days')),
            (4, 'Heathcliff: villain or victim?',
                'Every time I reread Wuthering Heights my sympathy for Heathcliff shifts. Where do you all land?',
                datetime('now', '-22 days')),
            (5, 'Cosy mysteries for a rainy weekend',
                'Looking for something light with a clever puzzle. I have already devoured all of Agatha Christie''s Poirot books.',
                datetime('now', '-20 days')),
            (1, 'Notes from our chat with a debut novelist',
                'At last week''s meetup we interviewed a local debut author about drafting, rejection letters and finding an agent.',
                datetime('now', '-17 days')),
            (2, 'Why does everyone skip the Silmarillion?',
                'It reads like mythology rather than a novel, and I think that is exactly its strength. Change my mind.',
                datetime('now', '-14 days')),
            (7, 'Unreliable narrators done right',
                'Remains of the Day, Gone Girl, The Murder of Roger Ackroyd... which unreliable narrator fooled you the hardest?',
                datetime('now', '-11 days')),
            (6, 'Hard sci-fi starter pack?',
                'A friend wants to get into hard science fiction. I was thinking The Martian, then Project Hail Mary, then Seveneves. Thoughts?',
                datetime('now', '-8 days')),
            (3, 'Review: Middlemarch',
                'George Eliot understood people better than most psychologists. Dorothea''s arc is quietly devastating. 4.5/5.',
                datetime('now', '-5 days')),
            (5, 'Best fantasy series that is actually finished',
                'I am tired of waiting for sequels. Recommend me a completed fantasy series, please!',
                datetime('now', '-3 days')),
            (4, 'Do you write in your books?',
                'I annotate everything in pencil. My partner thinks it is sacrilege. Where does the club stand?',
                datetime('now', '-1 days'));`
	_, err = db.Exec(sql)
	return err
}

func insertCategoryPostConnections(db *sql.DB) error {
	// genres: 1-9, authors: 10-22
	sql := `INSERT OR IGNORE INTO category_post (post_id, category_id) VALUES
            (1, 1), (1, 22),
            (2, 3), (2, 1), (2, 10),
            (3, 2), (3, 4), (3, 11),
            (4, 7), (4, 3), (4, 12),
            (5, 6), (5, 9), (5, 13),
            (6, 8),
            (7, 5), (7, 3), (7, 14),
            (8, 7), (8, 6), (8, 15), (8, 16), (8, 13),
            (9, 4), (9, 9), (9, 17), (9, 18),
            (10, 2), (10, 3), (10, 19),
            (11, 5), (11, 9), (11, 20), (11, 21), (11, 11),
            (12, 1);`
	_, err := db.Exec(sql)
	return err
}

func insertDataComments(db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM comment").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	sql := `INSERT INTO comment (user_id, post_id, body, created_at) VALUES
            (1, 1, 'Hi all! Currently reading Piranesi and loving it.',                 datetime('now', '-30 days', '+2 hours')),
            (6, 1, 'Finally, a place where my notes won''t get lost.',                  datetime('now', '-30 days', '+5 hours')),
            (4, 1, 'Great idea, Leo. Rereading Jane Eyre this month.',                  datetime('now', '-29 days')),
            (2, 2, 'The second half of Swann''s Way is where it really picks up.',      datetime('now', '-28 days', '+3 hours')),
            (7, 2, 'I finished all seven volumes last year. Worth every page.',         datetime('now', '-27 days')),
            (1, 3, 'Adding this to my to-read list right now.',                         datetime('now', '-25 days', '+4 hours')),
            (2, 3, 'Genly Ai is such a great outsider narrator.',                       datetime('now', '-24 days')),
            (1, 4, 'Heathcliff is both, and that is what makes him interesting.',      datetime('now', '-22 days', '+1 hours')),
            (7, 4, 'Victim of circumstance, villain by choice.',                        datetime('now', '-22 days', '+6 hours')),
            (3, 4, 'Hmm, I am not convinced he deserves any sympathy at all.',          datetime('now', '-21 days')),
            (7, 5, 'Try Richard Osman''s Thursday Murder Club.',                        datetime('now', '-20 days', '+2 hours')),
            (1, 5, 'Dorothy L. Sayers! Start with Strong Poison.',                      datetime('now', '-19 days')),
            (4, 6, 'Great write-up, thanks for sharing.',                               datetime('now', '-17 days', '+3 hours')),
            (5, 7, 'I tried twice and gave up both times, sorry Ivan.',                 datetime('now', '-14 days', '+2 hours')),
            (3, 7, 'Have you tried the audiobook? The narrator is fantastic.',          datetime('now', '-13 days')),
            (3, 8, 'Stevens in Remains of the Day broke my heart.',                     datetime('now', '-11 days', '+1 hours')),
            (5, 8, 'Roger Ackroyd, no contest. I threw the book across the room.',      datetime('now', '-11 days', '+4 hours')),
            (2, 9, 'Solid list. Maybe add The Three-Body Problem after Seveneves.',      datetime('now', '-8 days', '+2 hours')),
            (7, 10, 'We should make this next month''s group read.',                   datetime('now', '-5 days', '+3 hours')),
            (6, 11, 'Mistborn and The Earthsea Cycle are both finished and excellent.', datetime('now', '-3 days', '+1 hours')),
            (2, 11, 'The Broken Earth trilogy, hands down.',                            datetime('now', '-2 days')),
            (1, 12, 'Pencil only. Pen in books is a crime.',                            datetime('now', '-1 days', '+2 hours')),
            (6, 12, 'Sticky tabs for me, the paper-pandemonium habit dies hard.',       datetime('now', '-1 days', '+5 hours'));`
	_, err = db.Exec(sql)
	return err
}

func insertDataPostVotes(db *sql.DB) error {
	sql := `INSERT OR IGNORE INTO post_vote (user_id, post_id, vote) VALUES
            (1, 1, 1), (2, 1, 1), (3, 1, 1), (4, 1, 1), (6, 1, 1),
            (1, 2, 1), (7, 2, 1), (5, 2, -1),
            (1, 3, 1), (2, 3, 1), (4, 3, 1), (8, 3, 1),
            (1, 4, 1), (3, 4, -1), (7, 4, 1),
            (7, 5, 1), (1, 5, 1),
            (4, 6, 1), (8, 6, 1), (2, 6, 1),
            (3, 7, 1), (5, 7, -1), (6, 7, -1),
            (3, 8, 1), (5, 8, 1), (1, 8, 1), (2, 8, 1),
            (2, 9, 1), (8, 9, 1),
            (7, 10, 1), (4, 10, 1), (1, 10, 1),
            (6, 11, 1), (2, 11, 1), (3, 11, -1),
            (1, 12, 1), (6, 12, 1), (3, 12, -1), (8, 12, 1);`
	_, err := db.Exec(sql)
	return err
}

func insertDataCommentVotes(db *sql.DB) error {
	sql := `INSERT OR IGNORE INTO comment_vote (user_id, comment_id, vote) VALUES
            (8, 1, 1), (4, 1, 1),
            (8, 2, 1),
            (3, 4, 1), (7, 4, 1),
            (3, 5, 1), (1, 5, 1), (2, 5, 1),
            (6, 7, 1),
            (4, 8, 1), (7, 8, 1), (3, 8, -1),
            (1, 10, -1),
            (5, 11, 1), (5, 12, 1),
            (2, 14, -1),
            (7, 17, 1), (3, 17, 1),
            (6, 18, 1),
            (5, 20, 1), (2, 20, 1),
            (4, 22, 1), (3, 22, -1);`
	_, err := db.Exec(sql)
	return err
}
